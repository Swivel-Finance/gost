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

// SenseAdapterABI is the input ABI used to generate the binding from.
const SenseAdapterABI = "[{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseAdapterBin is the compiled bytecode used for deploying new contracts.
var SenseAdapterBin = "0x608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636f307dc31461003b578063e7ba677414610059575b600080fd5b610043610075565b6040516100509190610122565b60405180910390f35b610073600480360381019061006e919061016e565b61009e565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010c826100e1565b9050919050565b61011c81610101565b82525050565b60006020820190506101376000830184610113565b92915050565b600080fd5b61014b81610101565b811461015657600080fd5b50565b60008135905061016881610142565b92915050565b6000602082840312156101845761018361013d565b5b600061019284828501610159565b9150509291505056fea2646970667358221220519b635174382a2da803a68ddc8dcb7a802348dbf105a0f4f318d06d892e3ca964736f6c634300080d0033"

// DeploySenseAdapter deploys a new Ethereum contract, binding an instance of SenseAdapter to it.
func DeploySenseAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenseAdapter, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseAdapterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SenseAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenseAdapter{SenseAdapterCaller: SenseAdapterCaller{contract: contract}, SenseAdapterTransactor: SenseAdapterTransactor{contract: contract}, SenseAdapterFilterer: SenseAdapterFilterer{contract: contract}}, nil
}

// SenseAdapter is an auto generated Go binding around an Ethereum contract.
type SenseAdapter struct {
	SenseAdapterCaller     // Read-only binding to the contract
	SenseAdapterTransactor // Write-only binding to the contract
	SenseAdapterFilterer   // Log filterer for contract events
}

// SenseAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseAdapterSession struct {
	Contract     *SenseAdapter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseAdapterCallerSession struct {
	Contract *SenseAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SenseAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseAdapterTransactorSession struct {
	Contract     *SenseAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SenseAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseAdapterRaw struct {
	Contract *SenseAdapter // Generic contract binding to access the raw methods on
}

// SenseAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseAdapterCallerRaw struct {
	Contract *SenseAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// SenseAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseAdapterTransactorRaw struct {
	Contract *SenseAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenseAdapter creates a new instance of SenseAdapter, bound to a specific deployed contract.
func NewSenseAdapter(address common.Address, backend bind.ContractBackend) (*SenseAdapter, error) {
	contract, err := bindSenseAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenseAdapter{SenseAdapterCaller: SenseAdapterCaller{contract: contract}, SenseAdapterTransactor: SenseAdapterTransactor{contract: contract}, SenseAdapterFilterer: SenseAdapterFilterer{contract: contract}}, nil
}

// NewSenseAdapterCaller creates a new read-only instance of SenseAdapter, bound to a specific deployed contract.
func NewSenseAdapterCaller(address common.Address, caller bind.ContractCaller) (*SenseAdapterCaller, error) {
	contract, err := bindSenseAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseAdapterCaller{contract: contract}, nil
}

// NewSenseAdapterTransactor creates a new write-only instance of SenseAdapter, bound to a specific deployed contract.
func NewSenseAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseAdapterTransactor, error) {
	contract, err := bindSenseAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseAdapterTransactor{contract: contract}, nil
}

// NewSenseAdapterFilterer creates a new log filterer instance of SenseAdapter, bound to a specific deployed contract.
func NewSenseAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseAdapterFilterer, error) {
	contract, err := bindSenseAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseAdapterFilterer{contract: contract}, nil
}

// bindSenseAdapter binds a generic wrapper to an already deployed contract.
func bindSenseAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseAdapter *SenseAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseAdapter.Contract.SenseAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseAdapter *SenseAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseAdapter.Contract.SenseAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseAdapter *SenseAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseAdapter.Contract.SenseAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseAdapter *SenseAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseAdapter *SenseAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseAdapter *SenseAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseAdapter.Contract.contract.Transact(opts, method, params...)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseAdapter *SenseAdapterCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenseAdapter.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseAdapter *SenseAdapterSession) Underlying() (common.Address, error) {
	return _SenseAdapter.Contract.Underlying(&_SenseAdapter.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseAdapter *SenseAdapterCallerSession) Underlying() (common.Address, error) {
	return _SenseAdapter.Contract.Underlying(&_SenseAdapter.CallOpts)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseAdapter *SenseAdapterTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _SenseAdapter.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseAdapter *SenseAdapterSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseAdapter.Contract.UnderlyingReturns(&_SenseAdapter.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseAdapter *SenseAdapterTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseAdapter.Contract.UnderlyingReturns(&_SenseAdapter.TransactOpts, a)
}
