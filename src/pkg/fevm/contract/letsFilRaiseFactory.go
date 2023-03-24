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

// ILetsFilRaiseInfoNodeInfo is an auto generated low-level Go binding around an user-defined struct.
type ILetsFilRaiseInfoNodeInfo struct {
	NodeSize             *big.Int
	SectorSize           *big.Int
	SealTimeLimit        *big.Int
	NodePeriod           *big.Int
	OpsSecurityFund      *big.Int
	OpsSecurityFundPayer common.Address
}

// ILetsFilRaiseInfoRaiseInfo is an auto generated low-level Go binding around an user-defined struct.
type ILetsFilRaiseInfoRaiseInfo struct {
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
}

// LetsFilRaiseFactoryMetaData contains all meta data concerning the LetsFilRaiseFactory contract.
var LetsFilRaiseFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"raisePool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sponsorCompany\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structILetsFilRaiseInfo.RaiseInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealTimeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"opsSecurityFundPayer\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structILetsFilRaiseInfo.NodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"eCreateRaisePlan\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RaiseID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"securityFund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raiserShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"investorShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"servicerShare\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"minerID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"sponsorCompany\",\"type\":\"string\"}],\"internalType\":\"structILetsFilRaiseInfo.RaiseInfo\",\"name\":\"raiseInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sectorSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sealTimeLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"opsSecurityFund\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"opsSecurityFundPayer\",\"type\":\"address\"}],\"internalType\":\"structILetsFilRaiseInfo.NodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"createRaisePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"raisePool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getRaisePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LetsFilRaiseFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LetsFilRaiseFactoryMetaData.ABI instead.
var LetsFilRaiseFactoryABI = LetsFilRaiseFactoryMetaData.ABI

// LetsFilRaiseFactory is an auto generated Go binding around an Ethereum contract.
type LetsFilRaiseFactory struct {
	LetsFilRaiseFactoryCaller     // Read-only binding to the contract
	LetsFilRaiseFactoryTransactor // Write-only binding to the contract
	LetsFilRaiseFactoryFilterer   // Log filterer for contract events
}

// LetsFilRaiseFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LetsFilRaiseFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaiseFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LetsFilRaiseFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaiseFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LetsFilRaiseFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LetsFilRaiseFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LetsFilRaiseFactorySession struct {
	Contract     *LetsFilRaiseFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LetsFilRaiseFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LetsFilRaiseFactoryCallerSession struct {
	Contract *LetsFilRaiseFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LetsFilRaiseFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LetsFilRaiseFactoryTransactorSession struct {
	Contract     *LetsFilRaiseFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LetsFilRaiseFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LetsFilRaiseFactoryRaw struct {
	Contract *LetsFilRaiseFactory // Generic contract binding to access the raw methods on
}

// LetsFilRaiseFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LetsFilRaiseFactoryCallerRaw struct {
	Contract *LetsFilRaiseFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LetsFilRaiseFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LetsFilRaiseFactoryTransactorRaw struct {
	Contract *LetsFilRaiseFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLetsFilRaiseFactory creates a new instance of LetsFilRaiseFactory, bound to a specific deployed contract.
func NewLetsFilRaiseFactory(address common.Address, backend bind.ContractBackend) (*LetsFilRaiseFactory, error) {
	contract, err := bindLetsFilRaiseFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaiseFactory{LetsFilRaiseFactoryCaller: LetsFilRaiseFactoryCaller{contract: contract}, LetsFilRaiseFactoryTransactor: LetsFilRaiseFactoryTransactor{contract: contract}, LetsFilRaiseFactoryFilterer: LetsFilRaiseFactoryFilterer{contract: contract}}, nil
}

// NewLetsFilRaiseFactoryCaller creates a new read-only instance of LetsFilRaiseFactory, bound to a specific deployed contract.
func NewLetsFilRaiseFactoryCaller(address common.Address, caller bind.ContractCaller) (*LetsFilRaiseFactoryCaller, error) {
	contract, err := bindLetsFilRaiseFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaiseFactoryCaller{contract: contract}, nil
}

// NewLetsFilRaiseFactoryTransactor creates a new write-only instance of LetsFilRaiseFactory, bound to a specific deployed contract.
func NewLetsFilRaiseFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LetsFilRaiseFactoryTransactor, error) {
	contract, err := bindLetsFilRaiseFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaiseFactoryTransactor{contract: contract}, nil
}

// NewLetsFilRaiseFactoryFilterer creates a new log filterer instance of LetsFilRaiseFactory, bound to a specific deployed contract.
func NewLetsFilRaiseFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LetsFilRaiseFactoryFilterer, error) {
	contract, err := bindLetsFilRaiseFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LetsFilRaiseFactoryFilterer{contract: contract}, nil
}

// bindLetsFilRaiseFactory binds a generic wrapper to an already deployed contract.
func bindLetsFilRaiseFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LetsFilRaiseFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilRaiseFactory.Contract.LetsFilRaiseFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.LetsFilRaiseFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.LetsFilRaiseFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LetsFilRaiseFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.contract.Transact(opts, method, params...)
}

// RaiseID is a free data retrieval call binding the contract method 0x525f5e8d.
//
// Solidity: function RaiseID() view returns(uint256)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryCaller) RaiseID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LetsFilRaiseFactory.contract.Call(opts, &out, "RaiseID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaiseID is a free data retrieval call binding the contract method 0x525f5e8d.
//
// Solidity: function RaiseID() view returns(uint256)
func (_LetsFilRaiseFactory *LetsFilRaiseFactorySession) RaiseID() (*big.Int, error) {
	return _LetsFilRaiseFactory.Contract.RaiseID(&_LetsFilRaiseFactory.CallOpts)
}

// RaiseID is a free data retrieval call binding the contract method 0x525f5e8d.
//
// Solidity: function RaiseID() view returns(uint256)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryCallerSession) RaiseID() (*big.Int, error) {
	return _LetsFilRaiseFactory.Contract.RaiseID(&_LetsFilRaiseFactory.CallOpts)
}

// GetRaisePool is a free data retrieval call binding the contract method 0x283ed076.
//
// Solidity: function getRaisePool(address , bytes32 , uint256 ) view returns(address)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryCaller) GetRaisePool(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LetsFilRaiseFactory.contract.Call(opts, &out, "getRaisePool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRaisePool is a free data retrieval call binding the contract method 0x283ed076.
//
// Solidity: function getRaisePool(address , bytes32 , uint256 ) view returns(address)
func (_LetsFilRaiseFactory *LetsFilRaiseFactorySession) GetRaisePool(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (common.Address, error) {
	return _LetsFilRaiseFactory.Contract.GetRaisePool(&_LetsFilRaiseFactory.CallOpts, arg0, arg1, arg2)
}

// GetRaisePool is a free data retrieval call binding the contract method 0x283ed076.
//
// Solidity: function getRaisePool(address , bytes32 , uint256 ) view returns(address)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryCallerSession) GetRaisePool(arg0 common.Address, arg1 [32]byte, arg2 *big.Int) (common.Address, error) {
	return _LetsFilRaiseFactory.Contract.GetRaisePool(&_LetsFilRaiseFactory.CallOpts, arg0, arg1, arg2)
}

// CreateRaisePool is a paid mutator transaction binding the contract method 0x58664232.
//
// Solidity: function createRaisePool((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) raiseInfo, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo) returns(address raisePool)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryTransactor) CreateRaisePool(opts *bind.TransactOpts, raiseInfo ILetsFilRaiseInfoRaiseInfo, nodeInfo ILetsFilRaiseInfoNodeInfo) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.contract.Transact(opts, "createRaisePool", raiseInfo, nodeInfo)
}

// CreateRaisePool is a paid mutator transaction binding the contract method 0x58664232.
//
// Solidity: function createRaisePool((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) raiseInfo, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo) returns(address raisePool)
func (_LetsFilRaiseFactory *LetsFilRaiseFactorySession) CreateRaisePool(raiseInfo ILetsFilRaiseInfoRaiseInfo, nodeInfo ILetsFilRaiseInfoNodeInfo) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.CreateRaisePool(&_LetsFilRaiseFactory.TransactOpts, raiseInfo, nodeInfo)
}

// CreateRaisePool is a paid mutator transaction binding the contract method 0x58664232.
//
// Solidity: function createRaisePool((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) raiseInfo, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo) returns(address raisePool)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryTransactorSession) CreateRaisePool(raiseInfo ILetsFilRaiseInfoRaiseInfo, nodeInfo ILetsFilRaiseInfoNodeInfo) (*types.Transaction, error) {
	return _LetsFilRaiseFactory.Contract.CreateRaisePool(&_LetsFilRaiseFactory.TransactOpts, raiseInfo, nodeInfo)
}

// LetsFilRaiseFactoryECreateRaisePlanIterator is returned from FilterECreateRaisePlan and is used to iterate over the raw logs and unpacked data for ECreateRaisePlan events raised by the LetsFilRaiseFactory contract.
type LetsFilRaiseFactoryECreateRaisePlanIterator struct {
	Event *LetsFilRaiseFactoryECreateRaisePlan // Event containing the contract specifics and raw log

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
func (it *LetsFilRaiseFactoryECreateRaisePlanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LetsFilRaiseFactoryECreateRaisePlan)
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
		it.Event = new(LetsFilRaiseFactoryECreateRaisePlan)
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
func (it *LetsFilRaiseFactoryECreateRaisePlanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LetsFilRaiseFactoryECreateRaisePlanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LetsFilRaiseFactoryECreateRaisePlan represents a ECreateRaisePlan event raised by the LetsFilRaiseFactory contract.
type LetsFilRaiseFactoryECreateRaisePlan struct {
	RaisePool common.Address
	Info      ILetsFilRaiseInfoRaiseInfo
	NodeInfo  ILetsFilRaiseInfoNodeInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterECreateRaisePlan is a free log retrieval operation binding the contract event 0xab15acefbdebc9ac98213e99f7b739d25be74f5169b2c4728c0163266759bd87.
//
// Solidity: event eCreateRaisePlan(address raisePool, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) info, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryFilterer) FilterECreateRaisePlan(opts *bind.FilterOpts) (*LetsFilRaiseFactoryECreateRaisePlanIterator, error) {

	logs, sub, err := _LetsFilRaiseFactory.contract.FilterLogs(opts, "eCreateRaisePlan")
	if err != nil {
		return nil, err
	}
	return &LetsFilRaiseFactoryECreateRaisePlanIterator{contract: _LetsFilRaiseFactory.contract, event: "eCreateRaisePlan", logs: logs, sub: sub}, nil
}

// WatchECreateRaisePlan is a free log subscription operation binding the contract event 0xab15acefbdebc9ac98213e99f7b739d25be74f5169b2c4728c0163266759bd87.
//
// Solidity: event eCreateRaisePlan(address raisePool, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) info, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryFilterer) WatchECreateRaisePlan(opts *bind.WatchOpts, sink chan<- *LetsFilRaiseFactoryECreateRaisePlan) (event.Subscription, error) {

	logs, sub, err := _LetsFilRaiseFactory.contract.WatchLogs(opts, "eCreateRaisePlan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LetsFilRaiseFactoryECreateRaisePlan)
				if err := _LetsFilRaiseFactory.contract.UnpackLog(event, "eCreateRaisePlan", log); err != nil {
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

// ParseECreateRaisePlan is a log parse operation binding the contract event 0xab15acefbdebc9ac98213e99f7b739d25be74f5169b2c4728c0163266759bd87.
//
// Solidity: event eCreateRaisePlan(address raisePool, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bytes32,string) info, (uint256,uint256,uint256,uint256,uint256,address) nodeInfo)
func (_LetsFilRaiseFactory *LetsFilRaiseFactoryFilterer) ParseECreateRaisePlan(log types.Log) (*LetsFilRaiseFactoryECreateRaisePlan, error) {
	event := new(LetsFilRaiseFactoryECreateRaisePlan)
	if err := _LetsFilRaiseFactory.contract.UnpackLog(event, "eCreateRaisePlan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
