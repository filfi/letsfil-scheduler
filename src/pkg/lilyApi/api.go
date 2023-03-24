package lilyApi

import (
	"encoding/json"
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/httpUtils"
)

func GetFilBlockReward(miner string, height uint64) ([]GetBlockRewardResponse, error) {
	var header = make(map[string]string, 0)
	var param = make(map[string]string, 0)
	header["Content-Type"] = "application/json"

	param["miner"] = miner
	param["epoch"] = fmt.Sprintf("%d", height)

	url := configs.Get().FilDate.Url + "/GetBlockReward?"
	retStr := httpUtils.Get(url, param, header)
	var retList []GetBlockRewardResponse
	err := json.Unmarshal([]byte(retStr), &retList)
	if err != nil {
		return nil, err
	}
	return retList, nil
}
