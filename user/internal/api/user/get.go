package user

import (
	serv "user/internal/service"
	desc "user/pkg/user_v1"
	conv "user/internal/converter/user"
	"context"
	"database/sql"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Get(serv serv.UserService,ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	id := req.GetUserId()
	if id < 1 {
		return nil, status.Error(codes.InvalidArgument, "user id is incorrect or empty")
	}

	user, err := serv.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	res := conv.ToUserFromService(user)

	return res, nil
}