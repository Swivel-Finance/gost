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
const PoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"p\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalToken\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyPrincipalTokenCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"PrincipalTokenOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyPrincipalTokenPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"p\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"principalToken\",\"outputs\":[{\"internalType\":\"contractIErc5095\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalToken\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellPrincipalTokenCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellPrincipalTokenPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PoolBin is the compiled bytecode used for deploying new contracts.
var PoolBin = "0x608060405234801561001057600080fd5b506111fc806100206000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80639d92bc2f116100f9578063e42b16cb11610097578063e7ba677411610071578063e7ba677414610553578063f038a50f1461056f578063f332ea5e1461058b578063f3ec0d2c146105bb576101c4565b8063e42b16cb146104d4578063e525edce146104f2578063e78a9b0c14610523576101c4565b8063c59964b3116100d3578063c59964b314610439578063c9c917d714610455578063d9b03f8814610485578063dbc162de146104b6576101c4565b80639d92bc2f146103bd578063b03ceb15146103ed578063b45d88011461041d576101c4565b80634b266e28116101665780636fccc087116101405780636fccc08714610337578063757b2b761461035557806376a91253146103715780639b9c9fde146103a1576101c4565b80634b266e28146102e1578063556dd440146102fd5780636f307dc314610319576101c4565b806315c49330116101a257806315c49330146102455780633e23209114610263578063435bf33114610293578063463a861e146102b1576101c4565b806308e1cff5146101c95780630a7e546e146101e5578063134df58b14610215575b600080fd5b6101e360048036038101906101de9190610f84565b6105d7565b005b6101ff60048036038101906101fa9190610f84565b610613565b60405161020c9190610fc0565b60405180910390f35b61022f600480360381019061022a9190611039565b610674565b60405161023c9190610fc0565b60405180910390f35b61024d610713565b60405161025a9190610fc0565b60405180910390f35b61027d60048036038101906102789190611039565b610735565b60405161028a9190610fc0565b60405180910390f35b61029b6107d4565b6040516102a89190610fc0565b60405180910390f35b6102cb60048036038101906102c69190611079565b6107f6565b6040516102d89190610fc0565b60405180910390f35b6102fb60048036038101906102f69190610f84565b610825565b005b61031760048036038101906103129190610f84565b610861565b005b61032161089d565b60405161032e91906110b5565b60405180910390f35b61033f6108c7565b60405161034c9190610fc0565b60405180910390f35b61036f600480360381019061036a9190610f84565b6108e9565b005b61038b60048036038101906103869190610f84565b610925565b6040516103989190610fc0565b60405180910390f35b6103bb60048036038101906103b69190610f84565b610986565b005b6103d760048036038101906103d29190611079565b6109c2565b6040516103e49190610fc0565b60405180910390f35b610407600480360381019061040291906110d0565b6109f1565b6040516104149190610fc0565b60405180910390f35b61043760048036038101906104329190610f84565b610b14565b005b610453600480360381019061044e9190610f84565b610b50565b005b61046f600480360381019061046a9190610f84565b610b8c565b60405161047c9190610fc0565b60405180910390f35b61049f600480360381019061049a9190611079565b610bed565b6040516104ad929190611123565b60405180910390f35b6104be610c49565b6040516104cb91906111ab565b60405180910390f35b6104dc610c72565b6040516104e99190610fc0565b60405180910390f35b61050c60048036038101906105079190611079565b610c94565b60405161051a929190611123565b60405180910390f35b61053d600480360381019061053891906110d0565b610cf0565b60405161054a9190610fc0565b60405180910390f35b61056d60048036038101906105689190611079565b610e13565b005b61058960048036038101906105849190610f84565b610e57565b005b6105a560048036038101906105a09190610f84565b610e93565b6040516105b29190610fc0565b60405180910390f35b6105d560048036038101906105d09190611079565b610ef4565b005b80600260106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600b60006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600560009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600081600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600360009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b600b60009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600260009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b600b60109054906101000a90046fffffffffffffffffffffffffffffffff1681565b60066020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600460006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600560106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600a60109054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600260006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600a60106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600460109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b80600560006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b60086020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600360109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b80600360106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600b60106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600560109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b60076020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600a60009054906101000a90046fffffffffffffffffffffffffffffffff1681565b60096020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600260109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600460106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081600a60006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600460009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b60006fffffffffffffffffffffffffffffffff82169050919050565b610f6181610f3c565b8114610f6c57600080fd5b50565b600081359050610f7e81610f58565b92915050565b600060208284031215610f9a57610f99610f37565b5b6000610fa884828501610f6f565b91505092915050565b610fba81610f3c565b82525050565b6000602082019050610fd56000830184610fb1565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061100682610fdb565b9050919050565b61101681610ffb565b811461102157600080fd5b50565b6000813590506110338161100d565b92915050565b600080604083850312156110505761104f610f37565b5b600061105e85828601611024565b925050602061106f85828601610f6f565b9150509250929050565b60006020828403121561108f5761108e610f37565b5b600061109d84828501611024565b91505092915050565b6110af81610ffb565b82525050565b60006020820190506110ca60008301846110a6565b92915050565b6000806000606084860312156110e9576110e8610f37565b5b60006110f786828701611024565b935050602061110886828701610f6f565b925050604061111986828701610f6f565b9150509250925092565b60006040820190506111386000830185610fb1565b6111456020830184610fb1565b9392505050565b6000819050919050565b600061117161116c61116784610fdb565b61114c565b610fdb565b9050919050565b600061118382611156565b9050919050565b600061119582611178565b9050919050565b6111a58161118a565b82525050565b60006020820190506111c0600083018461119c565b9291505056fea264697066735822122098645af94df8bb9dda1b9cad004d7053c7e604d4aefcf2a7d8ddc067c8c7d5d764736f6c634300080d0033"

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

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolCaller) BuyPrincipalTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPrincipalTokenCalled", arg0)

	outstruct := new(struct {
		PrincipalTokenOut *big.Int
		Min               *big.Int
	})

	outstruct.PrincipalTokenOut = out[0].(*big.Int)
	outstruct.Min = out[1].(*big.Int)

	return *outstruct, err

}

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolSession) BuyPrincipalTokenCalled(arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	return _Pool.Contract.BuyPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolCallerSession) BuyPrincipalTokenCalled(arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	return _Pool.Contract.BuyPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) BuyPrincipalTokenPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPrincipalTokenPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) BuyPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) BuyPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewCalled(&_Pool.CallOpts)
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

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolCaller) PrincipalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "principalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolSession) PrincipalToken() (common.Address, error) {
	return _Pool.Contract.PrincipalToken(&_Pool.CallOpts)
}

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolCallerSession) PrincipalToken() (common.Address, error) {
	return _Pool.Contract.PrincipalToken(&_Pool.CallOpts)
}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolCaller) SellPrincipalTokenCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPrincipalTokenCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolCallerSession) SellPrincipalTokenCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) SellPrincipalTokenPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPrincipalTokenPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) SellPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewCalled(&_Pool.CallOpts)
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

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolSession) Underlying() (common.Address, error) {
	return _Pool.Contract.Underlying(&_Pool.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolCallerSession) Underlying() (common.Address, error) {
	return _Pool.Contract.Underlying(&_Pool.CallOpts)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) BuyPrincipalToken(opts *bind.TransactOpts, t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalToken", t, p, m)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolSession) BuyPrincipalToken(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalToken(&_Pool.TransactOpts, t, p, m)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPrincipalToken(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalToken(&_Pool.TransactOpts, t, p, m)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactor) BuyPrincipalTokenPreview(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenPreview", o)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolSession) BuyPrincipalTokenPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreview(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPrincipalTokenPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreview(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolTransactor) BuyPrincipalTokenPreviewReturns(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenPreviewReturns", o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolSession) BuyPrincipalTokenPreviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewReturns(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolTransactorSession) BuyPrincipalTokenPreviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewReturns(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolTransactor) BuyPrincipalTokenReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenReturns", b)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolSession) BuyPrincipalTokenReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenReturns(&_Pool.TransactOpts, b)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolTransactorSession) BuyPrincipalTokenReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenReturns(&_Pool.TransactOpts, b)
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

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolTransactor) BuyUnderlyingPreviewReturns(opts *bind.TransactOpts, p *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingPreviewReturns", p)
}

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolSession) BuyUnderlyingPreviewReturns(p *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreviewReturns(&_Pool.TransactOpts, p)
}

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolTransactorSession) BuyUnderlyingPreviewReturns(p *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreviewReturns(&_Pool.TransactOpts, p)
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

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolTransactor) PrincipalTokenReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "principalTokenReturns", p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PrincipalTokenReturns(&_Pool.TransactOpts, p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolTransactorSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PrincipalTokenReturns(&_Pool.TransactOpts, p)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) SellPrincipalToken(opts *bind.TransactOpts, t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalToken", t, m)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolSession) SellPrincipalToken(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalToken(&_Pool.TransactOpts, t, m)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) SellPrincipalToken(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalToken(&_Pool.TransactOpts, t, m)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactor) SellPrincipalTokenPreview(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenPreview", i)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreview(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactorSession) SellPrincipalTokenPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreview(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolTransactor) SellPrincipalTokenPreviewReturns(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenPreviewReturns", i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolSession) SellPrincipalTokenPreviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewReturns(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolTransactorSession) SellPrincipalTokenPreviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewReturns(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellPrincipalTokenReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenReturns", s)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolSession) SellPrincipalTokenReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenReturns(&_Pool.TransactOpts, s)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellPrincipalTokenReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenReturns(&_Pool.TransactOpts, s)
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

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellUnderlyingPreviewReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingPreviewReturns", s)
}

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolSession) SellUnderlyingPreviewReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreviewReturns(&_Pool.TransactOpts, s)
}

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellUnderlyingPreviewReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreviewReturns(&_Pool.TransactOpts, s)
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

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolTransactor) UnderlyingReturns(opts *bind.TransactOpts, u common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "underlyingReturns", u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UnderlyingReturns(&_Pool.TransactOpts, u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolTransactorSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UnderlyingReturns(&_Pool.TransactOpts, u)
}
