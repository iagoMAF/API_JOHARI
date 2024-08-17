package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/controller"
)

func HandleRequest() {
	r := gin.Default()

	// rotas de atleta
	r.GET("/atletas", controller.ExibeTodosAtletas)
	r.GET("/atleta/:cpf", controller.ExibeAtletaPorID)
	r.POST("/atleta", controller.CriaAtleta)
	r.PATCH("/atleta/:cpf", controller.AtualizaAtleta)

	// rotas de admin
	r.POST("/admin", controller.CriaAdmin)
	r.GET("/admin/:login", controller.ExibeAdminPorLogin)
	r.PATCH("/admin/:login", controller.AtualizaSenhaAdmin)
	r.DELETE("/admin/:login", controller.DeletaAdminPorLogin)

	// rotas de equipe
	r.POST("/equipe", controller.CriaEquipe)
	r.GET("/equipe", controller.ExibeTodasEquipes)
	r.GET("/equipe/:id", controller.ExibeEquipePorID)
	r.PATCH("/equipe/:id", controller.AtualizaEquipe)
	r.DELETE("/equipe/:id", controller.DeletaEquipePorID)

	// criar um .env para informar a porta em que a aplicação vai rodar
	r.Run(":8080")
}
