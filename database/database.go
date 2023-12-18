package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Con objetivo de tener una unica conexion a la base de datos.
Se utiliza el patron de diseno creacional Singleton, para inicializar
el objeto *gorm.DB una unica vez.
*/
var DB *gorm.DB
var once sync.Once

func InicializeDB(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//Con GORM no es necesario utilizar un defer para cerrar la conexion, pues esto lo realiza la bilioteca.

	/*La sincronización con sync.Once garantiza que la inicialización del objeto
	DB solo ocurra una vez */
	once.Do(func() {
		DB = db
		errorMigraciones := db.AutoMigrate(&Usuario{}, &Recreador{}, &Cliente{}, &Trabajador{})
		if errorMigraciones != nil {
			err = fmt.Errorf("error al realizar las migraciones en la base de datos: %w", errorMigraciones)
			return
		}
	})

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
