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

// TempusABI is the input ABI used to generate the binding from.
const TempusABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"depositAndFix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"depositAndFixReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBackingTokenCalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumReturnCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAMMCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusPoolCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TempusBin is the compiled bytecode used for deploying new contracts.
var TempusBin = "0x608060405234801561001057600080fd5b5061056b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80634f255a1f1161005b5780634f255a1f146101055780637822212d146101355780639683982814610153578063c203bd6b1461016f57610088565b806302d103c01461008d578063229204d9146100ab5780633ef943bb146100c95780634cf1cca7146100e7575b600080fd5b61009561018d565b6040516100a291906102e5565b60405180910390f35b6100b36101a0565b6040516100c09190610319565b60405180910390f35b6100d16101a6565b6040516100de9190610319565b60405180910390f35b6100ef6101ac565b6040516100fc9190610319565b60405180910390f35b61011f600480360381019061011a9190610401565b6101b2565b60405161012c9190610319565b60405180910390f35b61013d610274565b60405161014a91906104ed565b60405180910390f35b61016d60048036038101906101689190610508565b61029a565b005b6101776102a4565b60405161018491906104ed565b60405180910390f35b600460009054906101000a900460ff1681565b60055481565b60035481565b60065481565b600086600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460038190555083600460006101000a81548160ff021916908315150217905550826005819055508160068190555060005490509695505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060008190555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008115159050919050565b6102df816102ca565b82525050565b60006020820190506102fa60008301846102d6565b92915050565b6000819050919050565b61031381610300565b82525050565b600060208201905061032e600083018461030a565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061036482610339565b9050919050565b600061037682610359565b9050919050565b6103868161036b565b811461039157600080fd5b50565b6000813590506103a38161037d565b92915050565b6103b281610300565b81146103bd57600080fd5b50565b6000813590506103cf816103a9565b92915050565b6103de816102ca565b81146103e957600080fd5b50565b6000813590506103fb816103d5565b92915050565b60008060008060008060c0878903121561041e5761041d610334565b5b600061042c89828a01610394565b965050602061043d89828a01610394565b955050604061044e89828a016103c0565b945050606061045f89828a016103ec565b935050608061047089828a016103c0565b92505060a061048189828a016103c0565b9150509295509295509295565b6000819050919050565b60006104b36104ae6104a984610339565b61048e565b610339565b9050919050565b60006104c582610498565b9050919050565b60006104d7826104ba565b9050919050565b6104e7816104cc565b82525050565b600060208201905061050260008301846104de565b92915050565b60006020828403121561051e5761051d610334565b5b600061052c848285016103c0565b9150509291505056fea26469706673582212202e0c4f207b060685fbb590f2e10ef0e599242ea0072d945a05ef4a843d56765264736f6c634300080d0033"

// DeployTempus deploys a new Ethereum contract, binding an instance of Tempus to it.
func DeployTempus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tempus, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TempusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tempus{TempusCaller: TempusCaller{contract: contract}, TempusTransactor: TempusTransactor{contract: contract}, TempusFilterer: TempusFilterer{contract: contract}}, nil
}

// Tempus is an auto generated Go binding around an Ethereum contract.
type Tempus struct {
	TempusCaller     // Read-only binding to the contract
	TempusTransactor // Write-only binding to the contract
	TempusFilterer   // Log filterer for contract events
}

// TempusCaller is an auto generated read-only Go binding around an Ethereum contract.
type TempusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TempusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TempusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TempusSession struct {
	Contract     *Tempus           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TempusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TempusCallerSession struct {
	Contract *TempusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TempusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TempusTransactorSession struct {
	Contract     *TempusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TempusRaw is an auto generated low-level Go binding around an Ethereum contract.
type TempusRaw struct {
	Contract *Tempus // Generic contract binding to access the raw methods on
}

// TempusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TempusCallerRaw struct {
	Contract *TempusCaller // Generic read-only contract binding to access the raw methods on
}

// TempusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TempusTransactorRaw struct {
	Contract *TempusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTempus creates a new instance of Tempus, bound to a specific deployed contract.
func NewTempus(address common.Address, backend bind.ContractBackend) (*Tempus, error) {
	contract, err := bindTempus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tempus{TempusCaller: TempusCaller{contract: contract}, TempusTransactor: TempusTransactor{contract: contract}, TempusFilterer: TempusFilterer{contract: contract}}, nil
}

// NewTempusCaller creates a new read-only instance of Tempus, bound to a specific deployed contract.
func NewTempusCaller(address common.Address, caller bind.ContractCaller) (*TempusCaller, error) {
	contract, err := bindTempus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TempusCaller{contract: contract}, nil
}

// NewTempusTransactor creates a new write-only instance of Tempus, bound to a specific deployed contract.
func NewTempusTransactor(address common.Address, transactor bind.ContractTransactor) (*TempusTransactor, error) {
	contract, err := bindTempus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TempusTransactor{contract: contract}, nil
}

// NewTempusFilterer creates a new log filterer instance of Tempus, bound to a specific deployed contract.
func NewTempusFilterer(address common.Address, filterer bind.ContractFilterer) (*TempusFilterer, error) {
	contract, err := bindTempus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TempusFilterer{contract: contract}, nil
}

// bindTempus binds a generic wrapper to an already deployed contract.
func bindTempus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tempus *TempusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tempus.Contract.TempusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tempus *TempusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tempus.Contract.TempusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tempus *TempusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tempus.Contract.TempusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tempus *TempusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tempus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tempus *TempusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tempus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tempus *TempusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tempus.Contract.contract.Transact(opts, method, params...)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusSession) AmountCalled() (*big.Int, error) {
	return _Tempus.Contract.AmountCalled(&_Tempus.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) AmountCalled() (*big.Int, error) {
	return _Tempus.Contract.AmountCalled(&_Tempus.CallOpts)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusCaller) DeadlineCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "deadlineCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusSession) DeadlineCalled() (*big.Int, error) {
	return _Tempus.Contract.DeadlineCalled(&_Tempus.CallOpts)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) DeadlineCalled() (*big.Int, error) {
	return _Tempus.Contract.DeadlineCalled(&_Tempus.CallOpts)
}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusCaller) IsBackingTokenCalled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "isBackingTokenCalled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusSession) IsBackingTokenCalled() (bool, error) {
	return _Tempus.Contract.IsBackingTokenCalled(&_Tempus.CallOpts)
}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusCallerSession) IsBackingTokenCalled() (bool, error) {
	return _Tempus.Contract.IsBackingTokenCalled(&_Tempus.CallOpts)
}

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusCaller) MinimumReturnCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "minimumReturnCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusSession) MinimumReturnCalled() (*big.Int, error) {
	return _Tempus.Contract.MinimumReturnCalled(&_Tempus.CallOpts)
}

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) MinimumReturnCalled() (*big.Int, error) {
	return _Tempus.Contract.MinimumReturnCalled(&_Tempus.CallOpts)
}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusCaller) TempusAMMCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "tempusAMMCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusSession) TempusAMMCalled() (common.Address, error) {
	return _Tempus.Contract.TempusAMMCalled(&_Tempus.CallOpts)
}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusCallerSession) TempusAMMCalled() (common.Address, error) {
	return _Tempus.Contract.TempusAMMCalled(&_Tempus.CallOpts)
}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusCaller) TempusPoolCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "tempusPoolCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusSession) TempusPoolCalled() (common.Address, error) {
	return _Tempus.Contract.TempusPoolCalled(&_Tempus.CallOpts)
}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusCallerSession) TempusPoolCalled() (common.Address, error) {
	return _Tempus.Contract.TempusPoolCalled(&_Tempus.CallOpts)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusTransactor) DepositAndFix(opts *bind.TransactOpts, x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "depositAndFix", x, p, a, bt, mr, d)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusSession) DepositAndFix(x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFix(&_Tempus.TransactOpts, x, p, a, bt, mr, d)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusTransactorSession) DepositAndFix(x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFix(&_Tempus.TransactOpts, x, p, a, bt, mr, d)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusTransactor) DepositAndFixReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "depositAndFixReturns", r)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusSession) DepositAndFixReturns(r *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFixReturns(&_Tempus.TransactOpts, r)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusTransactorSession) DepositAndFixReturns(r *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFixReturns(&_Tempus.TransactOpts, r)
}
