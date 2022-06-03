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

// ElementTokenABI is the input ABI used to generate the binding from.
const ElementTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"unlockTimestampReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"}],\"name\":\"withdrawPrincipal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawPrincipalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ElementTokenBin is the compiled bytecode used for deploying new contracts.
var ElementTokenBin = "0x608060405234801561001057600080fd5b50610811806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063831e574f11610071578063831e574f1461019f578063884e17f3146101bb578063aa082a9d146101d7578063dea1a7e2146101f5578063e541efa214610213578063e7ba677414610244576100b4565b806323b872dd146100b957806339100838146100e95780633b1988dc146101055780636521b96a146101355780636f307dc31461015157806370a082311461016f575b600080fd5b6100d360048036038101906100ce91906105e2565b610260565b6040516100e09190610650565b60405180910390f35b61010360048036038101906100fe919061066b565b61035a565b005b61011f600480360381019061011a9190610698565b610364565b60405161012c91906106d4565b60405180910390f35b61014f600480360381019061014a919061071b565b61037c565b005b610159610399565b6040516101669190610757565b60405180910390f35b61018960048036038101906101849190610698565b6103c3565b60405161019691906106d4565b60405180910390f35b6101b960048036038101906101b4919061066b565b610410565b005b6101d560048036038101906101d09190610772565b61041a565b005b6101df610462565b6040516101ec91906106d4565b60405180910390f35b6101fd61046b565b60405161020a9190610757565b60405180910390f35b61022d60048036038101906102289190610698565b610491565b60405161023b9291906107b2565b60405180910390f35b61025e60048036038101906102599190610698565b6104d5565b005b600061026a610519565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260149054906101000a900460ff169150509392505050565b8060018190555050565b60056020528060005260406000206000915090505481565b80600260146101000a81548160ff02191690831515021790555050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001549050919050565b8060008190555050565b81600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60008054905090565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105798261054e565b9050919050565b6105898161056e565b811461059457600080fd5b50565b6000813590506105a681610580565b92915050565b6000819050919050565b6105bf816105ac565b81146105ca57600080fd5b50565b6000813590506105dc816105b6565b92915050565b6000806000606084860312156105fb576105fa610549565b5b600061060986828701610597565b935050602061061a86828701610597565b925050604061062b868287016105cd565b9150509250925092565b60008115159050919050565b61064a81610635565b82525050565b60006020820190506106656000830184610641565b92915050565b60006020828403121561068157610680610549565b5b600061068f848285016105cd565b91505092915050565b6000602082840312156106ae576106ad610549565b5b60006106bc84828501610597565b91505092915050565b6106ce816105ac565b82525050565b60006020820190506106e960008301846106c5565b92915050565b6106f881610635565b811461070357600080fd5b50565b600081359050610715816106ef565b92915050565b60006020828403121561073157610730610549565b5b600061073f84828501610706565b91505092915050565b6107518161056e565b82525050565b600060208201905061076c6000830184610748565b92915050565b6000806040838503121561078957610788610549565b5b6000610797858286016105cd565b92505060206107a885828601610597565b9150509250929050565b60006040820190506107c76000830185610748565b6107d460208301846106c5565b939250505056fea2646970667358221220d1344a9d5bb8e5f4c9bc73fed6fe53f153eff24a72ed0956a96114767e9cb46064736f6c634300080d0033"

// DeployElementToken deploys a new Ethereum contract, binding an instance of ElementToken to it.
func DeployElementToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ElementToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElementTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElementToken{ElementTokenCaller: ElementTokenCaller{contract: contract}, ElementTokenTransactor: ElementTokenTransactor{contract: contract}, ElementTokenFilterer: ElementTokenFilterer{contract: contract}}, nil
}

// ElementToken is an auto generated Go binding around an Ethereum contract.
type ElementToken struct {
	ElementTokenCaller     // Read-only binding to the contract
	ElementTokenTransactor // Write-only binding to the contract
	ElementTokenFilterer   // Log filterer for contract events
}

// ElementTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElementTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElementTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElementTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElementTokenSession struct {
	Contract     *ElementToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElementTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElementTokenCallerSession struct {
	Contract *ElementTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ElementTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElementTokenTransactorSession struct {
	Contract     *ElementTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ElementTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElementTokenRaw struct {
	Contract *ElementToken // Generic contract binding to access the raw methods on
}

// ElementTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElementTokenCallerRaw struct {
	Contract *ElementTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ElementTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElementTokenTransactorRaw struct {
	Contract *ElementTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElementToken creates a new instance of ElementToken, bound to a specific deployed contract.
func NewElementToken(address common.Address, backend bind.ContractBackend) (*ElementToken, error) {
	contract, err := bindElementToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElementToken{ElementTokenCaller: ElementTokenCaller{contract: contract}, ElementTokenTransactor: ElementTokenTransactor{contract: contract}, ElementTokenFilterer: ElementTokenFilterer{contract: contract}}, nil
}

// NewElementTokenCaller creates a new read-only instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenCaller(address common.Address, caller bind.ContractCaller) (*ElementTokenCaller, error) {
	contract, err := bindElementToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTokenCaller{contract: contract}, nil
}

// NewElementTokenTransactor creates a new write-only instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ElementTokenTransactor, error) {
	contract, err := bindElementToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTokenTransactor{contract: contract}, nil
}

// NewElementTokenFilterer creates a new log filterer instance of ElementToken, bound to a specific deployed contract.
func NewElementTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ElementTokenFilterer, error) {
	contract, err := bindElementToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElementTokenFilterer{contract: contract}, nil
}

// bindElementToken binds a generic wrapper to an already deployed contract.
func bindElementToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElementToken *ElementTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElementToken.Contract.ElementTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElementToken *ElementTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElementToken.Contract.ElementTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElementToken *ElementTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElementToken.Contract.ElementTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElementToken *ElementTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElementToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElementToken *ElementTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElementToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElementToken *ElementTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElementToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_ElementToken *ElementTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_ElementToken *ElementTokenSession) BalanceOfCalled() (common.Address, error) {
	return _ElementToken.Contract.BalanceOfCalled(&_ElementToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_ElementToken *ElementTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _ElementToken.Contract.BalanceOfCalled(&_ElementToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_ElementToken *ElementTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_ElementToken *ElementTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _ElementToken.Contract.TransferFromCalled(&_ElementToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_ElementToken *ElementTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _ElementToken.Contract.TransferFromCalled(&_ElementToken.CallOpts, arg0)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenSession) Underlying() (common.Address, error) {
	return _ElementToken.Contract.Underlying(&_ElementToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ElementToken *ElementTokenCallerSession) Underlying() (common.Address, error) {
	return _ElementToken.Contract.Underlying(&_ElementToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenCaller) UnlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenSession) UnlockTimestamp() (*big.Int, error) {
	return _ElementToken.Contract.UnlockTimestamp(&_ElementToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_ElementToken *ElementTokenCallerSession) UnlockTimestamp() (*big.Int, error) {
	return _ElementToken.Contract.UnlockTimestamp(&_ElementToken.CallOpts)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_ElementToken *ElementTokenCaller) WithdrawPrincipalCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ElementToken.contract.Call(opts, &out, "withdrawPrincipalCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_ElementToken *ElementTokenSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _ElementToken.Contract.WithdrawPrincipalCalled(&_ElementToken.CallOpts, arg0)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_ElementToken *ElementTokenCallerSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _ElementToken.Contract.WithdrawPrincipalCalled(&_ElementToken.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_ElementToken *ElementTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_ElementToken *ElementTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.BalanceOf(&_ElementToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_ElementToken *ElementTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.BalanceOf(&_ElementToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_ElementToken *ElementTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_ElementToken *ElementTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.BalanceOfReturns(&_ElementToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_ElementToken *ElementTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.BalanceOfReturns(&_ElementToken.TransactOpts, b)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ElementToken *ElementTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ElementToken *ElementTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.TransferFrom(&_ElementToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ElementToken *ElementTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.TransferFrom(&_ElementToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ElementToken *ElementTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ElementToken *ElementTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _ElementToken.Contract.TransferFromReturns(&_ElementToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ElementToken *ElementTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _ElementToken.Contract.TransferFromReturns(&_ElementToken.TransactOpts, b)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.UnderlyingReturns(&_ElementToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_ElementToken *ElementTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.UnderlyingReturns(&_ElementToken.TransactOpts, a)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenTransactor) UnlockTimestampReturns(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "unlockTimestampReturns", u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.UnlockTimestampReturns(&_ElementToken.TransactOpts, u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_ElementToken *ElementTokenTransactorSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _ElementToken.Contract.UnlockTimestampReturns(&_ElementToken.TransactOpts, u)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_ElementToken *ElementTokenTransactor) WithdrawPrincipal(opts *bind.TransactOpts, a *big.Int, d common.Address) (*types.Transaction, error) {
	return _ElementToken.contract.Transact(opts, "withdrawPrincipal", a, d)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_ElementToken *ElementTokenSession) WithdrawPrincipal(a *big.Int, d common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.WithdrawPrincipal(&_ElementToken.TransactOpts, a, d)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_ElementToken *ElementTokenTransactorSession) WithdrawPrincipal(a *big.Int, d common.Address) (*types.Transaction, error) {
	return _ElementToken.Contract.WithdrawPrincipal(&_ElementToken.TransactOpts, a, d)
}
