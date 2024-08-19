package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// ExibeTodasAsPerguntas godoc
// @Summary Exibe todas as perguntas
// @Description Retorna uma lista de todas as perguntas cadastradas
// @Tags Perguntas
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Pergunta
// @Router /perguntas [get]
func ExibeTodasAsPerguntas(c *gin.Context) {
	var perguntas []models.Pergunta
	if err := database.DB.Find(&perguntas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar perguntas"})
		return
	}
	c.JSON(http.StatusOK, perguntas)
}

// ExibePerguntaPorID godoc
// @Summary Exibe uma pergunta por ID
// @Description Busca uma pergunta pelo ID informado
// @Tags Perguntas
// @Accept  json
// @Produce  json
// @Param id path int true "ID da pergunta"
// @Success 200 {object} models.Pergunta
// @Failure 404 {object} map[string]interface{} "Pergunta não encontrada"
// @Router /pergunta/{id} [get]
func ExibePerguntaPorID(c *gin.Context) {
	var pergunta models.Pergunta
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&pergunta).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pergunta não encontrada"})
		return
	}
	c.JSON(http.StatusOK, pergunta)
}

// ExibePerguntaComRespostas godoc
// @Summary Exibe uma pergunta com suas respostas
// @Description Busca uma pergunta com todas as respostas associadas pelo ID informado
// @Tags Perguntas
// @Accept  json
// @Produce  json
// @Param id path int true "ID da pergunta"
// @Success 200 {object} models.Pergunta
// @Failure 404 {object} map[string]interface{} "Pergunta não encontrada"
// @Router /pergunta/{id}/respostas [get]
func ExibePerguntaComRespostas(c *gin.Context) {
	var pergunta models.Pergunta
	id := c.Param("id")

	if err := database.DB.Preload("Respostas").First(&pergunta, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pergunta não encontrada"})
		return
	}
	c.JSON(http.StatusOK, pergunta)
}

// ExibePerguntasComRespostas godoc
// @Summary Exibe todas as perguntas com suas respostas
// @Description Busca todas as perguntas com suas respostas associadas
// @Tags Perguntas
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Pergunta
// @Failure 500 {object} map[string]interface{} "Erro ao buscar perguntas"
// @Router /perguntas/respostas [get]
func ExibePerguntasComRespostas(c *gin.Context) {
	var perguntas []models.Pergunta

	if err := database.DB.Preload("Respostas").Find(&perguntas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar perguntas"})
		return
	}
	c.JSON(http.StatusOK, perguntas)
}
