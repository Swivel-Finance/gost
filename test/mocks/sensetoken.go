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

// SenseTokenMetaData contains all meta data concerning the SenseToken contract.
var SenseTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106aa806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610113578063dea1a7e214610143578063e541efa214610161578063e7ba67741461019257610088565b806323b872dd1461008d57806339100838146100bd5780636521b96a146100d95780636f307dc3146100f5575b600080fd5b6100a760048036038101906100a291906104bb565b6101ae565b6040516100b49190610529565b60405180910390f35b6100d760048036038101906100d29190610544565b6102a8565b005b6100f360048036038101906100ee919061059d565b6102b2565b005b6100fd6102cf565b60405161010a91906105d9565b60405180910390f35b61012d600480360381019061012891906105f4565b6102f8565b60405161013a9190610630565b60405180910390f35b61014b610345565b60405161015891906105d9565b60405180910390f35b61017b600480360381019061017691906105f4565b61036b565b60405161018992919061064b565b60405180910390f35b6101ac60048036038101906101a791906105f4565b6103af565b005b60006101b86103f2565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600060149054906101000a900460ff169150509392505050565b8060018190555050565b80600060146101000a81548160ff02191690831515021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001549050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061045282610427565b9050919050565b61046281610447565b811461046d57600080fd5b50565b60008135905061047f81610459565b92915050565b6000819050919050565b61049881610485565b81146104a357600080fd5b50565b6000813590506104b58161048f565b92915050565b6000806000606084860312156104d4576104d3610422565b5b60006104e286828701610470565b93505060206104f386828701610470565b9250506040610504868287016104a6565b9150509250925092565b60008115159050919050565b6105238161050e565b82525050565b600060208201905061053e600083018461051a565b92915050565b60006020828403121561055a57610559610422565b5b6000610568848285016104a6565b91505092915050565b61057a8161050e565b811461058557600080fd5b50565b60008135905061059781610571565b92915050565b6000602082840312156105b3576105b2610422565b5b60006105c184828501610588565b91505092915050565b6105d381610447565b82525050565b60006020820190506105ee60008301846105ca565b92915050565b60006020828403121561060a57610609610422565b5b600061061884828501610470565b91505092915050565b61062a81610485565b82525050565b60006020820190506106456000830184610621565b92915050565b600060408201905061066060008301856105ca565b61066d6020830184610621565b939250505056fea26469706673582212208e4fe81c0c1e2adbad258098cc4f515f53bef135515e15d09f3f1c54bf72d83264736f6c634300080d0033",
}

// SenseTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SenseTokenMetaData.ABI instead.
var SenseTokenABI = SenseTokenMetaData.ABI

// SenseTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenseTokenMetaData.Bin instead.
var SenseTokenBin = SenseTokenMetaData.Bin

// DeploySenseToken deploys a new Ethereum contract, binding an instance of SenseToken to it.
func DeploySenseToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SenseToken, error) {
	parsed, err := SenseTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenseTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenseToken{SenseTokenCaller: SenseTokenCaller{contract: contract}, SenseTokenTransactor: SenseTokenTransactor{contract: contract}, SenseTokenFilterer: SenseTokenFilterer{contract: contract}}, nil
}

// SenseToken is an auto generated Go binding around an Ethereum contract.
type SenseToken struct {
	SenseTokenCaller     // Read-only binding to the contract
	SenseTokenTransactor // Write-only binding to the contract
	SenseTokenFilterer   // Log filterer for contract events
}

// SenseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseTokenSession struct {
	Contract     *SenseToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseTokenCallerSession struct {
	Contract *SenseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SenseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseTokenTransactorSession struct {
	Contract     *SenseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SenseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseTokenRaw struct {
	Contract *SenseToken // Generic contract binding to access the raw methods on
}

// SenseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseTokenCallerRaw struct {
	Contract *SenseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SenseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseTokenTransactorRaw struct {
	Contract *SenseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenseToken creates a new instance of SenseToken, bound to a specific deployed contract.
func NewSenseToken(address common.Address, backend bind.ContractBackend) (*SenseToken, error) {
	contract, err := bindSenseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenseToken{SenseTokenCaller: SenseTokenCaller{contract: contract}, SenseTokenTransactor: SenseTokenTransactor{contract: contract}, SenseTokenFilterer: SenseTokenFilterer{contract: contract}}, nil
}

// NewSenseTokenCaller creates a new read-only instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenCaller(address common.Address, caller bind.ContractCaller) (*SenseTokenCaller, error) {
	contract, err := bindSenseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTokenCaller{contract: contract}, nil
}

// NewSenseTokenTransactor creates a new write-only instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseTokenTransactor, error) {
	contract, err := bindSenseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTokenTransactor{contract: contract}, nil
}

// NewSenseTokenFilterer creates a new log filterer instance of SenseToken, bound to a specific deployed contract.
func NewSenseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseTokenFilterer, error) {
	contract, err := bindSenseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseTokenFilterer{contract: contract}, nil
}

// bindSenseToken binds a generic wrapper to an already deployed contract.
func bindSenseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseToken *SenseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseToken.Contract.SenseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseToken *SenseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseToken.Contract.SenseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseToken *SenseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseToken.Contract.SenseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenseToken *SenseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenseToken *SenseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenseToken *SenseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenseToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_SenseToken *SenseTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenseToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_SenseToken *SenseTokenSession) BalanceOfCalled() (common.Address, error) {
	return _SenseToken.Contract.BalanceOfCalled(&_SenseToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_SenseToken *SenseTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _SenseToken.Contract.BalanceOfCalled(&_SenseToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_SenseToken *SenseTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _SenseToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_SenseToken *SenseTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _SenseToken.Contract.TransferFromCalled(&_SenseToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_SenseToken *SenseTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _SenseToken.Contract.TransferFromCalled(&_SenseToken.CallOpts, arg0)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenseToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenSession) Underlying() (common.Address, error) {
	return _SenseToken.Contract.Underlying(&_SenseToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_SenseToken *SenseTokenCallerSession) Underlying() (common.Address, error) {
	return _SenseToken.Contract.Underlying(&_SenseToken.CallOpts)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_SenseToken *SenseTokenTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_SenseToken *SenseTokenSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.BalanceOf(&_SenseToken.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_SenseToken *SenseTokenTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.BalanceOf(&_SenseToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_SenseToken *SenseTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_SenseToken *SenseTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _SenseToken.Contract.BalanceOfReturns(&_SenseToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_SenseToken *SenseTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _SenseToken.Contract.BalanceOfReturns(&_SenseToken.TransactOpts, b)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_SenseToken *SenseTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_SenseToken *SenseTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _SenseToken.Contract.TransferFrom(&_SenseToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_SenseToken *SenseTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _SenseToken.Contract.TransferFrom(&_SenseToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_SenseToken *SenseTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_SenseToken *SenseTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _SenseToken.Contract.TransferFromReturns(&_SenseToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_SenseToken *SenseTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _SenseToken.Contract.TransferFromReturns(&_SenseToken.TransactOpts, b)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _SenseToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.UnderlyingReturns(&_SenseToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_SenseToken *SenseTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _SenseToken.Contract.UnderlyingReturns(&_SenseToken.TransactOpts, a)
}
