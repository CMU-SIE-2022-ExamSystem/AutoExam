package utils

import (
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/middlewares"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, Id uint, token string) string {
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

func GetToken(c *gin.Context) (user models.UserToken) {

	token, err := c.Get("token")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no toke in gin.Context"})
	}
	id, err := c.Get("userId")
	if !err {
		response.ErrorInternalResponse(c, response.Error{Type: response.Authentication, Message: "there is no toke in gin.Context"})
	}

	user = models.UserToken{ID: id.(uint), Token: token.(string)}
	return
}
