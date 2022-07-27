package course

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func GetCourseBaseCourse(c *gin.Context) (string, string) {
	course := GetCourse(c)
	base, err := CourseToBase(course)
	if err != nil {
		fmt.Println(err)
		response.ErrMySQLReadResponse(c, "base_course")
	} else if base == "" {
		response.ErrCourseNoBaseCourseResponse(c, course)
	}
	return course, base
}

func CourseToBase(course_name string) (string, error) {
	var instance models.Base_Course_Relationship
	result := global.DB.Where(&models.Base_Course_Relationship{Course_name: course_name}).Find(&instance)
	return instance.Base_course, result.Error
}
