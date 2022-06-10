package controller

import (
	"fmt"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

var user models.User

func FindUserInfo(email string) (*models.User, bool) {
	rows := global.DB.Where(&models.User{Email: email}).Find(&user)
	fmt.Println(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}
