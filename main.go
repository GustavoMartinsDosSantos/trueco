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
	pairController := controllers.NewPairController(config.DB)

	r.GET("/player/", playerController.GetAllPlayers)
	r.GET("/player/:id", playerController.GetPlayerByID)
	r.POST("/player", playerController.CreatePlayer)
	r.POST("/pair", pairController.CreatePair)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
