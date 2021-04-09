package dto

import "gin-my/model"

type UserDto struct {
	Username string `json:"username"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
	}
}
