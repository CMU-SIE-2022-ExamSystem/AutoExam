package course

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Student_Model = "student"
)

func GetAssessmentStudent(c *gin.Context) dao.Assessment_Student {
	course_name, assessment_name := GetCourseAssessment(c)
	GetCourseBaseCourse(c)

	email := jwt.GetEmail(c)
	student, err := dao.ReadStudent(course_name, assessment_name, email.Email)

	// check whether student in the list
	if err != nil {
		if err == mongo.ErrNoDocuments {
			response.ErrStudentNotValidResponse(c, assessment_name, email.Email)
		}
		response.ErrMongoDBReadResponse(c, Student_Model)
	}

	// check question is ready
	if student.Generated == -1 || student.Questions == nil || len(student.Questions) == 0 {
		response.ErrStudentNotValidResponse(c, assessment_name, email.Email)
	}
	return student
}

func GenerateAssessmentStudent(c *gin.Context) {
	course_name, assessment_name := GetCourseAssessment(c)
	GetCourseBaseCourse(c)

	assessment := GetAssessment(c, course_name, assessment_name)
	users, _ := CourseUserData(c)
	var err error
	for _, user := range users {
		student := assessment.GenerateAssessmentStudent(user.Email, course_name, assessment_name)
		_, err = dao.CreateOrUpdateStudent(student)
		if err != nil {
			assessment.Generated = -1
			assessment.GeneratedError = "There is an error happen when generating " + student.Email + "'s exam with error message: " + err.Error()
		}
	}
	if err == nil {
		assessment.Generated = 1
		assessment.GeneratedError = ""
	}

	dao.UpdateExam(course_name, assessment_name, assessment)
}
