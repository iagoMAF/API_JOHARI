package models

type Atleta struct {
	Nome   string `json:"nome"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Senha  string `json:"senha"`
	Funcao string `json:"funcao"`
	Idade  int    `json:"indade"`
	Ativo  int    `json:"Ativo"`
	// CPF_Lider string `json:"cpf_lider"`
	// ID_Equipe int    `json:"id_equipe"`
}

var Atletas []Atleta
