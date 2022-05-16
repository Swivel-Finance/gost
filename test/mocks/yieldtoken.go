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

// YieldTokenABI is the input ABI used to generate the binding from.
const YieldTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"unlockTimestampReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawPrincipalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// YieldTokenBin is the compiled bytecode used for deploying new contracts.
var YieldTokenBin = "0x608060405234801561001057600080fd5b5061092d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636f307dc31161008c578063aa082a9d11610066578063aa082a9d14610223578063dea1a7e214610241578063e541efa21461025f578063e7ba677414610290576100cf565b80636f307dc3146101b957806370a08231146101d7578063831e574f14610207576100cf565b80630e6dfcd5146100d457806319e9be3a146100f057806323b872dd1461012157806339100838146101515780633b1988dc1461016d5780636521b96a1461019d575b600080fd5b6100ee60048036038101906100e9919061073e565b6102ac565b005b61010a60048036038101906101059190610791565b610390565b6040516101189291906107dc565b60405180910390f35b61013b6004803603810190610136919061073e565b6103d4565b6040516101489190610820565b60405180910390f35b61016b6004803603810190610166919061083b565b6104ce565b005b61018760048036038101906101829190610791565b6104d8565b6040516101949190610868565b60405180910390f35b6101b760048036038101906101b291906108af565b6104f0565b005b6101c161050d565b6040516101ce91906108dc565b60405180910390f35b6101f160048036038101906101ec9190610791565b610537565b6040516101fe9190610868565b60405180910390f35b610221600480360381019061021c919061083b565b610584565b005b61022b61058e565b6040516102389190610868565b60405180910390f35b610249610597565b60405161025691906108dc565b60405180910390f35b61027960048036038101906102749190610791565b6105bd565b6040516102879291906107dc565b60405180910390f35b6102aa60048036038101906102a59190610791565b610601565b005b6102b4610645565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508181602001818152505080600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015590505050505050565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b60006103de610675565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260149054906101000a900460ff169150509392505050565b8060018190555050565b60056020528060005260406000206000915090505481565b80600260146101000a81548160ff02191690831515021790555050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001549050919050565b8060008190555050565b60008054905090565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106d5826106aa565b9050919050565b6106e5816106ca565b81146106f057600080fd5b50565b600081359050610702816106dc565b92915050565b6000819050919050565b61071b81610708565b811461072657600080fd5b50565b60008135905061073881610712565b92915050565b600080600060608486031215610757576107566106a5565b5b6000610765868287016106f3565b9350506020610776868287016106f3565b925050604061078786828701610729565b9150509250925092565b6000602082840312156107a7576107a66106a5565b5b60006107b5848285016106f3565b91505092915050565b6107c7816106ca565b82525050565b6107d681610708565b82525050565b60006040820190506107f160008301856107be565b6107fe60208301846107cd565b9392505050565b60008115159050919050565b61081a81610805565b82525050565b60006020820190506108356000830184610811565b92915050565b600060208284031215610851576108506106a5565b5b600061085f84828501610729565b91505092915050565b600060208201905061087d60008301846107cd565b92915050565b61088c81610805565b811461089757600080fd5b50565b6000813590506108a981610883565b92915050565b6000602082840312156108c5576108c46106a5565b5b60006108d38482850161089a565b91505092915050565b60006020820190506108f160008301846107be565b9291505056fea2646970667358221220206456213add4e22c262ab27fcd92c9af8ff6a70e053b513c2be334c034cf24d64736f6c634300080d0033"

// DeployYieldToken deploys a new Ethereum contract, binding an instance of YieldToken to it.
func DeployYieldToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YieldToken, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(YieldTokenBin), backend)
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

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_YieldToken *YieldTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_YieldToken *YieldTokenSession) BalanceOfCalled() (common.Address, error) {
	return _YieldToken.Contract.BalanceOfCalled(&_YieldToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_YieldToken *YieldTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _YieldToken.Contract.BalanceOfCalled(&_YieldToken.CallOpts)
}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenCaller) RedeemCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "redeemCalled", arg0)

	outstruct := new(struct {
		From   common.Address
		Amount *big.Int
	})

	outstruct.From = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenSession) RedeemCalled(arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	return _YieldToken.Contract.RedeemCalled(&_YieldToken.CallOpts, arg0)
}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenCallerSession) RedeemCalled(arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	return _YieldToken.Contract.RedeemCalled(&_YieldToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		From   common.Address
		Amount *big.Int
	})

	outstruct.From = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	return _YieldToken.Contract.TransferFromCalled(&_YieldToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address from, uint256 amount)
func (_YieldToken *YieldTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	From   common.Address
	Amount *big.Int
}, error) {
	return _YieldToken.Contract.TransferFromCalled(&_YieldToken.CallOpts, arg0)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YieldToken *YieldTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YieldToken *YieldTokenSession) Underlying() (common.Address, error) {
	return _YieldToken.Contract.Underlying(&_YieldToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YieldToken *YieldTokenCallerSession) Underlying() (common.Address, error) {
	return _YieldToken.Contract.Underlying(&_YieldToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_YieldToken *YieldTokenCaller) UnlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "unlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_YieldToken *YieldTokenSession) UnlockTimestamp() (*big.Int, error) {
	return _YieldToken.Contract.UnlockTimestamp(&_YieldToken.CallOpts)
}

// UnlockTimestamp is a free data retrieval call binding the contract method 0xaa082a9d.
//
// Solidity: function unlockTimestamp() view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) UnlockTimestamp() (*big.Int, error) {
	return _YieldToken.Contract.UnlockTimestamp(&_YieldToken.CallOpts)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenCaller) WithdrawPrincipalCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "withdrawPrincipalCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.WithdrawPrincipalCalled(&_YieldToken.CallOpts, arg0)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.WithdrawPrincipalCalled(&_YieldToken.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_YieldToken *YieldTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_YieldToken *YieldTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.BalanceOf(&_YieldToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_YieldToken *YieldTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.BalanceOf(&_YieldToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_YieldToken *YieldTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_YieldToken *YieldTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.BalanceOfReturns(&_YieldToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_YieldToken *YieldTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.BalanceOfReturns(&_YieldToken.TransactOpts, b)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address t, address f, uint256 a) returns()
func (_YieldToken *YieldTokenTransactor) Redeem(opts *bind.TransactOpts, t common.Address, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "redeem", t, f, a)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address t, address f, uint256 a) returns()
func (_YieldToken *YieldTokenSession) Redeem(t common.Address, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Redeem(&_YieldToken.TransactOpts, t, f, a)
}

// Redeem is a paid mutator transaction binding the contract method 0x0e6dfcd5.
//
// Solidity: function redeem(address t, address f, uint256 a) returns()
func (_YieldToken *YieldTokenTransactorSession) Redeem(t common.Address, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Redeem(&_YieldToken.TransactOpts, t, f, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_YieldToken *YieldTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_YieldToken *YieldTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFrom(&_YieldToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_YieldToken *YieldTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFrom(&_YieldToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_YieldToken *YieldTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_YieldToken *YieldTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFromReturns(&_YieldToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_YieldToken *YieldTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFromReturns(&_YieldToken.TransactOpts, b)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YieldToken *YieldTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YieldToken *YieldTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.UnderlyingReturns(&_YieldToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YieldToken *YieldTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.UnderlyingReturns(&_YieldToken.TransactOpts, a)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_YieldToken *YieldTokenTransactor) UnlockTimestampReturns(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "unlockTimestampReturns", u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_YieldToken *YieldTokenSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.UnlockTimestampReturns(&_YieldToken.TransactOpts, u)
}

// UnlockTimestampReturns is a paid mutator transaction binding the contract method 0x831e574f.
//
// Solidity: function unlockTimestampReturns(uint256 u) returns()
func (_YieldToken *YieldTokenTransactorSession) UnlockTimestampReturns(u *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.UnlockTimestampReturns(&_YieldToken.TransactOpts, u)
}
