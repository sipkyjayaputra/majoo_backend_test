package user_usecase

import (
	"majoo/model/dto"
	"majoo/repository/user_repository"
)

type UserUsecase interface {
	Login(dto.LoginRequest) dto.Response
	Register(dto.RegisterRequest) dto.Response
}

type userUsecase struct {
	userRepo user_repository.UserRepository
}


func GetUserUsecase(userRepository user_repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepository,
	}
}
