package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func getUsers() []models.User {
	var users []models.User
	global.DB.Find(&users)
	return users
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} response.SwaggerResponse  "desc"
// @Security ApiKeyAuth
// @Router /test/users/ [get]
func GetUsers(c *gin.Context) {
	users := getUsers()
	fmt.Println("=====================")
	fmt.Println(jwt.GetEmail(c))
	fmt.Println("=====================")
	response.SuccessResponse(c, users)
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} response.SwaggerResponse "desc"
// @Security ApiKeyAuth
// @Router /test/user/ [get]
func GetUser(c *gin.Context) {
	users := getUsers()
	response.SuccessResponse(c, users)
	// c.JSON(http.StatusOK, users)
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} models.User_Token "desc"
// @Router /test/login/ [get]
func Login(c *gin.Context) {
	token := jwt.CreateToken(c, 1, "test@gmail")
	c.JSON(http.StatusOK, token)
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} response.SwaggerResponse "desc"
// @Security ApiKeyAuth
// @Router /test/cookie/ [get]
func CookieTest(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", int(jwt.Expire_time), "/", "localhost", false, true)
	}
	fmt.Println("============================")
	fmt.Printf("Cookie value: %s \n", cookie)
	fmt.Println("============================")
}

func FolderTest(c *gin.Context) {
	user_id := c.Param("user_id")
	course := c.Param("course")
	assessment := c.Param("assessment")

	path := utils.Find_folder(c, user_id, course, assessment)
	fmt.Println(path)
}

//todo: This is for all user auth-level in a course
func Course_all_Test(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	course_name := c.Param("course_name")

	body := autolab.AutolabGetHandler(c, token, "/courses/"+course_name+"/course_user_data")
	// fmt.Println(string(body))

	if strings.Contains(string(body), "error") {
		err_response := utils.Course_user_err_trans(string(body))
		response.ErrUnauthResponse(c, err_response.Error)
	} else {
		autolab_resp := utils.Course_user_trans(string(body))
		response.SuccessResponse(c, autolab_resp)
	}
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /test/exam [get]
func Test_exam(c *gin.Context) {
	user := jwt.GetEmail(c)
	courses := dao.Get_all_courses(user.ID)
	response.SuccessResponse(c, courses)
}

// AuthInfo godoc
// @Summary test
// @Schemes
// @Description test
// @Tags test
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=models.Submit} "desc"
// @Param		course_name			path	string	true	"Course Name"
// @Param		assessment_name		path	string	true	"Assessment name"
// @Security ApiKeyAuth
// @Router /test/{course_name}/assessments/{assessment_name}/submit [post]
//todo: This is for student to take exam
func Take_exam_Test(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)

	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	color.Yellow(string(course_name))
	color.Yellow(string(assessment_name))
	color.Yellow(strconv.Itoa(int((user.ID))))
}
