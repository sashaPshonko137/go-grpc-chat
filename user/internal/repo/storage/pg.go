package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
	userModel "user/internal/model/user"
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
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_users_id ON users(id);`)
	if err != nil {
	  return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(name string) error {
	_, err := s.db.Exec(`INSERT INTO users (name) VALUES ($1)`, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUser(user_id int32) (*userModel.User, error) {
	var name string
	err := s.db.QueryRow(`SELECT name FROM users WHERE id = $1`, user_id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &userModel.User{Id: user_id, Name: name}, nil
}

func (s *Storage) GetUsersFromChat(chat_id int32) ([]int32, error) {
	var users []int32

	rows, err := s.db.Query(`SELECT user_id FROM chat_users WHERE chat_id = $1`, chat_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userID int32
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		users = append(users, userID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, err
	}

	return users, nil
}