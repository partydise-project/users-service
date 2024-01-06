package server

import (
	"users-service/database"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userCreateDTO database.Usuario

	if err := c.BindJSON(&userCreateDTO); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Validate IDauth0 required.
	if userCreateDTO.Auth0ID == "" {
		c.JSON(404, gin.H{"error": "ID de Auth0 es obligatorio."})
		return
	}

	err := database.CreateUsuario(&userCreateDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(201, userCreateDTO)
}

func ReadUser(c *gin.Context) {
	usuario, err := database.ReadUsuario(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "No se encuentra registrado un usuario con ese id."})
		return
	}

	c.JSON(200, usuario)
}

func ReadUsers(c *gin.Context) {
	usuarios, err := database.ReadUsuarios()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, usuarios)
}

func UpdateUser(c *gin.Context) {
	user, err := database.ReadUsuario(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"No se encontro registros con ese id": err})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar el JSON"})
		return
	}

	err = database.UpdateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error al actualizar el usuario": err})
		return
	}

	c.JSON(201, user)
}
