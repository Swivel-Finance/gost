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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdrawReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610396806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806320a8b9cd1461006757806326cca786146100835780632e1a7d4d1461009f578063b6b55f25146100cf578063b79c449b146100ff578063d35147e41461012f575b600080fd5b610081600480360381019061007c919061027e565b61015f565b005b61009d6004803603810190610098919061027e565b610169565b005b6100b960048036038101906100b4919061027e565b610173565b6040516100c691906102ba565b60405180910390f35b6100e960048036038101906100e4919061027e565b6101c3565b6040516100f691906102ba565b60405180910390f35b61011960048036038101906101149190610333565b610213565b60405161012691906102ba565b60405180910390f35b61014960048036038101906101449190610333565b61022b565b60405161015691906102ba565b60405180910390f35b8060028190555050565b8060008190555050565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002549050919050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000549050919050565b60036020528060005260406000206000915090505481565b60016020528060005260406000206000915090505481565b600080fd5b6000819050919050565b61025b81610248565b811461026657600080fd5b50565b60008135905061027881610252565b92915050565b60006020828403121561029457610293610243565b5b60006102a284828501610269565b91505092915050565b6102b481610248565b82525050565b60006020820190506102cf60008301846102ab565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610300826102d5565b9050919050565b610310816102f5565b811461031b57600080fd5b50565b60008135905061032d81610307565b92915050565b60006020828403121561034957610348610243565b5b60006103578482850161031e565b9150509291505056fea26469706673582212206bbdc3d22a15de452fd6120cb7b90382360da9b506ea5e49c9c9c2638f69008b64736f6c634300080d0033",
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
