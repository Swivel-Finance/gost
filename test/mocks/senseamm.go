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

// SenseAMMABI is the input ABI used to generate the binding from.
const SenseAMMABI = "[{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseAMMBin is the compiled bytecode used for deploying new contracts.
var SenseAMMBin = "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063204f83f91461003b578063b4c4a4c814610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea26469706673582212205acc160e2ccfd11750ad37926187c289a56f7db9fdc1a4c6f7197d78f7fad7bf64736f6c634300080d0033"

// DeploySenseAMM deploys a new Ethereum contract, binding an instance of SenseAMM to it.
func DeploySenseAMM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenseAMM, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseAMMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SenseAMMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenseAMM{SenseAMMCaller: SenseAMMCaller{contract: contract}, SenseAMMTransactor: SenseAMMTransactor{contract: contract}, SenseAMMFilterer: SenseAMMFilterer{contract: contract}}, nil
}

// SenseAMM is an auto generated Go binding around an Ethereum contract.
type SenseAMM struct {
	SenseAMMCaller     // Read-only binding to the contract
	SenseAMMTransactor // Write-only binding to the contract
	SenseAMMFilterer   // Log filterer for contract events
}

// SenseAMMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseAMMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAMMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseAMMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAMMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseAMMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseAMMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseAMMSession struct {
	Contract     *SenseAMM         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseAMMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseAMMCallerSession struct {
	Contract *SenseAMMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SenseAMMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseAMMTransactorSession struct {
	Contract     *SenseAMMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SenseAMMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseAMMRaw struct {
	Contract *SenseAMM // Generic contract binding to access the raw methods on
}

// SenseAMMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseAMMCallerRaw struct {
	Contract *SenseAMMCaller // Generic read-only contract binding to access the raw methods on
}

// SenseAMMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseAMMTransactorRaw struct {
	Contract *SenseAMMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenseAMM creates a new instance of SenseAMM, bound to a specific deployed contract.
func NewSenseAMM(address common.Address, backend bind.ContractBackend) (*SenseAMM, error) {
	contract, err := bindSenseAMM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenseAMM{SenseAMMCaller: SenseAMMCaller{contract: contract}, SenseAMMTransactor: SenseAMMTransactor{contract: contract}, SenseAMMFilterer: SenseAMMFilterer{contract: contract}}, nil
}

// NewSenseAMMCaller creates a new read-only instance of SenseAMM, bound to a specific deployed contract.
func NewSenseAMMCaller(address common.Address, caller bind.ContractCaller) (*SenseAMMCaller, error) {
	contract, err := bindSenseAMM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseAMMCaller{contract: contract}, nil
}

// NewSenseAMMTransactor creates a new write-only instance of SenseAMM, bound to a specific deployed contract.
func NewSenseAMMTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseAMMTransactor, error) {
	contract, err := bindSenseAMM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseAMMTransactor{contract: contract}, nil
}

// NewSenseAMMFilterer creates a new log filterer instance of SenseAMM, bound to a specific deployed contract.
func NewSenseAMMFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseAMMFilterer, error) {
	contract, err := bindSenseAMM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseAMMFilterer{contract: contract}, nil
}

// bindSenseAMM binds a generic wrapper to an already deployed contract.
func bindSenseAMM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseAMMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseAMM *SenseAMMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseAMM.Contract.SenseAMMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseAMM *SenseAMMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseAMM.Contract.SenseAMMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseAMM *SenseAMMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseAMM.Contract.SenseAMMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseAMM *SenseAMMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseAMM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseAMM *SenseAMMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseAMM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseAMM *SenseAMMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseAMM.Contract.contract.Transact(opts, method, params...)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_SenseAMM *SenseAMMCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SenseAMM.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_SenseAMM *SenseAMMSession) Maturity() (*big.Int, error) {
	return _SenseAMM.Contract.Maturity(&_SenseAMM.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_SenseAMM *SenseAMMCallerSession) Maturity() (*big.Int, error) {
	return _SenseAMM.Contract.Maturity(&_SenseAMM.CallOpts)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 a) returns()
func (_SenseAMM *SenseAMMTransactor) MaturityReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _SenseAMM.contract.Transact(opts, "maturityReturns", a)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 a) returns()
func (_SenseAMM *SenseAMMSession) MaturityReturns(a *big.Int) (*types.Transaction, error) {
	return _SenseAMM.Contract.MaturityReturns(&_SenseAMM.TransactOpts, a)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 a) returns()
func (_SenseAMM *SenseAMMTransactorSession) MaturityReturns(a *big.Int) (*types.Transaction, error) {
	return _SenseAMM.Contract.MaturityReturns(&_SenseAMM.TransactOpts, a)
}
