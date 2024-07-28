package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCourse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
