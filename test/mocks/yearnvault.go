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

// YearnVaultMetaData contains all meta data concerning the YearnVault contract.
var YearnVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"p\",\"type\":\"uint256\"}],\"name\":\"pricePerShareReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdrawReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104f0806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b6b55f2511610066578063b6b55f2514610147578063b79c449b14610177578063d35147e4146101a7578063de7218a6146101d7578063e7ba6774146101f35761009e565b806320a8b9cd146100a357806326cca786146100bf5780632e1a7d4d146100db5780636f307dc31461010b57806399530b0614610129575b600080fd5b6100bd60048036038101906100b891906103ae565b61020f565b005b6100d960048036038101906100d491906103ae565b610219565b005b6100f560048036038101906100f091906103ae565b610223565b60405161010291906103ea565b60405180910390f35b610113610273565b6040516101209190610446565b60405180910390f35b61013161029c565b60405161013e91906103ea565b60405180910390f35b610161600480360381019061015c91906103ae565b6102a6565b60405161016e91906103ea565b60405180910390f35b610191600480360381019061018c919061048d565b6102f6565b60405161019e91906103ea565b60405180910390f35b6101c160048036038101906101bc919061048d565b61030e565b6040516101ce91906103ea565b60405180910390f35b6101f160048036038101906101ec91906103ae565b610326565b005b61020d6004803603810190610208919061048d565b610330565b005b8060048190555050565b8060028190555050565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506004549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002549050919050565b60056020528060005260406000206000915090505481565b60036020528060005260406000206000915090505481565b8060018190555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b6000819050919050565b61038b81610378565b811461039657600080fd5b50565b6000813590506103a881610382565b92915050565b6000602082840312156103c4576103c3610373565b5b60006103d284828501610399565b91505092915050565b6103e481610378565b82525050565b60006020820190506103ff60008301846103db565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061043082610405565b9050919050565b61044081610425565b82525050565b600060208201905061045b6000830184610437565b92915050565b61046a81610425565b811461047557600080fd5b50565b60008135905061048781610461565b92915050565b6000602082840312156104a3576104a2610373565b5b60006104b184828501610478565b9150509291505056fea2646970667358221220dd5a61eb4cff5d16243308fa9528e091a939f7fe46cb9f294ee877b00d8e3aec64736f6c634300080d0033",
}

// YearnVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnVaultMetaData.ABI instead.
var YearnVaultABI = YearnVaultMetaData.ABI

// YearnVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YearnVaultMetaData.Bin instead.
var YearnVaultBin = YearnVaultMetaData.Bin

// DeployYearnVault deploys a new Ethereum contract, binding an instance of YearnVault to it.
func DeployYearnVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YearnVault, error) {
	parsed, err := YearnVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YearnVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YearnVault{YearnVaultCaller: YearnVaultCaller{contract: contract}, YearnVaultTransactor: YearnVaultTransactor{contract: contract}, YearnVaultFilterer: YearnVaultFilterer{contract: contract}}, nil
}

// YearnVault is an auto generated Go binding around an Ethereum contract.
type YearnVault struct {
	YearnVaultCaller     // Read-only binding to the contract
	YearnVaultTransactor // Write-only binding to the contract
	YearnVaultFilterer   // Log filterer for contract events
}

// YearnVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnVaultSession struct {
	Contract     *YearnVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnVaultCallerSession struct {
	Contract *YearnVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YearnVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnVaultTransactorSession struct {
	Contract     *YearnVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YearnVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnVaultRaw struct {
	Contract *YearnVault // Generic contract binding to access the raw methods on
}

// YearnVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnVaultCallerRaw struct {
	Contract *YearnVaultCaller // Generic read-only contract binding to access the raw methods on
}

// YearnVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnVaultTransactorRaw struct {
	Contract *YearnVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnVault creates a new instance of YearnVault, bound to a specific deployed contract.
func NewYearnVault(address common.Address, backend bind.ContractBackend) (*YearnVault, error) {
	contract, err := bindYearnVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnVault{YearnVaultCaller: YearnVaultCaller{contract: contract}, YearnVaultTransactor: YearnVaultTransactor{contract: contract}, YearnVaultFilterer: YearnVaultFilterer{contract: contract}}, nil
}

// NewYearnVaultCaller creates a new read-only instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultCaller(address common.Address, caller bind.ContractCaller) (*YearnVaultCaller, error) {
	contract, err := bindYearnVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnVaultCaller{contract: contract}, nil
}

// NewYearnVaultTransactor creates a new write-only instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnVaultTransactor, error) {
	contract, err := bindYearnVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnVaultTransactor{contract: contract}, nil
}

// NewYearnVaultFilterer creates a new log filterer instance of YearnVault, bound to a specific deployed contract.
func NewYearnVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnVaultFilterer, error) {
	contract, err := bindYearnVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnVaultFilterer{contract: contract}, nil
}

// bindYearnVault binds a generic wrapper to an already deployed contract.
func bindYearnVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YearnVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnVault *YearnVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnVault.Contract.YearnVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnVault *YearnVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.Contract.YearnVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnVault *YearnVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnVault.Contract.YearnVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnVault *YearnVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnVault *YearnVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnVault *YearnVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnVault.Contract.contract.Transact(opts, method, params...)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "depositCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.DepositCalled(&_YearnVault.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.DepositCalled(&_YearnVault.CallOpts, arg0)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultSession) PricePerShare() (*big.Int, error) {
	return _YearnVault.Contract.PricePerShare(&_YearnVault.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) PricePerShare() (*big.Int, error) {
	return _YearnVault.Contract.PricePerShare(&_YearnVault.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YearnVault *YearnVaultCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YearnVault *YearnVaultSession) Underlying() (common.Address, error) {
	return _YearnVault.Contract.Underlying(&_YearnVault.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_YearnVault *YearnVaultCallerSession) Underlying() (common.Address, error) {
	return _YearnVault.Contract.Underlying(&_YearnVault.CallOpts)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultCaller) WithdrawCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnVault.contract.Call(opts, &out, "withdrawCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultSession) WithdrawCalled(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.WithdrawCalled(&_YearnVault.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_YearnVault *YearnVaultCallerSession) WithdrawCalled(arg0 common.Address) (*big.Int, error) {
	return _YearnVault.Contract.WithdrawCalled(&_YearnVault.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Deposit(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "deposit", a)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultSession) Deposit(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit(&_YearnVault.TransactOpts, a)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Deposit(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Deposit(&_YearnVault.TransactOpts, a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_YearnVault *YearnVaultTransactor) DepositReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "depositReturns", a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_YearnVault *YearnVaultSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.DepositReturns(&_YearnVault.TransactOpts, a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_YearnVault *YearnVaultTransactorSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.DepositReturns(&_YearnVault.TransactOpts, a)
}

// PricePerShareReturns is a paid mutator transaction binding the contract method 0xde7218a6.
//
// Solidity: function pricePerShareReturns(uint256 p) returns()
func (_YearnVault *YearnVaultTransactor) PricePerShareReturns(opts *bind.TransactOpts, p *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "pricePerShareReturns", p)
}

// PricePerShareReturns is a paid mutator transaction binding the contract method 0xde7218a6.
//
// Solidity: function pricePerShareReturns(uint256 p) returns()
func (_YearnVault *YearnVaultSession) PricePerShareReturns(p *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.PricePerShareReturns(&_YearnVault.TransactOpts, p)
}

// PricePerShareReturns is a paid mutator transaction binding the contract method 0xde7218a6.
//
// Solidity: function pricePerShareReturns(uint256 p) returns()
func (_YearnVault *YearnVaultTransactorSession) PricePerShareReturns(p *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.PricePerShareReturns(&_YearnVault.TransactOpts, p)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YearnVault *YearnVaultTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YearnVault *YearnVaultSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.UnderlyingReturns(&_YearnVault.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_YearnVault *YearnVaultTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _YearnVault.Contract.UnderlyingReturns(&_YearnVault.TransactOpts, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultTransactor) Withdraw(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdraw", a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultSession) Withdraw(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw(&_YearnVault.TransactOpts, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 a) returns(uint256)
func (_YearnVault *YearnVaultTransactorSession) Withdraw(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.Withdraw(&_YearnVault.TransactOpts, a)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_YearnVault *YearnVaultTransactor) WithdrawReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _YearnVault.contract.Transact(opts, "withdrawReturns", a)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_YearnVault *YearnVaultSession) WithdrawReturns(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.WithdrawReturns(&_YearnVault.TransactOpts, a)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_YearnVault *YearnVaultTransactorSession) WithdrawReturns(a *big.Int) (*types.Transaction, error) {
	return _YearnVault.Contract.WithdrawReturns(&_YearnVault.TransactOpts, a)
}
