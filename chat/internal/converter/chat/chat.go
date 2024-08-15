package chat

import (
	model "chat/internal/model/chat"
	desc "chat/pkg/chat_v1"
)

func ToChatFromService(chat *model.Chat) *desc.GetChatResponse {
		return &desc.GetChatResponse{
			ChatId: chat.Id,
			Name: chat.Name,
		}
}
