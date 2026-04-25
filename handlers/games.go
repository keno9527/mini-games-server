package handlers

import (
	"net/http"

	"server-go/storage"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetGames())
}

func GetGame(c *gin.Context) {
	id := c.Param("id")
	game, ok := storage.GetGame(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}
