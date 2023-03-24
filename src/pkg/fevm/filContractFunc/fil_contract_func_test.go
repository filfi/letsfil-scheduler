package filContractFunc

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestFilContractFunc(t *testing.T) {

	httpsClient, err := ethclient.Dial("xxxxxx")
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	log.Info("httpsClient: ", httpsClient)
	ret, err := PushBlockReward(httpsClient, "0x4247b30a4795c3F215274d2eE5666830E8f4f548")
	if err != nil {
		log.Fatalf("Blockchain - Failed dial http: %s", err.Error())
	}
	log.Info("ret: ", ret)
}
