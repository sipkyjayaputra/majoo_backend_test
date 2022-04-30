package user_delivery

import (
	"majoo/usecase/user_usecase"

	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type userDelivery struct {
	usecase user_usecase.UserUsecase
}


func GetUserDelivery(userUsecase user_usecase.UserUsecase) UserDelivery {
	return &userDelivery{
		usecase: userUsecase,
	}
}
