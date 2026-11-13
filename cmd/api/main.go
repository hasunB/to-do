package main

import (
	"log"
	"net/http"
	"todo-api/internals/infrastructure/config"
	"todo-api/internals/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {

  cfg, err := config.Load()
  if err != nil {
    log.Fatal("failed to load configuration", err)
  }

  pool, err := database.Connect(cfg.DatabaseURL)
  if err != nil {
    log.Fatal("failed to connect to database", err)
  }

  defer pool.Close() 

  // Create a Gin router with default middleware (logger and recovery)
  router := gin.Default()
  router.SetTrustedProxies(nil)

  // Define a simple GET endpoint
  router.GET("/", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong is pong",
      "status": "success",
      "database": "connected",
    })
  })

  router.Run(":" + cfg.Port)
  
}