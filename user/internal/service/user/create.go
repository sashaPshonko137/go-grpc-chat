package user

import (
	model "user/internal/model/user"
	"context"
)

func (s *userService) Create(ctx context.Context, info *model.UserInfo) error {
	err := s.storage.CreateUser(info.Name)
	if err != nil {
		return err
	}

	return nil
}