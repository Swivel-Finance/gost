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

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"inputs\":[],\"name\":\"PT\",\"outputs\":[{\"internalType\":\"contractIErc5095\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"PTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"p\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyPTCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"ptOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPTPReviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPTPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyPTPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyPTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellPTCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPTPReviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPTPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellPTPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellPTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
var PoolBin = "0x608060405234801561001057600080fd5b5061100e806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80639a7402ac116100c3578063d9b03f881161007c578063d9b03f88146103fe578063e412d6421461042f578063e42b16cb1461044b578063e463854d14610469578063e78a9b0c14610499578063f332ea5e146104c957610158565b80639a7402ac1461032a578063a367e7861461035a578063b249786214610376578063c45be1f214610392578063cf66288d146103c2578063d94073d4146103e057610158565b806370691fbd1161011557806370691fbd146102585780637123a5c414610274578063757b2b761461029057806376a91253146102ac5780637c7870ca146102dc57806387c09bbc146102fa57610158565b806308e1cff51461015d5780633e23209114610179578063463a861e146101a95780635c9f16b6146101d95780636045eae41461020a5780636fccc0871461023a575b600080fd5b61017760048036038101906101729190610dc0565b6104f9565b005b610193600480360381019061018e9190610e4b565b610535565b6040516101a09190610e9a565b60405180910390f35b6101c360048036038101906101be9190610eb5565b6105d4565b6040516101d09190610e9a565b60405180910390f35b6101f360048036038101906101ee9190610eb5565b610603565b604051610201929190610ee2565b60405180910390f35b610224600480360381019061021f9190610f0b565b61065f565b6040516102319190610e9a565b60405180910390f35b610242610782565b60405161024f9190610e9a565b60405180910390f35b610272600480360381019061026d9190610eb5565b6107a4565b005b61028e60048036038101906102899190610dc0565b6107e7565b005b6102aa60048036038101906102a59190610dc0565b610823565b005b6102c660048036038101906102c19190610dc0565b61085f565b6040516102d39190610e9a565b60405180910390f35b6102e46108c0565b6040516102f19190610e9a565b60405180910390f35b610314600480360381019061030f9190610dc0565b6108e2565b6040516103219190610e9a565b60405180910390f35b610344600480360381019061033f9190610eb5565b610943565b6040516103519190610e9a565b60405180910390f35b610374600480360381019061036f9190610dc0565b610972565b005b610390600480360381019061038b9190610dc0565b6109ae565b005b6103ac60048036038101906103a79190610e4b565b6109ea565b6040516103b99190610e9a565b60405180910390f35b6103ca610a89565b6040516103d79190610e9a565b60405180910390f35b6103e8610aab565b6040516103f59190610fbd565b60405180910390f35b61041860048036038101906104139190610eb5565b610ad4565b604051610426929190610ee2565b60405180910390f35b61044960048036038101906104449190610dc0565b610b30565b005b610453610b6c565b6040516104609190610e9a565b60405180910390f35b610483600480360381019061047e9190610dc0565b610b8e565b6040516104909190610e9a565b60405180910390f35b6104b360048036038101906104ae9190610f0b565b610bef565b6040516104c09190610e9a565b60405180910390f35b6104e360048036038101906104de9190610dc0565b610d12565b6040516104f09190610e9a565b60405180910390f35b80600160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600160009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b60056020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b60086020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600260109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b600960109054906101000a90046fffffffffffffffffffffffffffffffff1681565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600260006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600960106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600960109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600a60009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600a60106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600a60109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b60076020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600260106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600460006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600260009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b600a60109054906101000a90046fffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60066020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b80600460106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600960009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600a60006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600a60009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600160109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b600081600960006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600960009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600080fd5b60006fffffffffffffffffffffffffffffffff82169050919050565b610d9d81610d78565b8114610da857600080fd5b50565b600081359050610dba81610d94565b92915050565b600060208284031215610dd657610dd5610d73565b5b6000610de484828501610dab565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e1882610ded565b9050919050565b610e2881610e0d565b8114610e3357600080fd5b50565b600081359050610e4581610e1f565b92915050565b60008060408385031215610e6257610e61610d73565b5b6000610e7085828601610e36565b9250506020610e8185828601610dab565b9150509250929050565b610e9481610d78565b82525050565b6000602082019050610eaf6000830184610e8b565b92915050565b600060208284031215610ecb57610eca610d73565b5b6000610ed984828501610e36565b91505092915050565b6000604082019050610ef76000830185610e8b565b610f046020830184610e8b565b9392505050565b600080600060608486031215610f2457610f23610d73565b5b6000610f3286828701610e36565b9350506020610f4386828701610dab565b9250506040610f5486828701610dab565b9150509250925092565b6000819050919050565b6000610f83610f7e610f7984610ded565b610f5e565b610ded565b9050919050565b6000610f9582610f68565b9050919050565b6000610fa782610f8a565b9050919050565b610fb781610f9c565b82525050565b6000602082019050610fd26000830184610fae565b9291505056fea2646970667358221220f131f5d33d13da26675cc4b60148c3bcda577ea1528323b0c3be79b8dfe00f8964736f6c634300080d0033"

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolCaller) PT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "PT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolSession) PT() (common.Address, error) {
	return _Pool.Contract.PT(&_Pool.CallOpts)
}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolCallerSession) PT() (common.Address, error) {
	return _Pool.Contract.PT(&_Pool.CallOpts)
}

// BuyPTCalled is a free data retrieval call binding the contract method 0x5c9f16b6.
//
// Solidity: function buyPTCalled(address ) view returns(uint128 ptOut, uint128 min)
func (_Pool *PoolCaller) BuyPTCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	PtOut *big.Int
	Min   *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPTCalled", arg0)

	outstruct := new(struct {
		PtOut *big.Int
		Min   *big.Int
	})

	outstruct.PtOut = out[0].(*big.Int)
	outstruct.Min = out[1].(*big.Int)

	return *outstruct, err

}

// BuyPTCalled is a free data retrieval call binding the contract method 0x5c9f16b6.
//
// Solidity: function buyPTCalled(address ) view returns(uint128 ptOut, uint128 min)
func (_Pool *PoolSession) BuyPTCalled(arg0 common.Address) (struct {
	PtOut *big.Int
	Min   *big.Int
}, error) {
	return _Pool.Contract.BuyPTCalled(&_Pool.CallOpts, arg0)
}

// BuyPTCalled is a free data retrieval call binding the contract method 0x5c9f16b6.
//
// Solidity: function buyPTCalled(address ) view returns(uint128 ptOut, uint128 min)
func (_Pool *PoolCallerSession) BuyPTCalled(arg0 common.Address) (struct {
	PtOut *big.Int
	Min   *big.Int
}, error) {
	return _Pool.Contract.BuyPTCalled(&_Pool.CallOpts, arg0)
}

// BuyPTPreviewCalled is a free data retrieval call binding the contract method 0xcf66288d.
//
// Solidity: function buyPTPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) BuyPTPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPTPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyPTPreviewCalled is a free data retrieval call binding the contract method 0xcf66288d.
//
// Solidity: function buyPTPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) BuyPTPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPTPreviewCalled(&_Pool.CallOpts)
}

// BuyPTPreviewCalled is a free data retrieval call binding the contract method 0xcf66288d.
//
// Solidity: function buyPTPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) BuyPTPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPTPreviewCalled(&_Pool.CallOpts)
}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolCaller) BuyUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyUnderlyingCalled", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Min    *big.Int
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.Min = out[1].(*big.Int)

	return *outstruct, err

}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolSession) BuyUnderlyingCalled(arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	return _Pool.Contract.BuyUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolCallerSession) BuyUnderlyingCalled(arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	return _Pool.Contract.BuyUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) BuyUnderlyingPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyUnderlyingPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) BuyUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) BuyUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// SellPTCalled is a free data retrieval call binding the contract method 0x9a7402ac.
//
// Solidity: function sellPTCalled(address ) view returns(uint128)
func (_Pool *PoolCaller) SellPTCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPTCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPTCalled is a free data retrieval call binding the contract method 0x9a7402ac.
//
// Solidity: function sellPTCalled(address ) view returns(uint128)
func (_Pool *PoolSession) SellPTCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPTCalled(&_Pool.CallOpts, arg0)
}

// SellPTCalled is a free data retrieval call binding the contract method 0x9a7402ac.
//
// Solidity: function sellPTCalled(address ) view returns(uint128)
func (_Pool *PoolCallerSession) SellPTCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPTCalled(&_Pool.CallOpts, arg0)
}

// SellPTPreviewCalled is a free data retrieval call binding the contract method 0x7c7870ca.
//
// Solidity: function sellPTPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) SellPTPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPTPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPTPreviewCalled is a free data retrieval call binding the contract method 0x7c7870ca.
//
// Solidity: function sellPTPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) SellPTPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPTPreviewCalled(&_Pool.CallOpts)
}

// SellPTPreviewCalled is a free data retrieval call binding the contract method 0x7c7870ca.
//
// Solidity: function sellPTPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) SellPTPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPTPreviewCalled(&_Pool.CallOpts)
}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolCaller) SellUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellUnderlyingCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolSession) SellUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolCallerSession) SellUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) SellUnderlyingPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellUnderlyingPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) SellUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) SellUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// PTReturns is a paid mutator transaction binding the contract method 0x70691fbd.
//
// Solidity: function PTReturns(address p) returns()
func (_Pool *PoolTransactor) PTReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "PTReturns", p)
}

// PTReturns is a paid mutator transaction binding the contract method 0x70691fbd.
//
// Solidity: function PTReturns(address p) returns()
func (_Pool *PoolSession) PTReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PTReturns(&_Pool.TransactOpts, p)
}

// PTReturns is a paid mutator transaction binding the contract method 0x70691fbd.
//
// Solidity: function PTReturns(address p) returns()
func (_Pool *PoolTransactorSession) PTReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PTReturns(&_Pool.TransactOpts, p)
}

// BuyPT is a paid mutator transaction binding the contract method 0x6045eae4.
//
// Solidity: function buyPT(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) BuyPT(opts *bind.TransactOpts, t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPT", t, p, m)
}

// BuyPT is a paid mutator transaction binding the contract method 0x6045eae4.
//
// Solidity: function buyPT(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolSession) BuyPT(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPT(&_Pool.TransactOpts, t, p, m)
}

// BuyPT is a paid mutator transaction binding the contract method 0x6045eae4.
//
// Solidity: function buyPT(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPT(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPT(&_Pool.TransactOpts, t, p, m)
}

// BuyPTPReviewReturns is a paid mutator transaction binding the contract method 0xe412d642.
//
// Solidity: function buyPTPReviewReturns(uint128 o) returns()
func (_Pool *PoolTransactor) BuyPTPReviewReturns(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPTPReviewReturns", o)
}

// BuyPTPReviewReturns is a paid mutator transaction binding the contract method 0xe412d642.
//
// Solidity: function buyPTPReviewReturns(uint128 o) returns()
func (_Pool *PoolSession) BuyPTPReviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTPReviewReturns(&_Pool.TransactOpts, o)
}

// BuyPTPReviewReturns is a paid mutator transaction binding the contract method 0xe412d642.
//
// Solidity: function buyPTPReviewReturns(uint128 o) returns()
func (_Pool *PoolTransactorSession) BuyPTPReviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTPReviewReturns(&_Pool.TransactOpts, o)
}

// BuyPTPreview is a paid mutator transaction binding the contract method 0x87c09bbc.
//
// Solidity: function buyPTPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactor) BuyPTPreview(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPTPreview", o)
}

// BuyPTPreview is a paid mutator transaction binding the contract method 0x87c09bbc.
//
// Solidity: function buyPTPreview(uint128 o) returns(uint128)
func (_Pool *PoolSession) BuyPTPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTPreview(&_Pool.TransactOpts, o)
}

// BuyPTPreview is a paid mutator transaction binding the contract method 0x87c09bbc.
//
// Solidity: function buyPTPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPTPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTPreview(&_Pool.TransactOpts, o)
}

// BuyPTReturns is a paid mutator transaction binding the contract method 0xa367e786.
//
// Solidity: function buyPTReturns(uint128 b) returns()
func (_Pool *PoolTransactor) BuyPTReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPTReturns", b)
}

// BuyPTReturns is a paid mutator transaction binding the contract method 0xa367e786.
//
// Solidity: function buyPTReturns(uint128 b) returns()
func (_Pool *PoolSession) BuyPTReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTReturns(&_Pool.TransactOpts, b)
}

// BuyPTReturns is a paid mutator transaction binding the contract method 0xa367e786.
//
// Solidity: function buyPTReturns(uint128 b) returns()
func (_Pool *PoolTransactorSession) BuyPTReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPTReturns(&_Pool.TransactOpts, b)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) BuyUnderlying(opts *bind.TransactOpts, t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlying", t, u, m)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolSession) BuyUnderlying(t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlying(&_Pool.TransactOpts, t, u, m)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) BuyUnderlying(t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlying(&_Pool.TransactOpts, t, u, m)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactor) BuyUnderlyingPreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingPreview", u)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolSession) BuyUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreview(&_Pool.TransactOpts, u)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactorSession) BuyUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreview(&_Pool.TransactOpts, u)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolTransactor) BuyUnderlyingReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingReturns", b)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolSession) BuyUnderlyingReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingReturns(&_Pool.TransactOpts, b)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolTransactorSession) BuyUnderlyingReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingReturns(&_Pool.TransactOpts, b)
}

// SellPT is a paid mutator transaction binding the contract method 0xc45be1f2.
//
// Solidity: function sellPT(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) SellPT(opts *bind.TransactOpts, t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPT", t, m)
}

// SellPT is a paid mutator transaction binding the contract method 0xc45be1f2.
//
// Solidity: function sellPT(address t, uint128 m) returns(uint128)
func (_Pool *PoolSession) SellPT(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPT(&_Pool.TransactOpts, t, m)
}

// SellPT is a paid mutator transaction binding the contract method 0xc45be1f2.
//
// Solidity: function sellPT(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) SellPT(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPT(&_Pool.TransactOpts, t, m)
}

// SellPTPReviewReturns is a paid mutator transaction binding the contract method 0xb2497862.
//
// Solidity: function sellPTPReviewReturns(uint128 i) returns()
func (_Pool *PoolTransactor) SellPTPReviewReturns(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPTPReviewReturns", i)
}

// SellPTPReviewReturns is a paid mutator transaction binding the contract method 0xb2497862.
//
// Solidity: function sellPTPReviewReturns(uint128 i) returns()
func (_Pool *PoolSession) SellPTPReviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTPReviewReturns(&_Pool.TransactOpts, i)
}

// SellPTPReviewReturns is a paid mutator transaction binding the contract method 0xb2497862.
//
// Solidity: function sellPTPReviewReturns(uint128 i) returns()
func (_Pool *PoolTransactorSession) SellPTPReviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTPReviewReturns(&_Pool.TransactOpts, i)
}

// SellPTPreview is a paid mutator transaction binding the contract method 0xe463854d.
//
// Solidity: function sellPTPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactor) SellPTPreview(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPTPreview", i)
}

// SellPTPreview is a paid mutator transaction binding the contract method 0xe463854d.
//
// Solidity: function sellPTPreview(uint128 i) returns(uint128)
func (_Pool *PoolSession) SellPTPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTPreview(&_Pool.TransactOpts, i)
}

// SellPTPreview is a paid mutator transaction binding the contract method 0xe463854d.
//
// Solidity: function sellPTPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactorSession) SellPTPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTPreview(&_Pool.TransactOpts, i)
}

// SellPTReturns is a paid mutator transaction binding the contract method 0x7123a5c4.
//
// Solidity: function sellPTReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellPTReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPTReturns", s)
}

// SellPTReturns is a paid mutator transaction binding the contract method 0x7123a5c4.
//
// Solidity: function sellPTReturns(uint128 s) returns()
func (_Pool *PoolSession) SellPTReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTReturns(&_Pool.TransactOpts, s)
}

// SellPTReturns is a paid mutator transaction binding the contract method 0x7123a5c4.
//
// Solidity: function sellPTReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellPTReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPTReturns(&_Pool.TransactOpts, s)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) SellUnderlying(opts *bind.TransactOpts, t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlying", t, m)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolSession) SellUnderlying(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlying(&_Pool.TransactOpts, t, m)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) SellUnderlying(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlying(&_Pool.TransactOpts, t, m)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactor) SellUnderlyingPreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingPreview", u)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolSession) SellUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreview(&_Pool.TransactOpts, u)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactorSession) SellUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreview(&_Pool.TransactOpts, u)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellUnderlyingReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingReturns", s)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolSession) SellUnderlyingReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingReturns(&_Pool.TransactOpts, s)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellUnderlyingReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingReturns(&_Pool.TransactOpts, s)
}
