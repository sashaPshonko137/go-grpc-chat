package chat

import (
	model "chat/internal/model/chat"
	desc "chat/pkg/chat_v1"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ChatImplementation) Create(ctx context.Context, req *desc.CreateChatRequest) (res *emptypb.Empty, err error)  {
	name := req.GetName()
	if name == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "name is required")
	}

	info := &model.ChatInfo{
		Name: name,
	}

	err = i.ChatService.Create(ctx, info)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to create chat")
	}

	return &emptypb.Empty{}, nil
}