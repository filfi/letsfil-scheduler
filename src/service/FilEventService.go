package service

import (
	"fmt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/contract"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm/filContractEvent"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"math/big"
)

type FilEventService struct {
	Db            *gorm.DB            `inject:"-"`
	RedisClient   *redis.Client       `inject:"-"`
	HttpsClient   *ethclient.Client   `inject:"-"`
	WSClient      *ethclient.Client   `inject:"-"`
	RaiseContract *contract.RaiseInfo `inject:"-"`
}

func NewFilEventService() *FilEventService {
	return &FilEventService{}
}

func (this *FilEventService) SyncRaiseEvent() {
	// 获取所有募集计划
	var raisePlans []*models.RaisingPlan

	if err := this.Db.Find(&raisePlans).Error; err != nil {
		return
	}

	// 获取最新的区块高度
	latestBlockNumber, err := this.HttpsClient.HeaderByNumber(nil, nil)
	if err != nil {
		return
	}

	startBigInt := big.NewInt(0)
	startBigInt = startBigInt.SetUint64(latestBlockNumber.Number.Uint64() - 10000)

	endBigInt := big.NewInt(0)
	endBigInt = endBigInt.SetUint64(latestBlockNumber.Number.Uint64())

	contractAddresses := make([]string, 0)
	// 如果有工厂合约，就加入到合约地址列表中
	if configs.Get().FilFevm.FactoryContract != "" {
		contractAddresses = append(contractAddresses, configs.Get().FilFevm.FactoryContract)
	}

	// 循环每个募集计划的合约地址，去获取区块上面的事件
	for _, raisePlan := range raisePlans {

		contractAddresses = append(contractAddresses, raisePlan.RaiseAddress)
	}

	ret, err := filContractEvent.FilterEventEth(startBigInt.Uint64(), endBigInt.Uint64(), contractAddresses)
	if err != nil {
		fmt.Println("SubscribeFilterLogsfmt", err)
		log.Error("SubscribeFilterLogs ", err)
	}
	logs := *ret

	for i := 0; i < len(logs); i++ {
		vLog := logs[i]
		var oneLog *ethtypes.EthLog
		oneLog, err := fevm.ParseEthLogs(vLog)
		if err != nil {
			log.Error("ParseEthLogs ", err)
			continue
		}

		log.Info("oneLog:", oneLog) // pointer to event log
		switch oneLog.Topics[0] {
		case filContractEvent.PutValueSign(): // 投资人投资
			log.Info("Transfer Data", oneLog.Data)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLockABI(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error("CreateLockABI ", err)
			}
			event := new(contract.LockPutValue)
			contractPool.UnpackLog(event, "PutValue", vLog)
			fmt.Printf("event: %+v\n", event)
		case filContractEvent.ECreateRaisePlanSign(): // 创建募集计划
			log.Info("Transfer Topics", oneLog.Topics)
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer Address", oneLog.Address)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			log.Info("Transfer BlockHash", oneLog.BlockHash)
			log.Info("Transfer TxHash", oneLog.Data)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateRaiseFactory(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaiseFactoryECreateRaisePlan)
			contractPool.UnpackLog(event, "eCreateRaisePlan", vLog)
			fmt.Printf("event: %+v\n", event)

			// 存入数据库创建募集计划
			var raisePlan = models.RaisingPlan{
				RaiseAddress:      event.RaisePool.String(),
				TxHash:            oneLog.TransactionHash.String(),
				RaisingId:         event.Info.Id.Uint64(),
				Raiser:            event.Info.Sponsor.String(),
				SponsorCompany:    event.Info.SponsorCompany,
				SecurityFundRate:  event.Info.SecurityFund.String(),
				MinerId:           string(event.Info.MinerID[:]),
				TargetPower:       event.NodeInfo.NodeSize.String(),
				TargetAmount:      event.Info.TargetAmount.String(),
				ClosingTime:       event.Info.ClosingTime.Uint64(),
				RaiserShare:       event.Info.RaiserShare.Uint64(),
				InvestorShare:     event.Info.InvestorShare.Uint64(),
				ServicerShare:     event.Info.ServicerShare.Uint64(),
				SectorSize:        event.NodeInfo.SectorSize.Uint64(),
				SealTimeLimit:     event.NodeInfo.SealTimeLimit.Uint64(),
				Status:            0,
				Progress:          0,
				IncomeRate:        0,
				SealedStatus:      0,
				LatestBlockNumber: uint64(oneLog.BlockNumber),
			}

			// 创建minerinfo记录
			var minerInfo = models.MinerInfo{
				MinerId: string(event.Info.MinerID[:]),
				Height:  uint64(oneLog.BlockNumber),
			}
			err = this.Db.Transaction(func(tx *gorm.DB) error {
				if err := this.Db.Create(&raisePlan).Error; err != nil {
					log.Error(" ", err)
					return err
				}
				if err := this.Db.Create(&minerInfo).Error; err != nil {
					log.Error(" ", err)
					return err
				}
				return nil
			})
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.EStartRaisePlan(): // 开始募集
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)

			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEStartRaisePlan)
			contractPool.UnpackLog(event, "eStartRaisePlan", vLog)
			fmt.Printf("event: %+v\n", event)
			// 更新募集计划状态为募集中
			var raisePlan = models.RaisingPlan{
				Status: 1,
			}
			err = this.Db.Model(&models.RaisingPlan{}).Where("raising_id = ?", event.RaiseID.Uint64()).Updates(&raisePlan).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.ECloseRaisePlan(): // 结束募集
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolECloseRaisePlan)
			contractPool.UnpackLog(event, "eCloseRaisePlan", vLog)
			fmt.Printf("event: %+v\n", event)
			// 更新募集计划状态为关闭
			var raisePlan = models.RaisingPlan{
				Status: 2,
			}
			err = this.Db.Model(&models.RaisingPlan{}).Where("raising_id = ?", event.RaiseID.Uint64()).Updates(&raisePlan).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.EWithdrawRaiseSecurityFund(): // 提取募集保证金
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEWithdrawRaiseSecurityFund)
			contractPool.UnpackLog(event, "eWithdrawRaiseSecurityFund", vLog)
			fmt.Printf("event: %+v\n", event)
			// 插入资金记录DepositDetails
			var depositDetails = models.DepositDetails{
				TxHash:      oneLog.TransactionHash.String(),
				DepositType: 1,

				DepositStatus:   2,
				FromAddress:     oneLog.Address.String(),
				ToAddress:       event.Caller.String(),
				Value:           event.Amount.String(),
				TransactionTime: oneLog.BlockNumber.Hex(),
				RaiserId:        event.RaiseID.Uint64(),
			}
			err = this.Db.Create(&depositDetails).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.EWithdrawOPSSecurityFund(): // 提取OPS保证金
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEWithdrawOPSSecurityFund)
			contractPool.UnpackLog(event, "eWithdrawOPSSecurityFund", vLog)
			fmt.Printf("event: %+v\n", event)
			// 插入资金记录DepositDetails
			var depositDetails = models.DepositDetails{
				TxHash:          oneLog.TransactionHash.String(),
				DepositType:     2,
				DepositStatus:   2,
				FromAddress:     oneLog.Address.String(),
				ToAddress:       event.Caller.String(),
				Value:           event.Amount.String(),
				TransactionTime: oneLog.BlockNumber.Hex(),
				RaiserId:        event.RaiseID.Uint64(),
			}
			err = this.Db.Create(&depositDetails).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}
		case filContractEvent.EStackFromInvestor(): // 投资人投资
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEStackFromInvestor)
			contractPool.UnpackLog(event, "eStackFromInvestor", vLog)
			fmt.Printf("event: %+v\n", event)
			// 插入资金记录fundDetails
			var fundDetails = models.FundDetails{
				TxHash:      oneLog.TransactionHash.String(),
				TxType:      1,
				FromAddress: event.From.String(),
				ToAddress:   event.To.String(),
				Value:       event.Amount.String(),
				RaiserId:    event.RaiseID.Uint64(),
			}
			err = this.Db.Create(&fundDetails).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.EUnstackFromInverstor(): // 投资人赎回
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEUnstackFromInverstor)
			contractPool.UnpackLog(event, "eUnstackFromInverstor", vLog)
			fmt.Printf("event: %+v\n", event)
			// 插入资金记录fundDetails
			var fundDetails = models.FundDetails{
				TxHash:      oneLog.TransactionHash.String(),
				TxType:      2,
				FromAddress: event.From.String(),
				ToAddress:   event.To.String(),
				Value:       event.Amount.String(),
				RaiserId:    event.RaiseID.Uint64(),
			}
			err = this.Db.Create(&fundDetails).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}
		case filContractEvent.EWithdraw(): // 提取
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)
			vLog := logs[i]
			fmt.Printf("vLog: %+v\n", vLog)
			contractPool, err := contract.CreateLetsFilRaisePool(this.HttpsClient, oneLog.Address.String())
			if err != nil {
				log.Error(" ", err)
			}
			event := new(contract.LetsFilRaisePoolEWithdraw)
			contractPool.UnpackLog(event, "eWithdraw", vLog)
			fmt.Printf("event: %+v\n", event)
			// 插入资金记录fundDetails
			var fundDetails = models.FundDetails{
				TxHash:      oneLog.TransactionHash.String(),
				TxType:      3,
				FromAddress: event.From.String(),
				ToAddress:   event.To.String(),
				Value:       event.Amount.String(),
				//RaiserId:    event.RaiseID.Uint64(),
			}
			err = this.Db.Create(&fundDetails).Error
			if err != nil {
				log.Error(" ", err)
				continue
			}

		case filContractEvent.EPushBlockReward(): // 推送区块奖励
			log.Info("Transfer Data", oneLog.Data)
			log.Info("Transfer BlockNumber", oneLog.BlockNumber)

		}

	}
}
