package commonUtils

import (
	"crypto/md5"
	"encoding/hex"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/redisKey"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/bigInt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"strconv"
)

// Md5 md5加密
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func GetPageInfo(ctx *gin.Context) (pageNum, pageSize int) {
	pn := ctx.DefaultQuery(constant.PageNum, constant.DefaultPageNum)
	ps := ctx.DefaultQuery(constant.PageSize, constant.DefaultPageSize)
	var err error
	pageNum, err = strconv.Atoi(pn)
	if err != nil {
		panic(result.ParamError)
	}
	pageSize, err = strconv.Atoi(ps)
	if err != nil {
		panic(result.ParamError)
	}
	pageNum = (pageNum - 1) * pageSize
	return
}

func GetUserId(ctx *gin.Context) int64 {
	uid, exists := ctx.Get(redisKey.UserTokenValue)
	if !exists {
		panic(result.TokenInvalid)
	}
	return uid.(int64)
}

func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func String2Decimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		zap.L().Error("string to decimal error", zap.Error(err))
		panic(result.ParamError)
	}
	return d
}

func String2Int64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		zap.L().Error("string to int64 error", zap.Error(err))
		panic(result.ParamError)
	}
	return i
}

func CastToString(param interface{}, prec ...int) string {
	switch i := param.(type) {
	case nil:
		return ""
	case string:
		return i
	case bigInt.BigInt:
		return i.String()
	case float32:
		return strconv.FormatFloat(float64(i), 'f', prec[0], 32)
	case float64:
		return strconv.FormatFloat(i, 'f', prec[0], 64)
	case int:
		return strconv.Itoa(i)
	case int8:
		return strconv.Itoa(int(i))
	case int16:
		return strconv.Itoa(int(i))
	case int32:
		return strconv.Itoa(int(i))
	case int64:
		return bigInt.NewBigInt(i).String()
	default:
		panic("CastToString type not support")
	}
}
