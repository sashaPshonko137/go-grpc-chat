package user

import (
	"user/internal/repo/storage"
	"user/internal/service"
)

type userService struct {
	storage pg.Storage
}

func NewUserService(storage pg.Storage) service.UserService {
	return &userService{
		storage: storage,
	}
}
