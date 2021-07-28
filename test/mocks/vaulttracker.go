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
const VaultTrackerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"n\",\"type\":\"bool\"}],\"name\":\"addNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"m\",\"type\":\"bool\"}],\"name\":\"matureVaultReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemInterestReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"n\",\"type\":\"bool\"}],\"name\":\"removeNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferNotionalFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferNotionalFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
var VaultTrackerBin = "0x60806040523480156200001157600080fd5b5060405162000d6c38038062000d6c8339818101604052810190620000379190620000b5565b8160068190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000168565b600081519050620000988162000134565b92915050565b600081519050620000af816200014e565b92915050565b60008060408385031215620000c957600080fd5b6000620000d9858286016200009e565b9250506020620000ec8582860162000087565b9150509250929050565b600062000103826200010a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200013f81620000f6565b81146200014b57600080fd5b50565b62000159816200012a565b81146200016557600080fd5b50565b610bf480620001786000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80637ee01a1c116100ad578063d0b9d03211610071578063d0b9d0321461037e578063d6cb2c0d1461039a578063da3de9e9146103b6578063e590c362146103d2578063fb86ed76146104035761012c565b80637ee01a1c146102b357806391b9bf17146102e3578063a01cfffb14610314578063b4c4a4c814610344578063b7dd3483146103605761012c565b80635c70b7c1116100f45780635c70b7c11461020f5780635dfe12ac1461022b578063613a28d1146102475780636b868d511461027757806375d69a39146102955761012c565b80630aa93b9b14610131578063177946731461016157806319caf46c14610191578063204f83f9146101c15780633dfa1f41146101df575b600080fd5b61014b60048036038101906101469190610984565b61041f565b6040516101589190610b16565b60405180910390f35b61017b600480360381019061017691906109ad565b610437565b6040516101889190610afb565b60405180910390f35b6101ab60048036038101906101a69190610984565b610531565b6040516101b89190610b16565b60405180910390f35b6101c961057e565b6040516101d69190610b16565b60405180910390f35b6101f960048036038101906101f49190610984565b610588565b6040516102069190610b16565b60405180910390f35b61022960048036038101906102249190610a38565b6105a0565b005b61024560048036038101906102409190610a38565b6105bd565b005b610261600480360381019061025c91906109fc565b6105da565b60405161026e9190610afb565b60405180910390f35b61027f610638565b60405161028c9190610afb565b60405180910390f35b61029d61064f565b6040516102aa9190610ab7565b60405180910390f35b6102cd60048036038101906102c891906109ad565b610675565b6040516102da9190610afb565b60405180910390f35b6102fd60048036038101906102f89190610984565b61076f565b60405161030b929190610ad2565b60405180910390f35b61032e600480360381019061032991906109fc565b6107b3565b60405161033b9190610afb565b60405180910390f35b61035e60048036038101906103599190610a61565b610810565b005b61036861081a565b6040516103759190610ab7565b60405180910390f35b61039860048036038101906103939190610a38565b610840565b005b6103b460048036038101906103af9190610a61565b61085d565b005b6103d060048036038101906103cb9190610a38565b610867565b005b6103ec60048036038101906103e79190610984565b610884565b6040516103fa929190610ad2565b60405180910390f35b61041d60048036038101906104189190610a38565b6108c8565b005b60016020528060005260406000206000915090505481565b60006104416108e5565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600860039054906101000a900460ff169150509392505050565b600081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506007549050919050565b6000600654905090565b60006020528060005260406000206000915090505481565b80600860016101000a81548160ff02191690831515021790555050565b80600860026101000a81548160ff02191690831515021790555050565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600860029054906101000a900460ff16905092915050565b6000600860009054906101000a900460ff16905090565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061067f610915565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600860049054906101000a900460ff169150509392505050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600860019054906101000a900460ff16905092915050565b8060068190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600860036101000a81548160ff02191690831515021790555050565b8060078190555050565b80600860006101000a81548160ff02191690831515021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b80600860046101000a81548160ff02191690831515021790555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b60008135905061095481610b79565b92915050565b60008135905061096981610b90565b92915050565b60008135905061097e81610ba7565b92915050565b60006020828403121561099657600080fd5b60006109a484828501610945565b91505092915050565b6000806000606084860312156109c257600080fd5b60006109d086828701610945565b93505060206109e186828701610945565b92505060406109f28682870161096f565b9150509250925092565b60008060408385031215610a0f57600080fd5b6000610a1d85828601610945565b9250506020610a2e8582860161096f565b9150509250929050565b600060208284031215610a4a57600080fd5b6000610a588482850161095a565b91505092915050565b600060208284031215610a7357600080fd5b6000610a818482850161096f565b91505092915050565b610a9381610b31565b82525050565b610aa281610b43565b82525050565b610ab181610b6f565b82525050565b6000602082019050610acc6000830184610a8a565b92915050565b6000604082019050610ae76000830185610a8a565b610af46020830184610aa8565b9392505050565b6000602082019050610b106000830184610a99565b92915050565b6000602082019050610b2b6000830184610aa8565b92915050565b6000610b3c82610b4f565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610b8281610b31565b8114610b8d57600080fd5b50565b610b9981610b43565b8114610ba457600080fd5b50565b610bb081610b6f565b8114610bbb57600080fd5b5056fea26469706673582212204fb2732648146695b2406547b76ac1d3ea84e616b93755d2f3485d154183632664736f6c63430008040033"

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, m *big.Int, c common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultTrackerBin), backend, m, c)
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

// TransferNotionalCalled is a free data retrieval call binding the contract method 0x91b9bf17.
//
// Solidity: function transferNotionalCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerCaller) TransferNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "transferNotionalCalled", arg0)

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

// TransferNotionalCalled is a free data retrieval call binding the contract method 0x91b9bf17.
//
// Solidity: function transferNotionalCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerSession) TransferNotionalCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _VaultTracker.Contract.TransferNotionalCalled(&_VaultTracker.CallOpts, arg0)
}

// TransferNotionalCalled is a free data retrieval call binding the contract method 0x91b9bf17.
//
// Solidity: function transferNotionalCalled(address ) view returns(address to, uint256 amount)
func (_VaultTracker *VaultTrackerCallerSession) TransferNotionalCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _VaultTracker.Contract.TransferNotionalCalled(&_VaultTracker.CallOpts, arg0)
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
// Solidity: function addNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerTransactor) AddNotionalReturns(opts *bind.TransactOpts, n bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "addNotionalReturns", n)
}

// AddNotionalReturns is a paid mutator transaction binding the contract method 0x5c70b7c1.
//
// Solidity: function addNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerSession) AddNotionalReturns(n bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotionalReturns(&_VaultTracker.TransactOpts, n)
}

// AddNotionalReturns is a paid mutator transaction binding the contract method 0x5c70b7c1.
//
// Solidity: function addNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerTransactorSession) AddNotionalReturns(n bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotionalReturns(&_VaultTracker.TransactOpts, n)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool m) returns()
func (_VaultTracker *VaultTrackerTransactor) MatureVaultReturns(opts *bind.TransactOpts, m bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "matureVaultReturns", m)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool m) returns()
func (_VaultTracker *VaultTrackerSession) MatureVaultReturns(m bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVaultReturns(&_VaultTracker.TransactOpts, m)
}

// MatureVaultReturns is a paid mutator transaction binding the contract method 0xda3de9e9.
//
// Solidity: function matureVaultReturns(bool m) returns()
func (_VaultTracker *VaultTrackerTransactorSession) MatureVaultReturns(m bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVaultReturns(&_VaultTracker.TransactOpts, m)
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
// Solidity: function removeNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerTransactor) RemoveNotionalReturns(opts *bind.TransactOpts, n bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "removeNotionalReturns", n)
}

// RemoveNotionalReturns is a paid mutator transaction binding the contract method 0x5dfe12ac.
//
// Solidity: function removeNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerSession) RemoveNotionalReturns(n bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotionalReturns(&_VaultTracker.TransactOpts, n)
}

// RemoveNotionalReturns is a paid mutator transaction binding the contract method 0x5dfe12ac.
//
// Solidity: function removeNotionalReturns(bool n) returns()
func (_VaultTracker *VaultTrackerTransactorSession) RemoveNotionalReturns(n bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotionalReturns(&_VaultTracker.TransactOpts, n)
}

// TransferNotional is a paid mutator transaction binding the contract method 0x7ee01a1c.
//
// Solidity: function transferNotional(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotional(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotional", f, t, a)
}

// TransferNotional is a paid mutator transaction binding the contract method 0x7ee01a1c.
//
// Solidity: function transferNotional(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotional(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotional(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotional is a paid mutator transaction binding the contract method 0x7ee01a1c.
//
// Solidity: function transferNotional(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotional(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotional(&_VaultTracker.TransactOpts, f, t, a)
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

// TransferNotionalReturns is a paid mutator transaction binding the contract method 0xfb86ed76.
//
// Solidity: function transferNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalReturns", b)
}

// TransferNotionalReturns is a paid mutator transaction binding the contract method 0xfb86ed76.
//
// Solidity: function transferNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerSession) TransferNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalReturns(&_VaultTracker.TransactOpts, b)
}

// TransferNotionalReturns is a paid mutator transaction binding the contract method 0xfb86ed76.
//
// Solidity: function transferNotionalReturns(bool b) returns()
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalReturns(b bool) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalReturns(&_VaultTracker.TransactOpts, b)
}
