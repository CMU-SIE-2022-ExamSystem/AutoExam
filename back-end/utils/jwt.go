package utils

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, Id uint, email string) string {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:    uint(Id),
		Email: email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: GetNowTime(),
			ExpiresAt: GetNowTime() + 60*60*24,
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

func GetEmail(c *gin.Context) (user models.UserToken) {

	email, err := c.Get("email")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no token in gin.Context"})
	}
	id, err := c.Get("userId")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no token in gin.Context"})
	}

	user = models.UserToken{ID: id.(uint), Email: email.(string)}
	return
}
