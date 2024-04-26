package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Sets up the environment
func ConfigureEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading file .env: %v", err)
	}

	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
