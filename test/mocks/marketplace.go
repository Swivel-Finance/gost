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
const MarketPlaceABI = "[{\"inputs\":[],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"a\",\"type\":\"bool\"}],\"name\":\"setIsAuthorized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"zcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x60806040526000600260146101000a81548160ff02191690831515021790555034801561002b57600080fd5b506106798061003b6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632db46d281461006757806357e6c2f4146100985780636b079fbf146100b65780639e0a8e09146100d2578063ca1695f0146100ee578063f3ec0d2c1461011e575b600080fd5b610081600480360381019061007c9190610414565b61013a565b60405161008f929190610476565b60405180910390f35b6100a061016b565b6040516100ad91906104ba565b60405180910390f35b6100d060048036038101906100cb9190610501565b61017e565b005b6100ec60048036038101906100e79190610541565b610210565b005b610108600480360381019061010391906105c6565b61022d565b6040516101159190610628565b60405180910390f35b61013860048036038101906101339190610414565b61036e565b005b60036020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b600260149054906101000a900460ff1681565b80156101ca5781600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061020c565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050565b80600260146101000a81548160ff02191690831515021790555050565b600060405180604001604052808481526020018360ff16815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360ff16021790555090505060008260ff160361034357600260149054906101000a900460ff1615610319576000600260146101000a81548160ff021916908315150217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610367565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610367565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b9392505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103e1826103b6565b9050919050565b6103f1816103d6565b81146103fc57600080fd5b50565b60008135905061040e816103e8565b92915050565b60006020828403121561042a576104296103b1565b5b6000610438848285016103ff565b91505092915050565b6000819050919050565b61045481610441565b82525050565b600060ff82169050919050565b6104708161045a565b82525050565b600060408201905061048b600083018561044b565b6104986020830184610467565b9392505050565b60008115159050919050565b6104b48161049f565b82525050565b60006020820190506104cf60008301846104ab565b92915050565b6104de8161049f565b81146104e957600080fd5b50565b6000813590506104fb816104d5565b92915050565b60008060408385031215610518576105176103b1565b5b6000610526858286016103ff565b9250506020610537858286016104ec565b9150509250929050565b600060208284031215610557576105566103b1565b5b6000610565848285016104ec565b91505092915050565b61057781610441565b811461058257600080fd5b50565b6000813590506105948161056e565b92915050565b6105a38161045a565b81146105ae57600080fd5b50565b6000813590506105c08161059a565b92915050565b6000806000606084860312156105df576105de6103b1565b5b60006105ed868287016103ff565b93505060206105fe86828701610585565b925050604061060f868287016105b1565b9150509250925092565b610622816103d6565b82525050565b600060208201905061063d6000830184610619565b9291505056fea264697066735822122033e27b60c0c4bbee94971648efd3997f17519b33116b9ffcab2233342fa9d49364736f6c634300080d0033"

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

// SetIsAuthorized is a paid mutator transaction binding the contract method 0x9e0a8e09.
//
// Solidity: function setIsAuthorized(bool a) returns()
func (_MarketPlace *MarketPlaceTransactor) SetIsAuthorized(opts *bind.TransactOpts, a bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "setIsAuthorized", a)
}

// SetIsAuthorized is a paid mutator transaction binding the contract method 0x9e0a8e09.
//
// Solidity: function setIsAuthorized(bool a) returns()
func (_MarketPlace *MarketPlaceSession) SetIsAuthorized(a bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetIsAuthorized(&_MarketPlace.TransactOpts, a)
}

// SetIsAuthorized is a paid mutator transaction binding the contract method 0x9e0a8e09.
//
// Solidity: function setIsAuthorized(bool a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) SetIsAuthorized(a bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetIsAuthorized(&_MarketPlace.TransactOpts, a)
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
