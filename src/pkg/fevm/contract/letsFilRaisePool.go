// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LetsFilRaisePoolMetaData contains all meta data concerning the LetsFilRaisePool contract.
var LetsFilRaisePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sponsorCompany\",\"type\":\"string\"}],\"internalType\":\"structILetsFilRaiseInfo.RaiseInfo\",\"name\":\"raiseInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealTimeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"opsSecurityFundPayer\",\"type\":\"address\"}],\"internalType\":\"structILetsFilRaiseInfo.NodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotWithdrawRaiseSecurityFund\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientInvestMinAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientRaiseSecurityFund\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientStakeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaiseHaveSucceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumRaiseState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"RaiseNotInProcess\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"EPushRewardParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"}],\"name\":\"eCloseRaisePlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eDepositOPSSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"}],\"name\":\"eNotifyNodeExpiration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eStackFromInvestor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"}],\"name\":\"eStartRaisePlan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eUnstackFromInverstor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eWithdrawOPSSecurityFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raiseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"eWithdrawRaiseSecurityFund\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"InvestMinAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"InvestorStackRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODEINFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealTimeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"opsSecurityFundPayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NodeExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPSSecurityFundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RAISEINFO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sponsorCompany\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RaiseSecurityFundAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalInvestmentAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inverstor\",\"type\":\"address\"}],\"name\":\"availableIncomeOf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeRaisePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositOPSSecurityFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remain\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mAvailableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mWithdrawedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyNodeExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"willRelease\",\"type\":\"uint256\"}],\"name\":\"pushRewardParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curPro\",\"type\":\"uint256\"}],\"name\":\"pushSealProgress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raiseIsSuccessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raiseState\",\"outputs\":[{\"internalType\":\"enumRaiseState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stackFromInvestor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"startRaisePlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstackFromInverstor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawOPSSecurityFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRaiseSecurityFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LetsFilRaisePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LetsFilRaisePoolMetaData.ABI instead.
var LetsFilRaisePoolABI = LetsFilRaisePoolMetaData.ABI

// LetsFilRaisePool is an auto generated Go binding around an Ethereum contract.
type LetsFilRaisePool struct {
	LetsFilRaisePoolCaller     // Read-only binding to the contract
	LetsFilRaisePoolTransactor // Write-only binding to the contract
	LetsFilRaisePoolFilterer   // Log filterer for contract events
}

// LetsFilRaisePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LetsFilRaisePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaisePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LetsFilRaisePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaisePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LetsFilRaisePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaisePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LetsFilRaisePoolSession struct {
	Contract     *LetsFilRaisePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LetsFilRaisePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LetsFilRaisePoolCallerSession struct {
	Contract *LetsFilRaisePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LetsFilRaisePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LetsFilRaisePoolTransactorSession struct {
	Contract     *LetsFilRaisePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LetsFilRaisePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LetsFilRaisePoolRaw struct {
	Contract *LetsFilRaisePool // Generic contract binding to access the raw methods on
}

// LetsFilRaisePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LetsFilRaisePoolCallerRaw struct {
	Contract *LetsFilRaisePoolCaller // Generic read-only contract binding to access the raw methods on
}

// LetsFilRaisePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LetsFilRaisePoolTransactorRaw struct {
	Contract *LetsFilRaisePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLetsFilRaisePool creates a new instance of LetsFilRaisePool, bound to a specific deployed contract.
func NewLetsFilRaisePool(address common.Address, backend bind.ContractBackend) (*LetsFilRaisePool, error) {
	contract, err := bindLetsFilRaisePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePool{LetsFilRaisePoolCaller: LetsFilRaisePoolCaller{contract: contract}, LetsFilRaisePoolTransactor: LetsFilRaisePoolTransactor{contract: contract}, LetsFilRaisePoolFilterer: LetsFilRaisePoolFilterer{contract: contract}}, nil
}

// NewLetsFilRaisePoolCaller creates a new read-only instance of LetsFilRaisePool, bound to a specific deployed contract.
func NewLetsFilRaisePoolCaller(address common.Address, caller bind.ContractCaller) (*LetsFilRaisePoolCaller, error) {
	contract, err := bindLetsFilRaisePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolCaller{contract: contract}, nil
}

// NewLetsFilRaisePoolTransactor creates a new write-only instance of LetsFilRaisePool, bound to a specific deployed contract.
func NewLetsFilRaisePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LetsFilRaisePoolTransactor, error) {
	contract, err := bindLetsFilRaisePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolTransactor{contract: contract}, nil
}

// NewLetsFilRaisePoolFilterer creates a new log filterer instance of LetsFilRaisePool, bound to a specific deployed contract.
func NewLetsFilRaisePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LetsFilRaisePoolFilterer, error) {
	contract, err := bindLetsFilRaisePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolFilterer{contract: contract}, nil
}

// bindLetsFilRaisePool binds a generic wrapper to an already deployed contract.
func bindLetsFilRaisePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LetsFilRaisePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilRaisePool *LetsFilRaisePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilRaisePool.Contract.LetsFilRaisePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilRaisePool *LetsFilRaisePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.LetsFilRaisePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilRaisePool *LetsFilRaisePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.LetsFilRaisePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilRaisePool *LetsFilRaisePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilRaisePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.contract.Transact(opts, method, params...)
}

// InvestMinAmount is a free data retrieval call binding the contract method 0x14b90747.
//
// Solidity: function InvestMinAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) InvestMinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "InvestMinAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvestMinAmount is a free data retrieval call binding the contract method 0x14b90747.
//
// Solidity: function InvestMinAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) InvestMinAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.InvestMinAmount(&_LetsFilRaisePool.CallOpts)
}

// InvestMinAmount is a free data retrieval call binding the contract method 0x14b90747.
//
// Solidity: function InvestMinAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) InvestMinAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.InvestMinAmount(&_LetsFilRaisePool.CallOpts)
}

// InvestorStackRecord is a free data retrieval call binding the contract method 0x317d21fa.
//
// Solidity: function InvestorStackRecord(address ) view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) InvestorStackRecord(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "InvestorStackRecord", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvestorStackRecord is a free data retrieval call binding the contract method 0x317d21fa.
//
// Solidity: function InvestorStackRecord(address ) view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) InvestorStackRecord(arg0 common.Address) (*big.Int, error) {
	return _LetsFilRaisePool.Contract.InvestorStackRecord(&_LetsFilRaisePool.CallOpts, arg0)
}

// InvestorStackRecord is a free data retrieval call binding the contract method 0x317d21fa.
//
// Solidity: function InvestorStackRecord(address ) view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) InvestorStackRecord(arg0 common.Address) (*big.Int, error) {
	return _LetsFilRaisePool.Contract.InvestorStackRecord(&_LetsFilRaisePool.CallOpts, arg0)
}

// NODEINFO is a free data retrieval call binding the contract method 0x32f02c59.
//
// Solidity: function NODEINFO() view returns(uint256 nodeSize, uint256 sectorSize, uint256 sealTimeLimit, uint256 nodePeriod, uint256 opsSecurityFund, address opsSecurityFundPayer)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) NODEINFO(opts *bind.CallOpts) (struct {
	NodeSize             *big.Int
	SectorSize           *big.Int
	SealTimeLimit        *big.Int
	NodePeriod           *big.Int
	OpsSecurityFund      *big.Int
	OpsSecurityFundPayer common.Address
}, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "NODEINFO")

	outstruct := new(struct {
		NodeSize             *big.Int
		SectorSize           *big.Int
		SealTimeLimit        *big.Int
		NodePeriod           *big.Int
		OpsSecurityFund      *big.Int
		OpsSecurityFundPayer common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeSize = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SectorSize = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SealTimeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NodePeriod = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OpsSecurityFund = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OpsSecurityFundPayer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// NODEINFO is a free data retrieval call binding the contract method 0x32f02c59.
//
// Solidity: function NODEINFO() view returns(uint256 nodeSize, uint256 sectorSize, uint256 sealTimeLimit, uint256 nodePeriod, uint256 opsSecurityFund, address opsSecurityFundPayer)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) NODEINFO() (struct {
	NodeSize             *big.Int
	SectorSize           *big.Int
	SealTimeLimit        *big.Int
	NodePeriod           *big.Int
	OpsSecurityFund      *big.Int
	OpsSecurityFundPayer common.Address
}, error) {
	return _LetsFilRaisePool.Contract.NODEINFO(&_LetsFilRaisePool.CallOpts)
}

// NODEINFO is a free data retrieval call binding the contract method 0x32f02c59.
//
// Solidity: function NODEINFO() view returns(uint256 nodeSize, uint256 sectorSize, uint256 sealTimeLimit, uint256 nodePeriod, uint256 opsSecurityFund, address opsSecurityFundPayer)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) NODEINFO() (struct {
	NodeSize             *big.Int
	SectorSize           *big.Int
	SealTimeLimit        *big.Int
	NodePeriod           *big.Int
	OpsSecurityFund      *big.Int
	OpsSecurityFundPayer common.Address
}, error) {
	return _LetsFilRaisePool.Contract.NODEINFO(&_LetsFilRaisePool.CallOpts)
}

// NodeExpired is a free data retrieval call binding the contract method 0xdc5ec46d.
//
// Solidity: function NodeExpired() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) NodeExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "NodeExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NodeExpired is a free data retrieval call binding the contract method 0xdc5ec46d.
//
// Solidity: function NodeExpired() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) NodeExpired() (bool, error) {
	return _LetsFilRaisePool.Contract.NodeExpired(&_LetsFilRaisePool.CallOpts)
}

// NodeExpired is a free data retrieval call binding the contract method 0xdc5ec46d.
//
// Solidity: function NodeExpired() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) NodeExpired() (bool, error) {
	return _LetsFilRaisePool.Contract.NodeExpired(&_LetsFilRaisePool.CallOpts)
}

// OPSSecurityFundAmount is a free data retrieval call binding the contract method 0x53ea083e.
//
// Solidity: function OPSSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) OPSSecurityFundAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "OPSSecurityFundAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPSSecurityFundAmount is a free data retrieval call binding the contract method 0x53ea083e.
//
// Solidity: function OPSSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) OPSSecurityFundAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.OPSSecurityFundAmount(&_LetsFilRaisePool.CallOpts)
}

// OPSSecurityFundAmount is a free data retrieval call binding the contract method 0x53ea083e.
//
// Solidity: function OPSSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) OPSSecurityFundAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.OPSSecurityFundAmount(&_LetsFilRaisePool.CallOpts)
}

// RAISEINFO is a free data retrieval call binding the contract method 0x9604c09c.
//
// Solidity: function RAISEINFO() view returns(uint256 id, uint256 targetAmount, uint256 securityFund, uint256 closingTime, uint256 raiserShare, uint256 investorShare, uint256 servicerShare, address sponsor, bytes32 minerID, string sponsorCompany)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) RAISEINFO(opts *bind.CallOpts) (struct {
	Id             *big.Int
	TargetAmount   *big.Int
	SecurityFund   *big.Int
	ClosingTime    *big.Int
	RaiserShare    *big.Int
	InvestorShare  *big.Int
	ServicerShare  *big.Int
	Sponsor        common.Address
	MinerID        [32]byte
	SponsorCompany string
}, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "RAISEINFO")

	outstruct := new(struct {
		Id             *big.Int
		TargetAmount   *big.Int
		SecurityFund   *big.Int
		ClosingTime    *big.Int
		RaiserShare    *big.Int
		InvestorShare  *big.Int
		ServicerShare  *big.Int
		Sponsor        common.Address
		MinerID        [32]byte
		SponsorCompany string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TargetAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecurityFund = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ClosingTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RaiserShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.InvestorShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ServicerShare = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Sponsor = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.MinerID = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.SponsorCompany = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// RAISEINFO is a free data retrieval call binding the contract method 0x9604c09c.
//
// Solidity: function RAISEINFO() view returns(uint256 id, uint256 targetAmount, uint256 securityFund, uint256 closingTime, uint256 raiserShare, uint256 investorShare, uint256 servicerShare, address sponsor, bytes32 minerID, string sponsorCompany)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) RAISEINFO() (struct {
	Id             *big.Int
	TargetAmount   *big.Int
	SecurityFund   *big.Int
	ClosingTime    *big.Int
	RaiserShare    *big.Int
	InvestorShare  *big.Int
	ServicerShare  *big.Int
	Sponsor        common.Address
	MinerID        [32]byte
	SponsorCompany string
}, error) {
	return _LetsFilRaisePool.Contract.RAISEINFO(&_LetsFilRaisePool.CallOpts)
}

// RAISEINFO is a free data retrieval call binding the contract method 0x9604c09c.
//
// Solidity: function RAISEINFO() view returns(uint256 id, uint256 targetAmount, uint256 securityFund, uint256 closingTime, uint256 raiserShare, uint256 investorShare, uint256 servicerShare, address sponsor, bytes32 minerID, string sponsorCompany)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) RAISEINFO() (struct {
	Id             *big.Int
	TargetAmount   *big.Int
	SecurityFund   *big.Int
	ClosingTime    *big.Int
	RaiserShare    *big.Int
	InvestorShare  *big.Int
	ServicerShare  *big.Int
	Sponsor        common.Address
	MinerID        [32]byte
	SponsorCompany string
}, error) {
	return _LetsFilRaisePool.Contract.RAISEINFO(&_LetsFilRaisePool.CallOpts)
}

// RaiseSecurityFundAmount is a free data retrieval call binding the contract method 0x5eb07b21.
//
// Solidity: function RaiseSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) RaiseSecurityFundAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "RaiseSecurityFundAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiseSecurityFundAmount is a free data retrieval call binding the contract method 0x5eb07b21.
//
// Solidity: function RaiseSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) RaiseSecurityFundAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.RaiseSecurityFundAmount(&_LetsFilRaisePool.CallOpts)
}

// RaiseSecurityFundAmount is a free data retrieval call binding the contract method 0x5eb07b21.
//
// Solidity: function RaiseSecurityFundAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) RaiseSecurityFundAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.RaiseSecurityFundAmount(&_LetsFilRaisePool.CallOpts)
}

// TotalInvestmentAmount is a free data retrieval call binding the contract method 0xa659f4dc.
//
// Solidity: function TotalInvestmentAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) TotalInvestmentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "TotalInvestmentAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInvestmentAmount is a free data retrieval call binding the contract method 0xa659f4dc.
//
// Solidity: function TotalInvestmentAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) TotalInvestmentAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.TotalInvestmentAmount(&_LetsFilRaisePool.CallOpts)
}

// TotalInvestmentAmount is a free data retrieval call binding the contract method 0xa659f4dc.
//
// Solidity: function TotalInvestmentAmount() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) TotalInvestmentAmount() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.TotalInvestmentAmount(&_LetsFilRaisePool.CallOpts)
}

// AvailableIncome is a free data retrieval call binding the contract method 0xb718ee94.
//
// Solidity: function availableIncome() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) AvailableIncome(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "availableIncome")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableIncome is a free data retrieval call binding the contract method 0xb718ee94.
//
// Solidity: function availableIncome() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) AvailableIncome() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.AvailableIncome(&_LetsFilRaisePool.CallOpts)
}

// AvailableIncome is a free data retrieval call binding the contract method 0xb718ee94.
//
// Solidity: function availableIncome() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) AvailableIncome() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.AvailableIncome(&_LetsFilRaisePool.CallOpts)
}

// MAvailableReward is a free data retrieval call binding the contract method 0xa905debc.
//
// Solidity: function mAvailableReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) MAvailableReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "mAvailableReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAvailableReward is a free data retrieval call binding the contract method 0xa905debc.
//
// Solidity: function mAvailableReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) MAvailableReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MAvailableReward(&_LetsFilRaisePool.CallOpts)
}

// MAvailableReward is a free data retrieval call binding the contract method 0xa905debc.
//
// Solidity: function mAvailableReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) MAvailableReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MAvailableReward(&_LetsFilRaisePool.CallOpts)
}

// MTotalReward is a free data retrieval call binding the contract method 0xe6c5a626.
//
// Solidity: function mTotalReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) MTotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "mTotalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MTotalReward is a free data retrieval call binding the contract method 0xe6c5a626.
//
// Solidity: function mTotalReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) MTotalReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MTotalReward(&_LetsFilRaisePool.CallOpts)
}

// MTotalReward is a free data retrieval call binding the contract method 0xe6c5a626.
//
// Solidity: function mTotalReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) MTotalReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MTotalReward(&_LetsFilRaisePool.CallOpts)
}

// MWithdrawedReward is a free data retrieval call binding the contract method 0xe2b6e729.
//
// Solidity: function mWithdrawedReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) MWithdrawedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "mWithdrawedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWithdrawedReward is a free data retrieval call binding the contract method 0xe2b6e729.
//
// Solidity: function mWithdrawedReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) MWithdrawedReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MWithdrawedReward(&_LetsFilRaisePool.CallOpts)
}

// MWithdrawedReward is a free data retrieval call binding the contract method 0xe2b6e729.
//
// Solidity: function mWithdrawedReward() view returns(uint256)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) MWithdrawedReward() (*big.Int, error) {
	return _LetsFilRaisePool.Contract.MWithdrawedReward(&_LetsFilRaisePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) Owner() (common.Address, error) {
	return _LetsFilRaisePool.Contract.Owner(&_LetsFilRaisePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) Owner() (common.Address, error) {
	return _LetsFilRaisePool.Contract.Owner(&_LetsFilRaisePool.CallOpts)
}

// RaiseIsSuccessed is a free data retrieval call binding the contract method 0xc144b797.
//
// Solidity: function raiseIsSuccessed() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) RaiseIsSuccessed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "raiseIsSuccessed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RaiseIsSuccessed is a free data retrieval call binding the contract method 0xc144b797.
//
// Solidity: function raiseIsSuccessed() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) RaiseIsSuccessed() (bool, error) {
	return _LetsFilRaisePool.Contract.RaiseIsSuccessed(&_LetsFilRaisePool.CallOpts)
}

// RaiseIsSuccessed is a free data retrieval call binding the contract method 0xc144b797.
//
// Solidity: function raiseIsSuccessed() view returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) RaiseIsSuccessed() (bool, error) {
	return _LetsFilRaisePool.Contract.RaiseIsSuccessed(&_LetsFilRaisePool.CallOpts)
}

// RaiseState is a free data retrieval call binding the contract method 0xb6b6a6ae.
//
// Solidity: function raiseState() view returns(uint8)
func (_LetsFilRaisePool *LetsFilRaisePoolCaller) RaiseState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LetsFilRaisePool.contract.Call(opts, &out, "raiseState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RaiseState is a free data retrieval call binding the contract method 0xb6b6a6ae.
//
// Solidity: function raiseState() view returns(uint8)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) RaiseState() (uint8, error) {
	return _LetsFilRaisePool.Contract.RaiseState(&_LetsFilRaisePool.CallOpts)
}

// RaiseState is a free data retrieval call binding the contract method 0xb6b6a6ae.
//
// Solidity: function raiseState() view returns(uint8)
func (_LetsFilRaisePool *LetsFilRaisePoolCallerSession) RaiseState() (uint8, error) {
	return _LetsFilRaisePool.Contract.RaiseState(&_LetsFilRaisePool.CallOpts)
}

// AvailableIncomeOf is a paid mutator transaction binding the contract method 0xc123aaca.
//
// Solidity: function availableIncomeOf(address inverstor) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) AvailableIncomeOf(opts *bind.TransactOpts, inverstor common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "availableIncomeOf", inverstor)
}

// AvailableIncomeOf is a paid mutator transaction binding the contract method 0xc123aaca.
//
// Solidity: function availableIncomeOf(address inverstor) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) AvailableIncomeOf(inverstor common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.AvailableIncomeOf(&_LetsFilRaisePool.TransactOpts, inverstor)
}

// AvailableIncomeOf is a paid mutator transaction binding the contract method 0xc123aaca.
//
// Solidity: function availableIncomeOf(address inverstor) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) AvailableIncomeOf(inverstor common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.AvailableIncomeOf(&_LetsFilRaisePool.TransactOpts, inverstor)
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xd77acecd.
//
// Solidity: function closeRaisePlan() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) CloseRaisePlan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "closeRaisePlan")
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xd77acecd.
//
// Solidity: function closeRaisePlan() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) CloseRaisePlan() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.CloseRaisePlan(&_LetsFilRaisePool.TransactOpts)
}

// CloseRaisePlan is a paid mutator transaction binding the contract method 0xd77acecd.
//
// Solidity: function closeRaisePlan() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) CloseRaisePlan() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.CloseRaisePlan(&_LetsFilRaisePool.TransactOpts)
}

// DepositOPSSecurityFund is a paid mutator transaction binding the contract method 0x874eb233.
//
// Solidity: function depositOPSSecurityFund() payable returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) DepositOPSSecurityFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "depositOPSSecurityFund")
}

// DepositOPSSecurityFund is a paid mutator transaction binding the contract method 0x874eb233.
//
// Solidity: function depositOPSSecurityFund() payable returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) DepositOPSSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.DepositOPSSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// DepositOPSSecurityFund is a paid mutator transaction binding the contract method 0x874eb233.
//
// Solidity: function depositOPSSecurityFund() payable returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) DepositOPSSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.DepositOPSSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// GetTotalIncome is a paid mutator transaction binding the contract method 0x24b6046a.
//
// Solidity: function getTotalIncome() returns(uint256 total, uint256 remain)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) GetTotalIncome(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "getTotalIncome")
}

// GetTotalIncome is a paid mutator transaction binding the contract method 0x24b6046a.
//
// Solidity: function getTotalIncome() returns(uint256 total, uint256 remain)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) GetTotalIncome() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.GetTotalIncome(&_LetsFilRaisePool.TransactOpts)
}

// GetTotalIncome is a paid mutator transaction binding the contract method 0x24b6046a.
//
// Solidity: function getTotalIncome() returns(uint256 total, uint256 remain)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) GetTotalIncome() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.GetTotalIncome(&_LetsFilRaisePool.TransactOpts)
}

// NotifyNodeExpiration is a paid mutator transaction binding the contract method 0xbac00500.
//
// Solidity: function notifyNodeExpiration() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) NotifyNodeExpiration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "notifyNodeExpiration")
}

// NotifyNodeExpiration is a paid mutator transaction binding the contract method 0xbac00500.
//
// Solidity: function notifyNodeExpiration() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) NotifyNodeExpiration() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.NotifyNodeExpiration(&_LetsFilRaisePool.TransactOpts)
}

// NotifyNodeExpiration is a paid mutator transaction binding the contract method 0xbac00500.
//
// Solidity: function notifyNodeExpiration() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) NotifyNodeExpiration() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.NotifyNodeExpiration(&_LetsFilRaisePool.TransactOpts)
}

// PushRewardParams is a paid mutator transaction binding the contract method 0xd73cda01.
//
// Solidity: function pushRewardParams(uint256 released, uint256 willRelease) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) PushRewardParams(opts *bind.TransactOpts, released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "pushRewardParams", released, willRelease)
}

// PushRewardParams is a paid mutator transaction binding the contract method 0xd73cda01.
//
// Solidity: function pushRewardParams(uint256 released, uint256 willRelease) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) PushRewardParams(released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.PushRewardParams(&_LetsFilRaisePool.TransactOpts, released, willRelease)
}

// PushRewardParams is a paid mutator transaction binding the contract method 0xd73cda01.
//
// Solidity: function pushRewardParams(uint256 released, uint256 willRelease) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) PushRewardParams(released *big.Int, willRelease *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.PushRewardParams(&_LetsFilRaisePool.TransactOpts, released, willRelease)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0x9345dd26.
//
// Solidity: function pushSealProgress(uint256 curPro) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) PushSealProgress(opts *bind.TransactOpts, curPro *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "pushSealProgress", curPro)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0x9345dd26.
//
// Solidity: function pushSealProgress(uint256 curPro) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) PushSealProgress(curPro *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.PushSealProgress(&_LetsFilRaisePool.TransactOpts, curPro)
}

// PushSealProgress is a paid mutator transaction binding the contract method 0x9345dd26.
//
// Solidity: function pushSealProgress(uint256 curPro) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) PushSealProgress(curPro *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.PushSealProgress(&_LetsFilRaisePool.TransactOpts, curPro)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.RenounceOwnership(&_LetsFilRaisePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.RenounceOwnership(&_LetsFilRaisePool.TransactOpts)
}

// StackFromInvestor is a paid mutator transaction binding the contract method 0x7d26b6f4.
//
// Solidity: function stackFromInvestor() payable returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) StackFromInvestor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "stackFromInvestor")
}

// StackFromInvestor is a paid mutator transaction binding the contract method 0x7d26b6f4.
//
// Solidity: function stackFromInvestor() payable returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) StackFromInvestor() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StackFromInvestor(&_LetsFilRaisePool.TransactOpts)
}

// StackFromInvestor is a paid mutator transaction binding the contract method 0x7d26b6f4.
//
// Solidity: function stackFromInvestor() payable returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) StackFromInvestor() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StackFromInvestor(&_LetsFilRaisePool.TransactOpts)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0x8dd44429.
//
// Solidity: function startRaisePlan(bytes signature) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) StartRaisePlan(opts *bind.TransactOpts, signature []byte) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "startRaisePlan", signature)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0x8dd44429.
//
// Solidity: function startRaisePlan(bytes signature) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) StartRaisePlan(signature []byte) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StartRaisePlan(&_LetsFilRaisePool.TransactOpts, signature)
}

// StartRaisePlan is a paid mutator transaction binding the contract method 0x8dd44429.
//
// Solidity: function startRaisePlan(bytes signature) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) StartRaisePlan(signature []byte) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StartRaisePlan(&_LetsFilRaisePool.TransactOpts, signature)
}

// StartSeal is a paid mutator transaction binding the contract method 0x0c5211a5.
//
// Solidity: function startSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) StartSeal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "startSeal")
}

// StartSeal is a paid mutator transaction binding the contract method 0x0c5211a5.
//
// Solidity: function startSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) StartSeal() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StartSeal(&_LetsFilRaisePool.TransactOpts)
}

// StartSeal is a paid mutator transaction binding the contract method 0x0c5211a5.
//
// Solidity: function startSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) StartSeal() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StartSeal(&_LetsFilRaisePool.TransactOpts)
}

// StopSeal is a paid mutator transaction binding the contract method 0x1c0aef7a.
//
// Solidity: function stopSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) StopSeal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "stopSeal")
}

// StopSeal is a paid mutator transaction binding the contract method 0x1c0aef7a.
//
// Solidity: function stopSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) StopSeal() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StopSeal(&_LetsFilRaisePool.TransactOpts)
}

// StopSeal is a paid mutator transaction binding the contract method 0x1c0aef7a.
//
// Solidity: function stopSeal() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) StopSeal() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.StopSeal(&_LetsFilRaisePool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.TransferOwnership(&_LetsFilRaisePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.TransferOwnership(&_LetsFilRaisePool.TransactOpts, newOwner)
}

// UnstackFromInverstor is a paid mutator transaction binding the contract method 0x3e2f1f20.
//
// Solidity: function unstackFromInverstor(uint256 amount) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) UnstackFromInverstor(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "unstackFromInverstor", amount)
}

// UnstackFromInverstor is a paid mutator transaction binding the contract method 0x3e2f1f20.
//
// Solidity: function unstackFromInverstor(uint256 amount) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) UnstackFromInverstor(amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.UnstackFromInverstor(&_LetsFilRaisePool.TransactOpts, amount)
}

// UnstackFromInverstor is a paid mutator transaction binding the contract method 0x3e2f1f20.
//
// Solidity: function unstackFromInverstor(uint256 amount) returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) UnstackFromInverstor(amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.UnstackFromInverstor(&_LetsFilRaisePool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.Withdraw(&_LetsFilRaisePool.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns(bool)
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.Withdraw(&_LetsFilRaisePool.TransactOpts, to, amount)
}

// WithdrawOPSSecurityFund is a paid mutator transaction binding the contract method 0x995d0fe6.
//
// Solidity: function withdrawOPSSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) WithdrawOPSSecurityFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "withdrawOPSSecurityFund")
}

// WithdrawOPSSecurityFund is a paid mutator transaction binding the contract method 0x995d0fe6.
//
// Solidity: function withdrawOPSSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) WithdrawOPSSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.WithdrawOPSSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// WithdrawOPSSecurityFund is a paid mutator transaction binding the contract method 0x995d0fe6.
//
// Solidity: function withdrawOPSSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) WithdrawOPSSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.WithdrawOPSSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// WithdrawRaiseSecurityFund is a paid mutator transaction binding the contract method 0x6e4dadcd.
//
// Solidity: function withdrawRaiseSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactor) WithdrawRaiseSecurityFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaisePool.contract.Transact(opts, "withdrawRaiseSecurityFund")
}

// WithdrawRaiseSecurityFund is a paid mutator transaction binding the contract method 0x6e4dadcd.
//
// Solidity: function withdrawRaiseSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolSession) WithdrawRaiseSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.WithdrawRaiseSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// WithdrawRaiseSecurityFund is a paid mutator transaction binding the contract method 0x6e4dadcd.
//
// Solidity: function withdrawRaiseSecurityFund() returns()
func (_LetsFilRaisePool *LetsFilRaisePoolTransactorSession) WithdrawRaiseSecurityFund() (*types.Transaction, error) {
	return _LetsFilRaisePool.Contract.WithdrawRaiseSecurityFund(&_LetsFilRaisePool.TransactOpts)
}

// LetsFilRaisePoolEPushRewardParamsIterator is returned from FilterEPushRewardParams and is used to iterate over the raw logs and unpacked data for EPushRewardParams events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEPushRewardParamsIterator struct {
	Event *LetsFilRaisePoolEPushRewardParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEPushRewardParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEPushRewardParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEPushRewardParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEPushRewardParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEPushRewardParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEPushRewardParams represents a EPushRewardParams event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEPushRewardParams struct {
	Released    *big.Int
	WillRelease *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEPushRewardParams is a free log retrieval operation binding the contract event 0x490ea370552e584622420726b92efcbe4a900e2d2dc0545622a2073bf3c09c14.
//
// Solidity: event EPushRewardParams(uint256 released, uint256 willRelease)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEPushRewardParams(opts *bind.FilterOpts) (*LetsFilRaisePoolEPushRewardParamsIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "EPushRewardParams")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEPushRewardParamsIterator{contract: _LetsFilRaisePool.contract, event: "EPushRewardParams", logs: logs, sub: sub}, nil
}

// WatchEPushRewardParams is a free log subscription operation binding the contract event 0x490ea370552e584622420726b92efcbe4a900e2d2dc0545622a2073bf3c09c14.
//
// Solidity: event EPushRewardParams(uint256 released, uint256 willRelease)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEPushRewardParams(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEPushRewardParams) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "EPushRewardParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEPushRewardParams)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "EPushRewardParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEPushRewardParams is a log parse operation binding the contract event 0x490ea370552e584622420726b92efcbe4a900e2d2dc0545622a2073bf3c09c14.
//
// Solidity: event EPushRewardParams(uint256 released, uint256 willRelease)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEPushRewardParams(log types.Log) (*LetsFilRaisePoolEPushRewardParams, error) {
	event := new(LetsFilRaisePoolEPushRewardParams)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "EPushRewardParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolOwnershipTransferredIterator struct {
	Event *LetsFilRaisePoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolOwnershipTransferred represents a OwnershipTransferred event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LetsFilRaisePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolOwnershipTransferredIterator{contract: _LetsFilRaisePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolOwnershipTransferred)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseOwnershipTransferred(log types.Log) (*LetsFilRaisePoolOwnershipTransferred, error) {
	event := new(LetsFilRaisePoolOwnershipTransferred)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolECloseRaisePlanIterator is returned from FilterECloseRaisePlan and is used to iterate over the raw logs and unpacked data for ECloseRaisePlan events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolECloseRaisePlanIterator struct {
	Event *LetsFilRaisePoolECloseRaisePlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolECloseRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolECloseRaisePlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolECloseRaisePlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolECloseRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolECloseRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolECloseRaisePlan represents a ECloseRaisePlan event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolECloseRaisePlan struct {
	Caller  common.Address
	RaiseID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterECloseRaisePlan is a free log retrieval operation binding the contract event 0x96bea66445b93064905ab9484bca37afea06fc23524c53a1f8ac18f76d193f39.
//
// Solidity: event eCloseRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterECloseRaisePlan(opts *bind.FilterOpts) (*LetsFilRaisePoolECloseRaisePlanIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eCloseRaisePlan")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolECloseRaisePlanIterator{contract: _LetsFilRaisePool.contract, event: "eCloseRaisePlan", logs: logs, sub: sub}, nil
}

// WatchECloseRaisePlan is a free log subscription operation binding the contract event 0x96bea66445b93064905ab9484bca37afea06fc23524c53a1f8ac18f76d193f39.
//
// Solidity: event eCloseRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchECloseRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolECloseRaisePlan) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eCloseRaisePlan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolECloseRaisePlan)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eCloseRaisePlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseECloseRaisePlan is a log parse operation binding the contract event 0x96bea66445b93064905ab9484bca37afea06fc23524c53a1f8ac18f76d193f39.
//
// Solidity: event eCloseRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseECloseRaisePlan(log types.Log) (*LetsFilRaisePoolECloseRaisePlan, error) {
	event := new(LetsFilRaisePoolECloseRaisePlan)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eCloseRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEDepositOPSSecurityFundIterator is returned from FilterEDepositOPSSecurityFund and is used to iterate over the raw logs and unpacked data for EDepositOPSSecurityFund events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEDepositOPSSecurityFundIterator struct {
	Event *LetsFilRaisePoolEDepositOPSSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEDepositOPSSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEDepositOPSSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEDepositOPSSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEDepositOPSSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEDepositOPSSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEDepositOPSSecurityFund represents a EDepositOPSSecurityFund event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEDepositOPSSecurityFund struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEDepositOPSSecurityFund is a free log retrieval operation binding the contract event 0xc1ae3fd0f6bc23c996226fc7183a444799c943624383d4de87db8f2abf75969f.
//
// Solidity: event eDepositOPSSecurityFund(address sender, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEDepositOPSSecurityFund(opts *bind.FilterOpts) (*LetsFilRaisePoolEDepositOPSSecurityFundIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eDepositOPSSecurityFund")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEDepositOPSSecurityFundIterator{contract: _LetsFilRaisePool.contract, event: "eDepositOPSSecurityFund", logs: logs, sub: sub}, nil
}

// WatchEDepositOPSSecurityFund is a free log subscription operation binding the contract event 0xc1ae3fd0f6bc23c996226fc7183a444799c943624383d4de87db8f2abf75969f.
//
// Solidity: event eDepositOPSSecurityFund(address sender, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEDepositOPSSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEDepositOPSSecurityFund) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eDepositOPSSecurityFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEDepositOPSSecurityFund)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eDepositOPSSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEDepositOPSSecurityFund is a log parse operation binding the contract event 0xc1ae3fd0f6bc23c996226fc7183a444799c943624383d4de87db8f2abf75969f.
//
// Solidity: event eDepositOPSSecurityFund(address sender, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEDepositOPSSecurityFund(log types.Log) (*LetsFilRaisePoolEDepositOPSSecurityFund, error) {
	event := new(LetsFilRaisePoolEDepositOPSSecurityFund)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eDepositOPSSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolENotifyNodeExpirationIterator is returned from FilterENotifyNodeExpiration and is used to iterate over the raw logs and unpacked data for ENotifyNodeExpiration events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolENotifyNodeExpirationIterator struct {
	Event *LetsFilRaisePoolENotifyNodeExpiration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolENotifyNodeExpirationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolENotifyNodeExpiration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolENotifyNodeExpiration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolENotifyNodeExpirationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolENotifyNodeExpirationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolENotifyNodeExpiration represents a ENotifyNodeExpiration event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolENotifyNodeExpiration struct {
	Caller  common.Address
	RaiseID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterENotifyNodeExpiration is a free log retrieval operation binding the contract event 0x551da46ed2c76a73ab003dedb33a3b4b280e742fc5947ab357306e7df729d57b.
//
// Solidity: event eNotifyNodeExpiration(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterENotifyNodeExpiration(opts *bind.FilterOpts) (*LetsFilRaisePoolENotifyNodeExpirationIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eNotifyNodeExpiration")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolENotifyNodeExpirationIterator{contract: _LetsFilRaisePool.contract, event: "eNotifyNodeExpiration", logs: logs, sub: sub}, nil
}

// WatchENotifyNodeExpiration is a free log subscription operation binding the contract event 0x551da46ed2c76a73ab003dedb33a3b4b280e742fc5947ab357306e7df729d57b.
//
// Solidity: event eNotifyNodeExpiration(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchENotifyNodeExpiration(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolENotifyNodeExpiration) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eNotifyNodeExpiration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolENotifyNodeExpiration)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eNotifyNodeExpiration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseENotifyNodeExpiration is a log parse operation binding the contract event 0x551da46ed2c76a73ab003dedb33a3b4b280e742fc5947ab357306e7df729d57b.
//
// Solidity: event eNotifyNodeExpiration(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseENotifyNodeExpiration(log types.Log) (*LetsFilRaisePoolENotifyNodeExpiration, error) {
	event := new(LetsFilRaisePoolENotifyNodeExpiration)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eNotifyNodeExpiration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEStackFromInvestorIterator is returned from FilterEStackFromInvestor and is used to iterate over the raw logs and unpacked data for EStackFromInvestor events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEStackFromInvestorIterator struct {
	Event *LetsFilRaisePoolEStackFromInvestor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEStackFromInvestorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEStackFromInvestor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEStackFromInvestor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEStackFromInvestorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEStackFromInvestorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEStackFromInvestor represents a EStackFromInvestor event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEStackFromInvestor struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEStackFromInvestor is a free log retrieval operation binding the contract event 0xa448f26ffdb95dd281bc422c93f89e8a7a08ec6f81165a98bc87651921b1acf3.
//
// Solidity: event eStackFromInvestor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEStackFromInvestor(opts *bind.FilterOpts) (*LetsFilRaisePoolEStackFromInvestorIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eStackFromInvestor")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEStackFromInvestorIterator{contract: _LetsFilRaisePool.contract, event: "eStackFromInvestor", logs: logs, sub: sub}, nil
}

// WatchEStackFromInvestor is a free log subscription operation binding the contract event 0xa448f26ffdb95dd281bc422c93f89e8a7a08ec6f81165a98bc87651921b1acf3.
//
// Solidity: event eStackFromInvestor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEStackFromInvestor(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEStackFromInvestor) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eStackFromInvestor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEStackFromInvestor)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eStackFromInvestor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEStackFromInvestor is a log parse operation binding the contract event 0xa448f26ffdb95dd281bc422c93f89e8a7a08ec6f81165a98bc87651921b1acf3.
//
// Solidity: event eStackFromInvestor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEStackFromInvestor(log types.Log) (*LetsFilRaisePoolEStackFromInvestor, error) {
	event := new(LetsFilRaisePoolEStackFromInvestor)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eStackFromInvestor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEStartRaisePlanIterator is returned from FilterEStartRaisePlan and is used to iterate over the raw logs and unpacked data for EStartRaisePlan events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEStartRaisePlanIterator struct {
	Event *LetsFilRaisePoolEStartRaisePlan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEStartRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEStartRaisePlan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEStartRaisePlan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEStartRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEStartRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEStartRaisePlan represents a EStartRaisePlan event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEStartRaisePlan struct {
	Caller  common.Address
	RaiseID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEStartRaisePlan is a free log retrieval operation binding the contract event 0x232677777bd30fff3912b9961ead28df75f4cc82ea5fae5822c9fc748872cbb5.
//
// Solidity: event eStartRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEStartRaisePlan(opts *bind.FilterOpts) (*LetsFilRaisePoolEStartRaisePlanIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eStartRaisePlan")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEStartRaisePlanIterator{contract: _LetsFilRaisePool.contract, event: "eStartRaisePlan", logs: logs, sub: sub}, nil
}

// WatchEStartRaisePlan is a free log subscription operation binding the contract event 0x232677777bd30fff3912b9961ead28df75f4cc82ea5fae5822c9fc748872cbb5.
//
// Solidity: event eStartRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEStartRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEStartRaisePlan) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eStartRaisePlan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEStartRaisePlan)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eStartRaisePlan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEStartRaisePlan is a log parse operation binding the contract event 0x232677777bd30fff3912b9961ead28df75f4cc82ea5fae5822c9fc748872cbb5.
//
// Solidity: event eStartRaisePlan(address caller, uint256 raiseID)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEStartRaisePlan(log types.Log) (*LetsFilRaisePoolEStartRaisePlan, error) {
	event := new(LetsFilRaisePoolEStartRaisePlan)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eStartRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEUnstackFromInverstorIterator is returned from FilterEUnstackFromInverstor and is used to iterate over the raw logs and unpacked data for EUnstackFromInverstor events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEUnstackFromInverstorIterator struct {
	Event *LetsFilRaisePoolEUnstackFromInverstor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEUnstackFromInverstorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEUnstackFromInverstor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEUnstackFromInverstor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEUnstackFromInverstorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEUnstackFromInverstorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEUnstackFromInverstor represents a EUnstackFromInverstor event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEUnstackFromInverstor struct {
	RaiseID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEUnstackFromInverstor is a free log retrieval operation binding the contract event 0x995b34d30ffd80c800180c5b8900838f152b473de74057131c44b85cf1cb8476.
//
// Solidity: event eUnstackFromInverstor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEUnstackFromInverstor(opts *bind.FilterOpts) (*LetsFilRaisePoolEUnstackFromInverstorIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eUnstackFromInverstor")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEUnstackFromInverstorIterator{contract: _LetsFilRaisePool.contract, event: "eUnstackFromInverstor", logs: logs, sub: sub}, nil
}

// WatchEUnstackFromInverstor is a free log subscription operation binding the contract event 0x995b34d30ffd80c800180c5b8900838f152b473de74057131c44b85cf1cb8476.
//
// Solidity: event eUnstackFromInverstor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEUnstackFromInverstor(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEUnstackFromInverstor) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eUnstackFromInverstor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEUnstackFromInverstor)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eUnstackFromInverstor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEUnstackFromInverstor is a log parse operation binding the contract event 0x995b34d30ffd80c800180c5b8900838f152b473de74057131c44b85cf1cb8476.
//
// Solidity: event eUnstackFromInverstor(uint256 raiseID, address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEUnstackFromInverstor(log types.Log) (*LetsFilRaisePoolEUnstackFromInverstor, error) {
	event := new(LetsFilRaisePoolEUnstackFromInverstor)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eUnstackFromInverstor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEWithdrawIterator is returned from FilterEWithdraw and is used to iterate over the raw logs and unpacked data for EWithdraw events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdrawIterator struct {
	Event *LetsFilRaisePoolEWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEWithdraw represents a EWithdraw event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdraw struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEWithdraw is a free log retrieval operation binding the contract event 0xc4e9fd0b0814b142a8dce2da2fb5bbbc63e23ecc9dc77fbdf7e6785b65821073.
//
// Solidity: event eWithdraw(address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEWithdraw(opts *bind.FilterOpts) (*LetsFilRaisePoolEWithdrawIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eWithdraw")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEWithdrawIterator{contract: _LetsFilRaisePool.contract, event: "eWithdraw", logs: logs, sub: sub}, nil
}

// WatchEWithdraw is a free log subscription operation binding the contract event 0xc4e9fd0b0814b142a8dce2da2fb5bbbc63e23ecc9dc77fbdf7e6785b65821073.
//
// Solidity: event eWithdraw(address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEWithdraw(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEWithdraw) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEWithdraw)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEWithdraw is a log parse operation binding the contract event 0xc4e9fd0b0814b142a8dce2da2fb5bbbc63e23ecc9dc77fbdf7e6785b65821073.
//
// Solidity: event eWithdraw(address from, address to, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEWithdraw(log types.Log) (*LetsFilRaisePoolEWithdraw, error) {
	event := new(LetsFilRaisePoolEWithdraw)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEWithdrawOPSSecurityFundIterator is returned from FilterEWithdrawOPSSecurityFund and is used to iterate over the raw logs and unpacked data for EWithdrawOPSSecurityFund events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdrawOPSSecurityFundIterator struct {
	Event *LetsFilRaisePoolEWithdrawOPSSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEWithdrawOPSSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEWithdrawOPSSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEWithdrawOPSSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEWithdrawOPSSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEWithdrawOPSSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEWithdrawOPSSecurityFund represents a EWithdrawOPSSecurityFund event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdrawOPSSecurityFund struct {
	Caller  common.Address
	RaiseID *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEWithdrawOPSSecurityFund is a free log retrieval operation binding the contract event 0x9f5878cc3fcae7b43b9f02949d5828c9b460972d7e2cd4d55cc7298850fd17ff.
//
// Solidity: event eWithdrawOPSSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEWithdrawOPSSecurityFund(opts *bind.FilterOpts) (*LetsFilRaisePoolEWithdrawOPSSecurityFundIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eWithdrawOPSSecurityFund")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEWithdrawOPSSecurityFundIterator{contract: _LetsFilRaisePool.contract, event: "eWithdrawOPSSecurityFund", logs: logs, sub: sub}, nil
}

// WatchEWithdrawOPSSecurityFund is a free log subscription operation binding the contract event 0x9f5878cc3fcae7b43b9f02949d5828c9b460972d7e2cd4d55cc7298850fd17ff.
//
// Solidity: event eWithdrawOPSSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEWithdrawOPSSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEWithdrawOPSSecurityFund) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eWithdrawOPSSecurityFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEWithdrawOPSSecurityFund)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdrawOPSSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEWithdrawOPSSecurityFund is a log parse operation binding the contract event 0x9f5878cc3fcae7b43b9f02949d5828c9b460972d7e2cd4d55cc7298850fd17ff.
//
// Solidity: event eWithdrawOPSSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEWithdrawOPSSecurityFund(log types.Log) (*LetsFilRaisePoolEWithdrawOPSSecurityFund, error) {
	event := new(LetsFilRaisePoolEWithdrawOPSSecurityFund)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdrawOPSSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator is returned from FilterEWithdrawRaiseSecurityFund and is used to iterate over the raw logs and unpacked data for EWithdrawRaiseSecurityFund events raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator struct {
	Event *LetsFilRaisePoolEWithdrawRaiseSecurityFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaisePoolEWithdrawRaiseSecurityFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LetsFilRaisePoolEWithdrawRaiseSecurityFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaisePoolEWithdrawRaiseSecurityFund represents a EWithdrawRaiseSecurityFund event raised by the LetsFilRaisePool contract.
type LetsFilRaisePoolEWithdrawRaiseSecurityFund struct {
	Caller  common.Address
	RaiseID *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEWithdrawRaiseSecurityFund is a free log retrieval operation binding the contract event 0xa3a6324388fc0c4f50a89eb79a6adccc6c13e8d02f33c7bae2aafdd73dad28db.
//
// Solidity: event eWithdrawRaiseSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) FilterEWithdrawRaiseSecurityFund(opts *bind.FilterOpts) (*LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator, error) {

	logs, sub, err := _LetsFilRaisePool.contract.FilterLogs(opts, "eWithdrawRaiseSecurityFund")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaisePoolEWithdrawRaiseSecurityFundIterator{contract: _LetsFilRaisePool.contract, event: "eWithdrawRaiseSecurityFund", logs: logs, sub: sub}, nil
}

// WatchEWithdrawRaiseSecurityFund is a free log subscription operation binding the contract event 0xa3a6324388fc0c4f50a89eb79a6adccc6c13e8d02f33c7bae2aafdd73dad28db.
//
// Solidity: event eWithdrawRaiseSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) WatchEWithdrawRaiseSecurityFund(opts *bind.WatchOpts, sink chan<- *LetsFilRaisePoolEWithdrawRaiseSecurityFund) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaisePool.contract.WatchLogs(opts, "eWithdrawRaiseSecurityFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaisePoolEWithdrawRaiseSecurityFund)
				if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdrawRaiseSecurityFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEWithdrawRaiseSecurityFund is a log parse operation binding the contract event 0xa3a6324388fc0c4f50a89eb79a6adccc6c13e8d02f33c7bae2aafdd73dad28db.
//
// Solidity: event eWithdrawRaiseSecurityFund(address caller, uint256 raiseID, uint256 amount)
func (_LetsFilRaisePool *LetsFilRaisePoolFilterer) ParseEWithdrawRaiseSecurityFund(log types.Log) (*LetsFilRaisePoolEWithdrawRaiseSecurityFund, error) {
	event := new(LetsFilRaisePoolEWithdrawRaiseSecurityFund)
	if err := _LetsFilRaisePool.contract.UnpackLog(event, "eWithdrawRaiseSecurityFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
