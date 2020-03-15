package main

import (
	"os"

	"github.com/gin-contrib/cors"
)

func main() {
	port := os.Getenv("AUFBAU_PORT")
	if port == "" {
		port = "8080"
	}
	engine := createEngine()
	createRoutes(engine)
	engine.Use(cors.Default())
	engine.Run(":" + port)
}
