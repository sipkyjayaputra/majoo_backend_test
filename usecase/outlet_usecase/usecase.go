package outlet_usecase

import (
	"majoo/model/dto"
	"majoo/repository/outlet_repository"
)

type OutletUsecase interface {
	DailyTransaction(int) dto.Response
}

type outletUsecase struct {
	outletRepo outlet_repository.OutletRepository
}


func GetOutletUsecase(outletRepository outlet_repository.OutletRepository) OutletUsecase {
	return &outletUsecase{
		outletRepo: outletRepository,
	}
}
