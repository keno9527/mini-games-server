package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"server-go/handlers"
	"server-go/logfile"
	"server-go/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	closeLog, err := logfile.Init("log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init log file: %v\n", err)
		os.Exit(1)
	}
	defer closeLog()

	if err := storage.Init(); err != nil {
		log.Fatalf("failed to init storage: %v", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[API] %s %s %d %s",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(),
			time.Since(start).Round(time.Millisecond))
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: false,
	}))

	api := r.Group("/api")
	{
		api.GET("/games", handlers.GetGames)
		api.GET("/games/:id", handlers.GetGame)

		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
		api.GET("/users/:id/stats", handlers.GetUserStats)
		api.GET("/users/:id/records", handlers.GetRecords)
		api.POST("/users/:id/records", handlers.CreateRecord)
		api.GET("/records/ranking", handlers.GetPlayRanking)
	}

	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
