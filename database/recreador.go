package database

import "errors"

func CreateRecreador(recreador *Recreador) error {
	if recreador.UsuarioID == 0 {
		return errors.New("the userID must not be 0")
	}

	result := DB.Create(recreador)
	if result.Error != nil {
		return result.Error
	}
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
