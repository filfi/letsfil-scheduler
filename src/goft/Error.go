package goft

import (
	"bytes"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime"
)

const (
	ErrorStatus = "ErrorStatus"
)

func panicTrace(kb int) string {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return string(stack)
}

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				zap.L().Error(panicTrace(10))
				status := 200 //default status==400
				if value, exists := context.Get(ErrorStatus); exists {
					if v, ok := value.(int); ok {
						status = v
					}
				}

				switch t := e.(type) {
				case string:
					zap.L().Error(t)
					context.AbortWithStatusJSON(status, gin.H{"error": t})
				case result.Resp:
					zap.L().Error(t.Msg)
					context.JSON(http.StatusOK, t)
				default:
					if pe, ok := e.(error); ok {
						zap.L().Error(pe.Error())
						context.AbortWithStatusJSON(status, pe.Error())
					} else {
						context.AbortWithStatusJSON(status, e)
					}
				}

			}
		}()
		context.Next()
	}
}

func Throw(err interface{}, code int, context *gin.Context) {
	context.Set(ErrorStatus, code)
	panic(err)
}

func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
