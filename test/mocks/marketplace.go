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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50610fc7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80637d04cd1511610097578063c4c2572611610066578063c4c25726146102c0578063d557ee85146102f3578063eee6e8141461030f578063f8e51bcb14610342576100f5565b80637d04cd15146102285780638c6b9b4114610244578063bddbfbe414610274578063beaff41e146102a4576100f5565b80633f25be9d116100d35780633f25be9d146101625780633f73df2c146101955780634521d303146101c557806365a963aa146101f8576100f5565b806305e1dc25146100fa5780631cd9be911461012a5780632634de1914610146575b600080fd5b610114600480360381019061010f9190610d65565b610372565b6040516101219190610e6e565b60405180910390f35b610144600480360381019061013f9190610e18565b6103e1565b005b610160600480360381019061015b9190610e18565b6103fe565b005b61017c60048036038101906101779190610d3c565b61041b565b60405161018c9493929190610ebf565b60405180910390f35b6101af60048036038101906101aa9190610d3c565b61048b565b6040516101bc9190610ea4565b60405180910390f35b6101df60048036038101906101da9190610d3c565b6104a3565b6040516101ef9493929190610ebf565b60405180910390f35b610212600480360381019061020d9190610da1565b610513565b60405161021f9190610e89565b60405180910390f35b610242600480360381019061023d9190610e18565b6106a0565b005b61025e60048036038101906102599190610da1565b6106bd565b60405161026b9190610e89565b60405180910390f35b61028e60048036038101906102899190610da1565b61084c565b60405161029b9190610e89565b60405180910390f35b6102be60048036038101906102b99190610e18565b6109db565b005b6102da60048036038101906102d59190610d3c565b6109f7565b6040516102ea9493929190610ebf565b60405180910390f35b61030d60048036038101906103089190610d3c565b610a67565b005b61032960048036038101906103249190610d3c565b610aaa565b6040516103399493929190610ebf565b60405180910390f35b61035c60048036038101906103579190610da1565b610b1a565b6040516103699190610e89565b60405180910390f35b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b80600260016101000a81548160ff02191690831515021790555050565b80600260006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60016020528060005260406000206000915090505481565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b600061051d610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505060028054906101000a900460ff1691505095945050505050565b80600260036101000a81548160ff02191690831515021790555050565b60006106c7610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260019054906101000a900460ff1691505095945050505050565b6000610856610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260039054906101000a900460ff1691505095945050505050565b806002806101000a81548160ff02191690831515021790555050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610b24610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260009054906101000a900460ff1691505095945050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610d0c81610f4c565b92915050565b600081359050610d2181610f63565b92915050565b600081359050610d3681610f7a565b92915050565b600060208284031215610d4e57600080fd5b6000610d5c84828501610cfd565b91505092915050565b60008060408385031215610d7857600080fd5b6000610d8685828601610cfd565b9250506020610d9785828601610d27565b9150509250929050565b600080600080600060a08688031215610db957600080fd5b6000610dc788828901610cfd565b9550506020610dd888828901610d27565b9450506040610de988828901610cfd565b9350506060610dfa88828901610cfd565b9250506080610e0b88828901610d27565b9150509295509295909350565b600060208284031215610e2a57600080fd5b6000610e3884828501610d12565b91505092915050565b610e4a81610f04565b82525050565b610e5981610f16565b82525050565b610e6881610f42565b82525050565b6000602082019050610e836000830184610e41565b92915050565b6000602082019050610e9e6000830184610e50565b92915050565b6000602082019050610eb96000830184610e5f565b92915050565b6000608082019050610ed46000830187610e5f565b610ee16020830186610e41565b610eee6040830185610e41565b610efb6060830184610e5f565b95945050505050565b6000610f0f82610f22565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610f5581610f04565b8114610f6057600080fd5b50565b610f6c81610f16565b8114610f7757600080fd5b50565b610f8381610f42565b8114610f8e57600080fd5b5056fea2646970667358221220cbde010ceb7e85eca1269ef8d11c634553361f391b8996c5ad6151370467f04264736f6c63430008000033"

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

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialExitCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialExitCalled", arg0)

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

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialInitiateCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialInitiateCalled", arg0)

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

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pVaultExchangeCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pVaultExchangeCalled", arg0)

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

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pVaultExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pZcTokenExchangeCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pZcTokenExchangeCalled", arg0)

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

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pZcTokenExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
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

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialExit(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExit", u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) CustodialExitReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExitReturns", b)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) CustodialExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExitReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExitReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiate(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiate", u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiateReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiateReturns", b)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) CustodialInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchange(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchange", u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchangeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchangeReturns", b)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) P2pVaultExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchange(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchange", u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchangeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchangeReturns", b)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeReturns(&_MarketPlace.TransactOpts, b)
}
