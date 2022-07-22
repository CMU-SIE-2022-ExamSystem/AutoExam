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
		CourseRouter.GET("/assessments/config/categories", jwt.JWTAuth(), controller.AssessmentCategories_Handler)

		// assessment CRUD
		CourseRouter.POST("/:course_name/assessments", jwt.JWTAuth(), controller.CreateAssessment_Handler)
		CourseRouter.GET("/:course_name/assessments/:assessment_name/", jwt.JWTAuth(), controller.ReadAssessment_Handler)
		CourseRouter.PUT("/:course_name/assessments/:assessment_name/", jwt.JWTAuth(), controller.UpdateAssessment_Handler)
		CourseRouter.DELETE("/:course_name/assessments/:assessment_name/", jwt.JWTAuth(), controller.DeleteAssessment_Handler)

		// modify assessment's draft
		CourseRouter.PUT("/:course_name/assessments/:assessment_name/draft", jwt.JWTAuth(), controller.DraftAssessment_Handler)

		// tags CRUD
		CourseRouter.GET("/:course_name/tags", jwt.JWTAuth(), controller.ReadAllTag_Handler)
		CourseRouter.POST("/:course_name/tags", jwt.JWTAuth(), controller.CreateTag_Handler)
		CourseRouter.GET("/:course_name/tags/:tag_id/", jwt.JWTAuth(), controller.ReadTag_Handler)
		CourseRouter.PUT("/:course_name/tags/:tag_id/", jwt.JWTAuth(), controller.UpdateTag_Handler)
		CourseRouter.DELETE("/:course_name/tags/:tag_id/", jwt.JWTAuth(), controller.DeleteTag_Handler)

		// questions CRUD
		CourseRouter.GET("/:course_name/questions", jwt.JWTAuth(), controller.ReadAllQuestion_Handler)
		CourseRouter.POST("/:course_name/questions", jwt.JWTAuth(), controller.CreateQuestion_Handler)
		CourseRouter.GET("/:course_name/questions/:question_id/", jwt.JWTAuth(), controller.ReadQuestion_Handler)
		CourseRouter.PUT("/:course_name/questions/:question_id/", jwt.JWTAuth(), controller.UpdateQuestion_Handler)
		CourseRouter.DELETE("/:course_name/questions/:question_id/", jwt.JWTAuth(), controller.DeleteQuestion_Handler)

		// graders CRUD
		CourseRouter.GET("/:course_name/graders", jwt.JWTAuth(), controller.ReadAllGrader_Handler)
		CourseRouter.POST("/:course_name/graders", jwt.JWTAuth(), controller.CreateGrader_Handler)
		CourseRouter.GET("/:course_name/graders/:grader_name", jwt.JWTAuth(), controller.ReadGrader_Handler)
		CourseRouter.PUT("/:course_name/graders/:grader_name", jwt.JWTAuth(), controller.UpdateGrader_Handler)
		CourseRouter.DELETE("/:course_name/graders/:grader_name", jwt.JWTAuth(), controller.DeleteGrader_Handler)

		CourseRouter.GET("/:course_name/graders/list", jwt.JWTAuth(), controller.ReadGraderList_Handler)
		CourseRouter.PUT("/:course_name/graders/:grader_name/valid", jwt.JWTAuth(), controller.ValidGrader_Handler)
		CourseRouter.PUT("/:course_name/graders/:grader_name/force", jwt.JWTAuth(), controller.UpdateForceGrader_Handler)
	}
}
