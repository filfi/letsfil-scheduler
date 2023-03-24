package goft

import "github.com/gin-gonic/gin"

// Fairing 中间件接口
type Fairing interface {
	OnRequest(*gin.Context) error
	OnResponse(result interface{}) (interface{}, error)
}
