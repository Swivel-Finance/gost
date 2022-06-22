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

// TempusTokenMetaData contains all meta data concerning the TempusToken contract.
var TempusTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610641806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636521b96a1161005b5780636521b96a146100fe57806370a082311461011a578063dea1a7e21461014a578063e541efa2146101685761007d565b806323b872dd1461008257806327e235e3146100b257806339100838146100e2575b600080fd5b61009c60048036038101906100979190610452565b610199565b6040516100a991906104c0565b60405180910390f35b6100cc60048036038101906100c791906104db565b610293565b6040516100d99190610517565b60405180910390f35b6100fc60048036038101906100f79190610532565b6102ab565b005b6101186004803603810190610113919061058b565b6102b5565b005b610134600480360381019061012f91906104db565b6102d2565b6040516101419190610517565b60405180910390f35b61015261031f565b60405161015f91906105c7565b60405180910390f35b610182600480360381019061017d91906104db565b610345565b6040516101909291906105e2565b60405180910390f35b60006101a3610389565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600160009054906101000a900460ff169150509392505050565b60006020528060005260406000206000915090505481565b8060028190555050565b80600160006101000a81548160ff02191690831515021790555050565b600081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002549050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103e9826103be565b9050919050565b6103f9816103de565b811461040457600080fd5b50565b600081359050610416816103f0565b92915050565b6000819050919050565b61042f8161041c565b811461043a57600080fd5b50565b60008135905061044c81610426565b92915050565b60008060006060848603121561046b5761046a6103b9565b5b600061047986828701610407565b935050602061048a86828701610407565b925050604061049b8682870161043d565b9150509250925092565b60008115159050919050565b6104ba816104a5565b82525050565b60006020820190506104d560008301846104b1565b92915050565b6000602082840312156104f1576104f06103b9565b5b60006104ff84828501610407565b91505092915050565b6105118161041c565b82525050565b600060208201905061052c6000830184610508565b92915050565b600060208284031215610548576105476103b9565b5b60006105568482850161043d565b91505092915050565b610568816104a5565b811461057357600080fd5b50565b6000813590506105858161055f565b92915050565b6000602082840312156105a1576105a06103b9565b5b60006105af84828501610576565b91505092915050565b6105c1816103de565b82525050565b60006020820190506105dc60008301846105b8565b92915050565b60006040820190506105f760008301856105b8565b6106046020830184610508565b939250505056fea264697066735822122000fe7fc75d61a309502a52cff3c0fa8b31be9010fc095c6dad693e489966c71964736f6c634300080d0033",
}

// TempusTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TempusTokenMetaData.ABI instead.
var TempusTokenABI = TempusTokenMetaData.ABI

// TempusTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TempusTokenMetaData.Bin instead.
var TempusTokenBin = TempusTokenMetaData.Bin

// DeployTempusToken deploys a new Ethereum contract, binding an instance of TempusToken to it.
func DeployTempusToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TempusToken, error) {
	parsed, err := TempusTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TempusTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TempusToken{TempusTokenCaller: TempusTokenCaller{contract: contract}, TempusTokenTransactor: TempusTokenTransactor{contract: contract}, TempusTokenFilterer: TempusTokenFilterer{contract: contract}}, nil
}

// TempusToken is an auto generated Go binding around an Ethereum contract.
type TempusToken struct {
	TempusTokenCaller     // Read-only binding to the contract
	TempusTokenTransactor // Write-only binding to the contract
	TempusTokenFilterer   // Log filterer for contract events
}

// TempusTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TempusTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TempusTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TempusTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TempusTokenSession struct {
	Contract     *TempusToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TempusTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TempusTokenCallerSession struct {
	Contract *TempusTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TempusTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TempusTokenTransactorSession struct {
	Contract     *TempusTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TempusTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TempusTokenRaw struct {
	Contract *TempusToken // Generic contract binding to access the raw methods on
}

// TempusTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TempusTokenCallerRaw struct {
	Contract *TempusTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TempusTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TempusTokenTransactorRaw struct {
	Contract *TempusTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTempusToken creates a new instance of TempusToken, bound to a specific deployed contract.
func NewTempusToken(address common.Address, backend bind.ContractBackend) (*TempusToken, error) {
	contract, err := bindTempusToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TempusToken{TempusTokenCaller: TempusTokenCaller{contract: contract}, TempusTokenTransactor: TempusTokenTransactor{contract: contract}, TempusTokenFilterer: TempusTokenFilterer{contract: contract}}, nil
}

// NewTempusTokenCaller creates a new read-only instance of TempusToken, bound to a specific deployed contract.
func NewTempusTokenCaller(address common.Address, caller bind.ContractCaller) (*TempusTokenCaller, error) {
	contract, err := bindTempusToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TempusTokenCaller{contract: contract}, nil
}

// NewTempusTokenTransactor creates a new write-only instance of TempusToken, bound to a specific deployed contract.
func NewTempusTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TempusTokenTransactor, error) {
	contract, err := bindTempusToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TempusTokenTransactor{contract: contract}, nil
}

// NewTempusTokenFilterer creates a new log filterer instance of TempusToken, bound to a specific deployed contract.
func NewTempusTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TempusTokenFilterer, error) {
	contract, err := bindTempusToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TempusTokenFilterer{contract: contract}, nil
}

// bindTempusToken binds a generic wrapper to an already deployed contract.
func bindTempusToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TempusToken *TempusTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TempusToken.Contract.TempusTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TempusToken *TempusTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TempusToken.Contract.TempusTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TempusToken *TempusTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TempusToken.Contract.TempusTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TempusToken *TempusTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TempusToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TempusToken *TempusTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TempusToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TempusToken *TempusTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TempusToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_TempusToken *TempusTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TempusToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_TempusToken *TempusTokenSession) BalanceOfCalled() (common.Address, error) {
	return _TempusToken.Contract.BalanceOfCalled(&_TempusToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_TempusToken *TempusTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _TempusToken.Contract.BalanceOfCalled(&_TempusToken.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TempusToken *TempusTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TempusToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TempusToken *TempusTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TempusToken.Contract.Balances(&_TempusToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TempusToken *TempusTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TempusToken.Contract.Balances(&_TempusToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_TempusToken *TempusTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _TempusToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_TempusToken *TempusTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _TempusToken.Contract.TransferFromCalled(&_TempusToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_TempusToken *TempusTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _TempusToken.Contract.TransferFromCalled(&_TempusToken.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_TempusToken *TempusTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _TempusToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_TempusToken *TempusTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOf(&_TempusToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_TempusToken *TempusTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOf(&_TempusToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_TempusToken *TempusTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _TempusToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_TempusToken *TempusTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOfReturns(&_TempusToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_TempusToken *TempusTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOfReturns(&_TempusToken.TransactOpts, b)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_TempusToken *TempusTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _TempusToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_TempusToken *TempusTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.TransferFrom(&_TempusToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_TempusToken *TempusTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.TransferFrom(&_TempusToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_TempusToken *TempusTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _TempusToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_TempusToken *TempusTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _TempusToken.Contract.TransferFromReturns(&_TempusToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_TempusToken *TempusTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _TempusToken.Contract.TransferFromReturns(&_TempusToken.TransactOpts, b)
}