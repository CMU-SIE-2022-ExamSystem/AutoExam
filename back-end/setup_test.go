package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/initialize"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/utils"
	"github.com/DATA-DOG/go-sqlmock"
	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mock sqlmock.Sqlmock

func SetupTest(tb testing.TB, role string) (func(tb testing.TB), *gin.Engine, string) {
	log.Println("setup test")

	server := initialize.SetupServer()

	if role != "Instructor" && role != "TA" && role != "Student" {
		tb.Fatal("wrong role when setup the test")
	}

	migrateToken(tb, role)

	return func(tb testing.TB) {
		log.Println("teardown test")
	}, server, ""
}

func CreateToken(Id uint, email string) (string, error) {
	j := jwt.NewJWT()
	claims := jwt.CustomClaims{
		ID:    uint(Id),
		Email: email,
		StandardClaims: jwt_go.StandardClaims{
			NotBefore: utils.GetNowTime(),
			ExpiresAt: utils.GetNowTime() + jwt.Expire_time,
			Issuer:    "test",
		},
	}
	token, err := j.CreateToken(claims)
	return token, err
}

func CreateInstructorToken() (string, error) {
	return CreateToken(1, "sie2022@andrew.cmu.edu")
}

func CreateStudentToken() (string, error) {
	return CreateToken(2, "student@andrew.cmu.edu")
}

func CreateTAToken() (string, error) {
	return CreateToken(3, "ta@andrew.cmu.edu")
}

func migrateToken(tb testing.TB, role string) {
	var user models.User
	if role == "Instructor" {
		user = models.User{ID: 1}
	} else if role == "TA" {
		user = models.User{ID: 2}
	} else if role == "Student" {
		user = models.User{ID: 3}
	}
	global.DB.Find(&user)

	var db *sql.DB
	var err error
	db, Mock, err = sqlmock.New()

	if err != nil {
		tb.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		tb.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	Mock.ExpectBegin()
	Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `user` (`email`, `access_token`, `refresh_token`, `first_name`,`last_name`, `create_at`, `expires_at`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(user.Email, user.Access_token, user.Refresh_token, user.First_name, user.Last_name, user.Create_at, user.Expires_in).
		WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()

	defer db.Close()
	user.ID = 0
	global.DB.Create(&user)

	check := models.User{
		ID: 1,
	}
	err = global.DB.First(&check).Error
	fmt.Println(check)
	fmt.Println("==================")
	fmt.Println(err)
	fmt.Println("==================")
	// assert.NoError(tb, err)
}
