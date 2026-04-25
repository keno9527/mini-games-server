package handlers

import (
	"net/http"
	"time"

	"server-go/models"
	"server-go/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, storage.GetUsers())
}

func CreateUser(c *gin.Context) {
	var req struct {
		Name   string `json:"name" binding:"required"`
		Avatar string `json:"avatar"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	avatar := req.Avatar
	if avatar == "" {
		avatar = "default"
	}

	user := models.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Avatar:    avatar,
		CreatedAt: time.Now(),
	}

	if err := storage.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if _, ok := storage.GetUser(id); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err := storage.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func GetUserStats(c *gin.Context) {
	id := c.Param("id")
	if _, ok := storage.GetUser(id); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, storage.GetUserStats(id))
}
