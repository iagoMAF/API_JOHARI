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
	ID_Equipe *uint   `json:"id_equipe" gorm:"index"`
	Lider     Lider   `gorm:"foreignKey:CPF_Lider;references:CPF"`
	Equipe    Equipe  `gorm:"foreignKey:ID_Equipe;references:ID"`
}

var Atletas []Atleta
