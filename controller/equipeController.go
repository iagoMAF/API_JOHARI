package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// CriaEquipe godoc
// @Summary Cria uma nova equipe
// @Description Cadastra uma nova equipe no sistema
// @Tags Equipes
// @Accept  json
// @Produce  json
// @Param equipe body models.Equipe true "Dados da equipe"
// @Success 200 {object} models.Equipe
// @Failure 400 {object} map[string]interface{} "Falha ao cadastrar equipe"
// @Failure 500 {object} map[string]interface{} "Falha ao salvar equipe no banco de dados"
// @Router /equipe [post]
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

// ExibeTodasEquipes godoc
// @Summary Exibe todas as equipes
// @Description Retorna uma lista de todas as equipes cadastradas
// @Tags Equipes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Equipe
// @Router /equipes [get]
func ExibeTodasEquipes(c *gin.Context) {
	var equipes []models.Equipe
	database.DB.Find(&equipes)
	c.JSON(200, equipes)
}

// ExibeEquipePorID godoc
// @Summary Exibe uma equipe por ID
// @Description Busca uma equipe pelo ID informado
// @Tags Equipes
// @Accept  json
// @Produce  json
// @Param id path string true "ID da equipe"
// @Success 200 {object} models.Equipe
// @Failure 404 {object} map[string]interface{} "Equipe não encontrada"
// @Router /equipe/{id} [get]
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

// AtualizaEquipe godoc
// @Summary Atualiza uma equipe
// @Description Atualiza os dados de uma equipe existente pelo ID
// @Tags Equipes
// @Accept  json
// @Produce  json
// @Param id path string true "ID da equipe"
// @Param equipe body models.Equipe true "Dados atualizados da equipe"
// @Success 200 {object} models.Equipe
// @Failure 400 {object} map[string]interface{} "Dados inválidos para atualização"
// @Failure 404 {object} map[string]interface{} "Equipe não encontrada"
// @Failure 500 {object} map[string]interface{} "Falha ao atualizar a equipe"
// @Router /equipe/{id} [patch]
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

// DeletaEquipePorID godoc
// @Summary Deleta uma equipe por ID
// @Description Remove uma equipe do sistema pelo ID informado
// @Tags Equipes
// @Accept  json
// @Produce  json
// @Param id path string true "ID da equipe"
// @Success 200 {object} map[string]interface{} "Equipe deletada com sucesso"
// @Failure 404 {object} map[string]interface{} "Equipe não encontrada"
// @Failure 500 {object} map[string]interface{} "Falha ao deletar a equipe"
// @Router /equipe/{id} [delete]
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
