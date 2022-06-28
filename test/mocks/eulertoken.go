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

// EulerTokenMetaData contains all meta data concerning the EulerToken contract.
var EulerTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610258806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063022ddcdb14610051578063441a3e70146100815780635024c2791461009d578063e2bbb158146100cd575b600080fd5b61006b6004803603810190610066919061018b565b6100e9565b60405161007891906101c7565b60405180910390f35b61009b600480360381019061009691906101e2565b610101565b005b6100b760048036038101906100b2919061018b565b61011d565b6040516100c491906101c7565b60405180910390f35b6100e760048036038101906100e291906101e2565b610135565b005b60006020528060005260406000206000915090505481565b8060016000848152602001908152602001600020819055505050565b60016020528060005260406000206000915090505481565b80600080848152602001908152602001600020819055505050565b600080fd5b6000819050919050565b61016881610155565b811461017357600080fd5b50565b6000813590506101858161015f565b92915050565b6000602082840312156101a1576101a0610150565b5b60006101af84828501610176565b91505092915050565b6101c181610155565b82525050565b60006020820190506101dc60008301846101b8565b92915050565b600080604083850312156101f9576101f8610150565b5b600061020785828601610176565b925050602061021885828601610176565b915050925092905056fea2646970667358221220136200ac0b40878cf1f3afc658a9c5f45b361026b30623c525a13d2e99ca061d64736f6c634300080d0033",
}

// EulerTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerTokenMetaData.ABI instead.
var EulerTokenABI = EulerTokenMetaData.ABI

// EulerTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EulerTokenMetaData.Bin instead.
var EulerTokenBin = EulerTokenMetaData.Bin

// DeployEulerToken deploys a new Ethereum contract, binding an instance of EulerToken to it.
func DeployEulerToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EulerToken, error) {
	parsed, err := EulerTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EulerTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EulerToken{EulerTokenCaller: EulerTokenCaller{contract: contract}, EulerTokenTransactor: EulerTokenTransactor{contract: contract}, EulerTokenFilterer: EulerTokenFilterer{contract: contract}}, nil
}

// EulerToken is an auto generated Go binding around an Ethereum contract.
type EulerToken struct {
	EulerTokenCaller     // Read-only binding to the contract
	EulerTokenTransactor // Write-only binding to the contract
	EulerTokenFilterer   // Log filterer for contract events
}

// EulerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerTokenSession struct {
	Contract     *EulerToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerTokenCallerSession struct {
	Contract *EulerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EulerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerTokenTransactorSession struct {
	Contract     *EulerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EulerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerTokenRaw struct {
	Contract *EulerToken // Generic contract binding to access the raw methods on
}

// EulerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerTokenCallerRaw struct {
	Contract *EulerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EulerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerTokenTransactorRaw struct {
	Contract *EulerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerToken creates a new instance of EulerToken, bound to a specific deployed contract.
func NewEulerToken(address common.Address, backend bind.ContractBackend) (*EulerToken, error) {
	contract, err := bindEulerToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerToken{EulerTokenCaller: EulerTokenCaller{contract: contract}, EulerTokenTransactor: EulerTokenTransactor{contract: contract}, EulerTokenFilterer: EulerTokenFilterer{contract: contract}}, nil
}

// NewEulerTokenCaller creates a new read-only instance of EulerToken, bound to a specific deployed contract.
func NewEulerTokenCaller(address common.Address, caller bind.ContractCaller) (*EulerTokenCaller, error) {
	contract, err := bindEulerToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerTokenCaller{contract: contract}, nil
}

// NewEulerTokenTransactor creates a new write-only instance of EulerToken, bound to a specific deployed contract.
func NewEulerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerTokenTransactor, error) {
	contract, err := bindEulerToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerTokenTransactor{contract: contract}, nil
}

// NewEulerTokenFilterer creates a new log filterer instance of EulerToken, bound to a specific deployed contract.
func NewEulerTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerTokenFilterer, error) {
	contract, err := bindEulerToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerTokenFilterer{contract: contract}, nil
}

// bindEulerToken binds a generic wrapper to an already deployed contract.
func bindEulerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EulerTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerToken *EulerTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerToken.Contract.EulerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerToken *EulerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerToken.Contract.EulerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerToken *EulerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerToken.Contract.EulerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerToken *EulerTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerToken *EulerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerToken *EulerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerToken.Contract.contract.Transact(opts, method, params...)
}

// DepositCalled is a free data retrieval call binding the contract method 0x022ddcdb.
//
// Solidity: function depositCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCaller) DepositCalled(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerToken.contract.Call(opts, &out, "depositCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCalled is a free data retrieval call binding the contract method 0x022ddcdb.
//
// Solidity: function depositCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenSession) DepositCalled(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.DepositCalled(&_EulerToken.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0x022ddcdb.
//
// Solidity: function depositCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCallerSession) DepositCalled(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.DepositCalled(&_EulerToken.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0x5024c279.
//
// Solidity: function withdrawCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCaller) WithdrawCalled(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerToken.contract.Call(opts, &out, "withdrawCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCalled is a free data retrieval call binding the contract method 0x5024c279.
//
// Solidity: function withdrawCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenSession) WithdrawCalled(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.WithdrawCalled(&_EulerToken.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0x5024c279.
//
// Solidity: function withdrawCalled(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCallerSession) WithdrawCalled(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.WithdrawCalled(&_EulerToken.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenTransactor) Deposit(opts *bind.TransactOpts, i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.contract.Transact(opts, "deposit", i, a)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenSession) Deposit(i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.Deposit(&_EulerToken.TransactOpts, i, a)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenTransactorSession) Deposit(i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.Deposit(&_EulerToken.TransactOpts, i, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenTransactor) Withdraw(opts *bind.TransactOpts, i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.contract.Transact(opts, "withdraw", i, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenSession) Withdraw(i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.Withdraw(&_EulerToken.TransactOpts, i, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 i, uint256 a) returns()
func (_EulerToken *EulerTokenTransactorSession) Withdraw(i *big.Int, a *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.Withdraw(&_EulerToken.TransactOpts, i, a)
}
