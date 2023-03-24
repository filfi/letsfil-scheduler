package controller

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"gorm.io/gorm"
)

type AssetController struct {
	Db *gorm.DB `inject:"-"`
}

func NewAssetController() *AssetController {
	return &AssetController{}
}

func (this *AssetController) Name() string {
	return "AssetController"
}

func (this *AssetController) Build(goft *goft.Goft) {
	//goft.Handle("GET", "/list", this.PlanAllList)

}
