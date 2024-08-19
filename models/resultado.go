package models

import (
	"time"

	"github.com/google/uuid"
)

// Resultado representa o modelo de um resultado de desempenho
// @Description Resultado representa o modelo de um resultado de desempenho com informações do atleta e os dados de resultado
// @Property ID string `json:"id"` "ID único do resultado"
// @Property CPF string `json:"cpf"` "CPF do atleta"
// @Property EixoX string `json:"eixo_x"` "Valor do eixo X"
// @Property EixoY string `json:"eixo_y"` "Valor do eixo Y"
// @Property Data string `json:"data"` "Data em que o resultado foi registrado, no formato ISO 8601"
type Resultado struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CPF   string    `gorm:"type:varchar(11);not null" json:"cpf"`
	EixoX string    `json:"eixo_x"`
	EixoY string    `json:"eixo_y"`
	Data  time.Time `json:"data"`
}

var Resultados []Resultado
