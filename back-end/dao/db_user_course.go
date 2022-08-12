package dao

import (
	"database/sql/driver"
	"encoding/json"
	"sort"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

type Strings []string

type UserCourseRelationship struct {
	Id                  uint
	ClassesAsStudent    Strings
	ClassesAsTA         Strings
	ClassesAsInstructor Strings
}

func (c *Strings) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (c Strings) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func compareTwoStringSlice(slice1 Strings, slice2 Strings) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	sort.Strings(slice1)
	sort.Strings(slice2)
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func find_userinfo(id uint) (*UserCourseRelationship, bool) {
	var user UserCourseRelationship
	rows := global.DB.Where(&UserCourseRelationship{Id: id}).Find(&user)
	if rows.RowsAffected < 1 {
		return &user, false
	}
	return &user, true
}

func insertUserCourse(c *gin.Context, userCourse UserCourseRelationship) {
	if err := global.DB.Create(&userCourse).Error; err != nil {
		response.ErrDBResponse(c, "Can not insert user course relationship.")
	}
}

func updateUserCourse(c *gin.Context, userCourseInstance UserCourseRelationship, user_id uint, mapFromAutoLab map[string]Strings) {
	OriginalStudentCourses := userCourseInstance.ClassesAsStudent
	NewStudentCourses := mapFromAutoLab["student"]
	if !compareTwoStringSlice(OriginalStudentCourses, NewStudentCourses) {
		updateHelper(c, user_id, NewStudentCourses, "classes_as_student")
	}

	OriginalTACourses := userCourseInstance.ClassesAsTA
	NewTACourses := mapFromAutoLab["course_assistant"]
	if !compareTwoStringSlice(OriginalTACourses, NewTACourses) {
		updateHelper(c, user_id, NewTACourses, "classes_as_ta")
	}

	OriginalInstructorCourses := userCourseInstance.ClassesAsInstructor
	NewInstructorCourses := mapFromAutoLab["instructor"]
	if !compareTwoStringSlice(OriginalInstructorCourses, NewInstructorCourses) {
		updateHelper(c, user_id, NewInstructorCourses, "classes_as_instructor")
	}
}

func updateHelper(c *gin.Context, user_id uint, newCourses Strings, tag string) {

	if err := global.DB.Model(new(UserCourseRelationship)).Where("id=?", user_id).Update(tag, newCourses).Error; err != nil {
		response.ErrDBResponse(c, "Can not insert user course relationship.")
	}
}

func Check_authlevel(user_id uint, class_name string) string {
	userCourseInstanceAddress, flag := find_userinfo(user_id)
	userCourseInstance := *userCourseInstanceAddress
	if flag {
		User_InstructorCourses := userCourseInstance.ClassesAsInstructor
		for _, v := range User_InstructorCourses {
			if class_name == v {
				// color.Yellow("this user is instructor")
				return "instructor"
			}
		}

		User_TACourses := userCourseInstance.ClassesAsTA
		for _, v := range User_TACourses {
			if class_name == v {
				// color.Yellow("this user is a TA")
				return "course_assistant"
			}
		}

		User_StudentCourses := userCourseInstance.ClassesAsStudent
		for _, v := range User_StudentCourses {
			if class_name == v {
				// color.Yellow("this user is a student")
				return "student"
			}
		}
		// color.Yellow("this user is not one of the three roles of this course")
		return ""
	} else {
		// this user not existed
		// color.Yellow("this user is not existing")
		return ""
	}
}

func UpdateOrAddUser(c *gin.Context, user_id uint, mapFromAutoLab map[string]Strings) {
	userCourseInstance, flag := find_userinfo(user_id)
	if flag {
		// user already exists, check and update the information
		updateUserCourse(c, *userCourseInstance, user_id, mapFromAutoLab)
	} else {
		// insert this user to mysql
		NewStudentCourses := mapFromAutoLab["student"]
		NewTACourses := mapFromAutoLab["course_assistant"]
		NewInstructorCourses := mapFromAutoLab["instructor"]
		newUser := UserCourseRelationship{
			user_id, NewStudentCourses, NewTACourses, NewInstructorCourses,
		}
		insertUserCourse(c, newUser)
	}
}

func Get_all_courses(user_id uint) (courses Strings) {
	userCourseInstanceAddress, _ := find_userinfo(user_id)
	userCourseInstance := *userCourseInstanceAddress
	courses = append(userCourseInstance.ClassesAsInstructor, userCourseInstance.ClassesAsTA...)
	courses = append(courses, userCourseInstance.ClassesAsStudent...)
	return
}

func Get_Baseauth(user_id uint) bool {
	userCourseInstanceAddress, flag := find_userinfo(user_id)
	userCourseInstance := *userCourseInstanceAddress
	if flag {
		User_InstructorCourses := userCourseInstance.ClassesAsInstructor
		if User_InstructorCourses != nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
