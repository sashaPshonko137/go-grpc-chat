package chat

import (
	model "chat/internal/model/chat"
	"context"
)

func (s *chatService) Get(ctx context.Context, id int32) (*model.Chat, error) {
	chat, err := s.storage.GetChat(id)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
