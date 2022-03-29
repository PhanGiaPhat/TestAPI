package storage

import (
	config "TestAPI/config"
	"log"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // -> Log các câu lệnh truy vấn database trong terminal
	})

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
