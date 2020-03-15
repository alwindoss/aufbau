package main

import (
	"github.com/alwindoss/aufbau"
	"github.com/gin-gonic/gin"
)

func createEngine() *gin.Engine {
	return gin.Default()
}

func createRoutes(repo aufbau.Repository, e *gin.Engine) {
	e.GET("/org/:org_id/entity/:entity_id/config/:config_id", fetchConfigHandler(repo))
	e.GET("/org/:org_id/entity/:entity_id/config", fetchAllConfigsHandler(repo))
	e.POST("/org/:org_id/entity/:entity_id/config", createConfigHandler(repo))
	e.PUT("/org/:org_id/entity/:entity_id/config/:config_id", updateConfigHandler(repo))
	e.DELETE("/org/:org_id/entity/:entity_id/config/:config_id", deleteConfigHandler(repo))
}
