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
const PendleABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldTokenHolders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PendleBin is the compiled bytecode used for deploying new contracts.
var PendleBin = "0x608060405234801561001057600080fd5b50610289806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063af7705be14610046578063b4c4a4c814610065578063e7ba677414610081575b600080fd5b61004e61009d565b60405161005c929190610173565b60405180910390f35b61007f600480360381019061007a91906101cd565b6100cc565b005b61009b60048036038101906100969190610226565b6100d6565b005b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600154915091509091565b8060018190555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061014482610119565b9050919050565b61015481610139565b82525050565b6000819050919050565b61016d8161015a565b82525050565b6000604082019050610188600083018561014b565b6101956020830184610164565b9392505050565b600080fd5b6101aa8161015a565b81146101b557600080fd5b50565b6000813590506101c7816101a1565b92915050565b6000602082840312156101e3576101e261019c565b5b60006101f1848285016101b8565b91505092915050565b61020381610139565b811461020e57600080fd5b50565b600081359050610220816101fa565b92915050565b60006020828403121561023c5761023b61019c565b5b600061024a84828501610211565b9150509291505056fea26469706673582212202762a042f018cbec6ee145450d8ba06f62e77afc8fcbdc82426eec483384654d64736f6c634300080d0033"

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

// YieldTokenHolders is a free data retrieval call binding the contract method 0xaf7705be.
//
// Solidity: function yieldTokenHolders() view returns(address, uint256)
func (_Pendle *PendleCaller) YieldTokenHolders(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "yieldTokenHolders")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// YieldTokenHolders is a free data retrieval call binding the contract method 0xaf7705be.
//
// Solidity: function yieldTokenHolders() view returns(address, uint256)
func (_Pendle *PendleSession) YieldTokenHolders() (common.Address, *big.Int, error) {
	return _Pendle.Contract.YieldTokenHolders(&_Pendle.CallOpts)
}

// YieldTokenHolders is a free data retrieval call binding the contract method 0xaf7705be.
//
// Solidity: function yieldTokenHolders() view returns(address, uint256)
func (_Pendle *PendleCallerSession) YieldTokenHolders() (common.Address, *big.Int, error) {
	return _Pendle.Contract.YieldTokenHolders(&_Pendle.CallOpts)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Pendle *PendleTransactor) MaturityReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "maturityReturns", m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Pendle *PendleSession) MaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.MaturityReturns(&_Pendle.TransactOpts, m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Pendle *PendleTransactorSession) MaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.MaturityReturns(&_Pendle.TransactOpts, m)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Pendle *PendleTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Pendle *PendleSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _Pendle.Contract.UnderlyingReturns(&_Pendle.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_Pendle *PendleTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _Pendle.Contract.UnderlyingReturns(&_Pendle.TransactOpts, a)
}
