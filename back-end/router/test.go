package router

import (
	"github.com/gin-gonic/gin"
)

func TestRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("test")
	{
		AuthRouter.GET("/", func(c *gin.Context) {
			panic("123")
		})
	}
}
