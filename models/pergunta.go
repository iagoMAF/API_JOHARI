package models

type Pergunta struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Pergunta  string     `json:"pergunta"`
	Respostas []Resposta `gorm:"foreignKey:PerguntaID" json:"respostas"`
}

var Perguntas []Pergunta
