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

// PendleTokenMetaData contains all meta data concerning the PendleToken contract.
var PendleTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"expiryReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"yieldTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102c2806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806345c0684b1461005157806376d5de851461006d578063e184c9be1461008b578063ffd71860146100a9575b600080fd5b61006b600480360381019061006691906101a8565b6100c5565b005b610075610108565b60405161008291906101e4565b60405180910390f35b610093610131565b6040516100a09190610218565b60405180910390f35b6100c360048036038101906100be919061025f565b61013b565b005b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b8060018190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101758261014a565b9050919050565b6101858161016a565b811461019057600080fd5b50565b6000813590506101a28161017c565b92915050565b6000602082840312156101be576101bd610145565b5b60006101cc84828501610193565b91505092915050565b6101de8161016a565b82525050565b60006020820190506101f960008301846101d5565b92915050565b6000819050919050565b610212816101ff565b82525050565b600060208201905061022d6000830184610209565b92915050565b61023c816101ff565b811461024757600080fd5b50565b60008135905061025981610233565b92915050565b60006020828403121561027557610274610145565b5b60006102838482850161024a565b9150509291505056fea26469706673582212202d8ed28b79e7e90f5ba1158e590a9d75268a42dafc7bf6988fe7c40ec4f668d364736f6c634300080d0033",
}

// PendleTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PendleTokenMetaData.ABI instead.
var PendleTokenABI = PendleTokenMetaData.ABI

// PendleTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PendleTokenMetaData.Bin instead.
var PendleTokenBin = PendleTokenMetaData.Bin

// DeployPendleToken deploys a new Ethereum contract, binding an instance of PendleToken to it.
func DeployPendleToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleToken, error) {
	parsed, err := PendleTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PendleTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleToken{PendleTokenCaller: PendleTokenCaller{contract: contract}, PendleTokenTransactor: PendleTokenTransactor{contract: contract}, PendleTokenFilterer: PendleTokenFilterer{contract: contract}}, nil
}

// PendleToken is an auto generated Go binding around an Ethereum contract.
type PendleToken struct {
	PendleTokenCaller     // Read-only binding to the contract
	PendleTokenTransactor // Write-only binding to the contract
	PendleTokenFilterer   // Log filterer for contract events
}

// PendleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleTokenSession struct {
	Contract     *PendleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleTokenCallerSession struct {
	Contract *PendleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PendleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleTokenTransactorSession struct {
	Contract     *PendleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PendleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleTokenRaw struct {
	Contract *PendleToken // Generic contract binding to access the raw methods on
}

// PendleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleTokenCallerRaw struct {
	Contract *PendleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PendleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleTokenTransactorRaw struct {
	Contract *PendleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleToken creates a new instance of PendleToken, bound to a specific deployed contract.
func NewPendleToken(address common.Address, backend bind.ContractBackend) (*PendleToken, error) {
	contract, err := bindPendleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleToken{PendleTokenCaller: PendleTokenCaller{contract: contract}, PendleTokenTransactor: PendleTokenTransactor{contract: contract}, PendleTokenFilterer: PendleTokenFilterer{contract: contract}}, nil
}

// NewPendleTokenCaller creates a new read-only instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenCaller(address common.Address, caller bind.ContractCaller) (*PendleTokenCaller, error) {
	contract, err := bindPendleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTokenCaller{contract: contract}, nil
}

// NewPendleTokenTransactor creates a new write-only instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleTokenTransactor, error) {
	contract, err := bindPendleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTokenTransactor{contract: contract}, nil
}

// NewPendleTokenFilterer creates a new log filterer instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleTokenFilterer, error) {
	contract, err := bindPendleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleTokenFilterer{contract: contract}, nil
}

// bindPendleToken binds a generic wrapper to an already deployed contract.
func bindPendleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleToken *PendleTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleToken.Contract.PendleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleToken *PendleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleToken.Contract.PendleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleToken *PendleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleToken.Contract.PendleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleToken *PendleTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleToken *PendleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleToken *PendleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleToken.Contract.contract.Transact(opts, method, params...)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenSession) Expiry() (*big.Int, error) {
	return _PendleToken.Contract.Expiry(&_PendleToken.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenCallerSession) Expiry() (*big.Int, error) {
	return _PendleToken.Contract.Expiry(&_PendleToken.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenCaller) YieldToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "yieldToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenSession) YieldToken() (common.Address, error) {
	return _PendleToken.Contract.YieldToken(&_PendleToken.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenCallerSession) YieldToken() (common.Address, error) {
	return _PendleToken.Contract.YieldToken(&_PendleToken.CallOpts)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenTransactor) ExpiryReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "expiryReturns", m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.ExpiryReturns(&_PendleToken.TransactOpts, m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenTransactorSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.ExpiryReturns(&_PendleToken.TransactOpts, m)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenTransactor) YieldTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "yieldTokenReturns", a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.YieldTokenReturns(&_PendleToken.TransactOpts, a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenTransactorSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.YieldTokenReturns(&_PendleToken.TransactOpts, a)
}
