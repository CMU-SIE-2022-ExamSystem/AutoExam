package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/validate"
	"github.com/gin-gonic/gin"
)

// Exam_Handler godoc
// @Summary get exam question
// @Schemes
// @Description get exam question
// @Tags exam
// @Accept json
// @Produce json
// @Success 200 {object} response.Response "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/exam [get]
func Exam_Handler(c *gin.Context) {
	response.SuccessResponse(c, dao.GetQuestions())
}

// Usercourses_Handler godoc
// @Summary get user course information
// @Schemes
// @Description get user courses information from autolab
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Course_Info_Front} "desc"
// @Router /courses/{course_name}/info [get]
func Usercoursesinfo_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")

	body := autolab.AutolabGetHandler(c, token, "/courses")
	// fmt.Println(string(body))

	if strings.Contains(string(body), course_name) {
		autolab_resp := utils.User_courses_trans(string(body))

		autolab_map := utils.Map_course_info(autolab_resp)
		course_info := autolab_map[course_name]

		resp_body := models.Course_Info_Front{Name: course_info.Name, Display_name: course_info.Display_name, Auth_level: course_info.Auth_level}
		response.SuccessResponse(c, resp_body)
	} else {
		response.ErrorInternalResponse(c, response.Error{Type: "Auth-level", Message: "User is not registered for this course."})
	}
}

// Assessments_Handler godoc
// @Summary get course assessments
// @Schemes
// @Description get course assessments list
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.Autolab_Assessments} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments [get]
func Assessments_Handler(c *gin.Context) {
	assessments := course.GetFilteredAssessments(c)
	response.SuccessResponse(c, assessments)
}

// Submissions_Handler godoc
// @Summary get assessment submissions
// @Schemes
// @Description get assessment submissions list
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Submissions} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/submissions [get]
func Submissions_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submissions")
	// fmt.Println(string(body))

	autolab_resp := utils.Assessments_submissions_trans(string(body))

	response.SuccessResponse(c, autolab_resp)
}

// AssessmentConfigCategories_Handler godoc
// @Summary get the assessment's categories's possible list
// @Schemes
// @Description get the assessment's categories's possible list
// @Tags exam
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=dao.Categories_Return} "desc"
// @Security ApiKeyAuth
// @Router /courses/assessments/config/categories [get]
func AssessmentCategories_Handler(c *gin.Context) {
	data := dao.Categories_Return{Categories: dao.Assessment_Catergories}
	response.SuccessResponse(c, data)
}

// CreateAssessment_Handler godoc
// @Summary create an new exam configuration
// @Schemes
// @Description create an new exam configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param data body dao.AutoExam_Assessments_Create true "body data"
// @Success 201 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/ [post]
func CreateAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	var body dao.AutoExam_Assessments_Create
	validate.Validate(c, &body)
	course_name := course.GetCourse(c)
	course.Validate_assessment_name(c, course_name, body.Name)
	assessment := body.ToAutoExamAssessments(course_name)

	// check mongo error
	_, err := dao.CreateExam(assessment)
	if err != nil {
		response.ErrMongoDBCreateResponse(c, "assessment")
	}
	response.SuccessResponse(c, assessment)
}

// ReadAssessment_Handler godoc
// @Summary read an exam configuration
// @Schemes
// @Description read an exam configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [get]
func ReadAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	assessment := read_assessment(c, course_name, assessment_name)
	response.SuccessResponse(c, assessment)
}

// UpdateAssessment_Handler godoc
// @Summary update an exam configuration
// @Schemes
// @Description update an exam configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Param data body dao.AutoExam_Assessments_Update true "body data"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [put]
func UpdateAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)

	var body dao.AutoExam_Assessments_Update
	validate.Validate(c, &body)
	assessment := body.ToAutoExamAssessments(course_name)

	// check whether new data's name is same as the original assessment's name
	if !(assessment.Id == assessment_name) {
		if course.Validate_autoexam_assessment(c, course_name, assessment.Id) {
			response.ErrAssessmentNotValidResponse(c, course_name, assessment.Id)
		}
	}

	// check mongo error
	err := dao.UpdateExam(course_name, assessment_name, assessment)
	if err != nil {
		fmt.Println("==========")
		fmt.Println(err)
		fmt.Println("==========")
		response.ErrMongoDBUpdateResponse(c, "assessment")
	}

	response.SuccessResponse(c, assessment)
}

// DeleteAssessment_Handler godoc
// @Summary delete an exam configuration
// @Schemes
// @Description delete an exam configuration
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 204
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [delete]
func DeleteAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	err := dao.DeleteExam(course_name, assessment_name)
	if err != nil {
		response.ErrMongoDBDeleteResponse(c, "assessment")
	}
	response.NonContentResponse(c)
}

// DraftAssessment_Handler godoc
// @Summary edit an assessment's draft status
// @Schemes
// @Description edit an assessment's draft status
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Param data body dao.Draft true "body data"
// @Success 200 {object} response.Response{data=dao.Draft} "desc"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/draft [put]
func DraftAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)

	var body dao.Draft
	validate.Validate(c, &body)

	if !body.Draft {
		// assessment not in autolab
		if !course.Validate_autolab_assessment(c, course_name, assessment_name) {
			response.ErrAssessmentNotInAutolabResponse(c, course_name, assessment_name)
		}
	}

	assessment := read_assessment(c, course_name, assessment_name)
	assessment.Draft = body.Draft

	err := dao.UpdateExam(course_name, assessment_name, assessment)
	if err != nil {
		fmt.Println("==========")
		fmt.Println(err)
		fmt.Println("==========")
		response.ErrMongoDBUpdateResponse(c, "assessment")
	}
	response.SuccessResponse(c, assessment)
}

// DownloadAssessments_Handler godoc
// @Summary download an assessment tarball
// @Schemes
// @Description download an assessment tarball, only can be used for instructor or TA
// @Tags exam
// @Accept json
// @Produce mpfd
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/download [get]
// @Success 404 {object} response.Response{} "Not found course"
func DownloadAssessments_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	// course_name, assessment_name := course.GetCourseAssessment(c)
	// _ := read_assessment(c, course_name, assessment_name)
	// tar := course.Build_Assessment(c, course_name, assessment)
	// c.FileAttachment(tar, tar[strings.LastIndex(tar, "/")+1:])
}

// Usersubmit_Handler godoc
// @Summary submit answer
// @Schemes
// @Description submit answer to Tango
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Submit} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/submit [get]
func Usersubmit_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	body := autolab.AutolabSubmitHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submit", "./tmp/answer.tar")
	// fmt.Println(string(body))

	autolab_resp := utils.User_submit_trans(string(body))
	if autolab_resp.Version == 0 {
		response.ErrorResponseWithStatus(c, response.Error{Type: "Autolab", Message: "You are only allowed to submit once!"}, http.StatusForbidden)
	} else {
		response.SuccessResponse(c, autolab_resp)
	}
}

func read_assessment(c *gin.Context, course_name, assessment_name string) dao.AutoExam_Assessments {
	// read certain assessment
	assessment, err := dao.ReadExam(course_name, assessment_name)

	// check mongo error
	if err != nil {
		response.ErrMongoDBReadResponse(c, "assessment")
	}
	return assessment
}
