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

// TempusTokenABI is the input ABI used to generate the binding from.
const TempusTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TempusTokenBin is the compiled bytecode used for deploying new contracts.
var TempusTokenBin = "0x608060405234801561001057600080fd5b50610612806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806323b872dd1461006757806327e235e3146100975780636521b96a146100c757806370a08231146100e3578063c0c6e5ab14610113578063e541efa21461012f575b600080fd5b610081600480360381019061007c919061042b565b610160565b60405161008e9190610499565b60405180910390f35b6100b160048036038101906100ac91906104b4565b61025a565b6040516100be91906104f0565b60405180910390f35b6100e160048036038101906100dc9190610537565b610272565b005b6100fd60048036038101906100f891906104b4565b61028f565b60405161010a91906104f0565b60405180910390f35b61012d60048036038101906101289190610564565b6102d7565b005b610149600480360381019061014491906104b4565b61031e565b6040516101579291906105b3565b60405180910390f35b600061016a610362565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600160009054906101000a900460ff169150509392505050565b60006020528060005260406000206000915090505481565b80600160006101000a81548160ff02191690831515021790555050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103c282610397565b9050919050565b6103d2816103b7565b81146103dd57600080fd5b50565b6000813590506103ef816103c9565b92915050565b6000819050919050565b610408816103f5565b811461041357600080fd5b50565b600081359050610425816103ff565b92915050565b60008060006060848603121561044457610443610392565b5b6000610452868287016103e0565b9350506020610463868287016103e0565b925050604061047486828701610416565b9150509250925092565b60008115159050919050565b6104938161047e565b82525050565b60006020820190506104ae600083018461048a565b92915050565b6000602082840312156104ca576104c9610392565b5b60006104d8848285016103e0565b91505092915050565b6104ea816103f5565b82525050565b600060208201905061050560008301846104e1565b92915050565b6105148161047e565b811461051f57600080fd5b50565b6000813590506105318161050b565b92915050565b60006020828403121561054d5761054c610392565b5b600061055b84828501610522565b91505092915050565b6000806040838503121561057b5761057a610392565b5b6000610589858286016103e0565b925050602061059a85828601610416565b9150509250929050565b6105ad816103b7565b82525050565b60006040820190506105c860008301856105a4565b6105d560208301846104e1565b939250505056fea26469706673582212204942901a82bf14d9c7fa5b29706e3066d3c2948c2887d7a51014e1630a75e16164736f6c634300080d0033"

// DeployTempusToken deploys a new Ethereum contract, binding an instance of TempusToken to it.
func DeployTempusToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TempusToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TempusTokenBin), backend)
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_TempusToken *TempusTokenCaller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TempusToken.contract.Call(opts, &out, "balanceOf", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_TempusToken *TempusTokenSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _TempusToken.Contract.BalanceOf(&_TempusToken.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_TempusToken *TempusTokenCallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _TempusToken.Contract.BalanceOf(&_TempusToken.CallOpts, a)
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

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

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

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_TempusToken *TempusTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, a common.Address, b *big.Int) (*types.Transaction, error) {
	return _TempusToken.contract.Transact(opts, "balanceOfReturns", a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_TempusToken *TempusTokenSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOfReturns(&_TempusToken.TransactOpts, a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_TempusToken *TempusTokenTransactorSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _TempusToken.Contract.BalanceOfReturns(&_TempusToken.TransactOpts, a, b)
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
