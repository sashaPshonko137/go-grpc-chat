package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	apiUser "user/internal/api/user"
	config "user/internal/config"
	"user/internal/config/env"
	pg "user/internal/repo/storage"
	serv "user/internal/service"
	"user/internal/service/user"
	desc "user/pkg/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct{
	desc.UnimplementedChatV1Server
	storage *pg.Storage
	userService *serv.UserService
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config path")
}

func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (res *emptypb.Empty, err error) {
	return apiUser.Create(*s.userService, ctx, req)
}

func (s *server) GetUser(ctx context.Context, req *desc.GetUserRequest) (res *desc.GetUserResponse, err error) {
	fmt.Println(11)
	return apiUser.Get(*s.userService, ctx, req)
}

func main() {
	flag.Parse()
	err := config.Load(configPath)
	if err != nil {
		log.Fatal("failed to load config path", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatal("failed to get grpc config", err)
	}

	storage, err := pg.New(grpcConfig.GetDbUrl())
	if err != nil {
		log.Fatal("error to create storage", err)
		os.Exit(1)
	}

	userService := user.NewUserService(*storage)

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal("error to listen port", grpcConfig.Address())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{
		storage: storage,
		userService: &userService,
	},
)

	fmt.Printf("server started on %v\n", grpcConfig.Address())

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}