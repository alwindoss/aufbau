package main

import (
	"log"
	"net/http"

	"github.com/alwindoss/aufbau"
	"github.com/gin-gonic/gin"
)

func fetchConfigHandler(repository aufbau.Repository) gin.HandlerFunc {
	log.Printf("serving from fetchConfigHandler")
	return func(c *gin.Context) {
		repository.Fetch("", "", "")
		c.JSON(http.StatusOK, gin.H{
			"message": "fetchConfigHandler",
		})
	}
}
func fetchAllConfigsHandler(repository aufbau.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		repository.FetchAll("", "")
		c.JSON(http.StatusOK, gin.H{
			"message": "fetchAllConfigsHandler",
		})
	}
}
func createConfigHandler(repository aufbau.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		repository.Create(nil)
		c.JSON(http.StatusOK, gin.H{
			"message": "createConfigHandler",
		})
	}
}
func updateConfigHandler(repository aufbau.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		repository.Update(nil)
		c.JSON(http.StatusOK, gin.H{
			"message": "updateConfigHandler",
		})
	}
}
func deleteConfigHandler(repository aufbau.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		repository.Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"message": "deleteConfigHandler",
		})
	}
}
