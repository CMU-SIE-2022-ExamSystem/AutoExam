package controller

import (
	"fmt"
	"net/http"

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

	path := utils.Find_folder(c, user_id, course, assessment)
	fmt.Println(path)
}
