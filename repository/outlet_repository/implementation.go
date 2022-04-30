package outlet_repository

import (
	"majoo/model/entity"
)

// Register implements UserRepository
func (repo *outletRepository) DailyTransaction(userId int) ([]entity.OutletDailyTransaction, error) {
	merchants := []entity.Merchant{}
	outlet_transactions := []entity.OutletDailyTransaction{}

	if err := repo.mysqlConnection.Model(&entity.Merchant{}).Where("user_id = ?", userId).Scan(&merchants).Error; err != nil {
		return outlet_transactions, err
	}

	for _, merchant := range(merchants) {
		outlets := []entity.Outlet{}
		if err := repo.mysqlConnection.Model(&entity.Outlet{}).Where("merchant_id = ?", merchant.Id).Scan(&outlets).Error; err != nil {
			return outlet_transactions, err
		} 
		for _, outlet := range(outlets) {
			transactions := []entity.DailyTransactionForOutlet{}
			if err := repo.mysqlConnection.Select("A.id, B.merchant_name, C.outlet_name, SUM(A.bill_total) AS omzet, A.created_at AS date").Table("transactions A").Joins("left join merchants B ON A.merchant_id = B.id").Joins("left join outlets C ON A.outlet_id = C.id").Where("A.outlet_id = ?", outlet.Id).Where("A.created_at BETWEEN ? AND ?", "2022-04-01", "2022-12-30").Group("A.outlet_id").Find(&transactions).Error; err != nil {
				return outlet_transactions, err
			}
			outlet_transactions = append(outlet_transactions, entity.OutletDailyTransaction{
				Merchant: merchant,
				Outlet: outlet,
				Transactions: transactions,
			})
			
		}
	}

	return outlet_transactions, nil
}