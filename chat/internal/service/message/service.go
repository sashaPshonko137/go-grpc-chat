package message

import (
	pg "chat/internal/repo/storage"
	"chat/internal/service"
	userProto "chat/userclient/pkg/user_v1"
)

type messageService struct {
	storage pg.Storage
	clientUser userProto.UserV1Client
}

func NewMessageService(storage pg.Storage, clientUser userProto.UserV1Client) service.MessageService {
	return &messageService{
		storage: storage,
		clientUser: clientUser,
	}
}