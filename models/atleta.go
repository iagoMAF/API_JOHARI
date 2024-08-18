package models

import "gorm.io/gorm"

type Atleta struct {
	gorm.Model
	Nome      string  `json:"nome"`
	CPF       string  `json:"cpf"`
	Email     string  `json:"email"`
	Senha     string  `json:"senha"`
	Funcao    string  `json:"funcao"`
	Idade     int     `json:"indade"`
	Ativo     int     `json:"ativo"`
	CPF_Lider *string `json:"cpf_lider" gorm:"index"`
	ID_Equipe int     `json:"id_equipe"`
	Lider     Lider   `gorm:"foreignKey:CPF_Lider;references:CPF"`
}

var Atletas []Atleta
