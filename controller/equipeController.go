package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func CriaEquipe(c *gin.Context) {
	var equipe models.Equipe

	if err := c.ShouldBindJSON(&equipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Falha ao cadastrar equipe",
			"status": 400,
		})
		return
	}

	if err := database.DB.Create(&equipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao salvar equipe no banco de dados",
		})
		return
	}

	c.JSON(http.StatusOK, equipe)
}

func ExibeTodasEquipes(c *gin.Context) {
	var equipes []models.Equipe
	database.DB.Find(&equipes)
	c.JSON(200, equipes)
}

func ExibeEquipePorID(c *gin.Context) {
	var equipe models.Equipe
	id := c.Params.ByName("id")

	result := database.DB.First(&equipe, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Equipe não encontrada",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, equipe)
}

func AtualizaEquipe(c *gin.Context) {
	var equipe models.Equipe
	id := c.Params.ByName("id")

	result := database.DB.First(&equipe, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Equipe não encontrada",
			"status": 404,
		})
		return
	}

	if err := c.ShouldBindJSON(&equipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Dados inválidos para atualização",
			"status": 400,
		})
		return
	}

	if err := database.DB.Save(&equipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao atualizar a equipe",
		})
		return
	}

	c.JSON(http.StatusOK, equipe)
}

func DeletaEquipePorID(c *gin.Context) {
	var equipe models.Equipe
	id := c.Params.ByName("id")

	result := database.DB.First(&equipe, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Equipe não encontrada",
			"status": 404,
		})
		return
	}

	if err := database.DB.Delete(&equipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao deletar a equipe",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Equipe deletada com sucesso",
	})
}
