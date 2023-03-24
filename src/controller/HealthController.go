package controller

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/gin-gonic/gin"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

type HealthController struct {
	MarketClient *client.MarketClient `inject:"-"`
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (this *HealthController) Health(c *gin.Context) goft.Json {
	return result.OK
}

func (this *HealthController) Name() string {
	return "HealthController"
}

func (this *HealthController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/check", this.Health)
}
