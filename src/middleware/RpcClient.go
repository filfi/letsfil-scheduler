package middleware

import (
	"context"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"log"
	"net/http"
)

type RpcClient struct {
}

func NewRpcClient() *RpcClient {
	return &RpcClient{}
}

func (this *RpcClient) InitRpcClient() *api.FullNodeStruct {
	headers := http.Header{constant.RpcHeaderKey: []string{constant.RpcHeaderTokenPrefix + configs.GetValue(configs.LotusAuthToken, configs.Get().Lotus.AuthToken)}}
	var api api.FullNodeStruct
	_, err := jsonrpc.NewMergeClient(context.Background(), configs.GetValue(configs.LotusUrl, configs.Get().Lotus.Url), configs.Get().Lotus.NameSpace,
		[]interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}

	//todo...
	//name, err := api.Version(context.Background())
	//if err != nil {
	//	log.Fatalf("connecting with lotus failed: %s", err)
	//}
	//fmt.Println("当前rpc版本===========", name)
	return &api
}
