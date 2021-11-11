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

// CErc20ABI is the input ABI used to generate the binding from.
const CErc20ABI = "[{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCurrentReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CErc20Bin is the compiled bytecode used for deploying new contracts.
var CErc20Bin = "0x608060405234801561001057600080fd5b50610423806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063bd6d894d11610066578063bd6d894d14610159578063d4e7fdd414610177578063d6bcd7aa14610195578063e7a7b9ce146101b3578063e7ba6774146101cf5761009e565b806329d9ce3e146100a35780636f307dc3146100bf578063852a12e3146100dd5780639ff9f1d41461010d578063a0712d6814610129575b600080fd5b6100bd60048036038101906100b89190610306565b6101eb565b005b6100c76101f5565b6040516100d4919061034d565b60405180910390f35b6100f760048036038101906100f29190610306565b61021f565b6040516101049190610368565b60405180910390f35b61012760048036038101906101229190610306565b610232565b005b610143600480360381019061013e9190610306565b61023c565b6040516101509190610368565b60405180910390f35b61016161024f565b60405161016e9190610368565b60405180910390f35b61017f610259565b60405161018c9190610368565b60405180910390f35b61019d61025f565b6040516101aa9190610368565b60405180910390f35b6101cd60048036038101906101c89190610306565b610265565b005b6101e960048036038101906101e491906102dd565b61026f565b005b8060028190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000816003819055506002549050919050565b8060058190555050565b6000816001819055506000549050919050565b6000600554905090565b60015481565b60035481565b8060008190555050565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000813590506102c2816103bf565b92915050565b6000813590506102d7816103d6565b92915050565b6000602082840312156102ef57600080fd5b60006102fd848285016102b3565b91505092915050565b60006020828403121561031857600080fd5b6000610326848285016102c8565b91505092915050565b61033881610383565b82525050565b610347816103b5565b82525050565b6000602082019050610362600083018461032f565b92915050565b600060208201905061037d600083018461033e565b92915050565b600061038e82610395565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6103c881610383565b81146103d357600080fd5b50565b6103df816103b5565b81146103ea57600080fd5b5056fea2646970667358221220376396d5f3818556fa0fb639a8a07060eff04412ad95b660a0594d18285f7b4d64736f6c63430008040033"

// DeployCErc20 deploys a new Ethereum contract, binding an instance of CErc20 to it.
func DeployCErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CErc20, error) {
	parsed, err := abi.JSON(strings.NewReader(CErc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CErc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CErc20{CErc20Caller: CErc20Caller{contract: contract}, CErc20Transactor: CErc20Transactor{contract: contract}, CErc20Filterer: CErc20Filterer{contract: contract}}, nil
}

// CErc20 is an auto generated Go binding around an Ethereum contract.
type CErc20 struct {
	CErc20Caller     // Read-only binding to the contract
	CErc20Transactor // Write-only binding to the contract
	CErc20Filterer   // Log filterer for contract events
}

// CErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CErc20Session struct {
	Contract     *CErc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CErc20CallerSession struct {
	Contract *CErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CErc20TransactorSession struct {
	Contract     *CErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CErc20Raw struct {
	Contract *CErc20 // Generic contract binding to access the raw methods on
}

// CErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CErc20CallerRaw struct {
	Contract *CErc20Caller // Generic read-only contract binding to access the raw methods on
}

// CErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CErc20TransactorRaw struct {
	Contract *CErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCErc20 creates a new instance of CErc20, bound to a specific deployed contract.
func NewCErc20(address common.Address, backend bind.ContractBackend) (*CErc20, error) {
	contract, err := bindCErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CErc20{CErc20Caller: CErc20Caller{contract: contract}, CErc20Transactor: CErc20Transactor{contract: contract}, CErc20Filterer: CErc20Filterer{contract: contract}}, nil
}

// NewCErc20Caller creates a new read-only instance of CErc20, bound to a specific deployed contract.
func NewCErc20Caller(address common.Address, caller bind.ContractCaller) (*CErc20Caller, error) {
	contract, err := bindCErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CErc20Caller{contract: contract}, nil
}

// NewCErc20Transactor creates a new write-only instance of CErc20, bound to a specific deployed contract.
func NewCErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*CErc20Transactor, error) {
	contract, err := bindCErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CErc20Transactor{contract: contract}, nil
}

// NewCErc20Filterer creates a new log filterer instance of CErc20, bound to a specific deployed contract.
func NewCErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*CErc20Filterer, error) {
	contract, err := bindCErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CErc20Filterer{contract: contract}, nil
}

// bindCErc20 binds a generic wrapper to an already deployed contract.
func bindCErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CErc20 *CErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CErc20.Contract.CErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CErc20 *CErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.Contract.CErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CErc20 *CErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CErc20.Contract.CErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CErc20 *CErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CErc20 *CErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CErc20 *CErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CErc20.Contract.contract.Transact(opts, method, params...)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CErc20 *CErc20Caller) ExchangeRateCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "exchangeRateCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CErc20 *CErc20Session) ExchangeRateCurrent() (*big.Int, error) {
	return _CErc20.Contract.ExchangeRateCurrent(&_CErc20.CallOpts)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CErc20 *CErc20CallerSession) ExchangeRateCurrent() (*big.Int, error) {
	return _CErc20.Contract.ExchangeRateCurrent(&_CErc20.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CErc20 *CErc20Caller) MintCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "mintCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CErc20 *CErc20Session) MintCalled() (*big.Int, error) {
	return _CErc20.Contract.MintCalled(&_CErc20.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CErc20 *CErc20CallerSession) MintCalled() (*big.Int, error) {
	return _CErc20.Contract.MintCalled(&_CErc20.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CErc20 *CErc20Caller) RedeemUnderlyingCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "redeemUnderlyingCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CErc20 *CErc20Session) RedeemUnderlyingCalled() (*big.Int, error) {
	return _CErc20.Contract.RedeemUnderlyingCalled(&_CErc20.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CErc20 *CErc20CallerSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _CErc20.Contract.RedeemUnderlyingCalled(&_CErc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20Session) Underlying() (common.Address, error) {
	return _CErc20.Contract.Underlying(&_CErc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20CallerSession) Underlying() (common.Address, error) {
	return _CErc20.Contract.Underlying(&_CErc20.CallOpts)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CErc20 *CErc20Transactor) ExchangeRateCurrentReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "exchangeRateCurrentReturns", n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CErc20 *CErc20Session) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.ExchangeRateCurrentReturns(&_CErc20.TransactOpts, n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CErc20 *CErc20TransactorSession) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.ExchangeRateCurrentReturns(&_CErc20.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CErc20 *CErc20Transactor) Mint(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "mint", n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CErc20 *CErc20Session) Mint(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Mint(&_CErc20.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CErc20 *CErc20TransactorSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Mint(&_CErc20.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CErc20 *CErc20Transactor) MintReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "mintReturns", n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CErc20 *CErc20Session) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.MintReturns(&_CErc20.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CErc20 *CErc20TransactorSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.MintReturns(&_CErc20.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CErc20 *CErc20Transactor) RedeemUnderlying(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "redeemUnderlying", n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CErc20 *CErc20Session) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlying(&_CErc20.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CErc20 *CErc20TransactorSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlying(&_CErc20.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CErc20 *CErc20Transactor) RedeemUnderlyingReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "redeemUnderlyingReturns", n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CErc20 *CErc20Session) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlyingReturns(&_CErc20.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CErc20 *CErc20TransactorSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlyingReturns(&_CErc20.TransactOpts, n)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CErc20 *CErc20Transactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CErc20 *CErc20Session) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.UnderlyingReturns(&_CErc20.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CErc20 *CErc20TransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.UnderlyingReturns(&_CErc20.TransactOpts, a)
}
