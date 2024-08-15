package chat

import (
	"chat/internal/repo"
	"chat/internal/service"
)

type serv struct {
	chatRepository repo.Repo
}

func NewService(chatRepository repo.Repo) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
	}
}

func NewMockService(deps ...interface{}) service.ChatService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repo.Repo:
			srv.chatRepository = s
		}
	}

	return &srv
}
