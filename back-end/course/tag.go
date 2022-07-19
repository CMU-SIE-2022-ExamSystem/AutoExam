package course

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func GetCourseTagID(c *gin.Context) (string, string) {
	course := GetCourse(c)
	tag_id := c.Param("tag_id")
	Validate_tag(c, course, tag_id)

	return course, tag_id
}

func Validate_tag(c *gin.Context, course, tag_id string) {
	if status, err := dao.ValidateTagById(course, tag_id); err != nil {
		response.ErrMongoDBReadResponse(c, "tag")
	} else if status {
		response.ErrTagNotValidResponse(c, course, tag_id)
	}
}
