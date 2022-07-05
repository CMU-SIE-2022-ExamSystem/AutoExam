package controller

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

//todo: need some works
func Question_Handler(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	color.Yellow(string(data))
}
