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

// HashTestABI is the input ABI used to generate the binding from.
const HashTestABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"},{\"internalType\":\"uint256[6]\",\"name\":\"p\",\"type\":\"uint256[6]\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HashTestBin is the compiled bytecode used for deploying new contracts.
var HashTestBin = "0x608060405234801561001057600080fd5b50610990806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634c05b2a1146100465780638da4239e14610076578063af1ef90b146100a6575b600080fd5b610060600480360381019061005b919061063c565b6100d6565b60405161006d9190610747565b60405180910390f35b610090600480360381019061008b91906105c4565b6100ea565b60405161009d9190610747565b60405180910390f35b6100c060048036038101906100bb9190610678565b610104565b6040516100cd9190610747565b60405180910390f35b60006100e2838361011c565b905092915050565b60006100f9868686868661015d565b905095945050505050565b60006101128585858561048b565b9050949350505050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60007f982d366ee870e6c64d27e7b149dff6bf737fd1c909c2392b1b6dda92d31a24e260001b8686868686600060058111156101c2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106101f9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358760016005811115610239577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610270577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002013588600260058111156102b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106102e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358960036005811115610327577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061035e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358a6004600581111561039e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106103d5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358b600580811115610414577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061044b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002013560405160200161046a9b9a99989796959493929190610762565b60405160208183030381529060405280519060200120905095945050505050565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60006104fd6104f88461083e565b61080d565b90508281526020810184848401111561051557600080fd5b6105208482856108c0565b509392505050565b600081359050610537816108fe565b92915050565b60008190508260206006028201111561055557600080fd5b92915050565b60008135905061056a81610915565b92915050565b60008135905061057f8161092c565b92915050565b600082601f83011261059657600080fd5b81356105a68482602086016104ea565b91505092915050565b6000813590506105be81610943565b92915050565b600080600080600061014086880312156105dd57600080fd5b60006105eb88828901610570565b95505060206105fc88828901610528565b945050604061060d88828901610528565b935050606061061e8882890161055b565b925050608061062f8882890161053d565b9150509295509295909350565b6000806040838503121561064f57600080fd5b600061065d85828601610570565b925050602061066e85828601610570565b9150509250929050565b6000806000806080858703121561068e57600080fd5b600085013567ffffffffffffffff8111156106a857600080fd5b6106b487828801610585565b945050602085013567ffffffffffffffff8111156106d157600080fd5b6106dd87828801610585565b93505060406106ee878288016105af565b92505060606106ff87828801610528565b91505092959194509250565b6107148161086e565b82525050565b61072381610880565b82525050565b6107328161088c565b82525050565b610741816108b6565b82525050565b600060208201905061075c6000830184610729565b92915050565b600061016082019050610778600083018e610729565b610785602083018d610729565b610792604083018c61070b565b61079f606083018b61070b565b6107ac608083018a61071a565b6107b960a0830189610738565b6107c660c0830188610738565b6107d360e0830187610738565b6107e1610100830186610738565b6107ef610120830185610738565b6107fd610140830184610738565b9c9b505050505050505050505050565b6000604051905081810181811067ffffffffffffffff82111715610834576108336108cf565b5b8060405250919050565b600067ffffffffffffffff821115610859576108586108cf565b5b601f19601f8301169050602081019050919050565b600061087982610896565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6109078161086e565b811461091257600080fd5b50565b61091e81610880565b811461092957600080fd5b50565b6109358161088c565b811461094057600080fd5b50565b61094c816108b6565b811461095757600080fd5b5056fea2646970667358221220fbfd30265182e0c0655d88fc4ae6862fd3938f9e876b1680a79248362429fdc264736f6c63430008000033"

// DeployHashTest deploys a new Ethereum contract, binding an instance of HashTest to it.
func DeployHashTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashTest, error) {
	parsed, err := abi.JSON(strings.NewReader(HashTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashTest{HashTestCaller: HashTestCaller{contract: contract}, HashTestTransactor: HashTestTransactor{contract: contract}, HashTestFilterer: HashTestFilterer{contract: contract}}, nil
}

// HashTest is an auto generated Go binding around an Ethereum contract.
type HashTest struct {
	HashTestCaller     // Read-only binding to the contract
	HashTestTransactor // Write-only binding to the contract
	HashTestFilterer   // Log filterer for contract events
}

// HashTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashTestSession struct {
	Contract     *HashTest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashTestCallerSession struct {
	Contract *HashTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HashTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashTestTransactorSession struct {
	Contract     *HashTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HashTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashTestRaw struct {
	Contract *HashTest // Generic contract binding to access the raw methods on
}

// HashTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashTestCallerRaw struct {
	Contract *HashTestCaller // Generic read-only contract binding to access the raw methods on
}

// HashTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashTestTransactorRaw struct {
	Contract *HashTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashTest creates a new instance of HashTest, bound to a specific deployed contract.
func NewHashTest(address common.Address, backend bind.ContractBackend) (*HashTest, error) {
	contract, err := bindHashTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashTest{HashTestCaller: HashTestCaller{contract: contract}, HashTestTransactor: HashTestTransactor{contract: contract}, HashTestFilterer: HashTestFilterer{contract: contract}}, nil
}

// NewHashTestCaller creates a new read-only instance of HashTest, bound to a specific deployed contract.
func NewHashTestCaller(address common.Address, caller bind.ContractCaller) (*HashTestCaller, error) {
	contract, err := bindHashTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashTestCaller{contract: contract}, nil
}

// NewHashTestTransactor creates a new write-only instance of HashTest, bound to a specific deployed contract.
func NewHashTestTransactor(address common.Address, transactor bind.ContractTransactor) (*HashTestTransactor, error) {
	contract, err := bindHashTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashTestTransactor{contract: contract}, nil
}

// NewHashTestFilterer creates a new log filterer instance of HashTest, bound to a specific deployed contract.
func NewHashTestFilterer(address common.Address, filterer bind.ContractFilterer) (*HashTestFilterer, error) {
	contract, err := bindHashTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashTestFilterer{contract: contract}, nil
}

// bindHashTest binds a generic wrapper to an already deployed contract.
func bindHashTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashTest *HashTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashTest.Contract.HashTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashTest *HashTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashTest.Contract.HashTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashTest *HashTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashTest.Contract.HashTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashTest *HashTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashTest *HashTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashTest *HashTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashTest.Contract.contract.Transact(opts, method, params...)
}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashTest *HashTestCaller) DomainTest(opts *bind.CallOpts, n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	var out []interface{}
	err := _HashTest.contract.Call(opts, &out, "domainTest", n, version, c, verifier)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashTest *HashTestSession) DomainTest(n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	return _HashTest.Contract.DomainTest(&_HashTest.CallOpts, n, version, c, verifier)
}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashTest *HashTestCallerSession) DomainTest(n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	return _HashTest.Contract.DomainTest(&_HashTest.CallOpts, n, version, c, verifier)
}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashTest *HashTestCaller) MessageTest(opts *bind.CallOpts, d [32]byte, h [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HashTest.contract.Call(opts, &out, "messageTest", d, h)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashTest *HashTestSession) MessageTest(d [32]byte, h [32]byte) ([32]byte, error) {
	return _HashTest.Contract.MessageTest(&_HashTest.CallOpts, d, h)
}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashTest *HashTestCallerSession) MessageTest(d [32]byte, h [32]byte) ([32]byte, error) {
	return _HashTest.Contract.MessageTest(&_HashTest.CallOpts, d, h)
}

// OrderTest is a free data retrieval call binding the contract method 0x8da4239e.
//
// Solidity: function orderTest(bytes32 k, address m, address u, bool f, uint256[6] p) pure returns(bytes32)
func (_HashTest *HashTestCaller) OrderTest(opts *bind.CallOpts, k [32]byte, m common.Address, u common.Address, f bool, p [6]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _HashTest.contract.Call(opts, &out, "orderTest", k, m, u, f, p)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTest is a free data retrieval call binding the contract method 0x8da4239e.
//
// Solidity: function orderTest(bytes32 k, address m, address u, bool f, uint256[6] p) pure returns(bytes32)
func (_HashTest *HashTestSession) OrderTest(k [32]byte, m common.Address, u common.Address, f bool, p [6]*big.Int) ([32]byte, error) {
	return _HashTest.Contract.OrderTest(&_HashTest.CallOpts, k, m, u, f, p)
}

// OrderTest is a free data retrieval call binding the contract method 0x8da4239e.
//
// Solidity: function orderTest(bytes32 k, address m, address u, bool f, uint256[6] p) pure returns(bytes32)
func (_HashTest *HashTestCallerSession) OrderTest(k [32]byte, m common.Address, u common.Address, f bool, p [6]*big.Int) ([32]byte, error) {
	return _HashTest.Contract.OrderTest(&_HashTest.CallOpts, k, m, u, f, p)
}
