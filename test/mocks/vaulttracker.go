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

// VaultTrackerABI is the input ABI used to generate the binding from.
const VaultTrackerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"addNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"matureVaultReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemInterestReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"removeNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferNotionalFeeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferNotionalFeeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferNotionalFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferNotionalFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
var VaultTrackerBin = "0x60806040523480156200001157600080fd5b5060405162000d1938038062000d198339818101604052810190620000379190620000f7565b8260078190555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620001bf565b600081519050620000da816200018b565b92915050565b600081519050620000f181620001a5565b92915050565b6000806000606084860312156200010d57600080fd5b60006200011d86828701620000e0565b93505060206200013086828701620000c9565b92505060406200014386828701620000c9565b9150509250925092565b60006200015a8262000161565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b62000196816200014d565b8114620001a257600080fd5b50565b620001b08162000181565b8114620001bc57600080fd5b50565b610b4a80620001cf6000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806375d69a39116100b8578063b7dd34831161007c578063b7dd348314610374578063bbce238614610392578063d0b9d032146103c2578063d6cb2c0d146103de578063da3de9e9146103fa578063e590c3621461041657610137565b806375d69a39146102be578063a01cfffb146102dc578063a701da691461030c578063b326258d14610328578063b4c4a4c81461035857610137565b80633dfa1f41116100ff5780633dfa1f41146102085780635c70b7c1146102385780635dfe12ac14610254578063613a28d1146102705780636b868d51146102a057610137565b8063012b264a1461013c5780630aa93b9b1461015a578063177946731461018a57806319caf46c146101ba578063204f83f9146101ea575b600080fd5b610144610447565b6040516101519190610a0d565b60405180910390f35b610174600480360381019061016f91906108da565b61046d565b6040516101819190610a6c565b60405180910390f35b6101a4600480360381019061019f9190610903565b610485565b6040516101b19190610a51565b60405180910390f35b6101d460048036038101906101cf91906108da565b61057f565b6040516101e19190610a6c565b60405180910390f35b6101f26105cc565b6040516101ff9190610a6c565b60405180910390f35b610222600480360381019061021d91906108da565b6105d6565b60405161022f9190610a6c565b60405180910390f35b610252600480360381019061024d919061098e565b6105ee565b005b61026e6004803603810190610269919061098e565b61060b565b005b61028a60048036038101906102859190610952565b610628565b6040516102979190610a51565b60405180910390f35b6102a8610686565b6040516102b59190610a51565b60405180910390f35b6102c661069d565b6040516102d39190610a0d565b60405180910390f35b6102f660048036038101906102f19190610952565b6106c3565b6040516103039190610a51565b60405180910390f35b6103266004803603810190610321919061098e565b610720565b005b610342600480360381019061033d9190610952565b61073d565b60405161034f9190610a51565b60405180910390f35b610372600480360381019061036d91906109b7565b61079b565b005b61037c6107a5565b6040516103899190610a0d565b60405180910390f35b6103ac60048036038101906103a791906108da565b6107cb565b6040516103b99190610a6c565b60405180910390f35b6103dc60048036038101906103d7919061098e565b6107e3565b005b6103f860048036038101906103f391906109b7565b610800565b005b610414600480360381019061040f919061098e565b61080a565b005b610430600480360381019061042b91906108da565b610827565b60405161043e929190610a28565b60405180910390f35b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090505481565b600061048f61086b565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600960039054906101000a900460ff169150509392505050565b600081600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008549050919050565b6000600754905090565b60006020528060005260406000206000915090505481565b80600960016101000a81548160ff02191690831515021790555050565b80600960026101000a81548160ff02191690831515021790555050565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600960029054906101000a900460ff16905092915050565b6000600960009054906101000a900460ff16905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600960019054906101000a900460ff16905092915050565b80600960046101000a81548160ff02191690831515021790555050565b600081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600960049054906101000a900460ff16905092915050565b8060078190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b80600960036101000a81548160ff02191690831515021790555050565b8060088190555050565b80600960006101000a81548160ff02191690831515021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6000813590506108aa81610acf565b92915050565b6000813590506108bf81610ae6565b92915050565b6000813590506108d481610afd565b92915050565b6000602082840312156108ec57600080fd5b60006108fa8482850161089b565b91505092915050565b60008060006060848603121561091857600080fd5b60006109268682870161089b565b93505060206109378682870161089b565b9250506040610948868287016108c5565b9150509250925092565b6000806040838503121561096557600080fd5b60006109738582860161089b565b9250506020610984858286016108c5565b9150509250929050565b6000602082840312156109a057600080fd5b60006109ae848285016108b0565b91505092915050565b6000602082840312156109c957600080fd5b60006109d7848285016108c5565b91505092915050565b6109e981610a87565b82525050565b6109f881610a99565b82525050565b610a0781610ac5565b82525050565b6000602082019050610a2260008301846109e0565b92915050565b6000604082019050610a3d60008301856109e0565b610a4a60208301846109fe565b9392505050565b6000602082019050610a6660008301846109ef565b92915050565b6000602082019050610a8160008301846109fe565b92915050565b6000610a9282610aa5565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610ad881610a87565b8114610ae357600080fd5b50565b610aef81610a99565b8114610afa57600080fd5b50565b610b0681610ac5565b8114610b1157600080fd5b5056fea264697066735822122070f97c12705575cb4b6608c34d6fbc0367a3ead3e58d3b4fca60efdd587f15e764736f6c63430008040033"

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, m *big.Int, c common.Address, s common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultTrackerBin), backend, m, c, s)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// VaultTracker is an auto generated Go binding around an Ethereum contract.
type VaultTracker struct {
	VaultTrackerCaller     // Read-only binding to the contract
	VaultTrackerTransactor // Write-only binding to the contract
	VaultTrackerFilterer   // Log filterer for contract events
}

// VaultTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultTrackerSession struct {
	Contract     *VaultTracker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultTrackerCallerSession struct {
	Contract *VaultTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTrackerTransactorSession struct {
	Contract     *VaultTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultTrackerRaw struct {
	Contract *VaultTracker // Generic contract binding to access the raw methods on
}

// VaultTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultTrackerCallerRaw struct {
	Contract *VaultTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTrackerTransactorRaw struct {
	Contract *VaultTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultTracker creates a new instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTracker(address common.Address, backend bind.ContractBackend) (*VaultTracker, error) {
	contract, err := bindVaultTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// NewVaultTrackerCaller creates a new read-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerCaller(address common.Address, caller bind.ContractCaller) (*VaultTrackerCaller, error) {
	contract, err := bindVaultTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerCaller{contract: contract}, nil
}

// NewVaultTrackerTransactor creates a new write-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTrackerTransactor, error) {
	contract, err := bindVaultTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerTransactor{contract: contract}, nil
}

// NewVaultTrackerFilterer creates a new log filterer instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultTrackerFilterer, error) {
	contract, err := bindVaultTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerFilterer{contract: contract}, nil
}

// bindVaultTracker binds a generic wrapper to an already deployed contract.
func bindVaultTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.VaultTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transact(opts, method, params...)
}

// AddNotionalCalled is a free data retrieval call binding the contract method 0x3dfa1f41.
//
// Solidity: function addNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) AddNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "addNotionalCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddNotionalCalled is a free data retrieval call binding the contract method 0x3dfa1f41.
//
// Solidity: function addNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerSession) AddNotionalCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.AddNotionalCalled(&_VaultTracker.CallOpts, arg0)
}

// AddNotionalCalled is a free data retrieval call binding the contract method 0x3dfa1f41.
//
// Solidity: function addNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) AddNotionalCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.AddNotionalCalled(&_VaultTracker.CallOpts, arg0)
}

// BsAddr is a free data retrieval call binding the contract method 0x75d69a39.
//
// Solidity: function bsAddr() view returns(address)
func (_VaultTracker *VaultTrackerCaller) BsAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "bsAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BsAddr is a free data retrieval call binding the contract method 0x75d69a39.
//
// Solidity: function bsAddr() view returns(address)
func (_VaultTracker *VaultTrackerSession) BsAddr() (common.Address, error) {
	return _VaultTracker.Contract.BsAddr(&_VaultTracker.CallOpts)
}

// BsAddr is a free data retrieval call binding the contract method 0x75d69a39.
//
// Solidity: function bsAddr() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) BsAddr() (common.Address, error) {
	return _VaultTracker.Contract.BsAddr(&_VaultTracker.CallOpts)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCaller) CTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "cTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// MatureVault is a free data retrieval call binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() view returns(bool)
func (_VaultTracker *VaultTrackerCaller) MatureVault(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "matureVault")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MatureVault is a free data retrieval call binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() view returns(bool)
func (_VaultTracker *VaultTrackerSession) MatureVault() (bool, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.CallOpts)
}

// MatureVault is a free data retrieval call binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() view returns(bool)
func (_VaultTracker *VaultTrackerCallerSession) MatureVault() (bool, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// RemoveNotionalCalled is a free data retrieval call binding the contract method 0x0aa93b9b.
//
// Solidity: function removeNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) RemoveNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "removeNotionalCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveNotionalCalled is a free data retrieval call binding the contract method 0x0aa93b9b.
//
// Solidity: function removeNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerSession) RemoveNotionalCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.RemoveNotionalCalled(&_VaultTracker.CallOpts, arg0)
}

// RemoveNotionalCalled is a free data retrieval call binding the contract method 0x0aa93b9b.
//
// Solidity: function removeNotionalCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) RemoveNotionalCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.RemoveNotionalCalled(&_VaultTracker.CallOpts, arg0)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCaller) Swivel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "swivel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
}

// TransferNotionalFeeCalled is a free data retrieval call binding the contract method 0xbbce2386.
//
// Solidity: function transferNotionalFeeCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) TransferNotionalFeeCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "transferNotionalFeeCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferNotionalFeeCalled is a free data retrieval call binding the contract method 0xbbce2386.
//
// Solidity: function transferNotionalFeeCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFeeCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.TransferNotionalFeeCalled(&_VaultTracker.CallOpts, arg0)
}

// TransferNotionalFeeCalled is a free data retrieval call binding the contract method 0xbbce2386.
//
// Solidity: function transferNotionalFeeCalled(address ) view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) TransferNotionalFeeCalled(arg0 common.Address) (*big.Int, error) {
	return _VaultTracker.Contract.TransferNotionalFeeCalled(&_VaultTracker.CallOpts, arg0)
}

// TransferNotionalFromCalled is a free data retrieval call binding the contract method 0xe590c362.
//
// Solidity: function transferNotionalFromCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerCaller) TransferNotionalFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "transferNotionalFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferNotionalFromCalled is a free data retrieval call binding the contract method 0xe590c362.
//
// Solidity: function transferNotionalFromCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _VaultTracker.Contract.TransferNotionalFromCalled(&_VaultTracker.CallOpts, arg0)
}

// TransferNotionalFromCalled is a free data retrieval call binding the contract method 0xe590c362.
//
// Solidity: function transferNotionalFromCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerCallerSession) TransferNotionalFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _VaultTracker.Contract.TransferNotionalFromCalled(&_VaultTracker.CallOpts, arg0)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) AddNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "addNotional", o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// AddNotionalReturns is a paid mutator transaction binding the contract method 0x5c70b7c1.
//
// Solidity: function addNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) AddNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "addNotionalReturns", b)
}

// AddNotionalReturns is a paid mutator transaction binding the contract method 0x5c70b7c1.
//
// Solidity: function addNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) AddNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotionalReturns(&_VaultTracker.TransactOpts, b)
}

// AddNotionalReturns is a paid mutator transaction binding the contract method 0x5c70b7c1.
//
// Solidity: function addNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) AddNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotionalReturns(&_VaultTracker.TransactOpts, b)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) MatureVaultReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "matureVaultReturns", b)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) MatureVaultReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVaultReturns(&_VaultTracker.TransactOpts, b)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) MatureVaultReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVaultReturns(&_VaultTracker.TransactOpts, b)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_VaultTracker *VaultTrackerTransactor) MaturityReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "maturityReturns", n)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_VaultTracker *VaultTrackerSession) MaturityReturns(n *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MaturityReturns(&_VaultTracker.TransactOpts, n)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_VaultTracker *VaultTrackerTransactorSession) MaturityReturns(n *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MaturityReturns(&_VaultTracker.TransactOpts, n)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactor) RedeemInterest(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "redeemInterest", o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactorSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RedeemInterestReturns is a paid mutator transaction binding the contract method 0xd6cb2c0d.
//
// Solidity: function redeemInterestReturns(uint256 a) returns()
func (_VaultTracker *VaultTrackerTransactor) RedeemInterestReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "redeemInterestReturns", a)
}

// RedeemInterestReturns is a paid mutator transaction binding the contract method 0xd6cb2c0d.
//
// Solidity: function redeemInterestReturns(uint256 a) returns()
func (_VaultTracker *VaultTrackerSession) RedeemInterestReturns(a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterestReturns(&_VaultTracker.TransactOpts, a)
}

// RedeemInterestReturns is a paid mutator transaction binding the contract method 0xd6cb2c0d.
//
// Solidity: function redeemInterestReturns(uint256 a) returns()
func (_VaultTracker *VaultTrackerTransactorSession) RedeemInterestReturns(a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterestReturns(&_VaultTracker.TransactOpts, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) RemoveNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "removeNotional", o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// RemoveNotionalReturns is a paid mutator transaction binding the contract method 0x5dfe12ac.
//
// Solidity: function removeNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) RemoveNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "removeNotionalReturns", b)
}

// RemoveNotionalReturns is a paid mutator transaction binding the contract method 0x5dfe12ac.
//
// Solidity: function removeNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) RemoveNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotionalReturns(&_VaultTracker.TransactOpts, b)
}

// RemoveNotionalReturns is a paid mutator transaction binding the contract method 0x5dfe12ac.
//
// Solidity: function removeNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) RemoveNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotionalReturns(&_VaultTracker.TransactOpts, b)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFee(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFee", f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
}

// TransferNotionalFeeReturns is a paid mutator transaction binding the contract method 0xa701da69.
//
// Solidity: function transferNotionalFeeReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFeeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFeeReturns", b)
}

// TransferNotionalFeeReturns is a paid mutator transaction binding the contract method 0xa701da69.
//
// Solidity: function transferNotionalFeeReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) TransferNotionalFeeReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFeeReturns(&_VaultTracker.TransactOpts, b)
}

// TransferNotionalFeeReturns is a paid mutator transaction binding the contract method 0xa701da69.
//
// Solidity: function transferNotionalFeeReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFeeReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFeeReturns(&_VaultTracker.TransactOpts, b)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFrom", f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotionalFromReturns is a paid mutator transaction binding the contract method 0xd0b9d032.
//
// Solidity: function transferNotionalFromReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFromReturns", b)
}

// TransferNotionalFromReturns is a paid mutator transaction binding the contract method 0xd0b9d032.
//
// Solidity: function transferNotionalFromReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) TransferNotionalFromReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFromReturns(&_VaultTracker.TransactOpts, b)
}

// TransferNotionalFromReturns is a paid mutator transaction binding the contract method 0xd0b9d032.
//
// Solidity: function transferNotionalFromReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFromReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFromReturns(&_VaultTracker.TransactOpts, b)
}
