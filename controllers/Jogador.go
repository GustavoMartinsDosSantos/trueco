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

func (pc *PlayerController) GetJogadorByID(c *gin.Context) {
	jogadorID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var jogador models.Jogador
	if err := pc.db.First(&jogador, jogadorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogador not found"})
		return
	}

	c.JSON(http.StatusOK, jogador)
}

func (pc *PlayerController) CreateJogador(c *gin.Context) {
	var jogador models.Jogador

	if err := c.ShouldBindJSON(&jogador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.db.Create(&jogador).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create jogador"})
		return
	}

	c.JSON(http.StatusCreated, jogador)
}
