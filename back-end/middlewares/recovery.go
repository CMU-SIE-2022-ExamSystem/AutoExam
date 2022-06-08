package middlewares

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/CMU-SIE-2022-ExamSystem/exam-system/error"
	"github.com/CMU-SIE-2022-ExamSystem/exam-system/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err_type := (reflect.TypeOf(err) == reflect.TypeOf("123")); !err_type {
				err_temp := err.(error.ErrorResponse)
				message := gin.H{
					"error": gin.H{
						err_temp.Scope: err_temp.Message,
					},
				}
				c.JSON(err_temp.Status, message)
				data, _ := json.Marshal(message)
				c.Error(errors.New(string(data)))
				c.AbortWithStatus(http.StatusInternalServerError)
			} else {
				if err != nil {
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
								brokenPipe = true
							}
						}
					}

					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					if brokenPipe {
						global.Lg.Error(c.Request.URL.Path,
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
						c.Abort()
						return
					}

					if stack {
						zap.L().Error("[Recovery from panic]",
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
							zap.String("stack", string(debug.Stack())),
						)
					} else {
						zap.L().Error("[Recovery from panic]",
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
					}
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}

		}()
		c.Next()
	}
}
