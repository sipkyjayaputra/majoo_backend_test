package helpers

import (
	"fmt"
	"majoo/model/entity"

	"gorm.io/gorm"
)

func CheckUserExisted(mysql *gorm.DB, user_name string) bool {
	user := entity.User{}
	if err := mysql.Model(&entity.User{}).Where("user_name = ?", user_name).Scan(&user).Error; err != nil {
		return false
	}

	fmt.Println(user)
	if user.Id == 0 {
		return true
	}
	return false
}