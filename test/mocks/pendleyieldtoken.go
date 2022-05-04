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

// PendleYieldTokenABI is the input ABI used to generate the binding from.
const PendleYieldTokenABI = "[{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingAssetReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingYieldToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingYieldTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleYieldTokenBin is the compiled bytecode used for deploying new contracts.
var PendleYieldTokenBin = "0x608060405234801561001057600080fd5b5061028f806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634167cc2b146100515780637158da7c1461006d578063873c92be1461008b578063ddf0fa83146100a7575b600080fd5b61006b60048036038101906100669190610202565b6100c5565b005b610075610108565b604051610082919061023e565b60405180910390f35b6100a560048036038101906100a09190610202565b610131565b005b6100af610175565b6040516100bc919061023e565b60405180910390f35b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101cf826101a4565b9050919050565b6101df816101c4565b81146101ea57600080fd5b50565b6000813590506101fc816101d6565b92915050565b6000602082840312156102185761021761019f565b5b6000610226848285016101ed565b91505092915050565b610238816101c4565b82525050565b6000602082019050610253600083018461022f565b9291505056fea264697066735822122086a78bc939d23a95d4259c99ceaa145cb62451405e7f0c0cb9e6f5371d60f90e64736f6c634300080d0033"

// DeployPendleYieldToken deploys a new Ethereum contract, binding an instance of PendleYieldToken to it.
func DeployPendleYieldToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleYieldToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleYieldTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleYieldTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleYieldToken{PendleYieldTokenCaller: PendleYieldTokenCaller{contract: contract}, PendleYieldTokenTransactor: PendleYieldTokenTransactor{contract: contract}, PendleYieldTokenFilterer: PendleYieldTokenFilterer{contract: contract}}, nil
}

// PendleYieldToken is an auto generated Go binding around an Ethereum contract.
type PendleYieldToken struct {
	PendleYieldTokenCaller     // Read-only binding to the contract
	PendleYieldTokenTransactor // Write-only binding to the contract
	PendleYieldTokenFilterer   // Log filterer for contract events
}

// PendleYieldTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleYieldTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleYieldTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleYieldTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleYieldTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleYieldTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleYieldTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleYieldTokenSession struct {
	Contract     *PendleYieldToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleYieldTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleYieldTokenCallerSession struct {
	Contract *PendleYieldTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PendleYieldTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleYieldTokenTransactorSession struct {
	Contract     *PendleYieldTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PendleYieldTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleYieldTokenRaw struct {
	Contract *PendleYieldToken // Generic contract binding to access the raw methods on
}

// PendleYieldTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleYieldTokenCallerRaw struct {
	Contract *PendleYieldTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PendleYieldTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleYieldTokenTransactorRaw struct {
	Contract *PendleYieldTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleYieldToken creates a new instance of PendleYieldToken, bound to a specific deployed contract.
func NewPendleYieldToken(address common.Address, backend bind.ContractBackend) (*PendleYieldToken, error) {
	contract, err := bindPendleYieldToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleYieldToken{PendleYieldTokenCaller: PendleYieldTokenCaller{contract: contract}, PendleYieldTokenTransactor: PendleYieldTokenTransactor{contract: contract}, PendleYieldTokenFilterer: PendleYieldTokenFilterer{contract: contract}}, nil
}

// NewPendleYieldTokenCaller creates a new read-only instance of PendleYieldToken, bound to a specific deployed contract.
func NewPendleYieldTokenCaller(address common.Address, caller bind.ContractCaller) (*PendleYieldTokenCaller, error) {
	contract, err := bindPendleYieldToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleYieldTokenCaller{contract: contract}, nil
}

// NewPendleYieldTokenTransactor creates a new write-only instance of PendleYieldToken, bound to a specific deployed contract.
func NewPendleYieldTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleYieldTokenTransactor, error) {
	contract, err := bindPendleYieldToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleYieldTokenTransactor{contract: contract}, nil
}

// NewPendleYieldTokenFilterer creates a new log filterer instance of PendleYieldToken, bound to a specific deployed contract.
func NewPendleYieldTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleYieldTokenFilterer, error) {
	contract, err := bindPendleYieldToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleYieldTokenFilterer{contract: contract}, nil
}

// bindPendleYieldToken binds a generic wrapper to an already deployed contract.
func bindPendleYieldToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleYieldTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleYieldToken *PendleYieldTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleYieldToken.Contract.PendleYieldTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleYieldToken *PendleYieldTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.PendleYieldTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleYieldToken *PendleYieldTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.PendleYieldTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleYieldToken *PendleYieldTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleYieldToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleYieldToken *PendleYieldTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleYieldToken *PendleYieldTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.contract.Transact(opts, method, params...)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_PendleYieldToken *PendleYieldTokenCaller) UnderlyingAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleYieldToken.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_PendleYieldToken *PendleYieldTokenSession) UnderlyingAsset() (common.Address, error) {
	return _PendleYieldToken.Contract.UnderlyingAsset(&_PendleYieldToken.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_PendleYieldToken *PendleYieldTokenCallerSession) UnderlyingAsset() (common.Address, error) {
	return _PendleYieldToken.Contract.UnderlyingAsset(&_PendleYieldToken.CallOpts)
}

// UnderlyingYieldToken is a free data retrieval call binding the contract method 0xddf0fa83.
//
// Solidity: function underlyingYieldToken() view returns(address)
func (_PendleYieldToken *PendleYieldTokenCaller) UnderlyingYieldToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleYieldToken.contract.Call(opts, &out, "underlyingYieldToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingYieldToken is a free data retrieval call binding the contract method 0xddf0fa83.
//
// Solidity: function underlyingYieldToken() view returns(address)
func (_PendleYieldToken *PendleYieldTokenSession) UnderlyingYieldToken() (common.Address, error) {
	return _PendleYieldToken.Contract.UnderlyingYieldToken(&_PendleYieldToken.CallOpts)
}

// UnderlyingYieldToken is a free data retrieval call binding the contract method 0xddf0fa83.
//
// Solidity: function underlyingYieldToken() view returns(address)
func (_PendleYieldToken *PendleYieldTokenCallerSession) UnderlyingYieldToken() (common.Address, error) {
	return _PendleYieldToken.Contract.UnderlyingYieldToken(&_PendleYieldToken.CallOpts)
}

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenTransactor) UnderlyingAssetReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.contract.Transact(opts, "underlyingAssetReturns", a)
}

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenSession) UnderlyingAssetReturns(a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.UnderlyingAssetReturns(&_PendleYieldToken.TransactOpts, a)
}

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenTransactorSession) UnderlyingAssetReturns(a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.UnderlyingAssetReturns(&_PendleYieldToken.TransactOpts, a)
}

// UnderlyingYieldTokenReturns is a paid mutator transaction binding the contract method 0x873c92be.
//
// Solidity: function underlyingYieldTokenReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenTransactor) UnderlyingYieldTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.contract.Transact(opts, "underlyingYieldTokenReturns", a)
}

// UnderlyingYieldTokenReturns is a paid mutator transaction binding the contract method 0x873c92be.
//
// Solidity: function underlyingYieldTokenReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenSession) UnderlyingYieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.UnderlyingYieldTokenReturns(&_PendleYieldToken.TransactOpts, a)
}

// UnderlyingYieldTokenReturns is a paid mutator transaction binding the contract method 0x873c92be.
//
// Solidity: function underlyingYieldTokenReturns(address a) returns()
func (_PendleYieldToken *PendleYieldTokenTransactorSession) UnderlyingYieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleYieldToken.Contract.UnderlyingYieldTokenReturns(&_PendleYieldToken.TransactOpts, a)
}
