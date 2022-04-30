package user_repository

import (
	"errors"
	"fmt"
	"majoo/helpers"
	"majoo/model/entity"
)

// Login implements UserRepository
func (repo *userRepository) Login(userRequest entity.User) (*entity.User, error) {
	userName := userRequest.Username
	plainPassword := userRequest.Password
	fmt.Print(userName + "" + plainPassword)
	if err := repo.mysqlConnection.Model(&entity.User{}).Where("user_name = ?", userName).Scan(&userRequest).Error; err != nil {
		return &entity.User{}, errors.New("Username or Password is invalid!")
	}


	if helpers.CheckPasswordHash(plainPassword, userRequest.Password) != true {
		return &entity.User{}, errors.New("Username or Password is invalid!")
	}

	return &userRequest, nil
}

// Register implements UserRepository
func (repo *userRepository) Register(userRequest entity.User) (*entity.User, error) {
	if helpers.CheckUserExisted(repo.mysqlConnection, userRequest.Username) != true {
		return &entity.User{}, errors.New("Username is already exist!")
	}
	
	if err := repo.mysqlConnection.Create(&userRequest).Error; err != nil {
		return &entity.User{}, err
	}

	return &userRequest, nil
}