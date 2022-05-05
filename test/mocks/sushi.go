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

// SushiABI is the input ABI used to generate the binding from.
const SushiABI = "[{\"inputs\":[],\"name\":\"deadlineCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outMinimumCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pathCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"o\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"p\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"}],\"name\":\"swapExactTokensForTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SushiBin is the compiled bytecode used for deploying new contracts.
var SushiBin = "0x608060405234801561001057600080fd5b506108ae806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638678f4e71161005b5780638678f4e7146100ee578063b2f2260b1461010c578063cdd4255f14610128578063e5ee3924146101465761007d565b806338ed1739146100825780634cf1cca7146100b25780636a4f276f146100d0575b600080fd5b61009c600480360381019061009791906104e5565b610176565b6040516100a9919061063d565b60405180910390f35b6100ba61023d565b6040516100c7919061066e565b60405180910390f35b6100d8610243565b6040516100e5919061066e565b60405180910390f35b6100f6610249565b604051610103919061066e565b60405180910390f35b610126600480360381019061012191906107d8565b61024f565b005b610130610269565b60405161013d9190610830565b60405180910390f35b610160600480360381019061015b919061084b565b61028f565b60405161016d9190610830565b60405180910390f35b606086600181905550856002819055508484600391906101979291906102ce565b5082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600581905550600080548060200260200160405190810160405280929190818152602001828054801561022c57602002820191906000526020600020905b815481526020019060010190808311610218575b505050505090509695505050505050565b60055481565b60015481565b60025481565b806000908051906020019061026592919061036e565b5050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003818154811061029f57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b82805482825590600052602060002090810192821561035d579160200282015b8281111561035c57823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906102ee565b5b50905061036a91906103bb565b5090565b8280548282559060005260206000209081019282156103aa579160200282015b828111156103a957825182559160200191906001019061038e565b5b5090506103b791906103bb565b5090565b5b808211156103d45760008160009055506001016103bc565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6103ff816103ec565b811461040a57600080fd5b50565b60008135905061041c816103f6565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261044757610446610422565b5b8235905067ffffffffffffffff81111561046457610463610427565b5b6020830191508360208202830111156104805761047f61042c565b5b9250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104b282610487565b9050919050565b6104c2816104a7565b81146104cd57600080fd5b50565b6000813590506104df816104b9565b92915050565b60008060008060008060a08789031215610502576105016103e2565b5b600061051089828a0161040d565b965050602061052189828a0161040d565b955050604087013567ffffffffffffffff811115610542576105416103e7565b5b61054e89828a01610431565b9450945050606061056189828a016104d0565b925050608061057289828a0161040d565b9150509295509295509295565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6105b4816103ec565b82525050565b60006105c683836105ab565b60208301905092915050565b6000602082019050919050565b60006105ea8261057f565b6105f4818561058a565b93506105ff8361059b565b8060005b8381101561063057815161061788826105ba565b9750610622836105d2565b925050600181019050610603565b5085935050505092915050565b6000602082019050818103600083015261065781846105df565b905092915050565b610668816103ec565b82525050565b6000602082019050610683600083018461065f565b92915050565b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106d282610689565b810181811067ffffffffffffffff821117156106f1576106f061069a565b5b80604052505050565b60006107046103d8565b905061071082826106c9565b919050565b600067ffffffffffffffff8211156107305761072f61069a565b5b602082029050602081019050919050565b600061075461074f84610715565b6106fa565b905080838252602082019050602084028301858111156107775761077661042c565b5b835b818110156107a0578061078c888261040d565b845260208401935050602081019050610779565b5050509392505050565b600082601f8301126107bf576107be610422565b5b81356107cf848260208601610741565b91505092915050565b6000602082840312156107ee576107ed6103e2565b5b600082013567ffffffffffffffff81111561080c5761080b6103e7565b5b610818848285016107aa565b91505092915050565b61082a816104a7565b82525050565b60006020820190506108456000830184610821565b92915050565b600060208284031215610861576108606103e2565b5b600061086f8482850161040d565b9150509291505056fea26469706673582212208907d3b2d59892d4cf50489d025da09fd611a1a1c82d1f59dc336875f0ee274c64736f6c634300080d0033"

// DeploySushi deploys a new Ethereum contract, binding an instance of Sushi to it.
func DeploySushi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sushi, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SushiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sushi{SushiCaller: SushiCaller{contract: contract}, SushiTransactor: SushiTransactor{contract: contract}, SushiFilterer: SushiFilterer{contract: contract}}, nil
}

// Sushi is an auto generated Go binding around an Ethereum contract.
type Sushi struct {
	SushiCaller     // Read-only binding to the contract
	SushiTransactor // Write-only binding to the contract
	SushiFilterer   // Log filterer for contract events
}

// SushiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiSession struct {
	Contract     *Sushi            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiCallerSession struct {
	Contract *SushiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SushiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiTransactorSession struct {
	Contract     *SushiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiRaw struct {
	Contract *Sushi // Generic contract binding to access the raw methods on
}

// SushiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiCallerRaw struct {
	Contract *SushiCaller // Generic read-only contract binding to access the raw methods on
}

// SushiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiTransactorRaw struct {
	Contract *SushiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushi creates a new instance of Sushi, bound to a specific deployed contract.
func NewSushi(address common.Address, backend bind.ContractBackend) (*Sushi, error) {
	contract, err := bindSushi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sushi{SushiCaller: SushiCaller{contract: contract}, SushiTransactor: SushiTransactor{contract: contract}, SushiFilterer: SushiFilterer{contract: contract}}, nil
}

// NewSushiCaller creates a new read-only instance of Sushi, bound to a specific deployed contract.
func NewSushiCaller(address common.Address, caller bind.ContractCaller) (*SushiCaller, error) {
	contract, err := bindSushi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiCaller{contract: contract}, nil
}

// NewSushiTransactor creates a new write-only instance of Sushi, bound to a specific deployed contract.
func NewSushiTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiTransactor, error) {
	contract, err := bindSushi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiTransactor{contract: contract}, nil
}

// NewSushiFilterer creates a new log filterer instance of Sushi, bound to a specific deployed contract.
func NewSushiFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiFilterer, error) {
	contract, err := bindSushi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiFilterer{contract: contract}, nil
}

// bindSushi binds a generic wrapper to an already deployed contract.
func bindSushi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushi *SushiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushi.Contract.SushiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushi *SushiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushi.Contract.SushiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushi *SushiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushi.Contract.SushiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushi *SushiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushi *SushiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushi *SushiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushi.Contract.contract.Transact(opts, method, params...)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Sushi *SushiCaller) DeadlineCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sushi.contract.Call(opts, &out, "deadlineCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Sushi *SushiSession) DeadlineCalled() (*big.Int, error) {
	return _Sushi.Contract.DeadlineCalled(&_Sushi.CallOpts)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Sushi *SushiCallerSession) DeadlineCalled() (*big.Int, error) {
	return _Sushi.Contract.DeadlineCalled(&_Sushi.CallOpts)
}

// InCalled is a free data retrieval call binding the contract method 0x6a4f276f.
//
// Solidity: function inCalled() view returns(uint256)
func (_Sushi *SushiCaller) InCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sushi.contract.Call(opts, &out, "inCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InCalled is a free data retrieval call binding the contract method 0x6a4f276f.
//
// Solidity: function inCalled() view returns(uint256)
func (_Sushi *SushiSession) InCalled() (*big.Int, error) {
	return _Sushi.Contract.InCalled(&_Sushi.CallOpts)
}

// InCalled is a free data retrieval call binding the contract method 0x6a4f276f.
//
// Solidity: function inCalled() view returns(uint256)
func (_Sushi *SushiCallerSession) InCalled() (*big.Int, error) {
	return _Sushi.Contract.InCalled(&_Sushi.CallOpts)
}

// OutMinimumCalled is a free data retrieval call binding the contract method 0x8678f4e7.
//
// Solidity: function outMinimumCalled() view returns(uint256)
func (_Sushi *SushiCaller) OutMinimumCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sushi.contract.Call(opts, &out, "outMinimumCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutMinimumCalled is a free data retrieval call binding the contract method 0x8678f4e7.
//
// Solidity: function outMinimumCalled() view returns(uint256)
func (_Sushi *SushiSession) OutMinimumCalled() (*big.Int, error) {
	return _Sushi.Contract.OutMinimumCalled(&_Sushi.CallOpts)
}

// OutMinimumCalled is a free data retrieval call binding the contract method 0x8678f4e7.
//
// Solidity: function outMinimumCalled() view returns(uint256)
func (_Sushi *SushiCallerSession) OutMinimumCalled() (*big.Int, error) {
	return _Sushi.Contract.OutMinimumCalled(&_Sushi.CallOpts)
}

// PathCalled is a free data retrieval call binding the contract method 0xe5ee3924.
//
// Solidity: function pathCalled(uint256 ) view returns(address)
func (_Sushi *SushiCaller) PathCalled(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sushi.contract.Call(opts, &out, "pathCalled", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PathCalled is a free data retrieval call binding the contract method 0xe5ee3924.
//
// Solidity: function pathCalled(uint256 ) view returns(address)
func (_Sushi *SushiSession) PathCalled(arg0 *big.Int) (common.Address, error) {
	return _Sushi.Contract.PathCalled(&_Sushi.CallOpts, arg0)
}

// PathCalled is a free data retrieval call binding the contract method 0xe5ee3924.
//
// Solidity: function pathCalled(uint256 ) view returns(address)
func (_Sushi *SushiCallerSession) PathCalled(arg0 *big.Int) (common.Address, error) {
	return _Sushi.Contract.PathCalled(&_Sushi.CallOpts, arg0)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_Sushi *SushiCaller) ToCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sushi.contract.Call(opts, &out, "toCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_Sushi *SushiSession) ToCalled() (common.Address, error) {
	return _Sushi.Contract.ToCalled(&_Sushi.CallOpts)
}

// ToCalled is a free data retrieval call binding the contract method 0xcdd4255f.
//
// Solidity: function toCalled() view returns(address)
func (_Sushi *SushiCallerSession) ToCalled() (common.Address, error) {
	return _Sushi.Contract.ToCalled(&_Sushi.CallOpts)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.contract.Transact(opts, "swapExactTokensForTokens", i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokens(&_Sushi.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiTransactorSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokens(&_Sushi.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiTransactor) SwapExactTokensForTokensReturns(opts *bind.TransactOpts, r []*big.Int) (*types.Transaction, error) {
	return _Sushi.contract.Transact(opts, "swapExactTokensForTokensReturns", r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokensReturns(&_Sushi.TransactOpts, r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiTransactorSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokensReturns(&_Sushi.TransactOpts, r)
}
