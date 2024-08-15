package chat

import (
	"chat/internal/service"
	userProto "chat/userclient/pkg/user_v1"
	storage "chat/internal/repo"
)

type chatService struct {
	storage storage.Repo
	clientUser userProto.UserV1Client
}

func NewChatService(storage storage.Repo, clientUser userProto.UserV1Client) service.ChatService {
	return &chatService{
		storage: storage,
		clientUser: clientUser,
	}
}