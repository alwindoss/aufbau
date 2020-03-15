package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/alwindoss/aufbau/pkg/storage"
	"github.com/gin-contrib/cors"
)

func main() {
	port := os.Getenv("AUFBAU_PORT")
	if port == "" {
		port = "8080"
	}
	engine := createEngine()
	// connStr := "user=postgres dbname=aufbau sslmode=verify-false"
	connStr := "postgres://postgres:OpenSesame@localhost:5433/aufbau?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	repo := storage.NewRDBMS(db)
	createRoutes(repo, engine)
	engine.Use(cors.Default())
	engine.Run(":" + port)
}
