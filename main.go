package main

import (
	"fmt"
	"jagch/tokenGenerator/application"
	"jagch/tokenGenerator/config"
	"jagch/tokenGenerator/delivery"
	"jagch/tokenGenerator/infraestructure"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Config environment
	config.ConfigureEnvironment()

	// Gin-Gonic configuration
	router := gin.Default()

	// App configuration
	redisClient := infraestructure.NewRedisClient()
	tokenRepo := infraestructure.NewRedisTokenRepository(redisClient)
	tokenService := application.NewTokenService(tokenRepo)

	// Controllers HTTP configuration
	tokenController := delivery.NewTokennController(tokenService)

	// Defining routes
	api := router.Group("/api/v1/token")
	{
		api.POST("/:quantity/whitelabel/:whitelabelName", tokenController.GenerateTokens)
		api.GET("/:token/whitelabel/:whitelabelName", tokenController.CheckToken)
	}

	fmt.Println("port ", os.Getenv("SERVER_PORT"))

	// Run server
	err := router.Run(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
