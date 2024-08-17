package models

import (
	"github.com/google/uuid"
)

type Admin struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"-"`
	Login string    `json:"login"`
	Senha string    `json:"senha"`
}

var Admins []Admin
