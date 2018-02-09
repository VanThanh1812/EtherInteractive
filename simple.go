// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SimpleABI is the input ABI used to generate the binding from.
const SimpleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"multiply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"uint256\"},{\"name\":\"_b\",\"type\":\"uint256\"}],\"name\":\"arith\",\"outputs\":[{\"name\":\"o_sum\",\"type\":\"uint256\"},{\"name\":\"o_product\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SimpleBin is the compiled bytecode used for deploying new contracts.
const SimpleBin = `0x6060604052341561000f57600080fd5b60dc8061001d6000396000f30060606040526004361060485763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663165c4a168114604d578063d03deeac146075575b600080fd5b3415605757600080fd5b606360043560243560a3565b60405190815260200160405180910390f35b3415607f57600080fd5b608b60043560243560a7565b60405191825260208201526040908101905180910390f35b0290565b818101929102905600a165627a7a723058209fc4633ac77b74beb19f7c5b3f6f14028cd20aeda1e9dd7b529b9954d81279d90029`

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// Simple is an auto generated Go binding around an Ethereum contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// Arith is a paid mutator transaction binding the contract method 0xd03deeac.
//
// Solidity: function arith(_a uint256, _b uint256) returns(o_sum uint256, o_product uint256)
func (_Simple *SimpleTransactor) Arith(opts *bind.TransactOpts, _a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "arith", _a, _b)
}

// Arith is a paid mutator transaction binding the contract method 0xd03deeac.
//
// Solidity: function arith(_a uint256, _b uint256) returns(o_sum uint256, o_product uint256)
func (_Simple *SimpleSession) Arith(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Arith(&_Simple.TransactOpts, _a, _b)
}

// Arith is a paid mutator transaction binding the contract method 0xd03deeac.
//
// Solidity: function arith(_a uint256, _b uint256) returns(o_sum uint256, o_product uint256)
func (_Simple *SimpleTransactorSession) Arith(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Arith(&_Simple.TransactOpts, _a, _b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(_a uint256, _b uint256) returns(uint256)
func (_Simple *SimpleTransactor) Multiply(opts *bind.TransactOpts, _a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "multiply", _a, _b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(_a uint256, _b uint256) returns(uint256)
func (_Simple *SimpleSession) Multiply(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Multiply(&_Simple.TransactOpts, _a, _b)
}

// Multiply is a paid mutator transaction binding the contract method 0x165c4a16.
//
// Solidity: function multiply(_a uint256, _b uint256) returns(uint256)
func (_Simple *SimpleTransactorSession) Multiply(_a *big.Int, _b *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.Multiply(&_Simple.TransactOpts, _a, _b)
}
