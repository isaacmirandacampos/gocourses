package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, statusCode int, msg string) {
	ctx.JSON(statusCode, gin.H{"message": msg})
}
