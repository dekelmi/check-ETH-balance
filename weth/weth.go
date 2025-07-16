// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weth

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

// WethMetaData contains all meta data concerning the Weth contract.
var WethMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"type\":\"function\"}]",
}

// WethABI is the input ABI used to generate the binding from.
// Deprecated: Use WethMetaData.ABI instead.
var WethABI = WethMetaData.ABI

// Weth is an auto generated Go binding around an Ethereum contract.
type Weth struct {
	WethCaller     // Read-only binding to the contract
	WethTransactor // Write-only binding to the contract
	WethFilterer   // Log filterer for contract events
}

// WethCaller is an auto generated read-only Go binding around an Ethereum contract.
type WethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WethSession struct {
	Contract     *Weth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WethCallerSession struct {
	Contract *WethCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WethTransactorSession struct {
	Contract     *WethTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WethRaw is an auto generated low-level Go binding around an Ethereum contract.
type WethRaw struct {
	Contract *Weth // Generic contract binding to access the raw methods on
}

// WethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WethCallerRaw struct {
	Contract *WethCaller // Generic read-only contract binding to access the raw methods on
}

// WethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WethTransactorRaw struct {
	Contract *WethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeth creates a new instance of Weth, bound to a specific deployed contract.
func NewWeth(address common.Address, backend bind.ContractBackend) (*Weth, error) {
	contract, err := bindWeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Weth{WethCaller: WethCaller{contract: contract}, WethTransactor: WethTransactor{contract: contract}, WethFilterer: WethFilterer{contract: contract}}, nil
}

// NewWethCaller creates a new read-only instance of Weth, bound to a specific deployed contract.
func NewWethCaller(address common.Address, caller bind.ContractCaller) (*WethCaller, error) {
	contract, err := bindWeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WethCaller{contract: contract}, nil
}

// NewWethTransactor creates a new write-only instance of Weth, bound to a specific deployed contract.
func NewWethTransactor(address common.Address, transactor bind.ContractTransactor) (*WethTransactor, error) {
	contract, err := bindWeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WethTransactor{contract: contract}, nil
}

// NewWethFilterer creates a new log filterer instance of Weth, bound to a specific deployed contract.
func NewWethFilterer(address common.Address, filterer bind.ContractFilterer) (*WethFilterer, error) {
	contract, err := bindWeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WethFilterer{contract: contract}, nil
}

// bindWeth binds a generic wrapper to an already deployed contract.
func bindWeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WethMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Weth *WethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Weth.Contract.WethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Weth *WethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Weth.Contract.WethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Weth *WethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Weth.Contract.WethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Weth *WethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Weth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Weth *WethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Weth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Weth *WethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Weth.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Weth *WethCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Weth.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Weth *WethSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Weth.Contract.BalanceOf(&_Weth.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 balance)
func (_Weth *WethCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Weth.Contract.BalanceOf(&_Weth.CallOpts, _owner)
}
