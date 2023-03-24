package middleware

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HttpsFILClient struct {
}

func NewHttpsFILClient() *HttpsFILClient {
	return &HttpsFILClient{}
}

func (this *HttpsFILClient) HttpsFILClient() *ethclient.Client {
	httpsClient, err := ethclient.Dial(configs.Get().FilFevm.HttpsUrl)
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	return httpsClient
}

type WsFILClient struct {
}

func NewWsFILClient() *WsFILClient {
	return &WsFILClient{}
}

func (this *WsFILClient) WsFILClient() *ethclient.Client {
	httpsClient, err := ethclient.Dial(configs.Get().FilFevm.WsUrl)
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	return httpsClient
}
