package utils

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)

func CreateToken(c *gin.Context, Id int, token string) string {
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:    uint(Id),
		Token: token,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24,
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
