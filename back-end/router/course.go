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
		CourseRouter.GET("/:course_name/info", jwt.JWTAuth(), controller.Usercoursesinfo_Handler)
		CourseRouter.GET("/:course_name/assessments", jwt.JWTAuth(), controller.Assessments_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/submissions", jwt.JWTAuth(), controller.Submissions_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/download", jwt.JWTAuth(), controller.DownloadAssessments_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/submit", jwt.JWTAuth(), controller.Usersubmit_Handler)
		CourseRouter.GET("/:course_name/course_user_data", jwt.JWTAuth(), controller.Course_all_Test)
	}
}
