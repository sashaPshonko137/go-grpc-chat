package service

import (
	"context"
	userModel "user/internal/model/user"
)

type UserService interface {
	Create(ctx context.Context, info *userModel.UserInfo) error
	Get(ctx context.Context, id int32) (*userModel.User, error)
}