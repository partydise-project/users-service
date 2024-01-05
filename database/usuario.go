package database

func CreateUsuario(user *Usuario) error {
	result := DB.Create(user)
	if result.Error != nil {
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

func UpdateUser(id int, user *Usuario) error {
	var existingUser Usuario
	if err := DB.First(&existingUser, id).Error; err != nil {
		return err
	}

	existingUser = *user

	if err := DB.Save(&existingUser).Error; err != nil {
		return err
	}

	return nil
}
