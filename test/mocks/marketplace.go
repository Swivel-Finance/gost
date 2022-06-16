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

// MarketPlaceMetaData contains all meta data concerning the MarketPlace contract.
var MarketPlaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausedCalled\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"p\",\"type\":\"bool\"}],\"name\":\"pausedReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"poolsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"}],\"name\":\"zcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"zcTokenSpoofReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610984806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806368f8d5601161007157806368f8d560146101655780637763772a1461018157806386e69b521461019d578063ca1695f0146101cd578063ee3a5c9d146101fd578063f3ec0d2c1461022e576100a9565b80632cf443cb146100ae5780632db46d28146100ca57806331acf244146100fb5780634c5d037a146101195780635ac86ab714610135575b600080fd5b6100c860048036038101906100c391906106c4565b61024a565b005b6100e460048036038101906100df91906106c4565b61028e565b6040516100f2929190610726565b60405180910390f35b6101036102bf565b604051610110919061074f565b60405180910390f35b610133600480360381019061012e91906107a2565b6102d2565b005b61014f600480360381019061014a91906107fb565b6102ef565b60405161015c9190610837565b60405180910390f35b61017f600480360381019061017a91906106c4565b610323565b005b61019b600480360381019061019691906106c4565b610367565b005b6101b760048036038101906101b2919061087e565b6103ab565b6040516101c491906108e0565b60405180910390f35b6101e760048036038101906101e291906108fb565b61045e565b6040516101f491906108e0565b60405180910390f35b610217600480360381019061021291906106c4565b6105ed565b604051610225929190610726565b60405180910390f35b610248600480360381019061024391906106c4565b61061e565b005b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b600660009054906101000a900460ff1681565b80600360146101000a81548160ff02191690831515021790555050565b600081600660006101000a81548160ff021916908360ff160217905550600360149054906101000a900460ff169050919050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060405180604001604052808381526020018560ff16815250600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360ff160217905550905050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509392505050565b600060405180604001604052808481526020018360ff16815250600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff021916908360ff1602179055509050506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008360ff16036105c157600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610596576000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550809150506105e6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150506105e6565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150505b9392505050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16905082565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061069182610666565b9050919050565b6106a181610686565b81146106ac57600080fd5b50565b6000813590506106be81610698565b92915050565b6000602082840312156106da576106d9610661565b5b60006106e8848285016106af565b91505092915050565b6000819050919050565b610704816106f1565b82525050565b600060ff82169050919050565b6107208161070a565b82525050565b600060408201905061073b60008301856106fb565b6107486020830184610717565b9392505050565b60006020820190506107646000830184610717565b92915050565b60008115159050919050565b61077f8161076a565b811461078a57600080fd5b50565b60008135905061079c81610776565b92915050565b6000602082840312156107b8576107b7610661565b5b60006107c68482850161078d565b91505092915050565b6107d88161070a565b81146107e357600080fd5b50565b6000813590506107f5816107cf565b92915050565b60006020828403121561081157610810610661565b5b600061081f848285016107e6565b91505092915050565b6108318161076a565b82525050565b600060208201905061084c6000830184610828565b92915050565b61085b816106f1565b811461086657600080fd5b50565b60008135905061087881610852565b92915050565b60008060006060848603121561089757610896610661565b5b60006108a5868287016107e6565b93505060206108b6868287016106af565b92505060406108c786828701610869565b9150509250925092565b6108da81610686565b82525050565b60006020820190506108f560008301846108d1565b92915050565b60008060006060848603121561091457610913610661565b5b6000610922868287016106af565b935050602061093386828701610869565b9250506040610944868287016107e6565b915050925092509256fea26469706673582212207bb13450c849141f9d9e0645300c8617dc64c182b62540a92834e48e72f3381e64736f6c634300080d0033",
}

// MarketPlaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketPlaceMetaData.ABI instead.
var MarketPlaceABI = MarketPlaceMetaData.ABI

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketPlaceMetaData.Bin instead.
var MarketPlaceBin = MarketPlaceMetaData.Bin

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := MarketPlaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketPlaceBin), backend)
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
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Precision = *abi.ConvertType(out[1], new(uint8)).(*uint8)

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

// PausedCalled is a free data retrieval call binding the contract method 0x31acf244.
//
// Solidity: function pausedCalled() view returns(uint8)
func (_MarketPlace *MarketPlaceCaller) PausedCalled(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "pausedCalled")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PausedCalled is a free data retrieval call binding the contract method 0x31acf244.
//
// Solidity: function pausedCalled() view returns(uint8)
func (_MarketPlace *MarketPlaceSession) PausedCalled() (uint8, error) {
	return _MarketPlace.Contract.PausedCalled(&_MarketPlace.CallOpts)
}

// PausedCalled is a free data retrieval call binding the contract method 0x31acf244.
//
// Solidity: function pausedCalled() view returns(uint8)
func (_MarketPlace *MarketPlaceCallerSession) PausedCalled() (uint8, error) {
	return _MarketPlace.Contract.PausedCalled(&_MarketPlace.CallOpts)
}

// PoolsCalled is a free data retrieval call binding the contract method 0xee3a5c9d.
//
// Solidity: function poolsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceCaller) PoolsCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "poolsCalled", arg0)

	outstruct := new(struct {
		Maturity  *big.Int
		Precision uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Precision = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// PoolsCalled is a free data retrieval call binding the contract method 0xee3a5c9d.
//
// Solidity: function poolsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceSession) PoolsCalled(arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	return _MarketPlace.Contract.PoolsCalled(&_MarketPlace.CallOpts, arg0)
}

// PoolsCalled is a free data retrieval call binding the contract method 0xee3a5c9d.
//
// Solidity: function poolsCalled(address ) view returns(uint256 maturity, uint8 precision)
func (_MarketPlace *MarketPlaceCallerSession) PoolsCalled(arg0 common.Address) (struct {
	Maturity  *big.Int
	Precision uint8
}, error) {
	return _MarketPlace.Contract.PoolsCalled(&_MarketPlace.CallOpts, arg0)
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

// Paused is a paid mutator transaction binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 p) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) Paused(opts *bind.TransactOpts, p uint8) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "paused", p)
}

// Paused is a paid mutator transaction binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 p) returns(bool)
func (_MarketPlace *MarketPlaceSession) Paused(p uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.TransactOpts, p)
}

// Paused is a paid mutator transaction binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 p) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) Paused(p uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.TransactOpts, p)
}

// PausedReturns is a paid mutator transaction binding the contract method 0x4c5d037a.
//
// Solidity: function pausedReturns(bool p) returns()
func (_MarketPlace *MarketPlaceTransactor) PausedReturns(opts *bind.TransactOpts, p bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "pausedReturns", p)
}

// PausedReturns is a paid mutator transaction binding the contract method 0x4c5d037a.
//
// Solidity: function pausedReturns(bool p) returns()
func (_MarketPlace *MarketPlaceSession) PausedReturns(p bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.PausedReturns(&_MarketPlace.TransactOpts, p)
}

// PausedReturns is a paid mutator transaction binding the contract method 0x4c5d037a.
//
// Solidity: function pausedReturns(bool p) returns()
func (_MarketPlace *MarketPlaceTransactorSession) PausedReturns(p bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.PausedReturns(&_MarketPlace.TransactOpts, p)
}

// Pools is a paid mutator transaction binding the contract method 0x86e69b52.
//
// Solidity: function pools(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) Pools(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "pools", p, u, m)
}

// Pools is a paid mutator transaction binding the contract method 0x86e69b52.
//
// Solidity: function pools(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) Pools(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Pools(&_MarketPlace.TransactOpts, p, u, m)
}

// Pools is a paid mutator transaction binding the contract method 0x86e69b52.
//
// Solidity: function pools(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) Pools(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.Pools(&_MarketPlace.TransactOpts, p, u, m)
}

// PoolsReturns is a paid mutator transaction binding the contract method 0x7763772a.
//
// Solidity: function poolsReturns(address p) returns()
func (_MarketPlace *MarketPlaceTransactor) PoolsReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "poolsReturns", p)
}

// PoolsReturns is a paid mutator transaction binding the contract method 0x7763772a.
//
// Solidity: function poolsReturns(address p) returns()
func (_MarketPlace *MarketPlaceSession) PoolsReturns(p common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.PoolsReturns(&_MarketPlace.TransactOpts, p)
}

// PoolsReturns is a paid mutator transaction binding the contract method 0x7763772a.
//
// Solidity: function poolsReturns(address p) returns()
func (_MarketPlace *MarketPlaceTransactorSession) PoolsReturns(p common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.PoolsReturns(&_MarketPlace.TransactOpts, p)
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
