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

// HashOrder is an auto generated low-level Go binding around an user-defined struct.
type HashOrder struct {
	Key        [32]byte
	Maker      common.Address
	Underlying common.Address
	Floating   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Expiry     *big.Int
	Nonce      *big.Int
}

// HashTestABI is the input ABI used to generate the binding from.
const HashTestABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HashTestBin is the compiled bytecode used for deploying new contracts.
var HashTestBin = "0x608060405234801561001057600080fd5b50610717806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634c05b2a114610046578063af1ef90b14610076578063b7dc6a8d146100a6575b600080fd5b610060600480360381019061005b9190610399565b6100d6565b60405161006d91906104ce565b60405180910390f35b610090600480360381019061008b91906103d5565b6100ea565b60405161009d91906104ce565b60405180910390f35b6100c060048036038101906100bb9190610468565b610102565b6040516100cd91906104ce565b60405180910390f35b60006100e28383610114565b905092915050565b60006100f885858585610155565b9050949350505050565b600061010d826101b4565b9050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007f982d366ee870e6c64d27e7b149dff6bf737fd1c909c2392b1b6dda92d31a24e260001b82600001358360200160208101906101f29190610347565b8460400160208101906102059190610347565b8560600160208101906102189190610370565b86608001358760a001358860c001358960e001358a61010001358b61012001356040516020016102529b9a999897969594939291906104e9565b604051602081830303815290604052805190602001209050919050565b600061028261027d846105c5565b610594565b90508281526020810184848401111561029a57600080fd5b6102a5848285610647565b509392505050565b6000813590506102bc81610685565b92915050565b6000813590506102d18161069c565b92915050565b6000813590506102e6816106b3565b92915050565b600082601f8301126102fd57600080fd5b813561030d84826020860161026f565b91505092915050565b6000610140828403121561032957600080fd5b81905092915050565b600081359050610341816106ca565b92915050565b60006020828403121561035957600080fd5b6000610367848285016102ad565b91505092915050565b60006020828403121561038257600080fd5b6000610390848285016102c2565b91505092915050565b600080604083850312156103ac57600080fd5b60006103ba858286016102d7565b92505060206103cb858286016102d7565b9150509250929050565b600080600080608085870312156103eb57600080fd5b600085013567ffffffffffffffff81111561040557600080fd5b610411878288016102ec565b945050602085013567ffffffffffffffff81111561042e57600080fd5b61043a878288016102ec565b935050604061044b87828801610332565b925050606061045c878288016102ad565b91505092959194509250565b6000610140828403121561047b57600080fd5b600061048984828501610316565b91505092915050565b61049b816105f5565b82525050565b6104aa81610607565b82525050565b6104b981610613565b82525050565b6104c88161063d565b82525050565b60006020820190506104e360008301846104b0565b92915050565b6000610160820190506104ff600083018e6104b0565b61050c602083018d6104b0565b610519604083018c610492565b610526606083018b610492565b610533608083018a6104a1565b61054060a08301896104bf565b61054d60c08301886104bf565b61055a60e08301876104bf565b6105686101008301866104bf565b6105766101208301856104bf565b6105846101408301846104bf565b9c9b505050505050505050505050565b6000604051905081810181811067ffffffffffffffff821117156105bb576105ba610656565b5b8060405250919050565b600067ffffffffffffffff8211156105e0576105df610656565b5b601f19601f8301169050602081019050919050565b60006106008261061d565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61068e816105f5565b811461069957600080fd5b50565b6106a581610607565b81146106b057600080fd5b50565b6106bc81610613565b81146106c757600080fd5b50565b6106d38161063d565b81146106de57600080fd5b5056fea26469706673582212208b6684e5a9717a3f5ed98279e34a28476896afe9fd8ee8fbfcbc091be72b869e64736f6c63430008000033"

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

// OrderTest is a free data retrieval call binding the contract method 0xb7dc6a8d.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashTest *HashTestCaller) OrderTest(opts *bind.CallOpts, o HashOrder) ([32]byte, error) {
	var out []interface{}
	err := _HashTest.contract.Call(opts, &out, "orderTest", o)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTest is a free data retrieval call binding the contract method 0xb7dc6a8d.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashTest *HashTestSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashTest.Contract.OrderTest(&_HashTest.CallOpts, o)
}

// OrderTest is a free data retrieval call binding the contract method 0xb7dc6a8d.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashTest *HashTestCallerSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashTest.Contract.OrderTest(&_HashTest.CallOpts, o)
}
