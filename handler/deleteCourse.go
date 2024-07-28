package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/models"
)

func DeleteCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")
	if err := idCourseParameterValidator(id); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	course := models.Course{}
	if err := db.First(&course, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "course not found")
		return
	}

	if err := db.Delete(&course).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to delete course")
		return
	}

	sendSuccess(ctx, http.StatusOK, "course deleted successfully", course)
}
