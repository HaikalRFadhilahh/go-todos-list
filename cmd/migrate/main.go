package main

import (
	"fmt"
	"log"

	"github.com/HaikalRFadhilahh/go-todos-list/internal/db"
	"github.com/HaikalRFadhilahh/go-todos-list/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := db.InitDB()

	if err := db.AutoMigrate(models.Users{}, models.Tasks{}); err != nil {
		log.Fatalln("Migration Error, Has Error :", err.Error())
	}

	fmt.Println("Migration Success!")
}
