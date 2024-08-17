package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/controller"
)

func HandleRequest() {
	r := gin.Default()

	// rotas do atleta
	r.GET("/atletas", controller.ExibeTodosAtletas)
	r.GET("/atleta/:cpf", controller.ExibeAtletaPorID)
	r.POST("/atleta", controller.CriaAtleta)
	r.PATCH("/atleta/:cpf", controller.AtualizaAtleta)

	// criar um .env para informar a porta em que a aplicação vai rodar
	r.Run(":8080")
}
