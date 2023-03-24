package controller

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/global/constant"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/goft"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/models"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/bigInt"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/glif"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/log"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/result"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/gin-gonic/gin"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"gorm.io/gorm"
)

type RaisingPlanController struct {
	MarketClient *client.MarketClient `inject:"-"`
	Db           *gorm.DB             `inject:"-"`
}

func NewRaisingPlanController() *RaisingPlanController {
	return &RaisingPlanController{}
}

func FormatPlanList(pList []*models.RaisingPlan) {
	for _, plan := range pList {
		plan.IncomeRate = transferIncomeToRate(plan.IncomeRate)
	}
}

func transferIncomeToRate(income float64) float64 {
	return float64(income / constant.CALC_ACCURACY)
}

func (this *RaisingPlanController) PlanAllList(c *gin.Context) goft.Json {
	pagination := GeneratePaginationFromRequest(c)

	// 查询所有的服务商
	var rPList []*models.RaisingPlan

	// 分页查询
	offset := (pagination.Page - 1) * pagination.PageSize
	log.Infof("offset:%d", offset)
	log.Infof("pagination.PageSize:%+v", pagination)
	if err := this.Db.Find(&rPList).Limit(pagination.PageSize).Offset(int(offset)).Order(pagination.Sort).Error; err != nil {
		return result.SpNotFound
	}
	FormatPlanList(rPList)

	return result.OK.WithData(rPList)
}

func (this *RaisingPlanController) PlanRaiserList(c *gin.Context) goft.Json {
	pagination := GeneratePaginationFromRequest(c)
	raiserAddress := c.Query("raiser_address")

	// 查询所有的服务商
	var rPList []*models.RaisingPlan

	// 分页查询
	offset := (pagination.Page - 1) * pagination.PageSize
	if err := this.Db.Where("Raiser = ? ", raiserAddress).Find(&rPList).Limit(pagination.PageSize).Offset(int(offset)).Order(pagination.Sort).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}
	FormatPlanList(rPList)

	return result.OK.WithData(rPList)
}

func (this *RaisingPlanController) PlanInfo(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")

	// 查询所有的服务商
	var rPList models.RaisingPlan

	// 分页查询
	if err := this.Db.Where("raising_id = ? ", raising_id).Find(&rPList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}
	rPList.IncomeRate = transferIncomeToRate(rPList.IncomeRate)

	return result.OK.WithData(rPList)
}

func (this *RaisingPlanController) PlanInvesterList(c *gin.Context) goft.Json {
	pagination := GeneratePaginationFromRequest(c)
	investerAddress := c.Query("invest_address")

	var invester []*models.FundDetails
	if err := this.Db.Where("from_address = ? ", investerAddress).Find(&invester).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}

	panList := make([]uint, 0)
	for _, v := range invester {
		panList = append(panList, v.RaiserId)
	}

	// 查询所有的服务商
	var rPList []*models.RaisingPlan

	// 分页查询
	offset := (pagination.Page - 1) * pagination.PageSize
	if err := this.Db.Where("raising_id in ? ", panList).Find(&rPList).Limit(pagination.PageSize).Offset(int(offset)).Order(pagination.Sort).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}
	FormatPlanList(rPList)

	return result.OK.WithData(rPList)
}

func (this *RaisingPlanController) PlanIncome(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")

	// 查询所有的服务商
	var rPList models.RaisingPlan

	// 分页查询
	if err := this.Db.Where("raising_id = ? ", raising_id).Find(&rPList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}

	// 查询miner算力
	minerBaseInfo, err := glif.GetMinerPower(rPList.MinerId)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}

	// 计算理论收益
	powerRate := types.BigDivFloat(minerBaseInfo.MinerPower.QualityAdjPower, minerBaseInfo.TotalPower.QualityAdjPower)
	ecIncome := float64(build.BlocksPerEpoch*2880*365) * powerRate
	targetAmout, err := bigInt.FromString(rPList.TargetAmount)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}
	log.Infof("MinerPower.QualityAdjPower:%+v", minerBaseInfo.MinerPower.QualityAdjPower)
	log.Infof("minerBaseInfo.TotalPower.QualityAdjPower:%+v", minerBaseInfo.TotalPower.QualityAdjPower)
	log.Infof("ecIncome:%+v", ecIncome)
	log.Infof("targetAmout:%+v", targetAmout)
	ecIncomeRate := types.BigDiv(types.BigInt{types.NewInt(uint64(ecIncome * constant.CALC_ACCURACY)).Int},
		types.BigInt{targetAmout.Int})

	// 计算实际收益 todo。。

	// 返回
	var reward_each_block = 17.0000
	ret := make(map[string]interface{})
	ret["ec_income_rate"] = float64(ecIncomeRate.Uint64())*reward_each_block/constant.CALC_ACCURACY - 1
	ret["miner"] = rPList.MinerId

	return result.OK.WithData(ret)
}

func (this *RaisingPlanController) PlanAlloctedAsset(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")

	var total_amt float64
	sqlstr := `select COALESCE(SUM(value),0)  as total_amt from fund_details where raiser_id =  ? and created_at >= DATE_SUB(NOW(), INTERVAL 1 DAY) `
	if err := this.Db.Debug().Raw(sqlstr, raising_id).Pluck("total_amt", &total_amt).Error; err != nil {
		log.Errorf("search DerivedGasOutput %+v", err)
		return result.SystemError.WithParamError(err)
	}
	ret := make(map[string]interface{})
	ret["total_amt"] = total_amt

	return result.OK.WithData(ret)
}

func (this *RaisingPlanController) PlanProgress(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")

	// 查询所有的服务商
	var rPList models.RaisingPlan

	// 分页查询
	if err := this.Db.Where("raising_id = ? ", raising_id).Find(&rPList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}

	// 查询miner算力
	minerBaseInfo, err := glif.GetMinerPower(rPList.MinerId)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}
	targetPower, err := bigInt.FromString(rPList.TargetPower)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}
	// 进度
	progress := bigInt.Div(bigInt.Mul(bigInt.NewBigInt(minerBaseInfo.MinerPower.QualityAdjPower.Int64()),
		bigInt.NewBigInt(constant.CALC_ACCURACY)), bigInt.BigInt{(targetPower.Int)})

	// 返回
	ret := make(map[string]interface{})
	ret["progress"] = float64(progress.Int64()) / constant.CALC_ACCURACY
	ret["miner"] = rPList.MinerId

	return result.OK.WithData(ret)
}

func (this *RaisingPlanController) GetInvestorCount(raise_id uint) (int64, error) {
	var count int64
	err := this.Db.Model(&models.FundDetails{}).Where("raiser_id = ?", raise_id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (this *RaisingPlanController) PlanAssetReport(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")

	// 查询所有的服务商
	var rPList models.RaisingPlan

	// 分页查询
	if err := this.Db.Where("raising_id = ? ", raising_id).Find(&rPList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return result.SpNotFound
		}
		return result.SystemError.WithParamError(err)
	}
	// 查询miner状态信息
	minerStats, err := glif.GetMinerState(rPList.MinerId)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}
	// 查询miner算力
	minerPower, err := glif.GetMinerPower(rPList.MinerId)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}

	investorCount, err := this.GetInvestorCount(rPList.RaisingId)
	if err != nil {
		return result.SystemError.WithParamError(err)
	}

	var planAssetReport RaiseAssetReport
	planAssetReport.MinerTargetCapac = rPList.TargetPower
	planAssetReport.MinerCapacSize = minerPower.MinerPower.QualityAdjPower.String()
	planAssetReport.PledgeFund = minerStats.InitialPledge.String()
	planAssetReport.VestingFund = minerStats.LockedFunds.String()
	planAssetReport.InvestorCount = investorCount
	planAssetReport.TotalAsset = ""
	planAssetReport.Reward24h1T = "213343543"

	return result.OK.WithData(planAssetReport)
}

func (this *RaisingPlanController) Name() string {
	return "SProvierController"
}

func (this *RaisingPlanController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/list", this.PlanAllList)
	goft.Handle("GET", "/raiser/list", this.PlanRaiserList)
	goft.Handle("GET", "/info", this.PlanInfo)
	goft.Handle("GET", "/invest/list", this.PlanInvesterList)
	goft.Handle("GET", "/income", this.PlanIncome)
	goft.Handle("GET", "/progress", this.PlanProgress)
	goft.Handle("GET", "/asset/report", this.PlanAssetReport)
	goft.Handle("GET", "/allocated/asset", this.PlanAlloctedAsset)

}
