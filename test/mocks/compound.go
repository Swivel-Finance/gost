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

// CompoundMetaData contains all meta data concerning the Compound contract.
var CompoundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PROTOCOL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548160ff021916908360ff16021790555034801561002b57600080fd5b5061048d8061003b6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063dc3b7c8b1161005b578063dc3b7c8b146100da578063e6919c821461010a578063e7ba67741461013a578063f189b337146101565761007d565b8063589aeec7146100825780635c62eb2f1461009e57806391b9b827146100bc575b600080fd5b61009c60048036038101906100979190610314565b610174565b005b6100a661017e565b6040516100b39190610382565b60405180910390f35b6100c46101a4565b6040516100d191906103b9565b60405180910390f35b6100f460048036038101906100ef9190610400565b6101b5565b604051610101919061043c565b60405180910390f35b610124600480360381019061011f9190610400565b610202565b6040516101319190610382565b60405180910390f35b610154600480360381019061014f9190610400565b61026f565b005b61015e6102b3565b60405161016b9190610382565b60405180910390f35b8060038190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900460ff1681565b600081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003549050919050565b600081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080fd5b6000819050919050565b6102f1816102de565b81146102fc57600080fd5b50565b60008135905061030e816102e8565b92915050565b60006020828403121561032a576103296102d9565b5b6000610338848285016102ff565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061036c82610341565b9050919050565b61037c81610361565b82525050565b60006020820190506103976000830184610373565b92915050565b600060ff82169050919050565b6103b38161039d565b82525050565b60006020820190506103ce60008301846103aa565b92915050565b6103dd81610361565b81146103e857600080fd5b50565b6000813590506103fa816103d4565b92915050565b600060208284031215610416576104156102d9565b5b6000610424848285016103eb565b91505092915050565b610436816102de565b82525050565b6000602082019050610451600083018461042d565b9291505056fea26469706673582212201b22fa1c0c5b70f30d070c091f8b1e37166f8054ae7c55b12220e7d99310445264736f6c634300080d0033",
}

// CompoundABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundMetaData.ABI instead.
var CompoundABI = CompoundMetaData.ABI

// CompoundBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompoundMetaData.Bin instead.
var CompoundBin = CompoundMetaData.Bin

// DeployCompound deploys a new Ethereum contract, binding an instance of Compound to it.
func DeployCompound(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Compound, error) {
	parsed, err := CompoundMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompoundBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// Compound is an auto generated Go binding around an Ethereum contract.
type Compound struct {
	CompoundCaller     // Read-only binding to the contract
	CompoundTransactor // Write-only binding to the contract
	CompoundFilterer   // Log filterer for contract events
}

// CompoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundSession struct {
	Contract     *Compound         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundCallerSession struct {
	Contract *CompoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundTransactorSession struct {
	Contract     *CompoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRaw struct {
	Contract *Compound // Generic contract binding to access the raw methods on
}

// CompoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundCallerRaw struct {
	Contract *CompoundCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundTransactorRaw struct {
	Contract *CompoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompound creates a new instance of Compound, bound to a specific deployed contract.
func NewCompound(address common.Address, backend bind.ContractBackend) (*Compound, error) {
	contract, err := bindCompound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// NewCompoundCaller creates a new read-only instance of Compound, bound to a specific deployed contract.
func NewCompoundCaller(address common.Address, caller bind.ContractCaller) (*CompoundCaller, error) {
	contract, err := bindCompound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundCaller{contract: contract}, nil
}

// NewCompoundTransactor creates a new write-only instance of Compound, bound to a specific deployed contract.
func NewCompoundTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundTransactor, error) {
	contract, err := bindCompound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTransactor{contract: contract}, nil
}

// NewCompoundFilterer creates a new log filterer instance of Compound, bound to a specific deployed contract.
func NewCompoundFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundFilterer, error) {
	contract, err := bindCompound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundFilterer{contract: contract}, nil
}

// bindCompound binds a generic wrapper to an already deployed contract.
func bindCompound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.CompoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transact(opts, method, params...)
}

// PROTOCOL is a free data retrieval call binding the contract method 0x91b9b827.
//
// Solidity: function PROTOCOL() view returns(uint8)
func (_Compound *CompoundCaller) PROTOCOL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "PROTOCOL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PROTOCOL is a free data retrieval call binding the contract method 0x91b9b827.
//
// Solidity: function PROTOCOL() view returns(uint8)
func (_Compound *CompoundSession) PROTOCOL() (uint8, error) {
	return _Compound.Contract.PROTOCOL(&_Compound.CallOpts)
}

// PROTOCOL is a free data retrieval call binding the contract method 0x91b9b827.
//
// Solidity: function PROTOCOL() view returns(uint8)
func (_Compound *CompoundCallerSession) PROTOCOL() (uint8, error) {
	return _Compound.Contract.PROTOCOL(&_Compound.CallOpts)
}

// ExchangeRateCalled is a free data retrieval call binding the contract method 0x5c62eb2f.
//
// Solidity: function exchangeRateCalled() view returns(address)
func (_Compound *CompoundCaller) ExchangeRateCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "exchangeRateCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeRateCalled is a free data retrieval call binding the contract method 0x5c62eb2f.
//
// Solidity: function exchangeRateCalled() view returns(address)
func (_Compound *CompoundSession) ExchangeRateCalled() (common.Address, error) {
	return _Compound.Contract.ExchangeRateCalled(&_Compound.CallOpts)
}

// ExchangeRateCalled is a free data retrieval call binding the contract method 0x5c62eb2f.
//
// Solidity: function exchangeRateCalled() view returns(address)
func (_Compound *CompoundCallerSession) ExchangeRateCalled() (common.Address, error) {
	return _Compound.Contract.ExchangeRateCalled(&_Compound.CallOpts)
}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Compound *CompoundCaller) UnderlyingCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "underlyingCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Compound *CompoundSession) UnderlyingCalled() (common.Address, error) {
	return _Compound.Contract.UnderlyingCalled(&_Compound.CallOpts)
}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Compound *CompoundCallerSession) UnderlyingCalled() (common.Address, error) {
	return _Compound.Contract.UnderlyingCalled(&_Compound.CallOpts)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address a) returns(uint256)
func (_Compound *CompoundTransactor) ExchangeRate(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "exchangeRate", a)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address a) returns(uint256)
func (_Compound *CompoundSession) ExchangeRate(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRate(&_Compound.TransactOpts, a)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address a) returns(uint256)
func (_Compound *CompoundTransactorSession) ExchangeRate(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRate(&_Compound.TransactOpts, a)
}

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 n) returns()
func (_Compound *CompoundTransactor) ExchangeRateReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "exchangeRateReturns", n)
}

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 n) returns()
func (_Compound *CompoundSession) ExchangeRateReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRateReturns(&_Compound.TransactOpts, n)
}

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 n) returns()
func (_Compound *CompoundTransactorSession) ExchangeRateReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRateReturns(&_Compound.TransactOpts, n)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address a) returns(address)
func (_Compound *CompoundTransactor) Underlying(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "underlying", a)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address a) returns(address)
func (_Compound *CompoundSession) Underlying(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Underlying(&_Compound.TransactOpts, a)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address a) returns(address)
func (_Compound *CompoundTransactorSession) Underlying(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Underlying(&_Compound.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Compound *CompoundTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Compound *CompoundSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UnderlyingReturns(&_Compound.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Compound *CompoundTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _Compound.Contract.UnderlyingReturns(&_Compound.TransactOpts, a)
}
