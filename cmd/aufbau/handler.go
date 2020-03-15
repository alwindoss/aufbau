package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetchConfigHandler(ctx *gin.Context) {
	log.Printf("serving from fetchConfigHandler")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "fetchConfigHandler",
	})
}
func fetchAllConfigsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "fetchAllConfigsHandler",
	})

}
func createConfigHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "createConfigHandler",
	})

}
func updateConfigHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "updateConfigHandler",
	})

}
func deleteConfigHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleteConfigHandler",
	})

}
