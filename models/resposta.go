package models

type Resposta struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	PerguntaID uint   `json:"pergunta_id"`
	Resposta   string `json:"resposta"`
}

var Respostas []Resposta
