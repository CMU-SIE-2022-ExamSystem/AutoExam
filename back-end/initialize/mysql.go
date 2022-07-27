package initialize

import (
	"fmt"
	"os"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() {
	mysqlInfo := global.Settings.Mysqlinfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
		mysqlInfo.Port, mysqlInfo.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		color.Red("========================================================")
		color.Red("MySQL connection is not correct, please check settings-dev.yaml file and rerun the server again")
		color.Red("========================================================")
		os.Exit(3)
	}

	global.DB = db

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&dao.UserCourseRelationship{})
	db.AutoMigrate(&dao.Blanks{})
	db.AutoMigrate(&dao.PythonFile{})
}
