package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func ExibeTodasAsRespostas(c *gin.Context) {
	var respostas []models.Resposta
	database.DB.Find(&respostas)
	c.JSON(http.StatusOK, respostas)
}

func ExibeRespostaPorID(c *gin.Context) {
	var resposta models.Resposta
	id := c.Params.ByName("id")

	result := database.DB.Where("id = ?", id).First(&resposta)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Resposta não encontrada",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, resposta)
}
