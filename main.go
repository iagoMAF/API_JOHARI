package main

import (
	"github.com/iagoMAF/API_JOHARI/models"
	"github.com/iagoMAF/API_JOHARI/routes"
)

func main() {
	models.Atletas = []models.Atleta{
		{Nome: "Iago", CPF: "123", Email: "teste@gmail.com", Senha: "123", Funcao: "Atacante", Idade: 18, Ativo: 1},
		{Nome: "Gabi", CPF: "321", Email: "gabi@gmail.com", Senha: "123", Funcao: "Atacante", Idade: 18, Ativo: 0},
	}

	routes.HandleRequest()
}
