// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// PErc20ABI is the input ABI used to generate the binding from.
const PErc20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PErc20Bin is the compiled bytecode used for deploying new contracts.
var PErc20Bin = "0x60806040523480156200001157600080fd5b5060405162000e4538038062000e458339810160408190526200003491620001d6565b8251620000499060049060208601906200007d565b5081516200005f9060059060208501906200007d565b506002805460ff191660ff9290921691909117905550620002aa9050565b8280546200008b9062000257565b90600052602060002090601f016020900481019282620000af5760008555620000fa565b82601f10620000ca57805160ff1916838001178555620000fa565b82800160010185558215620000fa579182015b82811115620000fa578251825591602001919060010190620000dd565b50620001089291506200010c565b5090565b5b808211156200010857600081556001016200010d565b600082601f83011262000134578081fd5b81516001600160401b038082111562000151576200015162000294565b604051601f8301601f19908116603f011681019082821181831017156200017c576200017c62000294565b8160405283815260209250868385880101111562000198578485fd5b8491505b83821015620001bb57858201830151818301840152908201906200019c565b83821115620001cc57848385830101525b9695505050505050565b600080600060608486031215620001eb578283fd5b83516001600160401b038082111562000202578485fd5b620002108783880162000123565b9450602086015191508082111562000226578384fd5b50620002358682870162000123565b925050604084015160ff811681146200024c578182fd5b809150509250925092565b600181811c908216806200026c57607f821691505b602082108114156200028e57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610b8b80620002ba6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d7146101a9578063a9059cbb146101bc578063dd62ed3e146101cf57600080fd5b8063395093511461015857806370a082311461016b57806395d89b41146101a157600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610126578063313ce5671461013957600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d6610215565b6040516100e39190610a32565b60405180910390f35b6100ff6100fa366004610a09565b6102a3565b60405190151581526020016100e3565b61011860035481565b6040519081526020016100e3565b6100ff6101343660046109ce565b6102b9565b6002546101469060ff1681565b60405160ff90911681526020016100e3565b6100ff610166366004610a09565b6103ab565b61011861017936600461097b565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100d66103ef565b6100ff6101b7366004610a09565b6103fc565b6100ff6101ca366004610a09565b6104d5565b6101186101dd36600461099c565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6004805461022290610ad2565b80601f016020809104026020016040519081016040528092919081815260200182805461024e90610ad2565b801561029b5780601f106102705761010080835404028352916020019161029b565b820191906000526020600020905b81548152906001019060200180831161027e57829003601f168201915b505050505081565b60006102b03384846104e2565b50600192915050565b60006102c6848484610696565b73ffffffffffffffffffffffffffffffffffffffff841660009081526001602090815260408083203384529091529020548281101561038c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f6572633230207472616e7366657220616d6f756e74206578636565647320616c60448201527f6c6f77616e63650000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6103a0853361039b8685610abb565b6104e2565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490916102b091859061039b908690610aa3565b6005805461022290610ad2565b33600090815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152812054828110156104bc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f65726332302064656372656173656420616c6c6f77616e63652062656c6f772060448201527f7a65726f000000000000000000000000000000000000000000000000000000006064820152608401610383565b6104cb338561039b8685610abb565b5060019392505050565b60006102b0338484610696565b73ffffffffffffffffffffffffffffffffffffffff8316610585576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f657263323020617070726f76652066726f6d20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610383565b73ffffffffffffffffffffffffffffffffffffffff8216610628576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f657263323020617070726f766520746f20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610383565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610738576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6572633230207472616e736665722066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610383565b73ffffffffffffffffffffffffffffffffffffffff82166107db576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f6572633230207472616e7366657220746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610383565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610891576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f6572633230207472616e7366657220616d6f756e74206578636565647320626160448201527f6c616e63650000000000000000000000000000000000000000000000000000006064820152608401610383565b61089b8282610abb565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526020819052604080822093909355908516815290812080548492906108de908490610aa3565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161094491815260200190565b60405180910390a350505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461097657600080fd5b919050565b60006020828403121561098c578081fd5b61099582610952565b9392505050565b600080604083850312156109ae578081fd5b6109b783610952565b91506109c560208401610952565b90509250929050565b6000806000606084860312156109e2578081fd5b6109eb84610952565b92506109f960208501610952565b9150604084013590509250925092565b60008060408385031215610a1b578182fd5b610a2483610952565b946020939093013593505050565b6000602080835283518082850152825b81811015610a5e57858101830151858201604001528201610a42565b81811115610a6f5783604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610ab657610ab6610b26565b500190565b600082821015610acd57610acd610b26565b500390565b600181811c90821680610ae657607f821691505b60208210811415610b20577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220007b8ede98b9e79dc83f97b22e454bb39ade00a4fdd830391f558027a2f6b84864736f6c63430008040033"

// DeployPErc20 deploys a new Ethereum contract, binding an instance of PErc20 to it.
func DeployPErc20(auth *bind.TransactOpts, backend bind.ContractBackend, n string, s string, d uint8) (common.Address, *types.Transaction, *PErc20, error) {
	parsed, err := abi.JSON(strings.NewReader(PErc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PErc20Bin), backend, n, s, d)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PErc20{PErc20Caller: PErc20Caller{contract: contract}, PErc20Transactor: PErc20Transactor{contract: contract}, PErc20Filterer: PErc20Filterer{contract: contract}}, nil
}

// PErc20 is an auto generated Go binding around an Ethereum contract.
type PErc20 struct {
	PErc20Caller     // Read-only binding to the contract
	PErc20Transactor // Write-only binding to the contract
	PErc20Filterer   // Log filterer for contract events
}

// PErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type PErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PErc20Session struct {
	Contract     *PErc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PErc20CallerSession struct {
	Contract *PErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PErc20TransactorSession struct {
	Contract     *PErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type PErc20Raw struct {
	Contract *PErc20 // Generic contract binding to access the raw methods on
}

// PErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PErc20CallerRaw struct {
	Contract *PErc20Caller // Generic read-only contract binding to access the raw methods on
}

// PErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PErc20TransactorRaw struct {
	Contract *PErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPErc20 creates a new instance of PErc20, bound to a specific deployed contract.
func NewPErc20(address common.Address, backend bind.ContractBackend) (*PErc20, error) {
	contract, err := bindPErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PErc20{PErc20Caller: PErc20Caller{contract: contract}, PErc20Transactor: PErc20Transactor{contract: contract}, PErc20Filterer: PErc20Filterer{contract: contract}}, nil
}

// NewPErc20Caller creates a new read-only instance of PErc20, bound to a specific deployed contract.
func NewPErc20Caller(address common.Address, caller bind.ContractCaller) (*PErc20Caller, error) {
	contract, err := bindPErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PErc20Caller{contract: contract}, nil
}

// NewPErc20Transactor creates a new write-only instance of PErc20, bound to a specific deployed contract.
func NewPErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*PErc20Transactor, error) {
	contract, err := bindPErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PErc20Transactor{contract: contract}, nil
}

// NewPErc20Filterer creates a new log filterer instance of PErc20, bound to a specific deployed contract.
func NewPErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*PErc20Filterer, error) {
	contract, err := bindPErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PErc20Filterer{contract: contract}, nil
}

// bindPErc20 binds a generic wrapper to an already deployed contract.
func bindPErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PErc20 *PErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PErc20.Contract.PErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PErc20 *PErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PErc20.Contract.PErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PErc20 *PErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PErc20.Contract.PErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PErc20 *PErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PErc20 *PErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PErc20 *PErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_PErc20 *PErc20Caller) Allowance(opts *bind.CallOpts, o common.Address, s common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "allowance", o, s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_PErc20 *PErc20Session) Allowance(o common.Address, s common.Address) (*big.Int, error) {
	return _PErc20.Contract.Allowance(&_PErc20.CallOpts, o, s)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_PErc20 *PErc20CallerSession) Allowance(o common.Address, s common.Address) (*big.Int, error) {
	return _PErc20.Contract.Allowance(&_PErc20.CallOpts, o, s)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_PErc20 *PErc20Caller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "balanceOf", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_PErc20 *PErc20Session) BalanceOf(a common.Address) (*big.Int, error) {
	return _PErc20.Contract.BalanceOf(&_PErc20.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_PErc20 *PErc20CallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _PErc20.Contract.BalanceOf(&_PErc20.CallOpts, a)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PErc20 *PErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PErc20 *PErc20Session) Decimals() (uint8, error) {
	return _PErc20.Contract.Decimals(&_PErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PErc20 *PErc20CallerSession) Decimals() (uint8, error) {
	return _PErc20.Contract.Decimals(&_PErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PErc20 *PErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PErc20 *PErc20Session) Name() (string, error) {
	return _PErc20.Contract.Name(&_PErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PErc20 *PErc20CallerSession) Name() (string, error) {
	return _PErc20.Contract.Name(&_PErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PErc20 *PErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PErc20 *PErc20Session) Symbol() (string, error) {
	return _PErc20.Contract.Symbol(&_PErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PErc20 *PErc20CallerSession) Symbol() (string, error) {
	return _PErc20.Contract.Symbol(&_PErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PErc20 *PErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PErc20 *PErc20Session) TotalSupply() (*big.Int, error) {
	return _PErc20.Contract.TotalSupply(&_PErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PErc20 *PErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _PErc20.Contract.TotalSupply(&_PErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Transactor) Approve(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.contract.Transact(opts, "approve", s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Session) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.Approve(&_PErc20.TransactOpts, s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20TransactorSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.Approve(&_PErc20.TransactOpts, s, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Transactor) DecreaseAllowance(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.contract.Transact(opts, "decreaseAllowance", s, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Session) DecreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.DecreaseAllowance(&_PErc20.TransactOpts, s, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20TransactorSession) DecreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.DecreaseAllowance(&_PErc20.TransactOpts, s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Transactor) IncreaseAllowance(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.contract.Transact(opts, "increaseAllowance", s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20Session) IncreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.IncreaseAllowance(&_PErc20.TransactOpts, s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_PErc20 *PErc20TransactorSession) IncreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.IncreaseAllowance(&_PErc20.TransactOpts, s, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_PErc20 *PErc20Transactor) Transfer(opts *bind.TransactOpts, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.contract.Transact(opts, "transfer", r, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_PErc20 *PErc20Session) Transfer(r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.Transfer(&_PErc20.TransactOpts, r, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_PErc20 *PErc20TransactorSession) Transfer(r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.Transfer(&_PErc20.TransactOpts, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_PErc20 *PErc20Transactor) TransferFrom(opts *bind.TransactOpts, s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.contract.Transact(opts, "transferFrom", s, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_PErc20 *PErc20Session) TransferFrom(s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.TransferFrom(&_PErc20.TransactOpts, s, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_PErc20 *PErc20TransactorSession) TransferFrom(s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _PErc20.Contract.TransferFrom(&_PErc20.TransactOpts, s, r, a)
}

// PErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PErc20 contract.
type PErc20ApprovalIterator struct {
	Event *PErc20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PErc20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PErc20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PErc20Approval represents a Approval event raised by the PErc20 contract.
type PErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PErc20 *PErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PErc20ApprovalIterator{contract: _PErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PErc20 *PErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PErc20Approval)
				if err := _PErc20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PErc20 *PErc20Filterer) ParseApproval(log types.Log) (*PErc20Approval, error) {
	event := new(PErc20Approval)
	if err := _PErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PErc20 contract.
type PErc20TransferIterator struct {
	Event *PErc20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PErc20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PErc20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PErc20Transfer represents a Transfer event raised by the PErc20 contract.
type PErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PErc20 *PErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PErc20TransferIterator{contract: _PErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PErc20 *PErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PErc20Transfer)
				if err := _PErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PErc20 *PErc20Filterer) ParseTransfer(log types.Log) (*PErc20Transfer, error) {
	event := new(PErc20Transfer)
	if err := _PErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
