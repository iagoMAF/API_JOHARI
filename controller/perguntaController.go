package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func ExibeTodasAsPerguntas(c *gin.Context) {
	var perguntas []models.Pergunta
	database.DB.Find(&perguntas)
	c.JSON(http.StatusOK, perguntas)
}

func ExibePerguntaPorID(c *gin.Context) {
	var pergunta models.Pergunta
	id := c.Params.ByName("id")

	result := database.DB.Where("id = ?", id).First(&pergunta)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Pergunta não encontrada",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, pergunta)
}

func ExibePerguntaComRespostas(c *gin.Context) {
	var pergunta models.Pergunta
	id := c.Param("id") // Obter o ID da URL

	result := database.DB.Preload("Respostas").First(&pergunta, id)
	if result.Error != nil {

		fmt.Printf("Erro ao buscar pergunta com ID %s: %v\n", id, result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Pergunta não encontrada",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, pergunta)
}

func ExibePerguntasComRespostas(c *gin.Context) {
	var perguntas []models.Pergunta

	// Buscar todas as perguntas com suas respostas associadas
	result := database.DB.Preload("Respostas").Find(&perguntas)
	if result.Error != nil {
		// Adicione o log do erro para debugging
		fmt.Printf("Erro ao buscar todas as perguntas: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Erro ao buscar perguntas",
			"status": 500,
		})
		return
	}

	c.JSON(http.StatusOK, perguntas)
}
