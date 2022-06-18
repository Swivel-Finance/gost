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
	ABI: "[{\"inputs\":[],\"name\":\"PROTOCOL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060006101000a81548160ff021916908360ff16021790555034801561002b57600080fd5b506106fd8061003b6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063962941781161008c578063e7a7b9ce11610066578063e7a7b9ce14610238578063e7ba677414610254578063ee4db57014610270578063f189b337146102a0576100cf565b806396294178146101a8578063dc3b7c8b146101d8578063e6919c8214610208576100cf565b806329d9ce3e146100d457806340c10f19146100f0578063589aeec7146101205780635c62eb2f1461013c57806364f670ee1461015a57806391b9b8271461018a575b600080fd5b6100ee60048036038101906100e99190610544565b6102be565b005b61010a600480360381019061010591906105cf565b6102c8565b604051610117919061061e565b60405180910390f35b61013a60048036038101906101359190610544565b610319565b005b610144610323565b6040516101519190610648565b60405180910390f35b610174600480360381019061016f9190610663565b610349565b604051610181919061061e565b60405180910390f35b610192610361565b60405161019f91906106ac565b60405180910390f35b6101c260048036038101906101bd91906105cf565b610372565b6040516101cf919061061e565b60405180910390f35b6101f260048036038101906101ed9190610663565b6103c3565b6040516101ff919061061e565b60405180910390f35b610222600480360381019061021d9190610663565b610410565b60405161022f9190610648565b60405180910390f35b610252600480360381019061024d9190610544565b61047d565b005b61026e60048036038101906102699190610663565b610487565b005b61028a60048036038101906102859190610663565b6104cb565b604051610297919061061e565b60405180910390f35b6102a86104e3565b6040516102b59190610648565b60405180910390f35b8060038190555050565b600081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600154905092915050565b8060088190555050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915090505481565b60008054906101000a900460ff1681565b600081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600354905092915050565b600081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008549050919050565b600081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8060018190555050565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60026020528060005260406000206000915090505481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080fd5b6000819050919050565b6105218161050e565b811461052c57600080fd5b50565b60008135905061053e81610518565b92915050565b60006020828403121561055a57610559610509565b5b60006105688482850161052f565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061059c82610571565b9050919050565b6105ac81610591565b81146105b757600080fd5b50565b6000813590506105c9816105a3565b92915050565b600080604083850312156105e6576105e5610509565b5b60006105f4858286016105ba565b92505060206106058582860161052f565b9150509250929050565b6106188161050e565b82525050565b6000602082019050610633600083018461060f565b92915050565b61064281610591565b82525050565b600060208201905061065d6000830184610639565b92915050565b60006020828403121561067957610678610509565b5b6000610687848285016105ba565b91505092915050565b600060ff82169050919050565b6106a681610690565b82525050565b60006020820190506106c1600083018461069d565b9291505056fea2646970667358221220f96466f65ff7e9eef21a1176ba88af056c7fdcf83c88855dfdad99e0c2f8e0ba64736f6c634300080d0033",
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

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_Compound *CompoundCaller) MintCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "mintCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_Compound *CompoundSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.MintCalled(&_Compound.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_Compound *CompoundCallerSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.MintCalled(&_Compound.CallOpts, arg0)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_Compound *CompoundCaller) RedeemUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "redeemUnderlyingCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_Compound *CompoundSession) RedeemUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.RedeemUnderlyingCalled(&_Compound.CallOpts, arg0)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_Compound *CompoundCallerSession) RedeemUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.RedeemUnderlyingCalled(&_Compound.CallOpts, arg0)
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

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address c, uint256 n) returns(uint256)
func (_Compound *CompoundTransactor) Mint(opts *bind.TransactOpts, c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "mint", c, n)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address c, uint256 n) returns(uint256)
func (_Compound *CompoundSession) Mint(c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Mint(&_Compound.TransactOpts, c, n)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address c, uint256 n) returns(uint256)
func (_Compound *CompoundTransactorSession) Mint(c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.Mint(&_Compound.TransactOpts, c, n)
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

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address c, uint256 n) returns(uint256)
func (_Compound *CompoundTransactor) RedeemUnderlying(opts *bind.TransactOpts, c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeemUnderlying", c, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address c, uint256 n) returns(uint256)
func (_Compound *CompoundSession) RedeemUnderlying(c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlying(&_Compound.TransactOpts, c, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address c, uint256 n) returns(uint256)
func (_Compound *CompoundTransactorSession) RedeemUnderlying(c common.Address, n *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemUnderlying(&_Compound.TransactOpts, c, n)
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
