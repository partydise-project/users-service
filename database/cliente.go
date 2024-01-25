package database

func CreateCliente(user *Usuario, cliente *Cliente) error {
	// Iniciar la transacción
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(user).Error; err != nil {
		// Revertir la transacción en caso de error
		tx.Rollback()
		return err
	}

	cliente.UsuarioID = user.ID

	if err := tx.Create(cliente).Error; err != nil {
		// Revertir la transacción en caso de error
		tx.Rollback()
		return err
	}

	// Commit de la transacción si todo ha ido bien
	tx.Commit()
	return nil
}

func ReadCliente(id string) (*Cliente, error) {
	var cliente Cliente

	if err := DB.Preload("Usuario").First(&cliente, id).Error; err != nil {
		return nil, err
	}

	return &cliente, nil
}

func ReadClientes() ([]Cliente, error) {
	var cliente []Cliente

	if err := DB.Preload("Usuario").Find(&cliente).Error; err != nil {
		return nil, err
	}

	return cliente, nil
}

func UpdateCliente(cliente *Cliente) error {
	if err := DB.Save(&cliente).Error; err != nil {
		return err
	}

	return nil
}
