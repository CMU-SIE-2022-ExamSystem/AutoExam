package controller

import (
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/course"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/dao"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/jwt"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

const (
	BaseCourse_Model         = "base course"
	BaseCourseRelation_Model = "base course relation"
)

// ReadAllBaseCourse_Handler godoc
// @Summary get all base courses
// @Schemes
// @Description read all base courses in database
// @Tags basecourse
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.Course} "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Security ApiKeyAuth
// @Router /basecourse/list [get]
func ReadAllBaseCourse_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)

	courses, err := dao.ReadAllBaseCourse()
	if err != nil {
		response.ErrMySQLReadAllResponse(c, BaseCourse_Model)
	}
	response.SuccessResponse(c, courses)
}

// CreateBaseCourse_Handler godoc
// @Summary create a new base course
// @Schemes
// @Description create a new base course
// @Tags basecourse
// @Accept json
// @Produce json
// @Param		base		path	string	true	"Base Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Failure 400 "not valid"
// @Security ApiKeyAuth
// @Router /basecourse/{base} [post]
func CreateBaseCourse_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)

	name := c.Param("base")

	if !dao.ValidBaseCourse(name) {
		flag := dao.CreateBaseCourse(name)

		if flag {
			response.SuccessResponse(c, name)
		} else {
			response.ErrMySQLCreateResponse(c, BaseCourse_Model)
		}
	} else {
		response.ErrBasecourseNotValidResponse(c)
	}
}

// UpdateBaseCourse_Handler godoc
// @Summary update a base course name
// @Schemes
// @Description update a base course name
// @Tags basecourse
// @Accept json
// @Produce json
// @Param		base		path	string	true	"Base Course Name"
// @Param		new			path	string	true	"New Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Failure 400 "not valid"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /basecourse/{base}/{new} [put]
func UpdateBaseCourse_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)

	name := c.Param("base")
	new_name := c.Param("new")

	if dao.ValidBaseCourse(name) {
		flag, err := dao.UpdateBaseCourse(name, new_name)

		if err != nil {
			response.ErrMySQLUpdateResponse(c, BaseCourse_Model)
		} else {
			if flag {
				response.SuccessResponse(c, new_name)
			} else {
				response.ErrBasecourseNotValidResponse(c)
			}
		}
	} else {
		response.ErrBasecourseNotExistsResponse(c)
	}
}

// DeleteBaseCourse_Handler godoc
// @Summary delete a base course
// @Schemes
// @Description delete a base course
// @Tags basecourse
// @Accept json
// @Produce json
// @Param		base		path	string	true	"Base Course Name"
// @Success 204
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Failure 400 "not valid"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /basecourse/{base} [delete]
func DeleteBaseCourse_Handler(c *gin.Context) {
	jwt.Check_Baselevel(c)

	name := c.Param("base")

	if dao.ValidBaseCourse(name) {
		flag, err := dao.DeleteBaseCourse(name)

		if err != nil {
			response.ErrMySQLDeleteResponse(c, BaseCourse_Model)
		} else {
			if flag {
				response.NonContentResponse(c)
			} else {
				response.ErrBasecourseNotValidResponse(c)
			}
		}
	} else {
		response.ErrBasecourseNotExistsResponse(c)
	}
}

// ReadBaseCourseRelation_Handler godoc
// @Summary get base course relationship
// @Schemes
// @Description read base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLReadAllError} "mysql error"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /{course_name}/base [get]
func ReadBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)
	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		base_course, flag := dao.ReadBaseCourseRelation(coursename)
		if flag {
			response.SuccessResponse(c, base_course)
		} else {
			response.ErrMySQLReadResponse(c, BaseCourseRelation_Model)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}

}

// CreateBaseCourseRelation_Handler godoc
// @Summary create new base course relationship
// @Schemes
// @Description create new base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		base				path	string	true	"Base Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLCreateError} "mysql error"
// @Failure 400 "not valid"
// @Security ApiKeyAuth
// @Router /{course_name}/{base} [post]
func CreateBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base := c.Param("base")
	coursename := course.GetCourse(c)

	if !dao.ValidBaseCourseRelation(coursename) {
		flag := dao.CreateBaseCourseRelation(coursename, base)

		if flag {
			response.SuccessResponse(c, coursename+" "+base)
		} else {
			response.ErrMySQLCreateResponse(c, BaseCourseRelation_Model)
		}
	} else {
		response.ErrBasecourseRelationRecreatedResponse(c)
	}
}

// UpdateBaseCourseRelation_Handler godoc
// @Summary update base course relationship
// @Schemes
// @Description update base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Param		base				path	string	true	"Base Course Name"
// @Success 200 "success"
// @Failure 500 {object} response.DBesponse{error=response.MySQLUpdateError} "mysql error"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /{course_name}/{base} [put]
func UpdateBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	base := c.Param("base")
	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		if flag, _ := dao.ValidateAssessmentByCourse(coursename); flag {
			err := dao.UpdateBaseCourseRelation(coursename, base)

			if err != nil {
				response.ErrMySQLUpdateResponse(c, BaseCourseRelation_Model)
			} else {
				response.SuccessResponse(c, coursename+" "+base)
			}
		} else {
			response.ErrBasecourseRelationNotValidResponse(c)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}

}

// DeleteBaseCourseRelation_Handler godoc
// @Summary delete base course relationship
// @Schemes
// @Description delete base course relationship
// @Tags base
// @Accept json
// @Produce json
// @Param		course_name			path	string	true	"Course Name"
// @Success 204
// @Failure 500 {object} response.DBesponse{error=response.MySQLDeleteError} "mysql error"
// @Failure 404 "not exists"
// @Security ApiKeyAuth
// @Router /{course_name}/base [delete]
func DeleteBaseCourseRelation_Handler(c *gin.Context) {
	jwt.Check_authlevel_Instructor(c)

	coursename := course.GetCourse(c)

	if dao.ValidBaseCourseRelation(coursename) {
		if flag, _ := dao.ValidateAssessmentByCourse(coursename); flag {
			err := dao.DeleteBaseCourseRelation(coursename)

			if err != nil {
				response.ErrMySQLDeleteResponse(c, BaseCourse_Model)
			} else {
				response.NonContentResponse(c)
			}
		} else {
			response.ErrBasecourseRelationNotValidResponse(c)
		}
	} else {
		response.ErrBasecourseRelationNotExistsResponse(c)
	}
}
