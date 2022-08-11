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
		ImageRouter.GET("/getIDs/:base_course", jwt.JWTAuth(), controller.SearchImgIDs_Handler)
		ImageRouter.POST("/:base_course/:img_type/upload", jwt.JWTAuth(), controller.UploadImage_Handler)
		ImageRouter.PUT("/:base_course/:img_id/:img_type/update", jwt.JWTAuth(), controller.UpdateImage_Handler)
		ImageRouter.DELETE("/:base_course/:img_id/delete", jwt.JWTAuth(), controller.DeleteImg_Handler)
	}

}
