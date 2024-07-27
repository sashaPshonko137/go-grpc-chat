package pg

import (
	"database/sql"
	"fmt"
	"time"

	desc "chat/pkg/chat_v1"

	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Storage struct {
	db *sql.DB
}

func New(dbPath string) (*Storage, error) {
	const op = "storage.pg.New"
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("%s: ping failed: %w", op, err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: table creation failed: %w", op, err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_users_id ON users(id);`)
	if err != nil {
	  return nil, fmt.Errorf("%s: index creation failed: %w", op, err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chats (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: table creation failed: %w", op, err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_chats_id ON chats(id);`)
	if err != nil {
	  return nil, fmt.Errorf("%s: index creation failed: %w", op, err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			user_id INT,
			chat_id INT,
			content TEXT NOT NULL,
			created_at TIMESTAMP
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: table creation failed: %w", op, err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_messages_id ON messages(id);`)
	if err != nil {
	  return nil, fmt.Errorf("%s: user creation failed: %w", op, err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chat_users (
			id SERIAL PRIMARY KEY,
			user_id INT,
			chat_id INT
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: table creation failed: %w", op, err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_chat_users_id ON chat_users(id);`)
	if err != nil {
	  return nil, fmt.Errorf("%s: chat_user creation failed: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(name string) error {
	const op = "storage.pg.Users.Create"
	_, err := s.db.Exec(`INSERT INTO users (name) VALUES ($1)`, name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) CreateChat(name string) (int32, error) {
	const op = "storage.pg.Chats.Create"
	
	stmt, err := s.db.Prepare(`INSERT INTO chats (name) VALUES ($1) RETURNING id`)
	if err != nil {
		return 0, fmt.Errorf("%s: prepare statement failed: %w", op, err)
	}
	defer stmt.Close()

	var id int32
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: execution failed: %w", op, err)
	}
	return id, nil
}

func (s *Storage) CreateMessage(user_id, chat_id int32, content string, created_at *timestamppb.Timestamp) error {
	const op = "storage.pg.Messages.Create"
	_, err := s.db.Exec(`INSERT INTO messages (user_id, chat_id, content, created_at) VALUES ($1, $2, $3, $4)`, user_id, chat_id, content, created_at.AsTime())
	if err != nil {
		return fmt.Errorf("%s: message creation failed: %w", op, err)
	}

	return nil
}

func (s *Storage) CreateChatUser(chat_id, user_id int32) error {
	const op = "storage.pg.ChatUsers.Create"
	_, err := s.db.Exec(`INSERT INTO chat_users (chat_id, user_id) VALUES ($1, $2)`, chat_id, user_id)
	if err != nil {
		return fmt.Errorf("%s: prepare statement failed: %w", op, err)
	}

	return nil
}

func (s *Storage) GetUser(user_id int32) (string, error) {
	const op = "storage.pg.Users.Get"
	var name string
	err := s.db.QueryRow(`SELECT name FROM users WHERE id = $1`, user_id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return name, nil
}


func (s *Storage) GetChat(chat_id int32) (string, error) {
	const op = "storage.pg.Chats.Get"
	var name string
	err := s.db.QueryRow(`SELECT name FROM chats WHERE id = $1`, chat_id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return name, nil
}

func (s *Storage) GetUsersFromChat(chat_id int32) ([]int32, error) {
	const op = "storage.pg.Users.GetUsersFromChat"
	var users []int32

	rows, err := s.db.Query(`SELECT user_id FROM chat_users WHERE chat_id = $1`, chat_id)
	if err != nil {
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var userID int32
		if err := rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("%s: row scan failed: %w", op, err)
		}
		users = append(users, userID)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration failed: %w", op, err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("users not found in chat")
	}

	return users, nil
}

func (s *Storage) GetMessagesFromChat(chat_id int32) ([]*desc.Message, error) {
	const op = "storage.pg.Users.GetMessages"
	var messages []*desc.Message

	rows, err := s.db.Query(`SELECT user_id, content, created_at FROM messages WHERE chat_id = $1`, chat_id)
	if err != nil {
		return nil, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var userID int32
		var content string
		var createdAt time.Time

		err := rows.Scan(&userID, &content, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("%s: row scan failed: %w", op, err)
		}

		message := &desc.Message{
			UserId:    userID,
			Text:      content,
			Time:      timestamppb.New(createdAt),
		}

		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration failed: %w", op, err)
	}

	return messages, nil
}