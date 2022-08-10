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
// @Success 200 {object} response.Response{data=[]dao.Tags} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadAllError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags [get]
func ReadAllTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	_, base_course := course.GetCourseBaseCourse(c)
	tags, err := dao.ReadAllTags(base_course)
	if err != nil {
		response.ErrMongoDBReadAllResponse(c, Tag_Model)
	}
	if len(tags) == 0 {
		response.SuccessResponse(c, []string{})
	} else {
		response.SuccessResponse(c, tags)
	}
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
// @Success 201 {object} response.Response{data=dao.Tags} "created"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBCreateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags [post]
func CreateTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, base_course := course.GetCourseBaseCourse(c)

	var body dao.AutoExam_Tags_Create
	body.Course = base_course
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
// @Success 200 {object} response.Response{data=dao.Tags} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.TagNotValidError} "not valid of tag or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [get]
func ReadTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, tag_id := course.GetBaseCourseTagID(c)

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
// @Success 200 {object} response.Response{data=dao.Tags} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.TagNotValidError} "not valid of tag or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [put]
func UpdateTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base_course, tag_id := course.GetBaseCourseTagID(c)

	// do nothing to deal with the same name input because cannot use validate.Validate twice
	var update dao.AutoExam_Tags_Create
	update.Course = base_course
	validate.ValidateJson(c, &update)

	// autoexam_tags := body.ToAutoExamTagsCreate(course)
	err := dao.UpdateTag(tag_id, update)
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Tag_Model)
	}

	output := dao.Tags{
		Id:     tag_id,
		Name:   update.Name,
		Course: base_course,
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
// @Success 204 "no content"
// @Failure 400 {object} response.BadRequestResponse{error=response.TagDeleteNotSafeError} "not delete safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.TagNotValidError} "not valid of tag or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBDeleteError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/tags/{tag_id} [delete]
func DeleteTag_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	_, tag_id := course.GetBaseCourseTagID(c)
	if status, err := dao.ValidateTagUsedById(tag_id); err != nil {
		response.ErrMongoDBUpdateResponse(c, Tag_Model)
	} else if !status {
		response.ErrTagDeleteNotSafeResponse(c, dao.ReadTagName(tag_id))
	}

	err := dao.DeleteTagById(tag_id)
	if err != nil {
		response.ErrMongoDBDeleteResponse(c, Tag_Model)
	}
	response.NonContentResponse(c)
}
