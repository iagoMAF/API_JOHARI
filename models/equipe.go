package models

import "gorm.io/gorm"

type Equipe struct {
	gorm.Model
	Nome  string `json:"nome"`
	Ativo int    `json:"ativo"`
}

var Equipes []Equipe
