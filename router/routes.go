package router

import (
	"github.com/gin-gonic/gin"
	"github.com/isaacmirandacampos/gocourses/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	api := router.Group("/api")
	{
		api.GET("/health-check", handler.HealthCheck)
	}

	v1 := api.Group("/v1")
	{
		v1.GET("/courses", handler.ListCourse)

		v1.POST("/courses", handler.CreateCourse)

		v1.DELETE("/courses/:course_id", handler.DeleteCourse)

		v1.PATCH("/courses/:course_id", handler.PatchCourse)

		v1.GET("/courses/:course_id", handler.GetCourse)
	}
}
