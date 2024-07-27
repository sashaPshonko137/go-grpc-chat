package main

import (
	desc "chat/pkg/chat_v1"
	"chat/storage/pg"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	grpcPort = ":50051"
	dbUrrl = "postgresql://postgres:1242@localhost:5432/postgres?sslmode=disable"
)

type server struct{
	desc.UnimplementedChatV1Server
	storage *pg.Storage
}

func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (res *emptypb.Empty, err error) {
	name := req.GetName()
	if name == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "name is required")
	}

	err = s.storage.CreateUser(name)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "name is required")
	}

	return &emptypb.Empty{}, nil
}

func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (res *emptypb.Empty, err error) {
	name := req.GetName()
	if name == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "name is required")
	}

	chatId, err := s.storage.CreateChat(name)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to create chat")
	}

	userIds := req.GetUserIds()
	if len(userIds) == 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "userIds is required")
	}

	for _, userId := range userIds {
		_, err = s.storage.GetUser(userId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "user not exist")
			}
			return &emptypb.Empty{}, status.Error(codes.Internal, "failed to get user")
		}
		err = s.storage.CreateChatUser(chatId, userId)
		if err != nil {
			return &emptypb.Empty{}, status.Error(codes.Internal, "failed to add users")
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *server) WriteMessage(ctx context.Context, req *desc.WriteRequest) (res *emptypb.Empty, err error) {
	userId := req.GetUserId()
	if userId == 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "userId is required")
	}
	chatId := req.GetChatId()
	if chatId == 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "chatId is required")
	}
	message := req.GetMessage()
	if message == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "message is required")
	}
	_, err = s.storage.GetUser(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "user not exist")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to get user")
	}
	_, err = s.storage.GetChat(chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "chat not exist")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to get chat")
	}
	err = s.storage.CreateMessage(userId, chatId, message, timestamppb.Now())
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to send message")
	}
	return &emptypb.Empty{}, nil
}

func (s *server) GetMessagesFromChat(ctx context.Context, req *desc.GetMessagesRequest) (res *desc.GetMessagesResponse, err error) {
	chatId := req.GetChatId()
	if chatId == 0 {
		return nil, status.Error(codes.InvalidArgument, "chat is required")
	}
	_, err = s.storage.GetChat(chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.InvalidArgument, "chat not exist")
		}
		return nil, status.Error(codes.Internal, "failed to get chat")
	}
	messages, err := s.storage.GetMessagesFromChat(chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.InvalidArgument, "messages not exist")
		}
		return nil, status.Error(codes.Internal, "failed to get messages")
	}

	return &desc.GetMessagesResponse{
		Messages: messages,
	}, nil
}

func main() {
	storage, err := pg.New(dbUrrl)
	if err != nil {
		fmt.Println("error to create storage")
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("error to listen port%v\n", grpcPort)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{storage: storage})

	fmt.Printf("server started%v\n", grpcPort)

	if err = s.Serve(lis); err != nil {
		fmt.Println("error start server")
	}
}