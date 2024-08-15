package app

import (
	"fmt"

	chatImpl "chat/internal/api/chat"
	mesImpl "chat/internal/api/message"
	"chat/internal/config"
	"chat/internal/config/env"
	"chat/internal/repo/storage"
	"chat/internal/service"
	chatServ "chat/internal/service/chat"
	mesServ "chat/internal/service/message"
	userProto "chat/userclient/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	Config config.GRPCConfig
	Storage *pg.Storage
	ChatService service.ChatService
	MessageService service.MessageService
	ClientUser userProto.UserV1Client
	ChatImplementation *chatImpl.ChatImplementation
	MessageImplementation *mesImpl.MessageImplementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) init() error {
	var err error

	s.Config, err = env.NewGRPCConfig()
	if err != nil {
		return fmt.Errorf("failed to get grpc config: %w", err)
	}

	s.Storage, err = pg.New(s.Config.GetDbUrl())
	if err != nil {
		return fmt.Errorf("failed to create storage: %w", err)
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}

	s.ClientUser = userProto.NewUserV1Client(conn)

	s.ChatService = chatServ.NewChatService(s.Storage, s.ClientUser)
	s.MessageService = mesServ.NewMessageService(*s.Storage, s.ClientUser)

	s.ChatImplementation = chatImpl.NewChatImplementation(s.ChatService)
	s.MessageImplementation = mesImpl.NewMessageImplementation(s.MessageService)

	return nil
}