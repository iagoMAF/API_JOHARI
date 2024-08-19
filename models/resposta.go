package models

// Resposta representa uma resposta para uma pergunta.
// @Description Representa uma resposta associada a uma pergunta.
// @ID Resposta
// @Property id uint `json:"id"`
// @Property pergunta_id uint `json:"pergunta_id"`
// @Property resposta string `json:"resposta"`
type Resposta struct {
	ID         uint   `gorm:"primaryKey" json:"id"` // ID da resposta
	PerguntaID uint   `json:"pergunta_id"`          // ID da pergunta associada
	Resposta   string `json:"resposta"`             // Texto da resposta
}

var Respostas []Resposta
