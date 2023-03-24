package glif

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"strconv"
	"testing"
)

func TestGetMinerPower(t *testing.T) {
	minerId := "t01001"
	ret, err := GetMinerPower(minerId)
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
	t.Log(ret.MinerPower.QualityAdjPower.String())
	t.Log(ret.TotalPower.QualityAdjPower.String())
}

func TestGetGetFilterLogs(t *testing.T) {
	//GetMinerSectors("f01000", 0, 1000)
	// 获取最新的区块高度
	latestBlockNumber, err := lapi.ChainHead(context.Background())
	if err != nil {
		t.Fatalf("Blockchain - Failed get latest block number: %s", err.Error())
	}
	fmt.Printf("latestBlockNumber: %v\n", latestBlockNumber.Height().String())

	startBigInt := (latestBlockNumber.Height() - abi.ChainEpoch(2800))
	startBigStr := strconv.FormatUint(uint64(startBigInt), 16)
	endBigInt := (latestBlockNumber.Height())
	endBigStr := strconv.FormatUint(uint64(endBigInt), 16)

	contractAddresses := make(ethtypes.EthAddressList, 0)

	fmt.Printf("startBigInt: %v\n", startBigStr)
	fmt.Printf("endBigInt: %v\n", endBigStr)
	fmt.Printf("contractAddresses: %v\n", contractAddresses)

	// 获取当前募集计划的合约地址
	//contractAddress := common.HexToAddress("0x25BD83cd62802575f5a7F7D7F9a39b1CE247B035")
	//contractAddress := common.HexToAddress("0xa588B90FC814129FE10dFB93169246EE0064F27b")
	// 0x4247b30a4795c3F215274d2eE5666830E8f4f548  hyper 2k
	var contractAddress ethtypes.EthAddress

	contractAddress, err = ethtypes.ParseEthAddress("0x4247b30a4795c3F215274d2eE5666830E8f4f548")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("contractAddress: %v\n", contractAddress)

	contractAddresses = append(contractAddresses, contractAddress)
	//
	sign_str := "PutValue(uint256)"
	sign := fevm.EthTopicHash(sign_str)

	fmt.Printf("sign: %v\n", sign)

	query := ethtypes.EthFilterSpec{
		FromBlock: &startBigStr,
		ToBlock:   &endBigStr,
		Address:   contractAddresses,
		//Topics:    AllTopics(), // [][]common.Hash
		Topics: ethtypes.EthTopicSpec{
			{sign},
		},
	}
	f4Address, err := contractAddress.ToFilecoinAddress()
	log.Infof("f4Address: %s", f4Address.String())

	//log.Infof("topic: %+v", topic)
	//
	//	//
	//logs, err := httpsClient.FilterLogs(context.Background(), query)
	ret, err := GetGetFilterLogs(query)
	if err != nil {
		t.Error(err)

	}
	logs := ret.Results

	for i := 0; i < len(logs); i++ {
		vLog := logs[i]
		var oneLog *ethtypes.EthLog
		oneLog, err := fevm.ParseEthLogs(vLog)
		if err != nil {
			log.Error("vLog.(ethtypes.EthLog) error")
			continue
		}
		log.Infof("oneLog: %+v", oneLog)
		log.Infof("oneLog.Data.: %s", oneLog.Data.String())
	}

	log.Infof("sectors %+v", ret)
}
