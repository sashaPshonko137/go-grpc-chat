package chat

import (
	"chat/internal/service"
)

type ChatImplementation struct {
	ChatService service.ChatService
}

func NewChatImplementation(chatService service.ChatService) *ChatImplementation {
	return &ChatImplementation{
		ChatService: chatService,
	}
}