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

// PendleABI is the input ABI used to generate the binding from.
const PendleABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"o\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"p\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokensCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"}],\"name\":\"swapExactTokensForTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleBin is the compiled bytecode used for deploying new contracts.
var PendleBin = "0x608060405234801561001057600080fd5b50610802806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063199c4c341461004657806338ed173914610078578063b2f2260b146100a8575b600080fd5b610060600480360381019061005b91906103ac565b6100c4565b60405161006f939291906103f2565b60405180910390f35b610092600480360381019061008d91906104ba565b6100ee565b60405161009f9190610612565b60405180910390f35b6100c260048036038101906100bd9190610783565b61022c565b005b60016020528060005260406000206000915090508060000154908060010154908060030154905083565b60606040518060800160405280888152602001878152602001868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200183815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906101c1929190610246565b5060608201518160030155905050600080548060200260200160405190810160405280929190818152602001828054801561021b57602002820191906000526020600020905b815481526020019060010190808311610207575b505050505090509695505050505050565b80600090805190602001906102429291906102d0565b5050565b8280548282559060005260206000209081019282156102bf579160200282015b828111156102be5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610266565b5b5090506102cc919061031d565b5090565b82805482825590600052602060002090810192821561030c579160200282015b8281111561030b5782518255916020019190600101906102f0565b5b509050610319919061031d565b5090565b5b8082111561033657600081600090555060010161031e565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103798261034e565b9050919050565b6103898161036e565b811461039457600080fd5b50565b6000813590506103a681610380565b92915050565b6000602082840312156103c2576103c1610344565b5b60006103d084828501610397565b91505092915050565b6000819050919050565b6103ec816103d9565b82525050565b600060608201905061040760008301866103e3565b61041460208301856103e3565b61042160408301846103e3565b949350505050565b610432816103d9565b811461043d57600080fd5b50565b60008135905061044f81610429565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261047a57610479610455565b5b8235905067ffffffffffffffff8111156104975761049661045a565b5b6020830191508360208202830111156104b3576104b261045f565b5b9250929050565b60008060008060008060a087890312156104d7576104d6610344565b5b60006104e589828a01610440565b96505060206104f689828a01610440565b955050604087013567ffffffffffffffff81111561051757610516610349565b5b61052389828a01610464565b9450945050606061053689828a01610397565b925050608061054789828a01610440565b9150509295509295509295565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610589816103d9565b82525050565b600061059b8383610580565b60208301905092915050565b6000602082019050919050565b60006105bf82610554565b6105c9818561055f565b93506105d483610570565b8060005b838110156106055781516105ec888261058f565b97506105f7836105a7565b9250506001810190506105d8565b5085935050505092915050565b6000602082019050818103600083015261062c81846105b4565b905092915050565b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61067d82610634565b810181811067ffffffffffffffff8211171561069c5761069b610645565b5b80604052505050565b60006106af61033a565b90506106bb8282610674565b919050565b600067ffffffffffffffff8211156106db576106da610645565b5b602082029050602081019050919050565b60006106ff6106fa846106c0565b6106a5565b905080838252602082019050602084028301858111156107225761072161045f565b5b835b8181101561074b57806107378882610440565b845260208401935050602081019050610724565b5050509392505050565b600082601f83011261076a57610769610455565b5b813561077a8482602086016106ec565b91505092915050565b60006020828403121561079957610798610344565b5b600082013567ffffffffffffffff8111156107b7576107b6610349565b5b6107c384828501610755565b9150509291505056fea2646970667358221220d67ff3e499ce07141a62a3de27315214eb0680d9d1c5ee1f1e612d3f1890ab2464736f6c634300080d0033"

// DeployPendle deploys a new Ethereum contract, binding an instance of Pendle to it.
func DeployPendle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pendle, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pendle{PendleCaller: PendleCaller{contract: contract}, PendleTransactor: PendleTransactor{contract: contract}, PendleFilterer: PendleFilterer{contract: contract}}, nil
}

// Pendle is an auto generated Go binding around an Ethereum contract.
type Pendle struct {
	PendleCaller     // Read-only binding to the contract
	PendleTransactor // Write-only binding to the contract
	PendleFilterer   // Log filterer for contract events
}

// PendleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleSession struct {
	Contract     *Pendle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleCallerSession struct {
	Contract *PendleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PendleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleTransactorSession struct {
	Contract     *PendleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleRaw struct {
	Contract *Pendle // Generic contract binding to access the raw methods on
}

// PendleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleCallerRaw struct {
	Contract *PendleCaller // Generic read-only contract binding to access the raw methods on
}

// PendleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleTransactorRaw struct {
	Contract *PendleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendle creates a new instance of Pendle, bound to a specific deployed contract.
func NewPendle(address common.Address, backend bind.ContractBackend) (*Pendle, error) {
	contract, err := bindPendle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pendle{PendleCaller: PendleCaller{contract: contract}, PendleTransactor: PendleTransactor{contract: contract}, PendleFilterer: PendleFilterer{contract: contract}}, nil
}

// NewPendleCaller creates a new read-only instance of Pendle, bound to a specific deployed contract.
func NewPendleCaller(address common.Address, caller bind.ContractCaller) (*PendleCaller, error) {
	contract, err := bindPendle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleCaller{contract: contract}, nil
}

// NewPendleTransactor creates a new write-only instance of Pendle, bound to a specific deployed contract.
func NewPendleTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleTransactor, error) {
	contract, err := bindPendle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleTransactor{contract: contract}, nil
}

// NewPendleFilterer creates a new log filterer instance of Pendle, bound to a specific deployed contract.
func NewPendleFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleFilterer, error) {
	contract, err := bindPendle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleFilterer{contract: contract}, nil
}

// bindPendle binds a generic wrapper to an already deployed contract.
func bindPendle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pendle *PendleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pendle.Contract.PendleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pendle *PendleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pendle.Contract.PendleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pendle *PendleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pendle.Contract.PendleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pendle *PendleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pendle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pendle *PendleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pendle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pendle *PendleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pendle.Contract.contract.Transact(opts, method, params...)
}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 amount)
func (_Pendle *PendleCaller) SwapExactTokensForTokensCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id      *big.Int
	TokenIn *big.Int
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "swapExactTokensForTokensCalled", arg0)

	outstruct := new(struct {
		Id      *big.Int
		TokenIn *big.Int
		Amount  *big.Int
	})

	outstruct.Id = out[0].(*big.Int)
	outstruct.TokenIn = out[1].(*big.Int)
	outstruct.Amount = out[2].(*big.Int)

	return *outstruct, err

}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 amount)
func (_Pendle *PendleSession) SwapExactTokensForTokensCalled(arg0 common.Address) (struct {
	Id      *big.Int
	TokenIn *big.Int
	Amount  *big.Int
}, error) {
	return _Pendle.Contract.SwapExactTokensForTokensCalled(&_Pendle.CallOpts, arg0)
}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 amount)
func (_Pendle *PendleCallerSession) SwapExactTokensForTokensCalled(arg0 common.Address) (struct {
	Id      *big.Int
	TokenIn *big.Int
	Amount  *big.Int
}, error) {
	return _Pendle.Contract.SwapExactTokensForTokensCalled(&_Pendle.CallOpts, arg0)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "swapExactTokensForTokens", i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokens(&_Pendle.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleTransactorSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokens(&_Pendle.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Pendle *PendleTransactor) SwapExactTokensForTokensReturns(opts *bind.TransactOpts, r []*big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "swapExactTokensForTokensReturns", r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Pendle *PendleSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokensReturns(&_Pendle.TransactOpts, r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Pendle *PendleTransactorSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokensReturns(&_Pendle.TransactOpts, r)
}
