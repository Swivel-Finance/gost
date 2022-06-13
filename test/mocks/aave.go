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

// AaveABI is the input ABI used to generate the binding from.
const AaveABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"r\",\"type\":\"uint16\"}],\"name\":\"depsoit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AaveBin is the compiled bytecode used for deploying new contracts.
var AaveBin = "0x608060405234801561001057600080fd5b50610497806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632c9f606e146100465780633042a2b814610076578063d35147e414610092575b600080fd5b610060600480360381019061005b91906102e9565b6100c4565b60405161006d919061036b565b60405180910390f35b610090600480360381019061008b91906103b2565b6101b0565b005b6100ac60048036038101906100a791906103df565b6101cc565b6040516100bb9392919061042a565b60405180910390f35b600060405180606001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018361ffff16815250600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015590505060008054906101000a900460ff169050949350505050565b806000806101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102468261021b565b9050919050565b6102568161023b565b811461026157600080fd5b50565b6000813590506102738161024d565b92915050565b6000819050919050565b61028c81610279565b811461029757600080fd5b50565b6000813590506102a981610283565b92915050565b600061ffff82169050919050565b6102c6816102af565b81146102d157600080fd5b50565b6000813590506102e3816102bd565b92915050565b6000806000806080858703121561030357610302610216565b5b600061031187828801610264565b945050602061032287828801610264565b93505060406103338782880161029a565b9250506060610344878288016102d4565b91505092959194509250565b60008115159050919050565b61036581610350565b82525050565b6000602082019050610380600083018461035c565b92915050565b61038f81610350565b811461039a57600080fd5b50565b6000813590506103ac81610386565b92915050565b6000602082840312156103c8576103c7610216565b5b60006103d68482850161039d565b91505092915050565b6000602082840312156103f5576103f4610216565b5b600061040384828501610264565b91505092915050565b6104158161023b565b82525050565b61042481610279565b82525050565b600060608201905061043f600083018661040c565b61044c602083018561041b565b610459604083018461041b565b94935050505056fea2646970667358221220a7ec99c6160b513f33e20800c7c74f4e6ef155a84d4d8d22393af91beac6e31864736f6c634300080d0033"

// DeployAave deploys a new Ethereum contract, binding an instance of Aave to it.
func DeployAave(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Aave, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AaveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// Aave is an auto generated Go binding around an Ethereum contract.
type Aave struct {
	AaveCaller     // Read-only binding to the contract
	AaveTransactor // Write-only binding to the contract
	AaveFilterer   // Log filterer for contract events
}

// AaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveSession struct {
	Contract     *Aave             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveCallerSession struct {
	Contract *AaveCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTransactorSession struct {
	Contract     *AaveTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveRaw struct {
	Contract *Aave // Generic contract binding to access the raw methods on
}

// AaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveCallerRaw struct {
	Contract *AaveCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTransactorRaw struct {
	Contract *AaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAave creates a new instance of Aave, bound to a specific deployed contract.
func NewAave(address common.Address, backend bind.ContractBackend) (*Aave, error) {
	contract, err := bindAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// NewAaveCaller creates a new read-only instance of Aave, bound to a specific deployed contract.
func NewAaveCaller(address common.Address, caller bind.ContractCaller) (*AaveCaller, error) {
	contract, err := bindAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveCaller{contract: contract}, nil
}

// NewAaveTransactor creates a new write-only instance of Aave, bound to a specific deployed contract.
func NewAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTransactor, error) {
	contract, err := bindAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTransactor{contract: contract}, nil
}

// NewAaveFilterer creates a new log filterer instance of Aave, bound to a specific deployed contract.
func NewAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFilterer, error) {
	contract, err := bindAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFilterer{contract: contract}, nil
}

// bindAave binds a generic wrapper to an already deployed contract.
func bindAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.AaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transact(opts, method, params...)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint256 minimumAmount)
func (_Aave *AaveCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	OnBehalfOf    common.Address
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "depositCalled", arg0)

	outstruct := new(struct {
		OnBehalfOf    common.Address
		Amount        *big.Int
		MinimumAmount *big.Int
	})

	outstruct.OnBehalfOf = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)
	outstruct.MinimumAmount = out[2].(*big.Int)

	return *outstruct, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint256 minimumAmount)
func (_Aave *AaveSession) DepositCalled(arg0 common.Address) (struct {
	OnBehalfOf    common.Address
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	return _Aave.Contract.DepositCalled(&_Aave.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint256 minimumAmount)
func (_Aave *AaveCallerSession) DepositCalled(arg0 common.Address) (struct {
	OnBehalfOf    common.Address
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	return _Aave.Contract.DepositCalled(&_Aave.CallOpts, arg0)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x3042a2b8.
//
// Solidity: function depositReturns(bool b) returns()
func (_Aave *AaveTransactor) DepositReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "depositReturns", b)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x3042a2b8.
//
// Solidity: function depositReturns(bool b) returns()
func (_Aave *AaveSession) DepositReturns(b bool) (*types.Transaction, error) {
	return _Aave.Contract.DepositReturns(&_Aave.TransactOpts, b)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x3042a2b8.
//
// Solidity: function depositReturns(bool b) returns()
func (_Aave *AaveTransactorSession) DepositReturns(b bool) (*types.Transaction, error) {
	return _Aave.Contract.DepositReturns(&_Aave.TransactOpts, b)
}

// Depsoit is a paid mutator transaction binding the contract method 0x2c9f606e.
//
// Solidity: function depsoit(address a, address b, uint256 amount, uint16 r) returns(bool)
func (_Aave *AaveTransactor) Depsoit(opts *bind.TransactOpts, a common.Address, b common.Address, amount *big.Int, r uint16) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "depsoit", a, b, amount, r)
}

// Depsoit is a paid mutator transaction binding the contract method 0x2c9f606e.
//
// Solidity: function depsoit(address a, address b, uint256 amount, uint16 r) returns(bool)
func (_Aave *AaveSession) Depsoit(a common.Address, b common.Address, amount *big.Int, r uint16) (*types.Transaction, error) {
	return _Aave.Contract.Depsoit(&_Aave.TransactOpts, a, b, amount, r)
}

// Depsoit is a paid mutator transaction binding the contract method 0x2c9f606e.
//
// Solidity: function depsoit(address a, address b, uint256 amount, uint16 r) returns(bool)
func (_Aave *AaveTransactorSession) Depsoit(a common.Address, b common.Address, amount *big.Int, r uint16) (*types.Transaction, error) {
	return _Aave.Contract.Depsoit(&_Aave.TransactOpts, a, b, amount, r)
}
