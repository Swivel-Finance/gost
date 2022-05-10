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

// SenseTokenABI is the input ABI used to generate the binding from.
const SenseTokenABI = "[{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseTokenBin is the compiled bytecode used for deploying new contracts.
var SenseTokenBin = "0x608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636f307dc31461003b578063e7ba677414610059575b600080fd5b610043610075565b6040516100509190610122565b60405180910390f35b610073600480360381019061006e919061016e565b61009e565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010c826100e1565b9050919050565b61011c81610101565b82525050565b60006020820190506101376000830184610113565b92915050565b600080fd5b61014b81610101565b811461015657600080fd5b50565b60008135905061016881610142565b92915050565b6000602082840312156101845761018361013d565b5b600061019284828501610159565b9150509291505056fea26469706673582212207dd18821c1fa76b9ad427e5cb437ec6b7017b30e5ab672de71589b91bf23dc6864736f6c634300080d0033"

// DeploySenseToken deploys a new Ethereum contract, binding an instance of SenseToken to it.
func DeploySenseToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenseToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SenseTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenseToken{SenseTokenCaller: SenseTokenCaller{contract: contract}, SenseTokenTransactor: SenseTokenTransactor{contract: contract}, SenseTokenFilterer: SenseTokenFilterer{contract: contract}}, nil
}

// SenseToken is an auto generated Go binding around an Ethereum contract.
type SenseToken struct {
	SenseTokenCaller     // Read-only binding to the contract
	SenseTokenTransactor // Write-only binding to the contract
	SenseTokenFilterer   // Log filterer for contract events
}

// SenseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseTokenSession struct {
	Contract     *SenseToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseTokenCallerSession struct {
	Contract *SenseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SenseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseTokenTransactorSession struct {
	Contract     *SenseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SenseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseTokenRaw struct {
	Contract *SenseToken // Generic contract binding to access the raw methods on
}

// SenseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseTokenCallerRaw struct {
	Contract *SenseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SenseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseTokenTransactorRaw struct {
	Contract *SenseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenseToken creates a new instance of SenseToken, bound to a specific deployed contract.
func NewSenseToken(address common.Address, backend bind.ContractBackend) (*SenseToken, error) {
	contract, err := bindSenseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenseToken{SenseTokenCaller: SenseTokenCaller{contract: contract}, SenseTokenTransactor: SenseTokenTransactor{contract: contract}, SenseTokenFilterer: SenseTokenFilterer{contract: contract}}, nil
}

// NewSenseTokenCaller creates a new read-only instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenCaller(address common.Address, caller bind.ContractCaller) (*SenseTokenCaller, error) {
	contract, err := bindSenseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTokenCaller{contract: contract}, nil
}

// NewSenseTokenTransactor creates a new write-only instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseTokenTransactor, error) {
	contract, err := bindSenseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTokenTransactor{contract: contract}, nil
}

// NewSenseTokenFilterer creates a new log filterer instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseTokenFilterer, error) {
	contract, err := bindSenseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseTokenFilterer{contract: contract}, nil
}

// bindSenseToken binds a generic wrapper to an already deployed contract.
func bindSenseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseToken *SenseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseToken.Contract.SenseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseToken *SenseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseToken.Contract.SenseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseToken *SenseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseToken.Contract.SenseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseToken *SenseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseToken *SenseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseToken *SenseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseToken.Contract.contract.Transact(opts, method, params...)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenseToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenSession) Underlying() (common.Address, error) {
	return _SenseToken.Contract.Underlying(&_SenseToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenCallerSession) Underlying() (common.Address, error) {
	return _SenseToken.Contract.Underlying(&_SenseToken.CallOpts)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.UnderlyingReturns(&_SenseToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.UnderlyingReturns(&_SenseToken.TransactOpts, a)
}
