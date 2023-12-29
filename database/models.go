package database

import "gorm.io/gorm"

type genero int

const (
	Hombre genero = iota
	Mujer
)

type nivelHabilidad int

const (
	Basico nivelHabilidad = iota
	Intermedio
	Experto
)

type Usuario struct {
	gorm.Model
	Email               string `json:"email"`
	Auth0ID             string `json:"auth0_id"`
	NumeroCelular       uint64 `json:"numero_celular"`
	DireccionResidencia string `json:"direccion_residencia"`
	FechaNacimiento     string `json:"fecha_nacimiento"`
	Genero              genero `json:"genero"`
	Nombre              string `json:"nombre"`
}

type Recreador struct {
	gorm.Model
	UsuarioID            int            `json:"usuario_ID" gorm:"not null"`
	Usuario              Usuario        `json:"usuario"`
	ExperienciaRecreando string         `json:"experiencia_recreando"`
	Dotacion             bool           `json:"dotacion"`
	NivelRecreacion      nivelHabilidad `json:"nivel_recreacion"`
	NivelDecoracion      nivelHabilidad `json:"nivel_decoracion"`
	NivelMagia           nivelHabilidad `json:"nivel_magia"`
}

type Cliente struct {
	gorm.Model
	UsuarioID  int     `json:"usuario_ID" gorm:"not null"`
	Usuario    Usuario `json:"usuario"`
	Antiguedad string  `json:"antiguedad"`
}

type Trabajador struct {
	gorm.Model
	UsuarioID int     `json:"usuario_ID" gorm:"not null"`
	Usuario   Usuario `json:"usuario"`
	Cargo     string  `json:"cargo"`
}
