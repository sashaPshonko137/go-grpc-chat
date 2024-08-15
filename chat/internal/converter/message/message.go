package message

import (
	model "chat/internal/model/message"
	desc "chat/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(messages []*model.Message) *desc.GetMessagesResponse{
	var res []*desc.Message
	for _, mes := range messages {
		time := timestamppb.New(mes.Time)
		resultMessage := &desc.Message{
			UserId: mes.UserId,
			Text: mes.Text,
			Time: time,
			}
		res = append(res, resultMessage)
	}

	return &desc.GetMessagesResponse{
		Messages: res,
	}
}
