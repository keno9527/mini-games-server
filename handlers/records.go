package handlers

import (
	"net/http"
	"time"

	"server-go/models"
	"server-go/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRecords(c *gin.Context) {
	userID := c.Param("id")
	if _, ok := storage.GetUser(userID); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	records := storage.GetRecordsByUser(userID)
	if records == nil {
		records = []models.GameRecord{}
	}
	c.JSON(http.StatusOK, records)
}

func CreateRecord(c *gin.Context) {
	userID := c.Param("id")
	if _, ok := storage.GetUser(userID); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var req struct {
		GameID   string `json:"gameId" binding:"required"`
		Score    int    `json:"score"`
		Duration int    `json:"duration"`
		Result   string `json:"result" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record := models.GameRecord{
		ID:       uuid.New().String(),
		UserID:   userID,
		GameID:   req.GameID,
		Score:    req.Score,
		Duration: req.Duration,
		PlayedAt: time.Now(),
		Result:   req.Result,
	}

	if err := storage.CreateRecord(record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, record)
}

func GetPlayRanking(c *gin.Context) {
	ranking := storage.GetPlayRanking()
	if ranking == nil {
		ranking = []models.PlayRankItem{}
	}
	c.JSON(http.StatusOK, ranking)
}
