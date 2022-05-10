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

// ElementTokenABI is the input ABI used to generate the binding from.
const ElementTokenABI = "[{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"unlockTimestampReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ElementTokenBin is the compiled bytecode used for deploying new contracts.
var ElementTokenBin = "0x608060405234801561001057600080fd5b506102c3806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80636f307dc314610051578063831e574f1461006f578063aa082a9d1461008b578063e7ba6774146100a9575b600080fd5b6100596100c5565b6040516100669190610187565b60405180910390f35b610089600480360381019061008491906101dd565b6100ef565b005b6100936100f9565b6040516100a09190610219565b60405180910390f35b6100c360048036038101906100be9190610260565b610102565b005b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060008190555050565b60008054905090565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061017182610146565b9050919050565b61018181610166565b82525050565b600060208201905061019c6000830184610178565b92915050565b600080fd5b6000819050919050565b6101ba816101a7565b81146101c557600080fd5b50565b6000813590506101d7816101b1565b92915050565b6000602082840312156101f3576101f26101a2565b5b6000610201848285016101c8565b91505092915050565b610213816101a7565b82525050565b600060208201905061022e600083018461020a565b92915050565b61023d81610166565b811461024857600080fd5b50565b60008135905061025a81610234565b92915050565b600060208284031215610276576102756101a2565b5b60006102848482850161024b565b9150509291505056fea26469706673582212202bb534e08777966ba1178acf741e9d9581331b35cd46bb47ff070b11e90bb9b564736f6c634300080d0033"

// DeployElementToken deploys a new Ethereum contract, binding an instance of ElementToken to it.
func DeployElementToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ElementToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElementTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElementToken{ElementTokenCaller: ElementTokenCaller{contract: contract}, ElementTokenTransactor: ElementTokenTransactor{contract: contract}, ElementTokenFilterer: ElementTokenFilterer{contract: contract}}, nil
}

// ElementToken is an auto generated Go binding around an Ethereum contract.
type ElementToken struct {
	ElementTokenCaller     // Read-only binding to the contract
	ElementTokenTransactor // Write-only binding to the contract
	ElementTokenFilterer   // Log filterer for contract events
}

// ElementTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElementTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElementTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElementTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElementTokenSession struct {
	Contract     *ElementToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElementTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElementTokenCallerSession struct {
	Contract *ElementTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ElementTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElementTokenTransactorSession struct {
	Contract     *ElementTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ElementTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElementTokenRaw struct {
	Contract *ElementToken // Generic contract binding to access the raw methods on
}

// ElementTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElementTokenCallerRaw struct {
	Contract *ElementTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ElementTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElementTokenTransactorRaw struct {
	Contract *ElementTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElementToken creates a new instance of ElementToken, bound to a specific deployed contract.
func NewElementToken(address common.Address, backend bind.ContractBackend) (*ElementToken, error) {
	contract, err := bindElementToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElementToken{ElementTokenCaller: ElementTokenCaller{contract: contract}, ElementTokenTransactor: ElementTokenTransactor{contract: contract}, ElementTokenFilterer: ElementTokenFilterer{contract: contract}}, nil
}

// NewElementTokenCaller creates a new read-only instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenCaller(address common.Address, caller bind.ContractCaller) (*ElementTokenCaller, error) {
	contract, err := bindElementToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTokenCaller{contract: contract}, nil
}

// NewElementTokenTransactor creates a new write-only instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ElementTokenTransactor, error) {
	contract, err := bindElementToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTokenTransactor{contract: contract}, nil
}

// NewElementTokenFilterer creates a new log filterer instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ElementTokenFilterer, error) {
	contract, err := bindElementToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElementTokenFilterer{contract: contract}, nil
}

// bindElementToken binds a generic wrapper to an already deployed contract.
func bindElementToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElementToken *ElementTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElementToken.Contract.ElementTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElementToken *ElementTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElementToken.Contract.ElementTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElementToken *ElementTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElementToken.Contract.ElementTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElementToken *ElementTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElementToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElementToken *ElementTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElementToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElementToken *ElementTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElementToken.Contract.contract.Transact(opts, method, params...)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenSession) Underlying() (common.Address, error) {
	return _ElementToken.Contract.Underlying(&_ElementToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenCallerSession) Underlying() (common.Address, error) {
	return _ElementToken.Contract.Underlying(&_ElementToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenCaller) UnlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenSession) UnlockTimestamp() (*big.Int, error) {
	return _ElementToken.Contract.UnlockTimestamp(&_ElementToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenCallerSession) UnlockTimestamp() (*big.Int, error) {
	return _ElementToken.Contract.UnlockTimestamp(&_ElementToken.CallOpts)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.UnderlyingReturns(&_ElementToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.UnderlyingReturns(&_ElementToken.TransactOpts, a)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenTransactor) UnlockTimestampReturns(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "unlockTimestampReturns", u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.UnlockTimestampReturns(&_ElementToken.TransactOpts, u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenTransactorSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.UnlockTimestampReturns(&_ElementToken.TransactOpts, u)
}
