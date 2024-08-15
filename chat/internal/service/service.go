//go:generate minimock -i ChatService -o ./mocks/ -s "_minimock.go"

package service

import (
	"context"
	chatModel "chat/internal/model/chat"
	messageModel "chat/internal/model/message"
)

type ChatService interface {
	Create(ctx context.Context, info *chatModel.ChatInfo) error
	Get(ctx context.Context, id int32) (*chatModel.Chat, error)
}

type MessageService interface {
	Create(ctx context.Context, info *messageModel.MessageInfo) error
	GetManyFromChat(ctx context.Context, id int32) ([]*messageModel.Message, error)
}