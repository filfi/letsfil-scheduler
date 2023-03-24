package middleware

import (
	"github.com/huobirdcenter/huobi_golang/config"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

type HuobiClient struct {
}

func NewHuobiClient() *HuobiClient {
	return &HuobiClient{}
}

func (this *HuobiClient) HuobiClient() *client.MarketClient {
	hbClient := new(client.MarketClient).Init(config.Host)
	//_, err := hbClient.GetHistoricalTrade("filusdt", "1min", "1")
	//if err != nil {
	//	zap.L().Error("GetTimestamp error", zap.Error(err))
	//} else {
	//	zap.L().Info("huobi connect success")
	//}
	return hbClient
}
