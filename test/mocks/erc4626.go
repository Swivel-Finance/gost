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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"depositReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"withdrawReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061055c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806320a8b9cd1461006757806326cca786146100835780636e553f651461009f578063b460af94146100cf578063b79c449b146100ff578063d35147e414610130575b600080fd5b610081600480360381019061007c9190610379565b610160565b005b61009d60048036038101906100989190610379565b61016a565b005b6100b960048036038101906100b49190610404565b610174565b6040516100c69190610453565b60405180910390f35b6100e960048036038101906100e4919061046e565b6101c5565b6040516100f69190610453565b60405180910390f35b610119600480360381019061011491906104c1565b6102b2565b6040516101279291906104fd565b60405180910390f35b61014a600480360381019061014591906104c1565b6102f6565b6040516101579190610453565b60405180910390f35b8060028190555050565b8060008190555050565b600082600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600054905092915050565b60006101cf61030e565b8481600001818152505082816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050506002549150509392505050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b60016020528060005260406000206000915090505481565b604051806040016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b600080fd5b6000819050919050565b61035681610343565b811461036157600080fd5b50565b6000813590506103738161034d565b92915050565b60006020828403121561038f5761038e61033e565b5b600061039d84828501610364565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103d1826103a6565b9050919050565b6103e1816103c6565b81146103ec57600080fd5b50565b6000813590506103fe816103d8565b92915050565b6000806040838503121561041b5761041a61033e565b5b600061042985828601610364565b925050602061043a858286016103ef565b9150509250929050565b61044d81610343565b82525050565b60006020820190506104686000830184610444565b92915050565b6000806000606084860312156104875761048661033e565b5b600061049586828701610364565b93505060206104a6868287016103ef565b92505060406104b7868287016103ef565b9150509250925092565b6000602082840312156104d7576104d661033e565b5b60006104e5848285016103ef565b91505092915050565b6104f7816103c6565b82525050565b60006040820190506105126000830185610444565b61051f60208301846104ee565b939250505056fea264697066735822122005848b69d9e8fb00049e5ece0a05e003e834f9ca11bbda9e435064df99549cf664736f6c634300080d0033",
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
