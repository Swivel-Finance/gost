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

// MarketPlaceMetaData contains all meta data concerning the MarketPlace contract.
var MarketPlaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address[8]\",\"name\":\"\",\"type\":\"address[8]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[8]\",\"name\":\"m\",\"type\":\"address[8]\"}],\"name\":\"marketsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061050c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806317b3bba7146100465780632db46d2814610076578063f9ce9f14146100a6575b600080fd5b610060600480360381019061005b9190610323565b6100c2565b60405161006d919061040e565b60405180910390f35b610090600480360381019061008b919061042a565b61018a565b60405161009d9190610466565b60405180910390f35b6100c060048036038101906100bb91906104a8565b6101a2565b005b6100ca6101b7565b81600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600060088060200260405190810160405280929190826008801561017d576020028201915b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610133575b5050505050905092915050565b60086020528060005260406000206000915090505481565b8060009060086101b39291906101da565b5050565b604051806101000160405280600890602082028036833780820191505090505090565b826008810192821561025c579160200282015b8281111561025b57823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906101ed565b5b509050610269919061026d565b5090565b5b8082111561028657600081600090555060010161026e565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102ba8261028f565b9050919050565b6102ca816102af565b81146102d557600080fd5b50565b6000813590506102e7816102c1565b92915050565b6000819050919050565b610300816102ed565b811461030b57600080fd5b50565b60008135905061031d816102f7565b92915050565b6000806040838503121561033a5761033961028a565b5b6000610348858286016102d8565b92505060206103598582860161030e565b9150509250929050565b600060089050919050565b600081905092915050565b6000819050919050565b61038c816102af565b82525050565b600061039e8383610383565b60208301905092915050565b6000602082019050919050565b6103c081610363565b6103ca818461036e565b92506103d582610379565b8060005b838110156104065781516103ed8782610392565b96506103f8836103aa565b9250506001810190506103d9565b505050505050565b60006101008201905061042460008301846103b7565b92915050565b6000602082840312156104405761043f61028a565b5b600061044e848285016102d8565b91505092915050565b610460816102ed565b82525050565b600060208201905061047b6000830184610457565b92915050565b600080fd5b6000819050826020600802820111156104a2576104a1610481565b5b92915050565b600061010082840312156104bf576104be61028a565b5b60006104cd84828501610486565b9150509291505056fea26469706673582212204872a24b15420a5d1dd3f9d6a0a44aa2b66c5d83e3afc114d0f2c19f58213faa64736f6c634300080d0033",
}

// MarketPlaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketPlaceMetaData.ABI instead.
var MarketPlaceABI = MarketPlaceMetaData.ABI

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketPlaceMetaData.Bin instead.
var MarketPlaceBin = MarketPlaceMetaData.Bin

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := MarketPlaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketPlaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) MarketsCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "marketsCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_MarketPlace *MarketPlaceTransactor) Markets(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "markets", u, m)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_MarketPlace *MarketPlaceSession) Markets(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.TransactOpts, u, m)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_MarketPlace *MarketPlaceTransactorSession) Markets(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.TransactOpts, u, m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_MarketPlace *MarketPlaceTransactor) MarketsReturns(opts *bind.TransactOpts, m [8]common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "marketsReturns", m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_MarketPlace *MarketPlaceSession) MarketsReturns(m [8]common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketsReturns(&_MarketPlace.TransactOpts, m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MarketsReturns(m [8]common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketsReturns(&_MarketPlace.TransactOpts, m)
}
