package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Tag_Model = "tag"
)

// ReadTag_Handler godoc
// @Summary read all tags configuration
// @Schemes
// @Description read all tags configuration
// @Tags tag
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 {object} response.Response{data=[]dao.Tags} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags [get]
func ReadAllTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name := c.Param("course_name")
	tags, err := dao.ReadAllTags(course_name)
	if err != nil {
		response.ErrMongoDBReadAllResponse(c, Tag_Model)
	}
	response.SuccessResponse(c, tags)
}

// CreateTag_Handler godoc
// @Summary create a new tag configuration
// @Schemes
// @Description create a new tag configuration
// @Tags tag
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body dao.Tags_API true "body data"
// @Success 201 {object} response.Response{data=dao.Tags} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags [post]
func CreateTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course := course.GetCourse(c)

	var body dao.AutoExam_Tags_Create
	body.Course = course
	validate.ValidateJson(c, &body)

	// autoexam_tags := body.ToAutoExamTagsCreate(course)
	r, err := dao.CreateTag(body)
	if err != nil {
		response.ErrMongoDBCreateResponse(c, Tag_Model)
	}

	if id, ok := r.InsertedID.(primitive.ObjectID); ok {
		tags, err := dao.ReadTagById(id.Hex())
		if err != nil {
			response.ErrMongoDBCreateResponse(c, Tag_Model)
		}
		response.CreatedResponse(c, tags)
	} else {
		response.ErrMongoDBCreateResponse(c, Tag_Model)
	}
}

// ReadTag_Handler godoc
// @Summary read a tag configuration
// @Schemes
// @Description read a tag configuration
// @Tags tag
// @Accept json
// @Produce json
// @Param		course_name	    path	string	true	"Course Name"
// @Param		tag_id			path	string	true	"Tag Id"
// @Success 200 {object} response.Response{data=dao.Tags} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [get]
func ReadTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, tag_id := course.GetCourseTagID(c)

	tags, err := dao.ReadTagById(tag_id)
	if err != nil {
		response.ErrMongoDBReadResponse(c, Tag_Model)
	}
	response.SuccessResponse(c, tags)
}

// UpdateTag_Handler godoc
// @Summary update a tag configuration
// @Schemes
// @Description update a tag configuration
// @Tags tag
// @Accept json
// @Produce json
// @Param		course_name	    path	string	true	"Course Name"
// @Param		tag_id			path	string	true	"Tag Id"
// @Param data body dao.Tags_API true "body data"
// @Success 200 {object} response.Response{data=dao.Tags} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [put]
func UpdateTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course, tag_id := course.GetCourseTagID(c)

	// do nothing to deal with the same name input because cannot use validate.Validate twice
	var update dao.AutoExam_Tags_Create
	update.Course = course
	validate.ValidateJson(c, &update)

	// autoexam_tags := body.ToAutoExamTagsCreate(course)
	err := dao.UpdateTag(tag_id, update)
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Tag_Model)
	}

	output := dao.Tags{
		Id:     tag_id,
		Name:   update.Name,
		Course: course,
	}
	response.SuccessResponse(c, output)
}

// DeleteTag_Handler godoc
// @Summary delete a tag configuration
// @Schemes
// @Description delete a tag configuration
// @Tags tag
// @Accept json
// @Produce json
// @Param		course_name	    path	string	true	"Course Name"
// @Param		tag_id			path	string	true	"Tag Id"
// @Success 204
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [delete]
func DeleteTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, tag_id := course.GetCourseTagID(c)
	if status, err := dao.ValidateTagUsedById(tag_id); err != nil {
		response.ErrMongoDBUpdateResponse(c, Tag_Model)
	} else if !status {
		response.ErrTagNotSafeResponse(c, dao.ReadTagName(tag_id))
	}

	err := dao.DeleteTagById(tag_id)
	if err != nil {
		response.ErrMongoDBDeleteResponse(c, Tag_Model)
	}
	response.NonContentResponse(c)
}
