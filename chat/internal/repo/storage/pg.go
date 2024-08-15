package pg

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
	chatModel "chat/internal/model/chat"
	messageModel "chat/internal/model/message"
)

type Storage struct {
	db *sql.DB
}

func New(dbPath string) (*Storage, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chats (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_chats_id ON chats(id);`)
	if err != nil {
	  return nil, err
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
		return nil, err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_messages_id ON messages(id);`)
	if err != nil {
	  return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS chat_users (
			id SERIAL PRIMARY KEY,
			user_id INT,
			chat_id INT
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_chat_users_id ON chat_users(id);`)
	if err != nil {
	  return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateChat(name string) (int32, error) {	
	stmt, err := s.db.Prepare(`INSERT INTO chats (name) VALUES ($1) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id int32
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) CreateMessage(user_id, chat_id int32, content string, created_at time.Time) error {
	_, err := s.db.Exec(`INSERT INTO messages (user_id, chat_id, content, created_at) VALUES ($1, $2, $3, $4)`, user_id, chat_id, content, created_at)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateChatUser(chat_id, user_id int32) error {
	_, err := s.db.Exec(`INSERT INTO chat_users (chat_id, user_id) VALUES ($1, $2)`, chat_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetChat(chat_id int32) (*chatModel.Chat, error) {
	var name string
	err := s.db.QueryRow(`SELECT name FROM chats WHERE id = $1`, chat_id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &chatModel.Chat{Name: name}, nil
}

func (s *Storage) GetMessagesFromChat(chatId int32) ([]*messageModel.Message, error) {
	var messages []*messageModel.Message

	rows, err := s.db.Query(`SELECT id, user_id, content, created_at FROM messages WHERE chat_id = $1`, chatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var messageId, userId int32
		var content string
		var createdAt time.Time

		err := rows.Scan(&messageId, &userId, &content, &createdAt)
		if err != nil {
			return nil, err
		}

		message := &messageModel.Message{
			MessageId: messageId,
			UserId:    userId,
			Text:      content,
			Time:      createdAt,
		}

		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}