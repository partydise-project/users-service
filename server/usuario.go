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

	c.JSON(http.StatusCreated, userCreateDTO)
}

func ReadUser(c *gin.Context) {

}

func ReadUsers(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {
	var user database.Usuario

	if err := c.BindJSON(&user); err != nil {
		fmt.Println("error al recibir el objeto(user) en el request", err)

		c.JSON(404, "error error al crear el user")
		return
	}

	c.JSON(http.StatusCreated, user)
}
