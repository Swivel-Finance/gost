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
const SenseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sa\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapUnderlyingForPTsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBought\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseBin is the compiled bytecode used for deploying new contracts.
var SenseBin = "0x608060405234801561001057600080fd5b50610373806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063534ce00f146100465780638f1f30f014610078578063cbf6a44c146100a8575b600080fd5b610060600480360381019061005b91906101e5565b6100c4565b60405161006f9392919061022b565b60405180910390f35b610092600480360381019061008d919061028e565b6100ee565b60405161009f91906102f5565b60405180910390f35b6100c260048036038101906100bd9190610310565b610178565b005b60016020528060005260406000206000915090508060000154908060010154908060020154905083565b6000604051806060016040528085815260200184815260200183815250600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506000549050949350505050565b8060008190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101b282610187565b9050919050565b6101c2816101a7565b81146101cd57600080fd5b50565b6000813590506101df816101b9565b92915050565b6000602082840312156101fb576101fa610182565b5b6000610209848285016101d0565b91505092915050565b6000819050919050565b61022581610212565b82525050565b6000606082019050610240600083018661021c565b61024d602083018561021c565b61025a604083018461021c565b949350505050565b61026b81610212565b811461027657600080fd5b50565b60008135905061028881610262565b92915050565b600080600080608085870312156102a8576102a7610182565b5b60006102b6878288016101d0565b94505060206102c787828801610279565b93505060406102d887828801610279565b92505060606102e987828801610279565b91505092959194509250565b600060208201905061030a600083018461021c565b92915050565b60006020828403121561032657610325610182565b5b600061033484828501610279565b9150509291505056fea2646970667358221220f6b4b0b8b0410ab24799f19d655f8fb4d524105786c5d6ac82891f2c2979d88364736f6c634300080d0033"

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
