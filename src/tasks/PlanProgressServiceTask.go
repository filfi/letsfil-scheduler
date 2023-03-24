package tasks

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/service"
)

type PlanProgressTask struct {
	PlanProgressService *service.PlanProgressService `inject:"-"`
}

func NewPlanProgressTask() *PlanProgressTask {
	return &PlanProgressTask{}
}

func (this *PlanProgressTask) SyncMinerSealedProgress() {
	this.PlanProgressService.SyncMinerSealedProgress()
}
func (this *PlanProgressTask) SyncMinerSealedStart() {
	this.PlanProgressService.SyncMinerSealedStart()
}

func (this *PlanProgressTask) Build(goft *goft.Goft) {
	goft.Handle("GET", "/name", this.Name)
}

func (this *PlanProgressTask) Name() string {
	return "PlanProgressTask"
}
