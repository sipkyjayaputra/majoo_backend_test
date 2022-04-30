package outlet_usecase

import (
	"majoo/helpers"
	"majoo/model/dto"
	"net/http"
)

func (usecase *outletUsecase) DailyTransaction(userId int) dto.Response {	
	transactionData, err := usecase.outletRepo.DailyTransaction(userId)

	if err != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Internal server error", err.Error(), nil)
	}
	

	return helpers.SendResponse(http.StatusOK, "Get Merchant Daily Transaction success", nil, transactionData)
}