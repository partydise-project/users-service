package server

import (
	"users-service/database"

	"github.com/gin-gonic/gin"
)

func CreateRecreador(c *gin.Context) {
	var userCreateDTO database.Recreador
	user := userCreateDTO.Usuario

	if err := c.BindJSON(&userCreateDTO); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	err := database.CreateRecreador(&user, &userCreateDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(201, userCreateDTO)
}

func ReadRecreador(c *gin.Context) {
	usuario, err := database.ReadRecreador(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "No se encuentra registrado un recreador con ese id."})
		return
	}

	c.JSON(200, usuario)
}

func ReadRecreadores(c *gin.Context) {
	usuarios, err := database.ReadRecreadores()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, usuarios)
}

func UpdateRecreador(c *gin.Context) {
	user, err := database.ReadRecreador(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"No se encontro registros con ese id": err})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar el JSON"})
		return
	}

	err = database.UpdateRecreador(user)
	if err != nil {
		c.JSON(500, gin.H{"error al actualizar el recreador": err})
		return
	}

	c.JSON(201, user)
}
