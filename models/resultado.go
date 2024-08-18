package models

import (
	"time"

	"github.com/google/uuid"
)

type Resultado struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CPF   string    `gorm:"type:varchar(11);not null" json:"cpf"`
	EixoX string    `json:"eixo_x"`
	EixoY string    `json:"eixo_y"`
	Data  time.Time `json:"data"`
}

var Resultados []Resultado
