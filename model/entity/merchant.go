package entity

import "time"

type Merchant struct {
	Id           int       `json:"id" gorm:"column:id;NOT NULL;type:bigInt(20) AUTO_INCREMENT;autoIncrementIncrement:true;primaryKey"`
	UserId       int       `json:"user_id" gorm:"column:user_id;type:int(40);NOT NULL"`
	MerchantName string    `json:"merchant_name" gorm:"column:merchant_name;type:varchar(40);NOT NULL"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;NOT NULL"`
	CreatedBy    int       `json:"created_by" gorm:"type:bigInt(20);column:created_by"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;NOT NULL"`
	UpdatedBy    int       `json:"updated_by" gorm:"type:bigInt(20);column:updated_by"`
}

type MerchantDailyTransaction struct {
	Merchant	Merchant	`json:"merchant"`
	Transactions	[]DailyTransactionForMerchant	`json:"transactions"`
	StartDate	string	`json:"start_date"`
	EndDate		string	`json:"end_date"`
	Limit          int    	`json:"limit"`
	Offset         int    	`json:"offset"`
	Page           int    	`json:"page"`
}