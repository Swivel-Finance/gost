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

// SenseABI is the input ABI used to generate the binding from.
const SenseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sa\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapUnderlyingForPTsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBought\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseBin is the compiled bytecode used for deploying new contracts.
var SenseBin = "0x608060405234801561001057600080fd5b506104e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806319e9be3a1461005c5780632b83cccd1461008d578063534ce00f146100a95780638f1f30f0146100db578063cbf6a44c1461010b575b600080fd5b610076600480360381019061007191906102dc565b610127565b604051610084929190610322565b60405180910390f35b6100a760048036038101906100a29190610377565b61014b565b005b6100c360048036038101906100be91906102dc565b6101bb565b6040516100d2939291906103ca565b60405180910390f35b6100f560048036038101906100f09190610401565b6101e5565b6040516101029190610468565b60405180910390f35b61012560048036038101906101209190610483565b61026f565b005b60026020528060005260406000206000915090508060000154908060010154905082565b604051806040016040528083815260200182815250600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050505050565b60016020528060005260406000206000915090508060000154908060010154908060020154905083565b6000604051806060016040528085815260200184815260200183815250600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506000549050949350505050565b8060008190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102a98261027e565b9050919050565b6102b98161029e565b81146102c457600080fd5b50565b6000813590506102d6816102b0565b92915050565b6000602082840312156102f2576102f1610279565b5b6000610300848285016102c7565b91505092915050565b6000819050919050565b61031c81610309565b82525050565b60006040820190506103376000830185610313565b6103446020830184610313565b9392505050565b61035481610309565b811461035f57600080fd5b50565b6000813590506103718161034b565b92915050565b6000806000606084860312156103905761038f610279565b5b600061039e868287016102c7565b93505060206103af86828701610362565b92505060406103c086828701610362565b9150509250925092565b60006060820190506103df6000830186610313565b6103ec6020830185610313565b6103f96040830184610313565b949350505050565b6000806000806080858703121561041b5761041a610279565b5b6000610429878288016102c7565b945050602061043a87828801610362565b935050604061044b87828801610362565b925050606061045c87828801610362565b91505092959194509250565b600060208201905061047d6000830184610313565b92915050565b60006020828403121561049957610498610279565b5b60006104a784828501610362565b9150509291505056fea264697066735822122070ac2174b6fa1198002b662e338cfbf17d6d6b7ab75f829c719e25e286f0f16664736f6c634300080d0033"

// DeploySense deploys a new Ethereum contract, binding an instance of Sense to it.
func DeploySense(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sense, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SenseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sense{SenseCaller: SenseCaller{contract: contract}, SenseTransactor: SenseTransactor{contract: contract}, SenseFilterer: SenseFilterer{contract: contract}}, nil
}

// Sense is an auto generated Go binding around an Ethereum contract.
type Sense struct {
	SenseCaller     // Read-only binding to the contract
	SenseTransactor // Write-only binding to the contract
	SenseFilterer   // Log filterer for contract events
}

// SenseCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenseSession struct {
	Contract     *Sense            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenseCallerSession struct {
	Contract *SenseCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SenseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenseTransactorSession struct {
	Contract     *SenseTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenseRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenseRaw struct {
	Contract *Sense // Generic contract binding to access the raw methods on
}

// SenseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenseCallerRaw struct {
	Contract *SenseCaller // Generic read-only contract binding to access the raw methods on
}

// SenseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenseTransactorRaw struct {
	Contract *SenseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSense creates a new instance of Sense, bound to a specific deployed contract.
func NewSense(address common.Address, backend bind.ContractBackend) (*Sense, error) {
	contract, err := bindSense(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sense{SenseCaller: SenseCaller{contract: contract}, SenseTransactor: SenseTransactor{contract: contract}, SenseFilterer: SenseFilterer{contract: contract}}, nil
}

// NewSenseCaller creates a new read-only instance of Sense, bound to a specific deployed contract.
func NewSenseCaller(address common.Address, caller bind.ContractCaller) (*SenseCaller, error) {
	contract, err := bindSense(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenseCaller{contract: contract}, nil
}

// NewSenseTransactor creates a new write-only instance of Sense, bound to a specific deployed contract.
func NewSenseTransactor(address common.Address, transactor bind.ContractTransactor) (*SenseTransactor, error) {
	contract, err := bindSense(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenseTransactor{contract: contract}, nil
}

// NewSenseFilterer creates a new log filterer instance of Sense, bound to a specific deployed contract.
func NewSenseFilterer(address common.Address, filterer bind.ContractFilterer) (*SenseFilterer, error) {
	contract, err := bindSense(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenseFilterer{contract: contract}, nil
}

// bindSense binds a generic wrapper to an already deployed contract.
func bindSense(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SenseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sense *SenseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sense.Contract.SenseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sense *SenseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sense.Contract.SenseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sense *SenseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sense.Contract.SenseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sense *SenseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sense.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sense *SenseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sense.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sense *SenseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sense.Contract.contract.Transact(opts, method, params...)
}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(uint256 maturity, uint256 amount)
func (_Sense *SenseCaller) RedeemCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "redeemCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		Amount   *big.Int
	})

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(uint256 maturity, uint256 amount)
func (_Sense *SenseSession) RedeemCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
}, error) {
	return _Sense.Contract.RedeemCalled(&_Sense.CallOpts, arg0)
}

// RedeemCalled is a free data retrieval call binding the contract method 0x19e9be3a.
//
// Solidity: function redeemCalled(address ) view returns(uint256 maturity, uint256 amount)
func (_Sense *SenseCallerSession) RedeemCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
}, error) {
	return _Sense.Contract.RedeemCalled(&_Sense.CallOpts, arg0)
}

// SwapUnderlyingForPTsCalled is a free data retrieval call binding the contract method 0x534ce00f.
//
// Solidity: function swapUnderlyingForPTsCalled(address ) view returns(uint256 maturity, uint256 amount, uint256 minimumBought)
func (_Sense *SenseCaller) SwapUnderlyingForPTsCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity      *big.Int
	Amount        *big.Int
	MinimumBought *big.Int
}, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "swapUnderlyingForPTsCalled", arg0)

	outstruct := new(struct {
		Maturity      *big.Int
		Amount        *big.Int
		MinimumBought *big.Int
	})

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Amount = out[1].(*big.Int)
	outstruct.MinimumBought = out[2].(*big.Int)

	return *outstruct, err

}

// SwapUnderlyingForPTsCalled is a free data retrieval call binding the contract method 0x534ce00f.
//
// Solidity: function swapUnderlyingForPTsCalled(address ) view returns(uint256 maturity, uint256 amount, uint256 minimumBought)
func (_Sense *SenseSession) SwapUnderlyingForPTsCalled(arg0 common.Address) (struct {
	Maturity      *big.Int
	Amount        *big.Int
	MinimumBought *big.Int
}, error) {
	return _Sense.Contract.SwapUnderlyingForPTsCalled(&_Sense.CallOpts, arg0)
}

// SwapUnderlyingForPTsCalled is a free data retrieval call binding the contract method 0x534ce00f.
//
// Solidity: function swapUnderlyingForPTsCalled(address ) view returns(uint256 maturity, uint256 amount, uint256 minimumBought)
func (_Sense *SenseCallerSession) SwapUnderlyingForPTsCalled(arg0 common.Address) (struct {
	Maturity      *big.Int
	Amount        *big.Int
	MinimumBought *big.Int
}, error) {
	return _Sense.Contract.SwapUnderlyingForPTsCalled(&_Sense.CallOpts, arg0)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address a, uint256 m, uint256 amount) returns()
func (_Sense *SenseTransactor) Redeem(opts *bind.TransactOpts, a common.Address, m *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "redeem", a, m, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address a, uint256 m, uint256 amount) returns()
func (_Sense *SenseSession) Redeem(a common.Address, m *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.Redeem(&_Sense.TransactOpts, a, m, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address a, uint256 m, uint256 amount) returns()
func (_Sense *SenseTransactorSession) Redeem(a common.Address, m *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.Redeem(&_Sense.TransactOpts, a, m, amount)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseTransactor) SwapUnderlyingForPTs(opts *bind.TransactOpts, sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "swapUnderlyingForPTs", sa, m, a, mb)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseSession) SwapUnderlyingForPTs(sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTs(&_Sense.TransactOpts, sa, m, a, mb)
}

// SwapUnderlyingForPTs is a paid mutator transaction binding the contract method 0x8f1f30f0.
//
// Solidity: function swapUnderlyingForPTs(address sa, uint256 m, uint256 a, uint256 mb) returns(uint256)
func (_Sense *SenseTransactorSession) SwapUnderlyingForPTs(sa common.Address, m *big.Int, a *big.Int, mb *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTs(&_Sense.TransactOpts, sa, m, a, mb)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseTransactor) SwapUnderlyingForPTsReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "swapUnderlyingForPTsReturns", s)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseSession) SwapUnderlyingForPTsReturns(s *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTsReturns(&_Sense.TransactOpts, s)
}

// SwapUnderlyingForPTsReturns is a paid mutator transaction binding the contract method 0xcbf6a44c.
//
// Solidity: function swapUnderlyingForPTsReturns(uint256 s) returns()
func (_Sense *SenseTransactorSession) SwapUnderlyingForPTsReturns(s *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.SwapUnderlyingForPTsReturns(&_Sense.TransactOpts, s)
}
