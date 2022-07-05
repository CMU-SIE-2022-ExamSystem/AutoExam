package course

import (
	"github.com/gin-gonic/gin"
)

func GetCourseAssessment(c *gin.Context) (string, string) {
	course := c.Param("course_name")
	assessment := c.Param("assessment_name")

	// should validate coures and assessment

	// assessment no numeric start

	return course, assessment
}
