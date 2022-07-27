package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func BaseCourseRouter(Router *gin.RouterGroup) {
	BaseCourseRouter := Router.Group("basecourse")
	{
		BaseCourseRouter.POST("/:base", jwt.JWTAuth(), controller.CreateBaseCourse_Handler)
		BaseCourseRouter.GET("/list", jwt.JWTAuth(), controller.ReadAllBaseCourse_Handler)
		BaseCourseRouter.PUT("/:base/:new", jwt.JWTAuth(), controller.UpdateBaseCourse_Handler)
		BaseCourseRouter.DELETE("/:base", jwt.JWTAuth(), controller.DeleteBaseCourse_Handler)
	}
}
