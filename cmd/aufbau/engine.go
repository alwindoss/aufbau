package main

import "github.com/gin-gonic/gin"

func createEngine() *gin.Engine {
	return gin.Default()
}

func createRoutes(e *gin.Engine) {
	e.GET("/org/:org_id/entity/:entity_id/config/:config_id", fetchConfigHandler)
	e.GET("/org/:org_id/entity/:entity_id/config", fetchAllConfigsHandler)
	e.POST("/org/:org_id/entity/:entity_id/config", createConfigHandler)
	e.PUT("/org/:org_id/entity/:entity_id/config/:config_id", updateConfigHandler)
	e.DELETE("/org/:org_id/entity/:entity_id/config/:config_id", deleteConfigHandler)
}
