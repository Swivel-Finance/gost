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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[9]\",\"name\":\"m\",\"type\":\"address[9]\"}],\"name\":\"marketsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"principals\",\"outputs\":[{\"internalType\":\"address[9]\",\"name\":\"\",\"type\":\"address[9]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b506104c8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80632db46d28146100465780637794292e14610076578063ef8c64b6146100a6575b600080fd5b610060600480360381019061005b91906102a9565b6100c2565b60405161006d91906102ef565b60405180910390f35b610090600480360381019061008b9190610336565b6100da565b60405161009d9190610421565b60405180910390f35b6100c060048036038101906100bb9190610464565b61015e565b005b60096020528060005260406000206000915090505481565b6100e2610173565b6000600980602002604051908101604052809291908260098015610151576020028201915b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610107575b5050505050905092915050565b80600090600961016f929190610196565b5050565b604051806101200160405280600990602082028036833780820191505090505090565b8260098101928215610218579160200282015b8281111561021757823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906101a9565b5b5090506102259190610229565b5090565b5b8082111561024257600081600090555060010161022a565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102768261024b565b9050919050565b6102868161026b565b811461029157600080fd5b50565b6000813590506102a38161027d565b92915050565b6000602082840312156102bf576102be610246565b5b60006102cd84828501610294565b91505092915050565b6000819050919050565b6102e9816102d6565b82525050565b600060208201905061030460008301846102e0565b92915050565b610313816102d6565b811461031e57600080fd5b50565b6000813590506103308161030a565b92915050565b6000806040838503121561034d5761034c610246565b5b600061035b85828601610294565b925050602061036c85828601610321565b9150509250929050565b600060099050919050565b600081905092915050565b6000819050919050565b61039f8161026b565b82525050565b60006103b18383610396565b60208301905092915050565b6000602082019050919050565b6103d381610376565b6103dd8184610381565b92506103e88261038c565b8060005b8381101561041957815161040087826103a5565b965061040b836103bd565b9250506001810190506103ec565b505050505050565b60006101208201905061043760008301846103ca565b92915050565b600080fd5b60008190508260206009028201111561045e5761045d61043d565b5b92915050565b6000610120828403121561047b5761047a610246565b5b600061048984828501610442565b9150509291505056fea2646970667358221220f3e3ceaaa5c1f9057f880eaa12f5ae5416ddef00b441e92e13977e839d12236764736f6c634300080d0033"

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

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) MarketsCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "marketsCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketsCalled(&_MarketPlace.CallOpts, arg0)
}

// Principals is a free data retrieval call binding the contract method 0x7794292e.
//
// Solidity: function principals(address , uint256 ) view returns(address[9])
func (_MarketPlace *MarketPlaceCaller) Principals(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([9]common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "principals", arg0, arg1)

	if err != nil {
		return *new([9]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([9]common.Address)).(*[9]common.Address)

	return out0, err

}

// Principals is a free data retrieval call binding the contract method 0x7794292e.
//
// Solidity: function principals(address , uint256 ) view returns(address[9])
func (_MarketPlace *MarketPlaceSession) Principals(arg0 common.Address, arg1 *big.Int) ([9]common.Address, error) {
	return _MarketPlace.Contract.Principals(&_MarketPlace.CallOpts, arg0, arg1)
}

// Principals is a free data retrieval call binding the contract method 0x7794292e.
//
// Solidity: function principals(address , uint256 ) view returns(address[9])
func (_MarketPlace *MarketPlaceCallerSession) Principals(arg0 common.Address, arg1 *big.Int) ([9]common.Address, error) {
	return _MarketPlace.Contract.Principals(&_MarketPlace.CallOpts, arg0, arg1)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xef8c64b6.
//
// Solidity: function marketsReturns(address[9] m) returns()
func (_MarketPlace *MarketPlaceTransactor) MarketsReturns(opts *bind.TransactOpts, m [9]common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "marketsReturns", m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xef8c64b6.
//
// Solidity: function marketsReturns(address[9] m) returns()
func (_MarketPlace *MarketPlaceSession) MarketsReturns(m [9]common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketsReturns(&_MarketPlace.TransactOpts, m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xef8c64b6.
//
// Solidity: function marketsReturns(address[9] m) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MarketsReturns(m [9]common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketsReturns(&_MarketPlace.TransactOpts, m)
}
