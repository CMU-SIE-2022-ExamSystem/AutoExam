package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func CourseRouter(Router *gin.RouterGroup) {
	CourseRouter := Router.Group("courses")
	{
		CourseRouter.GET("/:course_name/assessments/:assessment_name/exam", controller.Exam_Handler)
		CourseRouter.GET("/:course_name/assessments", jwt.JWTAuth(), controller.Assessments_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/submissions", jwt.JWTAuth(), controller.Submissions_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/download", controller.DownloadAssessments_Handler)
	}
}
