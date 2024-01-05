package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusCreated, "Welcome to the users-services.")
}

func InicializeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", Hello)

	//User module.
	r.POST("/usuarios", CreateUser)
	r.GET()
	//read all
	r.PATCH("/usuario", UpdateUser)

	//Recreador module.
	//CLient module.
	//Trabajador module.
	return r
}
