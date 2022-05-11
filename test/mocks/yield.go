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

// YieldABI is the input ABI used to generate the binding from.
const YieldABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"i\",\"type\":\"address\"}],\"name\":\"baseReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellBaseCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBasePreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellBasePreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBasePreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBaseReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// YieldBin is the compiled bytecode used for deploying new contracts.
var YieldBin = "0x608060405234801561001057600080fd5b50610920806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063204f83f91161008c5780635001f3b5116100665780635001f3b5146102005780635e5dc52d1461021e578063808c58dd1461023a578063bcc1694f14610256576100cf565b8063204f83f9146101a85780633ef943bb146101c6578063439ec67d146101e4576100cf565b80631266b30c146100d457806313e7bc8c146100f257806318f154801461012257806319808486146101405780631cc52d77146101705780631e9a69501461018c575b600080fd5b6100dc610286565b6040516100e99190610564565b60405180910390f35b61010c600480360381019061010791906105b0565b6102a8565b6040516101199190610564565b60405180910390f35b61012a610309565b604051610137919061061e565b60405180910390f35b61015a60048036038101906101559190610665565b61032f565b60405161016791906106ab565b60405180910390f35b61018a60048036038101906101859190610702565b610347565b005b6101a660048036038101906101a1919061075b565b61036a565b005b6101b06103b6565b6040516101bd91906107aa565b60405180910390f35b6101ce6103cf565b6040516101db91906106ab565b60405180910390f35b6101fe60048036038101906101f99190610803565b6103d5565b005b610208610419565b604051610215919061088f565b60405180910390f35b610238600480360381019061023391906105b0565b610442565b005b610254600480360381019061024f91906105b0565b61047e565b005b610270600480360381019061026b91906108aa565b6104ba565b60405161027d9190610564565b60405180910390f35b600360009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600160109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090505481565b806000806101000a81548163ffffffff021916908363ffffffff16021790555050565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806005819055505050565b60008060009054906101000a900463ffffffff16905090565b60055481565b80600060046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b6000816fffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b61055e81610539565b82525050565b60006020820190506105796000830184610555565b92915050565b600080fd5b61058d81610539565b811461059857600080fd5b50565b6000813590506105aa81610584565b92915050565b6000602082840312156105c6576105c561057f565b5b60006105d48482850161059b565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610608826105dd565b9050919050565b610618816105fd565b82525050565b6000602082019050610633600083018461060f565b92915050565b610642816105fd565b811461064d57600080fd5b50565b60008135905061065f81610639565b92915050565b60006020828403121561067b5761067a61057f565b5b600061068984828501610650565b91505092915050565b6000819050919050565b6106a581610692565b82525050565b60006020820190506106c0600083018461069c565b92915050565b600063ffffffff82169050919050565b6106df816106c6565b81146106ea57600080fd5b50565b6000813590506106fc816106d6565b92915050565b6000602082840312156107185761071761057f565b5b6000610726848285016106ed565b91505092915050565b61073881610692565b811461074357600080fd5b50565b6000813590506107558161072f565b92915050565b600080604083850312156107725761077161057f565b5b600061078085828601610650565b925050602061079185828601610746565b9150509250929050565b6107a4816106c6565b82525050565b60006020820190506107bf600083018461079b565b92915050565b60006107d0826105fd565b9050919050565b6107e0816107c5565b81146107eb57600080fd5b50565b6000813590506107fd816107d7565b92915050565b6000602082840312156108195761081861057f565b5b6000610827848285016107ee565b91505092915050565b6000819050919050565b600061085561085061084b846105dd565b610830565b6105dd565b9050919050565b60006108678261083a565b9050919050565b60006108798261085c565b9050919050565b6108898161086e565b82525050565b60006020820190506108a46000830184610880565b92915050565b600080604083850312156108c1576108c061057f565b5b60006108cf85828601610650565b92505060206108e08582860161059b565b915050925092905056fea2646970667358221220fe390ae78f1cb7ccf09abb1699028e12a5da948ba12c59d6fdb2816d4c96700164736f6c634300080d0033"

// DeployYield deploys a new Ethereum contract, binding an instance of Yield to it.
func DeployYield(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Yield, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(YieldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Yield{YieldCaller: YieldCaller{contract: contract}, YieldTransactor: YieldTransactor{contract: contract}, YieldFilterer: YieldFilterer{contract: contract}}, nil
}

// Yield is an auto generated Go binding around an Ethereum contract.
type Yield struct {
	YieldCaller     // Read-only binding to the contract
	YieldTransactor // Write-only binding to the contract
	YieldFilterer   // Log filterer for contract events
}

// YieldCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldSession struct {
	Contract     *Yield            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldCallerSession struct {
	Contract *YieldCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YieldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldTransactorSession struct {
	Contract     *YieldTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldRaw struct {
	Contract *Yield // Generic contract binding to access the raw methods on
}

// YieldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldCallerRaw struct {
	Contract *YieldCaller // Generic read-only contract binding to access the raw methods on
}

// YieldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldTransactorRaw struct {
	Contract *YieldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYield creates a new instance of Yield, bound to a specific deployed contract.
func NewYield(address common.Address, backend bind.ContractBackend) (*Yield, error) {
	contract, err := bindYield(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yield{YieldCaller: YieldCaller{contract: contract}, YieldTransactor: YieldTransactor{contract: contract}, YieldFilterer: YieldFilterer{contract: contract}}, nil
}

// NewYieldCaller creates a new read-only instance of Yield, bound to a specific deployed contract.
func NewYieldCaller(address common.Address, caller bind.ContractCaller) (*YieldCaller, error) {
	contract, err := bindYield(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldCaller{contract: contract}, nil
}

// NewYieldTransactor creates a new write-only instance of Yield, bound to a specific deployed contract.
func NewYieldTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldTransactor, error) {
	contract, err := bindYield(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTransactor{contract: contract}, nil
}

// NewYieldFilterer creates a new log filterer instance of Yield, bound to a specific deployed contract.
func NewYieldFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldFilterer, error) {
	contract, err := bindYield(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldFilterer{contract: contract}, nil
}

// bindYield binds a generic wrapper to an already deployed contract.
func bindYield(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yield *YieldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yield.Contract.YieldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yield *YieldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yield.Contract.YieldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yield *YieldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yield.Contract.YieldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yield *YieldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yield.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yield *YieldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yield.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yield *YieldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yield.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Yield *YieldCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Yield *YieldSession) AmountCalled() (*big.Int, error) {
	return _Yield.Contract.AmountCalled(&_Yield.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Yield *YieldCallerSession) AmountCalled() (*big.Int, error) {
	return _Yield.Contract.AmountCalled(&_Yield.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldSession) Base() (common.Address, error) {
	return _Yield.Contract.Base(&_Yield.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldCallerSession) Base() (common.Address, error) {
	return _Yield.Contract.Base(&_Yield.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldCaller) Maturity(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldSession) Maturity() (uint32, error) {
	return _Yield.Contract.Maturity(&_Yield.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldCallerSession) Maturity() (uint32, error) {
	return _Yield.Contract.Maturity(&_Yield.CallOpts)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldCaller) SellBaseCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "sellBaseCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _Yield.Contract.SellBaseCalled(&_Yield.CallOpts, arg0)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldCallerSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _Yield.Contract.SellBaseCalled(&_Yield.CallOpts, arg0)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldCaller) SellBasePreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "sellBasePreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldSession) SellBasePreviewCalled() (*big.Int, error) {
	return _Yield.Contract.SellBasePreviewCalled(&_Yield.CallOpts)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldCallerSession) SellBasePreviewCalled() (*big.Int, error) {
	return _Yield.Contract.SellBasePreviewCalled(&_Yield.CallOpts)
}

// TokenCalled is a free data retrieval call binding the contract method 0x18f15480.
//
// Solidity: function tokenCalled() view returns(address)
func (_Yield *YieldCaller) TokenCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "tokenCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenCalled is a free data retrieval call binding the contract method 0x18f15480.
//
// Solidity: function tokenCalled() view returns(address)
func (_Yield *YieldSession) TokenCalled() (common.Address, error) {
	return _Yield.Contract.TokenCalled(&_Yield.CallOpts)
}

// TokenCalled is a free data retrieval call binding the contract method 0x18f15480.
//
// Solidity: function tokenCalled() view returns(address)
func (_Yield *YieldCallerSession) TokenCalled() (common.Address, error) {
	return _Yield.Contract.TokenCalled(&_Yield.CallOpts)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldTransactor) BaseReturns(opts *bind.TransactOpts, i common.Address) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "baseReturns", i)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldSession) BaseReturns(i common.Address) (*types.Transaction, error) {
	return _Yield.Contract.BaseReturns(&_Yield.TransactOpts, i)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldTransactorSession) BaseReturns(i common.Address) (*types.Transaction, error) {
	return _Yield.Contract.BaseReturns(&_Yield.TransactOpts, i)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldTransactor) MaturityReturns(opts *bind.TransactOpts, m uint32) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "maturityReturns", m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldSession) MaturityReturns(m uint32) (*types.Transaction, error) {
	return _Yield.Contract.MaturityReturns(&_Yield.TransactOpts, m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldTransactorSession) MaturityReturns(m uint32) (*types.Transaction, error) {
	return _Yield.Contract.MaturityReturns(&_Yield.TransactOpts, m)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address t, uint256 a) returns()
func (_Yield *YieldTransactor) Redeem(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "redeem", t, a)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address t, uint256 a) returns()
func (_Yield *YieldSession) Redeem(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.Redeem(&_Yield.TransactOpts, t, a)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address t, uint256 a) returns()
func (_Yield *YieldTransactorSession) Redeem(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.Redeem(&_Yield.TransactOpts, t, a)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldTransactor) SellBase(opts *bind.TransactOpts, a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBase", a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBase(&_Yield.TransactOpts, a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldTransactorSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBase(&_Yield.TransactOpts, a, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldTransactor) SellBasePreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBasePreview", u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreview(&_Yield.TransactOpts, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldTransactorSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreview(&_Yield.TransactOpts, u)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldTransactor) SellBasePreviewReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBasePreviewReturns", r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreviewReturns(&_Yield.TransactOpts, r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldTransactorSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreviewReturns(&_Yield.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldTransactor) SellBaseReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBaseReturns", r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBaseReturns(&_Yield.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldTransactorSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBaseReturns(&_Yield.TransactOpts, r)
}
