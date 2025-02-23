package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iagoMAF/API_JOHARI/controller"
	"github.com/iagoMAF/API_JOHARI/middleware"

	"github.com/joho/godotenv"

	_ "github.com/iagoMAF/API_JOHARI/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.Use(middleware.AuthMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// rotas de atleta
	r.GET("/atletas", controller.ExibeTodosAtletas)
	r.GET("/atleta/:cpf", controller.ExibeAtletaPorID)
	r.PATCH("/atleta/:cpf", controller.AtualizaAtleta)
	r.POST("/atleta", controller.CriaAtleta)
	r.GET("/atletas/lider/:cpf_lider", controller.ExibeAtletasPorLider)
	r.GET("/atletas/equipe/:id_equipe", controller.ExibeAtletasPorEquipe)

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

	// rotas de perguntas
	r.GET("/pergunta", controller.ExibeTodasAsPerguntas)
	r.GET("/perguntaComRespostas/:id", controller.ExibePerguntaComRespostas)
	r.GET("/perguntasComRespostas", controller.ExibePerguntasComRespostas)
	r.GET("/pergunta/:id", controller.ExibePerguntaPorID)

	// rotas de respostas
	r.GET("/resposta", controller.ExibeTodasAsRespostas)
	r.GET("/resposta/:id", controller.ExibeRespostaPorID)

	// rotas de resultado
	r.POST("/resultado", controller.CriaResultado)
	r.GET("/resultado/:cpf", controller.ExibeResultadosPorCPF)

	r.Run(":" + port)
}
