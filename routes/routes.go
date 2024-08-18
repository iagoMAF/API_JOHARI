package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/controller"
	"github.com/joho/godotenv"
)

func HandleRequest() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// rotas de atleta
	r.GET("/atletas", controller.ExibeTodosAtletas)
	r.GET("/atleta/:cpf", controller.ExibeAtletaPorID)
	r.PATCH("/atleta/:cpf", controller.AtualizaAtleta)
	r.POST("/atleta", controller.CriaAtleta)

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

	// rotas de lider
	r.POST("/lider", controller.CriaLider)
	r.GET("/lider", controller.ExibeTodosLideres)
	r.GET("/lider/:cpf", controller.ExibeLiderPorCPF)
	r.PATCH("/lider/:cpf", controller.AtualizaLider)
	r.DELETE("/lider/:cpf", controller.DeletaLider)

	r.Run(":" + port)
}
