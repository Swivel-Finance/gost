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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"}],\"name\":\"zcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"zcTokenSpoofReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50610595806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632cf443cb1461005c5780632db46d281461007857806368f8d560146100a9578063ca1695f0146100c5578063f3ec0d2c146100f5575b600080fd5b610076600480360381019061007191906103ff565b610111565b005b610092600480360381019061008d91906103ff565b610155565b6040516100a0929190610461565b60405180910390f35b6100c360048036038101906100be91906103ff565b610186565b005b6100df60048036038101906100da91906104e2565b6101ca565b6040516100ec9190610544565b60405180910390f35b61010f600480360381019061010a91906103ff565b610359565b005b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060405180604001604052808481526020018360ff16815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360ff1602179055509050506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008360ff160361032d57600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610302576000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080915050610352565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915050610352565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150505b9392505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103cc826103a1565b9050919050565b6103dc816103c1565b81146103e757600080fd5b50565b6000813590506103f9816103d3565b92915050565b6000602082840312156104155761041461039c565b5b6000610423848285016103ea565b91505092915050565b6000819050919050565b61043f8161042c565b82525050565b600060ff82169050919050565b61045b81610445565b82525050565b60006040820190506104766000830185610436565b6104836020830184610452565b9392505050565b6104938161042c565b811461049e57600080fd5b50565b6000813590506104b08161048a565b92915050565b6104bf81610445565b81146104ca57600080fd5b50565b6000813590506104dc816104b6565b92915050565b6000806000606084860312156104fb576104fa61039c565b5b6000610509868287016103ea565b935050602061051a868287016104a1565b925050604061052b868287016104cd565b9150509250925092565b61053e816103c1565b82525050565b60006020820190506105596000830184610535565b9291505056fea264697066735822122016343ff142d0e5a0492e19ade438ce62aa4330cfb060baa28df58a9ca9b809fe64736f6c634300080d0033"

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

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x68f8d560.
//
// Solidity: function zcTokenReturns(address z) returns()
func (_MarketPlace *MarketPlaceTransactor) ZcTokenReturns(opts *bind.TransactOpts, z common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "zcTokenReturns", z)
}

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x68f8d560.
//
// Solidity: function zcTokenReturns(address z) returns()
func (_MarketPlace *MarketPlaceSession) ZcTokenReturns(z common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenReturns(&_MarketPlace.TransactOpts, z)
}

// ZcTokenReturns is a paid mutator transaction binding the contract method 0x68f8d560.
//
// Solidity: function zcTokenReturns(address z) returns()
func (_MarketPlace *MarketPlaceTransactorSession) ZcTokenReturns(z common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenReturns(&_MarketPlace.TransactOpts, z)
}

// ZcTokenSpoofReturns is a paid mutator transaction binding the contract method 0x2cf443cb.
//
// Solidity: function zcTokenSpoofReturns(address s) returns()
func (_MarketPlace *MarketPlaceTransactor) ZcTokenSpoofReturns(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "zcTokenSpoofReturns", s)
}

// ZcTokenSpoofReturns is a paid mutator transaction binding the contract method 0x2cf443cb.
//
// Solidity: function zcTokenSpoofReturns(address s) returns()
func (_MarketPlace *MarketPlaceSession) ZcTokenSpoofReturns(s common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenSpoofReturns(&_MarketPlace.TransactOpts, s)
}

// ZcTokenSpoofReturns is a paid mutator transaction binding the contract method 0x2cf443cb.
//
// Solidity: function zcTokenSpoofReturns(address s) returns()
func (_MarketPlace *MarketPlaceTransactorSession) ZcTokenSpoofReturns(s common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.ZcTokenSpoofReturns(&_MarketPlace.TransactOpts, s)
}
