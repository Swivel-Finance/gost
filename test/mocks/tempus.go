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

// TempusMetaData contains all meta data concerning the Tempus contract.
var TempusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"depositAndFix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"depositAndFixReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBackingTokenCalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityTimeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumReturnCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAMMCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusPoolCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldBearingToken\",\"outputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"yieldBearingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610734806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80634f255a1f116100715780634f255a1f1461016d5780635d61761f1461019d5780637822212d146101b957806396839828146101d7578063c203bd6b146101f3578063d08dd24b14610211576100b4565b806302d103c0146100b9578063229204d9146100d75780632d81f838146100f55780633ef943bb146101135780634cf1cca7146101315780634e8bfdaa1461014f575b600080fd5b6100c161022d565b6040516100ce9190610407565b60405180910390f35b6100df610240565b6040516100ec919061043b565b60405180910390f35b6100fd610246565b60405161010a91906104d5565b60405180910390f35b61011b610270565b604051610128919061043b565b60405180910390f35b610139610276565b604051610146919061043b565b60405180910390f35b61015761027c565b604051610164919061043b565b60405180910390f35b6101876004803603810190610182919061059d565b610286565b604051610194919061043b565b60405180910390f35b6101b760048036038101906101b2919061062a565b610348565b005b6101c1610352565b6040516101ce9190610678565b60405180910390f35b6101f160048036038101906101ec919061062a565b610378565b005b6101fb610382565b6040516102089190610678565b60405180910390f35b61022b600480360381019061022691906106d1565b6103a8565b005b600660009054906101000a900460ff1681565b60075481565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055481565b60085481565b6000600154905090565b600086600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460058190555083600660006101000a81548160ff021916908315150217905550826007819055508160088190555060005490509695505050505050565b8060018190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060008190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008115159050919050565b610401816103ec565b82525050565b600060208201905061041c60008301846103f8565b92915050565b6000819050919050565b61043581610422565b82525050565b6000602082019050610450600083018461042c565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061049b61049661049184610456565b610476565b610456565b9050919050565b60006104ad82610480565b9050919050565b60006104bf826104a2565b9050919050565b6104cf816104b4565b82525050565b60006020820190506104ea60008301846104c6565b92915050565b600080fd5b600061050082610456565b9050919050565b6000610512826104f5565b9050919050565b61052281610507565b811461052d57600080fd5b50565b60008135905061053f81610519565b92915050565b61054e81610422565b811461055957600080fd5b50565b60008135905061056b81610545565b92915050565b61057a816103ec565b811461058557600080fd5b50565b60008135905061059781610571565b92915050565b60008060008060008060c087890312156105ba576105b96104f0565b5b60006105c889828a01610530565b96505060206105d989828a01610530565b95505060406105ea89828a0161055c565b94505060606105fb89828a01610588565b935050608061060c89828a0161055c565b92505060a061061d89828a0161055c565b9150509295509295509295565b6000602082840312156106405761063f6104f0565b5b600061064e8482850161055c565b91505092915050565b6000610662826104a2565b9050919050565b61067281610657565b82525050565b600060208201905061068d6000830184610669565b92915050565b600061069e826104f5565b9050919050565b6106ae81610693565b81146106b957600080fd5b50565b6000813590506106cb816106a5565b92915050565b6000602082840312156106e7576106e66104f0565b5b60006106f5848285016106bc565b9150509291505056fea2646970667358221220361dc7de1774bc2dbb92911851992a4d87b994371ce86dd6f5dc33604e4865ae64736f6c634300080d0033",
}

// TempusABI is the input ABI used to generate the binding from.
// Deprecated: Use TempusMetaData.ABI instead.
var TempusABI = TempusMetaData.ABI

// TempusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TempusMetaData.Bin instead.
var TempusBin = TempusMetaData.Bin

// DeployTempus deploys a new Ethereum contract, binding an instance of Tempus to it.
func DeployTempus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tempus, error) {
	parsed, err := TempusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TempusBin), backend)
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

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusSession) MaturityTime() (*big.Int, error) {
	return _Tempus.Contract.MaturityTime(&_Tempus.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusCallerSession) MaturityTime() (*big.Int, error) {
	return _Tempus.Contract.MaturityTime(&_Tempus.CallOpts)
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

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusCaller) YieldBearingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "yieldBearingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusSession) YieldBearingToken() (common.Address, error) {
	return _Tempus.Contract.YieldBearingToken(&_Tempus.CallOpts)
}

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusCallerSession) YieldBearingToken() (common.Address, error) {
	return _Tempus.Contract.YieldBearingToken(&_Tempus.CallOpts)
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

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusTransactor) MaturityTimeReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "maturityTimeReturns", m)
}

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusSession) MaturityTimeReturns(m *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.MaturityTimeReturns(&_Tempus.TransactOpts, m)
}

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusTransactorSession) MaturityTimeReturns(m *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.MaturityTimeReturns(&_Tempus.TransactOpts, m)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusTransactor) YieldBearingTokenReturns(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "yieldBearingTokenReturns", t)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusSession) YieldBearingTokenReturns(t common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.YieldBearingTokenReturns(&_Tempus.TransactOpts, t)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusTransactorSession) YieldBearingTokenReturns(t common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.YieldBearingTokenReturns(&_Tempus.TransactOpts, t)
}
