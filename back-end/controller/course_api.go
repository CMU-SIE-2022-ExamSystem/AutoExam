package controller

import (
	"math"
	"net/http"
	"strconv"
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

const (
	Student_Model = "student"
)

// Usercourses_Handler godoc
// @Summary get user course information
// @Schemes
// @Description get user courses information from autolab
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Course_Info_Front} "success"
// @Failure 500 "not registered for this course"
// @Param		course_name			path	string	true	"Course Name"
// @Security ApiKeyAuth
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
		base_course, _ := dao.ReadBaseCourseRelation(course_name)
		resp_body := models.Course_Info_Front{Name: course_info.Name, Display_name: course_info.Display_name, Auth_level: course_info.Auth_level, Base_course: base_course}
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
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	course_name := course.GetCourse(c)

	assessments := course.GetFilteredAssessments(c)
	for i := range assessments {
		student, err := dao.ReadStudent(course_name, assessments[i].Name, user.Email)
		if err == nil {
			assessments[i].Submitted = student.Submitted
			assessments[i].Can_submit = student.Can_submit
		}
	}
	response.SuccessResponse(c, assessments)
}

// Submissions_Handler godoc
// @Summary get assessment submissions
// @Schemes
// @Description get assessment submissions list
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.Submissions_Front} "success"
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
	student, _ := dao.ReadStudent(course_name, assessment_name, user.Email)
	maxsocre := make(map[string]float64)
	for i := range student.Problems {
		maxsocre[student.Problems[i].Name] = student.Problems[i].MaxScore
	}

	var resp []models.Submissions_Front
	for _, submission := range autolab_resp {
		resp = append(resp, course.ToFront(submission, maxsocre))
	}

	response.SuccessResponse(c, resp)
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
// @Summary read an exam configuration. student would only get start_at, end_at desciprition while ta or instructor can retrieve all details
// @Schemes
// @Description read an exam configuration. student would only get start_at, end_at desciprition while ta or instructor can retrieve all details
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=dao.AutoExam_Assessments} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name} [get]
func ReadAssessment_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)

	assessment := course.GetAssessment(c, course_name, assessment_name)

	auth_level := jwt.Get_authlevel_DB(c)
	if auth_level == "student" {
		student, err := dao.ReadStudent(course_name, assessment_name, user.Email)
		if err != nil {
			response.ErrMongoDBReadResponse(c, "student")
		}
		response.SuccessResponse(c, assessment.ToAssessmentsStudent(student))
	} else {
		response.SuccessResponse(c, assessment)
	}

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

	assessment := course.GetAssessment(c, course_name, assessment_name)
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
// @Failure 400 {object} response.BadRequestResponse{error=response.AssessmentNoSettingsError} "no settings or no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
func DownloadAssessments_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	course_name, assessment_name := course.GetCourseAssessment(c)
	_, base_course := course.GetCourseBaseCourse(c)
	assessment := course.GetAssessment(c, course_name, assessment_name)
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
// @Success 200 {object} response.Response{data=[]dao.Questions_Student} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.StudentNotValidError} "not valid question, assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/question [get]
// @Security ApiKeyAuth
func QuestionAssessments_Handler(c *gin.Context) {
	// valid student
	course.CheckAssessmentTime(c)
	student := course.GetAssessmentStudent(c)

	response.SuccessResponse(c, student.ToRealQuestions())
}

// ReadAnswersAssessments_Handler godoc
// @Summary get answers based on the user
// @Schemes
// @Description get answers based on the user
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=dao.Student_Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.StudentNotValidError} "not valid question, assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers [get]
// @Security ApiKeyAuth
func ReadAnswersAssessments_Handler(c *gin.Context) {
	course.CheckAssessmentTime(c)
	student := course.GetAssessmentStudent(c)
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
// @Success 200 {object} response.Response{data=dao.Student_Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.StudentNotValidError} "not valid question, assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers [put]
// @Security ApiKeyAuth
func UploadAnswersAssessments_Handler(c *gin.Context) {
	course.CheckAssessmentTime(c)
	student := course.GetAssessmentStudent(c)

	var body dao.Answers_Upload_Validate
	body.Student = student
	validate.ValidateJson(c, &body)

	student.Answers = body.Answers
	instance, err := dao.CreateOrUpdateStudent(student)
	if err != nil {
		response.ErrMongoDBUpdateResponse(c, Student_Model)
	}

	response.SuccessResponse(c, instance.Answers)
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
// @Success 200 {object} response.Response{data=dao.Student_Questions} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/answers/struct [get]
// @Security ApiKeyAuth
func ReadAnswersStructAssessments_Handler(c *gin.Context) {
	course.CheckAssessmentTime(c)
	student := course.GetAssessmentStudent(c)
	response.SuccessResponse(c, student.ToAnwerStruct())
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
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
	token := jwt.GetEmail(c)

	global.Redis.Delay("generate", course_name, assessment_name, token.Email)
	response.SuccessResponse(c, "The assessment is generated in the back-end's queue system")
}

// ReadStatisticAssessments_Handler godoc
// @Summary get the assessment's score statistic result
// @Schemes
// @Description get the assessment's score statistic result
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=dao.Statistic} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/statistic [get]
// @Security ApiKeyAuth
func ReadStatisticAssessments_Handler(c *gin.Context) {
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
	assessment := course.GetAssessment(c, course_name, assessment_name)

	response.SuccessResponse(c, assessment.Statistic)
}

// CreateStatisticAssessments_Handler godoc
// @Summary read all students' assessment's score and generate statistic result
// @Schemes
// @Description read all students' assessment's score and generate statistic result
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Param data body dao.Statistic_Create true "body data"
// @Success 200 {object} response.Response{data=dao.Statistic} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/statistic [post]
// @Security ApiKeyAuth
func CreateStatisticAssessments_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
	assessment := course.GetAssessment(c, course_name, assessment_name)

	var data dao.Statistic_Create
	validate.ValidateJson(c, &data)

	var statistic dao.Statistic
	// var score []float64
	var sum_score = 0.0
	highest := 0.0
	lowest := math.MaxFloat64
	// method 1 access every student's score by refresh token in db
	users, _ := course.CourseUserData(c)
	for _, user := range users {
		token := jwt.UserRefreshByEmailHandler(c, user.Email)
		body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submissions")
		submissions := utils.Assessments_submissions_trans(string(body))
		if len(submissions) != 0 {
			statistic.Number += 1
			var best_index = 0
			if data.Best {
				var best_score = 0.0
				for i, submission := range submissions {
					if submission.TotalScore > best_score {
						best_score = submission.TotalScore
						best_index = i
					}
				}
			}
			highest = math.Max(highest, submissions[best_index].TotalScore)
			lowest = math.Min(lowest, submissions[best_index].TotalScore)
			sum_score += submissions[best_index].TotalScore
		}
	}

	statistic.Highest = highest
	statistic.Lowest = lowest
	statistic.Mean = math.Round((sum_score/float64(statistic.Number))*100) / 100
	statistic.Best = data.Best
	assessment.Statistic = statistic
	dao.UpdateExam(course_name, assessment_name, assessment)

	// method 2 only use score in student's db

	response.SuccessResponse(c, statistic)
}

// ReadStatusAssessments_Handler godoc
// @Summary get all students' assessment status
// @Schemes
// @Description get all students' assessment status
// @Tags exam
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Success 200 {object} response.Response{data=[]dao.Student_Status} "success"
// @Failure 400 {object} response.BadRequestResponse{error=response.CourseNoBaseCourseError} "no base course"
// @Failure 403 {object} response.ForbiddenResponse{error=response.ForbiddenError} "not instructor"
// @Failure 404 {object} response.NotValidResponse{error=response.AssessmentNotValidError} "not valid of assessment or course"
// @Router /courses/{course_name}/assessments/{assessment_name}/status [get]
// @Security ApiKeyAuth
func ReadStatusAssessments_Handler(c *gin.Context) {
	course_name, assessment_name := course.GetCourseAssessment(c)
	course.GetCourseBaseCourse(c)
	status, _ := dao.ReadAllStudentsStatus(course_name, assessment_name)

	response.SuccessResponse(c, status)
}

// Usersubmit_Handler godoc
// @Summary submit answer
// @Schemes
// @Description submit answer to Tango
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Submit} "desc"
// @Failure 500 {object} response.DBesponse{error=response.MongoDBReadError} "mongo error"
// @Failure 403 "reached the maximum number"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/assessments/{assessment_name}/submit [post]
func Usersubmit_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name, assessment_name := course.GetCourseAssessment(c)
	answertar_path := utils.Find_assessment_folder(c, strconv.Itoa(int(user.ID)), course_name, assessment_name)

	if status, _ := course.CheckSubmission(c); !status {
		response.ErrorResponseWithStatus(c, response.Error{Type: "Autolab", Message: "You have reached the maximum number of submissions!"}, http.StatusForbidden)
	} else {
		student, err := dao.ReadStudent(course_name, assessment_name, user.Email)
		if err != nil {
			response.ErrMongoDBReadResponse(c, "student")
		}

		flag := course.CreateAnswerFolder(answertar_path)
		if !flag {
			response.ErrFileNotValidResponse(c)
		}

		err = course.PrepareSolution(student, answertar_path)
		if err != nil {
			response.ErrFileNotValidResponse(c)
		}
		err = course.PrepareConfig(student, answertar_path)
		if err != nil {
			response.ErrFileNotValidResponse(c)
		}
		err = course.PrepareAnswer(student, answertar_path)
		if err != nil {
			response.ErrFileNotValidResponse(c)
		}

		flag = utils.MakeAnswertar(answertar_path)
		if flag {
			body := autolab.AutolabSubmitHandler(c, token, "/courses/"+course_name+"/assessments/"+assessment_name+"/submit", answertar_path+"answer.tar")
			autolab_resp := utils.User_submit_trans(string(body))
			student.Submitted = true
			student.Can_submit, student.SubmitNumber = course.CheckSubmission(c)
			dao.CreateOrUpdateStudent(student)
			response.SuccessResponse(c, autolab_resp)
			// if autolab_resp.Version == 0 {
			// 	response.ErrorResponseWithStatus(c, response.Error{Type: "Autolab", Message: "You are only allowed to submit once!"}, http.StatusForbidden)
			// } else {
			// 	response.SuccessResponse(c, autolab_resp)
			// }
		} else {
			response.ErrFileNotValidResponse(c)
		}
	}
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

// ReadBaseCourseRelation_Handler godoc
// @Summary get base course relationship
// @Schemes
// @Description read base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/base [get]
func ReadBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		base_course, flag := dao.ReadBaseCourseRelation(coursename)
		if flag {
			response.SuccessResponse(c, base_course)
		} else {
			response.ErrMySQLReadResponse(c, BaseCourseRelation_Model)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}

}

// CreateBaseCourseRelation_Handler godoc
// @Summary create new base course relationship
// @Schemes
// @Description create new base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		base				path	string	true	"Base Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Failure 400 "not valid"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/base/{base_name} [post]
func CreateBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base := c.Param("base_name")
	coursename := course.GetCourse(c)

	if !dao.ValidBaseCourseRelation(coursename) {
		flag := dao.CreateBaseCourseRelation(coursename, base)

		if flag {
			response.SuccessResponse(c, coursename+" "+base)
		} else {
			response.ErrMySQLCreateResponse(c, BaseCourseRelation_Model)
		}
	} else {
		response.ErrBasecourseRelationRecreatedResponse(c)
	}
}

// UpdateBaseCourseRelation_Handler godoc
// @Summary update base course relationship
// @Schemes
// @Description update base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		base				path	string	true	"Base Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Failure 400 "not valid"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/base/{base_name} [put]
func UpdateBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base := c.Param("base_name")
	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		if flag, _ := dao.ValidateAssessmentByCourse(coursename); flag {
			err := dao.UpdateBaseCourseRelation(coursename, base)

			if err != nil {
				response.ErrMySQLUpdateResponse(c, BaseCourseRelation_Model)
			} else {
				response.SuccessResponse(c, coursename+" "+base)
			}
		} else {
			response.ErrBasecourseRelationNotValidResponse(c)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}

}

// DeleteBaseCourseRelation_Handler godoc
// @Summary delete base course relationship
// @Schemes
// @Description delete base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 204
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Failure 400 "not valid"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /courses/{course_name}/base [delete]
func DeleteBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		if flag, _ := dao.ValidateAssessmentByCourse(coursename); flag {
			err := dao.DeleteBaseCourseRelation(coursename)

			if err != nil {
				response.ErrMySQLDeleteResponse(c, BaseCourse_Model)
			} else {
				response.NonContentResponse(c)
			}
		} else {
			response.ErrBasecourseRelationNotValidResponse(c)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}
}
