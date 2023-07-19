package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializando o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Aqui está rodando a API
	router.Run(":8080")
}
