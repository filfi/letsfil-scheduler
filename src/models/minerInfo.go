package models

import "gorm.io/gorm"

type MinerInfo struct {
	gorm.Model
	MinerId        string `gorm:"column:miner_id;type:text" json:"miner_id"`               // Miner ID
	OwnerId        string `gorm:"column:owner_id;type:text" json:"owner_id"`               // Owner ID
	ManagerAddress string `gorm:"column:manager_address;type:text" json:"manager_address"` // Manager address
	Height         uint64  `gorm:"column:height;type:bigint;primaryKey" json:"height"`      // Epoch this rewards summary applies to.
}
