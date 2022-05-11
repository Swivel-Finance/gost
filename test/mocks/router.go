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

// RouterABI is the input ABI used to generate the binding from.
const RouterABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipientCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"redeemToBacking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldAmountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RouterBin is the compiled bytecode used for deploying new contracts.
var RouterBin = "0x608060405234801561001057600080fd5b50610365806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633ef943bb1461005c5780635d33d7081461007a5780636c8d4fa114610098578063a2f428f7146100b4578063e62bfee1146100d2575b600080fd5b6100646100f0565b60405161007191906101f4565b60405180910390f35b6100826100f6565b60405161008f9190610250565b60405180910390f35b6100b260048036038101906100ad91906102c8565b61011c565b005b6100bc6101b1565b6040516100c99190610250565b60405180910390f35b6100da6101d5565b6040516100e791906101f4565b60405180910390f35b60015481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826001819055508160028190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b6000819050919050565b6101ee816101db565b82525050565b600060208201905061020960008301846101e5565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061023a8261020f565b9050919050565b61024a8161022f565b82525050565b60006020820190506102656000830184610241565b92915050565b600080fd5b6102798161022f565b811461028457600080fd5b50565b60008135905061029681610270565b92915050565b6102a5816101db565b81146102b057600080fd5b50565b6000813590506102c28161029c565b92915050565b600080600080608085870312156102e2576102e161026b565b5b60006102f087828801610287565b9450506020610301878288016102b3565b9350506040610312878288016102b3565b925050606061032387828801610287565b9150509295919450925056fea2646970667358221220bf2a893b333138da7d31033ba63e17cb9f76bcaf671908cbc2a71c8eb6dbbe9d64736f6c634300080d0033"

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Router *RouterCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Router *RouterSession) AmountCalled() (*big.Int, error) {
	return _Router.Contract.AmountCalled(&_Router.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Router *RouterCallerSession) AmountCalled() (*big.Int, error) {
	return _Router.Contract.AmountCalled(&_Router.CallOpts)
}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Router *RouterCaller) OwnerCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "ownerCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Router *RouterSession) OwnerCalled() (common.Address, error) {
	return _Router.Contract.OwnerCalled(&_Router.CallOpts)
}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Router *RouterCallerSession) OwnerCalled() (common.Address, error) {
	return _Router.Contract.OwnerCalled(&_Router.CallOpts)
}

// RecipientCalled is a free data retrieval call binding the contract method 0x5d33d708.
//
// Solidity: function recipientCalled() view returns(address)
func (_Router *RouterCaller) RecipientCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "recipientCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientCalled is a free data retrieval call binding the contract method 0x5d33d708.
//
// Solidity: function recipientCalled() view returns(address)
func (_Router *RouterSession) RecipientCalled() (common.Address, error) {
	return _Router.Contract.RecipientCalled(&_Router.CallOpts)
}

// RecipientCalled is a free data retrieval call binding the contract method 0x5d33d708.
//
// Solidity: function recipientCalled() view returns(address)
func (_Router *RouterCallerSession) RecipientCalled() (common.Address, error) {
	return _Router.Contract.RecipientCalled(&_Router.CallOpts)
}

// YieldAmountCalled is a free data retrieval call binding the contract method 0xe62bfee1.
//
// Solidity: function yieldAmountCalled() view returns(uint256)
func (_Router *RouterCaller) YieldAmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "yieldAmountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldAmountCalled is a free data retrieval call binding the contract method 0xe62bfee1.
//
// Solidity: function yieldAmountCalled() view returns(uint256)
func (_Router *RouterSession) YieldAmountCalled() (*big.Int, error) {
	return _Router.Contract.YieldAmountCalled(&_Router.CallOpts)
}

// YieldAmountCalled is a free data retrieval call binding the contract method 0xe62bfee1.
//
// Solidity: function yieldAmountCalled() view returns(uint256)
func (_Router *RouterCallerSession) YieldAmountCalled() (*big.Int, error) {
	return _Router.Contract.YieldAmountCalled(&_Router.CallOpts)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 a, uint256 y, address r) returns()
func (_Router *RouterTransactor) RedeemToBacking(opts *bind.TransactOpts, o common.Address, a *big.Int, y *big.Int, r common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "redeemToBacking", o, a, y, r)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 a, uint256 y, address r) returns()
func (_Router *RouterSession) RedeemToBacking(o common.Address, a *big.Int, y *big.Int, r common.Address) (*types.Transaction, error) {
	return _Router.Contract.RedeemToBacking(&_Router.TransactOpts, o, a, y, r)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 a, uint256 y, address r) returns()
func (_Router *RouterTransactorSession) RedeemToBacking(o common.Address, a *big.Int, y *big.Int, r common.Address) (*types.Transaction, error) {
	return _Router.Contract.RedeemToBacking(&_Router.TransactOpts, o, a, y, r)
}
