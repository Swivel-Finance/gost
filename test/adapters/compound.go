// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapters

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
	ABI: "[{\"inputs\":[],\"name\":\"PROTOCOL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061025f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806391b9b82714610046578063dc3b7c8b14610065578063e6919c8214610086575b600080fd5b61004e600081565b60405160ff90911681526020015b60405180910390f35b6100786100733660046101cf565b6100be565b60405190815260200161005c565b6100996100943660046101cf565b610137565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161005c565b60008173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b81526004016020604051808303816000875af115801561010d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061013191906101f3565b92915050565b60008173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610186573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610131919061020c565b73ffffffffffffffffffffffffffffffffffffffff811681146101cc57600080fd5b50565b6000602082840312156101e157600080fd5b81356101ec816101aa565b9392505050565b60006020828403121561020557600080fd5b5051919050565b60006020828403121561021e57600080fd5b81516101ec816101aa56fea2646970667358221220aae43eaa7d3977788c5036cb6d936d02467f90438a09ff3db02de1256bffbaba64736f6c634300080d0033",
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

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address c) returns(uint256)
func (_Compound *CompoundTransactor) ExchangeRate(opts *bind.TransactOpts, c common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "exchangeRate", c)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address c) returns(uint256)
func (_Compound *CompoundSession) ExchangeRate(c common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRate(&_Compound.TransactOpts, c)
}

// ExchangeRate is a paid mutator transaction binding the contract method 0xdc3b7c8b.
//
// Solidity: function exchangeRate(address c) returns(uint256)
func (_Compound *CompoundTransactorSession) ExchangeRate(c common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExchangeRate(&_Compound.TransactOpts, c)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address c) returns(address)
func (_Compound *CompoundTransactor) Underlying(opts *bind.TransactOpts, c common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "underlying", c)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address c) returns(address)
func (_Compound *CompoundSession) Underlying(c common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Underlying(&_Compound.TransactOpts, c)
}

// Underlying is a paid mutator transaction binding the contract method 0xe6919c82.
//
// Solidity: function underlying(address c) returns(address)
func (_Compound *CompoundTransactorSession) Underlying(c common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Underlying(&_Compound.TransactOpts, c)
}
