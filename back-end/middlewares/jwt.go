package middlewares

import (
	"errors"
	"strings"
	"time"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	ID    uint   //
	Email string //
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			response.ErrUnauthResponse(c, "please use token to call api")
			c.Abort()
			return
		}
		token := strings.Split(auth, " ")[1]
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				if err == ErrTokenExpired {
					response.ErrUnauthResponse(c, "the token is expired")
					c.Abort()
					return
				}
			}
			response.ErrUnauthResponse(c, "couldn't handle this token")
			c.Abort()
			return
		}
		// fmt.Println(c)
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Set("email", claims.Email)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Settings.JWTKey.SigningKey),
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid

	} else {
		return nil, ErrTokenInvalid

	}

}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
