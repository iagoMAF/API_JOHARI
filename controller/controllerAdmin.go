package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

func CriaAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Falha ao cadastrar administrador",
			"status": 400,
		})
		return
	}

	// Gera um novo UUID para o administrador
	admin.ID = uuid.New()

	if err := database.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao salvar administrador no banco de dados",
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func ExibeAdminPorLogin(c *gin.Context) {
	var admin models.Admin
	login := c.Params.ByName("login")

	result := database.DB.Where("login = ?", login).First(&admin)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Administrador não encontrado",
			"status": 404,
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func AtualizaSenhaAdmin(c *gin.Context) {
	var admin models.Admin
	login := c.Params.ByName("login")

	result := database.DB.Where("login = ?", login).First(&admin)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Administrador não encontrado",
			"status": 404,
		})
		return
	}

	var novaSenha struct {
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&novaSenha); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Dados inválidos para atualização",
			"status": 400,
		})
		return
	}

	admin.Senha = novaSenha.Senha
	if err := database.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao atualizar a senha do administrador",
		})
		return
	}

	c.JSON(http.StatusOK, admin)
}

func DeletaAdminPorLogin(c *gin.Context) {
	var admin models.Admin
	login := c.Params.ByName("login")

	// Verifica se o administrador existe usando o Login
	result := database.DB.Where("login = ?", login).First(&admin)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Administrador não encontrado",
			"status": 404,
		})
		return
	}

	// Deleta o administrador
	if err := database.DB.Delete(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao deletar o administrador",
		})
		return
	}

	// Retorna uma resposta de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": "Administrador deletado com sucesso",
	})
}
