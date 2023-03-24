package rpcUtils

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"github.com/filecoin-project/go-address"
)

func CreateAddress(addr string) address.Address {
	adr, _ := address.NewFromString(addr)
	return adr
}

func GetCategoryByMethod(method string) int {
	switch method {
	case constant.RpcMethodSend:
		return constant.RpcMethodSendCategory
	case constant.RpcMethodWithdraw:
		return constant.RpcMethodWithdrawCategory
	default:
		return -1
	}
}
