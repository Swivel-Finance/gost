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
const NotionalTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"getMaturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getUnderlyingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NotionalTokenBin is the compiled bytecode used for deploying new contracts.
var NotionalTokenBin = "0x608060405234801561001057600080fd5b50610810806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063d35147e411610071578063d35147e414610189578063dea1a7e2146101b9578063e16695b5146101d7578063e541efa2146101f5578063ee719bc814610226578063fe68267014610244576100b4565b806323b872dd146100b957806339100838146100e957806347e7ef24146101055780636521b96a1461012157806370a082311461013d578063a7d8cf031461016d575b600080fd5b6100d360048036038101906100ce91906105e1565b610260565b6040516100e0919061064f565b60405180910390f35b61010360048036038101906100fe919061066a565b61035a565b005b61011f600480360381019061011a9190610697565b610364565b005b61013b60048036038101906101369190610703565b6103ac565b005b61015760048036038101906101529190610730565b6103c9565b604051610164919061076c565b60405180910390f35b61018760048036038101906101829190610730565b610416565b005b6101a3600480360381019061019e9190610730565b610459565b6040516101b0919061076c565b60405180910390f35b6101c1610471565b6040516101ce9190610796565b60405180910390f35b6101df610497565b6040516101ec919061076c565b60405180910390f35b61020f600480360381019061020a9190610730565b6104a1565b60405161021d9291906107b1565b60405180910390f35b61022e6104e5565b60405161023b9190610796565b60405180910390f35b61025e6004803603810190610259919061066a565b61050e565b005b600061026a610518565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260009054906101000a900460ff169150509392505050565b8060038190555050565b80600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b80600260006101000a81548160ff02191690831515021790555050565b600081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003549050919050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60056020528060005260406000206000915090505481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060018190555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105788261054d565b9050919050565b6105888161056d565b811461059357600080fd5b50565b6000813590506105a58161057f565b92915050565b6000819050919050565b6105be816105ab565b81146105c957600080fd5b50565b6000813590506105db816105b5565b92915050565b6000806000606084860312156105fa576105f9610548565b5b600061060886828701610596565b935050602061061986828701610596565b925050604061062a868287016105cc565b9150509250925092565b60008115159050919050565b61064981610634565b82525050565b60006020820190506106646000830184610640565b92915050565b6000602082840312156106805761067f610548565b5b600061068e848285016105cc565b91505092915050565b600080604083850312156106ae576106ad610548565b5b60006106bc85828601610596565b92505060206106cd858286016105cc565b9150509250929050565b6106e081610634565b81146106eb57600080fd5b50565b6000813590506106fd816106d7565b92915050565b60006020828403121561071957610718610548565b5b6000610727848285016106ee565b91505092915050565b60006020828403121561074657610745610548565b5b600061075484828501610596565b91505092915050565b610766816105ab565b82525050565b6000602082019050610781600083018461075d565b92915050565b6107908161056d565b82525050565b60006020820190506107ab6000830184610787565b92915050565b60006040820190506107c66000830185610787565b6107d3602083018461075d565b939250505056fea26469706673582212205dbd05f6535c907c775fb2c5f294dc5becdaf6e41a461a89a12c7c44390216b964736f6c634300080d0033"

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
// Solidity: function getUnderlyingToken() view returns(address)
func (_NotionalToken *NotionalTokenCaller) GetUnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NotionalToken.contract.Call(opts, &out, "getUnderlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address)
func (_NotionalToken *NotionalTokenSession) GetUnderlyingToken() (common.Address, error) {
	return _NotionalToken.Contract.GetUnderlyingToken(&_NotionalToken.CallOpts)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address)
func (_NotionalToken *NotionalTokenCallerSession) GetUnderlyingToken() (common.Address, error) {
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

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address a, uint256 b) returns()
func (_NotionalToken *NotionalTokenTransactor) Deposit(opts *bind.TransactOpts, a common.Address, b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.contract.Transact(opts, "deposit", a, b)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address a, uint256 b) returns()
func (_NotionalToken *NotionalTokenSession) Deposit(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.Deposit(&_NotionalToken.TransactOpts, a, b)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address a, uint256 b) returns()
func (_NotionalToken *NotionalTokenTransactorSession) Deposit(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _NotionalToken.Contract.Deposit(&_NotionalToken.TransactOpts, a, b)
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
