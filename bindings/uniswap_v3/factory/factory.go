// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package univ3factory

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
)

// Univ3factoryMetaData contains all meta data concerning the Univ3factory contract.
var Univ3factoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Univ3factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use Univ3factoryMetaData.ABI instead.
var Univ3factoryABI = Univ3factoryMetaData.ABI

// Univ3factory is an auto generated Go binding around an Ethereum contract.
type Univ3factory struct {
	Univ3factoryCaller     // Read-only binding to the contract
	Univ3factoryTransactor // Write-only binding to the contract
	Univ3factoryFilterer   // Log filterer for contract events
}

// Univ3factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Univ3factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ3factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Univ3factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ3factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Univ3factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Univ3factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Univ3factorySession struct {
	Contract     *Univ3factory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Univ3factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Univ3factoryCallerSession struct {
	Contract *Univ3factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Univ3factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Univ3factoryTransactorSession struct {
	Contract     *Univ3factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Univ3factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Univ3factoryRaw struct {
	Contract *Univ3factory // Generic contract binding to access the raw methods on
}

// Univ3factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Univ3factoryCallerRaw struct {
	Contract *Univ3factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Univ3factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Univ3factoryTransactorRaw struct {
	Contract *Univ3factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniv3factory creates a new instance of Univ3factory, bound to a specific deployed contract.
func NewUniv3factory(address common.Address, backend bind.ContractBackend) (*Univ3factory, error) {
	contract, err := bindUniv3factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Univ3factory{Univ3factoryCaller: Univ3factoryCaller{contract: contract}, Univ3factoryTransactor: Univ3factoryTransactor{contract: contract}, Univ3factoryFilterer: Univ3factoryFilterer{contract: contract}}, nil
}

// NewUniv3factoryCaller creates a new read-only instance of Univ3factory, bound to a specific deployed contract.
func NewUniv3factoryCaller(address common.Address, caller bind.ContractCaller) (*Univ3factoryCaller, error) {
	contract, err := bindUniv3factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Univ3factoryCaller{contract: contract}, nil
}

// NewUniv3factoryTransactor creates a new write-only instance of Univ3factory, bound to a specific deployed contract.
func NewUniv3factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Univ3factoryTransactor, error) {
	contract, err := bindUniv3factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Univ3factoryTransactor{contract: contract}, nil
}

// NewUniv3factoryFilterer creates a new log filterer instance of Univ3factory, bound to a specific deployed contract.
func NewUniv3factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Univ3factoryFilterer, error) {
	contract, err := bindUniv3factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Univ3factoryFilterer{contract: contract}, nil
}

// bindUniv3factory binds a generic wrapper to an already deployed contract.
func bindUniv3factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Univ3factoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Univ3factory *Univ3factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Univ3factory.Contract.Univ3factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Univ3factory *Univ3factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ3factory.Contract.Univ3factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Univ3factory *Univ3factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Univ3factory.Contract.Univ3factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Univ3factory *Univ3factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Univ3factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Univ3factory *Univ3factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Univ3factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Univ3factory *Univ3factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Univ3factory.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_Univ3factory *Univ3factoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Univ3factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_Univ3factory *Univ3factorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _Univ3factory.Contract.GetPool(&_Univ3factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_Univ3factory *Univ3factoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _Univ3factory.Contract.GetPool(&_Univ3factory.CallOpts, arg0, arg1, arg2)
}