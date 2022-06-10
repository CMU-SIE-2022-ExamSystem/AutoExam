package test

import (
	"fmt"
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/authentication"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = global.DB

func createUser() {
	var user models.User
	var users []models.User
	// 查看插入后的全部元素
	fmt.Printf("插入后元素:\n")
	db.Find(&users)
	fmt.Println(users)
	// 查询一条记录
	db.First(&user, "name = ?", "bgbiao")
	fmt.Println("查看查询记录:", user)
	// 更新记录(基于查出来的数据进行更新)
	db.Model(&user).Update("name", "biaoge")
	fmt.Println("更新后的记录:", user)
	// 删除记录
	db.Delete(&user)
	// 查看全部记录
	fmt.Println("查看全部记录:")
	db.Find(&users)
	fmt.Println(users)
}

func getUsers() []models.User {
	var users []models.User
	global.DB.Find(&users)
	return users
}

func getUser(id uint, email string) models.User {
	var user models.User
	if id == 0 {
		user = models.User{Email: email}
	} else {
		user = models.User{ID: id, Email: email}
	}
	global.DB.Find(&user)
	return user
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
	fmt.Println(authentication.GetToken(c))
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
// @Success 200 {object} models.UserToken "desc"
// @Router /test/login/ [get]
func Login(c *gin.Context) {
	token := authentication.CreateToken(c, 1, "test@gmail")
	c.JSON(http.StatusOK, token)
}
