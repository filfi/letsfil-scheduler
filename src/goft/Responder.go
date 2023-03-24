package goft

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var responderList []Responder
var once_resp_list sync.Once

func get_responder_list() []Responder {
	once_resp_list.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
			(ViewResponder)(nil),
			(VoidResponder)(nil),
		}
	})
	return responderList
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, getFairingHandler().handlerFairing(this, context).(string))
	}
}

type Json interface{}
type JsonResponder func(*gin.Context) Json

func (this JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, getFairingHandler().handlerFairing(this, context))
	}
}

type Void struct{}
type VoidResponder func(ctx *gin.Context) Void

func (this VoidResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getFairingHandler().handlerFairing(this, context)
	}
}

// Deprecated: 暂时不提供View的解析
type View string
type ViewResponder func(*gin.Context) View

func (this ViewResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(200, string(this(context))+".html", context.Keys)
	}
}
