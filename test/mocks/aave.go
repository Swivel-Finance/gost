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

// AaveMetaData contains all meta data concerning the Aave contract.
var AaveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"r\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103d2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063d35147e41461003b578063e8eda9df1461006d575b600080fd5b61005560048036038101906100509190610234565b610089565b604051610064939291906102a6565b60405180910390f35b61008760048036038101906100829190610335565b6100e1565b005b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900461ffff16905083565b60405180606001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018261ffff168152506000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548161ffff021916908361ffff16021790555090505050505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610201826101d6565b9050919050565b610211816101f6565b811461021c57600080fd5b50565b60008135905061022e81610208565b92915050565b60006020828403121561024a576102496101d1565b5b60006102588482850161021f565b91505092915050565b61026a816101f6565b82525050565b6000819050919050565b61028381610270565b82525050565b600061ffff82169050919050565b6102a081610289565b82525050565b60006060820190506102bb6000830186610261565b6102c8602083018561027a565b6102d56040830184610297565b949350505050565b6102e681610270565b81146102f157600080fd5b50565b600081359050610303816102dd565b92915050565b61031281610289565b811461031d57600080fd5b50565b60008135905061032f81610309565b92915050565b6000806000806080858703121561034f5761034e6101d1565b5b600061035d8782880161021f565b945050602061036e878288016102f4565b935050604061037f8782880161021f565b925050606061039087828801610320565b9150509295919450925056fea26469706673582212204a1d76525c6b996652c936a1b287636fa2e1e6275d56c1e7bbee31d3958b7f7d64736f6c634300080d0033",
}

// AaveABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveMetaData.ABI instead.
var AaveABI = AaveMetaData.ABI

// AaveBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AaveMetaData.Bin instead.
var AaveBin = AaveMetaData.Bin

// DeployAave deploys a new Ethereum contract, binding an instance of Aave to it.
func DeployAave(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Aave, error) {
	parsed, err := AaveMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AaveBin), backend)
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
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint16 referralCode)
func (_Aave *AaveCaller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "depositCalled", arg0)

	outstruct := new(struct {
		OnBehalfOf   common.Address
		Amount       *big.Int
		ReferralCode uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OnBehalfOf = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReferralCode = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint16 referralCode)
func (_Aave *AaveSession) DepositCalled(arg0 common.Address) (struct {
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
}, error) {
	return _Aave.Contract.DepositCalled(&_Aave.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(address onBehalfOf, uint256 amount, uint16 referralCode)
func (_Aave *AaveCallerSession) DepositCalled(arg0 common.Address) (struct {
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
}, error) {
	return _Aave.Contract.DepositCalled(&_Aave.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address a, uint256 amount, address b, uint16 r) returns()
func (_Aave *AaveTransactor) Deposit(opts *bind.TransactOpts, a common.Address, amount *big.Int, b common.Address, r uint16) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "deposit", a, amount, b, r)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address a, uint256 amount, address b, uint16 r) returns()
func (_Aave *AaveSession) Deposit(a common.Address, amount *big.Int, b common.Address, r uint16) (*types.Transaction, error) {
	return _Aave.Contract.Deposit(&_Aave.TransactOpts, a, amount, b, r)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address a, uint256 amount, address b, uint16 r) returns()
func (_Aave *AaveTransactorSession) Deposit(a common.Address, amount *big.Int, b common.Address, r uint16) (*types.Transaction, error) {
	return _Aave.Contract.Deposit(&_Aave.TransactOpts, a, amount, b, r)
}
