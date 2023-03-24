package controller

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/gin-gonic/gin"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"gorm.io/gorm"
)

type SProvierController struct {
	MarketClient *client.MarketClient `inject:"-"`
	Db           *gorm.DB             `inject:"-"`
}

func NewSProvierController() *SProvierController {
	return &SProvierController{}
}

func (this *SProvierController) spList(c *gin.Context) goft.Json {
	// 查询所有的服务商、
	var spList []models.ServiceProvider
	if err := this.Db.Find(&spList).Error; err != nil {
		return result.SpNotFound
	}
	return result.OK.WithData(spList)
}

func (this *SProvierController) Name() string {
	return "SProvierController"
}

func (this *SProvierController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/list", this.spList)
}
