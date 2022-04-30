package outlet_repository

import (
	"majoo/model/entity"

	"gorm.io/gorm"
)

type OutletRepository interface {
	DailyTransaction(int) ([]entity.OutletDailyTransaction, error)
}

type outletRepository struct {
	mysqlConnection *gorm.DB
}

func GetOutletRepository(mysqlConn *gorm.DB) OutletRepository {
	return &outletRepository{
		mysqlConnection: mysqlConn,
	}
}
