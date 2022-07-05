package controller

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

//todo: need some works
func Examconfig_Handler(c *gin.Context) {
	course_name := c.Param("course_name")
	assessment_name := c.Param("assessment_name")

	data, _ := ioutil.ReadAll(c.Request.Body)
	color.Yellow(string(course_name))
	color.Yellow(string(assessment_name))
	color.Yellow(string(data))
}
