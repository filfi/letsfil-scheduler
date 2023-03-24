package glif

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/v1api"
	"net/http"
)

// var LotusNodeAddr = "https://api.node.glif.io/rpc/v0"
//var LotusNodeAddr = "https://api.hyperspace.node.glif.io/rpc/v1"

var Token = configs.Get().Lotus.AuthToken
var LotusNodeAddr = configs.Get().Lotus.Url

var lapi api.FullNodeStruct
var Lv1api v1api.FullNode
var ctx = context.Background()
var httpsClient *ethclient.Client

func init() {
	//headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	headers := http.Header{
		"content-type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + Token},
	}
	closer, err := jsonrpc.NewMergeClient(context.Background(), LotusNodeAddr, "Filecoin", []interface{}{&lapi.Internal, &lapi.CommonStruct.Internal}, headers)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}
	defer closer()
	//httpsClient, err = ethclient.Dial("https://api.hyperspace.node.glif.io")
	httpsClient, err = ethclient.Dial(configs.Get().FilFevm.HttpsUrl)
	if err != nil {
		fmt.Errorf("connecting with lotus failed: %s", err)
	}

}
