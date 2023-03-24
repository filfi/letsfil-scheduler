package lilyApi

import "github.com/filecoin-project/go-state-types/big"

type GetBlockRewardResponse struct {
	MinerID string  `json:"miner_id"`
	Height  int64   `json:"height"`
	Reward  big.Int `json:"reward"`
}
