package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(ctx *gin.Context) {
	request := CreateCourseRequest{}
	ctx.BindJSON(&request)
	err := request.CreateCourseValidator()

	if err != nil {
		logger.Errorf("Error validating create course request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.Debugf("Creating a new course with the following data: %v", request)
	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating a new course: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating a new course")
		return
	}
}
