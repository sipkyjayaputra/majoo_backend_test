package entity

import "time"

type Outlet struct {
	Id         int       `json:"id" gorm:"column:id;NOT NULL;type:bigInt(20)  AUTO_INCREMENT;autoIncrementIncrement:true;primaryKey"`
	MerchantId   int       `json:"merchant_id" gorm:"column:merchant_id;type:bigInt(20);NOT NULL"`
	OutletName string    `json:"outlet_name" gorm:"column:outlet_name;type:varchar(40);NOT NULL"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;NOT NULL"`
	CreatedBy  int       `json:"created_by" gorm:"type:bigInt(20);column:created_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;NOT NULL"`
	UpdatedBy  int       `json:"updated_by" gorm:"type:bigInt(20);column:updated_by"`
}

type OutletDailyTransaction struct {
	Merchant	Merchant	`json:"merchant"`
	Outlet 		Outlet		`json:"outlet"`
	Transactions	[]DailyTransactionForOutlet	`json:"transactions"`
	StartDate	string	`json:"start_date"`
	EndDate		string	`json:"end_date"`
	Limit          int    	`json:"limit"`
	Offset         int    	`json:"offset"`
	Page           int    	`json:"page"`
}