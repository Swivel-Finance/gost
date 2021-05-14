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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"exitFillingExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exitFillingExitCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"exitFillingExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"initiateFillingInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateFillingInitiateCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"initiateFillingInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFromNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFromZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromZcTokenCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50610fc7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636d226daf11610097578063d370dbef11610066578063d370dbef146102d7578063d557ee85146102f3578063e750eedb1461030f578063fb1a0f5814610342576100f5565b80636d226daf1461023f5780637275047b1461026f578063b80ce9311461028b578063c129a1c2146102a7576100f5565b806330e5b1b9116100d357806330e5b1b9146101795780633f73df2c146101ac5780635d705de3146101dc5780636c9ab7831461020c576100f5565b806305e1dc25146100fa57806319178e711461012a5780632878278814610146575b600080fd5b610114600480360381019061010f9190610d65565b610372565b6040516101219190610e6e565b60405180910390f35b610144600480360381019061013f9190610e18565b6103e1565b005b610160600480360381019061015b9190610d3c565b6103fe565b6040516101709493929190610ebf565b60405180910390f35b610193600480360381019061018e9190610d3c565b61046e565b6040516101a39493929190610ebf565b60405180910390f35b6101c660048036038101906101c19190610d3c565b6104de565b6040516101d39190610ea4565b60405180910390f35b6101f660048036038101906101f19190610da1565b6104f6565b6040516102039190610e89565b60405180910390f35b61022660048036038101906102219190610d3c565b610685565b6040516102369493929190610ebf565b60405180910390f35b61025960048036038101906102549190610da1565b6106f5565b6040516102669190610e89565b60405180910390f35b61028960048036038101906102849190610e18565b610884565b005b6102a560048036038101906102a09190610e18565b6108a1565b005b6102c160048036038101906102bc9190610da1565b6108bd565b6040516102ce9190610e89565b60405180910390f35b6102f160048036038101906102ec9190610e18565b610a4a565b005b61030d60048036038101906103089190610d3c565b610a67565b005b61032960048036038101906103249190610d3c565b610aaa565b6040516103399493929190610ebf565b60405180910390f35b61035c60048036038101906103579190610da1565b610b1a565b6040516103699190610e89565b60405180910390f35b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b80600260036101000a81548160ff02191690831515021790555050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60016020528060005260406000206000915090505481565b6000610500610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260009054906101000a900460ff1691505095945050505050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006106ff610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260019054906101000a900460ff1691505095945050505050565b80600260016101000a81548160ff02191690831515021790555050565b806002806101000a81548160ff02191690831515021790555050565b60006108c7610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505060028054906101000a900460ff1691505095945050505050565b80600260006101000a81548160ff02191690831515021790555050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610b24610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260039054906101000a900460ff1691505095945050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d0c81610f4c565b92915050565b600081359050610d2181610f63565b92915050565b600081359050610d3681610f7a565b92915050565b600060208284031215610d4e57600080fd5b6000610d5c84828501610cfd565b91505092915050565b60008060408385031215610d7857600080fd5b6000610d8685828601610cfd565b9250506020610d9785828601610d27565b9150509250929050565b600080600080600060a08688031215610db957600080fd5b6000610dc788828901610cfd565b9550506020610dd888828901610d27565b9450506040610de988828901610cfd565b9350506060610dfa88828901610cfd565b9250506080610e0b88828901610d27565b9150509295509295909350565b600060208284031215610e2a57600080fd5b6000610e3884828501610d12565b91505092915050565b610e4a81610f04565b82525050565b610e5981610f16565b82525050565b610e6881610f42565b82525050565b6000602082019050610e836000830184610e41565b92915050565b6000602082019050610e9e6000830184610e50565b92915050565b6000602082019050610eb96000830184610e5f565b92915050565b6000608082019050610ed46000830187610e5f565b610ee16020830186610e41565b610eee6040830185610e41565b610efb6060830184610e5f565b95945050505050565b6000610f0f82610f22565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610f5581610f04565b8114610f6057600080fd5b50565b610f6c81610f16565b8114610f7757600080fd5b50565b610f8381610f42565b8114610f8e57600080fd5b5056fea26469706673582212207f3e91af4fc85c52007a34fa2c68bf03d5767c80b4eacabac58687a53406b8d264736f6c63430008000033"

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

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) CTokenAddressCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAddressCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// ExitFillingExitCalled is a free data retrieval call binding the contract method 0xe750eedb.
//
// Solidity: function exitFillingExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) ExitFillingExitCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "exitFillingExitCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// ExitFillingExitCalled is a free data retrieval call binding the contract method 0xe750eedb.
//
// Solidity: function exitFillingExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) ExitFillingExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.ExitFillingExitCalled(&_MarketPlace.CallOpts, arg0)
}

// ExitFillingExitCalled is a free data retrieval call binding the contract method 0xe750eedb.
//
// Solidity: function exitFillingExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) ExitFillingExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.ExitFillingExitCalled(&_MarketPlace.CallOpts, arg0)
}

// InitiateFillingInitiateCalled is a free data retrieval call binding the contract method 0x6c9ab783.
//
// Solidity: function initiateFillingInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) InitiateFillingInitiateCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "initiateFillingInitiateCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// InitiateFillingInitiateCalled is a free data retrieval call binding the contract method 0x6c9ab783.
//
// Solidity: function initiateFillingInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) InitiateFillingInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.InitiateFillingInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// InitiateFillingInitiateCalled is a free data retrieval call binding the contract method 0x6c9ab783.
//
// Solidity: function initiateFillingInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) InitiateFillingInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.InitiateFillingInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferFromNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferFromNotionalCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferFromNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferFromNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferFromZcTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferFromZcTokenCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferFromZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferFromZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) CTokenAddress(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddress", u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactor) CTokenAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddressReturns", a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// ExitFillingExit is a paid mutator transaction binding the contract method 0x6d226daf.
//
// Solidity: function exitFillingExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) ExitFillingExit(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "exitFillingExit", u, m, o, t, a)
}

// ExitFillingExit is a paid mutator transaction binding the contract method 0x6d226daf.
//
// Solidity: function exitFillingExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) ExitFillingExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExitFillingExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// ExitFillingExit is a paid mutator transaction binding the contract method 0x6d226daf.
//
// Solidity: function exitFillingExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) ExitFillingExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExitFillingExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// ExitFillingExitReturns is a paid mutator transaction binding the contract method 0x7275047b.
//
// Solidity: function exitFillingExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) ExitFillingExitReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "exitFillingExitReturns", b)
}

// ExitFillingExitReturns is a paid mutator transaction binding the contract method 0x7275047b.
//
// Solidity: function exitFillingExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) ExitFillingExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExitFillingExitReturns(&_MarketPlace.TransactOpts, b)
}

// ExitFillingExitReturns is a paid mutator transaction binding the contract method 0x7275047b.
//
// Solidity: function exitFillingExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) ExitFillingExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExitFillingExitReturns(&_MarketPlace.TransactOpts, b)
}

// InitiateFillingInitiate is a paid mutator transaction binding the contract method 0x5d705de3.
//
// Solidity: function initiateFillingInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) InitiateFillingInitiate(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "initiateFillingInitiate", u, m, o, t, a)
}

// InitiateFillingInitiate is a paid mutator transaction binding the contract method 0x5d705de3.
//
// Solidity: function initiateFillingInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) InitiateFillingInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.InitiateFillingInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// InitiateFillingInitiate is a paid mutator transaction binding the contract method 0x5d705de3.
//
// Solidity: function initiateFillingInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) InitiateFillingInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.InitiateFillingInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// InitiateFillingInitiateReturns is a paid mutator transaction binding the contract method 0xd370dbef.
//
// Solidity: function initiateFillingInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) InitiateFillingInitiateReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "initiateFillingInitiateReturns", b)
}

// InitiateFillingInitiateReturns is a paid mutator transaction binding the contract method 0xd370dbef.
//
// Solidity: function initiateFillingInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) InitiateFillingInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.InitiateFillingInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// InitiateFillingInitiateReturns is a paid mutator transaction binding the contract method 0xd370dbef.
//
// Solidity: function initiateFillingInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) InitiateFillingInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.InitiateFillingInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferFromNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromNotional", u, m, o, t, a)
}

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferFromNotional(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotional(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromNotional(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotional(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferFromNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromNotionalReturns", b)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferFromNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferFromZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromZcToken", u, m, o, t, a)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferFromZcToken(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcToken(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromZcToken(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcToken(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferFromZcTokenReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromZcTokenReturns", b)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferFromZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcTokenReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcTokenReturns(&_MarketPlace.TransactOpts, b)
}
