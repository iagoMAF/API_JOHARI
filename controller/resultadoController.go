package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// CriaResultado godoc
// @Summary Cria um novo resultado
// @Description Cadastra um novo resultado no sistema
// @Tags Resultados
// @Accept  json
// @Produce  json
// @Param resultado body models.Resultado true "Dados do resultado"
// @Success 200 {object} models.Resultado
// @Failure 400 {object} map[string]interface{} "Falha ao cadastrar resultado"
// @Failure 500 {object} map[string]interface{} "Falha ao criar resultado"
// @Router /resultado [post]
func CriaResultado(c *gin.Context) {
	var resultado models.Resultado

	if err := c.ShouldBindJSON(&resultado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Falha ao cadastrar resultado",
			"status": 400,
		})
		return
	}

	if resultado.Data.IsZero() {
		resultado.Data = time.Now()
	}

	if err := database.DB.Create(&resultado).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Falha ao criar resultado",
			"status": 500,
		})
		return
	}

	c.JSON(http.StatusOK, resultado)
}

// ExibeResultadosPorCPF godoc
// @Summary Exibe resultados por CPF
// @Description Busca todos os resultados associados ao CPF informado
// @Tags Resultados
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do atleta"
// @Success 200 {array} models.Resultado
// @Failure 404 {object} map[string]interface{} "Resultados não encontrados"
// @Router /resultados/{cpf} [get]
func ExibeResultadosPorCPF(c *gin.Context) {
	var resultados []models.Resultado
	cpf := c.Param("cpf")

	result := database.DB.Where("cpf = ?", cpf).Find(&resultados)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Resultados não encontrados",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, resultados)
}
