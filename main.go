package main

import (
	"net/http"

	"github.com/GustavoMartinsDosSantos/trueco/config"
	"github.com/GustavoMartinsDosSantos/trueco/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	playerController := controllers.NewPlayerController(config.DB)

	r.GET("/jogador/:id", playerController.GetJogadorByID)
	r.POST("/jogador", playerController.CreateJogador)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
