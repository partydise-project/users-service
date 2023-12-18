package database

import "fmt"

func CreateUsuario(user *Usuario) error {
	result := DB.Create(user)
	if result.Error != nil {
		fmt.Println("Error al crear el usuario.")
		return result.Error
	}
	return nil
}

func ReadUsuario(id int) (*Usuario, error) {
	var user Usuario
	result := DB.First(&user, id)
	if result.Error != nil {
		fmt.Println("Error al crear el usuario.")
		return nil, result.Error
	}

	return &user, nil
}

func ReadUsuarios() ([]Usuario, error) {
	var users []Usuario
	result := DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Error al crear el usuario.")
		return nil, result.Error
	}

	return users, nil
}
