package contract

import (
	"fmt"
	journal "git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// HeroInfo is an auto generated Go binding around an Ethereum contract.
type RaiseInfo struct {
}

func CreateLockABI(client *ethclient.Client, conAddress string) (*Lock, error) {
	ctrt, err := NewLock(common.HexToAddress(conAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate ERC721 contract: %v", err)
		return nil, err
	}

	return ctrt, nil
}

func CreateRaiseFactory(client *ethclient.Client, conAddress string) (*LetsFilRaiseFactory, error) {
	ctrt, err := NewLetsFilRaiseFactory(common.HexToAddress(conAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate ERC721 contract: %v", err)
		return nil, err
	}

	return ctrt, nil
}

func CreateLetsFilRaisePool(client *ethclient.Client, conAddress string) (*LetsFilRaisePool, error) {
	ctrt, err := NewLetsFilRaisePool(common.HexToAddress(conAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate ERC721 contract: %v", err)
		return nil, err
	}

	return ctrt, nil
}

func (pool *Lock) UnpackLog(out interface{}, event string, log types.Log) error {
	abiP, err := LockMetaData.GetAbi()
	if err != nil {
		journal.Info("UnpackLog err", err)
		return err
	}
	poolabi := *abiP
	if log.Topics[0] != poolabi.Events[event].ID {
		return fmt.Errorf("event UnpackLog signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := poolabi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolabi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func (pool *LetsFilRaiseFactory) UnpackLog(out interface{}, event string, log types.Log) error {
	abiP, err := LetsFilRaiseFactoryMetaData.GetAbi()
	if err != nil {
		journal.Info("UnpackLog err", err)
		return err
	}
	poolabi := *abiP
	if log.Topics[0] != poolabi.Events[event].ID {
		return fmt.Errorf("event UnpackLog signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := poolabi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolabi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func (pool *LetsFilRaisePool) UnpackLog(out interface{}, event string, log types.Log) error {
	abiP, err := LetsFilRaisePoolMetaData.GetAbi()
	if err != nil {
		journal.Info("UnpackLog err", err)
		return err
	}
	poolabi := *abiP
	if log.Topics[0] != poolabi.Events[event].ID {
		return fmt.Errorf("event UnpackLog signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := poolabi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolabi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
