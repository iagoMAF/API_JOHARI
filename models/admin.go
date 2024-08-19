package models

import (
	"github.com/google/uuid"
)

// Admin godoc
// @Description Representa um administrador do sistema
// @Property ID string `json:"-"` "Identificador único do administrador"
// @Property Login string `json:"login"` "Login do administrador"
// @Property Senha string `json:"senha"` "Senha do administrador"
type Admin struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"-"`
	Login string    `json:"login"`
	Senha string    `json:"senha"`
}

var Admins []Admin
