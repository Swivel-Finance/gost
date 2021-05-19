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

// Erc20ABI is the input ABI used to generate the binding from.
const Erc20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approveCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"approveReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Erc20Bin is the compiled bytecode used for deploying new contracts.
var Erc20Bin = "0x608060405234801561001057600080fd5b50610728806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80636581d543116100665780636581d543146101305780639dd0ff3714610160578063a9059cbb1461017c578063c1d2e9a1146101ac578063e541efa2146101dc57610093565b8063095ea7b31461009857806323b872dd146100c857806342b6cdbc146100f85780636521b96a14610114575b600080fd5b6100b260048036038101906100ad9190610574565b61020d565b6040516100bf919061062f565b60405180910390f35b6100e260048036038101906100dd9190610525565b61026a565b6040516100ef919061062f565b60405180910390f35b610112600480360381019061010d91906105b0565b610364565b005b61012e600480360381019061012991906105b0565b610381565b005b61014a600480360381019061014591906104fc565b61039e565b604051610157919061064a565b60405180910390f35b61017a600480360381019061017591906105b0565b6103b6565b005b61019660048036038101906101919190610574565b6103d3565b6040516101a3919061062f565b60405180910390f35b6101c660048036038101906101c191906104fc565b610431565b6040516101d3919061064a565b60405180910390f35b6101f660048036038101906101f191906104fc565b610449565b604051610204929190610606565b60405180910390f35b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160009054906101000a900460ff16905092915050565b600061027461048d565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600560009054906101000a900460ff169150509392505050565b80600360006101000a81548160ff02191690831515021790555050565b80600560006101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915090505481565b80600160006101000a81548160ff02191690831515021790555050565b600081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600360009054906101000a900460ff16905092915050565b60026020528060005260406000206000915090505481565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6000813590506104cc816106ad565b92915050565b6000813590506104e1816106c4565b92915050565b6000813590506104f6816106db565b92915050565b60006020828403121561050e57600080fd5b600061051c848285016104bd565b91505092915050565b60008060006060848603121561053a57600080fd5b6000610548868287016104bd565b9350506020610559868287016104bd565b925050604061056a868287016104e7565b9150509250925092565b6000806040838503121561058757600080fd5b6000610595858286016104bd565b92505060206105a6858286016104e7565b9150509250929050565b6000602082840312156105c257600080fd5b60006105d0848285016104d2565b91505092915050565b6105e281610665565b82525050565b6105f181610677565b82525050565b610600816106a3565b82525050565b600060408201905061061b60008301856105d9565b61062860208301846105f7565b9392505050565b600060208201905061064460008301846105e8565b92915050565b600060208201905061065f60008301846105f7565b92915050565b600061067082610683565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6106b681610665565b81146106c157600080fd5b50565b6106cd81610677565b81146106d857600080fd5b50565b6106e4816106a3565b81146106ef57600080fd5b5056fea264697066735822122030d6aee281c8408e865e225efcf6f63f84091f433b1d967b2d8278d408330ee464736f6c63430008000033"

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Erc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) ApproveCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "approveCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Session) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.ApproveCalled(&_Erc20.CallOpts, arg0)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.ApproveCalled(&_Erc20.CallOpts, arg0)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) TransferCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Session) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.TransferCalled(&_Erc20.CallOpts, arg0)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.TransferCalled(&_Erc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20Caller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20Session) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Erc20.Contract.TransferFromCalled(&_Erc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20CallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Erc20.Contract.TransferFromCalled(&_Erc20.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, s, a)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) ApproveReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approveReturns", b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20Session) ApproveReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.ApproveReturns(&_Erc20.TransactOpts, b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) ApproveReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.ApproveReturns(&_Erc20.TransactOpts, b)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20Session) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFromReturns(&_Erc20.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFromReturns(&_Erc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) TransferReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferReturns", b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20Session) TransferReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferReturns(&_Erc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) TransferReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferReturns(&_Erc20.TransactOpts, b)
}
