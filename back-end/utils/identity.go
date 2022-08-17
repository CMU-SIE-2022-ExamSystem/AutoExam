package utils

import (
	"errors"
	"os"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Map_course_info(user_courses []models.User_Courses) map[string]models.User_Courses {
	courses__info_map := make(map[string]models.User_Courses, len(user_courses))
	for i := 0; i < len(user_courses); i++ {
		courses__info_map[user_courses[i].Name] = user_courses[i]
	}
	return courses__info_map
}

func Map_user_authlevel(user_courses []models.User_Courses) map[string]string {
	user_courses_map := make(map[string]string, len(user_courses))
	for i := 0; i < len(user_courses); i++ {
		user_courses_map[user_courses[i].Name] = user_courses[i].Auth_level
	}
	return user_courses_map
}

func Map_DBcheck(user_courses_map map[string]string) map[string]dao.Strings {
	db_map := make(map[string]dao.Strings, 3)
	for k, v := range user_courses_map {
		db_map[v] = append(db_map[v], k)
	}
	return db_map
}

//user_id type should be uint and the uint to string: strconv.Itoa(int(user_id)
func Find_assessment_folder(c *gin.Context, user_id string, course string, assessment string) string {
	relative_path := "./tmp/assessment/"
	if _, err := os.Stat(relative_path + course + "/"); err == nil {
		if _, err := os.Stat(relative_path + course + "/" + assessment + "/"); err == nil {
			if _, err := os.Stat(relative_path + course + "/" + assessment + "/" + user_id + "/"); err == nil {
				color.Yellow("folder already exists!")
			} else if errors.Is(err, os.ErrNotExist) {
				err = os.Mkdir(relative_path+course+"/"+assessment+"/"+user_id+"/", 0777)
				if err != nil {
					response.ErrFileNotValidResponse(c)
				}
				color.Yellow("create folder for user!")
			}
		} else if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(relative_path+course+"/"+assessment+"/"+user_id+"/", 0777)
			if err != nil {
				response.ErrFileNotValidResponse(c)
			}
			color.Yellow("create folder for assessment!")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(relative_path+course+"/"+assessment+"/"+user_id+"/", 0777)
		if err != nil {
			response.ErrFileNotValidResponse(c)
		}
		color.Yellow("create folder for course!")
	}
	return relative_path + course + "/" + assessment + "/" + user_id + "/"
}
