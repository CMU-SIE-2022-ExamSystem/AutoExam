package jwt

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	Expire_time int64 = 60 * 60 * 24
)

func CreateToken(c *gin.Context, Id uint, email string) string {
	j := NewJWT()
	claims := CustomClaims{
		ID:    uint(Id),
		Email: email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: utils.GetNowTime(),
			ExpiresAt: utils.GetNowTime() + Expire_time,
			Issuer:    "test",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "token is not created correctly, please try again"})
		return ""
	}
	return token
}

func GetEmail(c *gin.Context) (user models.User_Token) {

	email, err := c.Get("email")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no token in gin.Context"})
	}
	id, err := c.Get("userId")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no token in gin.Context"})
	}
	user = models.User_Token{ID: id.(uint), Email: email.(string)}
	return
}

func GetAccessToken(id uint) string {
	user := models.User{ID: id}
	global.DB.Find(&user)
	token := user.Access_token

	return token
}
