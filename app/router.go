package app

import (
	"majoo/delivery/merchant_delivery"
	"majoo/delivery/outlet_delivery"
	"majoo/delivery/user_delivery"
	"majoo/middleware"
	"majoo/repository/merchant_repository"
	"majoo/repository/outlet_repository"
	"majoo/repository/user_repository"
	"majoo/usecase/merchant_usecase"
	"majoo/usecase/outlet_usecase"
	"majoo/usecase/user_usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(MySQL *gorm.DB) *gin.Engine {
	// USER SECTION
	userRepository := user_repository.GetUserRepository(MySQL)
	userUsecase := user_usecase.GetUserUsecase(userRepository)
	userDelivery := user_delivery.GetUserDelivery(userUsecase)
	// MERCHANT SECTION
	merchantRepository := merchant_repository.GetMerchantRepository(MySQL)
	merchantUsecase := merchant_usecase.GetMerchantUsecase(merchantRepository)
	merchantDelivery := merchant_delivery.GetMerchantDelivery(merchantUsecase)
	// OUTLET SECTION
	outletRepository := outlet_repository.GetOutletRepository(MySQL)
	outletUsecase := outlet_usecase.GetOutletUsecase(outletRepository)
	outletDelivery := outlet_delivery.GetOutletDelivery(outletUsecase)
	
	router := gin.Default()
	

	authRouter := router.Group("/")
	authRouter.Use(middleware.JWTauth())
	{
		authRouter.GET("/merchant-daily-transaction", merchantDelivery.DailyTransaction)
		authRouter.GET("/outlet-daily-transaction", outletDelivery.DailyTransaction)
	}
	router.POST("/login", userDelivery.Login)
	router.POST("/register", userDelivery.Register)

	return router
}