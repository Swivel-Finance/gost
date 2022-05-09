// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PendleABI is the input ABI used to generate the binding from.
const PendleABI = "[{\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"expiryReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"yieldTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleBin is the compiled bytecode used for deploying new contracts.
var PendleBin = "0x608060405234801561001057600080fd5b506102c2806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806345c0684b1461005157806376d5de851461006d578063e184c9be1461008b578063ffd71860146100a9575b600080fd5b61006b600480360381019061006691906101a8565b6100c5565b005b610075610108565b60405161008291906101e4565b60405180910390f35b610093610131565b6040516100a09190610218565b60405180910390f35b6100c360048036038101906100be919061025f565b61013b565b005b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b8060018190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101758261014a565b9050919050565b6101858161016a565b811461019057600080fd5b50565b6000813590506101a28161017c565b92915050565b6000602082840312156101be576101bd610145565b5b60006101cc84828501610193565b91505092915050565b6101de8161016a565b82525050565b60006020820190506101f960008301846101d5565b92915050565b6000819050919050565b610212816101ff565b82525050565b600060208201905061022d6000830184610209565b92915050565b61023c816101ff565b811461024757600080fd5b50565b60008135905061025981610233565b92915050565b60006020828403121561027557610274610145565b5b60006102838482850161024a565b9150509291505056fea2646970667358221220cbffdf7effe8bf3c74a006d39e818dfe937a72a71e5253a149ac280420fd7b3064736f6c634300080d0033"

// DeployPendle deploys a new Ethereum contract, binding an instance of Pendle to it.
func DeployPendle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pendle, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pendle{PendleCaller: PendleCaller{contract: contract}, PendleTransactor: PendleTransactor{contract: contract}, PendleFilterer: PendleFilterer{contract: contract}}, nil
}

// Pendle is an auto generated Go binding around an Ethereum contract.
type Pendle struct {
	PendleCaller     // Read-only binding to the contract
	PendleTransactor // Write-only binding to the contract
	PendleFilterer   // Log filterer for contract events
}

// PendleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleSession struct {
	Contract     *Pendle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleCallerSession struct {
	Contract *PendleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PendleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleTransactorSession struct {
	Contract     *PendleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleRaw struct {
	Contract *Pendle // Generic contract binding to access the raw methods on
}

// PendleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleCallerRaw struct {
	Contract *PendleCaller // Generic read-only contract binding to access the raw methods on
}

// PendleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleTransactorRaw struct {
	Contract *PendleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendle creates a new instance of Pendle, bound to a specific deployed contract.
func NewPendle(address common.Address, backend bind.ContractBackend) (*Pendle, error) {
	contract, err := bindPendle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pendle{PendleCaller: PendleCaller{contract: contract}, PendleTransactor: PendleTransactor{contract: contract}, PendleFilterer: PendleFilterer{contract: contract}}, nil
}

// NewPendleCaller creates a new read-only instance of Pendle, bound to a specific deployed contract.
func NewPendleCaller(address common.Address, caller bind.ContractCaller) (*PendleCaller, error) {
	contract, err := bindPendle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleCaller{contract: contract}, nil
}

// NewPendleTransactor creates a new write-only instance of Pendle, bound to a specific deployed contract.
func NewPendleTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleTransactor, error) {
	contract, err := bindPendle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTransactor{contract: contract}, nil
}

// NewPendleFilterer creates a new log filterer instance of Pendle, bound to a specific deployed contract.
func NewPendleFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleFilterer, error) {
	contract, err := bindPendle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleFilterer{contract: contract}, nil
}

// bindPendle binds a generic wrapper to an already deployed contract.
func bindPendle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pendle *PendleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pendle.Contract.PendleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pendle *PendleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pendle.Contract.PendleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pendle *PendleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pendle.Contract.PendleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pendle *PendleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pendle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pendle *PendleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pendle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pendle *PendleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pendle.Contract.contract.Transact(opts, method, params...)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Pendle *PendleCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Pendle *PendleSession) Expiry() (*big.Int, error) {
	return _Pendle.Contract.Expiry(&_Pendle.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_Pendle *PendleCallerSession) Expiry() (*big.Int, error) {
	return _Pendle.Contract.Expiry(&_Pendle.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_Pendle *PendleCaller) YieldToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "yieldToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_Pendle *PendleSession) YieldToken() (common.Address, error) {
	return _Pendle.Contract.YieldToken(&_Pendle.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_Pendle *PendleCallerSession) YieldToken() (common.Address, error) {
	return _Pendle.Contract.YieldToken(&_Pendle.CallOpts)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_Pendle *PendleTransactor) ExpiryReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "expiryReturns", m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_Pendle *PendleSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.ExpiryReturns(&_Pendle.TransactOpts, m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_Pendle *PendleTransactorSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.ExpiryReturns(&_Pendle.TransactOpts, m)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_Pendle *PendleTransactor) YieldTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "yieldTokenReturns", a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_Pendle *PendleSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _Pendle.Contract.YieldTokenReturns(&_Pendle.TransactOpts, a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_Pendle *PendleTransactorSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _Pendle.Contract.YieldTokenReturns(&_Pendle.TransactOpts, a)
}
