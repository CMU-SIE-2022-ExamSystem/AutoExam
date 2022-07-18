package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func GraderRouter(Router *gin.RouterGroup) {
	GraderRouter := Router.Group("graders")
	{
		// assessment CRUD
		GraderRouter.POST("/", jwt.JWTAuth(), controller.CreateGrader_Handler)
		GraderRouter.GET("/:grader_name", jwt.JWTAuth(), controller.ReadGrader_Handler)
		GraderRouter.PUT("/:grader_name", jwt.JWTAuth(), controller.UpdateGrader_Handler)
		GraderRouter.DELETE("/:grader_name", jwt.JWTAuth(), controller.DeleteGrader_Handler)
	}
}
