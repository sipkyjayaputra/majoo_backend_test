package merchant_repository

import (
	"majoo/model/entity"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	DailyTransaction(int) ([]entity.MerchantDailyTransaction, error)
}

type merchantRepository struct {
	mysqlConnection *gorm.DB
}

func GetMerchantRepository(mysqlConn *gorm.DB) MerchantRepository {
	return &merchantRepository{
		mysqlConnection: mysqlConn,
	}
}
