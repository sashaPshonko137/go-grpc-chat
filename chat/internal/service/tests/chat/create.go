package chat

import (
	"context"
	model "chat/internal/model/chat"
)

func (s *serv) Create(ctx context.Context, info *model.ChatInfo) error {
	_, err := s.chatRepository.CreateChat(info.Name)
	if err != nil {
		return err
	}

	return nil
}