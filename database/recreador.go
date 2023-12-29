package database

func CreateRecreador(recreador *Recreador) error {
	result := DB.Create(recreador)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ReadRecreador(id int) (*Recreador, error) {
	var recreador Recreador

	if err := DB.First(&recreador, id).Error; err != nil {
		return nil, err
	}

	return &recreador, nil
}

func ReadRecreadores() ([]Recreador, error) {
	var recreadores []Recreador

	if err := DB.Find(&recreadores).Error; err != nil {
		return nil, err
	}

	return recreadores, nil
}

func UpdateRecreador(id int, recreador *Recreador) error {
	var existingRecreador Recreador
	if err := DB.First(&existingRecreador, id).Error; err != nil {
		return err
	}

	existingRecreador = *recreador

	if err := DB.Save(&existingRecreador).Error; err != nil {
		return err
	}

	return nil
}
