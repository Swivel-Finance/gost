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

// NotionalTokenABI is the input ABI used to generate the binding from.
const NotionalTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"getMaturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingToken\",\"outputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getUnderlyingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NotionalTokenBin is the compiled bytecode used for deploying new contracts.
var NotionalTokenBin = "0x608060405234801561001057600080fd5b50610915806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063a7d8cf031161008c578063e16695b511610066578063e16695b514610222578063e541efa214610240578063ee719bc814610271578063fe68267014610290576100cf565b8063a7d8cf03146101b8578063d35147e4146101d4578063dea1a7e214610204576100cf565b806323b872dd146100d457806326cca7861461010457806339100838146101205780636521b96a1461013c5780636e553f651461015857806370a0823114610188575b600080fd5b6100ee60048036038101906100e99190610645565b6102ac565b6040516100fb91906106b3565b60405180910390f35b61011e600480360381019061011991906106ce565b6103a6565b005b61013a600480360381019061013591906106ce565b6103b0565b005b61015660048036038101906101519190610727565b6103ba565b005b610172600480360381019061016d9190610754565b6103d7565b60405161017f91906107a3565b60405180910390f35b6101a2600480360381019061019d91906107be565b610428565b6040516101af91906107a3565b60405180910390f35b6101d260048036038101906101cd91906107be565b610475565b005b6101ee60048036038101906101e991906107be565b6104b8565b6040516101fb91906107a3565b60405180910390f35b61020c6104d0565b60405161021991906107fa565b60405180910390f35b61022a6104f6565b60405161023791906107a3565b60405180910390f35b61025a600480360381019061025591906107be565b610500565b604051610268929190610815565b60405180910390f35b610279610544565b6040516102879291906108b6565b60405180910390f35b6102aa60048036038101906102a591906106ce565b610572565b005b60006102b661057c565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260009054906101000a900460ff169150509392505050565b8060048190555050565b8060038190555050565b80600260006101000a81548160ff02191690831515021790555050565b600082600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600454905092915050565b600081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003549050919050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60066020528060005260406000206000915090505481565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000915091509091565b8060018190555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105dc826105b1565b9050919050565b6105ec816105d1565b81146105f757600080fd5b50565b600081359050610609816105e3565b92915050565b6000819050919050565b6106228161060f565b811461062d57600080fd5b50565b60008135905061063f81610619565b92915050565b60008060006060848603121561065e5761065d6105ac565b5b600061066c868287016105fa565b935050602061067d868287016105fa565b925050604061068e86828701610630565b9150509250925092565b60008115159050919050565b6106ad81610698565b82525050565b60006020820190506106c860008301846106a4565b92915050565b6000602082840312156106e4576106e36105ac565b5b60006106f284828501610630565b91505092915050565b61070481610698565b811461070f57600080fd5b50565b600081359050610721816106fb565b92915050565b60006020828403121561073d5761073c6105ac565b5b600061074b84828501610712565b91505092915050565b6000806040838503121561076b5761076a6105ac565b5b600061077985828601610630565b925050602061078a858286016105fa565b9150509250929050565b61079d8161060f565b82525050565b60006020820190506107b86000830184610794565b92915050565b6000602082840312156107d4576107d36105ac565b5b60006107e2848285016105fa565b91505092915050565b6107f4816105d1565b82525050565b600060208201905061080f60008301846107eb565b92915050565b600060408201905061082a60008301856107eb565b6108376020830184610794565b9392505050565b6000819050919050565b600061086361085e610859846105b1565b61083e565b6105b1565b9050919050565b600061087582610848565b9050919050565b60006108878261086a565b9050919050565b6108978161087c565b82525050565b6000819050919050565b6108b08161089d565b82525050565b60006040820190506108cb600083018561088e565b6108d860208301846108a7565b939250505056fea2646970667358221220b4a7cec62f7853ad80f84efb6a503b2b3a36bdc53c4da7b7ffeeb63da5b6056a64736f6c634300080d0033"

// DeployNotionalToken deploys a new Ethereum contract, binding an instance of NotionalToken to it.
func DeployNotionalToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NotionalToken, error) {
	parsed, err := abi.JSON(strings.NewReader(NotionalTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NotionalTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NotionalToken{NotionalTokenCaller: NotionalTokenCaller{contract: contract}, NotionalTokenTransactor: NotionalTokenTransactor{contract: contract}, NotionalTokenFilterer: NotionalTokenFilterer{contract: contract}}, nil
}

// NotionalToken is an auto generated Go binding around an Ethereum contract.
type NotionalToken struct {
	NotionalTokenCaller     // Read-only binding to the contract
	NotionalTokenTransactor // Write-only binding to the contract
	NotionalTokenFilterer   // Log filterer for contract events
}

// NotionalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotionalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotionalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotionalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotionalTokenSession struct {
	Contract     *NotionalToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotionalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotionalTokenCallerSession struct {
	Contract *NotionalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NotionalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotionalTokenTransactorSession struct {
	Contract     *NotionalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NotionalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotionalTokenRaw struct {
	Contract *NotionalToken // Generic contract binding to access the raw methods on
}

// NotionalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotionalTokenCallerRaw struct {
	Contract *NotionalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NotionalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotionalTokenTransactorRaw struct {
	Contract *NotionalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotionalToken creates a new instance of NotionalToken, bound to a specific deployed contract.
func NewNotionalToken(address common.Address, backend bind.ContractBackend) (*NotionalToken, error) {
	contract, err := bindNotionalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NotionalToken{NotionalTokenCaller: NotionalTokenCaller{contract: contract}, NotionalTokenTransactor: NotionalTokenTransactor{contract: contract}, NotionalTokenFilterer: NotionalTokenFilterer{contract: contract}}, nil
}

// NewNotionalTokenCaller creates a new read-only instance of NotionalToken, bound to a specific deployed contract.
func NewNotionalTokenCaller(address common.Address, caller bind.ContractCaller) (*NotionalTokenCaller, error) {
	contract, err := bindNotionalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotionalTokenCaller{contract: contract}, nil
}

// NewNotionalTokenTransactor creates a new write-only instance of NotionalToken, bound to a specific deployed contract.
func NewNotionalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NotionalTokenTransactor, error) {
	contract, err := bindNotionalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotionalTokenTransactor{contract: contract}, nil
}

// NewNotionalTokenFilterer creates a new log filterer instance of NotionalToken, bound to a specific deployed contract.
func NewNotionalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NotionalTokenFilterer, error) {
	contract, err := bindNotionalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotionalTokenFilterer{contract: contract}, nil
}

// bindNotionalToken binds a generic wrapper to an already deployed contract.
func bindNotionalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NotionalTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotionalToken *NotionalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotionalToken.Contract.NotionalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotionalToken *NotionalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotionalToken.Contract.NotionalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotionalToken *NotionalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotionalToken.Contract.NotionalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotionalToken *NotionalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotionalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotionalToken *NotionalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotionalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotionalToken *NotionalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotionalToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_NotionalToken *NotionalTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_NotionalToken *NotionalTokenSession) BalanceOfCalled() (common.Address, error) {
	return _NotionalToken.Contract.BalanceOfCalled(&_NotionalToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_NotionalToken *NotionalTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _NotionalToken.Contract.BalanceOfCalled(&_NotionalToken.CallOpts)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_NotionalToken *NotionalTokenCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "depositCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_NotionalToken *NotionalTokenSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _NotionalToken.Contract.DepositCalled(&_NotionalToken.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_NotionalToken *NotionalTokenCallerSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _NotionalToken.Contract.DepositCalled(&_NotionalToken.CallOpts, arg0)
}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_NotionalToken *NotionalTokenCaller) GetMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "getMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_NotionalToken *NotionalTokenSession) GetMaturity() (*big.Int, error) {
	return _NotionalToken.Contract.GetMaturity(&_NotionalToken.CallOpts)
}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_NotionalToken *NotionalTokenCallerSession) GetMaturity() (*big.Int, error) {
	return _NotionalToken.Contract.GetMaturity(&_NotionalToken.CallOpts)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_NotionalToken *NotionalTokenCaller) GetUnderlyingToken(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "getUnderlyingToken")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_NotionalToken *NotionalTokenSession) GetUnderlyingToken() (common.Address, *big.Int, error) {
	return _NotionalToken.Contract.GetUnderlyingToken(&_NotionalToken.CallOpts)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_NotionalToken *NotionalTokenCallerSession) GetUnderlyingToken() (common.Address, *big.Int, error) {
	return _NotionalToken.Contract.GetUnderlyingToken(&_NotionalToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_NotionalToken *NotionalTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "transferFromCalled", arg0)

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
func (_NotionalToken *NotionalTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _NotionalToken.Contract.TransferFromCalled(&_NotionalToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_NotionalToken *NotionalTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _NotionalToken.Contract.TransferFromCalled(&_NotionalToken.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_NotionalToken *NotionalTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_NotionalToken *NotionalTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.BalanceOf(&_NotionalToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_NotionalToken *NotionalTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.BalanceOf(&_NotionalToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_NotionalToken *NotionalTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_NotionalToken *NotionalTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.BalanceOfReturns(&_NotionalToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_NotionalToken *NotionalTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.BalanceOfReturns(&_NotionalToken.TransactOpts, b)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_NotionalToken *NotionalTokenTransactor) Deposit(opts *bind.TransactOpts, a *big.Int, r common.Address) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "deposit", a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_NotionalToken *NotionalTokenSession) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.Deposit(&_NotionalToken.TransactOpts, a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_NotionalToken *NotionalTokenTransactorSession) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.Deposit(&_NotionalToken.TransactOpts, a, r)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_NotionalToken *NotionalTokenTransactor) DepositReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "depositReturns", a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_NotionalToken *NotionalTokenSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.DepositReturns(&_NotionalToken.TransactOpts, a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_NotionalToken *NotionalTokenTransactorSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.DepositReturns(&_NotionalToken.TransactOpts, a)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_NotionalToken *NotionalTokenTransactor) GetMaturityReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "getMaturityReturns", m)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_NotionalToken *NotionalTokenSession) GetMaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.GetMaturityReturns(&_NotionalToken.TransactOpts, m)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_NotionalToken *NotionalTokenTransactorSession) GetMaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.GetMaturityReturns(&_NotionalToken.TransactOpts, m)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_NotionalToken *NotionalTokenTransactor) GetUnderlyingTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "getUnderlyingTokenReturns", a)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_NotionalToken *NotionalTokenSession) GetUnderlyingTokenReturns(a common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.GetUnderlyingTokenReturns(&_NotionalToken.TransactOpts, a)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_NotionalToken *NotionalTokenTransactorSession) GetUnderlyingTokenReturns(a common.Address) (*types.Transaction, error) {
	return _NotionalToken.Contract.GetUnderlyingTokenReturns(&_NotionalToken.TransactOpts, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_NotionalToken *NotionalTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_NotionalToken *NotionalTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.TransferFrom(&_NotionalToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_NotionalToken *NotionalTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.TransferFrom(&_NotionalToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_NotionalToken *NotionalTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_NotionalToken *NotionalTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _NotionalToken.Contract.TransferFromReturns(&_NotionalToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_NotionalToken *NotionalTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _NotionalToken.Contract.TransferFromReturns(&_NotionalToken.TransactOpts, b)
}
