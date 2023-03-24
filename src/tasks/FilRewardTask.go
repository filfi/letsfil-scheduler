package tasks

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/service"
)

type FilRewardTask struct {
	FilRewardService *service.FilRewardService `inject:"-"`
}

func NewFilRewardTask() *FilRewardTask {
	return &FilRewardTask{}
}

func (this *FilRewardTask) SyncFilBlockReward() {
	this.FilRewardService.SyncFilBlockReward()
}

func (this *FilRewardTask) SyncFilBlockRewardDistribution() {
	this.FilRewardService.SyncFilBlockRewardDistribution()
}

func (this *FilRewardTask) CheckSyncFilBlockReward() {
	this.FilRewardService.CheckFilBlockRewardResult()
}

func (this *FilRewardTask) Build(goft *goft.Goft) {
	goft.Handle("GET", "/name", this.Name)
}

func (this *FilRewardTask) Name() string {
	return "FilRewardTask"
}
