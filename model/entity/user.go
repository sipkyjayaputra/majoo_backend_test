package entity

import "time"

type User struct {
	Id	int	`json:"id" gorm:"column:id;NOT NULL;type:bigInt(20) AUTO_INCREMENT;autoIncrement;primaryKey"`
	Name	string	`json:"name" gorm:"column:name;type:varchar(45);default:NULL"`
	Username	string	`json:"user_name" gorm:"column:user_name;type:varchar(45);default:NULL"`
	Password	string	`json:"password" gorm:"column:password;type:varchar(255);default:NULL"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at;NOT NULL"`
	CreatedBy       int    `json:"created_by" gorm:"type:bigInt(20);column:created_by"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at;NOT NULL"`
	UpdatedBy       int    `json:"updated_by" gorm:"type:bigInt(20);column:updated_by"`
}
