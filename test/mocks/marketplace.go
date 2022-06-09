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

// MarketPlaceABI is the input ABI used to generate the binding from.
const MarketPlaceABI = "[{\"inputs\":[],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"zcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b506105ed806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632db46d281461005c57806357e6c2f41461008d5780636b079fbf146100ab578063ca1695f0146100c7578063f3ec0d2c146100f7575b600080fd5b610076600480360381019061007191906103b5565b610113565b604051610084929190610417565b60405180910390f35b610095610144565b6040516100a2919061045b565b60405180910390f35b6100c560048036038101906100c091906104a2565b610157565b005b6100e160048036038101906100dc919061053a565b6101e9565b6040516100ee919061059c565b60405180910390f35b610111600480360381019061010c91906103b5565b61030f565b005b60036020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b600260149054906101000a900460ff1681565b80156101a35781600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101e5565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050565b600060405180604001604052808481526020018360ff16815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360ff16021790555090505060008260ff16036102e457600260149054906101000a900460ff16156102ba57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610308565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610308565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b9392505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061038282610357565b9050919050565b61039281610377565b811461039d57600080fd5b50565b6000813590506103af81610389565b92915050565b6000602082840312156103cb576103ca610352565b5b60006103d9848285016103a0565b91505092915050565b6000819050919050565b6103f5816103e2565b82525050565b600060ff82169050919050565b610411816103fb565b82525050565b600060408201905061042c60008301856103ec565b6104396020830184610408565b9392505050565b60008115159050919050565b61045581610440565b82525050565b6000602082019050610470600083018461044c565b92915050565b61047f81610440565b811461048a57600080fd5b50565b60008135905061049c81610476565b92915050565b600080604083850312156104b9576104b8610352565b5b60006104c7858286016103a0565b92505060206104d88582860161048d565b9150509250929050565b6104eb816103e2565b81146104f657600080fd5b50565b600081359050610508816104e2565b92915050565b610517816103fb565b811461052257600080fd5b50565b6000813590506105348161050e565b92915050565b60008060006060848603121561055357610552610352565b5b6000610561868287016103a0565b9350506020610572868287016104f9565b925050604061058386828701610525565b9150509250925092565b61059681610377565b82525050565b60006020820190506105b1600083018461058d565b9291505056fea2646970667358221220700e7529d1cd41bfd9bdceaa85685a4de21dd0f318495e4d6a86ee050c3c880f64736f6c634300080d0033"

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketPlaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x57e6c2f4.
//
// Solidity: function isAuthorized() view returns(bool)
func (_MarketPlace *MarketPlaceCaller) IsAuthorized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "isAuthorized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0x57e6c2f4.
//
// Solidity: function isAuthorized() view returns(bool)
func (_MarketPlace *MarketPlaceSession) IsAuthorized() (bool, error) {
	return _MarketPlace.Contract.IsAuthorized(&_MarketPlace.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x57e6c2f4.
//
// Solidity: function isAuthorized() view returns(bool)
func (_MarketPlace *MarketPlaceCallerSession) IsAuthorized() (bool, error) {
	return _MarketPlace.Contract.IsAuthorized(&_MarketPlace.CallOpts)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceCaller) MarketsCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "marketsCalled", arg0)

	outstruct := new(struct {
		Maturity  *big.Int
		Precision uint8
	})

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Precision = out[1].(uint8)

	return *outstruct, err

}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceSession) MarketsCalled(arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceCallerSession) MarketsCalled(arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// Markets is a paid mutator transaction binding the contract method 0xca1695f0.
//
// Solidity: function markets(address u, uint256 m, uint8 p) returns(address)
func (_MarketPlace *MarketPlaceTransactor) Markets(opts *bind.TransactOpts, u common.Address, m *big.Int, p uint8) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "markets", u, m, p)
}

// Markets is a paid mutator transaction binding the contract method 0xca1695f0.
//
// Solidity: function markets(address u, uint256 m, uint8 p) returns(address)
func (_MarketPlace *MarketPlaceSession) Markets(u common.Address, m *big.Int, p uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.TransactOpts, u, m, p)
}

// Markets is a paid mutator transaction binding the contract method 0xca1695f0.
//
// Solidity: function markets(address u, uint256 m, uint8 p) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) Markets(u common.Address, m *big.Int, p uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.TransactOpts, u, m, p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_MarketPlace *MarketPlaceTransactor) PrincipalTokenReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "principalTokenReturns", p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_MarketPlace *MarketPlaceSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.PrincipalTokenReturns(&_MarketPlace.TransactOpts, p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_MarketPlace *MarketPlaceTransactorSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.PrincipalTokenReturns(&_MarketPlace.TransactOpts, p)
}

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x6b079fbf.
//
// Solidity: function zcTokenReturns(address z, bool authorized) returns()
func (_MarketPlace *MarketPlaceTransactor) ZcTokenReturns(opts *bind.TransactOpts, z common.Address, authorized bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "zcTokenReturns", z, authorized)
}

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x6b079fbf.
//
// Solidity: function zcTokenReturns(address z, bool authorized) returns()
func (_MarketPlace *MarketPlaceSession) ZcTokenReturns(z common.Address, authorized bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenReturns(&_MarketPlace.TransactOpts, z, authorized)
}

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x6b079fbf.
//
// Solidity: function zcTokenReturns(address z, bool authorized) returns()
func (_MarketPlace *MarketPlaceTransactorSession) ZcTokenReturns(z common.Address, authorized bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenReturns(&_MarketPlace.TransactOpts, z, authorized)
}
