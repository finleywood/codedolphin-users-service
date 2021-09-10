package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_CONNECTION")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Could not init db connection, err: " + err.Error())
	}
	DB = db
	initModels()
}

func initModels() {
	DB.AutoMigrate(User{})
}
