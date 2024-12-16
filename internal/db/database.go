package db

import (
	"fmt"
	"log"

	"github.com/HaikalRFadhilahh/go-todos-list/pkg/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	HOST := helper.GetEnv("DB_HOST", "127.0.0.1")
	PORT := helper.GetEnv("DB_PORT", "3306")
	USERNAME := helper.GetEnv("DB_USERNAME", "root")
	PASSWORD := helper.GetEnv("DB_PASSWORD", "")
	NAME := helper.GetEnv("DB_NAME", "")

	// Dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Database Connection Error, Error : ", err.Error())
	}

	fmt.Println("Database Success Connected!")

	return db
}
