package service

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/bigInt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/filContractFunc"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/glif"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/lilyApi"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"math/big"
)

type FilRewardService struct {
	Db          *gorm.DB          `inject:"-"`
	RedisClient *redis.Client     `inject:"-"`
	HttpsClient *ethclient.Client `inject:"-"`
	WSClient    *ethclient.Client `inject:"-"`
}

func NewFilRewardService() *FilRewardService {
	return &FilRewardService{}
}

type SyncFilReward struct {
	TotalReward   *big.Int `json:"total_reward"`
	LockReward    *big.Int `json:"lock_reward"`
	ReleaseReward *big.Int `json:"release_reward"`
}

func (this *FilRewardService) SyncFilBlockReward() {
	// 查询需要同步的minerId列表和高度列表
	var searchList []models.MinerInfo
	if err := this.Db.Find(&searchList).Error; err != nil {
		log.Errorf("查询需要同步的minerId列表和高度列表失败:%v", err)
		return
	}

	// 查询高度对应的奖励
	for _, minerInfo := range searchList {
		// 查询高度对应的奖励
		newReward, err := lilyApi.GetFilBlockReward(minerInfo.MinerId, minerInfo.Height)
		if err != nil {
			log.Errorf("查询高度对应的奖励失败:%v", err)
			return
		}
		var now_height uint64

		var rewards = make([]models.BlockReward, 0)
		for _, reward := range newReward {
			var releaseReward big.Int
			releaseReward.Div(releaseReward.Mul(reward.Reward.Int, big.NewInt(25)), big.NewInt(100))
			one_reward := models.BlockReward{
				MinerId:          reward.MinerID,
				Height:           uint64(reward.Height),
				TotalReward:      reward.Reward.String(),
				ReleaseReward:    releaseReward.String(),
				NextReleasHeight: uint64(reward.Height + 2880),
				ReleaseCount:     0,
			}
			rewards = append(rewards, one_reward)

			if now_height < uint64(reward.Height) {
				now_height = uint64(reward.Height)
			}
		}
		//  使用事物插入数据
		tx := this.Db.Begin()
		if err := tx.Create(&rewards).Error; err != nil {
			log.Errorf("记录合约调用参数失败:%v", err)
			return
		}

		// 更新minerInfo表中的高度
		if err := tx.Model(&models.MinerInfo{}).Where("miner_id = ?", minerInfo.MinerId).Update("height", now_height).Error; err != nil {
			log.Errorf("更新minerInfo表中的高度失败:%v", err)
			return
		}

		tx.Commit()

	}

}

// 同步奖励分配到合约中
func (this *FilRewardService) SyncFilBlockRewardDistribution() {
	// 获取当前链上区块高度
	chainHead, err := glif.GetChainHead()
	if err != nil {
		log.Errorf("获取当前链上区块高度失败:%v", err)
		return
	}

	// 查询需要释放的区块奖励
	var blockReward []*models.BlockReward
	if err := this.Db.Find(&blockReward, "next_releas_height <= ?", chainHead.Height).Error; err != nil {
		log.Errorf("查询需要释放的区块奖励失败:%v", err)
		return
	}

	// 计算需要释放的奖励并更新数据库
	for _, reward := range blockReward {

		epoch := uint64(chainHead.Height()) - reward.NextReleasHeight
		days := uint64(epoch / 2880)

		if reward.ReleaseCount > 180 {
			continue
		}

		for i := 0; i < int(days); i++ {
			// 计算需要释放的奖励
			var needReleaseReward big.Int

			totalReward, err := bigInt.FromString(reward.TotalReward)
			if err != nil {
				log.Errorf("计算需要释放的奖励失败:%v", err)
				return
			}
			releaseReward, err := bigInt.FromString(reward.ReleaseReward)
			if err != nil {
				log.Errorf("计算需要释放的奖励失败:%v", err)
				return
			}

			needReleaseReward.Div(needReleaseReward.Sub(totalReward.Int, releaseReward.Int), big.NewInt(180))

			reward.ReleaseReward = releaseReward.Add(releaseReward.Int, &needReleaseReward).String()
			reward.ReleaseCount += 1
			reward.NextReleasHeight += 2880
		}
		// 更新数据库
		if err := this.Db.Save(reward).Error; err != nil {
			log.Errorf("更新数据库失败:%v", err)
			return
		}
	}

	// 获取需要同步的募集计划，mineriD
	var raisePlan []*models.RaisingPlan
	if err := this.Db.Find(&raisePlan, "status = 0").Error; err != nil {
		log.Errorf("获取需要同步的募集计划失败:%v", err)
		return
	}
	// 获取miner 对应的待释放奖励、已释放奖励、待释放高度、已释放高度
	for _, plan := range raisePlan {
		// 查询 blockReward 中sum(releaseReward)、sum(TotalReward)
		var syncFilReward SyncFilReward
		if err := this.Db.Model(&models.BlockReward{}).Where("miner_id = ?",
			plan.MinerId).Select("COALESCE(sum(release_reward),0),COALESCE(sum(total_reward),0)").Scan(&syncFilReward).Error; err != nil {
			log.Errorf("查询 blockReward 中sum(releaseReward)、sum(TotalReward)失败:%v", err)
			return
		}
		// 调用合约推送释放奖励
		// todo...
		receipt, err := filContractFunc.PostBlockChainData(this.HttpsClient, plan.RaiseAddress, "putFilReward", syncFilReward)
		if err != nil {
			log.Errorf("调用合约推送释放奖励失败:%v", err)
			return
		}
		// 记录合约调用参数
		log.Infof("调用合约推送释放奖励成功:%v", receipt)
	}

}

func (this *FilRewardService) CheckFilBlockRewardResult() {

}
