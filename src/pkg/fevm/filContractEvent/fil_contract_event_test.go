package filContractEvent

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/contract"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestFilClient(t *testing.T) {
	httpsClient, err := ethclient.Dial("https://api.hyperspace.node.glif.io")
	if err != nil {
		t.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}

	// 获取最新的区块高度
	latestBlockNumber, err := httpsClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("Blockchain - Failed get latest block number: %s", err.Error())
	}
	fmt.Printf("latestBlockNumber: %v\n", latestBlockNumber.Number.Uint64())

	startBigInt := big.NewInt(0)
	startBigInt = startBigInt.SetUint64(latestBlockNumber.Number.Uint64() - 2800)

	endBigInt := big.NewInt(0)
	endBigInt = endBigInt.SetUint64(latestBlockNumber.Number.Uint64())

	contractAddresses := make([]common.Address, 0)

	// 获取当前募集计划的合约地址
	//contractAddress := common.HexToAddress("0x25BD83cd62802575f5a7F7D7F9a39b1CE247B035")
	contractAddress := common.HexToAddress("0x4247b30a4795c3F215274d2eE5666830E8f4f548")

	contractAddresses = append(contractAddresses, contractAddress)
	//
	sign_str := []byte("PutValue(uint256)")
	sign := ethcrypto.Keccak256Hash(sign_str)
	log.Infof("sign: %s", sign.String())

	query := ethereum.FilterQuery{
		FromBlock: startBigInt,
		ToBlock:   endBigInt,
		Addresses: contractAddresses,
		//Topics:    AllTopics(),             // [][]common.Hash
		Topics: [][]common.Hash{{sign}}, // [][]common.Hash
	}

	//
	logs, err := httpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Println("SubscribeFilterLogsfmt", err)
		t.Error("SubscribeFilterLogs ", err)
	}
	//logs, err := glif.GetGetFilterLogs(query)
	for i := 0; i < len(logs); i++ {
		vLog := logs[i]
		fmt.Printf("vLog: %+v\n", vLog)
		rewardPool, err := contract.CreateLockABI(httpsClient, vLog.Address.String())
		if err != nil {
			fmt.Println("CreateLockABI", err)
			t.Error("CreateLockABI ", err)
		}
		event := new(contract.LockPutValue)
		rewardPool.UnpackLog(event, "PutValue", vLog)
		fmt.Printf("event: %+v\n", event)

	}

}

func TestFilterEventEth(t *testing.T) {
	httpsClient, err := ethclient.Dial("https://api.hyperspace.node.glif.io")
	if err != nil {
		t.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}

	// 获取最新的区块高度
	latestBlockNumber, err := httpsClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("Blockchain - Failed get latest block number: %s", err.Error())
	}
	fmt.Printf("latestBlockNumber: %v\n", latestBlockNumber.Number.Uint64())

	startBigInt := big.NewInt(0)
	startBigInt = startBigInt.SetUint64(latestBlockNumber.Number.Uint64() - 2800)

	endBigInt := big.NewInt(0)
	endBigInt = endBigInt.SetUint64(latestBlockNumber.Number.Uint64())

	contractAddresses := make([]string, 0)

	// 获取当前募集计划的合约地址
	//contractAddress := common.HexToAddress("0x25BD83cd62802575f5a7F7D7F9a39b1CE247B035")
	contractAddress := "0x4247b30a4795c3F215274d2eE5666830E8f4f548"

	contractAddresses = append(contractAddresses, contractAddress)

	logs, err := FilterEventEth(startBigInt.Uint64(), endBigInt.Uint64(), contractAddresses)
	if err != nil {
		fmt.Println("SubscribeFilterLogsfmt", err)
		t.Error("SubscribeFilterLogs ", err)
	}
	//logs, err := glif.GetGetFilterLogs(query)
	for i := 0; i < len(*logs); i++ {
		vLog := (*logs)[i]
		fmt.Printf("vLog: %+v\n", vLog)
		rewardPool, err := contract.CreateLockABI(httpsClient, vLog.Address.String())
		if err != nil {
			fmt.Println("CreateLockABI", err)
			t.Error("CreateLockABI ", err)
		}
		event := new(contract.LockPutValue)
		rewardPool.UnpackLog(event, "PutValue", vLog)
		fmt.Printf("event: %+v\n", event)

	}

}
