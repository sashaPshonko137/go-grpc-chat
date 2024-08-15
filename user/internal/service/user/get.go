package user

import (
	model "user/internal/model/user"
	"context"
	"fmt"
)

func (s *userService) Get(ctx context.Context, id int32) (*model.User, error) {
	user, err := s.storage.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user:%w", err)
	}
	return user, nil
}
