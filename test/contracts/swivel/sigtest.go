// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swivel

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

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SigTestABI is the input ABI used to generate the binding from.
const SigTestABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"recoverTest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitTest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigTestBin is the compiled bytecode used for deploying new contracts.
var SigTestBin = "0x608060405234801561001057600080fd5b5061071d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063458b017a1461003b578063ecf888711461006d575b600080fd5b61005560048036038101906100509190610372565b61009d565b60405161006493929190610589565b60405180910390f35b61008760048036038101906100829190610336565b6100b8565b60405161009491906104c9565b60405180910390f35b60008060006100ab846100cc565b9250925092509193909250565b60006100c48383610134565b905092915050565b60008060006041845114610115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010c90610549565b60405180910390fd5b6020840151915060408401519050606084015160001a92509193909250565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0826040013560001c11156101a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019790610569565b60405180910390fd5b601b8260000160208101906101b591906103b3565b60ff1614806101d95750601c8260000160208101906101d491906103b3565b60ff16145b610218576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020f90610529565b60405180910390fd5b60018383600001602081019061022e91906103b3565b846020013585604001356040516000815260200160405260405161025594939291906104e4565b6020604051602081039080840390855afa158015610277573d6000803e3d6000fd5b50505060206040510351905092915050565b600061029c610297846105f1565b6105c0565b9050828152602081018484840111156102b457600080fd5b6102bf84828561067b565b509392505050565b6000813590506102d6816106b9565b92915050565b600082601f8301126102ed57600080fd5b81356102fd848260208601610289565b91505092915050565b60006060828403121561031857600080fd5b81905092915050565b600081359050610330816106d0565b92915050565b6000806080838503121561034957600080fd5b6000610357858286016102c7565b925050602061036885828601610306565b9150509250929050565b60006020828403121561038457600080fd5b600082013567ffffffffffffffff81111561039e57600080fd5b6103aa848285016102dc565b91505092915050565b6000602082840312156103c557600080fd5b60006103d384828501610321565b91505092915050565b6103e581610632565b82525050565b6103f481610644565b82525050565b6000610407601b83610621565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b6000610447601883610621565b91507f696e76616c6964207369676e6174757265206c656e67746800000000000000006000830152602082019050919050565b6000610487601b83610621565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b6104c38161066e565b82525050565b60006020820190506104de60008301846103dc565b92915050565b60006080820190506104f960008301876103eb565b61050660208301866104ba565b61051360408301856103eb565b61052060608301846103eb565b95945050505050565b60006020820190508181036000830152610542816103fa565b9050919050565b600060208201905081810360008301526105628161043a565b9050919050565b600060208201905081810360008301526105828161047a565b9050919050565b600060608201905061059e60008301866104ba565b6105ab60208301856103eb565b6105b860408301846103eb565b949350505050565b6000604051905081810181811067ffffffffffffffff821117156105e7576105e661068a565b5b8060405250919050565b600067ffffffffffffffff82111561060c5761060b61068a565b5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b600061063d8261064e565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106c281610644565b81146106cd57600080fd5b50565b6106d98161066e565b81146106e457600080fd5b5056fea264697066735822122014019cfc97a3e215d466206ec5502d18c3bc975d1a5833c18d6d668e37c6eeda64736f6c63430008000033"

// DeploySigTest deploys a new Ethereum contract, binding an instance of SigTest to it.
func DeploySigTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigTest, error) {
	parsed, err := abi.JSON(strings.NewReader(SigTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigTest{SigTestCaller: SigTestCaller{contract: contract}, SigTestTransactor: SigTestTransactor{contract: contract}, SigTestFilterer: SigTestFilterer{contract: contract}}, nil
}

// SigTest is an auto generated Go binding around an Ethereum contract.
type SigTest struct {
	SigTestCaller     // Read-only binding to the contract
	SigTestTransactor // Write-only binding to the contract
	SigTestFilterer   // Log filterer for contract events
}

// SigTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigTestSession struct {
	Contract     *SigTest          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigTestCallerSession struct {
	Contract *SigTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SigTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigTestTransactorSession struct {
	Contract     *SigTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SigTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigTestRaw struct {
	Contract *SigTest // Generic contract binding to access the raw methods on
}

// SigTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigTestCallerRaw struct {
	Contract *SigTestCaller // Generic read-only contract binding to access the raw methods on
}

// SigTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigTestTransactorRaw struct {
	Contract *SigTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigTest creates a new instance of SigTest, bound to a specific deployed contract.
func NewSigTest(address common.Address, backend bind.ContractBackend) (*SigTest, error) {
	contract, err := bindSigTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigTest{SigTestCaller: SigTestCaller{contract: contract}, SigTestTransactor: SigTestTransactor{contract: contract}, SigTestFilterer: SigTestFilterer{contract: contract}}, nil
}

// NewSigTestCaller creates a new read-only instance of SigTest, bound to a specific deployed contract.
func NewSigTestCaller(address common.Address, caller bind.ContractCaller) (*SigTestCaller, error) {
	contract, err := bindSigTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigTestCaller{contract: contract}, nil
}

// NewSigTestTransactor creates a new write-only instance of SigTest, bound to a specific deployed contract.
func NewSigTestTransactor(address common.Address, transactor bind.ContractTransactor) (*SigTestTransactor, error) {
	contract, err := bindSigTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigTestTransactor{contract: contract}, nil
}

// NewSigTestFilterer creates a new log filterer instance of SigTest, bound to a specific deployed contract.
func NewSigTestFilterer(address common.Address, filterer bind.ContractFilterer) (*SigTestFilterer, error) {
	contract, err := bindSigTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigTestFilterer{contract: contract}, nil
}

// bindSigTest binds a generic wrapper to an already deployed contract.
func bindSigTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigTest *SigTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SigTest.Contract.SigTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigTest *SigTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigTest.Contract.SigTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigTest *SigTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigTest.Contract.SigTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigTest *SigTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SigTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigTest *SigTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigTest *SigTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigTest.Contract.contract.Transact(opts, method, params...)
}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigTest *SigTestCaller) RecoverTest(opts *bind.CallOpts, h [32]byte, c SigComponents) (common.Address, error) {
	var out []interface{}
	err := _SigTest.contract.Call(opts, &out, "recoverTest", h, c)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigTest *SigTestSession) RecoverTest(h [32]byte, c SigComponents) (common.Address, error) {
	return _SigTest.Contract.RecoverTest(&_SigTest.CallOpts, h, c)
}

// RecoverTest is a free data retrieval call binding the contract method 0xecf88871.
//
// Solidity: function recoverTest(bytes32 h, (uint8,bytes32,bytes32) c) pure returns(address)
func (_SigTest *SigTestCallerSession) RecoverTest(h [32]byte, c SigComponents) (common.Address, error) {
	return _SigTest.Contract.RecoverTest(&_SigTest.CallOpts, h, c)
}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigTest *SigTestCaller) SplitTest(opts *bind.CallOpts, sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _SigTest.contract.Call(opts, &out, "splitTest", sig)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = out[0].(uint8)
	outstruct.R = out[1].([32]byte)
	outstruct.S = out[2].([32]byte)

	return *outstruct, err

}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigTest *SigTestSession) SplitTest(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigTest.Contract.SplitTest(&_SigTest.CallOpts, sig)
}

// SplitTest is a free data retrieval call binding the contract method 0x458b017a.
//
// Solidity: function splitTest(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_SigTest *SigTestCallerSession) SplitTest(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _SigTest.Contract.SplitTest(&_SigTest.CallOpts, sig)
}
