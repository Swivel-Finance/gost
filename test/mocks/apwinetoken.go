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

// APWineTokenABI is the input ABI used to generate the binding from.
const APWineTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getPTAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// APWineTokenBin is the compiled bytecode used for deploying new contracts.
var APWineTokenBin = "0x608060405234801561001057600080fd5b506106fb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a08231146101275780639a8cb8f314610157578063c0c6e5ab14610173578063e541efa21461018f57610088565b80631e5f74a11461008d57806323b872dd146100ab57806327e235e3146100db5780636521b96a1461010b575b600080fd5b6100956101c0565b6040516100a291906104a1565b60405180910390f35b6100c560048036038101906100c09190610523565b6101ea565b6040516100d29190610591565b60405180910390f35b6100f560048036038101906100f091906105ac565b6102e4565b60405161010291906105e8565b60405180910390f35b6101256004803603810190610120919061062f565b6102fc565b005b610141600480360381019061013c91906105ac565b610319565b60405161014e91906105e8565b60405180910390f35b610171600480360381019061016c91906105ac565b610361565b005b61018d6004803603810190610188919061065c565b6103a5565b005b6101a960048036038101906101a491906105ac565b6103ec565b6040516101b792919061069c565b60405180910390f35b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006101f4610430565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600160149054906101000a900460ff169150509392505050565b60006020528060005260406000206000915090505481565b80600160146101000a81548160ff02191690831515021790555050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061048b82610460565b9050919050565b61049b81610480565b82525050565b60006020820190506104b66000830184610492565b92915050565b600080fd5b6104ca81610480565b81146104d557600080fd5b50565b6000813590506104e7816104c1565b92915050565b6000819050919050565b610500816104ed565b811461050b57600080fd5b50565b60008135905061051d816104f7565b92915050565b60008060006060848603121561053c5761053b6104bc565b5b600061054a868287016104d8565b935050602061055b868287016104d8565b925050604061056c8682870161050e565b9150509250925092565b60008115159050919050565b61058b81610576565b82525050565b60006020820190506105a66000830184610582565b92915050565b6000602082840312156105c2576105c16104bc565b5b60006105d0848285016104d8565b91505092915050565b6105e2816104ed565b82525050565b60006020820190506105fd60008301846105d9565b92915050565b61060c81610576565b811461061757600080fd5b50565b60008135905061062981610603565b92915050565b600060208284031215610645576106446104bc565b5b60006106538482850161061a565b91505092915050565b60008060408385031215610673576106726104bc565b5b6000610681858286016104d8565b92505060206106928582860161050e565b9150509250929050565b60006040820190506106b16000830185610492565b6106be60208301846105d9565b939250505056fea2646970667358221220635cf58715a0ff25de5515134c556bec1ddc106d77e384764ce624ed11913e5864736f6c634300080d0033"

// DeployAPWineToken deploys a new Ethereum contract, binding an instance of APWineToken to it.
func DeployAPWineToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWineToken, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APWineTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// APWineToken is an auto generated Go binding around an Ethereum contract.
type APWineToken struct {
	APWineTokenCaller     // Read-only binding to the contract
	APWineTokenTransactor // Write-only binding to the contract
	APWineTokenFilterer   // Log filterer for contract events
}

// APWineTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineTokenSession struct {
	Contract     *APWineToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineTokenCallerSession struct {
	Contract *APWineTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// APWineTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineTokenTransactorSession struct {
	Contract     *APWineTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// APWineTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineTokenRaw struct {
	Contract *APWineToken // Generic contract binding to access the raw methods on
}

// APWineTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineTokenCallerRaw struct {
	Contract *APWineTokenCaller // Generic read-only contract binding to access the raw methods on
}

// APWineTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineTokenTransactorRaw struct {
	Contract *APWineTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWineToken creates a new instance of APWineToken, bound to a specific deployed contract.
func NewAPWineToken(address common.Address, backend bind.ContractBackend) (*APWineToken, error) {
	contract, err := bindAPWineToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// NewAPWineTokenCaller creates a new read-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenCaller(address common.Address, caller bind.ContractCaller) (*APWineTokenCaller, error) {
	contract, err := bindAPWineToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenCaller{contract: contract}, nil
}

// NewAPWineTokenTransactor creates a new write-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineTokenTransactor, error) {
	contract, err := bindAPWineToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenTransactor{contract: contract}, nil
}

// NewAPWineTokenFilterer creates a new log filterer instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineTokenFilterer, error) {
	contract, err := bindAPWineToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineTokenFilterer{contract: contract}, nil
}

// bindAPWineToken binds a generic wrapper to an already deployed contract.
func bindAPWineToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.APWineTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_APWineToken *APWineTokenCaller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "balanceOf", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_APWineToken *APWineTokenSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _APWineToken.Contract.BalanceOf(&_APWineToken.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_APWineToken *APWineTokenCallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _APWineToken.Contract.BalanceOf(&_APWineToken.CallOpts, a)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_APWineToken *APWineTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_APWineToken *APWineTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _APWineToken.Contract.Balances(&_APWineToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_APWineToken *APWineTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _APWineToken.Contract.Balances(&_APWineToken.CallOpts, arg0)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenCaller) GetPTAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "getPTAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenSession) GetPTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetPTAddress(&_APWineToken.CallOpts)
}

// GetPTAddress is a free data retrieval call binding the contract method 0x1e5f74a1.
//
// Solidity: function getPTAddress() view returns(address)
func (_APWineToken *APWineTokenCallerSession) GetPTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetPTAddress(&_APWineToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _APWineToken.Contract.TransferFromCalled(&_APWineToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _APWineToken.Contract.TransferFromCalled(&_APWineToken.CallOpts, arg0)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_APWineToken *APWineTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, a common.Address, b *big.Int) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "balanceOfReturns", a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_APWineToken *APWineTokenSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOfReturns(&_APWineToken.TransactOpts, a, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0xc0c6e5ab.
//
// Solidity: function balanceOfReturns(address a, uint256 b) returns()
func (_APWineToken *APWineTokenTransactorSession) BalanceOfReturns(a common.Address, b *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOfReturns(&_APWineToken.TransactOpts, a, b)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactor) GetPTAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "getPTAddressReturns", a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetPTAddressReturns(&_APWineToken.TransactOpts, a)
}

// GetPTAddressReturns is a paid mutator transaction binding the contract method 0x9a8cb8f3.
//
// Solidity: function getPTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactorSession) GetPTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetPTAddressReturns(&_APWineToken.TransactOpts, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFrom(&_APWineToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFrom(&_APWineToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFromReturns(&_APWineToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFromReturns(&_APWineToken.TransactOpts, b)
}
