package dao

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

func GetAccessToken(id uint) string {
	user := models.User{ID: id}
	global.DB.Find(&user)
	token := user.Access_token

	return token
}
