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

// PendleRouterABI is the input ABI used to generate the binding from.
const PendleRouterABI = "[{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"contractIPendleData\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPendleData\",\"name\":\"x\",\"type\":\"address\"}],\"name\":\"dataReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"tokenizeYield\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenizeYieldAmountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenizeYieldIdCalled\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenizeYieldMaturityCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"tokenizeYieldReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenizeYieldTokenCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenizeYieldUnderlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PendleRouterBin is the compiled bytecode used for deploying new contracts.
var PendleRouterBin = "0x608060405234801561001057600080fd5b50610686806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063bae3b11211610066578063bae3b11214610110578063bd7389701461012c578063d29c2f351461015e578063d8088d031461017c578063e43e0e261461019857610093565b806331d36d8b146100985780633f0d9b07146100b6578063721ea14b146100d457806373d4a13a146100f2575b600080fd5b6100a06101b6565b6040516100ad9190610358565b60405180910390f35b6100be6101bc565b6040516100cb91906103b4565b60405180910390f35b6100dc6101e2565b6040516100e99190610358565b60405180910390f35b6100fa6101e8565b604051610107919061042e565b60405180910390f35b61012a6004803603810190610125919061047a565b610212565b005b61014660048036038101906101419190610509565b61021c565b60405161015593929190610584565b60405180910390f35b6101666102cf565b60405161017391906103b4565b60405180910390f35b610196600480360381019061019191906105f9565b6102f5565b005b6101a0610339565b6040516101ad9190610635565b60405180910390f35b60055481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060008190555050565b60008060008760028190555086600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550856004819055508460058190555083600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080600054925092509250955095509592505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025481565b6000819050919050565b6103528161033f565b82525050565b600060208201905061036d6000830184610349565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061039e82610373565b9050919050565b6103ae81610393565b82525050565b60006020820190506103c960008301846103a5565b92915050565b6000819050919050565b60006103f46103ef6103ea84610373565b6103cf565b610373565b9050919050565b6000610406826103d9565b9050919050565b6000610418826103fb565b9050919050565b6104288161040d565b82525050565b6000602082019050610443600083018461041f565b92915050565b600080fd5b6104578161033f565b811461046257600080fd5b50565b6000813590506104748161044e565b92915050565b6000602082840312156104905761048f610449565b5b600061049e84828501610465565b91505092915050565b6000819050919050565b6104ba816104a7565b81146104c557600080fd5b50565b6000813590506104d7816104b1565b92915050565b6104e681610393565b81146104f157600080fd5b50565b600081359050610503816104dd565b92915050565b600080600080600060a0868803121561052557610524610449565b5b6000610533888289016104c8565b9550506020610544888289016104f4565b945050604061055588828901610465565b935050606061056688828901610465565b9250506080610577888289016104f4565b9150509295509295909350565b600060608201905061059960008301866103a5565b6105a660208301856103a5565b6105b36040830184610349565b949350505050565b60006105c682610393565b9050919050565b6105d6816105bb565b81146105e157600080fd5b50565b6000813590506105f3816105cd565b92915050565b60006020828403121561060f5761060e610449565b5b600061061d848285016105e4565b91505092915050565b61062f816104a7565b82525050565b600060208201905061064a6000830184610626565b9291505056fea26469706673582212209b1d3781d5162abb04bf2a4bbd8c81d2a3ac60734ee6148add990e16a8ecbad764736f6c634300080d0033"

// DeployPendleRouter deploys a new Ethereum contract, binding an instance of PendleRouter to it.
func DeployPendleRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleRouter{PendleRouterCaller: PendleRouterCaller{contract: contract}, PendleRouterTransactor: PendleRouterTransactor{contract: contract}, PendleRouterFilterer: PendleRouterFilterer{contract: contract}}, nil
}

// PendleRouter is an auto generated Go binding around an Ethereum contract.
type PendleRouter struct {
	PendleRouterCaller     // Read-only binding to the contract
	PendleRouterTransactor // Write-only binding to the contract
	PendleRouterFilterer   // Log filterer for contract events
}

// PendleRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleRouterSession struct {
	Contract     *PendleRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleRouterCallerSession struct {
	Contract *PendleRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PendleRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleRouterTransactorSession struct {
	Contract     *PendleRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PendleRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleRouterRaw struct {
	Contract *PendleRouter // Generic contract binding to access the raw methods on
}

// PendleRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleRouterCallerRaw struct {
	Contract *PendleRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PendleRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleRouterTransactorRaw struct {
	Contract *PendleRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleRouter creates a new instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouter(address common.Address, backend bind.ContractBackend) (*PendleRouter, error) {
	contract, err := bindPendleRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleRouter{PendleRouterCaller: PendleRouterCaller{contract: contract}, PendleRouterTransactor: PendleRouterTransactor{contract: contract}, PendleRouterFilterer: PendleRouterFilterer{contract: contract}}, nil
}

// NewPendleRouterCaller creates a new read-only instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterCaller(address common.Address, caller bind.ContractCaller) (*PendleRouterCaller, error) {
	contract, err := bindPendleRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterCaller{contract: contract}, nil
}

// NewPendleRouterTransactor creates a new write-only instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleRouterTransactor, error) {
	contract, err := bindPendleRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterTransactor{contract: contract}, nil
}

// NewPendleRouterFilterer creates a new log filterer instance of PendleRouter, bound to a specific deployed contract.
func NewPendleRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleRouterFilterer, error) {
	contract, err := bindPendleRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleRouterFilterer{contract: contract}, nil
}

// bindPendleRouter binds a generic wrapper to an already deployed contract.
func bindPendleRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouter *PendleRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouter.Contract.PendleRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouter *PendleRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouter.Contract.PendleRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouter *PendleRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouter.Contract.PendleRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouter *PendleRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouter *PendleRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouter *PendleRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouter.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterSession) Data() (common.Address, error) {
	return _PendleRouter.Contract.Data(&_PendleRouter.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(address)
func (_PendleRouter *PendleRouterCallerSession) Data() (common.Address, error) {
	return _PendleRouter.Contract.Data(&_PendleRouter.CallOpts)
}

// TokenizeYieldAmountCalled is a free data retrieval call binding the contract method 0x31d36d8b.
//
// Solidity: function tokenizeYieldAmountCalled() view returns(uint256)
func (_PendleRouter *PendleRouterCaller) TokenizeYieldAmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "tokenizeYieldAmountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenizeYieldAmountCalled is a free data retrieval call binding the contract method 0x31d36d8b.
//
// Solidity: function tokenizeYieldAmountCalled() view returns(uint256)
func (_PendleRouter *PendleRouterSession) TokenizeYieldAmountCalled() (*big.Int, error) {
	return _PendleRouter.Contract.TokenizeYieldAmountCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldAmountCalled is a free data retrieval call binding the contract method 0x31d36d8b.
//
// Solidity: function tokenizeYieldAmountCalled() view returns(uint256)
func (_PendleRouter *PendleRouterCallerSession) TokenizeYieldAmountCalled() (*big.Int, error) {
	return _PendleRouter.Contract.TokenizeYieldAmountCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldIdCalled is a free data retrieval call binding the contract method 0xe43e0e26.
//
// Solidity: function tokenizeYieldIdCalled() view returns(bytes32)
func (_PendleRouter *PendleRouterCaller) TokenizeYieldIdCalled(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "tokenizeYieldIdCalled")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenizeYieldIdCalled is a free data retrieval call binding the contract method 0xe43e0e26.
//
// Solidity: function tokenizeYieldIdCalled() view returns(bytes32)
func (_PendleRouter *PendleRouterSession) TokenizeYieldIdCalled() ([32]byte, error) {
	return _PendleRouter.Contract.TokenizeYieldIdCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldIdCalled is a free data retrieval call binding the contract method 0xe43e0e26.
//
// Solidity: function tokenizeYieldIdCalled() view returns(bytes32)
func (_PendleRouter *PendleRouterCallerSession) TokenizeYieldIdCalled() ([32]byte, error) {
	return _PendleRouter.Contract.TokenizeYieldIdCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldMaturityCalled is a free data retrieval call binding the contract method 0x721ea14b.
//
// Solidity: function tokenizeYieldMaturityCalled() view returns(uint256)
func (_PendleRouter *PendleRouterCaller) TokenizeYieldMaturityCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "tokenizeYieldMaturityCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenizeYieldMaturityCalled is a free data retrieval call binding the contract method 0x721ea14b.
//
// Solidity: function tokenizeYieldMaturityCalled() view returns(uint256)
func (_PendleRouter *PendleRouterSession) TokenizeYieldMaturityCalled() (*big.Int, error) {
	return _PendleRouter.Contract.TokenizeYieldMaturityCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldMaturityCalled is a free data retrieval call binding the contract method 0x721ea14b.
//
// Solidity: function tokenizeYieldMaturityCalled() view returns(uint256)
func (_PendleRouter *PendleRouterCallerSession) TokenizeYieldMaturityCalled() (*big.Int, error) {
	return _PendleRouter.Contract.TokenizeYieldMaturityCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldTokenCalled is a free data retrieval call binding the contract method 0x3f0d9b07.
//
// Solidity: function tokenizeYieldTokenCalled() view returns(address)
func (_PendleRouter *PendleRouterCaller) TokenizeYieldTokenCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "tokenizeYieldTokenCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenizeYieldTokenCalled is a free data retrieval call binding the contract method 0x3f0d9b07.
//
// Solidity: function tokenizeYieldTokenCalled() view returns(address)
func (_PendleRouter *PendleRouterSession) TokenizeYieldTokenCalled() (common.Address, error) {
	return _PendleRouter.Contract.TokenizeYieldTokenCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldTokenCalled is a free data retrieval call binding the contract method 0x3f0d9b07.
//
// Solidity: function tokenizeYieldTokenCalled() view returns(address)
func (_PendleRouter *PendleRouterCallerSession) TokenizeYieldTokenCalled() (common.Address, error) {
	return _PendleRouter.Contract.TokenizeYieldTokenCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldUnderlyingCalled is a free data retrieval call binding the contract method 0xd29c2f35.
//
// Solidity: function tokenizeYieldUnderlyingCalled() view returns(address)
func (_PendleRouter *PendleRouterCaller) TokenizeYieldUnderlyingCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleRouter.contract.Call(opts, &out, "tokenizeYieldUnderlyingCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenizeYieldUnderlyingCalled is a free data retrieval call binding the contract method 0xd29c2f35.
//
// Solidity: function tokenizeYieldUnderlyingCalled() view returns(address)
func (_PendleRouter *PendleRouterSession) TokenizeYieldUnderlyingCalled() (common.Address, error) {
	return _PendleRouter.Contract.TokenizeYieldUnderlyingCalled(&_PendleRouter.CallOpts)
}

// TokenizeYieldUnderlyingCalled is a free data retrieval call binding the contract method 0xd29c2f35.
//
// Solidity: function tokenizeYieldUnderlyingCalled() view returns(address)
func (_PendleRouter *PendleRouterCallerSession) TokenizeYieldUnderlyingCalled() (common.Address, error) {
	return _PendleRouter.Contract.TokenizeYieldUnderlyingCalled(&_PendleRouter.CallOpts)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterTransactor) DataReturns(opts *bind.TransactOpts, x common.Address) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "dataReturns", x)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterSession) DataReturns(x common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.DataReturns(&_PendleRouter.TransactOpts, x)
}

// DataReturns is a paid mutator transaction binding the contract method 0xd8088d03.
//
// Solidity: function dataReturns(address x) returns()
func (_PendleRouter *PendleRouterTransactorSession) DataReturns(x common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.DataReturns(&_PendleRouter.TransactOpts, x)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterTransactor) TokenizeYield(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "tokenizeYield", f, u, m, a, t)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterSession) TokenizeYield(f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYield(&_PendleRouter.TransactOpts, f, u, m, a, t)
}

// TokenizeYield is a paid mutator transaction binding the contract method 0xbd738970.
//
// Solidity: function tokenizeYield(bytes32 f, address u, uint256 m, uint256 a, address t) returns(address, address, uint256)
func (_PendleRouter *PendleRouterTransactorSession) TokenizeYield(f [32]byte, u common.Address, m *big.Int, a *big.Int, t common.Address) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYield(&_PendleRouter.TransactOpts, f, u, m, a, t)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterTransactor) TokenizeYieldReturns(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.contract.Transact(opts, "tokenizeYieldReturns", x)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterSession) TokenizeYieldReturns(x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYieldReturns(&_PendleRouter.TransactOpts, x)
}

// TokenizeYieldReturns is a paid mutator transaction binding the contract method 0xbae3b112.
//
// Solidity: function tokenizeYieldReturns(uint256 x) returns()
func (_PendleRouter *PendleRouterTransactorSession) TokenizeYieldReturns(x *big.Int) (*types.Transaction, error) {
	return _PendleRouter.Contract.TokenizeYieldReturns(&_PendleRouter.TransactOpts, x)
}
