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
	ABI: "[{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCurrentReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"supplyRatePerBlockReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610494806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063bd6d894d11610071578063bd6d894d1461018d578063c6bf1552146101ab578063d4e7fdd4146101c7578063d6bcd7aa146101e5578063e7a7b9ce14610203578063e7ba67741461021f576100b4565b806329d9ce3e146100b95780636f307dc3146100d5578063852a12e3146100f35780639ff9f1d414610123578063a0712d681461013f578063ae9d70b01461016f575b600080fd5b6100d360048036038101906100ce9190610352565b61023b565b005b6100dd610245565b6040516100ea91906103c0565b60405180910390f35b61010d60048036038101906101089190610352565b61026f565b60405161011a91906103ea565b60405180910390f35b61013d60048036038101906101389190610352565b610282565b005b61015960048036038101906101549190610352565b61028c565b60405161016691906103ea565b60405180910390f35b61017761029f565b60405161018491906103ea565b60405180910390f35b6101956102a9565b6040516101a291906103ea565b60405180910390f35b6101c560048036038101906101c09190610352565b6102b3565b005b6101cf6102bd565b6040516101dc91906103ea565b60405180910390f35b6101ed6102c3565b6040516101fa91906103ea565b60405180910390f35b61021d60048036038101906102189190610352565b6102c9565b005b61023960048036038101906102349190610431565b6102d3565b005b8060028190555050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000816003819055506002549050919050565b8060058190555050565b6000816001819055506000549050919050565b6000600654905090565b6000600554905090565b8060068190555050565b60015481565b60035481565b8060008190555050565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080fd5b6000819050919050565b61032f8161031c565b811461033a57600080fd5b50565b60008135905061034c81610326565b92915050565b60006020828403121561036857610367610317565b5b60006103768482850161033d565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103aa8261037f565b9050919050565b6103ba8161039f565b82525050565b60006020820190506103d560008301846103b1565b92915050565b6103e48161031c565b82525050565b60006020820190506103ff60008301846103db565b92915050565b61040e8161039f565b811461041957600080fd5b50565b60008135905061042b81610405565b92915050565b60006020828403121561044757610446610317565b5b60006104558482850161041c565b9150509291505056fea2646970667358221220c4e2b710412a907943694ad048d54febe9fa5c4c5b055c5997f9e4d15cfaed3e64736f6c634300080d0033",
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

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) MintCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "mintCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenSession) MintCalled() (*big.Int, error) {
	return _CompoundToken.Contract.MintCalled(&_CompoundToken.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) MintCalled() (*big.Int, error) {
	return _CompoundToken.Contract.MintCalled(&_CompoundToken.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenCaller) RedeemUnderlyingCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundToken.contract.Call(opts, &out, "redeemUnderlyingCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _CompoundToken.Contract.RedeemUnderlyingCalled(&_CompoundToken.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_CompoundToken *CompoundTokenCallerSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _CompoundToken.Contract.RedeemUnderlyingCalled(&_CompoundToken.CallOpts)
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
