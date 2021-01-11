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

// SigABI is the input ABI used to generate the binding from.
const SigABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"split\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitAndRecover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SigBin is the compiled bytecode used for deploying new contracts.
var SigBin = "0x6107b7610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c806322d1f49e1461005057806365d894e914610080578063c2bf17b0146100b2575b600080fd5b61006a60048036038101906100659190610360565b6100e2565b6040516100779190610563565b60405180910390f35b61009a60048036038101906100959190610417565b610151565b6040516100a993929190610623565b60405180910390f35b6100cc60048036038101906100c791906103b4565b6101b9565b6040516100d99190610563565b60405180910390f35b6000806000806100f185610151565b9250925092506001868484846040516000815260200160405260405161011a949392919061057e565b6020604051602081039080840390855afa15801561013c573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000806000604184511461019a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610191906105e3565b60405180910390fd5b6020840151915060408401519050606084015160001a92509193909250565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610221576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021890610603565b60405180910390fd5b601b8460ff1614806102365750601c8460ff16145b610275576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026c906105c3565b60405180910390fd5b60018585858560405160008152602001604052604051610298949392919061057e565b6020604051602081039080840390855afa1580156102ba573d6000803e3d6000fd5b505050602060405103519050949350505050565b60006102e16102dc8461068b565b61065a565b9050828152602081018484840111156102f957600080fd5b610304848285610715565b509392505050565b60008135905061031b81610753565b92915050565b600082601f83011261033257600080fd5b81356103428482602086016102ce565b91505092915050565b60008135905061035a8161076a565b92915050565b6000806040838503121561037357600080fd5b60006103818582860161030c565b925050602083013567ffffffffffffffff81111561039e57600080fd5b6103aa85828601610321565b9150509250929050565b600080600080608085870312156103ca57600080fd5b60006103d88782880161030c565b94505060206103e98782880161034b565b93505060406103fa8782880161030c565b925050606061040b8782880161030c565b91505092959194509250565b60006020828403121561042957600080fd5b600082013567ffffffffffffffff81111561044357600080fd5b61044f84828501610321565b91505092915050565b610461816106cc565b82525050565b610470816106de565b82525050565b61047f816106de565b82525050565b6000610492601b836106bb565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b60006104d26018836106bb565b91507f696e76616c6964207369676e6174757265206c656e67746800000000000000006000830152602082019050919050565b6000610512601b836106bb565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b61054e81610708565b82525050565b61055d81610708565b82525050565b60006020820190506105786000830184610458565b92915050565b60006080820190506105936000830187610467565b6105a06020830186610545565b6105ad6040830185610467565b6105ba6060830184610467565b95945050505050565b600060208201905081810360008301526105dc81610485565b9050919050565b600060208201905081810360008301526105fc816104c5565b9050919050565b6000602082019050818103600083015261061c81610505565b9050919050565b60006060820190506106386000830186610554565b6106456020830185610476565b6106526040830184610476565b949350505050565b6000604051905081810181811067ffffffffffffffff8211171561068157610680610724565b5b8060405250919050565b600067ffffffffffffffff8211156106a6576106a5610724565b5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b60006106d7826106e8565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600060ff82169050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61075c816106de565b811461076757600080fd5b50565b61077381610708565b811461077e57600080fd5b5056fea2646970667358221220cecd1bae7094f28c5034591c1f34fc9484e8c5a3c04faecd1e243ad387a418a664736f6c63430008000033"

// DeploySig deploys a new Ethereum contract, binding an instance of Sig to it.
func DeploySig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sig, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// Sig is an auto generated Go binding around an Ethereum contract.
type Sig struct {
	SigCaller     // Read-only binding to the contract
	SigTransactor // Write-only binding to the contract
	SigFilterer   // Log filterer for contract events
}

// SigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigSession struct {
	Contract     *Sig              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigCallerSession struct {
	Contract *SigCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigTransactorSession struct {
	Contract     *SigTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigRaw struct {
	Contract *Sig // Generic contract binding to access the raw methods on
}

// SigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigCallerRaw struct {
	Contract *SigCaller // Generic read-only contract binding to access the raw methods on
}

// SigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigTransactorRaw struct {
	Contract *SigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSig creates a new instance of Sig, bound to a specific deployed contract.
func NewSig(address common.Address, backend bind.ContractBackend) (*Sig, error) {
	contract, err := bindSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sig{SigCaller: SigCaller{contract: contract}, SigTransactor: SigTransactor{contract: contract}, SigFilterer: SigFilterer{contract: contract}}, nil
}

// NewSigCaller creates a new read-only instance of Sig, bound to a specific deployed contract.
func NewSigCaller(address common.Address, caller bind.ContractCaller) (*SigCaller, error) {
	contract, err := bindSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigCaller{contract: contract}, nil
}

// NewSigTransactor creates a new write-only instance of Sig, bound to a specific deployed contract.
func NewSigTransactor(address common.Address, transactor bind.ContractTransactor) (*SigTransactor, error) {
	contract, err := bindSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigTransactor{contract: contract}, nil
}

// NewSigFilterer creates a new log filterer instance of Sig, bound to a specific deployed contract.
func NewSigFilterer(address common.Address, filterer bind.ContractFilterer) (*SigFilterer, error) {
	contract, err := bindSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigFilterer{contract: contract}, nil
}

// bindSig binds a generic wrapper to an already deployed contract.
func bindSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.SigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.SigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sig *SigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sig *SigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sig *SigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sig.Contract.contract.Transact(opts, method, params...)
}

// Recover is a free data retrieval call binding the contract method 0xc2bf17b0.
//
// Solidity: function recover(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Sig *SigCaller) Recover(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Sig.contract.Call(opts, &out, "recover", h, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0xc2bf17b0.
//
// Solidity: function recover(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Sig *SigSession) Recover(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Sig.Contract.Recover(&_Sig.CallOpts, h, v, r, s)
}

// Recover is a free data retrieval call binding the contract method 0xc2bf17b0.
//
// Solidity: function recover(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Sig *SigCallerSession) Recover(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Sig.Contract.Recover(&_Sig.CallOpts, h, v, r, s)
}

// Split is a free data retrieval call binding the contract method 0x65d894e9.
//
// Solidity: function split(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Sig *SigCaller) Split(opts *bind.CallOpts, sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Sig.contract.Call(opts, &out, "split", sig)

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

// Split is a free data retrieval call binding the contract method 0x65d894e9.
//
// Solidity: function split(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Sig *SigSession) Split(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Sig.Contract.Split(&_Sig.CallOpts, sig)
}

// Split is a free data retrieval call binding the contract method 0x65d894e9.
//
// Solidity: function split(bytes sig) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Sig *SigCallerSession) Split(sig []byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Sig.Contract.Split(&_Sig.CallOpts, sig)
}

// SplitAndRecover is a free data retrieval call binding the contract method 0x22d1f49e.
//
// Solidity: function splitAndRecover(bytes32 h, bytes sig) pure returns(address)
func (_Sig *SigCaller) SplitAndRecover(opts *bind.CallOpts, h [32]byte, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Sig.contract.Call(opts, &out, "splitAndRecover", h, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SplitAndRecover is a free data retrieval call binding the contract method 0x22d1f49e.
//
// Solidity: function splitAndRecover(bytes32 h, bytes sig) pure returns(address)
func (_Sig *SigSession) SplitAndRecover(h [32]byte, sig []byte) (common.Address, error) {
	return _Sig.Contract.SplitAndRecover(&_Sig.CallOpts, h, sig)
}

// SplitAndRecover is a free data retrieval call binding the contract method 0x22d1f49e.
//
// Solidity: function splitAndRecover(bytes32 h, bytes sig) pure returns(address)
func (_Sig *SigCallerSession) SplitAndRecover(h [32]byte, sig []byte) (common.Address, error) {
	return _Sig.Contract.SplitAndRecover(&_Sig.CallOpts, h, sig)
}
