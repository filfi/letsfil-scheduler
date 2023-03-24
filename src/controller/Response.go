package controller

type RaiseAssetReport struct {
	//已产出
	TotalAsset       string `json:"total_asset"`        //已产出
	Reward24h1T      string `json:"reward_24h_1t"`      //24小时产出
	TotalAlloc       string `json:"total_alloc"`        //已分配
	MinerActive      int64  `json:"miner_active"`       //活跃矿工
	VestingFund      string `json:"vesting_fund"`       //锁仓基金
	PledgeFund       string `json:"pledge_fund"`        //质押基金
	UsedGas          string `json:"used_gas"`           //已使用gas
	Pelnaty          string `json:"pelnaty"`            //罚金
	MinerTargetCapac string `json:"miner_target_capac"` //矿工目标容量
	MinerCapacSize   string `json:"miner_capac_size"`   //矿工容量
	InvestorCount    int64  `json:"investor_count"`     //投资者数量
}

type InvestorAssetReport struct {
	//已产出
	TotalAsset  int64 `json:"total_asset"`   //已产出
	Reward24h1T int64 `json:"reward_24h_1t"` //24小时产出
	TotalAlloc  int64 `json:"total_alloc"`   //已分配
	MinerActive int64 `json:"miner_active"`  //活跃矿工
	VestingFund int64 `json:"vesting_fund"`  //锁仓基金
	PledgeFund  int64 `json:"pledge_fund"`   //质押基金
	UsedGas     int64 `json:"used_gas"`      //已使用gas
	Pelnaty     int64 `json:"pelnaty"`       //罚金

	MinerCapacSize string `json:"miner_capac_size"` //矿工容量
	InvestorCount  int64  `json:"investor_count"`   //投资者数量
}
