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

// APWineMetaData contains all meta data concerning the APWine contract.
var APWineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAmountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountInReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenInCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenOutCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061042d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639eedc5b01161005b5780639eedc5b014610115578063cb2df34614610133578063cdd4255f14610151578063d179c9c51461016f57610088565b80633ef943bb1461008d5780635275215b146100ab5780635dde4ba6146100c757806362a01c1f146100f7575b600080fd5b61009561018d565b6040516100a29190610269565b60405180910390f35b6100c560048036038101906100c091906102b5565b610193565b005b6100e160048036038101906100dc9190610340565b61019d565b6040516100ee9190610269565b60405180910390f35b6100ff610212565b60405161010c9190610269565b60405180910390f35b61011d610218565b60405161012a9190610269565b60405180910390f35b61013b61021e565b6040516101489190610269565b60405180910390f35b610159610224565b60405161016691906103dc565b60405180910390f35b61017761024a565b6040516101849190610269565b60405180910390f35b60035481565b8060008190555050565b6000866001819055508560028190555084600381905550836004819055508260058190555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005490509695505050505050565b60045481565b60055481565b60025481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6000819050919050565b61026381610250565b82525050565b600060208201905061027e600083018461025a565b92915050565b600080fd5b61029281610250565b811461029d57600080fd5b50565b6000813590506102af81610289565b92915050565b6000602082840312156102cb576102ca610284565b5b60006102d9848285016102a0565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061030d826102e2565b9050919050565b61031d81610302565b811461032857600080fd5b50565b60008135905061033a81610314565b92915050565b60008060008060008060c0878903121561035d5761035c610284565b5b600061036b89828a016102a0565b965050602061037c89828a016102a0565b955050604061038d89828a016102a0565b945050606061039e89828a016102a0565b93505060806103af89828a016102a0565b92505060a06103c089828a0161032b565b9150509295509295509295565b6103d681610302565b82525050565b60006020820190506103f160008301846103cd565b9291505056fea2646970667358221220b3a6d31326cfd02d134cb348a26f7dd7987ee831c2374d3e63b8b3c26e7953ef64736f6c634300080d0033",
}

// APWineABI is the input ABI used to generate the binding from.
// Deprecated: Use APWineMetaData.ABI instead.
var APWineABI = APWineMetaData.ABI

// APWineBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use APWineMetaData.Bin instead.
var APWineBin = APWineMetaData.Bin

// DeployAPWine deploys a new Ethereum contract, binding an instance of APWine to it.
func DeployAPWine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWine, error) {
	parsed, err := APWineMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(APWineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWine{APWineCaller: APWineCaller{contract: contract}, APWineTransactor: APWineTransactor{contract: contract}, APWineFilterer: APWineFilterer{contract: contract}}, nil
}

// APWine is an auto generated Go binding around an Ethereum contract.
type APWine struct {
	APWineCaller     // Read-only binding to the contract
	APWineTransactor // Write-only binding to the contract
	APWineFilterer   // Log filterer for contract events
}

// APWineCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineSession struct {
	Contract     *APWine           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineCallerSession struct {
	Contract *APWineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// APWineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineTransactorSession struct {
	Contract     *APWineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineRaw struct {
	Contract *APWine // Generic contract binding to access the raw methods on
}

// APWineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineCallerRaw struct {
	Contract *APWineCaller // Generic read-only contract binding to access the raw methods on
}

// APWineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineTransactorRaw struct {
	Contract *APWineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWine creates a new instance of APWine, bound to a specific deployed contract.
func NewAPWine(address common.Address, backend bind.ContractBackend) (*APWine, error) {
	contract, err := bindAPWine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWine{APWineCaller: APWineCaller{contract: contract}, APWineTransactor: APWineTransactor{contract: contract}, APWineFilterer: APWineFilterer{contract: contract}}, nil
}

// NewAPWineCaller creates a new read-only instance of APWine, bound to a specific deployed contract.
func NewAPWineCaller(address common.Address, caller bind.ContractCaller) (*APWineCaller, error) {
	contract, err := bindAPWine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineCaller{contract: contract}, nil
}

// NewAPWineTransactor creates a new write-only instance of APWine, bound to a specific deployed contract.
func NewAPWineTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineTransactor, error) {
	contract, err := bindAPWine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTransactor{contract: contract}, nil
}

// NewAPWineFilterer creates a new log filterer instance of APWine, bound to a specific deployed contract.
func NewAPWineFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineFilterer, error) {
	contract, err := bindAPWine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineFilterer{contract: contract}, nil
}

// bindAPWine binds a generic wrapper to an already deployed contract.
func bindAPWine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWine *APWineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWine.Contract.APWineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWine *APWineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWine.Contract.APWineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWine *APWineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWine.Contract.APWineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWine *APWineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWine *APWineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWine *APWineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWine.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWine *APWineCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWine *APWineSession) AmountCalled() (*big.Int, error) {
	return _APWine.Contract.AmountCalled(&_APWine.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWine *APWineCallerSession) AmountCalled() (*big.Int, error) {
	return _APWine.Contract.AmountCalled(&_APWine.CallOpts)
}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWine *APWineCaller) IdCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "idCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWine *APWineSession) IdCalled() (*big.Int, error) {
	return _APWine.Contract.IdCalled(&_APWine.CallOpts)
}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWine *APWineCallerSession) IdCalled() (*big.Int, error) {
	return _APWine.Contract.IdCalled(&_APWine.CallOpts)
}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWine *APWineCaller) MinimumAmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "minimumAmountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWine *APWineSession) MinimumAmountCalled() (*big.Int, error) {
	return _APWine.Contract.MinimumAmountCalled(&_APWine.CallOpts)
}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWine *APWineCallerSession) MinimumAmountCalled() (*big.Int, error) {
	return _APWine.Contract.MinimumAmountCalled(&_APWine.CallOpts)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWine *APWineCaller) ToCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "toCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWine *APWineSession) ToCalled() (common.Address, error) {
	return _APWine.Contract.ToCalled(&_APWine.CallOpts)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWine *APWineCallerSession) ToCalled() (common.Address, error) {
	return _APWine.Contract.ToCalled(&_APWine.CallOpts)
}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWine *APWineCaller) TokenInCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "tokenInCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWine *APWineSession) TokenInCalled() (*big.Int, error) {
	return _APWine.Contract.TokenInCalled(&_APWine.CallOpts)
}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWine *APWineCallerSession) TokenInCalled() (*big.Int, error) {
	return _APWine.Contract.TokenInCalled(&_APWine.CallOpts)
}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWine *APWineCaller) TokenOutCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "tokenOutCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWine *APWineSession) TokenOutCalled() (*big.Int, error) {
	return _APWine.Contract.TokenOutCalled(&_APWine.CallOpts)
}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWine *APWineCallerSession) TokenOutCalled() (*big.Int, error) {
	return _APWine.Contract.TokenOutCalled(&_APWine.CallOpts)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWine *APWineTransactor) SwapExactAmountIn(opts *bind.TransactOpts, id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.contract.Transact(opts, "swapExactAmountIn", id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWine *APWineSession) SwapExactAmountIn(id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountIn(&_APWine.TransactOpts, id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWine *APWineTransactorSession) SwapExactAmountIn(id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountIn(&_APWine.TransactOpts, id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWine *APWineTransactor) SwapExactAmountInReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _APWine.contract.Transact(opts, "swapExactAmountInReturns", s)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWine *APWineSession) SwapExactAmountInReturns(s *big.Int) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountInReturns(&_APWine.TransactOpts, s)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWine *APWineTransactorSession) SwapExactAmountInReturns(s *big.Int) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountInReturns(&_APWine.TransactOpts, s)
}
