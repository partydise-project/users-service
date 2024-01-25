package database

func CreateRecreador(user *Usuario, recreador *Recreador) error {
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

	recreador.UsuarioID = user.ID

	if err := tx.Create(recreador).Error; err != nil {
		// Revertir la transacción en caso de error
		tx.Rollback()
		return err
	}

	// Commit de la transacción si todo ha ido bien
	tx.Commit()
	return nil
}

func ReadRecreador(id string) (*Recreador, error) {
	var recreador Recreador

	if err := DB.Preload("Usuario").First(&recreador, id).Error; err != nil {
		return nil, err
	}

	return &recreador, nil
}

func ReadRecreadores() ([]Recreador, error) {
	var recreadores []Recreador

	if err := DB.Preload("Usuario").Find(&recreadores).Error; err != nil {
		return nil, err
	}

	return recreadores, nil
}

func UpdateRecreador(recreador *Recreador) error {
	if err := DB.Save(&recreador).Error; err != nil {
		return err
	}

	return nil
}
