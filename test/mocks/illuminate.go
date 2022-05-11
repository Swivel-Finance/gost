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

// IlluminateABI is the input ABI used to generate the binding from.
const IlluminateABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address[8]\",\"name\":\"\",\"type\":\"address[8]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[8]\",\"name\":\"m\",\"type\":\"address[8]\"}],\"name\":\"marketsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"t\",\"type\":\"bool\"}],\"name\":\"tranferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IlluminateBin is the compiled bytecode used for deploying new contracts.
var IlluminateBin = "0x608060405234801561001057600080fd5b5061061f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806317b3bba71461005c5780632db46d281461008c578063811c34d3146100bc578063f9bc41d9146100da578063f9ce9f14146100f6575b600080fd5b610076600480360381019061007191906103a7565b610112565b6040516100839190610492565b60405180910390f35b6100a660048036038101906100a191906104ae565b6101da565b6040516100b391906104ea565b60405180910390f35b6100c46101f2565b6040516100d19190610520565b60405180910390f35b6100f460048036038101906100ef9190610567565b610209565b005b610110600480360381019061010b91906105bb565b610226565b005b61011a61023b565b81600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060006008806020026040519081016040528092919082600880156101cd576020028201915b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610183575b5050505050905092915050565b60096020528060005260406000206000915090505481565b6000600860009054906101000a900460ff16905090565b80600860006101000a81548160ff02191690831515021790555050565b80600090600861023792919061025e565b5050565b604051806101000160405280600890602082028036833780820191505090505090565b82600881019282156102e0579160200282015b828111156102df57823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610271565b5b5090506102ed91906102f1565b5090565b5b8082111561030a5760008160009055506001016102f2565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061033e82610313565b9050919050565b61034e81610333565b811461035957600080fd5b50565b60008135905061036b81610345565b92915050565b6000819050919050565b61038481610371565b811461038f57600080fd5b50565b6000813590506103a18161037b565b92915050565b600080604083850312156103be576103bd61030e565b5b60006103cc8582860161035c565b92505060206103dd85828601610392565b9150509250929050565b600060089050919050565b600081905092915050565b6000819050919050565b61041081610333565b82525050565b60006104228383610407565b60208301905092915050565b6000602082019050919050565b610444816103e7565b61044e81846103f2565b9250610459826103fd565b8060005b8381101561048a5781516104718782610416565b965061047c8361042e565b92505060018101905061045d565b505050505050565b6000610100820190506104a8600083018461043b565b92915050565b6000602082840312156104c4576104c361030e565b5b60006104d28482850161035c565b91505092915050565b6104e481610371565b82525050565b60006020820190506104ff60008301846104db565b92915050565b60008115159050919050565b61051a81610505565b82525050565b60006020820190506105356000830184610511565b92915050565b61054481610505565b811461054f57600080fd5b50565b6000813590506105618161053b565b92915050565b60006020828403121561057d5761057c61030e565b5b600061058b84828501610552565b91505092915050565b600080fd5b6000819050826020600802820111156105b5576105b4610594565b5b92915050565b600061010082840312156105d2576105d161030e565b5b60006105e084828501610599565b9150509291505056fea2646970667358221220cdd6eb0e8b6711693a637c2a129065b620056c712d7040b49c5fc94f60505d3064736f6c634300080d0033"

// DeployIlluminate deploys a new Ethereum contract, binding an instance of Illuminate to it.
func DeployIlluminate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Illuminate, error) {
	parsed, err := abi.JSON(strings.NewReader(IlluminateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IlluminateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// Illuminate is an auto generated Go binding around an Ethereum contract.
type Illuminate struct {
	IlluminateCaller     // Read-only binding to the contract
	IlluminateTransactor // Write-only binding to the contract
	IlluminateFilterer   // Log filterer for contract events
}

// IlluminateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IlluminateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IlluminateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IlluminateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IlluminateSession struct {
	Contract     *Illuminate       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IlluminateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IlluminateCallerSession struct {
	Contract *IlluminateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IlluminateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IlluminateTransactorSession struct {
	Contract     *IlluminateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IlluminateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IlluminateRaw struct {
	Contract *Illuminate // Generic contract binding to access the raw methods on
}

// IlluminateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IlluminateCallerRaw struct {
	Contract *IlluminateCaller // Generic read-only contract binding to access the raw methods on
}

// IlluminateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IlluminateTransactorRaw struct {
	Contract *IlluminateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIlluminate creates a new instance of Illuminate, bound to a specific deployed contract.
func NewIlluminate(address common.Address, backend bind.ContractBackend) (*Illuminate, error) {
	contract, err := bindIlluminate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// NewIlluminateCaller creates a new read-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateCaller(address common.Address, caller bind.ContractCaller) (*IlluminateCaller, error) {
	contract, err := bindIlluminate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateCaller{contract: contract}, nil
}

// NewIlluminateTransactor creates a new write-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateTransactor(address common.Address, transactor bind.ContractTransactor) (*IlluminateTransactor, error) {
	contract, err := bindIlluminate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateTransactor{contract: contract}, nil
}

// NewIlluminateFilterer creates a new log filterer instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateFilterer(address common.Address, filterer bind.ContractFilterer) (*IlluminateFilterer, error) {
	contract, err := bindIlluminate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IlluminateFilterer{contract: contract}, nil
}

// bindIlluminate binds a generic wrapper to an already deployed contract.
func bindIlluminate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IlluminateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.IlluminateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transact(opts, method, params...)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_Illuminate *IlluminateCaller) MarketsCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "marketsCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_Illuminate *IlluminateSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _Illuminate.Contract.MarketsCalled(&_Illuminate.CallOpts, arg0)
}

// MarketsCalled is a free data retrieval call binding the contract method 0x2db46d28.
//
// Solidity: function marketsCalled(address ) view returns(uint256)
func (_Illuminate *IlluminateCallerSession) MarketsCalled(arg0 common.Address) (*big.Int, error) {
	return _Illuminate.Contract.MarketsCalled(&_Illuminate.CallOpts, arg0)
}

// TransferFrom is a free data retrieval call binding the contract method 0x811c34d3.
//
// Solidity: function transferFrom() view returns(bool)
func (_Illuminate *IlluminateCaller) TransferFrom(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "transferFrom")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferFrom is a free data retrieval call binding the contract method 0x811c34d3.
//
// Solidity: function transferFrom() view returns(bool)
func (_Illuminate *IlluminateSession) TransferFrom() (bool, error) {
	return _Illuminate.Contract.TransferFrom(&_Illuminate.CallOpts)
}

// TransferFrom is a free data retrieval call binding the contract method 0x811c34d3.
//
// Solidity: function transferFrom() view returns(bool)
func (_Illuminate *IlluminateCallerSession) TransferFrom() (bool, error) {
	return _Illuminate.Contract.TransferFrom(&_Illuminate.CallOpts)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_Illuminate *IlluminateTransactor) Markets(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Illuminate.contract.Transact(opts, "markets", u, m)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_Illuminate *IlluminateSession) Markets(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.TransactOpts, u, m)
}

// Markets is a paid mutator transaction binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address u, uint256 m) returns(address[8])
func (_Illuminate *IlluminateTransactorSession) Markets(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.TransactOpts, u, m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_Illuminate *IlluminateTransactor) MarketsReturns(opts *bind.TransactOpts, m [8]common.Address) (*types.Transaction, error) {
	return _Illuminate.contract.Transact(opts, "marketsReturns", m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_Illuminate *IlluminateSession) MarketsReturns(m [8]common.Address) (*types.Transaction, error) {
	return _Illuminate.Contract.MarketsReturns(&_Illuminate.TransactOpts, m)
}

// MarketsReturns is a paid mutator transaction binding the contract method 0xf9ce9f14.
//
// Solidity: function marketsReturns(address[8] m) returns()
func (_Illuminate *IlluminateTransactorSession) MarketsReturns(m [8]common.Address) (*types.Transaction, error) {
	return _Illuminate.Contract.MarketsReturns(&_Illuminate.TransactOpts, m)
}

// TranferFromReturns is a paid mutator transaction binding the contract method 0xf9bc41d9.
//
// Solidity: function tranferFromReturns(bool t) returns()
func (_Illuminate *IlluminateTransactor) TranferFromReturns(opts *bind.TransactOpts, t bool) (*types.Transaction, error) {
	return _Illuminate.contract.Transact(opts, "tranferFromReturns", t)
}

// TranferFromReturns is a paid mutator transaction binding the contract method 0xf9bc41d9.
//
// Solidity: function tranferFromReturns(bool t) returns()
func (_Illuminate *IlluminateSession) TranferFromReturns(t bool) (*types.Transaction, error) {
	return _Illuminate.Contract.TranferFromReturns(&_Illuminate.TransactOpts, t)
}

// TranferFromReturns is a paid mutator transaction binding the contract method 0xf9bc41d9.
//
// Solidity: function tranferFromReturns(bool t) returns()
func (_Illuminate *IlluminateTransactorSession) TranferFromReturns(t bool) (*types.Transaction, error) {
	return _Illuminate.Contract.TranferFromReturns(&_Illuminate.TransactOpts, t)
}
