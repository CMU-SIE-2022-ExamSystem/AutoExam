package dao

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/models"
)

func insertBaseCourse(course models.Course) error {
	if err := global.DB.Create(&course).Error; err != nil {
		// color.Yellow("Insert course table failed")
		return err
	}
	// color.Yellow("Insert course table successfully")
	return nil
}

func insertBaseCourseRelation(course models.Base_Course_Relationship) error {
	if err := global.DB.Create(&course).Error; err != nil {
		// color.Yellow("Insert course relationship table failed")
		return err
	}
	// color.Yellow("Insert course relationship table successfully")
	return nil
}

func ValidBaseCourse(base string) bool {
	var course models.Course
	rows := global.DB.Where("name = ?", base).Find(&course)
	if rows.RowsAffected < 1 {
		return false
	} else {
		return true
	}
}

func ValidBaseCourseRelation(coursename string) bool {
	var course models.Base_Course_Relationship
	rows := global.DB.Where("course_name = ?", coursename).Find(&course)
	if rows.RowsAffected < 1 {
		return false
	} else {
		return true
	}
}

func usedBaseCourse(base string) bool {
	var course models.Base_Course_Relationship
	rows := global.DB.Where("base_course = ?", base).Find(&course)
	if rows.RowsAffected < 1 {
		return false
	} else {
		return true
	}
}

func CreateBaseCourse(name string) bool {
	var course models.Course
	rows := global.DB.Where("name = ?", name).Find(&course)
	if rows.RowsAffected < 1 {
		new_course := models.Course{
			Name: name,
		}
		if err := insertBaseCourse(new_course); err != nil {
			// color.Yellow("There was an error adding this base course, please try again.")
			return false
		}
		// color.Yellow("Create course table successfully.")
		return true
	} else {
		// color.Yellow("This base course course already exists.")
		return false
	}
}

func ReadAllBaseCourse() ([]models.Course, error) {
	var courses []models.Course
	result := global.DB.Find(&courses)
	if result.Error != nil {
		// color.Yellow("Read all course table failed")
		return courses, result.Error
	}
	// color.Yellow("Read all course table successfully")
	return courses, nil
}

func UpdateBaseCourse(name, new_name string) (bool, error) {
	if !usedBaseCourse(name) {
		var course models.Course
		if err := global.DB.Where("name = ?", name).Find(&course).Update("name", new_name).Error; err != nil {
			// color.Yellow("Update course table failed")
			return false, err
		}
		// color.Yellow("Update course table successfully")
		return true, nil
	}
	return false, nil
}

func DeleteBaseCourse(name string) (bool, error) {
	if !usedBaseCourse(name) {
		result := global.DB.Where(&models.Course{Name: name}).Delete(&models.Course{})
		return true, result.Error
	} else {
		return false, nil
	}
}

func CreateBaseCourseRelation(coursename, base string) bool {
	if !ValidBaseCourse(base) {
		new_course := models.Course{
			Name: base,
		}
		if err := insertBaseCourse(new_course); err != nil {
			// color.Yellow("There was an error adding this base course, please try again.")
			return false
		}
	}

	var course models.Base_Course_Relationship
	rows := global.DB.Where("course_name = ?", coursename).Find(&course)
	if rows.RowsAffected < 1 {
		course_relation := models.Base_Course_Relationship{
			Course_name: coursename,
			Base_course: base,
		}

		if err := insertBaseCourseRelation(course_relation); err != nil {
			// color.Yellow("There was an error adding this course relationship, please try again.")
			return false
		} else {
			return true
		}
	} else {
		// color.Yellow("This course relationship already exists.")
		return false
	}
}

func ReadBaseCourseRelation(coursename string) (string, bool) {
	var course models.Base_Course_Relationship
	rows := global.DB.Where(&models.Base_Course_Relationship{Course_name: coursename}).Find(&course)
	if rows.RowsAffected < 1 {
		// color.Yellow("Read course relationship table failed.")
		return course.Base_course, false
	}
	// color.Yellow("Read course relationship table successfully.")
	return course.Base_course, true
}

func UpdateBaseCourseRelation(coursename, new_base string) error {
	if !ValidBaseCourse(new_base) {
		new_course := models.Course{
			Name: new_base,
		}
		if err := insertBaseCourse(new_course); err != nil {
			// color.Yellow("There was an error adding this base course, please try again.")
			return err
		}
	}

	var course models.Base_Course_Relationship
	if err := global.DB.Where("course_name = ?", coursename).Find(&course).Update("base_course", new_base).Error; err != nil {
		// color.Yellow("Update course relationship table failed.")
		return err
	}
	// color.Yellow("Update course relationship table successfully.")
	return nil
}

func DeleteBaseCourseRelation(coursename string) error {
	result := global.DB.Where(&models.Base_Course_Relationship{Course_name: coursename}).Delete(&models.Base_Course_Relationship{})
	return result.Error
}
