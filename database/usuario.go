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

	if err := DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func ReadUsuarios() ([]Usuario, error) {
	var users []Usuario

	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
