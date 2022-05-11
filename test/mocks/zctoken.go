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

// ZcTokenMetaData contains all meta data concerning the ZcToken contract.
var ZcTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806327e235e31461006757806340c10f191461009757806370a08231146100c7578063c0c6e5ab146100f7578063ee4db57014610113578063fdfe5f4d14610143575b600080fd5b610081600480360381019061007c919061033f565b61015f565b60405161008e9190610385565b60405180910390f35b6100b160048036038101906100ac91906103cc565b610177565b6040516100be9190610427565b60405180910390f35b6100e160048036038101906100dc919061033f565b610218565b6040516100ee9190610385565b60405180910390f35b610111600480360381019061010c91906103cc565b610260565b005b61012d6004803603810190610128919061033f565b6102a7565b60405161013a9190610385565b60405180910390f35b61015d6004803603810190610158919061046e565b6102bf565b005b60006020528060005260406000206000915090505481565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160009054906101000a900460ff16905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60026020528060005260406000206000915090505481565b80600160006101000a81548160ff02191690831515021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061030c826102e1565b9050919050565b61031c81610301565b811461032757600080fd5b50565b60008135905061033981610313565b92915050565b600060208284031215610355576103546102dc565b5b60006103638482850161032a565b91505092915050565b6000819050919050565b61037f8161036c565b82525050565b600060208201905061039a6000830184610376565b92915050565b6103a98161036c565b81146103b457600080fd5b50565b6000813590506103c6816103a0565b92915050565b600080604083850312156103e3576103e26102dc565b5b60006103f18582860161032a565b9250506020610402858286016103b7565b9150509250929050565b60008115159050919050565b6104218161040c565b82525050565b600060208201905061043c6000830184610418565b92915050565b61044b8161040c565b811461045657600080fd5b50565b60008135905061046881610442565b92915050565b600060208284031215610484576104836102dc565b5b600061049284828501610459565b9150509291505056fea26469706673582212208b34f2e698f24550ba3dd1cececbbfa23050796bcf04c668a260499468e3c75064736f6c634300080d0033",
}

// ZcTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ZcTokenMetaData.ABI instead.
var ZcTokenABI = ZcTokenMetaData.ABI

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZcTokenMetaData.Bin instead.
var ZcTokenBin = ZcTokenMetaData.Bin

// DeployZcToken deploys a new Ethereum contract, binding an instance of ZcToken to it.
func DeployZcToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZcToken, error) {
	parsed, err := ZcTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZcTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// ZcToken is an auto generated Go binding around an Ethereum contract.
type ZcToken struct {
	ZcTokenCaller     // Read-only binding to the contract
	ZcTokenTransactor // Write-only binding to the contract
	ZcTokenFilterer   // Log filterer for contract events
}

// ZcTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZcTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZcTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZcTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZcTokenSession struct {
	Contract     *ZcToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZcTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZcTokenCallerSession struct {
	Contract *ZcTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZcTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZcTokenTransactorSession struct {
	Contract     *ZcTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZcTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZcTokenRaw struct {
	Contract *ZcToken // Generic contract binding to access the raw methods on
}

// ZcTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZcTokenCallerRaw struct {
	Contract *ZcTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZcTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZcTokenTransactorRaw struct {
	Contract *ZcTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZcToken creates a new instance of ZcToken, bound to a specific deployed contract.
func NewZcToken(address common.Address, backend bind.ContractBackend) (*ZcToken, error) {
	contract, err := bindZcToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// NewZcTokenCaller creates a new read-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenCaller(address common.Address, caller bind.ContractCaller) (*ZcTokenCaller, error) {
	contract, err := bindZcToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenCaller{contract: contract}, nil
}

// NewZcTokenTransactor creates a new write-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZcTokenTransactor, error) {
	contract, err := bindZcToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransactor{contract: contract}, nil
}

// NewZcTokenFilterer creates a new log filterer instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZcTokenFilterer, error) {
	contract, err := bindZcToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZcTokenFilterer{contract: contract}, nil
}

// bindZcToken binds a generic wrapper to an already deployed contract.
func bindZcToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.ZcTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balanceOf", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, a)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Balances(&_ZcToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Balances(&_ZcToken.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) MintCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "mintCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MintCalled(&_ZcToken.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MintCalled(&_ZcToken.CallOpts, arg0)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_ZcToken *ZcTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, a common.Address, b *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "balanceOfReturns", a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_ZcToken *ZcTokenSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.BalanceOfReturns(&_ZcToken.TransactOpts, a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_ZcToken *ZcTokenTransactorSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.BalanceOfReturns(&_ZcToken.TransactOpts, a, b)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address a, uint256 u) returns(bool)
func (_ZcToken *ZcTokenTransactor) Mint(opts *bind.TransactOpts, a common.Address, u *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mint", a, u)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address a, uint256 u) returns(bool)
func (_ZcToken *ZcTokenSession) Mint(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, a, u)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address a, uint256 u) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Mint(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, a, u)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool r) returns()
func (_ZcToken *ZcTokenTransactor) MintReturns(opts *bind.TransactOpts, r bool) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mintReturns", r)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool r) returns()
func (_ZcToken *ZcTokenSession) MintReturns(r bool) (*types.Transaction, error) {
	return _ZcToken.Contract.MintReturns(&_ZcToken.TransactOpts, r)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool r) returns()
func (_ZcToken *ZcTokenTransactorSession) MintReturns(r bool) (*types.Transaction, error) {
	return _ZcToken.Contract.MintReturns(&_ZcToken.TransactOpts, r)
}
