package controllers

import (
	"net/http"

	"github.com/GustavoMartinsDosSantos/trueco/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PairController struct {
	db *gorm.DB
}

func NewPairController(db *gorm.DB) *PairController {
	return &PairController{db: db}
}

func (pc *PairController) CreatePair(c *gin.Context) {
	var pair models.Pair
	var player models.Player

	if err := c.ShouldBindJSON(&pair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player.ID = pair.Player1ID

	if err := player.IsPlayerExists(pc.db); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	player.ID = pair.Player2ID

	if err := player.IsPlayerExists(pc.db); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := pair.BeforeSave(pc.db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := pair.BeforeSave(pc.db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := pc.db.Create(&pair).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pair"})
		return
	}

	c.JSON(http.StatusCreated, pair)
}
