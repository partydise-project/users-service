package server

import (
	"users-service/database"

	"github.com/gin-gonic/gin"
)

func CreateTrabajador(c *gin.Context) {
	var userCreateDTO database.Trabajador

	if err := c.BindJSON(&userCreateDTO); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Validate UserID required.
	if userCreateDTO.UsuarioID == 0 {
		c.JSON(404, gin.H{"error": "the userID must not be 0"})
		return
	}

	err := database.CreateTrabajador(&userCreateDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(201, userCreateDTO)
}

func ReadTrabajador(c *gin.Context) {
	usuario, err := database.ReadTrabajador(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "No se encuentra registrado un trabajador con ese id."})
		return
	}

	c.JSON(200, usuario)
}

func ReadTrabajadores(c *gin.Context) {
	usuarios, err := database.ReadTrabajadores()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, usuarios)
}

func UpdateTrabajador(c *gin.Context) {
	user, err := database.ReadTrabajador(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"No se encontro registros con ese id": err})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Error al decodificar el JSON"})
		return
	}

	err = database.UpdateTrabajador(user)
	if err != nil {
		c.JSON(500, gin.H{"error al actualizar el trabajador": err})
		return
	}

	c.JSON(201, user)
}
