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

// SenseABI is the input ABI used to generate the binding from.
const SenseABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBoughtCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"senseAdapterCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sa\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseBin is the compiled bytecode used for deploying new contracts.
var SenseBin = "0x608060405234801561001057600080fd5b50610389806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806336a9449e146100675780633ef943bb1461008557806357e5b682146100a35780638f1f30f0146100c157806393824eca146100f1578063cbf6a44c1461010f575b600080fd5b61006f61012b565b60405161007c9190610213565b60405180910390f35b61008d610151565b60405161009a9190610247565b60405180910390f35b6100ab610157565b6040516100b89190610247565b60405180910390f35b6100db60048036038101906100d691906102bf565b61015d565b6040516100e89190610247565b60405180910390f35b6100f96101c2565b6040516101069190610247565b60405180910390f35b61012960048036038101906101249190610326565b6101c8565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60025481565b600084600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360028190555082600381905550816004819055506000549050949350505050565b60045481565b8060008190555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101fd826101d2565b9050919050565b61020d816101f2565b82525050565b60006020820190506102286000830184610204565b92915050565b6000819050919050565b6102418161022e565b82525050565b600060208201905061025c6000830184610238565b92915050565b600080fd5b610270816101f2565b811461027b57600080fd5b50565b60008135905061028d81610267565b92915050565b61029c8161022e565b81146102a757600080fd5b50565b6000813590506102b981610293565b92915050565b600080600080608085870312156102d9576102d8610262565b5b60006102e78782880161027e565b94505060206102f8878288016102aa565b9350506040610309878288016102aa565b925050606061031a878288016102aa565b91505092959194509250565b60006020828403121561033c5761033b610262565b5b600061034a848285016102aa565b9150509291505056fea264697066735822122045a2c0844d273f5476732417580ea252cc5e5877a5d3cf3b123aa01a5e8673d464736f6c634300080d0033"

// DeploySense deploys a new Ethereum contract, binding an instance of Sense to it.
func DeploySense(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sense, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SenseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sense{SenseCaller: SenseCaller{contract: contract}, SenseTransactor: SenseTransactor{contract: contract}, SenseFilterer: SenseFilterer{contract: contract}}, nil
}

// Sense is an auto generated Go binding around an Ethereum contract.
type Sense struct {
	SenseCaller     // Read-only binding to the contract
	SenseTransactor // Write-only binding to the contract
	SenseFilterer   // Log filterer for contract events
}

// SenseCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseSession struct {
	Contract     *Sense            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseCallerSession struct {
	Contract *SenseCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SenseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseTransactorSession struct {
	Contract     *SenseTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseRaw struct {
	Contract *Sense // Generic contract binding to access the raw methods on
}

// SenseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseCallerRaw struct {
	Contract *SenseCaller // Generic read-only contract binding to access the raw methods on
}

// SenseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseTransactorRaw struct {
	Contract *SenseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSense creates a new instance of Sense, bound to a specific deployed contract.
func NewSense(address common.Address, backend bind.ContractBackend) (*Sense, error) {
	contract, err := bindSense(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sense{SenseCaller: SenseCaller{contract: contract}, SenseTransactor: SenseTransactor{contract: contract}, SenseFilterer: SenseFilterer{contract: contract}}, nil
}

// NewSenseCaller creates a new read-only instance of Sense, bound to a specific deployed contract.
func NewSenseCaller(address common.Address, caller bind.ContractCaller) (*SenseCaller, error) {
	contract, err := bindSense(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseCaller{contract: contract}, nil
}

// NewSenseTransactor creates a new write-only instance of Sense, bound to a specific deployed contract.
func NewSenseTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseTransactor, error) {
	contract, err := bindSense(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTransactor{contract: contract}, nil
}

// NewSenseFilterer creates a new log filterer instance of Sense, bound to a specific deployed contract.
func NewSenseFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseFilterer, error) {
	contract, err := bindSense(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseFilterer{contract: contract}, nil
}

// bindSense binds a generic wrapper to an already deployed contract.
func bindSense(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sense *SenseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sense.Contract.SenseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sense *SenseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sense.Contract.SenseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sense *SenseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sense.Contract.SenseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sense *SenseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sense.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sense *SenseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sense.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sense *SenseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sense.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Sense *SenseCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Sense *SenseSession) AmountCalled() (*big.Int, error) {
	return _Sense.Contract.AmountCalled(&_Sense.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Sense *SenseCallerSession) AmountCalled() (*big.Int, error) {
	return _Sense.Contract.AmountCalled(&_Sense.CallOpts)
}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Sense *SenseCaller) MaturityCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "maturityCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Sense *SenseSession) MaturityCalled() (*big.Int, error) {
	return _Sense.Contract.MaturityCalled(&_Sense.CallOpts)
}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Sense *SenseCallerSession) MaturityCalled() (*big.Int, error) {
	return _Sense.Contract.MaturityCalled(&_Sense.CallOpts)
}

// MinimumBoughtCalled is a free data retrieval call binding the contract method 0x93824eca.
//
// Solidity: function minimumBoughtCalled() view returns(uint256)
func (_Sense *SenseCaller) MinimumBoughtCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "minimumBoughtCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumBoughtCalled is a free data retrieval call binding the contract method 0x93824eca.
//
// Solidity: function minimumBoughtCalled() view returns(uint256)
func (_Sense *SenseSession) MinimumBoughtCalled() (*big.Int, error) {
	return _Sense.Contract.MinimumBoughtCalled(&_Sense.CallOpts)
}

// MinimumBoughtCalled is a free data retrieval call binding the contract method 0x93824eca.
//
// Solidity: function minimumBoughtCalled() view returns(uint256)
func (_Sense *SenseCallerSession) MinimumBoughtCalled() (*big.Int, error) {
	return _Sense.Contract.MinimumBoughtCalled(&_Sense.CallOpts)
}

// SenseAdapterCalled is a free data retrieval call binding the contract method 0x36a9449e.
//
// Solidity: function senseAdapterCalled() view returns(address)
func (_Sense *SenseCaller) SenseAdapterCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "senseAdapterCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SenseAdapterCalled is a free data retrieval call binding the contract method 0x36a9449e.
//
// Solidity: function senseAdapterCalled() view returns(address)
func (_Sense *SenseSession) SenseAdapterCalled() (common.Address, error) {
	return _Sense.Contract.SenseAdapterCalled(&_Sense.CallOpts)
}

// SenseAdapterCalled is a free data retrieval call binding the contract method 0x36a9449e.
//
// Solidity: function senseAdapterCalled() view returns(address)
func (_Sense *SenseCallerSession) SenseAdapterCalled() (common.Address, error) {
	return _Sense.Contract.SenseAdapterCalled(&_Sense.CallOpts)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseTransactor) SwapUnderlyingForPTs(opts *bind.TransactOpts, sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "swapUnderlyingForPTs", sa, m, a, mb)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseSession) SwapUnderlyingForPTs(sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTs(&_Sense.TransactOpts, sa, m, a, mb)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseTransactorSession) SwapUnderlyingForPTs(sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTs(&_Sense.TransactOpts, sa, m, a, mb)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseTransactor) SwapUnderlyingForPTsReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "swapUnderlyingForPTsReturns", s)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseSession) SwapUnderlyingForPTsReturns(s *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTsReturns(&_Sense.TransactOpts, s)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseTransactorSession) SwapUnderlyingForPTsReturns(s *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTsReturns(&_Sense.TransactOpts, s)
}
