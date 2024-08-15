package message

import (
	model "chat/internal/model/message"
	"context"
	"fmt"
)

func (s *messageService) GetManyFromChat(ctx context.Context, id int32) ([]*model.Message, error) {
	_, err := s.storage.GetChat(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat:%w", err)
	}
	messages, err := s.storage.GetMessagesFromChat(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages:%w", err)
	}

	return messages, nil
}