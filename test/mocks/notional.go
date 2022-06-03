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

// NotionalABI is the input ABI used to generate the binding from.
const NotionalABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"getMaturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingToken\",\"outputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getUnderlyingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxRedeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"maxRedeemReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NotionalBin is the compiled bytecode used for deploying new contracts.
var NotionalBin = "0x608060405234801561001057600080fd5b50610a2d806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063a7d8cf0311610097578063e16695b511610066578063e16695b5146102bd578063e541efa2146102db578063ee719bc81461030c578063fe6826701461032b57610100565b8063a7d8cf0314610223578063d35147e41461023f578063d905777e1461026f578063dea1a7e21461029f57610100565b80636521b96a116100d35780636521b96a1461018b5780636e553f65146101a757806370a08231146101d7578063813875a51461020757610100565b806323b872dd1461010557806326cca786146101355780633443b7d014610151578063391008381461016f575b600080fd5b61011f600480360381019061011a919061075d565b610347565b60405161012c91906107cb565b60405180910390f35b61014f600480360381019061014a91906107e6565b610441565b005b61015961044b565b6040516101669190610822565b60405180910390f35b610189600480360381019061018491906107e6565b610471565b005b6101a560048036038101906101a09190610869565b61047b565b005b6101c160048036038101906101bc9190610896565b610498565b6040516101ce91906108e5565b60405180910390f35b6101f160048036038101906101ec9190610900565b6104e9565b6040516101fe91906108e5565b60405180910390f35b610221600480360381019061021c91906107e6565b610536565b005b61023d60048036038101906102389190610900565b610540565b005b61025960048036038101906102549190610900565b610583565b60405161026691906108e5565b60405180910390f35b61028960048036038101906102849190610900565b61059b565b60405161029691906108e5565b60405180910390f35b6102a76105e8565b6040516102b49190610822565b60405180910390f35b6102c561060e565b6040516102d291906108e5565b60405180910390f35b6102f560048036038101906102f09190610900565b610618565b60405161030392919061092d565b60405180910390f35b61031461065c565b6040516103229291906109ce565b60405180910390f35b610345600480360381019061034091906107e6565b61068a565b005b6000610351610694565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600260009054906101000a900460ff169150509392505050565b8060048190555050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060038190555050565b80600260006101000a81548160ff02191690831515021790555050565b600082600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600454905092915050565b600081600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003549050919050565b8060058190555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60076020528060005260406000206000915090505481565b600081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506005549050919050565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600154905090565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000915091509091565b8060018190555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106f4826106c9565b9050919050565b610704816106e9565b811461070f57600080fd5b50565b600081359050610721816106fb565b92915050565b6000819050919050565b61073a81610727565b811461074557600080fd5b50565b60008135905061075781610731565b92915050565b600080600060608486031215610776576107756106c4565b5b600061078486828701610712565b935050602061079586828701610712565b92505060406107a686828701610748565b9150509250925092565b60008115159050919050565b6107c5816107b0565b82525050565b60006020820190506107e060008301846107bc565b92915050565b6000602082840312156107fc576107fb6106c4565b5b600061080a84828501610748565b91505092915050565b61081c816106e9565b82525050565b60006020820190506108376000830184610813565b92915050565b610846816107b0565b811461085157600080fd5b50565b6000813590506108638161083d565b92915050565b60006020828403121561087f5761087e6106c4565b5b600061088d84828501610854565b91505092915050565b600080604083850312156108ad576108ac6106c4565b5b60006108bb85828601610748565b92505060206108cc85828601610712565b9150509250929050565b6108df81610727565b82525050565b60006020820190506108fa60008301846108d6565b92915050565b600060208284031215610916576109156106c4565b5b600061092484828501610712565b91505092915050565b60006040820190506109426000830185610813565b61094f60208301846108d6565b9392505050565b6000819050919050565b600061097b610976610971846106c9565b610956565b6106c9565b9050919050565b600061098d82610960565b9050919050565b600061099f82610982565b9050919050565b6109af81610994565b82525050565b6000819050919050565b6109c8816109b5565b82525050565b60006040820190506109e360008301856109a6565b6109f060208301846109bf565b939250505056fea264697066735822122017bfc4b491af693d6b3431db38cbe2dd30d29010e92fdb82d94efed37125c63864736f6c634300080d0033"

// DeployNotional deploys a new Ethereum contract, binding an instance of Notional to it.
func DeployNotional(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Notional, error) {
	parsed, err := abi.JSON(strings.NewReader(NotionalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NotionalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Notional{NotionalCaller: NotionalCaller{contract: contract}, NotionalTransactor: NotionalTransactor{contract: contract}, NotionalFilterer: NotionalFilterer{contract: contract}}, nil
}

// Notional is an auto generated Go binding around an Ethereum contract.
type Notional struct {
	NotionalCaller     // Read-only binding to the contract
	NotionalTransactor // Write-only binding to the contract
	NotionalFilterer   // Log filterer for contract events
}

// NotionalCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotionalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotionalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotionalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotionalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotionalSession struct {
	Contract     *Notional         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotionalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotionalCallerSession struct {
	Contract *NotionalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NotionalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotionalTransactorSession struct {
	Contract     *NotionalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NotionalRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotionalRaw struct {
	Contract *Notional // Generic contract binding to access the raw methods on
}

// NotionalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotionalCallerRaw struct {
	Contract *NotionalCaller // Generic read-only contract binding to access the raw methods on
}

// NotionalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotionalTransactorRaw struct {
	Contract *NotionalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotional creates a new instance of Notional, bound to a specific deployed contract.
func NewNotional(address common.Address, backend bind.ContractBackend) (*Notional, error) {
	contract, err := bindNotional(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Notional{NotionalCaller: NotionalCaller{contract: contract}, NotionalTransactor: NotionalTransactor{contract: contract}, NotionalFilterer: NotionalFilterer{contract: contract}}, nil
}

// NewNotionalCaller creates a new read-only instance of Notional, bound to a specific deployed contract.
func NewNotionalCaller(address common.Address, caller bind.ContractCaller) (*NotionalCaller, error) {
	contract, err := bindNotional(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotionalCaller{contract: contract}, nil
}

// NewNotionalTransactor creates a new write-only instance of Notional, bound to a specific deployed contract.
func NewNotionalTransactor(address common.Address, transactor bind.ContractTransactor) (*NotionalTransactor, error) {
	contract, err := bindNotional(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotionalTransactor{contract: contract}, nil
}

// NewNotionalFilterer creates a new log filterer instance of Notional, bound to a specific deployed contract.
func NewNotionalFilterer(address common.Address, filterer bind.ContractFilterer) (*NotionalFilterer, error) {
	contract, err := bindNotional(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotionalFilterer{contract: contract}, nil
}

// bindNotional binds a generic wrapper to an already deployed contract.
func bindNotional(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NotionalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notional *NotionalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notional.Contract.NotionalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notional *NotionalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notional.Contract.NotionalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notional *NotionalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notional.Contract.NotionalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notional *NotionalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notional.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notional *NotionalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notional.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notional *NotionalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notional.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_Notional *NotionalCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_Notional *NotionalSession) BalanceOfCalled() (common.Address, error) {
	return _Notional.Contract.BalanceOfCalled(&_Notional.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_Notional *NotionalCallerSession) BalanceOfCalled() (common.Address, error) {
	return _Notional.Contract.BalanceOfCalled(&_Notional.CallOpts)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Notional *NotionalCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "depositCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Notional *NotionalSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _Notional.Contract.DepositCalled(&_Notional.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Notional *NotionalCallerSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _Notional.Contract.DepositCalled(&_Notional.CallOpts, arg0)
}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_Notional *NotionalCaller) GetMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "getMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_Notional *NotionalSession) GetMaturity() (*big.Int, error) {
	return _Notional.Contract.GetMaturity(&_Notional.CallOpts)
}

// GetMaturity is a free data retrieval call binding the contract method 0xe16695b5.
//
// Solidity: function getMaturity() view returns(uint256)
func (_Notional *NotionalCallerSession) GetMaturity() (*big.Int, error) {
	return _Notional.Contract.GetMaturity(&_Notional.CallOpts)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_Notional *NotionalCaller) GetUnderlyingToken(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "getUnderlyingToken")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_Notional *NotionalSession) GetUnderlyingToken() (common.Address, *big.Int, error) {
	return _Notional.Contract.GetUnderlyingToken(&_Notional.CallOpts)
}

// GetUnderlyingToken is a free data retrieval call binding the contract method 0xee719bc8.
//
// Solidity: function getUnderlyingToken() view returns(address, int256)
func (_Notional *NotionalCallerSession) GetUnderlyingToken() (common.Address, *big.Int, error) {
	return _Notional.Contract.GetUnderlyingToken(&_Notional.CallOpts)
}

// MaxRedeemCalled is a free data retrieval call binding the contract method 0x3443b7d0.
//
// Solidity: function maxRedeemCalled() view returns(address)
func (_Notional *NotionalCaller) MaxRedeemCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "maxRedeemCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaxRedeemCalled is a free data retrieval call binding the contract method 0x3443b7d0.
//
// Solidity: function maxRedeemCalled() view returns(address)
func (_Notional *NotionalSession) MaxRedeemCalled() (common.Address, error) {
	return _Notional.Contract.MaxRedeemCalled(&_Notional.CallOpts)
}

// MaxRedeemCalled is a free data retrieval call binding the contract method 0x3443b7d0.
//
// Solidity: function maxRedeemCalled() view returns(address)
func (_Notional *NotionalCallerSession) MaxRedeemCalled() (common.Address, error) {
	return _Notional.Contract.MaxRedeemCalled(&_Notional.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Notional *NotionalCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Notional.contract.Call(opts, &out, "transferFromCalled", arg0)

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
func (_Notional *NotionalSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Notional.Contract.TransferFromCalled(&_Notional.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Notional *NotionalCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Notional.Contract.TransferFromCalled(&_Notional.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_Notional *NotionalTransactor) BalanceOf(opts *bind.TransactOpts, b common.Address) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "balanceOf", b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_Notional *NotionalSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _Notional.Contract.BalanceOf(&_Notional.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address b) returns(uint256)
func (_Notional *NotionalTransactorSession) BalanceOf(b common.Address) (*types.Transaction, error) {
	return _Notional.Contract.BalanceOf(&_Notional.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Notional *NotionalTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Notional *NotionalSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.BalanceOfReturns(&_Notional.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Notional *NotionalTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.BalanceOfReturns(&_Notional.TransactOpts, b)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Notional *NotionalTransactor) Deposit(opts *bind.TransactOpts, a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "deposit", a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Notional *NotionalSession) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Notional.Contract.Deposit(&_Notional.TransactOpts, a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Notional *NotionalTransactorSession) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Notional.Contract.Deposit(&_Notional.TransactOpts, a, r)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_Notional *NotionalTransactor) DepositReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "depositReturns", a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_Notional *NotionalSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.DepositReturns(&_Notional.TransactOpts, a)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 a) returns()
func (_Notional *NotionalTransactorSession) DepositReturns(a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.DepositReturns(&_Notional.TransactOpts, a)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_Notional *NotionalTransactor) GetMaturityReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "getMaturityReturns", m)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_Notional *NotionalSession) GetMaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.GetMaturityReturns(&_Notional.TransactOpts, m)
}

// GetMaturityReturns is a paid mutator transaction binding the contract method 0xfe682670.
//
// Solidity: function getMaturityReturns(uint256 m) returns()
func (_Notional *NotionalTransactorSession) GetMaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.GetMaturityReturns(&_Notional.TransactOpts, m)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_Notional *NotionalTransactor) GetUnderlyingTokenReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "getUnderlyingTokenReturns", a)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_Notional *NotionalSession) GetUnderlyingTokenReturns(a common.Address) (*types.Transaction, error) {
	return _Notional.Contract.GetUnderlyingTokenReturns(&_Notional.TransactOpts, a)
}

// GetUnderlyingTokenReturns is a paid mutator transaction binding the contract method 0xa7d8cf03.
//
// Solidity: function getUnderlyingTokenReturns(address a) returns()
func (_Notional *NotionalTransactorSession) GetUnderlyingTokenReturns(a common.Address) (*types.Transaction, error) {
	return _Notional.Contract.GetUnderlyingTokenReturns(&_Notional.TransactOpts, a)
}

// MaxRedeem is a paid mutator transaction binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address o) returns(uint256)
func (_Notional *NotionalTransactor) MaxRedeem(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "maxRedeem", o)
}

// MaxRedeem is a paid mutator transaction binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address o) returns(uint256)
func (_Notional *NotionalSession) MaxRedeem(o common.Address) (*types.Transaction, error) {
	return _Notional.Contract.MaxRedeem(&_Notional.TransactOpts, o)
}

// MaxRedeem is a paid mutator transaction binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address o) returns(uint256)
func (_Notional *NotionalTransactorSession) MaxRedeem(o common.Address) (*types.Transaction, error) {
	return _Notional.Contract.MaxRedeem(&_Notional.TransactOpts, o)
}

// MaxRedeemReturns is a paid mutator transaction binding the contract method 0x813875a5.
//
// Solidity: function maxRedeemReturns(uint256 a) returns()
func (_Notional *NotionalTransactor) MaxRedeemReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "maxRedeemReturns", a)
}

// MaxRedeemReturns is a paid mutator transaction binding the contract method 0x813875a5.
//
// Solidity: function maxRedeemReturns(uint256 a) returns()
func (_Notional *NotionalSession) MaxRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.MaxRedeemReturns(&_Notional.TransactOpts, a)
}

// MaxRedeemReturns is a paid mutator transaction binding the contract method 0x813875a5.
//
// Solidity: function maxRedeemReturns(uint256 a) returns()
func (_Notional *NotionalTransactorSession) MaxRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.MaxRedeemReturns(&_Notional.TransactOpts, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Notional *NotionalTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Notional *NotionalSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.TransferFrom(&_Notional.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Notional *NotionalTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Notional.Contract.TransferFrom(&_Notional.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Notional *NotionalTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Notional.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Notional *NotionalSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Notional.Contract.TransferFromReturns(&_Notional.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Notional *NotionalTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Notional.Contract.TransferFromReturns(&_Notional.TransactOpts, b)
}
