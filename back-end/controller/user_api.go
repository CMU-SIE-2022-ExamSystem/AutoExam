package controller

import (
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

func find_userinfo(email string) (*models.User, bool) {
	var user models.User
	rows := global.DB.Where(&models.User{Email: email}).Find(&user)
	// fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func set_cookie(c *gin.Context, cookie string) {
	// c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("gin_cookie", "Bearer "+cookie, int(jwt.Expire_time), "/", "", false, false)
}

func Userinfo_Handler(c *gin.Context, autolab_resp models.Autolab_Response) {
	body := autolab.AutolabGetHandler(c, autolab_resp.Access_token, "/user")
	// fmt.Println(string(body))

	userinfo_resp := utils.User_info_trans(string(body))

	user, flag := find_userinfo(userinfo_resp.Email)

	if flag {
		color.Yellow("User is already in our DB!")
		user.Access_token = autolab_resp.Access_token
		user.Refresh_token = autolab_resp.Refresh_token
		user.Create_at = utils.GetNowTime()
		user.Expires_in = autolab_resp.Expires_in

		global.DB.Save(&user)
		jwt_token := jwt.CreateToken(c, user.ID, user.Email)
		set_cookie(c, jwt_token)

		user_info := models.User_Info_Front{Token: jwt_token, First_name: user.First_name, Last_name: user.Last_name}
		response.SuccessResponse(c, user_info)
	} else {
		color.Yellow("User is not in our DB!")
		new_user := models.User{
			Email:         userinfo_resp.Email,
			First_name:    userinfo_resp.First_name,
			Last_name:     userinfo_resp.Last_name,
			Access_token:  autolab_resp.Access_token,
			Refresh_token: autolab_resp.Refresh_token,
			Create_at:     utils.GetNowTime(),
			Expires_in:    autolab_resp.Expires_in,
		}

		global.DB.Create(&new_user)
		jwt_token := jwt.CreateToken(c, new_user.ID, new_user.Email)
		set_cookie(c, jwt_token)

		user_info := models.User_Info_Front{Token: jwt_token, First_name: new_user.First_name, Last_name: new_user.Last_name}
		response.SuccessResponse(c, user_info)
	}
}

// Usercourses_Handler godoc
// @Summary get user courses
// @Schemes
// @Description get user courses from autolab
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.User_Courses} "success"
// @Security ApiKeyAuth
// @Router /user/courses [get]
func Usercourses_Handler(c *gin.Context) {
	user_email := jwt.GetEmail(c)
	user := models.User{ID: user_email.ID}
	global.DB.Find(&user)
	token := user.Access_token

	body := autolab.AutolabGetHandler(c, token, "/courses")
	// fmt.Println(string(body))

	autolab_resp := utils.User_courses_trans(string(body))

	autolab_map := utils.Map_user_authlevel(autolab_resp)
	autolab_map_DB := utils.Map_DBcheck(autolab_map)
	dao.UpdateOrAddUser(c, user.ID, autolab_map_DB)

	response.SuccessResponse(c, autolab_resp)
}
