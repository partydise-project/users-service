package server

import (
	"net/http"
	"users-service/middleware"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusCreated, "Welcome to the users-services.")
}

func InicializeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", Hello)

	//Auth0 test.
	r.GET("/auth0-test", middleware.CheckJWT(), Hello)

	//User module.
	r.POST("/usuario", CreateUser)
	r.GET("/usuario/:id", ReadUser)
	r.GET("/usuarios", ReadUsers)
	r.PATCH("/usuario/:id", UpdateUser)

	//Recreador module.
	r.POST("/recreador", CreateRecreador)
	r.GET("/recreador/:id", ReadRecreador)
	r.GET("/recreadores", ReadRecreadores)
	r.PATCH("/recreador/:id", UpdateRecreador)

	//CLient module.
	//Trabajador module.
	return r
}
