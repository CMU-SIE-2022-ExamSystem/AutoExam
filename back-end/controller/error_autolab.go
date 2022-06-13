package controller

import (
	"net/http"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/response"
	"github.com/gin-gonic/gin"
)

func Autolab_Error_Hander(c *gin.Context, resp *http.Response, err error) {
	if err != nil {
		response.ErrUnauthResponse(c, "There may be something wrong with Autolab's web server, please try again later.")
	}

	status := resp.StatusCode
	if status >= http.StatusOK && status <= 299 {
		return
	} else {
		if status == http.StatusUnauthorized {
		} else {
			response.ErrorInternalResponse(c, response.Error{Type: "Autolab", Message: "Unknown error"})
		}
	}
}
