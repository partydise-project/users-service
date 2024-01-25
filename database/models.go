package database

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Email               string `json:"email"`
	Auth0ID             string `json:"auth0_id"`
	NumeroCelular       uint64 `json:"numero_celular"`
	DireccionResidencia string `json:"direccion_residencia"`
	FechaNacimiento     string `json:"fecha_nacimiento"`
	Genero              string `json:"genero"`
	Nombre              string `json:"nombre"`
}

type Recreador struct {
	gorm.Model
	UsuarioID            int     `json:"usuario_ID" gorm:"not null"`
	Usuario              Usuario `json:"usuario"`
	ExperienciaRecreando string  `json:"experiencia_recreando"`
	Dotacion             bool    `json:"dotacion"`
	NivelRecreacion      string  `json:"nivel_recreacion"`
	NivelDecoracion      string  `json:"nivel_decoracion"`
	NivelMagia           string  `json:"nivel_magia"`
}

type Cliente struct {
	gorm.Model
	UsuarioID  int     `json:"usuario_ID" gorm:"not null"`
	Usuario    Usuario `json:"usuario"`
	Antiguedad string  `json:"antiguedad"`
}

type Trabajador struct {
	gorm.Model
	UsuarioID uint    `json:"usuario_ID" gorm:"not null"`
	Usuario   Usuario `json:"usuario"`
	Cargo     string  `json:"cargo"`
}
