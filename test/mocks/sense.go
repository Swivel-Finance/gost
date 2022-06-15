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
const SenseABI = "[{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sa\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapUnderlyingForPTsCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumBought\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapUnderlyingForPTsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SenseBin is the compiled bytecode used for deploying new contracts.
var SenseBin = "0x608060405234801561001057600080fd5b5061055a806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063534ce00f1161005b578063534ce00f146100ed5780638f1f30f01461011f578063b4c4a4c81461014f578063cbf6a44c1461016b5761007d565b806319e9be3a14610082578063204f83f9146100b35780632b83cccd146100d1575b600080fd5b61009c60048036038101906100979190610350565b610187565b6040516100aa929190610396565b60405180910390f35b6100bb6101ab565b6040516100c891906103bf565b60405180910390f35b6100eb60048036038101906100e69190610406565b6101b5565b005b61010760048036038101906101029190610350565b610225565b60405161011693929190610459565b60405180910390f35b61013960048036038101906101349190610490565b61024f565b60405161014691906103bf565b60405180910390f35b610169600480360381019061016491906104f7565b6102d9565b005b610185600480360381019061018091906104f7565b6102e3565b005b60036020528060005260406000206000915090508060000154908060010154905082565b6000600154905090565b604051806040016040528083815260200182815250600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050505050565b60026020528060005260406000206000915090508060000154908060010154908060020154905083565b6000604051806060016040528085815260200184815260200183815250600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506000549050949350505050565b8060018190555050565b8060008190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061031d826102f2565b9050919050565b61032d81610312565b811461033857600080fd5b50565b60008135905061034a81610324565b92915050565b600060208284031215610366576103656102ed565b5b60006103748482850161033b565b91505092915050565b6000819050919050565b6103908161037d565b82525050565b60006040820190506103ab6000830185610387565b6103b86020830184610387565b9392505050565b60006020820190506103d46000830184610387565b92915050565b6103e38161037d565b81146103ee57600080fd5b50565b600081359050610400816103da565b92915050565b60008060006060848603121561041f5761041e6102ed565b5b600061042d8682870161033b565b935050602061043e868287016103f1565b925050604061044f868287016103f1565b9150509250925092565b600060608201905061046e6000830186610387565b61047b6020830185610387565b6104886040830184610387565b949350505050565b600080600080608085870312156104aa576104a96102ed565b5b60006104b88782880161033b565b94505060206104c9878288016103f1565b93505060406104da878288016103f1565b92505060606104eb878288016103f1565b91505092959194509250565b60006020828403121561050d5761050c6102ed565b5b600061051b848285016103f1565b9150509291505056fea26469706673582212202d79cd4c3a829f56cb5f9547c8a01fe9b58d033dfcdd661abb710f06d16ce59264736f6c634300080d0033"

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

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_Sense *SenseCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sense.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_Sense *SenseSession) Maturity() (*big.Int, error) {
	return _Sense.Contract.Maturity(&_Sense.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_Sense *SenseCallerSession) Maturity() (*big.Int, error) {
	return _Sense.Contract.Maturity(&_Sense.CallOpts)
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

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Sense *SenseTransactor) MaturityReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Sense.contract.Transact(opts, "maturityReturns", m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Sense *SenseSession) MaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.MaturityReturns(&_Sense.TransactOpts, m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 m) returns()
func (_Sense *SenseTransactorSession) MaturityReturns(m *big.Int) (*types.Transaction, error) {
	return _Sense.Contract.MaturityReturns(&_Sense.TransactOpts, m)
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
