package service

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/bigInt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/filContractFunc"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/glif"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/gorm"
)

type PlanProgressService struct {
	Db             *gorm.DB             `inject:"-"`
	RedisClient    *redis.Client        `inject:"-"`
	HttpsClient    *ethclient.Client    `inject:"-"`
	WSClient       *ethclient.Client    `inject:"-"`
	MarketClient   *client.MarketClient `inject:"-"`
	InfluxWriteApi api.WriteAPI         `inject:"-"`
}

func NewPlanProgressService() *PlanProgressService {
	return &PlanProgressService{}
}

func (this *PlanProgressService) SyncMinerSealedProgress() {
	// 查询需要同步的募集计划
	var searchList []models.RaisingPlan
	if err := this.Db.Find(&searchList).Error; err != nil {
		log.Errorf("查询需要同步的募集计划失败:%v", err)
		return
	}
	// 查询每个计划对应的进度
	for _, plan := range searchList {
		// 查询miner对应的进度
		// 查询miner算力
		minerBaseInfo, err := glif.GetMinerPower(plan.MinerId)
		if err != nil {
			return
		}
		targetPower, err := bigInt.FromString(plan.TargetPower)
		if err != nil {
			return
		}
		// 进度
		progress := bigInt.Div(bigInt.Mul(bigInt.NewBigInt(minerBaseInfo.MinerPower.QualityAdjPower.Int64()),
			bigInt.NewBigInt(constant.CALC_ACCURACY)), bigInt.BigInt{(targetPower.Int)})

		// 返回
		ret := make(map[string]interface{})
		ret["progress"] = float64(progress.Int64()) / constant.CALC_ACCURACY

		// 更新数据库
		if err := this.Db.Model(&models.RaisingPlan{}).Where("id = ?", plan.ID).Updates(ret).Error; err != nil {
			log.Errorf("更新数据库失败:%v", err)
			return
		}

		// todo 同步进度
		receipt, err := filContractFunc.PostBlockChainData(this.HttpsClient, plan.RaiseAddress, "PushSealProgress", nil)
		if err != nil {
			log.Errorf("调用合约推送释放奖励失败:%v", err)
			continue
		}
		log.Infof("调用合约推送释放奖励成功:%v", receipt)
		if receipt.Status == 1 {

		}
	}

}

func (this *PlanProgressService) SyncMinerSealedStart() {
	// 查询需要同步的募集计划
	var searchList []models.RaisingPlan
	if err := this.Db.Find(&searchList, " 1 = 1").Error; err != nil {
		log.Errorf("查询需要同步的募集计划失败:%v", err)
		return
	}
	// 查询每个计划对应的进度
	for _, plan := range searchList {
		// 查询miner对应的进度
		// 查询miner算力
		minerSectors, err := glif.GetMinerSectors(plan.MinerId)
		if err != nil {
			return
		}
		if minerSectors.Live == 0 {
			continue
		}
		if minerSectors.Live+minerSectors.Active+minerSectors.Faulty > 0 && plan.SealedStatus == 0 {
			// 调用合约推送释放奖励
			// todo...
			receipt, err := filContractFunc.PostBlockChainData(this.HttpsClient, plan.RaiseAddress, "StartSeal", nil)
			if err != nil {
				log.Errorf("调用合约推送释放奖励失败:%v", err)
				continue
			}
			log.Infof("调用合约推送释放奖励成功:%v", receipt)
			if receipt.Status == 1 {
				// 更新数据库
				if err := this.Db.Model(&models.RaisingPlan{}).Where("id = ?", plan.ID).Update("sealed_status", 2).Error; err != nil {
					log.Errorf("更新数据库失败:%v", err)
					continue
				}
			}
		}
		if minerSectors.Live+minerSectors.Active+minerSectors.Faulty == 0 && plan.SealedStatus == 1 {
			// 更新数据库
			// 调用合约推送释放奖励
			// todo...
			receipt, err := filContractFunc.PostBlockChainData(this.HttpsClient, plan.RaiseAddress, "EndMiner", nil)
			if err != nil {
				log.Errorf("调用合约推送释放奖励失败:%v", err)
				continue
			}
			if receipt.Status == 1 {
				// 更新募集计划为结束
				if err := this.Db.Model(&models.RaisingPlan{}).Where("id = ?", plan.ID).Update("status", 9).Error; err != nil {
					log.Errorf("更新数据库失败:%v", err)
					continue
				}
			}
			// todo 通知合约节点结束
		}

	}

}

func (this *PlanProgressService) SyncMinerSealedEnd() {
	// 查询需要同步的募集计划
	var searchList []models.RaisingPlan
	if err := this.Db.Find(&searchList).Error; err != nil {
		log.Errorf("查询需要同步的募集计划失败:%v", err)
		return
	}
	// 查询每个计划对应的进度
	for _, plan := range searchList {
		// 查询miner对应的进度
		// 查询miner算力
		minerPower, err := glif.GetMinerPower(plan.MinerId)
		if err != nil {
			return
		}
		log.Infof("minerPower:%v", minerPower)

		// 查询质押币数量
		minerState, err := glif.GetMinerState(plan.MinerId)
		if err != nil {
			return
		}
		plegeAmount := minerState.InitialPledge //
		var balance bigInt.BigInt
		targetAmount, err := bigInt.FromString(plan.TargetAmount)

		balance.Sub(plegeAmount.Int, targetAmount.Int)
		sealedPlageAmt := bigInt.NewBigInt(constant.SealedPlageAmt)

		// 质押币是否耗尽
		if balance.Abs().Cmp(sealedPlageAmt.Int) <= 0 {
			// 通知合约封装完成
			// 调用合约推送释放奖励
			// todo...
			receipt, err := filContractFunc.PostBlockChainData(this.HttpsClient, plan.RaiseAddress, "EndSeal", nil)
			if err != nil {
				log.Errorf("调用合约推送释放奖励失败:%v", err)
				continue
			}
			log.Infof("调用合约推送释放奖励成功:%v", receipt)
			// 更新数据库
			if err := this.Db.Model(&models.RaisingPlan{}).Where("id = ?", plan.ID).Update("sealed_status", 2).Error; err != nil {
				log.Errorf("更新数据库失败:%v", err)
				continue
			}
		}

	}

}
