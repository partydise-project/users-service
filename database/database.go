package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InicializeDB(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Usuario{}, &Recreador{}, &Cliente{}, &Trabajador{})

	DB = db
	return nil
}

func Create(user *Usuario) error {
	result := DB.Create(user)
	if result.Error != nil {
		fmt.Println("Error en el servicio de la base de datos.")
		return result.Error
	}
	return nil
}
