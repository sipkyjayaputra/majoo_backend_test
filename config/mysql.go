package config

import (
	"fmt"
	"log"
	"majoo/model/entity"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mySQL *gorm.DB

func InitMySQL() (*gorm.DB, error) {
	var err error
	// CONNECTION STRING TO MYSQL
	connectionString := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		CONFIG["MYSQL_USER"],
		CONFIG["MYSQL_PASS"],
		CONFIG["MYSQL_HOST"],
		CONFIG["MYSQL_PORT"],
		CONFIG["MYSQL_SCHEMA"],
	)

	// INITIALIZE LOGGER
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // IO WRITER
		logger.Config{
			SlowThreshold: time.Second,			// SLOW SQL THRESHOLD
			LogLevel: logger.Info,				// LOG LEVEL
			IgnoreRecordNotFoundError: false,	// IGNORE ErrRecordNotFound for logger
			Colorful: false,					// DISABLE COLOR
		},
	)

	mySQL, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, err
	}else {
		sqlDB, err_sql := mySQL.DB()
		if err_sql != nil {
			sqlDB.SetMaxIdleConns(2)
			sqlDB.SetMaxOpenConns(1000)
		}
		log.Println(mySQL)
		log.Println("MySQL is connected.")
	}

	// MIGRATE MODEL
	mySQL.AutoMigrate(
		entity.User{}, 
		entity.Merchant{}, 
		entity.Outlet{}, 
		entity.Transaction{},
	)
	
	return mySQL, err
}

func GetMySQL() *gorm.DB {
	return mySQL
}