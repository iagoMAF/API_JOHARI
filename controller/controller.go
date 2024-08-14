package controller

import "github.com/gin-gonic/gin"

func ExibeTodosAtletas(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Teste",
	})
}
