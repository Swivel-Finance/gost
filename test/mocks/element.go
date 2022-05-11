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

// FundManagement is an auto generated low-level Go binding around an user-defined struct.
type FundManagement struct {
	Sender              common.Address
	Recipient           common.Address
	FromInternalBalance bool
	ToInternalBalance   bool
}

// SingleSwap is an auto generated low-level Go binding around an user-defined struct.
type SingleSwap struct {
	UserData []byte
	PoolId   [32]byte
	Amount   *big.Int
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
}

// ElementABI is the input ABI used to generate the binding from.
const ElementABI = "[{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"deadlineReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundManagementSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"fundManagementSenderReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"returnReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"return_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"singleSwapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"singleSwapAmountReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"f\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapCalled\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"fundManagement\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"swap\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_return\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ElementBin is the compiled bytecode used for deploying new contracts.
var ElementBin = "0x608060405234801561001057600080fd5b50611120806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80635f65334c116100665780635f65334c1461014a57806369bbd8cd1461016657806384ea21c81461019657806391d51144146101b2578063a4b15a5d146101d05761009e565b8063127f0f26146100a357806328386b83146100bf57806329dcb0cf146100dd5780633b09e714146100fb5780635658724f14610117575b600080fd5b6100bd60048036038101906100b89190610895565b6101ee565b005b6100c76101f8565b6040516100d49190610903565b60405180910390f35b6100e5610222565b6040516100f2919061092d565b60405180910390f35b61011560048036038101906101109190610895565b61022c565b005b610131600480360381019061012c9190610974565b610236565b6040516101419493929190610c61565b60405180910390f35b610164600480360381019061015f9190610974565b6104e9565b005b610180600480360381019061017b9190611007565b61052d565b60405161018d919061092d565b60405180910390f35b6101b060048036038101906101ab9190610895565b61078a565b005b6101ba610794565b6040516101c7919061092d565b60405180910390f35b6101d861079e565b6040516101e5919061092d565b60405180910390f35b8060028190555050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b8060048190555050565b6005602052806000526040600020600091509050806000016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900460ff161515151581526020016001820160159054906101000a900460ff16151515158152505090806002016040518060c001604052908160008201805461035b906110b9565b80601f0160208091040260200160405190810160405280929190818152602001828054610387906110b9565b80156103d45780601f106103a9576101008083540402835291602001916103d4565b820191906000526020600020905b8154815290600101906020018083116103b757829003601f168201915b5050505050815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff16600181111561041657610415610b02565b5b600181111561042857610427610b02565b5b81526020016003820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050908060070154908060080154905084565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060405180608001604052808581526020018681526020018481526020018381525060056000866000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff02191690831515021790555060608201518160010160156101000a81548160ff021916908315150217905550505060208201518160020160008201518160000190805190602001906106919291906107a8565b50602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908360018111156106d1576106d0610b02565b5b021790555060808201518160030160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505060408201518160070155606082015181600801559050506000549050949350505050565b8060018190555050565b6000600254905090565b6000600454905090565b8280546107b4906110b9565b90600052602060002090601f0160209004810192826107d6576000855561081d565b82601f106107ef57805160ff191683800117855561081d565b8280016001018555821561081d579182015b8281111561081c578251825591602001919060010190610801565b5b50905061082a919061082e565b5090565b5b8082111561084757600081600090555060010161082f565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6108728161085f565b811461087d57600080fd5b50565b60008135905061088f81610869565b92915050565b6000602082840312156108ab576108aa610855565b5b60006108b984828501610880565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108ed826108c2565b9050919050565b6108fd816108e2565b82525050565b600060208201905061091860008301846108f4565b92915050565b6109278161085f565b82525050565b6000602082019050610942600083018461091e565b92915050565b610951816108e2565b811461095c57600080fd5b50565b60008135905061096e81610948565b92915050565b60006020828403121561098a57610989610855565b5b60006109988482850161095f565b91505092915050565b6109aa816108e2565b82525050565b60006109bb826108c2565b9050919050565b6109cb816109b0565b82525050565b60008115159050919050565b6109e6816109d1565b82525050565b608082016000820151610a0260008501826109a1565b506020820151610a1560208501826109c2565b506040820151610a2860408501826109dd565b506060820151610a3b60608501826109dd565b50505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a7b578082015181840152602081019050610a60565b83811115610a8a576000848401525b50505050565b6000601f19601f8301169050919050565b6000610aac82610a41565b610ab68185610a4c565b9350610ac6818560208601610a5d565b610acf81610a90565b840191505092915050565b6000819050919050565b610aed81610ada565b82525050565b610afc8161085f565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028110610b4257610b41610b02565b5b50565b6000819050610b5382610b31565b919050565b6000610b6382610b45565b9050919050565b610b7381610b58565b82525050565b6000819050919050565b6000610b9e610b99610b94846108c2565b610b79565b6108c2565b9050919050565b6000610bb082610b83565b9050919050565b6000610bc282610ba5565b9050919050565b610bd281610bb7565b82525050565b600060c0830160008301518482036000860152610bf58282610aa1565b9150506020830151610c0a6020860182610ae4565b506040830151610c1d6040860182610af3565b506060830151610c306060860182610b6a565b506080830151610c436080860182610bc9565b5060a0830151610c5660a0860182610bc9565b508091505092915050565b600060e082019050610c7660008301876109ec565b8181036080830152610c888186610bd8565b9050610c9760a083018561091e565b610ca460c083018461091e565b95945050505050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610cea82610a90565b810181811067ffffffffffffffff82111715610d0957610d08610cb2565b5b80604052505050565b6000610d1c61084b565b9050610d288282610ce1565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610d5757610d56610cb2565b5b610d6082610a90565b9050602081019050919050565b82818337600083830152505050565b6000610d8f610d8a84610d3c565b610d12565b905082815260208101848484011115610dab57610daa610d37565b5b610db6848285610d6d565b509392505050565b600082601f830112610dd357610dd2610d32565b5b8135610de3848260208601610d7c565b91505092915050565b610df581610ada565b8114610e0057600080fd5b50565b600081359050610e1281610dec565b92915050565b60028110610e2557600080fd5b50565b600081359050610e3781610e18565b92915050565b6000610e48826108e2565b9050919050565b610e5881610e3d565b8114610e6357600080fd5b50565b600081359050610e7581610e4f565b92915050565b600060c08284031215610e9157610e90610cad565b5b610e9b60c0610d12565b9050600082013567ffffffffffffffff811115610ebb57610eba610d2d565b5b610ec784828501610dbe565b6000830152506020610edb84828501610e03565b6020830152506040610eef84828501610880565b6040830152506060610f0384828501610e28565b6060830152506080610f1784828501610e66565b60808301525060a0610f2b84828501610e66565b60a08301525092915050565b610f40816109b0565b8114610f4b57600080fd5b50565b600081359050610f5d81610f37565b92915050565b610f6c816109d1565b8114610f7757600080fd5b50565b600081359050610f8981610f63565b92915050565b600060808284031215610fa557610fa4610cad565b5b610faf6080610d12565b90506000610fbf8482850161095f565b6000830152506020610fd384828501610f4e565b6020830152506040610fe784828501610f7a565b6040830152506060610ffb84828501610f7a565b60608301525092915050565b60008060008060e0858703121561102157611020610855565b5b600085013567ffffffffffffffff81111561103f5761103e61085a565b5b61104b87828801610e7b565b945050602061105c87828801610f8f565b93505060a061106d87828801610880565b92505060c061107e87828801610880565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806110d157607f821691505b6020821081036110e4576110e361108a565b5b5091905056fea2646970667358221220cd76f6a3eca237bcb2d3d306674c5658fd8554415e5d0243aa66cdb92c15172264736f6c634300080d0033"

// DeployElement deploys a new Ethereum contract, binding an instance of Element to it.
func DeployElement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Element, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Element{ElementCaller: ElementCaller{contract: contract}, ElementTransactor: ElementTransactor{contract: contract}, ElementFilterer: ElementFilterer{contract: contract}}, nil
}

// Element is an auto generated Go binding around an Ethereum contract.
type Element struct {
	ElementCaller     // Read-only binding to the contract
	ElementTransactor // Write-only binding to the contract
	ElementFilterer   // Log filterer for contract events
}

// ElementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElementSession struct {
	Contract     *Element          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElementCallerSession struct {
	Contract *ElementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ElementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElementTransactorSession struct {
	Contract     *ElementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ElementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElementRaw struct {
	Contract *Element // Generic contract binding to access the raw methods on
}

// ElementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElementCallerRaw struct {
	Contract *ElementCaller // Generic read-only contract binding to access the raw methods on
}

// ElementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElementTransactorRaw struct {
	Contract *ElementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElement creates a new instance of Element, bound to a specific deployed contract.
func NewElement(address common.Address, backend bind.ContractBackend) (*Element, error) {
	contract, err := bindElement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Element{ElementCaller: ElementCaller{contract: contract}, ElementTransactor: ElementTransactor{contract: contract}, ElementFilterer: ElementFilterer{contract: contract}}, nil
}

// NewElementCaller creates a new read-only instance of Element, bound to a specific deployed contract.
func NewElementCaller(address common.Address, caller bind.ContractCaller) (*ElementCaller, error) {
	contract, err := bindElement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElementCaller{contract: contract}, nil
}

// NewElementTransactor creates a new write-only instance of Element, bound to a specific deployed contract.
func NewElementTransactor(address common.Address, transactor bind.ContractTransactor) (*ElementTransactor, error) {
	contract, err := bindElement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTransactor{contract: contract}, nil
}

// NewElementFilterer creates a new log filterer instance of Element, bound to a specific deployed contract.
func NewElementFilterer(address common.Address, filterer bind.ContractFilterer) (*ElementFilterer, error) {
	contract, err := bindElement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElementFilterer{contract: contract}, nil
}

// bindElement binds a generic wrapper to an already deployed contract.
func bindElement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Element *ElementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Element.Contract.ElementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Element *ElementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Element.Contract.ElementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Element *ElementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Element.Contract.ElementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Element *ElementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Element.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Element *ElementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Element.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Element *ElementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Element.Contract.contract.Transact(opts, method, params...)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementSession) Deadline() (*big.Int, error) {
	return _Element.Contract.Deadline(&_Element.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementCallerSession) Deadline() (*big.Int, error) {
	return _Element.Contract.Deadline(&_Element.CallOpts)
}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementCaller) FundManagementSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "fundManagementSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementSession) FundManagementSender() (common.Address, error) {
	return _Element.Contract.FundManagementSender(&_Element.CallOpts)
}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementCallerSession) FundManagementSender() (common.Address, error) {
	return _Element.Contract.FundManagementSender(&_Element.CallOpts)
}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementCaller) Return(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "return_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementSession) Return() (*big.Int, error) {
	return _Element.Contract.Return(&_Element.CallOpts)
}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementCallerSession) Return() (*big.Int, error) {
	return _Element.Contract.Return(&_Element.CallOpts)
}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementCaller) SingleSwapAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "singleSwapAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementSession) SingleSwapAmount() (*big.Int, error) {
	return _Element.Contract.SingleSwapAmount(&_Element.CallOpts)
}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementCallerSession) SingleSwapAmount() (*big.Int, error) {
	return _Element.Contract.SingleSwapAmount(&_Element.CallOpts)
}

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns((address,address,bool,bool) fundManagement, (bytes,bytes32,uint256,uint8,address,address) swap, uint256 _return, uint256 deadline)
func (_Element *ElementCaller) SwapCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	FundManagement FundManagement
	Swap           SingleSwap
	Return         *big.Int
	Deadline       *big.Int
}, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "swapCalled", arg0)

	outstruct := new(struct {
		FundManagement FundManagement
		Swap           SingleSwap
		Return         *big.Int
		Deadline       *big.Int
	})

	outstruct.FundManagement = out[0].(FundManagement)
	outstruct.Swap = out[1].(SingleSwap)
	outstruct.Return = out[2].(*big.Int)
	outstruct.Deadline = out[3].(*big.Int)

	return *outstruct, err

}

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns((address,address,bool,bool) fundManagement, (bytes,bytes32,uint256,uint8,address,address) swap, uint256 _return, uint256 deadline)
func (_Element *ElementSession) SwapCalled(arg0 common.Address) (struct {
	FundManagement FundManagement
	Swap           SingleSwap
	Return         *big.Int
	Deadline       *big.Int
}, error) {
	return _Element.Contract.SwapCalled(&_Element.CallOpts, arg0)
}

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns((address,address,bool,bool) fundManagement, (bytes,bytes32,uint256,uint8,address,address) swap, uint256 _return, uint256 deadline)
func (_Element *ElementCallerSession) SwapCalled(arg0 common.Address) (struct {
	FundManagement FundManagement
	Swap           SingleSwap
	Return         *big.Int
	Deadline       *big.Int
}, error) {
	return _Element.Contract.SwapCalled(&_Element.CallOpts, arg0)
}

// DeadlineReturns is a paid mutator transaction binding the contract method 0x84ea21c8.
//
// Solidity: function deadlineReturns(uint256 d) returns()
func (_Element *ElementTransactor) DeadlineReturns(opts *bind.TransactOpts, d *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "deadlineReturns", d)
}

// DeadlineReturns is a paid mutator transaction binding the contract method 0x84ea21c8.
//
// Solidity: function deadlineReturns(uint256 d) returns()
func (_Element *ElementSession) DeadlineReturns(d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.DeadlineReturns(&_Element.TransactOpts, d)
}

// DeadlineReturns is a paid mutator transaction binding the contract method 0x84ea21c8.
//
// Solidity: function deadlineReturns(uint256 d) returns()
func (_Element *ElementTransactorSession) DeadlineReturns(d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.DeadlineReturns(&_Element.TransactOpts, d)
}

// FundManagementSenderReturns is a paid mutator transaction binding the contract method 0x5f65334c.
//
// Solidity: function fundManagementSenderReturns(address s) returns()
func (_Element *ElementTransactor) FundManagementSenderReturns(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "fundManagementSenderReturns", s)
}

// FundManagementSenderReturns is a paid mutator transaction binding the contract method 0x5f65334c.
//
// Solidity: function fundManagementSenderReturns(address s) returns()
func (_Element *ElementSession) FundManagementSenderReturns(s common.Address) (*types.Transaction, error) {
	return _Element.Contract.FundManagementSenderReturns(&_Element.TransactOpts, s)
}

// FundManagementSenderReturns is a paid mutator transaction binding the contract method 0x5f65334c.
//
// Solidity: function fundManagementSenderReturns(address s) returns()
func (_Element *ElementTransactorSession) FundManagementSenderReturns(s common.Address) (*types.Transaction, error) {
	return _Element.Contract.FundManagementSenderReturns(&_Element.TransactOpts, s)
}

// ReturnReturns is a paid mutator transaction binding the contract method 0x127f0f26.
//
// Solidity: function returnReturns(uint256 r) returns()
func (_Element *ElementTransactor) ReturnReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "returnReturns", r)
}

// ReturnReturns is a paid mutator transaction binding the contract method 0x127f0f26.
//
// Solidity: function returnReturns(uint256 r) returns()
func (_Element *ElementSession) ReturnReturns(r *big.Int) (*types.Transaction, error) {
	return _Element.Contract.ReturnReturns(&_Element.TransactOpts, r)
}

// ReturnReturns is a paid mutator transaction binding the contract method 0x127f0f26.
//
// Solidity: function returnReturns(uint256 r) returns()
func (_Element *ElementTransactorSession) ReturnReturns(r *big.Int) (*types.Transaction, error) {
	return _Element.Contract.ReturnReturns(&_Element.TransactOpts, r)
}

// SingleSwapAmountReturns is a paid mutator transaction binding the contract method 0x3b09e714.
//
// Solidity: function singleSwapAmountReturns(uint256 a) returns()
func (_Element *ElementTransactor) SingleSwapAmountReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "singleSwapAmountReturns", a)
}

// SingleSwapAmountReturns is a paid mutator transaction binding the contract method 0x3b09e714.
//
// Solidity: function singleSwapAmountReturns(uint256 a) returns()
func (_Element *ElementSession) SingleSwapAmountReturns(a *big.Int) (*types.Transaction, error) {
	return _Element.Contract.SingleSwapAmountReturns(&_Element.TransactOpts, a)
}

// SingleSwapAmountReturns is a paid mutator transaction binding the contract method 0x3b09e714.
//
// Solidity: function singleSwapAmountReturns(uint256 a) returns()
func (_Element *ElementTransactorSession) SingleSwapAmountReturns(a *big.Int) (*types.Transaction, error) {
	return _Element.Contract.SingleSwapAmountReturns(&_Element.TransactOpts, a)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementTransactor) Swap(opts *bind.TransactOpts, s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "swap", s, f, r, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementSession) Swap(s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, r, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementTransactorSession) Swap(s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, r, d)
}
