package database

import "errors"

func CreateTrabajador(trabajador *Trabajador) error {
	if trabajador.UsuarioID == 0 {
		return errors.New("the userID must not be 0")
	}

	result := DB.Create(trabajador)
	if result.Error != nil {
		return result.Error
	}

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

	if err := DB.Find(&trabajadores).Error; err != nil {
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
