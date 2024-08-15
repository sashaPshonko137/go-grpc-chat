package chat

import (
	"context"
	model "chat/internal/model/chat"
)

func (s *serv) Get(ctx context.Context, id int32 ) (*model.Chat, error) {
	chat, err := s.chatRepository.GetChat(id)
	if err != nil {
		return nil, err
	}
	return chat, nil
}