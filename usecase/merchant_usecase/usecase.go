package merchant_usecase

import (
	"majoo/model/dto"
	"majoo/repository/merchant_repository"
)

type MerchantUsecase interface {
	DailyTransaction(int) dto.Response
}

type merchantUsecase struct {
	merchantRepo merchant_repository.MerchantRepository
}


func GetMerchantUsecase(merchantRepository merchant_repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{
		merchantRepo: merchantRepository,
	}
}
