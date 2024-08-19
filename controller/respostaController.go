package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// ExibeTodasAsRespostas godoc
// @Summary Exibe todas as respostas
// @Description Retorna uma lista de todas as respostas cadastradas
// @Tags Respostas
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Resposta
// @Router /respostas [get]
func ExibeTodasAsRespostas(c *gin.Context) {
	var respostas []models.Resposta
	if err := database.DB.Find(&respostas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar respostas"})
		return
	}
	c.JSON(http.StatusOK, respostas)
}

// ExibeRespostaPorID godoc
// @Summary Exibe uma resposta por ID
// @Description Busca uma resposta pelo ID informado
// @Tags Respostas
// @Accept  json
// @Produce  json
// @Param id path int true "ID da resposta"
// @Success 200 {object} models.Resposta
// @Failure 404 {object} map[string]interface{} "Resposta não encontrada"
// @Router /resposta/{id} [get]
func ExibeRespostaPorID(c *gin.Context) {
	var resposta models.Resposta
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&resposta).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resposta não encontrada"})
		return
	}
	c.JSON(http.StatusOK, resposta)
}
