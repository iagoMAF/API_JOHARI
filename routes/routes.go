package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/controller"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/atletas", controller.ExibeTodosAtletas)

	// criar um .env para informar a porta em que a aplicação vai rodar
	r.Run(":8080")
}
