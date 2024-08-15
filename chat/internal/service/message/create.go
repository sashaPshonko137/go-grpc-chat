package message

import (
	model "chat/internal/model/message"
	userProto "chat/userclient/pkg/user_v1"
	"context"
	"fmt"
)

func (s *messageService) Create(ctx context.Context, info *model.MessageInfo) error {
	_, err := s.clientUser.GetUser(ctx, &userProto.GetUserRequest{UserId: info.UserId})
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	_, err = s.storage.GetChat(info.ChatId)
	if err != nil {
		return err
	}
	err = s.storage.CreateMessage(info.UserId, info.ChatId, info.Text, info.Time)
	if err != nil {
		return err
	}
	return nil
}