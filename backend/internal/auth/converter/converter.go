package converter

import (
	"auth-app/internal/dto"
	"auth-app/internal/entity"
)

type LoginConverter struct{}

func (c LoginConverter) DTOToEntity(dto *dto.LoginRequestBody) *entity.User {
	return &entity.User{
		Email: dto.Email,
		Password: dto.Password,
	}
}

type UserConverter struct{}

func (c UserConverter) EntityToDTO(user *entity.User) *dto.UserResponseBody {
	return &dto.UserResponseBody{
		Email: user.Email,
		Username: user.Username,
	}
}

type RegisterConverter struct{}

func (c RegisterConverter) DTOToEntity(dto *dto.RegisterRequestBody) *entity.User {
	return &entity.User{
		Username: dto.Username,
		Email: dto.Email,
		Password: dto.Password,
	}
}