package main

import (
	"fmt"
	"log"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/db"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := db.InitDB()

	if err := db.AutoMigrate(model.Users{}, model.Tasks{}); err != nil {
		log.Fatalln("Migration Error, Has Error :", err.Error())
	}

	fmt.Println("Migration Success!")
}
