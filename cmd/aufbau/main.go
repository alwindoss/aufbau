package main

import (
	"os"
)

func main() {
	port := os.Getenv("AUFBAU_PORT")
	if port == "" {
		port = "8080"
	}
	engine := createEngine()
	createRoutes(engine)
	engine.Run(":" + port)
}
