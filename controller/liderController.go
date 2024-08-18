package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func ExibeTodosLideres(c *gin.Context) {
	var lideres []models.Lider
	database.DB.Find(&lideres)
	c.JSON(http.StatusOK, lideres)
}

func CriaLider(c *gin.Context) {
	var lider models.Lider

	if err := c.ShouldBindJSON(&lider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Falha ao cadastrar líder",
			"status": 404,
		})
		return
	}

	database.DB.Create(&lider)
	c.JSON(http.StatusOK, lider)
}

func ExibeLiderPorCPF(c *gin.Context) {
	var lider models.Lider
	cpf := c.Params.ByName("cpf")

	result := database.DB.Where("cpf = ?", cpf).First(&lider)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Líder não encontrado",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, lider)
}

func AtualizaLider(c *gin.Context) {
	var lider models.Lider
	cpf := c.Params.ByName("cpf")

	if err := database.DB.Where("cpf = ?", cpf).First(&lider).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Líder não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&lider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Dados inválidos para atualização",
			"status": 400,
		})
		return
	}

	if err := database.DB.Model(&lider).Updates(lider).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar o líder"})
		return
	}

	c.JSON(http.StatusOK, lider)
}

func DeletaLider(c *gin.Context) {
	var lider models.Lider
	cpf := c.Params.ByName("cpf")

	if err := database.DB.Where("cpf = ?", cpf).First(&lider).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Líder não encontrado"})
		return
	}

	if err := database.DB.Delete(&lider).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao deletar o líder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Líder deletado com sucesso"})
}
