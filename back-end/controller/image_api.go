package controller

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/gin-gonic/gin"
)

type Image_Upload struct {
	File *multipart.FileHeader `form:"file" json:"file" binding:"required" swaggerignore:"true"`
}

var temp_path = "/tmp/images"

// SearchImg_Handler godoc
// @Summary get the image by image id
// @Schemes
// @Description get the image by image id
// @Tags images
// @Accept json
// @Produce json
// @Param		img_id		path	string	false	"Image ID"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Failure 500 "file system error"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /images/{img_id}/search [get]
func SearchOneImg_Handler(c *gin.Context) {
	ImgID := c.Param("img_id")
	flag, imgType, imgContent, err := dao.SearchOnePictureBasedOnID(ImgID)
	if !flag {
		response.ErrImageNotExistsResponse(c)
	} else {
		if err != nil {
			response.ErrMySQLReadResponse(c, "picture read failed")
		} else {
			// save the image file in temp
			utils.CreateFolder(temp_path)
			file_name := fmt.Sprintf("%s.%s", ImgID, imgType)
			pathAndName := fmt.Sprintf("%s/%s", temp_path, file_name)

			err := ioutil.WriteFile(pathAndName, imgContent, 0666)
			if err != nil {
				response.ErrFileStoreResponse(c)
				fmt.Println(err)
			} else {
				c.File(pathAndName)
				// c.File(pathAndName)
				// response.SuccessResponse(c, "picture read success")
				os.Remove(pathAndName)
			}

		}
	}

}

// SearchImgIDs_Handler godoc
// @Summary get all image ids of the course input
// @Schemes
// @Description get all image ids of the course input, and get a string array
// @Tags images
// @Accept json
// @Produce json
// @Param		course_name 	path	string	true	"Course Name"
// @Success 200 "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 "not exists"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Security ApiKeyAuth
// @Router /images/getIDs/{course_name}  [get]
func SearchImgIDs_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)

	_, baseCourse := course.GetCourseBaseCourse(c)

	flag, ids, err := dao.SearchCoursePictureIDs(baseCourse)
	if !flag {
		response.ErrImageNotExistsResponse(c)
	} else {
		if err != nil {
			response.ErrMySQLReadAllResponse(c, "fail to read instances of this base course!")
		}
		response.SuccessResponse(c, ids)
	}
}

// func SuccessResponseWithByteArray(c *gin.Context, fitlType string, content []byte) {
// 	c.Data(200, fitlType, content)
// }

// UploadImage_Handler godoc
// @Summary upload the image and its type, its base course to database, and get the image id
// @Schemes
// @Description upload the image and its type, base course to database, and get the image id
// @Tags images
// @Accept mpfd
// @Produce json
// @Param		course_name 	path	string	true	"Course Name"
// @Param		img_type		path	string	true	"Image Type"
// @Param file formData file true "the image file"
// @Success 200 "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Security ApiKeyAuth
// @Router /images/{course_name}/{img_type}/upload [post]
func UploadImage_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)
	_, baseCourse := course.GetCourseBaseCourse(c)

	imgType := c.Param("img_type")
	var imgBody Image_Upload
	validate.ValidateForm(c, &imgBody)
	id, err := dao.InsertOnePicture(baseCourse, imgType, course.FileToByte(c, imgBody.File))
	if err != nil {
		response.ErrMySQLCreateResponse(c, "failed to upload picture")
	}
	response.SuccessResponse(c, id)

}

// UpdateImage_Handler godoc
// @Summary update the image content and its type in the database
// @Schemes
// @Description update the image content and its type in the database
// @Tags images
// @Accept mpfd
// @Produce json
// @Param		course_name 	path	string	true	"Course Name"
// @Param		img_id		path	string	true	"Image ID"
// @Param		img_type		path	string	true	"New Image Type"
// @Param file formData file true "the image file"
// @Success 200 "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 "not exists"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Security ApiKeyAuth
// @Router /images/{course_name}/{img_id}/{img_type}/update [put]
func UpdateImage_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)
	_, baseCourse := course.GetCourseBaseCourse(c)
	imgID := c.Param("img_id")
	imgType := c.Param("img_type")
	var imgBody Image_Upload
	validate.ValidateForm(c, &imgBody)
	flag, err := dao.UpadateOnePicture(baseCourse, imgID, imgType, course.FileToByte(c, imgBody.File))
	if !flag {
		response.ErrImageNotExistsResponse(c)
	} else {
		if err != nil {
			response.ErrMySQLUpdateResponse(c, "update image failed")
		}
		response.SuccessResponse(c, "update image succeed")
	}
}

// DeleteImg_Handler godoc
// @Summary delete an image according to input id and its base course
// @Schemes
// @Description delete an image according to input id and its base course
// @Tags images
// @Accept json
// @Produce json
// @Param		course_name 	path	string	true	"Course Name"
// @Param		img_id		path	string	true	"Image ID"
// @Success 204
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadError} "mysql error"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /images/{course_name}/{img_id}/delete [delete]
func DeleteImg_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)
	_, baseCourse := course.GetCourseBaseCourse(c)

	ImgID := c.Param("img_id")
	flag, err := dao.DeletePicture(baseCourse, ImgID)
	if !flag {
		response.ErrImageNotExistsResponse(c)
	} else {
		if err != nil {
			response.ErrMySQLDeleteResponse(c, "delete picture failed")
		} else {
			response.NonContentResponse(c)
		}
	}

}
