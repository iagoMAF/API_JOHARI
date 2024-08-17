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

	// rota de admin
	r.POST("/admin", controller.CriaAdmin)
	r.GET("/admin/:login", controller.ExibeAdminPorLogin)
	r.PATCH("/admin/:login", controller.AtualizaSenhaAdmin)
	r.DELETE("/admin/:login", controller.DeletaAdminPorLogin)

	// criar um .env para informar a porta em que a aplicação vai rodar
	r.Run(":8080")
}
