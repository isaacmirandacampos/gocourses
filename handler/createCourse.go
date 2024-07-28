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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Debugf("Creating a new course with the following data: %v", request)
	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating a new course: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating a new course"})
		return
	}
}
