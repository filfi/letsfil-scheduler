package models

import "gorm.io/gorm"

type BlockReward struct {
	gorm.Model
	MinerId          string `gorm:"column:miner_id;type:text" json:"miner_id"`                       // Miner ID
	Height           uint64 `gorm:"column:height;type:bigint" json:"height"`                         // Block height
	TotalReward      string `gorm:"column:reward;type:text" json:"reward"`                           // Reward
	ReleaseReward    string `gorm:"column:release_reward;type:text" json:"release_reward"`           // 释放奖励
	NextReleasHeight uint64 `gorm:"column:next_releas_height;type:bigint" json:"next_releas_height"` // 下次释放高度
	ReleaseCount     uint64 `gorm:"column:release_count;type:bigint" json:"release_count"`           // 释放次数
}
