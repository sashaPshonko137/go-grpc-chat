package message

import (
	model "chat/internal/model/message"
	serv "chat/internal/service"
	desc "chat/pkg/chat_v1"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Create(serv serv.MessageService,ctx context.Context, req *desc.WriteRequest) (*emptypb.Empty, error)  {
	userId := req.GetChatId()
	if userId <= 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "user id is incorrect or empty")
	}
	chatId := req.GetChatId()
	if chatId <= 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "chat id is incorrect or empty")
	}
	text := req.GetMessage()
	if text == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "message is incorrect or empty")
	}
	time := req.GetTime()
	if time == nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "time is incorrect or empty")
	}

	info := &model.MessageInfo{
		UserId: userId,
		ChatId: chatId,
		Text: text,
		Time: time.AsTime(),
	}

	err := serv.Create(ctx, info)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to send message")
	}

	return &emptypb.Empty{}, nil
}