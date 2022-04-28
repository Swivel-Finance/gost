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

// YieldTokenMetaData contains all meta data concerning the YieldToken contract.
var YieldTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellBaseCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBasePreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellBasePreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBasePreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBaseReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610610806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80635001f3b51161005b5780635001f3b5146101295780635e5dc52d14610147578063808c58dd14610163578063bcc1694f1461017f57610088565b80631266b30c1461008d57806313e7bc8c146100ab57806319808486146100db578063204f83f91461010b575b600080fd5b6100956101af565b6040516100a291906103ae565b60405180910390f35b6100c560048036038101906100c091906103fa565b6101d1565b6040516100d291906103ae565b60405180910390f35b6100f560048036038101906100f09190610485565b610232565b60405161010291906104cb565b60405180910390f35b61011361024a565b6040516101209190610505565b60405180910390f35b610131610263565b60405161013e919061057f565b60405180910390f35b610161600480360381019061015c91906103fa565b61028c565b005b61017d600480360381019061017891906103fa565b6102c8565b005b6101996004803603810190610194919061059a565b610304565b6040516101a691906103ae565b60405180910390f35b600360009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600160109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b60026020528060005260406000206000915090505481565b60008060009054906101000a900463ffffffff16905090565b60008060049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b6000816fffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b6103a881610383565b82525050565b60006020820190506103c3600083018461039f565b92915050565b600080fd5b6103d781610383565b81146103e257600080fd5b50565b6000813590506103f4816103ce565b92915050565b6000602082840312156104105761040f6103c9565b5b600061041e848285016103e5565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061045282610427565b9050919050565b61046281610447565b811461046d57600080fd5b50565b60008135905061047f81610459565b92915050565b60006020828403121561049b5761049a6103c9565b5b60006104a984828501610470565b91505092915050565b6000819050919050565b6104c5816104b2565b82525050565b60006020820190506104e060008301846104bc565b92915050565b600063ffffffff82169050919050565b6104ff816104e6565b82525050565b600060208201905061051a60008301846104f6565b92915050565b6000819050919050565b600061054561054061053b84610427565b610520565b610427565b9050919050565b60006105578261052a565b9050919050565b60006105698261054c565b9050919050565b6105798161055e565b82525050565b60006020820190506105946000830184610570565b92915050565b600080604083850312156105b1576105b06103c9565b5b60006105bf85828601610470565b92505060206105d0858286016103e5565b915050925092905056fea2646970667358221220d74db6c94247c4b58a1ce034e78d605fb5917d348c551425a7532fcdf8a6fe4564736f6c634300080d0033",
}

// YieldTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use YieldTokenMetaData.ABI instead.
var YieldTokenABI = YieldTokenMetaData.ABI

// YieldTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YieldTokenMetaData.Bin instead.
var YieldTokenBin = YieldTokenMetaData.Bin

// DeployYieldToken deploys a new Ethereum contract, binding an instance of YieldToken to it.
func DeployYieldToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YieldToken, error) {
	parsed, err := YieldTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YieldTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YieldToken{YieldTokenCaller: YieldTokenCaller{contract: contract}, YieldTokenTransactor: YieldTokenTransactor{contract: contract}, YieldTokenFilterer: YieldTokenFilterer{contract: contract}}, nil
}

// YieldToken is an auto generated Go binding around an Ethereum contract.
type YieldToken struct {
	YieldTokenCaller     // Read-only binding to the contract
	YieldTokenTransactor // Write-only binding to the contract
	YieldTokenFilterer   // Log filterer for contract events
}

// YieldTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldTokenSession struct {
	Contract     *YieldToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldTokenCallerSession struct {
	Contract *YieldTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YieldTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldTokenTransactorSession struct {
	Contract     *YieldTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YieldTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldTokenRaw struct {
	Contract *YieldToken // Generic contract binding to access the raw methods on
}

// YieldTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldTokenCallerRaw struct {
	Contract *YieldTokenCaller // Generic read-only contract binding to access the raw methods on
}

// YieldTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldTokenTransactorRaw struct {
	Contract *YieldTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldToken creates a new instance of YieldToken, bound to a specific deployed contract.
func NewYieldToken(address common.Address, backend bind.ContractBackend) (*YieldToken, error) {
	contract, err := bindYieldToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldToken{YieldTokenCaller: YieldTokenCaller{contract: contract}, YieldTokenTransactor: YieldTokenTransactor{contract: contract}, YieldTokenFilterer: YieldTokenFilterer{contract: contract}}, nil
}

// NewYieldTokenCaller creates a new read-only instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenCaller(address common.Address, caller bind.ContractCaller) (*YieldTokenCaller, error) {
	contract, err := bindYieldToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTokenCaller{contract: contract}, nil
}

// NewYieldTokenTransactor creates a new write-only instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldTokenTransactor, error) {
	contract, err := bindYieldToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTokenTransactor{contract: contract}, nil
}

// NewYieldTokenFilterer creates a new log filterer instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldTokenFilterer, error) {
	contract, err := bindYieldToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldTokenFilterer{contract: contract}, nil
}

// bindYieldToken binds a generic wrapper to an already deployed contract.
func bindYieldToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldToken *YieldTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldToken.Contract.YieldTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldToken *YieldTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldToken.Contract.YieldTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldToken *YieldTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldToken.Contract.YieldTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldToken *YieldTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldToken *YieldTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldToken *YieldTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldToken.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_YieldToken *YieldTokenCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_YieldToken *YieldTokenSession) Base() (common.Address, error) {
	return _YieldToken.Contract.Base(&_YieldToken.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_YieldToken *YieldTokenCallerSession) Base() (common.Address, error) {
	return _YieldToken.Contract.Base(&_YieldToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_YieldToken *YieldTokenCaller) Maturity(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_YieldToken *YieldTokenSession) Maturity() (uint32, error) {
	return _YieldToken.Contract.Maturity(&_YieldToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_YieldToken *YieldTokenCallerSession) Maturity() (uint32, error) {
	return _YieldToken.Contract.Maturity(&_YieldToken.CallOpts)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenCaller) SellBaseCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "sellBaseCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.SellBaseCalled(&_YieldToken.CallOpts, arg0)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.SellBaseCalled(&_YieldToken.CallOpts, arg0)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_YieldToken *YieldTokenCaller) SellBasePreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "sellBasePreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_YieldToken *YieldTokenSession) SellBasePreviewCalled() (*big.Int, error) {
	return _YieldToken.Contract.SellBasePreviewCalled(&_YieldToken.CallOpts)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_YieldToken *YieldTokenCallerSession) SellBasePreviewCalled() (*big.Int, error) {
	return _YieldToken.Contract.SellBasePreviewCalled(&_YieldToken.CallOpts)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_YieldToken *YieldTokenTransactor) SellBase(opts *bind.TransactOpts, a common.Address, u *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "sellBase", a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_YieldToken *YieldTokenSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBase(&_YieldToken.TransactOpts, a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_YieldToken *YieldTokenTransactorSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBase(&_YieldToken.TransactOpts, a, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_YieldToken *YieldTokenTransactor) SellBasePreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "sellBasePreview", u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_YieldToken *YieldTokenSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBasePreview(&_YieldToken.TransactOpts, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_YieldToken *YieldTokenTransactorSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBasePreview(&_YieldToken.TransactOpts, u)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_YieldToken *YieldTokenTransactor) SellBasePreviewReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "sellBasePreviewReturns", r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_YieldToken *YieldTokenSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBasePreviewReturns(&_YieldToken.TransactOpts, r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_YieldToken *YieldTokenTransactorSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBasePreviewReturns(&_YieldToken.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_YieldToken *YieldTokenTransactor) SellBaseReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "sellBaseReturns", r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_YieldToken *YieldTokenSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBaseReturns(&_YieldToken.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_YieldToken *YieldTokenTransactorSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.SellBaseReturns(&_YieldToken.TransactOpts, r)
}
