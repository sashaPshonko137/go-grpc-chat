package message

import (
	"chat/internal/service"
)

type MessageImplementation struct {
	messageService service.MessageService
}

func NewMessageImplementation(messageService service.MessageService) *MessageImplementation {
	return &MessageImplementation{
		messageService: messageService,
	}
}