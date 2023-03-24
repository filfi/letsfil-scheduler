package models

import "gorm.io/gorm"

type RewardSync struct {
	gorm.Model
	Tx          string `gorm:"column:tx;type:text" json:"tx"`                     // Transaction ID
	Pyload      string `gorm:"column:pyload;type:text" json:"pyload"`             // Payload
	FromAddress string `gorm:"column:from_address;type:text" json:"from_address"` // from
	ToAddress   string `gorm:"column:to_address;type:text" json:"to_address"`     // to
	Status      int64  `gorm:"column:status;type:int" json:"status"`              // Status
}
