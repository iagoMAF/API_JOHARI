package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

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
