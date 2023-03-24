package influxUtils

import (
	"bytes"
	"context"
	"fmt"
	//"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models/entity"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/commonUtils"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"go.uber.org/zap"
	"strings"
	"time"
)

func JoinOrBySlice(key string, arr []string) string {
	return fmt.Sprint(`r["`+key+`"]=="`, strings.Join(arr, `" or r["`+key+`"]=="`), `"`)
}

func JoinMachineOrBySlice(key string, arr []*entity.NodeMachine) string {
	var buffer bytes.Buffer
	for k, v := range arr {
		if k == 0 {
			buffer.WriteString(fmt.Sprint(`r["`+key+`"]=="`, v.Tag))
		} else {
			buffer.WriteString(`" or r["`)
			buffer.WriteString(key)
			buffer.WriteString(`"]=="`)
			buffer.WriteString(v.Tag)
			buffer.WriteString(`"`)
		}
	}
	return buffer.String()
}

func GetStringValue(str interface{}, defaultValue string) string {
	if str == nil {
		return defaultValue
	} else {
		return str.(string)
	}
}

func GetFloat64(str interface{}, defaultValue float64) float64 {
	if str == nil {
		return defaultValue
	} else {
		return str.(float64)
	}
}

func GetInt64(str interface{}, defaultValue int64) int64 {
	if str == nil {
		return defaultValue
	} else {
		return str.(int64)
	}
}

func Float2String(str interface{}, defaultValue string) string {
	if str == nil {
		return defaultValue
	} else {
		return fmt.Sprint(str.(float64))
	}
}

func GetTimeValue(str interface{}) time.Time {
	if str == nil {
		zap.L().Error("influxdb get time value error, time is nil")
		panic(result.SystemError)
	} else {
		return str.(time.Time)
	}
}

func GetQueryValue(influxResp *api.QueryTableResult, err error, method string) string {
	var resp string
	if err == nil {
		for influxResp.Next() {
			values := influxResp.Record().Values()
			value := values["_value"]
			if value != nil {
				resp = commonUtils.CastToString(value)
			}
		}
		// Check for an error
		if influxResp.Err() != nil {
			zap.L().Error(method + " query parsing error: " + influxResp.Err().Error())
			panic(result.SystemError)
		}
	} else {
		zap.L().Error(method + " influxdb query error: " + err.Error())
		panic(result.SystemError)
	}
	return resp
}

func Query(queryApi api.QueryAPI, sql string, method string) (*api.QueryTableResult, error) {
	influxResp, err := queryApi.Query(context.Background(), sql)
	if err == nil {
		// Check for an error
		if influxResp.Err() != nil {
			zap.L().Error(method + " query parsing error: " + influxResp.Err().Error())
			return nil, influxResp.Err()
		}
	} else {
		zap.L().Error(method + " influxdb query error: " + err.Error())
		return nil, influxResp.Err()
	}
	return influxResp, nil

}
