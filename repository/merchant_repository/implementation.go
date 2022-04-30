package merchant_repository

import (
	"majoo/model/entity"
)

// Register implements UserRepository
func (repo *merchantRepository) DailyTransaction(userId int) ([]entity.MerchantDailyTransaction, error) {
	merchants := []entity.Merchant{}
	merchant_transactions := []entity.MerchantDailyTransaction{}

	if err := repo.mysqlConnection.Model(&entity.Merchant{}).Where("user_id = ?", userId).Scan(&merchants).Error; err != nil {
		return merchant_transactions, err
	}

	for _, merchant := range(merchants) {
		transactions := []entity.DailyTransactionForMerchant{}
		if err := repo.mysqlConnection.Select("A.id, B.merchant_name, SUM(A.bill_total) AS omzet, A.created_at AS date").Table("transactions A").Joins("left join merchants B ON A.merchant_id = B.id").Where("A.merchant_id = ?", merchant.Id).Where("A.created_at BETWEEN ? AND ?", "2022-04-01", "2022-12-30").Group("A.merchant_id").Find(&transactions).Error; err != nil {
			return merchant_transactions, err
		}
		merchant_transactions = append(merchant_transactions, entity.MerchantDailyTransaction{
			Merchant: merchant,
			Transactions: transactions,
		})
	}

	return merchant_transactions, nil
}