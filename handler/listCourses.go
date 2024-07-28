package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/models"
)

func ListCourse(ctx *gin.Context) {
	courses := []models.Course{}
	if err := db.Find(&courses).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to get courses")
		return
	}
	sendSuccess(ctx, http.StatusOK, "courses found", courses)
}
