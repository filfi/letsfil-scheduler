package validtorUtils

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateUsername ValidateMobile 校验用户名
func ValidateUsername(fl validator.FieldLevel) bool {
	// 利用反射拿到结构体tag含有mobile的key字段
	mobile := fl.Field().String()
	//使用正则表达式判断是否合法
	ok, _ := regexp.MatchString(`^1\d{10}$`, mobile)
	if !ok {
		return false
	}
	return true
}

func ValidateMinerId(ctx *gin.Context) string {
	minerId := ctx.Query("minerId")
	if minerId == "" {
		panic(result.ParamError)
	} else {
		return minerId
	}
}
