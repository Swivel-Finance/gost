// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// UnderlyingABI is the input ABI used to generate the binding from.
const UnderlyingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UnderlyingBin is the compiled bytecode used for deploying new contracts.
var UnderlyingBin = "0x608060405234801561001057600080fd5b506101ec806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806327e235e31461003b57806370a082311461006b575b600080fd5b61005560048036038101906100509190610110565b61009b565b6040516100629190610148565b60405180910390f35b61008560048036038101906100809190610110565b6100b3565b6040516100929190610148565b60405180910390f35b60006020528060005260406000206000915090505481565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008135905061010a8161019f565b92915050565b60006020828403121561012257600080fd5b6000610130848285016100fb565b91505092915050565b61014281610195565b82525050565b600060208201905061015d6000830184610139565b92915050565b600061016e82610175565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6101a881610163565b81146101b357600080fd5b5056fea26469706673582212208728a7fc7bab72c759f0cd0f09b5a20219b8b07a5414242a39a91d3bef5987b464736f6c63430008000033"

// DeployUnderlying deploys a new Ethereum contract, binding an instance of Underlying to it.
func DeployUnderlying(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Underlying, error) {
	parsed, err := abi.JSON(strings.NewReader(UnderlyingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnderlyingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Underlying{UnderlyingCaller: UnderlyingCaller{contract: contract}, UnderlyingTransactor: UnderlyingTransactor{contract: contract}, UnderlyingFilterer: UnderlyingFilterer{contract: contract}}, nil
}

// Underlying is an auto generated Go binding around an Ethereum contract.
type Underlying struct {
	UnderlyingCaller     // Read-only binding to the contract
	UnderlyingTransactor // Write-only binding to the contract
	UnderlyingFilterer   // Log filterer for contract events
}

// UnderlyingCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnderlyingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnderlyingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnderlyingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnderlyingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnderlyingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnderlyingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnderlyingSession struct {
	Contract     *Underlying       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnderlyingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnderlyingCallerSession struct {
	Contract *UnderlyingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnderlyingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnderlyingTransactorSession struct {
	Contract     *UnderlyingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnderlyingRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnderlyingRaw struct {
	Contract *Underlying // Generic contract binding to access the raw methods on
}

// UnderlyingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnderlyingCallerRaw struct {
	Contract *UnderlyingCaller // Generic read-only contract binding to access the raw methods on
}

// UnderlyingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnderlyingTransactorRaw struct {
	Contract *UnderlyingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnderlying creates a new instance of Underlying, bound to a specific deployed contract.
func NewUnderlying(address common.Address, backend bind.ContractBackend) (*Underlying, error) {
	contract, err := bindUnderlying(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Underlying{UnderlyingCaller: UnderlyingCaller{contract: contract}, UnderlyingTransactor: UnderlyingTransactor{contract: contract}, UnderlyingFilterer: UnderlyingFilterer{contract: contract}}, nil
}

// NewUnderlyingCaller creates a new read-only instance of Underlying, bound to a specific deployed contract.
func NewUnderlyingCaller(address common.Address, caller bind.ContractCaller) (*UnderlyingCaller, error) {
	contract, err := bindUnderlying(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnderlyingCaller{contract: contract}, nil
}

// NewUnderlyingTransactor creates a new write-only instance of Underlying, bound to a specific deployed contract.
func NewUnderlyingTransactor(address common.Address, transactor bind.ContractTransactor) (*UnderlyingTransactor, error) {
	contract, err := bindUnderlying(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnderlyingTransactor{contract: contract}, nil
}

// NewUnderlyingFilterer creates a new log filterer instance of Underlying, bound to a specific deployed contract.
func NewUnderlyingFilterer(address common.Address, filterer bind.ContractFilterer) (*UnderlyingFilterer, error) {
	contract, err := bindUnderlying(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnderlyingFilterer{contract: contract}, nil
}

// bindUnderlying binds a generic wrapper to an already deployed contract.
func bindUnderlying(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnderlyingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Underlying *UnderlyingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Underlying.Contract.UnderlyingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Underlying *UnderlyingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Underlying.Contract.UnderlyingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Underlying *UnderlyingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Underlying.Contract.UnderlyingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Underlying *UnderlyingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Underlying.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Underlying *UnderlyingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Underlying.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Underlying *UnderlyingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Underlying.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address o) view returns(uint256)
func (_Underlying *UnderlyingCaller) BalanceOf(opts *bind.CallOpts, o common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Underlying.contract.Call(opts, &out, "balanceOf", o)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address o) view returns(uint256)
func (_Underlying *UnderlyingSession) BalanceOf(o common.Address) (*big.Int, error) {
	return _Underlying.Contract.BalanceOf(&_Underlying.CallOpts, o)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address o) view returns(uint256)
func (_Underlying *UnderlyingCallerSession) BalanceOf(o common.Address) (*big.Int, error) {
	return _Underlying.Contract.BalanceOf(&_Underlying.CallOpts, o)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Underlying *UnderlyingCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Underlying.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Underlying *UnderlyingSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Underlying.Contract.Balances(&_Underlying.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Underlying *UnderlyingCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Underlying.Contract.Balances(&_Underlying.CallOpts, arg0)
}
