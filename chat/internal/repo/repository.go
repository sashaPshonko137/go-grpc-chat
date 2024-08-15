package repo

import (
	"time"
	chatModel "chat/internal/model/chat"
	messageModel "chat/internal/model/message"
)

type Repo interface {
	CreateChat(name string) (int32, error)
	CreateMessage(user_id, chat_id int32, content string, created_at time.Time) error
	CreateChatUser(chat_id, user_id int32) error
	GetChat(chat_id int32) (*chatModel.Chat, error)
	GetMessagesFromChat(chatId int32) ([]*messageModel.Message, error)
}