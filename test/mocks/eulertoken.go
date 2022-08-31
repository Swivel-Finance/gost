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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"convertBalanceToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"convertBalanceToUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingAssetReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610452806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063441a3e701161005b578063441a3e70146101255780635024c279146101415780637158da7c14610171578063e2bbb1581461018f57610088565b8063010ad6d11461008d578063022ddcdb146100bd5780632a646f07146100ed5780634167cc2b14610109575b600080fd5b6100a760048036038101906100a291906102d0565b6101ab565b6040516100b4919061030c565b60405180910390f35b6100d760048036038101906100d291906102d0565b6101b7565b6040516100e4919061030c565b60405180910390f35b610107600480360381019061010291906102d0565b6101cf565b005b610123600480360381019061011e9190610385565b6101d9565b005b61013f600480360381019061013a91906103b2565b61021c565b005b61015b600480360381019061015691906102d0565b610238565b604051610168919061030c565b60405180910390f35b610179610250565b6040516101869190610401565b60405180910390f35b6101a960048036038101906101a491906103b2565b610279565b005b60006001549050919050565b60026020528060005260406000206000915090505481565b8060018190555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b8060036000848152602001908152602001600020819055505050565b60036020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060026000848152602001908152602001600020819055505050565b600080fd5b6000819050919050565b6102ad8161029a565b81146102b857600080fd5b50565b6000813590506102ca816102a4565b92915050565b6000602082840312156102e6576102e5610295565b5b60006102f4848285016102bb565b91505092915050565b6103068161029a565b82525050565b600060208201905061032160008301846102fd565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061035282610327565b9050919050565b61036281610347565b811461036d57600080fd5b50565b60008135905061037f81610359565b92915050565b60006020828403121561039b5761039a610295565b5b60006103a984828501610370565b91505092915050565b600080604083850312156103c9576103c8610295565b5b60006103d7858286016102bb565b92505060206103e8858286016102bb565b9150509250929050565b6103fb81610347565b82525050565b600060208201905061041660008301846103f2565b9291505056fea2646970667358221220a69aa00690cbb6ee58ef2d10648e4a4bbdf89174fb493db52f16c5221fa11d3964736f6c63430008100033",
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

// ConvertBalanceToUnderlying is a free data retrieval call binding the contract method 0x010ad6d1.
//
// Solidity: function convertBalanceToUnderlying(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCaller) ConvertBalanceToUnderlying(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerToken.contract.Call(opts, &out, "convertBalanceToUnderlying", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertBalanceToUnderlying is a free data retrieval call binding the contract method 0x010ad6d1.
//
// Solidity: function convertBalanceToUnderlying(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenSession) ConvertBalanceToUnderlying(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.ConvertBalanceToUnderlying(&_EulerToken.CallOpts, arg0)
}

// ConvertBalanceToUnderlying is a free data retrieval call binding the contract method 0x010ad6d1.
//
// Solidity: function convertBalanceToUnderlying(uint256 ) view returns(uint256)
func (_EulerToken *EulerTokenCallerSession) ConvertBalanceToUnderlying(arg0 *big.Int) (*big.Int, error) {
	return _EulerToken.Contract.ConvertBalanceToUnderlying(&_EulerToken.CallOpts, arg0)
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

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_EulerToken *EulerTokenCaller) UnderlyingAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerToken.contract.Call(opts, &out, "underlyingAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_EulerToken *EulerTokenSession) UnderlyingAsset() (common.Address, error) {
	return _EulerToken.Contract.UnderlyingAsset(&_EulerToken.CallOpts)
}

// UnderlyingAsset is a free data retrieval call binding the contract method 0x7158da7c.
//
// Solidity: function underlyingAsset() view returns(address)
func (_EulerToken *EulerTokenCallerSession) UnderlyingAsset() (common.Address, error) {
	return _EulerToken.Contract.UnderlyingAsset(&_EulerToken.CallOpts)
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

// ConvertBalanceToUnderlyingReturns is a paid mutator transaction binding the contract method 0x2a646f07.
//
// Solidity: function convertBalanceToUnderlyingReturns(uint256 n) returns()
func (_EulerToken *EulerTokenTransactor) ConvertBalanceToUnderlyingReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _EulerToken.contract.Transact(opts, "convertBalanceToUnderlyingReturns", n)
}

// ConvertBalanceToUnderlyingReturns is a paid mutator transaction binding the contract method 0x2a646f07.
//
// Solidity: function convertBalanceToUnderlyingReturns(uint256 n) returns()
func (_EulerToken *EulerTokenSession) ConvertBalanceToUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.ConvertBalanceToUnderlyingReturns(&_EulerToken.TransactOpts, n)
}

// ConvertBalanceToUnderlyingReturns is a paid mutator transaction binding the contract method 0x2a646f07.
//
// Solidity: function convertBalanceToUnderlyingReturns(uint256 n) returns()
func (_EulerToken *EulerTokenTransactorSession) ConvertBalanceToUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _EulerToken.Contract.ConvertBalanceToUnderlyingReturns(&_EulerToken.TransactOpts, n)
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

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_EulerToken *EulerTokenTransactor) UnderlyingAssetReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _EulerToken.contract.Transact(opts, "underlyingAssetReturns", a)
}

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_EulerToken *EulerTokenSession) UnderlyingAssetReturns(a common.Address) (*types.Transaction, error) {
	return _EulerToken.Contract.UnderlyingAssetReturns(&_EulerToken.TransactOpts, a)
}

// UnderlyingAssetReturns is a paid mutator transaction binding the contract method 0x4167cc2b.
//
// Solidity: function underlyingAssetReturns(address a) returns()
func (_EulerToken *EulerTokenTransactorSession) UnderlyingAssetReturns(a common.Address) (*types.Transaction, error) {
	return _EulerToken.Contract.UnderlyingAssetReturns(&_EulerToken.TransactOpts, a)
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
