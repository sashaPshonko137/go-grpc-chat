package message

import (
	serv "chat/internal/service"
	desc "chat/pkg/chat_v1"
	conv "chat/internal/converter/message"
	"context"
	"database/sql"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetMany(serv serv.MessageService,ctx context.Context, req *desc.GetMessagesRequest) (*desc.GetMessagesResponse, error) {
	id := req.GetChatId()
	if id < 1 {
		return nil, status.Error(codes.InvalidArgument, "chat id is incorrect or empty")
	}

	messages, err := serv.GetManyFromChat(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "failed to get messages")
	}

	res := conv.ToUserFromService(messages)

	return res, nil
}