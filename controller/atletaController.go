package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func ExibeTodosAtletas(c *gin.Context) {
	var atletas []models.Atleta
	database.DB.Find(&atletas)
	c.JSON(200, atletas)
}

func CriaAtleta(c *gin.Context) {
	var atleta models.Atleta

	if err := c.ShouldBindJSON(&atleta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Falha ao cadastrar atleta",
			"status": 404,
		})
		return
	}

	database.DB.Create(&atleta)
	c.JSON(http.StatusOK, atleta)
}

func ExibeAtletaPorID(c *gin.Context) {
	var atleta models.Atleta
	cpf := c.Params.ByName("cpf")

	result := database.DB.Where("cpf = ?", cpf).First(&atleta)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Atleta não encontrado",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, atleta)
}

func AtualizaAtleta(c *gin.Context) {
	var atleta models.Atleta
	cpf := c.Params.ByName("cpf")

	if err := database.DB.Where("cpf = ?", cpf).First(&atleta).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Atleta não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&atleta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Dados inválidos para atualização",
			"status": 400,
		})
		return
	}

	if err := database.DB.Model(&atleta).Updates(atleta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar o atleta"})
		return
	}

	c.JSON(http.StatusOK, atleta)
}
