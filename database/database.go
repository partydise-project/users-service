package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Usuario es el modelo para la entidad Usuario
type Usuario struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func InicializeDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Usuario{})
}
