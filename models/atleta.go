package models

import (
	"time"
)

// Atleta representa um atleta na aplicação.
// @Description Representa um atleta com informações básicas.
// @ID Atleta
// @Property id integer `json:"id"`
// @Property createdAt string `json:"createdAt"`
// @Property updatedAt string `json:"updatedAt"`
// @Property deletedAt string `json:"deletedAt"`
// @Property nome string `json:"nome"`
// @Property cpf string `json:"cpf"`
// @Property email string `json:"email"`
// @Property senha string `json:"senha"`
// @Property funcao string `json:"funcao"`
// @Property idade integer `json:"idade"`
// @Property ativo integer `json:"ativo"`
// @Property cpf_lider string `json:"cpf_lider"`
// @Property id_equipe integer `json:"id_equipe"`
type Atleta struct {
	ID        uint       `json:"id" gorm:"primaryKey"`                // ID do atleta
	CreatedAt time.Time  `json:"createdAt"`                           // Data de criação
	UpdatedAt time.Time  `json:"updatedAt"`                           // Data de atualização
	DeletedAt *time.Time `json:"deletedAt"`                           // Data de exclusão (opcional)
	Nome      string     `json:"nome"`                                // Nome do atleta
	CPF       string     `json:"cpf"`                                 // CPF do atleta
	Email     string     `json:"email"`                               // Email do atleta
	Senha     string     `json:"senha"`                               // Senha do atleta
	Funcao    string     `json:"funcao"`                              // Função do atleta
	Idade     int        `json:"idade"`                               // Idade do atleta
	Ativo     int        `json:"ativo"`                               // Status ativo do atleta
	CPF_Lider *string    `json:"cpf_lider" gorm:"index"`              // CPF do líder associado
	ID_Equipe *uint      `json:"id_equipe" gorm:"index"`              // ID da equipe associada
	Lider     Lider      `gorm:"foreignKey:CPF_Lider;references:CPF"` // Relacionamento com o modelo Lider
	Equipe    Equipe     `gorm:"foreignKey:ID_Equipe;references:ID"`  // Relacionamento com o modelo Equipe
}
