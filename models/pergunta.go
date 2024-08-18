package models

type Pergunta struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Pergunta string `json:"pergunta"`
}

var Perguntas []Pergunta
