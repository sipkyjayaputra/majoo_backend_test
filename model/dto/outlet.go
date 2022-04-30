package dto

import "time"

type Outlet struct {
	Id         int       `json:"id" gorm:"column:id;NOT NULL;type:bigInt(20);autoIncrementIncrement:true;primaryKey"`
	OutletId   int       `json:"outlet_id" gorm:"column:outlet_id;type:bigInt(20);NOT NULL"`
	OutletName string    `json:"outlet_name" gorm:"column:outlet_name;type:varchar(40);NOT NULL"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;NOT NULL"`
	CreatedBy  int       `json:"created_by" gorm:"type:bigInt(20);column:created_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;NOT NULL"`
	UpdatedBy  int       `json:"updated_by" gorm:"type:bigInt(20);column:updated_by"`
}