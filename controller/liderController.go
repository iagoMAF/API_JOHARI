package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// ExibeTodosLideres godoc
// @Summary Exibe todos os líderes
// @Description Retorna uma lista de todos os líderes cadastrados
// @Tags Líderes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Lider
// @Router /lideres [get]
func ExibeTodosLideres(c *gin.Context) {
	var lideres []models.Lider
	database.DB.Find(&lideres)
	c.JSON(http.StatusOK, lideres)
}

// CriaLider godoc
// @Summary Cria um novo líder
// @Description Cadastra um novo líder no sistema
// @Tags Líderes
// @Accept  json
// @Produce  json
// @Param lider body models.Lider true "Dados do líder"
// @Success 200 {object} models.Lider
// @Failure 400 {object} map[string]interface{} "Falha ao cadastrar líder"
// @Router /lider [post]
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

// ExibeLiderPorCPF godoc
// @Summary Exibe um líder por CPF
// @Description Busca um líder pelo CPF informado
// @Tags Líderes
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do líder"
// @Success 200 {object} models.Lider
// @Failure 404 {object} map[string]interface{} "Líder não encontrado"
// @Router /lider/{cpf} [get]
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

// AtualizaLider godoc
// @Summary Atualiza um líder
// @Description Atualiza os dados de um líder existente pelo CPF
// @Tags Líderes
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do líder"
// @Param lider body models.Lider true "Dados atualizados do líder"
// @Success 200 {object} models.Lider
// @Failure 400 {object} map[string]interface{} "Dados inválidos para atualização"
// @Failure 404 {object} map[string]interface{} "Líder não encontrado"
// @Failure 500 {object} map[string]interface{} "Falha ao atualizar o líder"
// @Router /lider/{cpf} [patch]
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

// DeletaLider godoc
// @Summary Deleta um líder por CPF
// @Description Remove um líder do sistema pelo CPF informado
// @Tags Líderes
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF do líder"
// @Success 200 {object} map[string]interface{} "Líder deletado com sucesso"
// @Failure 404 {object} map[string]interface{} "Líder não encontrado"
// @Failure 500 {object} map[string]interface{} "Falha ao deletar o líder"
// @Router /lider/{cpf} [delete]
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
