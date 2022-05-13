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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapExactAmountInCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountInReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506104e7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80634308b5b31461005c5780635275215b146100905780635dde4ba6146100ac578063b79c449b146100dc578063f3fef3a31461010c575b600080fd5b610076600480360381019061007191906102d7565b610128565b60405161008795949392919061031d565b60405180910390f35b6100aa60048036038101906100a5919061039c565b61015e565b005b6100c660048036038101906100c191906103c9565b610168565b6040516100d39190610456565b60405180910390f35b6100f660048036038101906100f191906102d7565b610214565b6040516101039190610456565b60405180910390f35b61012660048036038101906101219190610471565b61022c565b005b60026020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154905085565b8060008190555050565b60006040518060a0016040528088815260200187815260200185815260200186815260200184815250600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015590505060005490509695505050505050565b60016020528060005260406000206000915090505481565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102a482610279565b9050919050565b6102b481610299565b81146102bf57600080fd5b50565b6000813590506102d1816102ab565b92915050565b6000602082840312156102ed576102ec610274565b5b60006102fb848285016102c2565b91505092915050565b6000819050919050565b61031781610304565b82525050565b600060a082019050610332600083018861030e565b61033f602083018761030e565b61034c604083018661030e565b610359606083018561030e565b610366608083018461030e565b9695505050505050565b61037981610304565b811461038457600080fd5b50565b60008135905061039681610370565b92915050565b6000602082840312156103b2576103b1610274565b5b60006103c084828501610387565b91505092915050565b60008060008060008060c087890312156103e6576103e5610274565b5b60006103f489828a01610387565b965050602061040589828a01610387565b955050604061041689828a01610387565b945050606061042789828a01610387565b935050608061043889828a01610387565b92505060a061044989828a016102c2565b9150509295509295509295565b600060208201905061046b600083018461030e565b92915050565b6000806040838503121561048857610487610274565b5b6000610496858286016102c2565b92505060206104a785828601610387565b915050925092905056fea2646970667358221220997d0fd02bdc81e9f8b51436a884c3b323234464febc4ecf4a8dccb62d70809e64736f6c634300080d0033",
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

// SwapExactAmountInCalled is a free data retrieval call binding the contract method 0x4308b5b3.
//
// Solidity: function swapExactAmountInCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 tokenOut, uint256 amount, uint256 minimumAmount)
func (_APWine *APWineCaller) SwapExactAmountInCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id            *big.Int
	TokenIn       *big.Int
	TokenOut      *big.Int
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "swapExactAmountInCalled", arg0)

	outstruct := new(struct {
		Id            *big.Int
		TokenIn       *big.Int
		TokenOut      *big.Int
		Amount        *big.Int
		MinimumAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenOut = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinimumAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SwapExactAmountInCalled is a free data retrieval call binding the contract method 0x4308b5b3.
//
// Solidity: function swapExactAmountInCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 tokenOut, uint256 amount, uint256 minimumAmount)
func (_APWine *APWineSession) SwapExactAmountInCalled(arg0 common.Address) (struct {
	Id            *big.Int
	TokenIn       *big.Int
	TokenOut      *big.Int
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	return _APWine.Contract.SwapExactAmountInCalled(&_APWine.CallOpts, arg0)
}

// SwapExactAmountInCalled is a free data retrieval call binding the contract method 0x4308b5b3.
//
// Solidity: function swapExactAmountInCalled(address ) view returns(uint256 id, uint256 tokenIn, uint256 tokenOut, uint256 amount, uint256 minimumAmount)
func (_APWine *APWineCallerSession) SwapExactAmountInCalled(arg0 common.Address) (struct {
	Id            *big.Int
	TokenIn       *big.Int
	TokenOut      *big.Int
	Amount        *big.Int
	MinimumAmount *big.Int
}, error) {
	return _APWine.Contract.SwapExactAmountInCalled(&_APWine.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_APWine *APWineCaller) WithdrawCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APWine.contract.Call(opts, &out, "withdrawCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_APWine *APWineSession) WithdrawCalled(arg0 common.Address) (*big.Int, error) {
	return _APWine.Contract.WithdrawCalled(&_APWine.CallOpts, arg0)
}

// WithdrawCalled is a free data retrieval call binding the contract method 0xb79c449b.
//
// Solidity: function withdrawCalled(address ) view returns(uint256)
func (_APWine *APWineCallerSession) WithdrawCalled(arg0 common.Address) (*big.Int, error) {
	return _APWine.Contract.WithdrawCalled(&_APWine.CallOpts, arg0)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 i, uint256 tokenIn, uint256 a, uint256 tokenOut, uint256 m, address to) returns(uint256)
func (_APWine *APWineTransactor) SwapExactAmountIn(opts *bind.TransactOpts, i *big.Int, tokenIn *big.Int, a *big.Int, tokenOut *big.Int, m *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.contract.Transact(opts, "swapExactAmountIn", i, tokenIn, a, tokenOut, m, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 i, uint256 tokenIn, uint256 a, uint256 tokenOut, uint256 m, address to) returns(uint256)
func (_APWine *APWineSession) SwapExactAmountIn(i *big.Int, tokenIn *big.Int, a *big.Int, tokenOut *big.Int, m *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountIn(&_APWine.TransactOpts, i, tokenIn, a, tokenOut, m, to)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x5dde4ba6.
//
// Solidity: function swapExactAmountIn(uint256 i, uint256 tokenIn, uint256 a, uint256 tokenOut, uint256 m, address to) returns(uint256)
func (_APWine *APWineTransactorSession) SwapExactAmountIn(i *big.Int, tokenIn *big.Int, a *big.Int, tokenOut *big.Int, m *big.Int, to common.Address) (*types.Transaction, error) {
	return _APWine.Contract.SwapExactAmountIn(&_APWine.TransactOpts, i, tokenIn, a, tokenOut, m, to)
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

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address o, uint256 a) returns()
func (_APWine *APWineTransactor) Withdraw(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWine.contract.Transact(opts, "withdraw", o, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address o, uint256 a) returns()
func (_APWine *APWineSession) Withdraw(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWine.Contract.Withdraw(&_APWine.TransactOpts, o, a)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address o, uint256 a) returns()
func (_APWine *APWineTransactorSession) Withdraw(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWine.Contract.Withdraw(&_APWine.TransactOpts, o, a)
}
