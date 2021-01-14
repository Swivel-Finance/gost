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

// SigTestABI is the input ABI used to generate the binding from.
const SigTestABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recoverTest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitTest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigTestBin is the compiled bytecode used for deploying new contracts.
var SigTestBin = "0x608060405234801561001057600080fd5b506106c4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063279758d81461003b578063458b017a1461006b575b600080fd5b610055600480360381019061005091906102df565b61009d565b6040516100629190610470565b60405180910390f35b61008560048036038101906100809190610342565b6100b5565b60405161009493929190610530565b60405180910390f35b60006100ab858585856100d0565b9050949350505050565b60008060006100c3846101e5565b9250925092509193909250565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610138576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012f90610510565b60405180910390fd5b601b8460ff16148061014d5750601c8460ff16145b61018c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610183906104d0565b60405180910390fd5b600185858585604051600081526020016040526040516101af949392919061048b565b6020604051602081039080840390855afa1580156101d1573d6000803e3d6000fd5b505050602060405103519050949350505050565b6000806000604184511461022e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610225906104f0565b60405180910390fd5b6020840151915060408401519050606084015160001a92509193909250565b600061026061025b84610598565b610567565b90508281526020810184848401111561027857600080fd5b610283848285610622565b509392505050565b60008135905061029a81610660565b92915050565b600082601f8301126102b157600080fd5b81356102c184826020860161024d565b91505092915050565b6000813590506102d981610677565b92915050565b600080600080608085870312156102f557600080fd5b60006103038782880161028b565b9450506020610314878288016102ca565b93505060406103258782880161028b565b92505060606103368782880161028b565b91505092959194509250565b60006020828403121561035457600080fd5b600082013567ffffffffffffffff81111561036e57600080fd5b61037a848285016102a0565b91505092915050565b61038c816105d9565b82525050565b61039b816105eb565b82525050565b60006103ae601b836105c8565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b60006103ee6018836105c8565b91507f696e76616c6964207369676e6174757265206c656e67746800000000000000006000830152602082019050919050565b600061042e601b836105c8565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b61046a81610615565b82525050565b60006020820190506104856000830184610383565b92915050565b60006080820190506104a06000830187610392565b6104ad6020830186610461565b6104ba6040830185610392565b6104c76060830184610392565b95945050505050565b600060208201905081810360008301526104e9816103a1565b9050919050565b60006020820190508181036000830152610509816103e1565b9050919050565b6000602082019050818103600083015261052981610421565b9050919050565b60006060820190506105456000830186610461565b6105526020830185610392565b61055f6040830184610392565b949350505050565b6000604051905081810181811067ffffffffffffffff8211171561058e5761058d610631565b5b8060405250919050565b600067ffffffffffffffff8211156105b3576105b2610631565b5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b60006105e4826105f5565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610669816105eb565b811461067457600080fd5b50565b61068081610615565b811461068b57600080fd5b5056fea2646970667358221220e3964c4ba257e7eb4519c5af5c4c546752a57f8fd45d0ce9cc32a92642eb4c5d64736f6c63430008000033"

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

// RecoverTest is a free data retrieval call binding the contract method 0x279758d8.
//
// Solidity: function recoverTest(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_SigTest *SigTestCaller) RecoverTest(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SigTest.contract.Call(opts, &out, "recoverTest", h, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverTest is a free data retrieval call binding the contract method 0x279758d8.
//
// Solidity: function recoverTest(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_SigTest *SigTestSession) RecoverTest(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _SigTest.Contract.RecoverTest(&_SigTest.CallOpts, h, v, r, s)
}

// RecoverTest is a free data retrieval call binding the contract method 0x279758d8.
//
// Solidity: function recoverTest(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_SigTest *SigTestCallerSession) RecoverTest(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _SigTest.Contract.RecoverTest(&_SigTest.CallOpts, h, v, r, s)
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
