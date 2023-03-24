package models

import "gorm.io/gorm"

type FundDetails struct {
	gorm.Model
	TxHash      string `gorm:"column:tx_hash;type:text" json:"tx_hash"`           // 交易哈希
	TxType      uint   `gorm:"column:tx_type;type:text" json:"tx_type"`           // 交易类型 1 质押 2 赎回 3 提取收益
	FromAddress string `gorm:"column:from_address;type:text" json:"from_address"` // from
	ToAddress   string `gorm:"column:to_address;type:text" json:"to_address"`     // to
	Value       string `gorm:"column:value;type:text" json:"value"`               // 交易金额
	RaiserId    uint64 `gorm:"column:raiser_id;type:bigint" json:"raiser_id"`     // 募集计划id

}
