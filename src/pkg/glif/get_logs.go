package glif

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

func GetGetFilterLogs(filter ethtypes.EthFilterSpec) (mbs *ethtypes.EthFilterResult, err error) {

	log.Infof("GetGetFilterLogs start %+v", filter)
	filt, err := lapi.EthNewFilter(context.TODO(), &filter)
	if err != nil {
		log.Infof("EthNewFilter failed: %s", err)
		return nil, err
	}

	log.Infof("filt: %+v\n", filt)
	ma, err := lapi.EthGetFilterLogs(context.Background(), filt)
	if err != nil {
		log.Infof("GetGetFilterLogs failed: %s", err)
		//fmt.Println(err)
		return nil, err
	}
	fmt.Println(ma)

	return ma, nil

}

func GetGetFilterLogsEth(filter ethereum.FilterQuery) (mbs *[]types.Log, err error) {

	if err != nil {
		log.Error("ethclient.Dial ", err)
		return
	}

	logs, err := httpsClient.FilterLogs(context.Background(), filter)
	if err != nil {
		log.Error("SubscribeFilterLogs ", err)
		return
	}
	fmt.Printf("logs: %+v\n", logs)

	return &logs, nil

}
