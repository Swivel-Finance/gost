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

// CErc20MetaData contains all meta data concerning the CErc20 contract.
var CErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCurrentReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"supplyRatePerBlockReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610494806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063bd6d894d11610071578063bd6d894d1461018d578063c6bf1552146101ab578063d4e7fdd4146101c7578063d6bcd7aa146101e5578063e7a7b9ce14610203578063e7ba67741461021f576100b4565b806329d9ce3e146100b95780636f307dc3146100d5578063852a12e3146100f35780639ff9f1d414610123578063a0712d681461013f578063ae9d70b01461016f575b600080fd5b6100d360048036038101906100ce9190610352565b61023b565b005b6100dd610245565b6040516100ea91906103c0565b60405180910390f35b61010d60048036038101906101089190610352565b61026f565b60405161011a91906103ea565b60405180910390f35b61013d60048036038101906101389190610352565b610282565b005b61015960048036038101906101549190610352565b61028c565b60405161016691906103ea565b60405180910390f35b61017761029f565b60405161018491906103ea565b60405180910390f35b6101956102a9565b6040516101a291906103ea565b60405180910390f35b6101c560048036038101906101c09190610352565b6102b3565b005b6101cf6102bd565b6040516101dc91906103ea565b60405180910390f35b6101ed6102c3565b6040516101fa91906103ea565b60405180910390f35b61021d60048036038101906102189190610352565b6102c9565b005b61023960048036038101906102349190610431565b6102d3565b005b8060028190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000816003819055506002549050919050565b8060058190555050565b6000816001819055506000549050919050565b6000600654905090565b6000600554905090565b8060068190555050565b60015481565b60035481565b8060008190555050565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b6000819050919050565b61032f8161031c565b811461033a57600080fd5b50565b60008135905061034c81610326565b92915050565b60006020828403121561036857610367610317565b5b60006103768482850161033d565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103aa8261037f565b9050919050565b6103ba8161039f565b82525050565b60006020820190506103d560008301846103b1565b92915050565b6103e48161031c565b82525050565b60006020820190506103ff60008301846103db565b92915050565b61040e8161039f565b811461041957600080fd5b50565b60008135905061042b81610405565b92915050565b60006020828403121561044757610446610317565b5b60006104558482850161041c565b9150509291505056fea264697066735822122095ba17bc1bac2957efb0b1ae576ab4c30b775d581159c1a79d00ac8bc39e2e9964736f6c634300080d0033",
}

// CErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CErc20MetaData.ABI instead.
var CErc20ABI = CErc20MetaData.ABI

// CErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CErc20MetaData.Bin instead.
var CErc20Bin = CErc20MetaData.Bin

// DeployCErc20 deploys a new Ethereum contract, binding an instance of CErc20 to it.
func DeployCErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CErc20, error) {
	parsed, err := CErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CErc20Bin), backend)
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

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Caller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Session) SupplyRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.SupplyRatePerBlock(&_CErc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20CallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.SupplyRatePerBlock(&_CErc20.CallOpts)
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

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CErc20 *CErc20Transactor) SupplyRatePerBlockReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "supplyRatePerBlockReturns", n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CErc20 *CErc20Session) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.SupplyRatePerBlockReturns(&_CErc20.TransactOpts, n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CErc20 *CErc20TransactorSession) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.SupplyRatePerBlockReturns(&_CErc20.TransactOpts, n)
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
