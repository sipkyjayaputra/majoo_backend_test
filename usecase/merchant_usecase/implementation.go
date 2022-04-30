package merchant_usecase

import (
	"majoo/helpers"
	"majoo/model/dto"
	"net/http"
)

func (usecase *merchantUsecase) DailyTransaction(userId int) dto.Response {	
	transactionData, err := usecase.merchantRepo.DailyTransaction(userId)

	if err != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Internal server error", err.Error(), nil)
	}
	

	return helpers.SendResponse(http.StatusOK, "Get Merchant Daily Transaction success", nil, transactionData)
}