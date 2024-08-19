package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iagoMAF/API_JOHARI/database"
	"github.com/iagoMAF/API_JOHARI/models"
)

// CriaAdmin godoc
// @Summary Cria um novo administrador
// @Description Cadastra um novo administrador no sistema
// @Tags Administradores
// @Accept  json
// @Produce  json
// @Param admin body models.Admin true "Dados do administrador"
// @Success 200 {object} models.Admin
// @Failure 400 {object} map[string]interface{} "Falha ao cadastrar administrador"
// @Failure 500 {object} map[string]interface{} "Falha ao salvar administrador no banco de dados"
// @Router /admin [post]
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

// ExibeAdminPorLogin godoc
// @Summary Exibe um administrador por login
// @Description Busca um administrador pelo login informado
// @Tags Administradores
// @Accept  json
// @Produce  json
// @Param login path string true "Login do administrador"
// @Success 200 {object} models.Admin
// @Failure 404 {object} map[string]interface{} "Administrador não encontrado"
// @Router /admin/{login} [get]
func ExibeAdminPorLogin(c *gin.Context) {
	var admin models.Admin
	login := c.Param("login")

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

// AtualizaSenhaAdmin godoc
// @Summary Atualiza a senha de um administrador
// @Description Atualiza a senha de um administrador existente pelo login
// @Tags Administradores
// @Accept  json
// @Produce  json
// @Param login path string true "Login do administrador"
// @Param senha body string true "Nova senha"
// @Success 200 {object} models.Admin
// @Failure 400 {object} map[string]interface{} "Dados inválidos para atualização"
// @Failure 404 {object} map[string]interface{} "Administrador não encontrado"
// @Failure 500 {object} map[string]interface{} "Falha ao atualizar a senha do administrador"
// @Router /admin/{login}/senha [patch]
func AtualizaSenhaAdmin(c *gin.Context) {
	var admin models.Admin
	login := c.Param("login")

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

// DeletaAdminPorLogin godoc
// @Summary Deleta um administrador por login
// @Description Remove um administrador do sistema pelo login informado
// @Tags Administradores
// @Accept  json
// @Produce  json
// @Param login path string true "Login do administrador"
// @Success 200 {object} map[string]interface{} "Administrador deletado com sucesso"
// @Failure 404 {object} map[string]interface{} "Administrador não encontrado"
// @Failure 500 {object} map[string]interface{} "Falha ao deletar o administrador"
// @Router /admin/{login} [delete]
func DeletaAdminPorLogin(c *gin.Context) {
	var admin models.Admin
	login := c.Param("login")

	result := database.DB.Where("login = ?", login).First(&admin)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "Administrador não encontrado",
			"status": 404,
		})
		return
	}

	if err := database.DB.Delete(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao deletar o administrador",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Administrador deletado com sucesso",
	})
}
