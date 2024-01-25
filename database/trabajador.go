package database

func CreateTrabajador(user *Usuario, trabajador *Trabajador) error {
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

	trabajador.UsuarioID = user.ID

	if err := tx.Create(trabajador).Error; err != nil {
		// Revertir la transacción en caso de error
		tx.Rollback()
		return err
	}

	// Commit de la transacción si todo ha ido bien
	tx.Commit()
	return nil
}

func ReadTrabajador(id string) (*Trabajador, error) {
	var trabajador Trabajador

	if err := DB.Preload("Usuario").First(&trabajador, id).Error; err != nil {
		return nil, err
	}

	return &trabajador, nil
}

func ReadTrabajadores() ([]Trabajador, error) {
	var trabajadores []Trabajador

	if err := DB.Preload("Usuario").Find(&trabajadores).Error; err != nil {
		return nil, err
	}

	return trabajadores, nil
}

func UpdateTrabajador(trabajador *Trabajador) error {
	if err := DB.Save(&trabajador).Error; err != nil {
		return err
	}

	return nil
}
