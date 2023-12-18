package main

import (
	"fmt"
	"log"
	"os"
	"users-service/database"
	"users-service/server"

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

	err := database.InicializeDB(dns)
	if err != nil {
		fmt.Println("Error al inicializar la base de datos:", err)
		return
	}

	r := server.InicializeRouter()

	puerto := 8080
	direccion := fmt.Sprintf(":%d", puerto)

	fmt.Println("Hello word", dns)
	log.Fatal(r.Run(direccion))
}
