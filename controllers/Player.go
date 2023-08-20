package controllers

import (
	"net/http"
	"strconv"

	"github.com/GustavoMartinsDosSantos/trueco/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlayerController struct {
	db *gorm.DB
}

func NewPlayerController(db *gorm.DB) *PlayerController {
	return &PlayerController{db: db}
}

func (pc *PlayerController) GetAllPlayers(c *gin.Context) {
	var players []models.Player

	if err := pc.db.Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, players)

}

func (pc *PlayerController) GetPlayerByID(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var player models.Player
	if err := pc.db.First(&player, playerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "player not found"})
		return
	}

	c.JSON(http.StatusOK, player)
}

func (pc *PlayerController) CreatePlayer(c *gin.Context) {
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.db.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}

	c.JSON(http.StatusCreated, player)
}
