package models

import "gorm.io/gorm"

type DepositDetails struct {
	gorm.Model
	TxHash          string `gorm:"column:tx_hash;type:text" json:"tx_hash"`                   // 交易哈希
	DepositType     uint   `gorm:"column:deposit_type;type:text" json:"deposit_type"`         // 保证金类型
	DepositStatus   uint   `gorm:"column:deposit_status;type:int" json:"deposit_status"`      // 保证金状态
	FromAddress     string `gorm:"column:from_address;type:text" json:"from_address"`         // from
	ToAddress       string `gorm:"column:to_address;type:text" json:"to_address"`             // to
	Value           string `gorm:"column:value;type:text" json:"value"`                       // value
	TransactionTime string `gorm:"column:transaction_time;type:text" json:"transaction_time"` // transaction_time
	RaiserId        uint64 `gorm:"column:raiser_id;type:bigint" json:"raiser_id"`             // 募集计划id
}
