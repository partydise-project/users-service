package main

import (
	"fmt"
	"log"
	"os"
	"users-service/database"

	"github.com/joho/godotenv"
)

func initializeEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initializeEnv()

	dns := os.Getenv("DSN")

	database.InicializeDB(dns)

	fmt.Println("Hello word", dns)
}
