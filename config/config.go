package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Sets up the environment
func ConfigureEnvironment() {
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
