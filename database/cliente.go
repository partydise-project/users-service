package database

import "errors"

func CreateCliente(cliente *Cliente) error {
	if cliente.UsuarioID == 0 {
		return errors.New("the userID must not be 0")
	}

	result := DB.Create(cliente)
	if result.Error != nil {
		return result.Error
	}

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
