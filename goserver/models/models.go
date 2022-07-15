package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

type DBModel struct {
	DB *sql.DB
}

// NewModels return models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Nacimiento struct {
	Mes  int `json:"mes"`
	Anio int `json:"anio"`
}

type Usuario struct {
	UsuarioId int64  `json:"usuarioid"`
	Email     string `json:"email"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	// FechaRegistro string            `json:"fecharegistro"`
	Pass   string `json:"pass"`
	PaisId string `json:"paisid"`
	// Grupos        []GrupoIntegrante `json:"grupos"`
}

type Auditoria struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UsuCreate string    `json:"usu_create"`
	UsuUpdate string    `json:"usuUpdate"`
}

type Ramas struct {
	Lobatos bool `json:"lobatos"`
	Scouts  bool `json:"scouts"`
	Raiders bool `json:"raiders"`
	Rovers  bool `json:"rovers"`
}

type Grupo struct {
	GrupoId int64  `json:"id"`
	Nombre  string `json:"nombre"`
	// Anio            int       `json:"anio"`
	// Audit           Auditoria `json:"audit"`
	// RamasUtilizadas Ramas     `json:"ramas"`
}

type GrupoIntegrante struct {
	GrupoId   int64 `json:"grupoId"`
	UsuarioId int64 `json:"usuarioId"`
}

type Pais struct {
	PaisId string `json:"paisId"`
	Nombre string `json:"nombre"`
}
