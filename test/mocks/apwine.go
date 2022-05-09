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

// APWineABI is the input ABI used to generate the binding from.
const APWineABI = "[{\"inputs\":[],\"name\":\"getPTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getPTAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// APWineBin is the compiled bytecode used for deploying new contracts.
var APWineBin = "0x608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631e5f74a11461003b5780639a8cb8f314610059575b600080fd5b610043610075565b6040516100509190610122565b60405180910390f35b610073600480360381019061006e919061016e565b61009e565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010c826100e1565b9050919050565b61011c81610101565b82525050565b60006020820190506101376000830184610113565b92915050565b600080fd5b61014b81610101565b811461015657600080fd5b50565b60008135905061016881610142565b92915050565b6000602082840312156101845761018361013d565b5b600061019284828501610159565b9150509291505056fea26469706673582212204cc9955fc23a662161d6f25d6d99dd142a1ed71fde42cb8ad958fbfdff05d50564736f6c634300080d0033"

// DeployAPWine deploys a new Ethereum contract, binding an instance of APWine to it.
func DeployAPWine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWine, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APWineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWine{APWineCaller: APWineCaller{contract: contract}, APWineTransactor: APWineTransactor{contract: contract}, APWineFilterer: APWineFilterer{contract: contract}}, nil
}

// APWine is an auto generated Go binding around an Ethereum contract.
type APWine struct {
	APWineCaller     // Read-only binding to the contract
	APWineTransactor // Write-only binding to the contract
	APWineFilterer   // Log filterer for contract events
}

// APWineCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineSession struct {
	Contract     *APWine           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineCallerSession struct {
	Contract *APWineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// APWineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineTransactorSession struct {
	Contract     *APWineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineRaw struct {
	Contract *APWine // Generic contract binding to access the raw methods on
}

// APWineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineCallerRaw struct {
	Contract *APWineCaller // Generic read-only contract binding to access the raw methods on
}

// APWineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineTransactorRaw struct {
	Contract *APWineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWine creates a new instance of APWine, bound to a specific deployed contract.
func NewAPWine(address common.Address, backend bind.ContractBackend) (*APWine, error) {
	contract, err := bindAPWine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWine{APWineCaller: APWineCaller{contract: contract}, APWineTransactor: APWineTransactor{contract: contract}, APWineFilterer: APWineFilterer{contract: contract}}, nil
}

// NewAPWineCaller creates a new read-only instance of APWine, bound to a specific deployed contract.
func NewAPWineCaller(address common.Address, caller bind.ContractCaller) (*APWineCaller, error) {
	contract, err := bindAPWine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineCaller{contract: contract}, nil
}

// NewAPWineTransactor creates a new write-only instance of APWine, bound to a specific deployed contract.
func NewAPWineTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineTransactor, error) {
	contract, err := bindAPWine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTransactor{contract: contract}, nil
}

// NewAPWineFilterer creates a new log filterer instance of APWine, bound to a specific deployed contract.
func NewAPWineFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineFilterer, error) {
	contract, err := bindAPWine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineFilterer{contract: contract}, nil
}

// bindAPWine binds a generic wrapper to an already deployed contract.
func bindAPWine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWine *APWineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWine.Contract.APWineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWine *APWineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWine.Contract.APWineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWine *APWineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWine.Contract.APWineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWine *APWineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWine *APWineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWine *APWineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWine.Contract.contract.Transact(opts, method, params...)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWine *APWineCaller) GetPTAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "getPTAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWine *APWineSession) GetPTAddress() (common.Address, error) {
	return _APWine.Contract.GetPTAddress(&_APWine.CallOpts)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWine *APWineCallerSession) GetPTAddress() (common.Address, error) {
	return _APWine.Contract.GetPTAddress(&_APWine.CallOpts)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWine *APWineTransactor) GetPTAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _APWine.contract.Transact(opts, "getPTAddressReturns", a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWine *APWineSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWine.Contract.GetPTAddressReturns(&_APWine.TransactOpts, a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWine *APWineTransactorSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWine.Contract.GetPTAddressReturns(&_APWine.TransactOpts, a)
}
