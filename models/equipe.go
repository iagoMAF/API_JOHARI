package models

import "time"

// Equipe representa uma equipe na aplicação.
// @Description Representa uma equipe com informações básicas.
// @ID Equipe
// @Property id integer `json:"id"`
// @Property createdAt string `json:"createdAt"`
// @Property updatedAt string `json:"updatedAt"`
// @Property deletedAt string `json:"deletedAt"`
// @Property nome string `json:"nome"`
// @Property ativo integer `json:"ativo"`
type Equipe struct {
	ID        uint       `json:"id" gorm:"primaryKey"` // ID da equipe
	CreatedAt time.Time  `json:"createdAt"`            // Data de criação
	UpdatedAt time.Time  `json:"updatedAt"`            // Data de atualização
	DeletedAt *time.Time `json:"deletedAt"`            // Data de exclusão (opcional)
	Nome      string     `json:"nome"`                 // Nome da equipe
	Ativo     int        `json:"ativo"`                // Status ativo da equipe
}
