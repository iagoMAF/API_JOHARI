package models

import "gorm.io/gorm"

type Lider struct {
	gorm.Model
	CPF   string `gorm:"primaryKey" json:"cpf"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

var Lideres []Lider
