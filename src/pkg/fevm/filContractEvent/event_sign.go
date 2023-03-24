package filContractEvent

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/pkg/fevm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
)

func AllTopics() ethtypes.EthTopicSpec {
	topics := []ethtypes.EthHashList{{
		PutValueSign(),
		InvokeContractSign(),
		ECreateRaisePlanSign(),
		EStartRaisePlan(),
		ECloseRaisePlan(),
		EWithdrawRaiseSecurityFund(),
		EWithdrawOPSSecurityFund(),
		EStackFromInvestor(),
		EUnstackFromInverstor(),
		EWithdraw(),
		EPushBlockReward(),
		ERaiseSecurityFund(),
		EDepositOPSSecurityFund(),
		ERaiseFailed()}}

	return topics
}

func AllTopicsEth() [][]common.Hash {
	topics := [][]common.Hash{{
		common.HexToHash(PutValueSign().String()),
		common.HexToHash(InvokeContractSign().String()),
		common.HexToHash(ECreateRaisePlanSign().String()),
		common.HexToHash(EStartRaisePlan().String()),
		common.HexToHash(ECloseRaisePlan().String()),
		common.HexToHash(EWithdrawRaiseSecurityFund().String()),
		common.HexToHash(EWithdrawOPSSecurityFund().String()),
		common.HexToHash(EStackFromInvestor().String()),
		common.HexToHash(EUnstackFromInverstor().String()),
		common.HexToHash(EWithdraw().String()),
		common.HexToHash(EPushBlockReward().String()),
		common.HexToHash(ERaiseSecurityFund().String()),
		common.HexToHash(EDepositOPSSecurityFund().String()),
		common.HexToHash(ERaiseFailed().String())}}

	return topics
}

func PutValueSign() ethtypes.EthHash {
	sign := "PutValue(uint256)"
	return fevm.EthTopicHash(sign)
}

func InvokeContractSign() ethtypes.EthHash {
	sign := "InvokeContract(address,RaiseInfo)"
	return fevm.EthTopicHash(sign)
}

func ECreateRaisePlanSign() ethtypes.EthHash {
	sign := "eCreateRaisePlan(address,RaiseInfo)"
	return fevm.EthTopicHash(sign)
}

func EStartRaisePlan() ethtypes.EthHash {
	sign := "eStartRaisePlan(address,uint256)"
	return fevm.EthTopicHash(sign)
}

func ECloseRaisePlan() ethtypes.EthHash {
	sign := "eCloseRaisePlan(address,uint256)"
	return fevm.EthTopicHash(sign)
}

func EWithdrawRaiseSecurityFund() ethtypes.EthHash {
	sign := "eWithdrawRaiseSecurityFund(address,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

func EWithdrawOPSSecurityFund() ethtypes.EthHash {
	sign := "eWithdrawOPSSecurityFund(address,uint256,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者质押
func EStackFromInvestor() ethtypes.EthHash {
	sign := "eStackFromInvestor(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者取消质押
func EUnstackFromInverstor() ethtypes.EthHash {
	sign := "eUnstackFromInverstor(uint256,address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 投资者提取收益
func EWithdraw() ethtypes.EthHash {
	sign := "eWithdraw(address,address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 向合约压入爆块信息
func EPushBlockReward() ethtypes.EthHash {
	sign := "ePushBlockReward(bytes,uint256[],uint256[])"
	return fevm.EthTopicHash(sign)
}

// 缴纳募集计划保证金
func ERaiseSecurityFund() ethtypes.EthHash {
	sign := "eRaiseSecurityFund(address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 缴纳运维保证金
func EDepositOPSSecurityFund() ethtypes.EthHash {
	sign := "eDepositOPSSecurityFund(address,uint256)"
	return fevm.EthTopicHash(sign)
}

// 募集失败
func ERaiseFailed() ethtypes.EthHash {
	sign := "eRaiseFailed(uint256)"
	return fevm.EthTopicHash(sign)
}
