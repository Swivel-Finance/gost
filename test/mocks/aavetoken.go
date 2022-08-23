// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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

// AaveTokenMetaData contains all meta data concerning the AaveToken contract.
var AaveTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"POOLReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"UNDERLYINGReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061028f806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633f419ebb146100515780637535d2461461006d578063b16a19de1461008b578063f136cc7b146100a9575b600080fd5b61006b60048036038101906100669190610202565b6100c5565b005b610075610109565b604051610082919061023e565b60405180910390f35b610093610132565b6040516100a0919061023e565b60405180910390f35b6100c360048036038101906100be9190610202565b61015c565b005b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101cf826101a4565b9050919050565b6101df816101c4565b81146101ea57600080fd5b50565b6000813590506101fc816101d6565b92915050565b6000602082840312156102185761021761019f565b5b6000610226848285016101ed565b91505092915050565b610238816101c4565b82525050565b6000602082019050610253600083018461022f565b9291505056fea264697066735822122025f50d9a4e64c7d9feb4027a25cbfa46853e36a588cb0d44a7e229b8e975eef364736f6c63430008100033",
}

// AaveTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveTokenMetaData.ABI instead.
var AaveTokenABI = AaveTokenMetaData.ABI

// AaveTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AaveTokenMetaData.Bin instead.
var AaveTokenBin = AaveTokenMetaData.Bin

// DeployAaveToken deploys a new Ethereum contract, binding an instance of AaveToken to it.
func DeployAaveToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AaveToken, error) {
	parsed, err := AaveTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AaveTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AaveToken{AaveTokenCaller: AaveTokenCaller{contract: contract}, AaveTokenTransactor: AaveTokenTransactor{contract: contract}, AaveTokenFilterer: AaveTokenFilterer{contract: contract}}, nil
}

// AaveToken is an auto generated Go binding around an Ethereum contract.
type AaveToken struct {
	AaveTokenCaller     // Read-only binding to the contract
	AaveTokenTransactor // Write-only binding to the contract
	AaveTokenFilterer   // Log filterer for contract events
}

// AaveTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveTokenSession struct {
	Contract     *AaveToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveTokenCallerSession struct {
	Contract *AaveTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AaveTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTokenTransactorSession struct {
	Contract     *AaveTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AaveTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveTokenRaw struct {
	Contract *AaveToken // Generic contract binding to access the raw methods on
}

// AaveTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveTokenCallerRaw struct {
	Contract *AaveTokenCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTokenTransactorRaw struct {
	Contract *AaveTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveToken creates a new instance of AaveToken, bound to a specific deployed contract.
func NewAaveToken(address common.Address, backend bind.ContractBackend) (*AaveToken, error) {
	contract, err := bindAaveToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveToken{AaveTokenCaller: AaveTokenCaller{contract: contract}, AaveTokenTransactor: AaveTokenTransactor{contract: contract}, AaveTokenFilterer: AaveTokenFilterer{contract: contract}}, nil
}

// NewAaveTokenCaller creates a new read-only instance of AaveToken, bound to a specific deployed contract.
func NewAaveTokenCaller(address common.Address, caller bind.ContractCaller) (*AaveTokenCaller, error) {
	contract, err := bindAaveToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTokenCaller{contract: contract}, nil
}

// NewAaveTokenTransactor creates a new write-only instance of AaveToken, bound to a specific deployed contract.
func NewAaveTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTokenTransactor, error) {
	contract, err := bindAaveToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTokenTransactor{contract: contract}, nil
}

// NewAaveTokenFilterer creates a new log filterer instance of AaveToken, bound to a specific deployed contract.
func NewAaveTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveTokenFilterer, error) {
	contract, err := bindAaveToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveTokenFilterer{contract: contract}, nil
}

// bindAaveToken binds a generic wrapper to an already deployed contract.
func bindAaveToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveToken *AaveTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveToken.Contract.AaveTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveToken *AaveTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveToken.Contract.AaveTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveToken *AaveTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveToken.Contract.AaveTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveToken *AaveTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveToken *AaveTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveToken *AaveTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveToken.Contract.contract.Transact(opts, method, params...)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveToken *AaveTokenCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveToken.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveToken *AaveTokenSession) POOL() (common.Address, error) {
	return _AaveToken.Contract.POOL(&_AaveToken.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveToken *AaveTokenCallerSession) POOL() (common.Address, error) {
	return _AaveToken.Contract.POOL(&_AaveToken.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveToken *AaveTokenCaller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveToken.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveToken *AaveTokenSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveToken.Contract.UNDERLYINGASSETADDRESS(&_AaveToken.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveToken *AaveTokenCallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveToken.Contract.UNDERLYINGASSETADDRESS(&_AaveToken.CallOpts)
}

// POOLReturns is a paid mutator transaction binding the contract method 0xf136cc7b.
//
// Solidity: function POOLReturns(address a) returns()
func (_AaveToken *AaveTokenTransactor) POOLReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _AaveToken.contract.Transact(opts, "POOLReturns", a)
}

// POOLReturns is a paid mutator transaction binding the contract method 0xf136cc7b.
//
// Solidity: function POOLReturns(address a) returns()
func (_AaveToken *AaveTokenSession) POOLReturns(a common.Address) (*types.Transaction, error) {
	return _AaveToken.Contract.POOLReturns(&_AaveToken.TransactOpts, a)
}

// POOLReturns is a paid mutator transaction binding the contract method 0xf136cc7b.
//
// Solidity: function POOLReturns(address a) returns()
func (_AaveToken *AaveTokenTransactorSession) POOLReturns(a common.Address) (*types.Transaction, error) {
	return _AaveToken.Contract.POOLReturns(&_AaveToken.TransactOpts, a)
}

// UNDERLYINGReturns is a paid mutator transaction binding the contract method 0x3f419ebb.
//
// Solidity: function UNDERLYINGReturns(address a) returns()
func (_AaveToken *AaveTokenTransactor) UNDERLYINGReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _AaveToken.contract.Transact(opts, "UNDERLYINGReturns", a)
}

// UNDERLYINGReturns is a paid mutator transaction binding the contract method 0x3f419ebb.
//
// Solidity: function UNDERLYINGReturns(address a) returns()
func (_AaveToken *AaveTokenSession) UNDERLYINGReturns(a common.Address) (*types.Transaction, error) {
	return _AaveToken.Contract.UNDERLYINGReturns(&_AaveToken.TransactOpts, a)
}

// UNDERLYINGReturns is a paid mutator transaction binding the contract method 0x3f419ebb.
//
// Solidity: function UNDERLYINGReturns(address a) returns()
func (_AaveToken *AaveTokenTransactorSession) UNDERLYINGReturns(a common.Address) (*types.Transaction, error) {
	return _AaveToken.Contract.UNDERLYINGReturns(&_AaveToken.TransactOpts, a)
}
