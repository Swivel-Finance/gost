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

// APWineRouterABI is the input ABI used to generate the binding from.
const APWineRouterABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAmountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountInReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenInCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenOutCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// APWineRouterBin is the compiled bytecode used for deploying new contracts.
var APWineRouterBin = "0x608060405234801561001057600080fd5b5061042d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639eedc5b01161005b5780639eedc5b014610115578063cb2df34614610133578063cdd4255f14610151578063d179c9c51461016f57610088565b80633ef943bb1461008d5780635275215b146100ab5780635dde4ba6146100c757806362a01c1f146100f7575b600080fd5b61009561018d565b6040516100a29190610269565b60405180910390f35b6100c560048036038101906100c091906102b5565b610193565b005b6100e160048036038101906100dc9190610340565b61019d565b6040516100ee9190610269565b60405180910390f35b6100ff610212565b60405161010c9190610269565b60405180910390f35b61011d610218565b60405161012a9190610269565b60405180910390f35b61013b61021e565b6040516101489190610269565b60405180910390f35b610159610224565b60405161016691906103dc565b60405180910390f35b61017761024a565b6040516101849190610269565b60405180910390f35b60035481565b8060008190555050565b6000866001819055508560028190555084600381905550836004819055508260058190555081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005490509695505050505050565b60045481565b60055481565b60025481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6000819050919050565b61026381610250565b82525050565b600060208201905061027e600083018461025a565b92915050565b600080fd5b61029281610250565b811461029d57600080fd5b50565b6000813590506102af81610289565b92915050565b6000602082840312156102cb576102ca610284565b5b60006102d9848285016102a0565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061030d826102e2565b9050919050565b61031d81610302565b811461032857600080fd5b50565b60008135905061033a81610314565b92915050565b60008060008060008060c0878903121561035d5761035c610284565b5b600061036b89828a016102a0565b965050602061037c89828a016102a0565b955050604061038d89828a016102a0565b945050606061039e89828a016102a0565b93505060806103af89828a016102a0565b92505060a06103c089828a0161032b565b9150509295509295509295565b6103d681610302565b82525050565b60006020820190506103f160008301846103cd565b9291505056fea2646970667358221220a85ab6bcc39e071d909d8b23d3cb9d4c29c98bd764f192fb9b1bdf9d89ff3e7264736f6c634300080d0033"

// DeployAPWineRouter deploys a new Ethereum contract, binding an instance of APWineRouter to it.
func DeployAPWineRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWineRouter, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineRouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APWineRouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWineRouter{APWineRouterCaller: APWineRouterCaller{contract: contract}, APWineRouterTransactor: APWineRouterTransactor{contract: contract}, APWineRouterFilterer: APWineRouterFilterer{contract: contract}}, nil
}

// APWineRouter is an auto generated Go binding around an Ethereum contract.
type APWineRouter struct {
	APWineRouterCaller     // Read-only binding to the contract
	APWineRouterTransactor // Write-only binding to the contract
	APWineRouterFilterer   // Log filterer for contract events
}

// APWineRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineRouterSession struct {
	Contract     *APWineRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineRouterCallerSession struct {
	Contract *APWineRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// APWineRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineRouterTransactorSession struct {
	Contract     *APWineRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// APWineRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineRouterRaw struct {
	Contract *APWineRouter // Generic contract binding to access the raw methods on
}

// APWineRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineRouterCallerRaw struct {
	Contract *APWineRouterCaller // Generic read-only contract binding to access the raw methods on
}

// APWineRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineRouterTransactorRaw struct {
	Contract *APWineRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWineRouter creates a new instance of APWineRouter, bound to a specific deployed contract.
func NewAPWineRouter(address common.Address, backend bind.ContractBackend) (*APWineRouter, error) {
	contract, err := bindAPWineRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWineRouter{APWineRouterCaller: APWineRouterCaller{contract: contract}, APWineRouterTransactor: APWineRouterTransactor{contract: contract}, APWineRouterFilterer: APWineRouterFilterer{contract: contract}}, nil
}

// NewAPWineRouterCaller creates a new read-only instance of APWineRouter, bound to a specific deployed contract.
func NewAPWineRouterCaller(address common.Address, caller bind.ContractCaller) (*APWineRouterCaller, error) {
	contract, err := bindAPWineRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineRouterCaller{contract: contract}, nil
}

// NewAPWineRouterTransactor creates a new write-only instance of APWineRouter, bound to a specific deployed contract.
func NewAPWineRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineRouterTransactor, error) {
	contract, err := bindAPWineRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineRouterTransactor{contract: contract}, nil
}

// NewAPWineRouterFilterer creates a new log filterer instance of APWineRouter, bound to a specific deployed contract.
func NewAPWineRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineRouterFilterer, error) {
	contract, err := bindAPWineRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineRouterFilterer{contract: contract}, nil
}

// bindAPWineRouter binds a generic wrapper to an already deployed contract.
func bindAPWineRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineRouter *APWineRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineRouter.Contract.APWineRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineRouter *APWineRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineRouter.Contract.APWineRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineRouter *APWineRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineRouter.Contract.APWineRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineRouter *APWineRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineRouter *APWineRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineRouter *APWineRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineRouter.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterSession) AmountCalled() (*big.Int, error) {
	return _APWineRouter.Contract.AmountCalled(&_APWineRouter.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCallerSession) AmountCalled() (*big.Int, error) {
	return _APWineRouter.Contract.AmountCalled(&_APWineRouter.CallOpts)
}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCaller) IdCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "idCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWineRouter *APWineRouterSession) IdCalled() (*big.Int, error) {
	return _APWineRouter.Contract.IdCalled(&_APWineRouter.CallOpts)
}

// IdCalled is a free data retrieval call binding the contract method 0xd179c9c5.
//
// Solidity: function idCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCallerSession) IdCalled() (*big.Int, error) {
	return _APWineRouter.Contract.IdCalled(&_APWineRouter.CallOpts)
}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCaller) MinimumAmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "minimumAmountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterSession) MinimumAmountCalled() (*big.Int, error) {
	return _APWineRouter.Contract.MinimumAmountCalled(&_APWineRouter.CallOpts)
}

// MinimumAmountCalled is a free data retrieval call binding the contract method 0x9eedc5b0.
//
// Solidity: function minimumAmountCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCallerSession) MinimumAmountCalled() (*big.Int, error) {
	return _APWineRouter.Contract.MinimumAmountCalled(&_APWineRouter.CallOpts)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWineRouter *APWineRouterCaller) ToCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "toCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWineRouter *APWineRouterSession) ToCalled() (common.Address, error) {
	return _APWineRouter.Contract.ToCalled(&_APWineRouter.CallOpts)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_APWineRouter *APWineRouterCallerSession) ToCalled() (common.Address, error) {
	return _APWineRouter.Contract.ToCalled(&_APWineRouter.CallOpts)
}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCaller) TokenInCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "tokenInCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWineRouter *APWineRouterSession) TokenInCalled() (*big.Int, error) {
	return _APWineRouter.Contract.TokenInCalled(&_APWineRouter.CallOpts)
}

// TokenInCalled is a free data retrieval call binding the contract method 0xcb2df346.
//
// Solidity: function tokenInCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCallerSession) TokenInCalled() (*big.Int, error) {
	return _APWineRouter.Contract.TokenInCalled(&_APWineRouter.CallOpts)
}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCaller) TokenOutCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _APWineRouter.contract.Call(opts, &out, "tokenOutCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWineRouter *APWineRouterSession) TokenOutCalled() (*big.Int, error) {
	return _APWineRouter.Contract.TokenOutCalled(&_APWineRouter.CallOpts)
}

// TokenOutCalled is a free data retrieval call binding the contract method 0x62a01c1f.
//
// Solidity: function tokenOutCalled() view returns(uint256)
func (_APWineRouter *APWineRouterCallerSession) TokenOutCalled() (*big.Int, error) {
	return _APWineRouter.Contract.TokenOutCalled(&_APWineRouter.CallOpts)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWineRouter *APWineRouterTransactor) SwapExactAmountIn(opts *bind.TransactOpts, id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWineRouter.contract.Transact(opts, "swapExactAmountIn", id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWineRouter *APWineRouterSession) SwapExactAmountIn(id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWineRouter.Contract.SwapExactAmountIn(&_APWineRouter.TransactOpts, id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 id, uint256 tokenIn, uint256 amount, uint256 tokenOut, uint256 minimumAmount, address to) returns(uint256)
func (_APWineRouter *APWineRouterTransactorSession) SwapExactAmountIn(id *big.Int, tokenIn *big.Int, amount *big.Int, tokenOut *big.Int, minimumAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWineRouter.Contract.SwapExactAmountIn(&_APWineRouter.TransactOpts, id, tokenIn, amount, tokenOut, minimumAmount, to)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWineRouter *APWineRouterTransactor) SwapExactAmountInReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _APWineRouter.contract.Transact(opts, "swapExactAmountInReturns", s)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWineRouter *APWineRouterSession) SwapExactAmountInReturns(s *big.Int) (*types.Transaction, error) {
	return _APWineRouter.Contract.SwapExactAmountInReturns(&_APWineRouter.TransactOpts, s)
}

// SwapExactAmountInReturns is a paid mutator transaction binding the contract method 0x5275215b.
//
// Solidity: function swapExactAmountInReturns(uint256 s) returns()
func (_APWineRouter *APWineRouterTransactorSession) SwapExactAmountInReturns(s *big.Int) (*types.Transaction, error) {
	return _APWineRouter.Contract.SwapExactAmountInReturns(&_APWineRouter.TransactOpts, s)
}
