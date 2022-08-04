package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/autolab"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
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

	path := utils.Find_assessment_folder(c, user_id, course, assessment)
	fmt.Println(path)
}

// This is for all user auth-level in a course
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

func Autograder_Test(c *gin.Context) {
	question_type := c.Param("question_type")
	// dao.SearchAndStore_grader(c, question_type, "./autograder/exec/autograders/")

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("python", "main.py", question_type)
	cmd.Dir = "./autograder/exec/"
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	os.Remove("./autograder/exec/autograders/" + question_type + ".py")
	if err != nil {
		// color.Red(err.Error())
		// color.Red(stdout.String())
		// color.Red(stderr.String())
		response.ErrorInternaWithData(c, err.Error(), stdout.String()+stderr.String())
	} else {
		// color.Yellow(stdout.String())
		response.SuccessResponse(c, stdout.String())
	}
}

func Answertar_Test(c *gin.Context) {
	path := "./tmp/"
	flag := utils.MakeAnswertar(path)
	if flag {
		response.SuccessResponse(c, "Get it!")
	} else {
		response.ErrFileNotValidResponse(c)
	}
}
