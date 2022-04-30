package merchant_delivery

import (
	"majoo/usecase/merchant_usecase"

	"github.com/gin-gonic/gin"
)

type MerchantDelivery interface {
	DailyTransaction(*gin.Context)
}

type merchantDelivery struct {
	usecase merchant_usecase.MerchantUsecase
}

func GetMerchantDelivery(merchantUc merchant_usecase.MerchantUsecase) MerchantDelivery {
	return &merchantDelivery{
		usecase: merchantUc,
	}
}
