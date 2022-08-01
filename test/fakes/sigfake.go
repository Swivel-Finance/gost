// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fakes

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

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SigFakeMetaData contains all meta data concerning the SigFake contract.
var SigFakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Length\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"S\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"recoverTest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitTest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106be806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063458b017a1461003b578063ecf888711461006d575b600080fd5b61005560048036038101906100509190610449565b61009d565b604051610064939291906104c7565b60405180910390f35b6100876004803603810190610082919061054e565b6100b8565b60405161009491906105cf565b60405180910390f35b60008060006100ab846100cc565b9250925092509193909250565b60006100c4838361013c565b905092915050565b6000806000604184511461010c576040517f8217288200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060006020870151925060408701519150606087015160001a90508083839550955095505050509193909250565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0826040013560001c111561019f576040517f4be1c79600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601b8260000160208101906101b49190610616565b60ff16141580156101db5750601c8260000160208101906101d59190610616565b60ff1614155b15610212576040517f205db2c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060018484600001602081019061022a9190610616565b85602001358660400135604051600081526020016040526040516102519493929190610643565b6020604051602081039080840390855afa158015610273573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102e5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8091505092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103568261030d565b810181811067ffffffffffffffff821117156103755761037461031e565b5b80604052505050565b60006103886102ef565b9050610394828261034d565b919050565b600067ffffffffffffffff8211156103b4576103b361031e565b5b6103bd8261030d565b9050602081019050919050565b82818337600083830152505050565b60006103ec6103e784610399565b61037e565b90508281526020810184848401111561040857610407610308565b5b6104138482856103ca565b509392505050565b600082601f8301126104305761042f610303565b5b81356104408482602086016103d9565b91505092915050565b60006020828403121561045f5761045e6102f9565b5b600082013567ffffffffffffffff81111561047d5761047c6102fe565b5b6104898482850161041b565b91505092915050565b600060ff82169050919050565b6104a881610492565b82525050565b6000819050919050565b6104c1816104ae565b82525050565b60006060820190506104dc600083018661049f565b6104e960208301856104b8565b6104f660408301846104b8565b949350505050565b610507816104ae565b811461051257600080fd5b50565b600081359050610524816104fe565b92915050565b600080fd5b6000606082840312156105455761054461052a565b5b81905092915050565b60008060808385031215610565576105646102f9565b5b600061057385828601610515565b92505060206105848582860161052f565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105b98261058e565b9050919050565b6105c9816105ae565b82525050565b60006020820190506105e460008301846105c0565b92915050565b6105f381610492565b81146105fe57600080fd5b50565b600081359050610610816105ea565b92915050565b60006020828403121561062c5761062b6102f9565b5b600061063a84828501610601565b91505092915050565b600060808201905061065860008301876104b8565b610665602083018661049f565b61067260408301856104b8565b61067f60608301846104b8565b9594505050505056fea2646970667358221220a7afcf172492943fb7e289e153444447b6af75a7b95eda06a75bb1c49c1a2db664736f6c634300080d0033",
}

// SigFakeABI is the input ABI used to generate the binding from.
// Deprecated: Use SigFakeMetaData.ABI instead.
var SigFakeABI = SigFakeMetaData.ABI

// SigFakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SigFakeMetaData.Bin instead.
var SigFakeBin = SigFakeMetaData.Bin

// DeploySigFake deploys a new Ethereum contract, binding an instance of SigFake to it.
func DeploySigFake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigFake, error) {
	parsed, err := SigFakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SigFakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigFake{SigFakeCaller: SigFakeCaller{contract: contract}, SigFakeTransactor: SigFakeTransactor{contract: contract}, SigFakeFilterer: SigFakeFilterer{contract: contract}}, nil
}

// SigFake is an auto generated Go binding around an Ethereum contract.
type SigFake struct {
	SigFakeCaller     // Read-only binding to the contract
	SigFakeTransactor // Write-only binding to the contract
	SigFakeFilterer   // Log filterer for contract events
}

// SigFakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigFakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigFakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigFakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigFakeSession struct {
	Contract     *SigFake          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigFakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigFakeCallerSession struct {
	Contract *SigFakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SigFakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigFakeTransactorSession struct {
	Contract     *SigFakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SigFakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigFakeRaw struct {
	Contract *SigFake // Generic contract binding to access the raw methods on
}

// SigFakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigFakeCallerRaw struct {
	Contract *SigFakeCaller // Generic read-only contract binding to access the raw methods on
}

// SigFakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigFakeTransactorRaw struct {
	Contract *SigFakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigFake creates a new instance of SigFake, bound to a specific deployed contract.
func NewSigFake(address common.Address, backend bind.ContractBackend) (*SigFake, error) {
	contract, err := bindSigFake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigFake{SigFakeCaller: SigFakeCaller{contract: contract}, SigFakeTransactor: SigFakeTransactor{contract: contract}, SigFakeFilterer: SigFakeFilterer{contract: contract}}, nil
}

// NewSigFakeCaller creates a new read-only instance of SigFake, bound to a specific deployed contract.
func NewSigFakeCaller(address common.Address, caller bind.ContractCaller) (*SigFakeCaller, error) {
	contract, err := bindSigFake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigFakeCaller{contract: contract}, nil
}

// NewSigFakeTransactor creates a new write-only instance of SigFake, bound to a specific deployed contract.
func NewSigFakeTransactor(address common.Address, transactor bind.ContractTransactor) (*SigFakeTransactor, error) {
	contract, err := bindSigFake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigFakeTransactor{contract: contract}, nil
}

// NewSigFakeFilterer creates a new log filterer instance of SigFake, bound to a specific deployed contract.
func NewSigFakeFilterer(address common.Address, filterer bind.ContractFilterer) (*SigFakeFilterer, error) {
	contract, err := bindSigFake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigFakeFilterer{contract: contract}, nil
}

// bindSigFake binds a generic wrapper to an already deployed contract.
func bindSigFake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigFakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigFake *SigFakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SigFake.Contract.SigFakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigFake *SigFakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigFake.Contract.SigFakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigFake *SigFakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigFake.Contract.SigFakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigFake *SigFakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SigFake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigFake *SigFakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigFake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigFake *SigFakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigFake.Contract.contract.Transact(opts, method, params...)
}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigFake *SigFakeCaller) RecoverTest(opts *bind.CallOpts, h [32]byte, c SigComponents) (common.Address, error) {
	var out []interface{}
	err := _SigFake.contract.Call(opts, &out, "recoverTest", h, c)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigFake *SigFakeSession) RecoverTest(h [32]byte, c SigComponents) (common.Address, error) {
	return _SigFake.Contract.RecoverTest(&_SigFake.CallOpts, h, c)
}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigFake *SigFakeCallerSession) RecoverTest(h [32]byte, c SigComponents) (common.Address, error) {
	return _SigFake.Contract.RecoverTest(&_SigFake.CallOpts, h, c)
}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigFake *SigFakeCaller) SplitTest(opts *bind.CallOpts, sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _SigFake.contract.Call(opts, &out, "splitTest", sig)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigFake *SigFakeSession) SplitTest(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigFake.Contract.SplitTest(&_SigFake.CallOpts, sig)
}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigFake *SigFakeCallerSession) SplitTest(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigFake.Contract.SplitTest(&_SigFake.CallOpts, sig)
}
