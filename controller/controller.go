package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/models"
)

func ExibeTodosAtletas(c *gin.Context) {
	c.JSON(200, models.Atletas)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API": "Bem vindo " + nome + ".",
	})
}
