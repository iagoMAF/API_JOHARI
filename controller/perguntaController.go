package controller

import (
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
