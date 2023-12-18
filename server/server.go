package server

import (
	"fmt"
	"net/http"
	"users-service/database"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userCreateDTO database.Usuario

	if err := c.BindJSON(&userCreateDTO); err != nil {
		fmt.Println("error al recibir el objeto(user) en el request", err)

		c.JSON(404, "error error al crear el user")
		return
	}

	database.Create(&userCreateDTO)
	c.JSON(http.StatusCreated, userCreateDTO)
}
func Hello(c *gin.Context) {
	c.JSON(http.StatusCreated, "Hola mundo desde un endpoint de gin ☻.")
}

func InicializeRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/usuarios", CreateUser)
	r.GET("/", Hello)
	return r
}
