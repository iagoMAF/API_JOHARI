package models

import "time"

// Lider representa um líder na aplicação.
// @Description Representa um líder com informações básicas.
// @ID Lider
// @Property cpf string `json:"cpf"`
// @Property nome string `json:"nome"`
// @Property email string `json:"email"`
// @Property senha string `json:"senha"`
// @Property createdAt string `json:"createdAt"`
// @Property updatedAt string `json:"updatedAt"`
// @Property deletedAt string `json:"deletedAt"`
type Lider struct {
	CPF       string     `json:"cpf" gorm:"primaryKey"` // CPF do líder
	Nome      string     `json:"nome"`                  // Nome do líder
	Email     string     `json:"email"`                 // Email do líder
	Senha     string     `json:"senha"`                 // Senha do líder
	CreatedAt time.Time  `json:"createdAt"`             // Data de criação
	UpdatedAt time.Time  `json:"updatedAt"`             // Data de atualização
	DeletedAt *time.Time `json:"deletedAt"`             // Data de exclusão (opcional)
}
