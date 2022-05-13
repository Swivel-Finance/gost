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

// PendleTokenABI is the input ABI used to generate the binding from.
const PendleTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"expiryReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"yieldTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleTokenBin is the compiled bytecode used for deploying new contracts.
var PendleTokenBin = "0x608060405234801561001057600080fd5b5061070e806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806376d5de851161006657806376d5de8514610157578063dea1a7e214610175578063e184c9be14610193578063e541efa2146101b1578063ffd71860146101e25761009e565b806323b872dd146100a357806339100838146100d357806345c0684b146100ef5780636521b96a1461010b57806370a0823114610127575b600080fd5b6100bd60048036038101906100b8919061051f565b6101fe565b6040516100ca919061058d565b60405180910390f35b6100ed60048036038101906100e891906105a8565b6102f8565b005b610109600480360381019061010491906105d5565b610302565b005b6101256004803603810190610120919061062e565b610345565b005b610141600480360381019061013c91906105d5565b610362565b60405161014e919061066a565b60405180910390f35b61015f6103af565b60405161016c9190610694565b60405180910390f35b61017d6103d8565b60405161018a9190610694565b60405180910390f35b61019b6103fe565b6040516101a8919061066a565b60405180910390f35b6101cb60048036038101906101c691906105d5565b610408565b6040516101d99291906106af565b60405180910390f35b6101fc60048036038101906101f791906105a8565b61044c565b005b6000610208610456565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260009054906101000a900460ff169150509392505050565b8060038190555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600260006101000a81548160ff02191690831515021790555050565b600081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b8060018190555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104b68261048b565b9050919050565b6104c6816104ab565b81146104d157600080fd5b50565b6000813590506104e3816104bd565b92915050565b6000819050919050565b6104fc816104e9565b811461050757600080fd5b50565b600081359050610519816104f3565b92915050565b60008060006060848603121561053857610537610486565b5b6000610546868287016104d4565b9350506020610557868287016104d4565b92505060406105688682870161050a565b9150509250925092565b60008115159050919050565b61058781610572565b82525050565b60006020820190506105a2600083018461057e565b92915050565b6000602082840312156105be576105bd610486565b5b60006105cc8482850161050a565b91505092915050565b6000602082840312156105eb576105ea610486565b5b60006105f9848285016104d4565b91505092915050565b61060b81610572565b811461061657600080fd5b50565b60008135905061062881610602565b92915050565b60006020828403121561064457610643610486565b5b600061065284828501610619565b91505092915050565b610664816104e9565b82525050565b600060208201905061067f600083018461065b565b92915050565b61068e816104ab565b82525050565b60006020820190506106a96000830184610685565b92915050565b60006040820190506106c46000830185610685565b6106d1602083018461065b565b939250505056fea26469706673582212209eb99b6a971f107301d46f35c3a6148d51fc938bdb85c3ab9a308a343be2ec3464736f6c634300080d0033"

// DeployPendleToken deploys a new Ethereum contract, binding an instance of PendleToken to it.
func DeployPendleToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleToken{PendleTokenCaller: PendleTokenCaller{contract: contract}, PendleTokenTransactor: PendleTokenTransactor{contract: contract}, PendleTokenFilterer: PendleTokenFilterer{contract: contract}}, nil
}

// PendleToken is an auto generated Go binding around an Ethereum contract.
type PendleToken struct {
	PendleTokenCaller     // Read-only binding to the contract
	PendleTokenTransactor // Write-only binding to the contract
	PendleTokenFilterer   // Log filterer for contract events
}

// PendleTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleTokenSession struct {
	Contract     *PendleToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleTokenCallerSession struct {
	Contract *PendleTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PendleTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleTokenTransactorSession struct {
	Contract     *PendleTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PendleTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleTokenRaw struct {
	Contract *PendleToken // Generic contract binding to access the raw methods on
}

// PendleTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleTokenCallerRaw struct {
	Contract *PendleTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PendleTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleTokenTransactorRaw struct {
	Contract *PendleTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleToken creates a new instance of PendleToken, bound to a specific deployed contract.
func NewPendleToken(address common.Address, backend bind.ContractBackend) (*PendleToken, error) {
	contract, err := bindPendleToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleToken{PendleTokenCaller: PendleTokenCaller{contract: contract}, PendleTokenTransactor: PendleTokenTransactor{contract: contract}, PendleTokenFilterer: PendleTokenFilterer{contract: contract}}, nil
}

// NewPendleTokenCaller creates a new read-only instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenCaller(address common.Address, caller bind.ContractCaller) (*PendleTokenCaller, error) {
	contract, err := bindPendleToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTokenCaller{contract: contract}, nil
}

// NewPendleTokenTransactor creates a new write-only instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleTokenTransactor, error) {
	contract, err := bindPendleToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTokenTransactor{contract: contract}, nil
}

// NewPendleTokenFilterer creates a new log filterer instance of PendleToken, bound to a specific deployed contract.
func NewPendleTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleTokenFilterer, error) {
	contract, err := bindPendleToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleTokenFilterer{contract: contract}, nil
}

// bindPendleToken binds a generic wrapper to an already deployed contract.
func bindPendleToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleToken *PendleTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleToken.Contract.PendleTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleToken *PendleTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleToken.Contract.PendleTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleToken *PendleTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleToken.Contract.PendleTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleToken *PendleTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleToken *PendleTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleToken *PendleTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_PendleToken *PendleTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_PendleToken *PendleTokenSession) BalanceOfCalled() (common.Address, error) {
	return _PendleToken.Contract.BalanceOfCalled(&_PendleToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_PendleToken *PendleTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _PendleToken.Contract.BalanceOfCalled(&_PendleToken.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenSession) Expiry() (*big.Int, error) {
	return _PendleToken.Contract.Expiry(&_PendleToken.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xe184c9be.
//
// Solidity: function expiry() view returns(uint256)
func (_PendleToken *PendleTokenCallerSession) Expiry() (*big.Int, error) {
	return _PendleToken.Contract.Expiry(&_PendleToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_PendleToken *PendleTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "transferFromCalled", arg0)

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
func (_PendleToken *PendleTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _PendleToken.Contract.TransferFromCalled(&_PendleToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_PendleToken *PendleTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _PendleToken.Contract.TransferFromCalled(&_PendleToken.CallOpts, arg0)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenCaller) YieldToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleToken.contract.Call(opts, &out, "yieldToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenSession) YieldToken() (common.Address, error) {
	return _PendleToken.Contract.YieldToken(&_PendleToken.CallOpts)
}

// YieldToken is a free data retrieval call binding the contract method 0x76d5de85.
//
// Solidity: function yieldToken() view returns(address)
func (_PendleToken *PendleTokenCallerSession) YieldToken() (common.Address, error) {
	return _PendleToken.Contract.YieldToken(&_PendleToken.CallOpts)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_PendleToken *PendleTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_PendleToken *PendleTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.BalanceOf(&_PendleToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_PendleToken *PendleTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.BalanceOf(&_PendleToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_PendleToken *PendleTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_PendleToken *PendleTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.BalanceOfReturns(&_PendleToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_PendleToken *PendleTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.BalanceOfReturns(&_PendleToken.TransactOpts, b)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenTransactor) ExpiryReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "expiryReturns", m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.ExpiryReturns(&_PendleToken.TransactOpts, m)
}

// ExpiryReturns is a paid mutator transaction binding the contract method 0xffd71860.
//
// Solidity: function expiryReturns(uint256 m) returns()
func (_PendleToken *PendleTokenTransactorSession) ExpiryReturns(m *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.ExpiryReturns(&_PendleToken.TransactOpts, m)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_PendleToken *PendleTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_PendleToken *PendleTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.TransferFrom(&_PendleToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_PendleToken *PendleTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _PendleToken.Contract.TransferFrom(&_PendleToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_PendleToken *PendleTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_PendleToken *PendleTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _PendleToken.Contract.TransferFromReturns(&_PendleToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_PendleToken *PendleTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _PendleToken.Contract.TransferFromReturns(&_PendleToken.TransactOpts, b)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenTransactor) YieldTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleToken.contract.Transact(opts, "yieldTokenReturns", a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.YieldTokenReturns(&_PendleToken.TransactOpts, a)
}

// YieldTokenReturns is a paid mutator transaction binding the contract method 0x45c0684b.
//
// Solidity: function yieldTokenReturns(address a) returns()
func (_PendleToken *PendleTokenTransactorSession) YieldTokenReturns(a common.Address) (*types.Transaction, error) {
	return _PendleToken.Contract.YieldTokenReturns(&_PendleToken.TransactOpts, a)
}
