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

// YieldABI is the input ABI used to generate the binding from.
const YieldABI = "[{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20\",\"name\":\"i\",\"type\":\"address\"}],\"name\":\"baseReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBase\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellBaseCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellBasePreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellBasePreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBasePreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"r\",\"type\":\"uint128\"}],\"name\":\"sellBaseReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// YieldBin is the compiled bytecode used for deploying new contracts.
var YieldBin = "0x608060405234801561001057600080fd5b50610789806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063439ec67d11610066578063439ec67d1461015b5780635001f3b5146101775780635e5dc52d14610195578063808c58dd146101b1578063bcc1694f146101cd5761009e565b80631266b30c146100a357806313e7bc8c146100c157806319808486146100f15780631cc52d7714610121578063204f83f91461013d575b600080fd5b6100ab6101fd565b6040516100b89190610463565b60405180910390f35b6100db60048036038101906100d691906104af565b61021f565b6040516100e89190610463565b60405180910390f35b61010b6004803603810190610106919061053a565b610280565b6040516101189190610580565b60405180910390f35b61013b600480360381019061013691906105d7565b610298565b005b6101456102bb565b6040516101529190610613565b60405180910390f35b6101756004803603810190610170919061066c565b6102d4565b005b61017f610318565b60405161018c91906106f8565b60405180910390f35b6101af60048036038101906101aa91906104af565b610341565b005b6101cb60048036038101906101c691906104af565b61037d565b005b6101e760048036038101906101e29190610713565b6103b9565b6040516101f49190610463565b60405180910390f35b600360009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600160109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b60026020528060005260406000206000915090505481565b806000806101000a81548163ffffffff021916908363ffffffff16021790555050565b60008060009054906101000a900463ffffffff16905090565b80600060046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b80600160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b80600160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b6000816fffffffffffffffffffffffffffffffff16600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b61045d81610438565b82525050565b60006020820190506104786000830184610454565b92915050565b600080fd5b61048c81610438565b811461049757600080fd5b50565b6000813590506104a981610483565b92915050565b6000602082840312156104c5576104c461047e565b5b60006104d38482850161049a565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610507826104dc565b9050919050565b610517816104fc565b811461052257600080fd5b50565b6000813590506105348161050e565b92915050565b6000602082840312156105505761054f61047e565b5b600061055e84828501610525565b91505092915050565b6000819050919050565b61057a81610567565b82525050565b60006020820190506105956000830184610571565b92915050565b600063ffffffff82169050919050565b6105b48161059b565b81146105bf57600080fd5b50565b6000813590506105d1816105ab565b92915050565b6000602082840312156105ed576105ec61047e565b5b60006105fb848285016105c2565b91505092915050565b61060d8161059b565b82525050565b60006020820190506106286000830184610604565b92915050565b6000610639826104fc565b9050919050565b6106498161062e565b811461065457600080fd5b50565b60008135905061066681610640565b92915050565b6000602082840312156106825761068161047e565b5b600061069084828501610657565b91505092915050565b6000819050919050565b60006106be6106b96106b4846104dc565b610699565b6104dc565b9050919050565b60006106d0826106a3565b9050919050565b60006106e2826106c5565b9050919050565b6106f2816106d7565b82525050565b600060208201905061070d60008301846106e9565b92915050565b6000806040838503121561072a5761072961047e565b5b600061073885828601610525565b92505060206107498582860161049a565b915050925092905056fea2646970667358221220b4ea7b7c14b0313b86ae4cd8a9b149b62d07606039d92f341acb0db97be083ec64736f6c634300080d0033"

// DeployYield deploys a new Ethereum contract, binding an instance of Yield to it.
func DeployYield(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Yield, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(YieldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Yield{YieldCaller: YieldCaller{contract: contract}, YieldTransactor: YieldTransactor{contract: contract}, YieldFilterer: YieldFilterer{contract: contract}}, nil
}

// Yield is an auto generated Go binding around an Ethereum contract.
type Yield struct {
	YieldCaller     // Read-only binding to the contract
	YieldTransactor // Write-only binding to the contract
	YieldFilterer   // Log filterer for contract events
}

// YieldCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldSession struct {
	Contract     *Yield            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldCallerSession struct {
	Contract *YieldCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YieldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldTransactorSession struct {
	Contract     *YieldTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldRaw struct {
	Contract *Yield // Generic contract binding to access the raw methods on
}

// YieldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldCallerRaw struct {
	Contract *YieldCaller // Generic read-only contract binding to access the raw methods on
}

// YieldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldTransactorRaw struct {
	Contract *YieldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYield creates a new instance of Yield, bound to a specific deployed contract.
func NewYield(address common.Address, backend bind.ContractBackend) (*Yield, error) {
	contract, err := bindYield(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yield{YieldCaller: YieldCaller{contract: contract}, YieldTransactor: YieldTransactor{contract: contract}, YieldFilterer: YieldFilterer{contract: contract}}, nil
}

// NewYieldCaller creates a new read-only instance of Yield, bound to a specific deployed contract.
func NewYieldCaller(address common.Address, caller bind.ContractCaller) (*YieldCaller, error) {
	contract, err := bindYield(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldCaller{contract: contract}, nil
}

// NewYieldTransactor creates a new write-only instance of Yield, bound to a specific deployed contract.
func NewYieldTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldTransactor, error) {
	contract, err := bindYield(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTransactor{contract: contract}, nil
}

// NewYieldFilterer creates a new log filterer instance of Yield, bound to a specific deployed contract.
func NewYieldFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldFilterer, error) {
	contract, err := bindYield(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldFilterer{contract: contract}, nil
}

// bindYield binds a generic wrapper to an already deployed contract.
func bindYield(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YieldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yield *YieldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yield.Contract.YieldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yield *YieldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yield.Contract.YieldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yield *YieldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yield.Contract.YieldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yield *YieldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yield.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yield *YieldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yield.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yield *YieldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yield.Contract.contract.Transact(opts, method, params...)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldSession) Base() (common.Address, error) {
	return _Yield.Contract.Base(&_Yield.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_Yield *YieldCallerSession) Base() (common.Address, error) {
	return _Yield.Contract.Base(&_Yield.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldCaller) Maturity(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldSession) Maturity() (uint32, error) {
	return _Yield.Contract.Maturity(&_Yield.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint32)
func (_Yield *YieldCallerSession) Maturity() (uint32, error) {
	return _Yield.Contract.Maturity(&_Yield.CallOpts)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldCaller) SellBaseCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "sellBaseCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _Yield.Contract.SellBaseCalled(&_Yield.CallOpts, arg0)
}

// SellBaseCalled is a free data retrieval call binding the contract method 0x19808486.
//
// Solidity: function sellBaseCalled(address ) view returns(uint256)
func (_Yield *YieldCallerSession) SellBaseCalled(arg0 common.Address) (*big.Int, error) {
	return _Yield.Contract.SellBaseCalled(&_Yield.CallOpts, arg0)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldCaller) SellBasePreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yield.contract.Call(opts, &out, "sellBasePreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldSession) SellBasePreviewCalled() (*big.Int, error) {
	return _Yield.Contract.SellBasePreviewCalled(&_Yield.CallOpts)
}

// SellBasePreviewCalled is a free data retrieval call binding the contract method 0x1266b30c.
//
// Solidity: function sellBasePreviewCalled() view returns(uint128)
func (_Yield *YieldCallerSession) SellBasePreviewCalled() (*big.Int, error) {
	return _Yield.Contract.SellBasePreviewCalled(&_Yield.CallOpts)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldTransactor) BaseReturns(opts *bind.TransactOpts, i common.Address) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "baseReturns", i)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldSession) BaseReturns(i common.Address) (*types.Transaction, error) {
	return _Yield.Contract.BaseReturns(&_Yield.TransactOpts, i)
}

// BaseReturns is a paid mutator transaction binding the contract method 0x439ec67d.
//
// Solidity: function baseReturns(address i) returns()
func (_Yield *YieldTransactorSession) BaseReturns(i common.Address) (*types.Transaction, error) {
	return _Yield.Contract.BaseReturns(&_Yield.TransactOpts, i)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldTransactor) MaturityReturns(opts *bind.TransactOpts, m uint32) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "maturityReturns", m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldSession) MaturityReturns(m uint32) (*types.Transaction, error) {
	return _Yield.Contract.MaturityReturns(&_Yield.TransactOpts, m)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0x1cc52d77.
//
// Solidity: function maturityReturns(uint32 m) returns()
func (_Yield *YieldTransactorSession) MaturityReturns(m uint32) (*types.Transaction, error) {
	return _Yield.Contract.MaturityReturns(&_Yield.TransactOpts, m)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldTransactor) SellBase(opts *bind.TransactOpts, a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBase", a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBase(&_Yield.TransactOpts, a, u)
}

// SellBase is a paid mutator transaction binding the contract method 0xbcc1694f.
//
// Solidity: function sellBase(address a, uint128 u) returns(uint128)
func (_Yield *YieldTransactorSession) SellBase(a common.Address, u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBase(&_Yield.TransactOpts, a, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldTransactor) SellBasePreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBasePreview", u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreview(&_Yield.TransactOpts, u)
}

// SellBasePreview is a paid mutator transaction binding the contract method 0x13e7bc8c.
//
// Solidity: function sellBasePreview(uint128 u) returns(uint128)
func (_Yield *YieldTransactorSession) SellBasePreview(u *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreview(&_Yield.TransactOpts, u)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldTransactor) SellBasePreviewReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBasePreviewReturns", r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreviewReturns(&_Yield.TransactOpts, r)
}

// SellBasePreviewReturns is a paid mutator transaction binding the contract method 0x5e5dc52d.
//
// Solidity: function sellBasePreviewReturns(uint128 r) returns()
func (_Yield *YieldTransactorSession) SellBasePreviewReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBasePreviewReturns(&_Yield.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldTransactor) SellBaseReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Yield.contract.Transact(opts, "sellBaseReturns", r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBaseReturns(&_Yield.TransactOpts, r)
}

// SellBaseReturns is a paid mutator transaction binding the contract method 0x808c58dd.
//
// Solidity: function sellBaseReturns(uint128 r) returns()
func (_Yield *YieldTransactorSession) SellBaseReturns(r *big.Int) (*types.Transaction, error) {
	return _Yield.Contract.SellBaseReturns(&_Yield.TransactOpts, r)
}
