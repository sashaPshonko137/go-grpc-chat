package user

import (
	model "user/internal/model/user"
	serv "user/internal/service"
	desc "user/pkg/user_v1"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Create(serv serv.UserService,ctx context.Context, req *desc.CreateUserRequest) (*emptypb.Empty, error)  {
	name := req.GetName()
	if name == "" {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "name is required")
	}

	info := &model.UserInfo{
		Name: name,
	}

	err := serv.Create(ctx, info)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "failed to create user")
	}

	return &emptypb.Empty{}, nil
}