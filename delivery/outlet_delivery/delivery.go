package outlet_delivery

import (
	"majoo/usecase/outlet_usecase"

	"github.com/gin-gonic/gin"
)

type OutletDelivery interface {
	DailyTransaction(*gin.Context)
}

type outletDelivery struct {
	usecase outlet_usecase.OutletUsecase
}

func GetOutletDelivery(outletUc outlet_usecase.OutletUsecase) OutletDelivery {
	return &outletDelivery{
		usecase: outletUc,
	}
}
