package chat

import (
	desc "chat/pkg/chat_v1"
	conv "chat/internal/converter/chat"
	"context"
	"database/sql"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ChatImplementation) Get(ctx context.Context, req *desc.GetChatRequest) (res *desc.GetChatResponse, err error) {
	id := req.GetChatId()
	if id < 1 {
		return nil, status.Error(codes.InvalidArgument, "chat id is incorrect or empty")
	}

	chat, err := i.ChatService.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "failed to get chat")
	}

	res = conv.ToChatFromService(chat)

	return res, nil
}