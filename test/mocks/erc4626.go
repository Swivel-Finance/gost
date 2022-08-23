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

// Erc4626MetaData contains all meta data concerning the Erc4626 contract.
var Erc4626MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"assetReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"convertToAssetsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"withdrawReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106bb806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a326d8a811610066578063a326d8a814610159578063b460af9414610175578063b79c449b146101a5578063cf631143146101d6578063d35147e4146101f25761009e565b806307a2d13a146100a357806320a8b9cd146100d357806326cca786146100ef57806338d52e0f1461010b5780636e553f6514610129575b600080fd5b6100bd60048036038101906100b891906104bd565b610222565b6040516100ca91906104f9565b60405180910390f35b6100ed60048036038101906100e891906104bd565b61022e565b005b610109600480360381019061010491906104bd565b610238565b005b610113610242565b6040516101209190610555565b60405180910390f35b610143600480360381019061013e919061059c565b61026b565b60405161015091906104f9565b60405180910390f35b610173600480360381019061016e91906104bd565b6102bc565b005b61018f600480360381019061018a91906105dc565b6102c6565b60405161019c91906104f9565b60405180910390f35b6101bf60048036038101906101ba919061062f565b6103b3565b6040516101cd92919061065c565b60405180910390f35b6101f060048036038101906101eb919061062f565b6103f7565b005b61020c6004803603810190610207919061062f565b61043a565b60405161021991906104f9565b60405180910390f35b60006001549050919050565b8060048190555050565b8060028190555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600082600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600254905092915050565b8060018190555050565b60006102d0610452565b8481600001818152505082816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050506004549150509392505050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60036020528060005260406000206000915090505481565b604051806040016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b600080fd5b6000819050919050565b61049a81610487565b81146104a557600080fd5b50565b6000813590506104b781610491565b92915050565b6000602082840312156104d3576104d2610482565b5b60006104e1848285016104a8565b91505092915050565b6104f381610487565b82525050565b600060208201905061050e60008301846104ea565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061053f82610514565b9050919050565b61054f81610534565b82525050565b600060208201905061056a6000830184610546565b92915050565b61057981610534565b811461058457600080fd5b50565b60008135905061059681610570565b92915050565b600080604083850312156105b3576105b2610482565b5b60006105c1858286016104a8565b92505060206105d285828601610587565b9150509250929050565b6000806000606084860312156105f5576105f4610482565b5b6000610603868287016104a8565b935050602061061486828701610587565b925050604061062586828701610587565b9150509250925092565b60006020828403121561064557610644610482565b5b600061065384828501610587565b91505092915050565b600060408201905061067160008301856104ea565b61067e6020830184610546565b939250505056fea2646970667358221220cf88f958a28fc82000e5f924a2ec71a09646f9c015241b5ffda3050feb46e6fc64736f6c63430008100033",
}

// Erc4626ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc4626MetaData.ABI instead.
var Erc4626ABI = Erc4626MetaData.ABI

// Erc4626Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc4626MetaData.Bin instead.
var Erc4626Bin = Erc4626MetaData.Bin

// DeployErc4626 deploys a new Ethereum contract, binding an instance of Erc4626 to it.
func DeployErc4626(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc4626, error) {
	parsed, err := Erc4626MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc4626Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc4626{Erc4626Caller: Erc4626Caller{contract: contract}, Erc4626Transactor: Erc4626Transactor{contract: contract}, Erc4626Filterer: Erc4626Filterer{contract: contract}}, nil
}

// Erc4626 is an auto generated Go binding around an Ethereum contract.
type Erc4626 struct {
	Erc4626Caller     // Read-only binding to the contract
	Erc4626Transactor // Write-only binding to the contract
	Erc4626Filterer   // Log filterer for contract events
}

// Erc4626Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc4626Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc4626Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc4626Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc4626Session struct {
	Contract     *Erc4626          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc4626CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc4626CallerSession struct {
	Contract *Erc4626Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Erc4626TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc4626TransactorSession struct {
	Contract     *Erc4626Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Erc4626Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc4626Raw struct {
	Contract *Erc4626 // Generic contract binding to access the raw methods on
}

// Erc4626CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc4626CallerRaw struct {
	Contract *Erc4626Caller // Generic read-only contract binding to access the raw methods on
}

// Erc4626TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc4626TransactorRaw struct {
	Contract *Erc4626Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc4626 creates a new instance of Erc4626, bound to a specific deployed contract.
func NewErc4626(address common.Address, backend bind.ContractBackend) (*Erc4626, error) {
	contract, err := bindErc4626(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc4626{Erc4626Caller: Erc4626Caller{contract: contract}, Erc4626Transactor: Erc4626Transactor{contract: contract}, Erc4626Filterer: Erc4626Filterer{contract: contract}}, nil
}

// NewErc4626Caller creates a new read-only instance of Erc4626, bound to a specific deployed contract.
func NewErc4626Caller(address common.Address, caller bind.ContractCaller) (*Erc4626Caller, error) {
	contract, err := bindErc4626(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4626Caller{contract: contract}, nil
}

// NewErc4626Transactor creates a new write-only instance of Erc4626, bound to a specific deployed contract.
func NewErc4626Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc4626Transactor, error) {
	contract, err := bindErc4626(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4626Transactor{contract: contract}, nil
}

// NewErc4626Filterer creates a new log filterer instance of Erc4626, bound to a specific deployed contract.
func NewErc4626Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc4626Filterer, error) {
	contract, err := bindErc4626(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc4626Filterer{contract: contract}, nil
}

// bindErc4626 binds a generic wrapper to an already deployed contract.
func bindErc4626(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc4626ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4626 *Erc4626Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4626.Contract.Erc4626Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4626 *Erc4626Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4626.Contract.Erc4626Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4626 *Erc4626Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4626.Contract.Erc4626Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4626 *Erc4626CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4626.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4626 *Erc4626TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4626.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4626 *Erc4626TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4626.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626 *Erc4626Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626 *Erc4626Session) Asset() (common.Address, error) {
	return _Erc4626.Contract.Asset(&_Erc4626.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626 *Erc4626CallerSession) Asset() (common.Address, error) {
	return _Erc4626.Contract.Asset(&_Erc4626.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 ) view returns(uint256)
func (_Erc4626 *Erc4626Caller) ConvertToAssets(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc4626.contract.Call(opts, &out, "convertToAssets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 ) view returns(uint256)
func (_Erc4626 *Erc4626Session) ConvertToAssets(arg0 *big.Int) (*big.Int, error) {
	return _Erc4626.Contract.ConvertToAssets(&_Erc4626.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 ) view returns(uint256)
func (_Erc4626 *Erc4626CallerSession) ConvertToAssets(arg0 *big.Int) (*big.Int, error) {
	return _Erc4626.Contract.ConvertToAssets(&_Erc4626.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Erc4626 *Erc4626Caller) DepositCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc4626.contract.Call(opts, &out, "depositCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Erc4626 *Erc4626Session) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc4626.Contract.DepositCalled(&_Erc4626.CallOpts, arg0)
}

// DepositCalled is a free data retrieval call binding the contract method 0xd35147e4.
//
// Solidity: function depositCalled(address ) view returns(uint256)
func (_Erc4626 *Erc4626CallerSession) DepositCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc4626.Contract.DepositCalled(&_Erc4626.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 assets, address owner)
func (_Erc4626 *Erc4626Caller) WithdrawCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Assets *big.Int
	Owner  common.Address
}, error) {
	var out []interface{}
	err := _Erc4626.contract.Call(opts, &out, "withdrawCalled", arg0)

	outstruct := new(struct {
		Assets *big.Int
		Owner  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Assets = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 assets, address owner)
func (_Erc4626 *Erc4626Session) WithdrawCalled(arg0 common.Address) (struct {
	Assets *big.Int
	Owner  common.Address
}, error) {
	return _Erc4626.Contract.WithdrawCalled(&_Erc4626.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256 assets, address owner)
func (_Erc4626 *Erc4626CallerSession) WithdrawCalled(arg0 common.Address) (struct {
	Assets *big.Int
	Owner  common.Address
}, error) {
	return _Erc4626.Contract.WithdrawCalled(&_Erc4626.CallOpts, arg0)
}

// AssetReturns is a paid mutator transaction binding the contract method 0xcf631143.
//
// Solidity: function assetReturns(address a) returns()
func (_Erc4626 *Erc4626Transactor) AssetReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "assetReturns", a)
}

// AssetReturns is a paid mutator transaction binding the contract method 0xcf631143.
//
// Solidity: function assetReturns(address a) returns()
func (_Erc4626 *Erc4626Session) AssetReturns(a common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.AssetReturns(&_Erc4626.TransactOpts, a)
}

// AssetReturns is a paid mutator transaction binding the contract method 0xcf631143.
//
// Solidity: function assetReturns(address a) returns()
func (_Erc4626 *Erc4626TransactorSession) AssetReturns(a common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.AssetReturns(&_Erc4626.TransactOpts, a)
}

// ConvertToAssetsReturns is a paid mutator transaction binding the contract method 0xa326d8a8.
//
// Solidity: function convertToAssetsReturns(uint256 a) returns()
func (_Erc4626 *Erc4626Transactor) ConvertToAssetsReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "convertToAssetsReturns", a)
}

// ConvertToAssetsReturns is a paid mutator transaction binding the contract method 0xa326d8a8.
//
// Solidity: function convertToAssetsReturns(uint256 a) returns()
func (_Erc4626 *Erc4626Session) ConvertToAssetsReturns(a *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.ConvertToAssetsReturns(&_Erc4626.TransactOpts, a)
}

// ConvertToAssetsReturns is a paid mutator transaction binding the contract method 0xa326d8a8.
//
// Solidity: function convertToAssetsReturns(uint256 a) returns()
func (_Erc4626 *Erc4626TransactorSession) ConvertToAssetsReturns(a *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.ConvertToAssetsReturns(&_Erc4626.TransactOpts, a)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Erc4626 *Erc4626Transactor) Deposit(opts *bind.TransactOpts, a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "deposit", a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Erc4626 *Erc4626Session) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.Deposit(&_Erc4626.TransactOpts, a, r)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 a, address r) returns(uint256)
func (_Erc4626 *Erc4626TransactorSession) Deposit(a *big.Int, r common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.Deposit(&_Erc4626.TransactOpts, a, r)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 s) returns()
func (_Erc4626 *Erc4626Transactor) DepositReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "depositReturns", s)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 s) returns()
func (_Erc4626 *Erc4626Session) DepositReturns(s *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.DepositReturns(&_Erc4626.TransactOpts, s)
}

// DepositReturns is a paid mutator transaction binding the contract method 0x26cca786.
//
// Solidity: function depositReturns(uint256 s) returns()
func (_Erc4626 *Erc4626TransactorSession) DepositReturns(s *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.DepositReturns(&_Erc4626.TransactOpts, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 a, address r, address o) returns(uint256)
func (_Erc4626 *Erc4626Transactor) Withdraw(opts *bind.TransactOpts, a *big.Int, r common.Address, o common.Address) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "withdraw", a, r, o)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 a, address r, address o) returns(uint256)
func (_Erc4626 *Erc4626Session) Withdraw(a *big.Int, r common.Address, o common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.Withdraw(&_Erc4626.TransactOpts, a, r, o)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 a, address r, address o) returns(uint256)
func (_Erc4626 *Erc4626TransactorSession) Withdraw(a *big.Int, r common.Address, o common.Address) (*types.Transaction, error) {
	return _Erc4626.Contract.Withdraw(&_Erc4626.TransactOpts, a, r, o)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 s) returns()
func (_Erc4626 *Erc4626Transactor) WithdrawReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Erc4626.contract.Transact(opts, "withdrawReturns", s)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 s) returns()
func (_Erc4626 *Erc4626Session) WithdrawReturns(s *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.WithdrawReturns(&_Erc4626.TransactOpts, s)
}

// WithdrawReturns is a paid mutator transaction binding the contract method 0x20a8b9cd.
//
// Solidity: function withdrawReturns(uint256 s) returns()
func (_Erc4626 *Erc4626TransactorSession) WithdrawReturns(s *big.Int) (*types.Transaction, error) {
	return _Erc4626.Contract.WithdrawReturns(&_Erc4626.TransactOpts, s)
}
