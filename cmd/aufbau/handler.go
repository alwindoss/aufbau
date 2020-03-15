package main

import (
	"encoding/json"
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
		orgID := c.Param("org_id")
		entityID := c.Param("entity_id")
		// c.Request.Body
		cfg := &aufbau.Configuration{}
		err := json.NewDecoder(c.Request.Body).Decode(cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		cfg.OrgID = orgID
		cfg.EntityID = entityID
		respCfg, err := repository.Create(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"config": respCfg,
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
