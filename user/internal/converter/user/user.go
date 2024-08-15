package user

import (
	model "user/internal/model/user"
	desc "user/pkg/user_v1"
)

func ToUserFromService(user *model.User) *desc.GetUserResponse {
		return &desc.GetUserResponse{
			UserId: user.Id,
			Name: user.Name,
		}
}