package user_repository

import (
	"majoo/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(entity.User) (*entity.User, error)
	Login(entity.User) (*entity.User, error)
}

type userRepository struct {
	mysqlConnection *gorm.DB
}

func GetUserRepository(mysqlConn *gorm.DB) UserRepository {
	return &userRepository{
		mysqlConnection: mysqlConn,
	}
}
