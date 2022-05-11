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

// APWineTokenMetaData contains all meta data concerning the APWineToken contract.
var APWineTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getPTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getPTAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631e5f74a11461003b5780639a8cb8f314610059575b600080fd5b610043610075565b6040516100509190610122565b60405180910390f35b610073600480360381019061006e919061016e565b61009e565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010c826100e1565b9050919050565b61011c81610101565b82525050565b60006020820190506101376000830184610113565b92915050565b600080fd5b61014b81610101565b811461015657600080fd5b50565b60008135905061016881610142565b92915050565b6000602082840312156101845761018361013d565b5b600061019284828501610159565b9150509291505056fea2646970667358221220170efc283760638e14d30eb332a0b929da823596810d82cbbd38f903526a97ac64736f6c634300080d0033",
}

// APWineTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use APWineTokenMetaData.ABI instead.
var APWineTokenABI = APWineTokenMetaData.ABI

// APWineTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use APWineTokenMetaData.Bin instead.
var APWineTokenBin = APWineTokenMetaData.Bin

// DeployAPWineToken deploys a new Ethereum contract, binding an instance of APWineToken to it.
func DeployAPWineToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWineToken, error) {
	parsed, err := APWineTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(APWineTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// APWineToken is an auto generated Go binding around an Ethereum contract.
type APWineToken struct {
	APWineTokenCaller     // Read-only binding to the contract
	APWineTokenTransactor // Write-only binding to the contract
	APWineTokenFilterer   // Log filterer for contract events
}

// APWineTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineTokenSession struct {
	Contract     *APWineToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineTokenCallerSession struct {
	Contract *APWineTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// APWineTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineTokenTransactorSession struct {
	Contract     *APWineTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// APWineTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineTokenRaw struct {
	Contract *APWineToken // Generic contract binding to access the raw methods on
}

// APWineTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineTokenCallerRaw struct {
	Contract *APWineTokenCaller // Generic read-only contract binding to access the raw methods on
}

// APWineTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineTokenTransactorRaw struct {
	Contract *APWineTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWineToken creates a new instance of APWineToken, bound to a specific deployed contract.
func NewAPWineToken(address common.Address, backend bind.ContractBackend) (*APWineToken, error) {
	contract, err := bindAPWineToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// NewAPWineTokenCaller creates a new read-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenCaller(address common.Address, caller bind.ContractCaller) (*APWineTokenCaller, error) {
	contract, err := bindAPWineToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenCaller{contract: contract}, nil
}

// NewAPWineTokenTransactor creates a new write-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineTokenTransactor, error) {
	contract, err := bindAPWineToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenTransactor{contract: contract}, nil
}

// NewAPWineTokenFilterer creates a new log filterer instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineTokenFilterer, error) {
	contract, err := bindAPWineToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineTokenFilterer{contract: contract}, nil
}

// bindAPWineToken binds a generic wrapper to an already deployed contract.
func bindAPWineToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.APWineTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transact(opts, method, params...)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenCaller) GetPTAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "getPTAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenSession) GetPTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetPTAddress(&_APWineToken.CallOpts)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenCallerSession) GetPTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetPTAddress(&_APWineToken.CallOpts)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactor) GetPTAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "getPTAddressReturns", a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetPTAddressReturns(&_APWineToken.TransactOpts, a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactorSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetPTAddressReturns(&_APWineToken.TransactOpts, a)
}
