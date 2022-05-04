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

// PendleRouterABI is the input ABI used to generate the binding from.
const PendleRouterABI = "[{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"contractIPendleData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPendleData\",\"name\":\"x\",\"type\":\"address\"}],\"name\":\"dataReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"tokenizeYield\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"tokenizeYieldReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleRouterBin is the compiled bytecode used for deploying new contracts.
var PendleRouterBin = "0x608060405234801561001057600080fd5b506104eb806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806373d4a13a14610051578063bae3b1121461006f578063bd7389701461008b578063d8088d03146100bd575b600080fd5b6100596100d9565b6040516100669190610283565b60405180910390f35b610089600480360381019061008491906102d9565b610103565b005b6100a560048036038101906100a0919061037a565b61010d565b6040516100b493929190610413565b60405180910390f35b6100d760048036038101906100d29190610488565b6101c0565b005b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060008190555050565b60008060008760028190555086600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856004819055508460058190555083600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080600054925092509250955095509592505050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061024961024461023f84610204565b610224565b610204565b9050919050565b600061025b8261022e565b9050919050565b600061026d82610250565b9050919050565b61027d81610262565b82525050565b60006020820190506102986000830184610274565b92915050565b600080fd5b6000819050919050565b6102b6816102a3565b81146102c157600080fd5b50565b6000813590506102d3816102ad565b92915050565b6000602082840312156102ef576102ee61029e565b5b60006102fd848285016102c4565b91505092915050565b6000819050919050565b61031981610306565b811461032457600080fd5b50565b60008135905061033681610310565b92915050565b600061034782610204565b9050919050565b6103578161033c565b811461036257600080fd5b50565b6000813590506103748161034e565b92915050565b600080600080600060a086880312156103965761039561029e565b5b60006103a488828901610327565b95505060206103b588828901610365565b94505060406103c6888289016102c4565b93505060606103d7888289016102c4565b92505060806103e888828901610365565b9150509295509295909350565b6103fe8161033c565b82525050565b61040d816102a3565b82525050565b600060608201905061042860008301866103f5565b61043560208301856103f5565b6104426040830184610404565b949350505050565b60006104558261033c565b9050919050565b6104658161044a565b811461047057600080fd5b50565b6000813590506104828161045c565b92915050565b60006020828403121561049e5761049d61029e565b5b60006104ac84828501610473565b9150509291505056fea26469706673582212206216252873f5ccf0c0ea42ef89035d175cd112c24e0e1a6a0c8b9c2344503c4064736f6c634300080d0033"

// DeployPendleRouter deploys a new Ethereum contract, binding an instance of PendleRouter to it.
func DeployPendleRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleRouter{PendleRouterCaller: PendleRouterCaller{contract: contract}, PendleRouterTransactor: PendleRouterTransactor{contract: contract}, PendleRouterFilterer: PendleRouterFilterer{contract: contract}}, nil
}

// PendleRouter is an auto generated Go binding around an Ethereum contract.
type PendleRouter struct {
	PendleRouterCaller     // Read-only binding to the contract
	PendleRouterTransactor // Write-only binding to the contract
	PendleRouterFilterer   // Log filterer for contract events
}

// PendleRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleRouterSession struct {
	Contract     *PendleRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleRouterCallerSession struct {
	Contract *PendleRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PendleRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleRouterTransactorSession struct {
	Contract     *PendleRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PendleRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleRouterRaw struct {
	Contract *PendleRouter // Generic contract binding to access the raw methods on
}

// PendleRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleRouterCallerRaw struct {
	Contract *PendleRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PendleRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleRouterTransactorRaw struct {
	Contract *PendleRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleRouter creates a new instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouter(address common.Address, backend bind.ContractBackend) (*PendleRouter, error) {
	contract, err := bindPendleRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleRouter{PendleRouterCaller: PendleRouterCaller{contract: contract}, PendleRouterTransactor: PendleRouterTransactor{contract: contract}, PendleRouterFilterer: PendleRouterFilterer{contract: contract}}, nil
}

// NewPendleRouterCaller creates a new read-only instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterCaller(address common.Address, caller bind.ContractCaller) (*PendleRouterCaller, error) {
	contract, err := bindPendleRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterCaller{contract: contract}, nil
}

// NewPendleRouterTransactor creates a new write-only instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleRouterTransactor, error) {
	contract, err := bindPendleRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterTransactor{contract: contract}, nil
}

// NewPendleRouterFilterer creates a new log filterer instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleRouterFilterer, error) {
	contract, err := bindPendleRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleRouterFilterer{contract: contract}, nil
}

// bindPendleRouter binds a generic wrapper to an already deployed contract.
func bindPendleRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouter *PendleRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouter.Contract.PendleRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouter *PendleRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouter.Contract.PendleRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouter *PendleRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouter.Contract.PendleRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouter *PendleRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouter *PendleRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouter *PendleRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouter.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterSession) Data() (common.Address, error) {
	return _PendleRouter.Contract.Data(&_PendleRouter.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterCallerSession) Data() (common.Address, error) {
	return _PendleRouter.Contract.Data(&_PendleRouter.CallOpts)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterTransactor) DataReturns(opts *bind.TransactOpts, x common.Address) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "dataReturns", x)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterSession) DataReturns(x common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.DataReturns(&_PendleRouter.TransactOpts, x)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterTransactorSession) DataReturns(x common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.DataReturns(&_PendleRouter.TransactOpts, x)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterTransactor) TokenizeYield(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "tokenizeYield", f, u, m, a, t)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterSession) TokenizeYield(f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYield(&_PendleRouter.TransactOpts, f, u, m, a, t)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterTransactorSession) TokenizeYield(f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYield(&_PendleRouter.TransactOpts, f, u, m, a, t)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterTransactor) TokenizeYieldReturns(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "tokenizeYieldReturns", x)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterSession) TokenizeYieldReturns(x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYieldReturns(&_PendleRouter.TransactOpts, x)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterTransactorSession) TokenizeYieldReturns(x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYieldReturns(&_PendleRouter.TransactOpts, x)
}
