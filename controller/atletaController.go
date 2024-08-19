package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// ExibeTodosAtletas godoc
// @Summary Exibe todos os atletas
// @Description Retorna uma lista de todos os atletas cadastrados
// @Tags Atletas
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Atleta
// @Router /atletas [get]
func ExibeTodosAtletas(c *gin.Context) {
	var atletas []models.Atleta
	database.DB.Find(&atletas)
	c.JSON(http.StatusOK, atletas)
}

// CriaAtleta godoc
// @Summary Cria um novo atleta
// @Description Cadastra um novo atleta no sistema
// @Tags Atletas
// @Accept  json
// @Produce  json
// @Param atleta body models.Atleta true "Dados do atleta"
// @Success 200 {object} models.Atleta
// @Failure 400 {object} models.Atleta "Falha ao cadastrar atleta"
// @Router /atleta [post]
func CriaAtleta(c *gin.Context) {
	var atleta models.Atleta

	if err := c.ShouldBindJSON(&atleta); err != nil {
		c.JSON(http.StatusBadRequest, models.Atleta{
			Nome:  "Erro",
			Email: err.Error(),
		})
		return
	}

	database.DB.Create(&atleta)
	c.JSON(http.StatusOK, atleta)
}

// ExibeAtletaPorID godoc
// @Summary Exibe um atleta por CPF
// @Description Busca um atleta pelo CPF informado
// @Tags Atletas
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do atleta"
// @Success 200 {object} models.Atleta
// @Failure 404 {object} models.Atleta "Atleta não encontrado"
// @Router /atleta/{cpf} [get]
func ExibeAtletaPorID(c *gin.Context) {
	var atleta models.Atleta
	cpf := c.Param("cpf")

	result := database.DB.Preload("Lider").Preload("Equipe").Where("cpf = ?", cpf).First(&atleta)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.Atleta{
			Nome:  "Erro",
			Email: "Atleta não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, atleta)
}

// AtualizaAtleta godoc
// @Summary Atualiza um atleta
// @Description Atualiza os dados de um atleta existente pelo CPF
// @Tags Atletas
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do atleta"
// @Param atleta body models.Atleta true "Dados atualizados do atleta"
// @Success 200 {object} models.Atleta
// @Failure 400 {object} models.Atleta "Dados inválidos para atualização"
// @Failure 404 {object} models.Atleta "Atleta não encontrado"
// @Router /atleta/{cpf} [patch]
func AtualizaAtleta(c *gin.Context) {
	var atleta models.Atleta
	cpf := c.Param("cpf")

	if err := database.DB.Where("cpf = ?", cpf).First(&atleta).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Atleta{
			Nome:  "Erro",
			Email: "Atleta não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&atleta); err != nil {
		c.JSON(http.StatusBadRequest, models.Atleta{
			Nome:  "Erro",
			Email: "Dados inválidos para atualização",
		})
		return
	}

	if err := database.DB.Model(&atleta).Updates(atleta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Atleta{
			Nome:  "Erro",
			Email: "Falha ao atualizar o atleta",
		})
		return
	}

	c.JSON(http.StatusOK, atleta)
}
