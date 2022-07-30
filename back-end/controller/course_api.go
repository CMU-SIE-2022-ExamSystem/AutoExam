package controller

import (
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
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Student_Model = "student"
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

	// course_name := c.Param("course_name")
	course_name := course.GetCourse(c)

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
// @Success 200 {object} response.Response{data=[]models.Autolab_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadAllError} "mongo error"
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
// @Success 200 {object} response.Response{data=models.Submissions} "success"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/submissions [get]
func Submissions_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	// course_name := c.Param("course_name")
	// assessment_name := c.Param("assessment_name")
	course_name, assessment_name := course.GetCourseAssessment(c)

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
// @Success 200 {object} response.Response{data=dao.Categories_Return} "success"
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
// @Success 201 {object} response.Response{data=dao.AutoExam_Assessments} "created"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.CourseNotValidError} "not valid of course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBCreateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/ [post]
func CreateAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name, _ := course.GetCourseBaseCourse(c)

	var body dao.AutoExam_Assessments_Create
	validate.ValidateJson(c, &body)

	course.Validate_assessment_name(c, course_name, body.Name)
	assessment := body.ToAutoExamAssessments(course_name)

	// check mongo error
	_, err := dao.CreateExam(assessment)
	if err != nil {
		response.ErrMongoDBCreateResponse(c, "assessment")
	}
	response.CreatedResponse(c, assessment)
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
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [get]
func ReadAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

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
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.AssessmentModifyNotSafeError} "not update safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [put]
func UpdateAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	var body dao.AutoExam_Assessments_Update_Validate
	_, body.BaseCourse = course.GetCourseBaseCourse(c)
	validate.ValidateJson(c, &body)
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
// @Success 204 "no content"
// @Failure 400 {object} response.BadRequestResponse{error=response.AssessmentModifyNotSafeError} "not delete safe or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBDeleteError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [delete]
func DeleteAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
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
// @Success 200 {object} response.Response{data=dao.Draft} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.AssessmentNotInAutolabError} "assessment not in autolab or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBUpdateError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/draft [put]
func DraftAssessment_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	var body dao.Draft
	validate.ValidateJson(c, &body)

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
// @Success 200 {object} response.Response{data=dao.Draft} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.AssessmentNoSettingsbError} "no settings or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
func DownloadAssessments_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	_, base_course := course.GetCourseBaseCourse(c)
	assessment := read_assessment(c, course_name, assessment_name)
	if len(assessment.Settings) == 0 {
		response.ErrAssessmentNoSettingsResponse(c, assessment_name)
	}
	tar := course.Build_Assessment(c, course_name, base_course, assessment)
	c.FileAttachment(tar, tar[strings.LastIndex(tar, "/")+1:])
}

// QuestionAssessments_Handler godoc
// @Summary get all the questions based on the user
// @Schemes
// @Description get all the questions based on the user
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200  "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/question [get]
// @Security ApiKeyAuth
func QuestionAssessments_Handler(c *gin.Context) {
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
	email := jwt.GetEmail(c)
	student, err := dao.ReadStudent(course_name, assessment_name, email.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// TODO response for no student instance
		}
		response.ErrMongoDBReadResponse(c, Student_Model)
	}
	response.SuccessResponse(c, student.ToRealQuestions())
}

// GetAnswersAssessments_Handler godoc
// @Summary get answers based on the user
// @Schemes
// @Description get answers based on the user
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=dao.Answers_Upload} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers [get]
// @Security ApiKeyAuth
func GetAnswersAssessments_Handler(c *gin.Context) {
	student := read_assessment_student(c)
	response.SuccessResponse(c, student.Answers)
}

// GetAnswersAssessments_Handler godoc
// @Summary upload user's answers
// @Schemes
// @Description upload user's answers
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Param data body dao.Answers_Upload true "body data"
// @Success 200 {object} response.Response{data=dao.Answers_Upload} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers [put]
// @Security ApiKeyAuth
func UploadAnswersAssessments_Handler(c *gin.Context) {
	student := read_assessment_student(c)

	var body dao.Answers_Upload_Validate
	body.Student = student
	validate.ValidateJson(c, &body)

	student.Answers = body.Answers
	instance, err := dao.CreateOrUpdateStudent(student)
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Student_Model)
	}

	response.SuccessResponse(c, dao.Answers_Upload{Answers: instance.Answers})
}

// GetAnswersAssessments_Handler godoc
// @Summary get answers structure based on the user
// @Schemes
// @Description get answers structure based on the user
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=[]dao.Student_Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers/struct [get]
// @Security ApiKeyAuth
func GetAnswersStructAssessments_Handler(c *gin.Context) {
	student := read_assessment_student(c)
	student.ToAnwerStruct()
	response.SuccessResponse(c, student.Solutions)
}

// GenerateAssessments_Handler godoc
// @Summary generate the assessment for all student in the course
// @Schemes
// @Description generate the assessment for all student in the course
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200  "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/generate [get]
// @Security ApiKeyAuth
func GenerateAssessments_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	generate_assessment_student(c)
	// TODO que system for back process
	response.SuccessResponse(c, nil)
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

	// course_name := c.Param("course_name")
	// assessment_name := c.Param("assessment_name")
	course_name, assessment_name := course.GetCourseAssessment(c)

	body := autolab.AutolabSubmitHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submit", "./tmp/answer.tar")
	// fmt.Println(string(body))

	autolab_resp := utils.User_submit_trans(string(body))
	response.SuccessResponse(c, autolab_resp)
	// if autolab_resp.Version == 0 {
	// 	response.ErrorResponseWithStatus(c, response.Error{Type: "Autolab", Message: "You are only allowed to submit once!"}, http.StatusForbidden)
	// } else {
	// 	response.SuccessResponse(c, autolab_resp)
	// }
}

// CheckSubmission_Handler godoc
// @Summary check assessment submission
// @Schemes
// @Description check assessment submission
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 "This user has already submitted."
// @Failure 404 "This user has no submission records."
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/check [get]
func CheckSubmission_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name, assessment_name := course.GetCourseAssessment(c)

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submissions")
	// fmt.Println(string(body))

	_, flag := utils.Assessments_submissionscheck_trans(string(body))

	if flag {
		response.SuccessResponse(c, "This user has already submitted.")
	} else {
		response.ErrorResponseWithStatus(c, "This user has no submission records.", http.StatusNotFound)
	}
}

func read_assessment(c *gin.Context, course_name, assessment_name string) dao.AutoExam_Assessments {
	// read certain assessment
	assessment, err := dao.ReadExam(course_name, assessment_name)

	// check mongo error
	if err != nil {
		response.ErrMongoDBReadResponse(c, Student_Model)
	}
	return assessment
}

func generate_assessment_student(c *gin.Context) {
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	assessment := read_assessment(c, course_name, assessment_name)
	users, _ := course.CourseUserData(c)
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

func read_assessment_student(c *gin.Context) dao.Assessment_Student {
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	email := jwt.GetEmail(c)
	student, err := dao.ReadStudent(course_name, assessment_name, email.Email)

	// check mongo error
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// TODO response for no student instance
		}
		response.ErrMongoDBReadResponse(c, Student_Model)
	}
	return student
}
