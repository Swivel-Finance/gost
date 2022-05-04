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

// PendleDataABI is the input ABI used to generate the binding from.
const PendleDataABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"isValidOT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"o\",\"type\":\"bool\"}],\"name\":\"isValidOTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"isValidXYT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"x\",\"type\":\"bool\"}],\"name\":\"isValidXYTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"xytTokens\",\"outputs\":[{\"internalType\":\"contractIPendleYieldToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"xytTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleDataBin is the compiled bytecode used for deploying new contracts.
var PendleDataBin = "0x608060405234801561001057600080fd5b506105a5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806305fc12e3146100675780633815f2a2146100835780634ec8911c1461009f5780635dc65934146100cf57806394c0c1e0146100ff578063a10cae601461012f575b600080fd5b610081600480360381019061007c9190610354565b61014b565b005b61009d60048036038101906100989190610354565b610168565b005b6100b960048036038101906100b4919061044b565b610184565b6040516100c691906104ad565b60405180910390f35b6100e960048036038101906100e4919061044b565b6101ee565b6040516100f69190610527565b60405180910390f35b6101196004803603810190610114919061044b565b61026b565b60405161012691906104ad565b60405180910390f35b61014960048036038101906101449190610542565b6102d3565b005b80600060016101000a81548160ff02191690831515021790555050565b806000806101000a81548160ff02191690831515021790555050565b60008360018190555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600381905550600060019054906101000a900460ff1690509392505050565b60008360078190555082600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600981905550600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509392505050565b60008360048190555082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160068190555060008054906101000a900460ff1690509392505050565b80600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b60008115159050919050565b6103318161031c565b811461033c57600080fd5b50565b60008135905061034e81610328565b92915050565b60006020828403121561036a57610369610317565b5b60006103788482850161033f565b91505092915050565b6000819050919050565b61039481610381565b811461039f57600080fd5b50565b6000813590506103b18161038b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103e2826103b7565b9050919050565b6103f2816103d7565b81146103fd57600080fd5b50565b60008135905061040f816103e9565b92915050565b6000819050919050565b61042881610415565b811461043357600080fd5b50565b6000813590506104458161041f565b92915050565b60008060006060848603121561046457610463610317565b5b6000610472868287016103a2565b935050602061048386828701610400565b925050604061049486828701610436565b9150509250925092565b6104a78161031c565b82525050565b60006020820190506104c2600083018461049e565b92915050565b6000819050919050565b60006104ed6104e86104e3846103b7565b6104c8565b6103b7565b9050919050565b60006104ff826104d2565b9050919050565b6000610511826104f4565b9050919050565b61052181610506565b82525050565b600060208201905061053c6000830184610518565b92915050565b60006020828403121561055857610557610317565b5b600061056684828501610400565b9150509291505056fea2646970667358221220d0b55e8bbb30f41dd8a2d69d00d3024f04325ede4a0a86702abf95ad3b7d443664736f6c634300080d0033"

// DeployPendleData deploys a new Ethereum contract, binding an instance of PendleData to it.
func DeployPendleData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleData, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleData{PendleDataCaller: PendleDataCaller{contract: contract}, PendleDataTransactor: PendleDataTransactor{contract: contract}, PendleDataFilterer: PendleDataFilterer{contract: contract}}, nil
}

// PendleData is an auto generated Go binding around an Ethereum contract.
type PendleData struct {
	PendleDataCaller     // Read-only binding to the contract
	PendleDataTransactor // Write-only binding to the contract
	PendleDataFilterer   // Log filterer for contract events
}

// PendleDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleDataSession struct {
	Contract     *PendleData       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleDataCallerSession struct {
	Contract *PendleDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PendleDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleDataTransactorSession struct {
	Contract     *PendleDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PendleDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleDataRaw struct {
	Contract *PendleData // Generic contract binding to access the raw methods on
}

// PendleDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleDataCallerRaw struct {
	Contract *PendleDataCaller // Generic read-only contract binding to access the raw methods on
}

// PendleDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleDataTransactorRaw struct {
	Contract *PendleDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleData creates a new instance of PendleData, bound to a specific deployed contract.
func NewPendleData(address common.Address, backend bind.ContractBackend) (*PendleData, error) {
	contract, err := bindPendleData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleData{PendleDataCaller: PendleDataCaller{contract: contract}, PendleDataTransactor: PendleDataTransactor{contract: contract}, PendleDataFilterer: PendleDataFilterer{contract: contract}}, nil
}

// NewPendleDataCaller creates a new read-only instance of PendleData, bound to a specific deployed contract.
func NewPendleDataCaller(address common.Address, caller bind.ContractCaller) (*PendleDataCaller, error) {
	contract, err := bindPendleData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleDataCaller{contract: contract}, nil
}

// NewPendleDataTransactor creates a new write-only instance of PendleData, bound to a specific deployed contract.
func NewPendleDataTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleDataTransactor, error) {
	contract, err := bindPendleData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleDataTransactor{contract: contract}, nil
}

// NewPendleDataFilterer creates a new log filterer instance of PendleData, bound to a specific deployed contract.
func NewPendleDataFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleDataFilterer, error) {
	contract, err := bindPendleData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleDataFilterer{contract: contract}, nil
}

// bindPendleData binds a generic wrapper to an already deployed contract.
func bindPendleData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleData *PendleDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleData.Contract.PendleDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleData *PendleDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleData.Contract.PendleDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleData *PendleDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleData.Contract.PendleDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleData *PendleDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleData *PendleDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleData *PendleDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleData.Contract.contract.Transact(opts, method, params...)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactor) IsValidOT(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidOT", f, u, m)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataSession) IsValidOT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactorSession) IsValidOT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataTransactor) IsValidOTReturns(opts *bind.TransactOpts, o bool) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidOTReturns", o)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataSession) IsValidOTReturns(o bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOTReturns(&_PendleData.TransactOpts, o)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataTransactorSession) IsValidOTReturns(o bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOTReturns(&_PendleData.TransactOpts, o)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactor) IsValidXYT(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidXYT", f, u, m)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataSession) IsValidXYT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactorSession) IsValidXYT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataTransactor) IsValidXYTReturns(opts *bind.TransactOpts, x bool) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidXYTReturns", x)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataSession) IsValidXYTReturns(x bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYTReturns(&_PendleData.TransactOpts, x)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataTransactorSession) IsValidXYTReturns(x bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYTReturns(&_PendleData.TransactOpts, x)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataTransactor) XytTokens(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "xytTokens", f, u, m)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataSession) XytTokens(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokens(&_PendleData.TransactOpts, f, u, m)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataTransactorSession) XytTokens(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokens(&_PendleData.TransactOpts, f, u, m)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataTransactor) XytTokensReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "xytTokensReturns", a)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataSession) XytTokensReturns(a common.Address) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokensReturns(&_PendleData.TransactOpts, a)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataTransactorSession) XytTokensReturns(a common.Address) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokensReturns(&_PendleData.TransactOpts, a)
}
