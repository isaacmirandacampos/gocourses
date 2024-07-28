package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/models"
)

func PatchCourse(ctx *gin.Context) {
	course_id := ctx.Param("course_id")
	var request CourseRequest
	if err := idCourseParameterValidator(course_id); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.BindJSON(&request)
	if err := request.updateCourseBodyValidator(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	course := models.Course{}
	if err := db.First(&course, course_id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "course not found")
		return
	}
	if request.Name != "" {
		course.Name = request.Name
	}
	if request.Slug != "" {
		course.Slug = request.Slug
	}
	if request.Level != "" {
		course.Level = request.Level
	}
	if request.Description != "" {
		course.Description = request.Description
	}
	if request.DurationInHours > 0 {
		course.DurationInHours = request.DurationInHours
	}
	if err := db.Save(&course).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to update course")
		return
	}

	sendSuccess(ctx, http.StatusOK, "course updated", course)
}
