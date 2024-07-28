package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/models"
)

func CreateCourse(ctx *gin.Context) {
	request := CourseRequest{}
	ctx.BindJSON(&request)
	err := request.createCourseBodyValidator()

	if err != nil {
		logger.Errorf("Error validating create course request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.Debugf("Creating a new course with the following data: %v", request)
	course := models.Course{
		Name:            request.Name,
		Slug:            request.Slug,
		Level:           request.Level,
		Description:     request.Description,
		DurationInHours: request.DurationInHours,
	}

	if err := db.Create(&course).Error; err != nil {
		logger.Errorf("Error creating a new course: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating a new course")
		return
	}
	sendSuccess(ctx, http.StatusCreated, "create-course", course)
}
