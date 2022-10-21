package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	. "github.com/wkrzyzanowski/notion-task-integrator/logging"
	"github.com/wkrzyzanowski/notion-task-integrator/server"
	"io/ioutil"
)

func NewRequestLoggerMiddleware() server.ApiMiddleware {
	return server.ApiMiddleware{
		Name:     "Logging Middleware",
		Function: logRequest(),
	}
}

func logRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Read the Body content
		var bodyBytes []byte
		if ctx.Request.GetBody != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}
		// Restore the io.ReadCloser to its original state
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		if bodyBytes == nil {
			Logger.Sugar().Infof("Request URL: %v, Method: %v", ctx.Request.URL, ctx.Request.Method)
		} else {
			Logger.Sugar().Infof("Request URL: %v, Method: %v, Body: %v", ctx.Request.URL, ctx.Request.Method, string(bodyBytes))
		}

		ctx.Next()
	}
}
