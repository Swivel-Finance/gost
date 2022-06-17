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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610578806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063d6bcd7aa11610071578063d6bcd7aa1461018d578063dc3b7c8b146101ab578063e6919c82146101db578063e7a7b9ce1461020b578063e7ba677414610227578063f189b33714610243576100b4565b806329d9ce3e146100b9578063589aeec7146100d55780635c62eb2f146100f1578063852a12e31461010f578063a0712d681461013f578063d4e7fdd41461016f575b600080fd5b6100d360048036038101906100ce9190610436565b610261565b005b6100ef60048036038101906100ea9190610436565b61026b565b005b6100f9610275565b60405161010691906104a4565b60405180910390f35b61012960048036038101906101249190610436565b61029b565b60405161013691906104ce565b60405180910390f35b61015960048036038101906101549190610436565b6102ae565b60405161016691906104ce565b60405180910390f35b6101776102c1565b60405161018491906104ce565b60405180910390f35b6101956102c7565b6040516101a291906104ce565b60405180910390f35b6101c560048036038101906101c09190610515565b6102cd565b6040516101d291906104ce565b60405180910390f35b6101f560048036038101906101f09190610515565b61031a565b60405161020291906104a4565b60405180910390f35b61022560048036038101906102209190610436565b610387565b005b610241600480360381019061023c9190610515565b610391565b005b61024b6103d5565b60405161025891906104a4565b60405180910390f35b8060028190555050565b8060078190555050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000816003819055506002549050919050565b6000816001819055506000549050919050565b60015481565b60035481565b600081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506007549050919050565b600081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8060008190555050565b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080fd5b6000819050919050565b61041381610400565b811461041e57600080fd5b50565b6000813590506104308161040a565b92915050565b60006020828403121561044c5761044b6103fb565b5b600061045a84828501610421565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061048e82610463565b9050919050565b61049e81610483565b82525050565b60006020820190506104b96000830184610495565b92915050565b6104c881610400565b82525050565b60006020820190506104e360008301846104bf565b92915050565b6104f281610483565b81146104fd57600080fd5b50565b60008135905061050f816104e9565b92915050565b60006020828403121561052b5761052a6103fb565b5b600061053984828501610500565b9150509291505056fea26469706673582212208f51169f53729f5f779701b86c8e86bde4f0dc97e49ace9882ddf899c827bac664736f6c634300080d0033",
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

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_Compound *CompoundCaller) MintCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "mintCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_Compound *CompoundSession) MintCalled() (*big.Int, error) {
	return _Compound.Contract.MintCalled(&_Compound.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_Compound *CompoundCallerSession) MintCalled() (*big.Int, error) {
	return _Compound.Contract.MintCalled(&_Compound.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_Compound *CompoundCaller) RedeemUnderlyingCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "redeemUnderlyingCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_Compound *CompoundSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _Compound.Contract.RedeemUnderlyingCalled(&_Compound.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_Compound *CompoundCallerSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _Compound.Contract.RedeemUnderlyingCalled(&_Compound.CallOpts)
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

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_Compound *CompoundTransactor) Mint(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "mint", n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_Compound *CompoundSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Mint(&_Compound.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_Compound *CompoundTransactorSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Mint(&_Compound.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_Compound *CompoundTransactor) MintReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "mintReturns", n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_Compound *CompoundSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintReturns(&_Compound.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_Compound *CompoundTransactorSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintReturns(&_Compound.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_Compound *CompoundTransactor) RedeemUnderlying(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeemUnderlying", n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_Compound *CompoundSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlying(&_Compound.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_Compound *CompoundTransactorSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlying(&_Compound.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_Compound *CompoundTransactor) RedeemUnderlyingReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeemUnderlyingReturns", n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_Compound *CompoundSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlyingReturns(&_Compound.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_Compound *CompoundTransactorSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlyingReturns(&_Compound.TransactOpts, n)
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
