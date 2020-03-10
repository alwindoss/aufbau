package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("AUFBAU_PORT")
	if port == "" {
		port = "8080"
	}
	engine := gin.Default()
	engine.Run(":" + port)
}
