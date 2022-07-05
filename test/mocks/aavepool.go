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

// AavePoolMetaData contains all meta data concerning the AavePool contract.
var AavePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"c\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"code\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getReserveNormalizedIncomeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"code\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdrawReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610873806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b79c449b1161005b578063b79c449b146100ea578063d15e00531461011d578063d35147e41461014d578063e8eda9df146101805761007d565b80630d8f89391461008257806320a8b9cd1461009e57806369328dec146100ba575b600080fd5b61009c60048036038101906100979190610604565b61019c565b005b6100b860048036038101906100b39190610604565b6101a6565b005b6100d460048036038101906100cf919061068f565b6101b0565b6040516100e191906106f1565b60405180910390f35b61010460048036038101906100ff919061070c565b610307565b6040516101149493929190610765565b60405180910390f35b6101376004803603810190610132919061070c565b610385565b60405161014491906106f1565b60405180910390f35b6101676004803603810190610162919061070c565b610390565b6040516101779493929190610765565b60405180910390f35b61019a600480360381019061019591906107d6565b61040e565b005b8060008190555050565b8060028190555050565b60006101ba610571565b8381600001818152505082816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548161ffff021916908361ffff1602179055509050506002549150509392505050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900461ffff16905084565b600080549050919050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900461ffff16905084565b610416610571565b8381600001818152505082816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505081816060019061ffff16908161ffff168152505080600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548161ffff021916908361ffff1602179055509050505050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600061ffff1681525090565b600080fd5b6000819050919050565b6105e1816105ce565b81146105ec57600080fd5b50565b6000813590506105fe816105d8565b92915050565b60006020828403121561061a576106196105c9565b5b6000610628848285016105ef565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061065c82610631565b9050919050565b61066c81610651565b811461067757600080fd5b50565b60008135905061068981610663565b92915050565b6000806000606084860312156106a8576106a76105c9565b5b60006106b68682870161067a565b93505060206106c7868287016105ef565b92505060406106d88682870161067a565b9150509250925092565b6106eb816105ce565b82525050565b600060208201905061070660008301846106e2565b92915050565b600060208284031215610722576107216105c9565b5b60006107308482850161067a565b91505092915050565b61074281610651565b82525050565b600061ffff82169050919050565b61075f81610748565b82525050565b600060808201905061077a60008301876106e2565b6107876020830186610739565b6107946040830185610739565b6107a16060830184610756565b95945050505050565b6107b381610748565b81146107be57600080fd5b50565b6000813590506107d0816107aa565b92915050565b600080600080608085870312156107f0576107ef6105c9565b5b60006107fe8782880161067a565b945050602061080f878288016105ef565b93505060406108208782880161067a565b9250506060610831878288016107c1565b9150509295919450925056fea264697066735822122019385d33395d3dc5114348507ae759170f4524ccdc1e32a5346776f408d6bd6564736f6c634300080d0033",
}

// AavePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AavePoolMetaData.ABI instead.
var AavePoolABI = AavePoolMetaData.ABI

// AavePoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AavePoolMetaData.Bin instead.
var AavePoolBin = AavePoolMetaData.Bin

// DeployAavePool deploys a new Ethereum contract, binding an instance of AavePool to it.
func DeployAavePool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AavePool, error) {
	parsed, err := AavePoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AavePoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AavePool{AavePoolCaller: AavePoolCaller{contract: contract}, AavePoolTransactor: AavePoolTransactor{contract: contract}, AavePoolFilterer: AavePoolFilterer{contract: contract}}, nil
}

// AavePool is an auto generated Go binding around an Ethereum contract.
type AavePool struct {
	AavePoolCaller     // Read-only binding to the contract
	AavePoolTransactor // Write-only binding to the contract
	AavePoolFilterer   // Log filterer for contract events
}

// AavePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AavePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AavePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AavePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AavePoolSession struct {
	Contract     *AavePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AavePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AavePoolCallerSession struct {
	Contract *AavePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AavePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AavePoolTransactorSession struct {
	Contract     *AavePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AavePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AavePoolRaw struct {
	Contract *AavePool // Generic contract binding to access the raw methods on
}

// AavePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AavePoolCallerRaw struct {
	Contract *AavePoolCaller // Generic read-only contract binding to access the raw methods on
}

// AavePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AavePoolTransactorRaw struct {
	Contract *AavePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavePool creates a new instance of AavePool, bound to a specific deployed contract.
func NewAavePool(address common.Address, backend bind.ContractBackend) (*AavePool, error) {
	contract, err := bindAavePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AavePool{AavePoolCaller: AavePoolCaller{contract: contract}, AavePoolTransactor: AavePoolTransactor{contract: contract}, AavePoolFilterer: AavePoolFilterer{contract: contract}}, nil
}

// NewAavePoolCaller creates a new read-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolCaller(address common.Address, caller bind.ContractCaller) (*AavePoolCaller, error) {
	contract, err := bindAavePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolCaller{contract: contract}, nil
}

// NewAavePoolTransactor creates a new write-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AavePoolTransactor, error) {
	contract, err := bindAavePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolTransactor{contract: contract}, nil
}

// NewAavePoolFilterer creates a new log filterer instance of AavePool, bound to a specific deployed contract.
func NewAavePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AavePoolFilterer, error) {
	contract, err := bindAavePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AavePoolFilterer{contract: contract}, nil
}

// bindAavePool binds a generic wrapper to an already deployed contract.
func bindAavePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AavePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.AavePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transact(opts, method, params...)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "depositCalled", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		OnBehalfOf common.Address
		To         common.Address
		Code       uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OnBehalfOf = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Code = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolSession) DepositCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	return _AavePool.Contract.DepositCalled(&_AavePool.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolCallerSession) DepositCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	return _AavePool.Contract.DepositCalled(&_AavePool.CallOpts, arg0)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address ) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveNormalizedIncome", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address ) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveNormalizedIncome(arg0 common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, arg0)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address ) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveNormalizedIncome(arg0 common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolCaller) WithdrawCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "withdrawCalled", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		OnBehalfOf common.Address
		To         common.Address
		Code       uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OnBehalfOf = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Code = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolSession) WithdrawCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	return _AavePool.Contract.WithdrawCalled(&_AavePool.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 amount, address onBehalfOf, address to, uint16 code)
func (_AavePool *AavePoolCallerSession) WithdrawCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	OnBehalfOf common.Address
	To         common.Address
	Code       uint16
}, error) {
	return _AavePool.Contract.WithdrawCalled(&_AavePool.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address u, uint256 a, address o, uint16 c) returns()
func (_AavePool *AavePoolTransactor) Deposit(opts *bind.TransactOpts, u common.Address, a *big.Int, o common.Address, c uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "deposit", u, a, o, c)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address u, uint256 a, address o, uint16 c) returns()
func (_AavePool *AavePoolSession) Deposit(u common.Address, a *big.Int, o common.Address, c uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, u, a, o, c)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address u, uint256 a, address o, uint16 c) returns()
func (_AavePool *AavePoolTransactorSession) Deposit(u common.Address, a *big.Int, o common.Address, c uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, u, a, o, c)
}

// GetReserveNormalizedIncomeReturns is a paid mutator transaction binding the contract method 0x0d8f8939.
//
// Solidity: function getReserveNormalizedIncomeReturns(uint256 n) returns()
func (_AavePool *AavePoolTransactor) GetReserveNormalizedIncomeReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "getReserveNormalizedIncomeReturns", n)
}

// GetReserveNormalizedIncomeReturns is a paid mutator transaction binding the contract method 0x0d8f8939.
//
// Solidity: function getReserveNormalizedIncomeReturns(uint256 n) returns()
func (_AavePool *AavePoolSession) GetReserveNormalizedIncomeReturns(n *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.GetReserveNormalizedIncomeReturns(&_AavePool.TransactOpts, n)
}

// GetReserveNormalizedIncomeReturns is a paid mutator transaction binding the contract method 0x0d8f8939.
//
// Solidity: function getReserveNormalizedIncomeReturns(uint256 n) returns()
func (_AavePool *AavePoolTransactorSession) GetReserveNormalizedIncomeReturns(n *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.GetReserveNormalizedIncomeReturns(&_AavePool.TransactOpts, n)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address u, uint256 a, address t) returns(uint256)
func (_AavePool *AavePoolTransactor) Withdraw(opts *bind.TransactOpts, u common.Address, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "withdraw", u, a, t)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address u, uint256 a, address t) returns(uint256)
func (_AavePool *AavePoolSession) Withdraw(u common.Address, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, u, a, t)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address u, uint256 a, address t) returns(uint256)
func (_AavePool *AavePoolTransactorSession) Withdraw(u common.Address, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, u, a, t)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_AavePool *AavePoolTransactor) WithdrawReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "withdrawReturns", a)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_AavePool *AavePoolSession) WithdrawReturns(a *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.WithdrawReturns(&_AavePool.TransactOpts, a)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 a) returns()
func (_AavePool *AavePoolTransactorSession) WithdrawReturns(a *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.WithdrawReturns(&_AavePool.TransactOpts, a)
}
