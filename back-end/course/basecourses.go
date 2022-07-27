package course

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

type Base_Course_Relationships struct {
	Course_Name string `json:"name" gorm:"type:varchar(255)"`
	Base_Course string `json:"base_course" gorm:"type:varchar(255)"`
}

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
	var instance Base_Course_Relationships
	result := global.DB.Where(&Base_Course_Relationships{Course_Name: course_name}).Find(&instance)
	return instance.Base_Course, result.Error
}
