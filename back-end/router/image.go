package router

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/controller"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/gin-gonic/gin"
)

func ImageRouter(Router *gin.RouterGroup) {
	ImageRouter := Router.Group("images")
	{
		// image CRUD
		ImageRouter.GET("/:img_id/search", controller.SearchOneImg_Handler)
		ImageRouter.GET("/getIDs/:course_name", jwt.JWTAuth(), controller.SearchImgIDs_Handler)
		ImageRouter.POST("/:course_name/:img_type/upload", jwt.JWTAuth(), controller.UploadImage_Handler)
		ImageRouter.PUT("/:course_name/:img_id/:img_type/update", jwt.JWTAuth(), controller.UpdateImage_Handler)
		ImageRouter.DELETE("/:course_name/:img_id/delete", jwt.JWTAuth(), controller.DeleteImg_Handler)
	}

}
