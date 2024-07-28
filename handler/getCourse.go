package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/models"
)

func GetCourse(ctx *gin.Context) {
	course_id := ctx.Param("course_id")
	if err := idCourseParameterValidator(course_id); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	course := models.Course{}

	if err := db.First(&course, course_id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "course not found")
		return
	}

	sendSuccess(ctx, http.StatusOK, "course found", course)
}
