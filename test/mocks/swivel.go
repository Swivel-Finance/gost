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

// Components is an auto generated low-level Go binding around an user-defined struct.
type Components struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Key        [32]byte
	Maker      common.Address
	Underlying common.Address
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// SwivelMetaData contains all meta data concerning the Swivel contract.
var SwivelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structComponents[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"i\",\"type\":\"bool\"}],\"name\":\"initiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610751806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806338599fc8146100515780634578b3ce14610081578063c268dc2a1461009d578063d2144f58146100cd575b600080fd5b61006b6004803603810190610066919061032f565b6100fd565b6040516100789190610378565b60405180910390f35b61009b600480360381019061009691906103cb565b61011d565b005b6100b760048036038101906100b2919061032f565b610139565b6040516100c49190610411565b60405180910390f35b6100e760048036038101906100e2919061053e565b610151565b6040516100f49190610601565b60405180910390f35b60026020528060005260406000206000915054906101000a900460ff1681565b806000806101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600080600090505b878790508110156102ac578585828181106101775761017661061c565b5b90506020020135600160008a8a858181106101955761019461061c565b5b9050610120020160200160208101906101ae919061032f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508383828181106101ff576101fe61061c565b5b90506060020160000160208101906102179190610677565b600260008a8a8581811061022e5761022d61061c565b5b905061012002016020016020810190610247919061032f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555080806102a4906106d3565b915050610159565b5060008054906101000a900460ff1690509695505050505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102fc826102d1565b9050919050565b61030c816102f1565b811461031757600080fd5b50565b60008135905061032981610303565b92915050565b600060208284031215610345576103446102c7565b5b60006103538482850161031a565b91505092915050565b600060ff82169050919050565b6103728161035c565b82525050565b600060208201905061038d6000830184610369565b92915050565b60008115159050919050565b6103a881610393565b81146103b357600080fd5b50565b6000813590506103c58161039f565b92915050565b6000602082840312156103e1576103e06102c7565b5b60006103ef848285016103b6565b91505092915050565b6000819050919050565b61040b816103f8565b82525050565b60006020820190506104266000830184610402565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126104515761045061042c565b5b8235905067ffffffffffffffff81111561046e5761046d610431565b5b6020830191508361012082028301111561048b5761048a610436565b5b9250929050565b60008083601f8401126104a8576104a761042c565b5b8235905067ffffffffffffffff8111156104c5576104c4610431565b5b6020830191508360208202830111156104e1576104e0610436565b5b9250929050565b60008083601f8401126104fe576104fd61042c565b5b8235905067ffffffffffffffff81111561051b5761051a610431565b5b60208301915083606082028301111561053757610536610436565b5b9250929050565b6000806000806000806060878903121561055b5761055a6102c7565b5b600087013567ffffffffffffffff811115610579576105786102cc565b5b61058589828a0161043b565b9650965050602087013567ffffffffffffffff8111156105a8576105a76102cc565b5b6105b489828a01610492565b9450945050604087013567ffffffffffffffff8111156105d7576105d66102cc565b5b6105e389828a016104e8565b92509250509295509295509295565b6105fb81610393565b82525050565b600060208201905061061660008301846105f2565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6106548161035c565b811461065f57600080fd5b50565b6000813590506106718161064b565b92915050565b60006020828403121561068d5761068c6102c7565b5b600061069b84828501610662565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106de826103f8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107105761070f6106a4565b5b60018201905091905056fea26469706673582212207373274d8823ca64ae780e0272489ce4ee0542d9b6838bb7e2523974a954f9dc64736f6c634300080d0033",
}

// SwivelABI is the input ABI used to generate the binding from.
// Deprecated: Use SwivelMetaData.ABI instead.
var SwivelABI = SwivelMetaData.ABI

// SwivelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwivelMetaData.Bin instead.
var SwivelBin = SwivelMetaData.Bin

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := SwivelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwivelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// Swivel is an auto generated Go binding around an Ethereum contract.
type Swivel struct {
	SwivelCaller     // Read-only binding to the contract
	SwivelTransactor // Write-only binding to the contract
	SwivelFilterer   // Log filterer for contract events
}

// SwivelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelSession struct {
	Contract     *Swivel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelCallerSession struct {
	Contract *SwivelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwivelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTransactorSession struct {
	Contract     *SwivelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelRaw struct {
	Contract *Swivel // Generic contract binding to access the raw methods on
}

// SwivelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelCallerRaw struct {
	Contract *SwivelCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTransactorRaw struct {
	Contract *SwivelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivel creates a new instance of Swivel, bound to a specific deployed contract.
func NewSwivel(address common.Address, backend bind.ContractBackend) (*Swivel, error) {
	contract, err := bindSwivel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// NewSwivelCaller creates a new read-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelCaller(address common.Address, caller bind.ContractCaller) (*SwivelCaller, error) {
	contract, err := bindSwivel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelCaller{contract: contract}, nil
}

// NewSwivelTransactor creates a new write-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTransactor, error) {
	contract, err := bindSwivel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTransactor{contract: contract}, nil
}

// NewSwivelFilterer creates a new log filterer instance of Swivel, bound to a specific deployed contract.
func NewSwivelFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelFilterer, error) {
	contract, err := bindSwivel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelFilterer{contract: contract}, nil
}

// bindSwivel binds a generic wrapper to an already deployed contract.
func bindSwivel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.SwivelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transact(opts, method, params...)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCaller) InitiateCalledAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCallerSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCaller) InitiateCalledSignature(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledSignature", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCallerSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactor) Initiate(opts *bind.TransactOpts, o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiate", o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactorSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactor) InitiateReturns(opts *bind.TransactOpts, i bool) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiateReturns", i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactorSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}
