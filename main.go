package main

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/controller"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	. "git.sxxfuture.net/filfi/letsfil/letsfil-job/src/middleware"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/tasks"
)

func main() {
	filRewardTask := tasks.NewFilRewardTask()
	planProgressTask := tasks.NewPlanProgressTask()
	//middleware.InitRpcClient()
	goft.Ignite().
		Config(NewDbConfig(),
			//NewRedisClient(),
			NewServiceConfig(),
			NewHttpsFILClient(),
			NewWsFILClient()).
		Mount("health", controller.NewHealthController()).
		Mount("service-provier", controller.NewSProvierController()).
		Mount("raising-plan", controller.NewRaisingPlanController()).
		ClassTask(filRewardTask, "0 */3 * * * *", filRewardTask.SyncFilBlockReward).
		ClassTask(filRewardTask, "0 0 */1 * * *", filRewardTask.SyncFilBlockRewardDistribution).
		ClassTask(planProgressTask, "0 0 */1 * * *", planProgressTask.SyncMinerSealedProgress).
		ClassTask(planProgressTask, "0 0 */1 * * *", planProgressTask.SyncMinerSealedStart).
		Launch()

}
