package filContractFunc

import (
	"context"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/contract"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func PushBlockReward(cli *ethclient.Client, conAddress string) (receipt *types.Receipt, err error) {

	// 创建指向智能合约地址的实例
	contract, err := contract.CreateLockABI(cli, conAddress)
	if err != nil {
		return nil, err
	}

	// 创建交易
	privateKey, err := ethcrypto.HexToECDSA(configs.GetValue(configs.Private_Key, ""))
	if err != nil {
		panic(err)
	}
	nonce, err := cli.PendingNonceAt(context.Background(), ethcrypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		panic(err)
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	log.Infof("gasPrice: %s", gasPrice.String())

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))
	auth.GasTipCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))

	// 调用 PutValue 方法
	tx, err := contract.PutValue(auth, big.NewInt(123))
	if err != nil {
		log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
		return nil, err

	}
	log.Infof("tx: %+v", tx)
	log.Infof("tx.Hash(): %s", tx.Hash().Hex())

	timer := time.NewTicker(30 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("Running task...")
			// 执行任务的代码，例如向 API 发送请求并处理响应
			// 如果收到结果则退出循环
			// 等待交易被打包并确认
			receipt, err = cli.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
				continue
				//return nil, err
			}
			fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
			fmt.Printf("Gas used: %d\n", receipt.GasUsed)
			fmt.Printf("Block number: %d\n", receipt.BlockNumber)
			fmt.Printf("Block hash: %s\n", receipt.BlockHash.Hex())
			fmt.Printf("Contract address: %s\n", receipt.ContractAddress.Hex())
			fmt.Printf("Status: %d\n", receipt.Status)
			fmt.Printf("Logs: %+v\n", receipt.Logs)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].Topics)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].Data)
			fmt.Printf("Logs: %+v\n", receipt.Logs[0].BlockNumber)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].BlockHash.Hex())
			return nil, nil
		}
	}

	return receipt, nil
}

func PostBlockChainData(cli *ethclient.Client, conAddress string, methmod string, payload interface{}) (receipt *types.Receipt, err error) {

	// 创建交易
	privateKey, err := ethcrypto.HexToECDSA(configs.GetValue(configs.Private_Key, ""))
	if err != nil {
		panic(err)
	}
	nonce, err := cli.PendingNonceAt(context.Background(), ethcrypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		panic(err)
	}
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	log.Infof("gasPrice: %s", gasPrice.String())

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))
	auth.GasTipCap = big.NewInt(0).Add(gasPrice, big.NewInt(300000))

	var tx *types.Transaction
	switch methmod {
	case "PutValue":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLockABI(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		tx, err := contract.PutValue(auth, big.NewInt(123))
		if err != nil {
			log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
			return nil, err

		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "StartSeal":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		tx, err := contract.StartRaisePlan(auth, nil)
		if err != nil {
			log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
			return nil, err

		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "EndSeal":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		tx, err := contract.StopSeal(auth)
		if err != nil {
			log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
			return nil, err

		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	case "PushSealProgress":
		// 创建指向智能合约地址的实例
		contract, err := contract.CreateLetsFilRaisePool(cli, conAddress)
		if err != nil {
			return nil, err
		}
		// 调用 PutValue 方法
		tx, err := contract.PushSealProgress(auth, nil)
		if err != nil {
			log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
			return nil, err

		}
		log.Infof("tx: %+v", tx)
		log.Infof("tx.Hash(): %s", tx.Hash().Hex())
	}

	timer := time.NewTicker(30 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("Running task...")
			// 执行任务的代码，例如向 API 发送请求并处理响应
			// 如果收到结果则退出循环
			// 等待交易被打包并确认
			receipt, err = cli.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Errorf("Blockchain - Failed PutValue: %s", err.Error())
				continue
				//return nil, err
			}
			fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
			fmt.Printf("Gas used: %d\n", receipt.GasUsed)
			fmt.Printf("Block number: %d\n", receipt.BlockNumber)
			fmt.Printf("Block hash: %s\n", receipt.BlockHash.Hex())
			fmt.Printf("Contract address: %s\n", receipt.ContractAddress.Hex())
			fmt.Printf("Status: %d\n", receipt.Status)
			fmt.Printf("Logs: %+v\n", receipt.Logs)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].Topics)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].Data)
			fmt.Printf("Logs: %+v\n", receipt.Logs[0].BlockNumber)
			fmt.Printf("Logs: %s\n", receipt.Logs[0].BlockHash.Hex())
			return nil, nil
		}
	}

	return receipt, nil
}
