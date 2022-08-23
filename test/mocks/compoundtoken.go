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

// CompoundTokenMetaData contains all meta data concerning the CompoundToken contract.
var CompoundTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCurrentReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"supplyRatePerBlockReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610556806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063ae9d70b011610071578063ae9d70b01461019f578063bd6d894d146101bd578063c6bf1552146101db578063e7a7b9ce146101f7578063e7ba677414610213578063ee4db5701461022f576100b4565b806329d9ce3e146100b957806364f670ee146100d55780636f307dc314610105578063852a12e3146101235780639ff9f1d414610153578063a0712d681461016f575b600080fd5b6100d360048036038101906100ce9190610414565b61025f565b005b6100ef60048036038101906100ea919061049f565b610269565b6040516100fc91906104db565b60405180910390f35b61010d610281565b60405161011a9190610505565b60405180910390f35b61013d60048036038101906101389190610414565b6102ab565b60405161014a91906104db565b60405180910390f35b61016d60048036038101906101689190610414565b6102fb565b005b61018960048036038101906101849190610414565b610305565b60405161019691906104db565b60405180910390f35b6101a7610355565b6040516101b491906104db565b60405180910390f35b6101c561035f565b6040516101d291906104db565b60405180910390f35b6101f560048036038101906101f09190610414565b610369565b005b610211600480360381019061020c9190610414565b610373565b005b61022d6004803603810190610228919061049f565b61037d565b005b6102496004803603810190610244919061049f565b6103c1565b60405161025691906104db565b60405180910390f35b8060028190555050565b60036020528060005260406000206000915090505481565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506002549050919050565b8060058190555050565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000549050919050565b6000600654905090565b6000600554905090565b8060068190555050565b8060008190555050565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60016020528060005260406000206000915090505481565b600080fd5b6000819050919050565b6103f1816103de565b81146103fc57600080fd5b50565b60008135905061040e816103e8565b92915050565b60006020828403121561042a576104296103d9565b5b6000610438848285016103ff565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061046c82610441565b9050919050565b61047c81610461565b811461048757600080fd5b50565b60008135905061049981610473565b92915050565b6000602082840312156104b5576104b46103d9565b5b60006104c38482850161048a565b91505092915050565b6104d5816103de565b82525050565b60006020820190506104f060008301846104cc565b92915050565b6104ff81610461565b82525050565b600060208201905061051a60008301846104f6565b9291505056fea2646970667358221220127224491fdab093ceb493cbce578629768d10f07295506e9fac8675517fb7e764736f6c63430008100033",
}

// CompoundTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundTokenMetaData.ABI instead.
var CompoundTokenABI = CompoundTokenMetaData.ABI

// CompoundTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompoundTokenMetaData.Bin instead.
var CompoundTokenBin = CompoundTokenMetaData.Bin

// DeployCompoundToken deploys a new Ethereum contract, binding an instance of CompoundToken to it.
func DeployCompoundToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundToken, error) {
	parsed, err := CompoundTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompoundTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompoundToken{CompoundTokenCaller: CompoundTokenCaller{contract: contract}, CompoundTokenTransactor: CompoundTokenTransactor{contract: contract}, CompoundTokenFilterer: CompoundTokenFilterer{contract: contract}}, nil
}

// CompoundToken is an auto generated Go binding around an Ethereum contract.
type CompoundToken struct {
	CompoundTokenCaller     // Read-only binding to the contract
	CompoundTokenTransactor // Write-only binding to the contract
	CompoundTokenFilterer   // Log filterer for contract events
}

// CompoundTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundTokenSession struct {
	Contract     *CompoundToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundTokenCallerSession struct {
	Contract *CompoundTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CompoundTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundTokenTransactorSession struct {
	Contract     *CompoundTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CompoundTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundTokenRaw struct {
	Contract *CompoundToken // Generic contract binding to access the raw methods on
}

// CompoundTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundTokenCallerRaw struct {
	Contract *CompoundTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundTokenTransactorRaw struct {
	Contract *CompoundTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundToken creates a new instance of CompoundToken, bound to a specific deployed contract.
func NewCompoundToken(address common.Address, backend bind.ContractBackend) (*CompoundToken, error) {
	contract, err := bindCompoundToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundToken{CompoundTokenCaller: CompoundTokenCaller{contract: contract}, CompoundTokenTransactor: CompoundTokenTransactor{contract: contract}, CompoundTokenFilterer: CompoundTokenFilterer{contract: contract}}, nil
}

// NewCompoundTokenCaller creates a new read-only instance of CompoundToken, bound to a specific deployed contract.
func NewCompoundTokenCaller(address common.Address, caller bind.ContractCaller) (*CompoundTokenCaller, error) {
	contract, err := bindCompoundToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTokenCaller{contract: contract}, nil
}

// NewCompoundTokenTransactor creates a new write-only instance of CompoundToken, bound to a specific deployed contract.
func NewCompoundTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundTokenTransactor, error) {
	contract, err := bindCompoundToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTokenTransactor{contract: contract}, nil
}

// NewCompoundTokenFilterer creates a new log filterer instance of CompoundToken, bound to a specific deployed contract.
func NewCompoundTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundTokenFilterer, error) {
	contract, err := bindCompoundToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundTokenFilterer{contract: contract}, nil
}

// bindCompoundToken binds a generic wrapper to an already deployed contract.
func bindCompoundToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundToken *CompoundTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundToken.Contract.CompoundTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundToken *CompoundTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundToken.Contract.CompoundTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundToken *CompoundTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundToken.Contract.CompoundTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundToken *CompoundTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundToken *CompoundTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundToken *CompoundTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundToken.Contract.contract.Transact(opts, method, params...)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) ExchangeRateCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "exchangeRateCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CompoundToken *CompoundTokenSession) ExchangeRateCurrent() (*big.Int, error) {
	return _CompoundToken.Contract.ExchangeRateCurrent(&_CompoundToken.CallOpts)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) ExchangeRateCurrent() (*big.Int, error) {
	return _CompoundToken.Contract.ExchangeRateCurrent(&_CompoundToken.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) MintCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "mintCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _CompoundToken.Contract.MintCalled(&_CompoundToken.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _CompoundToken.Contract.MintCalled(&_CompoundToken.CallOpts, arg0)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) RedeemUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "redeemUnderlyingCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenSession) RedeemUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _CompoundToken.Contract.RedeemUnderlyingCalled(&_CompoundToken.CallOpts, arg0)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0x64f670ee.
//
// Solidity: function redeemUnderlyingCalled(address ) view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) RedeemUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _CompoundToken.Contract.RedeemUnderlyingCalled(&_CompoundToken.CallOpts, arg0)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CompoundToken *CompoundTokenSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CompoundToken.Contract.SupplyRatePerBlock(&_CompoundToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CompoundToken.Contract.SupplyRatePerBlock(&_CompoundToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CompoundToken *CompoundTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CompoundToken *CompoundTokenSession) Underlying() (common.Address, error) {
	return _CompoundToken.Contract.Underlying(&_CompoundToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CompoundToken *CompoundTokenCallerSession) Underlying() (common.Address, error) {
	return _CompoundToken.Contract.Underlying(&_CompoundToken.CallOpts)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactor) ExchangeRateCurrentReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "exchangeRateCurrentReturns", n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenSession) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.ExchangeRateCurrentReturns(&_CompoundToken.TransactOpts, n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactorSession) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.ExchangeRateCurrentReturns(&_CompoundToken.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenTransactor) Mint(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "mint", n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.Mint(&_CompoundToken.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenTransactorSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.Mint(&_CompoundToken.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactor) MintReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "mintReturns", n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.MintReturns(&_CompoundToken.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactorSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.MintReturns(&_CompoundToken.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "redeemUnderlying", n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.RedeemUnderlying(&_CompoundToken.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_CompoundToken *CompoundTokenTransactorSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.RedeemUnderlying(&_CompoundToken.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactor) RedeemUnderlyingReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "redeemUnderlyingReturns", n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.RedeemUnderlyingReturns(&_CompoundToken.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactorSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.RedeemUnderlyingReturns(&_CompoundToken.TransactOpts, n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactor) SupplyRatePerBlockReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "supplyRatePerBlockReturns", n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenSession) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.SupplyRatePerBlockReturns(&_CompoundToken.TransactOpts, n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_CompoundToken *CompoundTokenTransactorSession) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _CompoundToken.Contract.SupplyRatePerBlockReturns(&_CompoundToken.TransactOpts, n)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CompoundToken *CompoundTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _CompoundToken.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CompoundToken *CompoundTokenSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _CompoundToken.Contract.UnderlyingReturns(&_CompoundToken.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_CompoundToken *CompoundTokenTransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _CompoundToken.Contract.UnderlyingReturns(&_CompoundToken.TransactOpts, a)
}
