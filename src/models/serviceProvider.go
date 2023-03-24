package models

import "gorm.io/gorm"

type ServiceProvider struct {
	gorm.Model
	FullName string `gorm:"column:full_name;type:text" json:"full_name"` // full_name
	// 简称
	ShortName string `gorm:"column:short_name;type:text" json:"short_name"` // short_name
	// 简介
	Introduction string `gorm:"column:introduction;type:text" json:"introduction"` // introduction
	// logo
	LogoUrl string `gorm:"column:logo_url;type:text" json:"logo_url"` // logo
	// 钱包地址
	WalletAddress string `gorm:"column:wallet_address;type:text" json:"wallet_address"` // wallet_address

}
