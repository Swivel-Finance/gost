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
const PendleABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeemAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemAfterExpiryCalled\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"forgeId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"p\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokensCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBought\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"}],\"name\":\"swapExactTokensForTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleBin is the compiled bytecode used for deploying new contracts.
var PendleBin = "0x608060405234801561001057600080fd5b506109ba806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063199c4c341461005c57806338ed17391461008e578063b2f2260b146100be578063cc1d44db146100da578063d230566c1461010b575b600080fd5b610076600480360381019061007191906104a3565b610127565b604051610085939291906104e9565b60405180910390f35b6100a860048036038101906100a391906105b1565b610151565b6040516100b59190610709565b60405180910390f35b6100d860048036038101906100d3919061087a565b61028f565b005b6100f460048036038101906100ef91906104a3565b6102a9565b6040516101029291906108dc565b60405180910390f35b61012560048036038101906101209190610931565b6102cd565b005b60016020528060005260406000206000915090508060000154908060010154908060030154905083565b60606040518060800160405280888152602001878152602001868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200183815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201908051906020019061022492919061033d565b5060608201518160030155905050600080548060200260200160405190810160405280929190818152602001828054801561027e57602002820191906000526020600020905b81548152602001906001019080831161026a575b505050505090509695505050505050565b80600090805190602001906102a59291906103c7565b5050565b60026020528060005260406000206000915090508060000154908060010154905082565b604051806040016040528084815260200182815250600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050505050565b8280548282559060005260206000209081019282156103b6579160200282015b828111156103b55782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509160200191906001019061035d565b5b5090506103c39190610414565b5090565b828054828255906000526020600020908101928215610403579160200282015b828111156104025782518255916020019190600101906103e7565b5b5090506104109190610414565b5090565b5b8082111561042d576000816000905550600101610415565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061047082610445565b9050919050565b61048081610465565b811461048b57600080fd5b50565b60008135905061049d81610477565b92915050565b6000602082840312156104b9576104b861043b565b5b60006104c78482850161048e565b91505092915050565b6000819050919050565b6104e3816104d0565b82525050565b60006060820190506104fe60008301866104da565b61050b60208301856104da565b61051860408301846104da565b949350505050565b610529816104d0565b811461053457600080fd5b50565b60008135905061054681610520565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126105715761057061054c565b5b8235905067ffffffffffffffff81111561058e5761058d610551565b5b6020830191508360208202830111156105aa576105a9610556565b5b9250929050565b60008060008060008060a087890312156105ce576105cd61043b565b5b60006105dc89828a01610537565b96505060206105ed89828a01610537565b955050604087013567ffffffffffffffff81111561060e5761060d610440565b5b61061a89828a0161055b565b9450945050606061062d89828a0161048e565b925050608061063e89828a01610537565b9150509295509295509295565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610680816104d0565b82525050565b60006106928383610677565b60208301905092915050565b6000602082019050919050565b60006106b68261064b565b6106c08185610656565b93506106cb83610667565b8060005b838110156106fc5781516106e38882610686565b97506106ee8361069e565b9250506001810190506106cf565b5085935050505092915050565b6000602082019050818103600083015261072381846106ab565b905092915050565b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107748261072b565b810181811067ffffffffffffffff821117156107935761079261073c565b5b80604052505050565b60006107a6610431565b90506107b2828261076b565b919050565b600067ffffffffffffffff8211156107d2576107d161073c565b5b602082029050602081019050919050565b60006107f66107f1846107b7565b61079c565b9050808382526020820190506020840283018581111561081957610818610556565b5b835b81811015610842578061082e8882610537565b84526020840193505060208101905061081b565b5050509392505050565b600082601f8301126108615761086061054c565b5b81356108718482602086016107e3565b91505092915050565b6000602082840312156108905761088f61043b565b5b600082013567ffffffffffffffff8111156108ae576108ad610440565b5b6108ba8482850161084c565b91505092915050565b6000819050919050565b6108d6816108c3565b82525050565b60006040820190506108f160008301856108cd565b6108fe60208301846104da565b9392505050565b61090e816108c3565b811461091957600080fd5b50565b60008135905061092b81610905565b92915050565b60008060006060848603121561094a5761094961043b565b5b60006109588682870161091c565b93505060206109698682870161048e565b925050604061097a86828701610537565b915050925092509256fea2646970667358221220fdd91c21e5bf21beee2a15e1b287fe4c3db00a98dce2b4f34baba7a963cd128164736f6c634300080d0033"

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

// RedeemAfterExpiryCalled is a free data retrieval call binding the contract method 0xcc1d44db.
//
// Solidity: function redeemAfterExpiryCalled(address ) view returns(bytes32 forgeId, uint256 maturity)
func (_Pendle *PendleCaller) RedeemAfterExpiryCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	ForgeId  [32]byte
	Maturity *big.Int
}, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "redeemAfterExpiryCalled", arg0)

	outstruct := new(struct {
		ForgeId  [32]byte
		Maturity *big.Int
	})

	outstruct.ForgeId = out[0].([32]byte)
	outstruct.Maturity = out[1].(*big.Int)

	return *outstruct, err

}

// RedeemAfterExpiryCalled is a free data retrieval call binding the contract method 0xcc1d44db.
//
// Solidity: function redeemAfterExpiryCalled(address ) view returns(bytes32 forgeId, uint256 maturity)
func (_Pendle *PendleSession) RedeemAfterExpiryCalled(arg0 common.Address) (struct {
	ForgeId  [32]byte
	Maturity *big.Int
}, error) {
	return _Pendle.Contract.RedeemAfterExpiryCalled(&_Pendle.CallOpts, arg0)
}

// RedeemAfterExpiryCalled is a free data retrieval call binding the contract method 0xcc1d44db.
//
// Solidity: function redeemAfterExpiryCalled(address ) view returns(bytes32 forgeId, uint256 maturity)
func (_Pendle *PendleCallerSession) RedeemAfterExpiryCalled(arg0 common.Address) (struct {
	ForgeId  [32]byte
	Maturity *big.Int
}, error) {
	return _Pendle.Contract.RedeemAfterExpiryCalled(&_Pendle.CallOpts, arg0)
}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 amount, uint256 minimumBought, uint256 deadline)
func (_Pendle *PendleCaller) SwapExactTokensForTokensCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount        *big.Int
	MinimumBought *big.Int
	Deadline      *big.Int
}, error) {
	var out []interface{}
	err := _Pendle.contract.Call(opts, &out, "swapExactTokensForTokensCalled", arg0)

	outstruct := new(struct {
		Amount        *big.Int
		MinimumBought *big.Int
		Deadline      *big.Int
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.MinimumBought = out[1].(*big.Int)
	outstruct.Deadline = out[2].(*big.Int)

	return *outstruct, err

}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 amount, uint256 minimumBought, uint256 deadline)
func (_Pendle *PendleSession) SwapExactTokensForTokensCalled(arg0 common.Address) (struct {
	Amount        *big.Int
	MinimumBought *big.Int
	Deadline      *big.Int
}, error) {
	return _Pendle.Contract.SwapExactTokensForTokensCalled(&_Pendle.CallOpts, arg0)
}

// SwapExactTokensForTokensCalled is a free data retrieval call binding the contract method 0x199c4c34.
//
// Solidity: function swapExactTokensForTokensCalled(address ) view returns(uint256 amount, uint256 minimumBought, uint256 deadline)
func (_Pendle *PendleCallerSession) SwapExactTokensForTokensCalled(arg0 common.Address) (struct {
	Amount        *big.Int
	MinimumBought *big.Int
	Deadline      *big.Int
}, error) {
	return _Pendle.Contract.SwapExactTokensForTokensCalled(&_Pendle.CallOpts, arg0)
}

// RedeemAfterExpiry is a paid mutator transaction binding the contract method 0xd230566c.
//
// Solidity: function redeemAfterExpiry(bytes32 i, address u, uint256 m) returns()
func (_Pendle *PendleTransactor) RedeemAfterExpiry(opts *bind.TransactOpts, i [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "redeemAfterExpiry", i, u, m)
}

// RedeemAfterExpiry is a paid mutator transaction binding the contract method 0xd230566c.
//
// Solidity: function redeemAfterExpiry(bytes32 i, address u, uint256 m) returns()
func (_Pendle *PendleSession) RedeemAfterExpiry(i [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.RedeemAfterExpiry(&_Pendle.TransactOpts, i, u, m)
}

// RedeemAfterExpiry is a paid mutator transaction binding the contract method 0xd230566c.
//
// Solidity: function redeemAfterExpiry(bytes32 i, address u, uint256 m) returns()
func (_Pendle *PendleTransactorSession) RedeemAfterExpiry(i [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.RedeemAfterExpiry(&_Pendle.TransactOpts, i, u, m)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 a, uint256 m, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, a *big.Int, m *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.contract.Transact(opts, "swapExactTokensForTokens", a, m, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 a, uint256 m, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleSession) SwapExactTokensForTokens(a *big.Int, m *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokens(&_Pendle.TransactOpts, a, m, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 a, uint256 m, address[] p, address t, uint256 d) returns(uint256[])
func (_Pendle *PendleTransactorSession) SwapExactTokensForTokens(a *big.Int, m *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Pendle.Contract.SwapExactTokensForTokens(&_Pendle.TransactOpts, a, m, p, t, d)
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
