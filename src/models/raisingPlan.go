package models

import (
	"gorm.io/gorm"
)

type RaisingPlan struct {
	gorm.Model
	TxHash           string  `gorm:"column:tx_hash;type:text" json:"tx_hash"`                   // 交易哈希
	RaisingId        uint64  `gorm:"column:raising_id;type:bigint" json:"raising_id"`           // 募集计划id
	Raiser           string  `gorm:"column:raiser;type:text" json:"raiser"`                     // 募集人 发起人地址
	SponsorCompany   string  `gorm:"column:sponsor_company;type:text" json:"sponsor_company"`   // 发起单位
	SecurityFundRate string  `gorm:"column:security_fund;type:text" json:"security_fund"`       // 保证金占比
	MinerId          string  `gorm:"column:miner_id;type:text" json:"miner_id"`                 // miner
	RaiseAddress     string  `gorm:"column:raise_address;type:text" json:"raise_address"`       // 募集计划合约地址
	TargetPower      string  `gorm:"column:target_power;type:text" json:"target_power"`         // 目标算力
	TargetAmount     string  `gorm:"column:target_amount;type:text" json:"target_amount"`       // 目标金额
	ClosingTime      uint64  `gorm:"column:closing_time;type:bigint" json:"closing_time"`       //募集截止时间
	RaiserShare      uint64  `gorm:"column:raiser_share;type:bigint" json:"raiser_share"`       // 募集者权益
	InvestorShare    uint64  `gorm:"column:investor_share;type:bigint" json:"investor_share"`   // 投资者权益
	ServicerShare    uint64  `gorm:"column:servicer_share;type:bigint" json:"servicer_share"`   // 服务商权益
	SectorSize       uint64  `gorm:"column:sector_size;type:bigint" json:"sector_size"`         // 扇区大小
	SealTimeLimit    uint64  `gorm:"column:seal_time_limit;type:bigint" json:"seal_time_limit"` // 封装时间限制
	Status           uint    `gorm:"column:status;type:int" json:"status"`                      // 状态 0:募集中 1:募集成功 2:募集失败 3:等待服务商签名 9:节点结束
	Progress         uint    `gorm:"column:progress;type:int" json:"progress"`                  // 进度 0:未开始 1:进行中 2:已完成
	IncomeRate       float64 `gorm:"column:income_rate;type:int" json:"income_rate"`            // 收益率
	SealedStatus     uint64  `gorm:"column:sealed_status;type:bigint" json:"sealed_status"`     // 封装开始 0:未开始 1:进行中 2:已完成

	LatestBlockNumber uint64 `gorm:"column:latest_block_number;type:bigint" json:"latest_block_number"` // 最新区块高度
}
