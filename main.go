package main

import (
	"fmt"
	"log"
	"os"
	"users-service/database"
	"users-service/server"

	"github.com/gin-gonic/gin"
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

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "puerto donde ejecuta front")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})

	puerto := 8080
	direccion := fmt.Sprintf(":%d", puerto)

	fmt.Println("Hello server listing in:", puerto)
	log.Fatal(r.Run(direccion))
}
