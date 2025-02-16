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

func (c LoginConverter) EntityToDTO(user entity.User) *dto.LoginResponseBody {
	return &dto.LoginResponseBody{
		Username: user.Username,
		AccesToken: user.AccesToken,
	}
}