// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fakes

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
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Expiry     *big.Int
}

// HashFakeABI is the input ABI used to generate the binding from.
const HashFakeABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HashFakeBin is the compiled bytecode used for deploying new contracts.
var HashFakeBin = "0x608060405234801561001057600080fd5b50610cdf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063207d220c1461005c5780634c05b2a11461008c57806391c2600f146100bc57806397995ecb146100da578063af1ef90b146100f8575b600080fd5b61007660048036038101906100719190610504565b610128565b6040516100839190610ac0565b60405180910390f35b6100a660048036038101906100a19190610435565b61013a565b6040516100b39190610ac0565b60405180910390f35b6100c461014e565b6040516100d19190610ac0565b60405180910390f35b6100e261017a565b6040516100ef9190610ac0565b60405180910390f35b610112600480360381019061010d9190610471565b6101a6565b60405161011f9190610ac0565b60405180910390f35b6000610133826101be565b9050919050565b6000610146838361026b565b905092915050565b600060405160200161015f906109fc565b60405160208183030381529060405280519060200120905090565b600060405160200161018b90610a74565b60405160208183030381529060405280519060200120905090565b60006101b4858585856102ac565b9050949350505050565b60007f6d7f43a0d330aacab713d630fa9dad09c6e92c02743a1d7a6172d2307b612f6160001b82600001358360200160208101906101fc91906103e3565b84604001602081019061020f91906103e3565b856060016020810190610222919061040c565b86608001358760a001358860c001358960e0013560405160200161024e99989796959493929190610959565b604051602081830303815290604052805190602001209050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b600061031e61031984610b0c565b610adb565b90508281526020810184848401111561033657600080fd5b610341848285610b99565b509392505050565b60008135905061035881610c4d565b92915050565b60008135905061036d81610c64565b92915050565b60008135905061038281610c7b565b92915050565b600082601f83011261039957600080fd5b81356103a984826020860161030b565b91505092915050565b600061010082840312156103c557600080fd5b81905092915050565b6000813590506103dd81610c92565b92915050565b6000602082840312156103f557600080fd5b600061040384828501610349565b91505092915050565b60006020828403121561041e57600080fd5b600061042c8482850161035e565b91505092915050565b6000806040838503121561044857600080fd5b600061045685828601610373565b925050602061046785828601610373565b9150509250929050565b6000806000806080858703121561048757600080fd5b600085013567ffffffffffffffff8111156104a157600080fd5b6104ad87828801610388565b945050602085013567ffffffffffffffff8111156104ca57600080fd5b6104d687828801610388565b93505060406104e7878288016103ce565b92505060606104f887828801610349565b91505092959194509250565b6000610100828403121561051757600080fd5b6000610525848285016103b2565b91505092915050565b61053f61053a82610b47565b610ba8565b82525050565b61055661055182610b59565b610bba565b82525050565b61056581610b65565b82525050565b61057c61057782610b65565b610bcc565b82525050565b600061058f600c83610b3c565b91507f737472696e67206e616d652c00000000000000000000000000000000000000006000830152600c82019050919050565b60006105cf600c83610b3c565b91507f62797465733332206b65792c00000000000000000000000000000000000000006000830152600c82019050919050565b600061060f601183610b3c565b91507f75696e7432353620696e7465726573742c0000000000000000000000000000006000830152601182019050919050565b600061064f600683610b3c565b91507f4f726465722800000000000000000000000000000000000000000000000000006000830152600682019050919050565b600061068f601383610b3c565b91507f6164647265737320756e6465726c79696e672c000000000000000000000000006000830152601382019050919050565b60006106cf601983610b3c565b91507f6164647265737320766572696679696e67436f6e7472616374000000000000006000830152601982019050919050565b600061070f601083610b3c565b91507f75696e7432353620636861696e49642c000000000000000000000000000000006000830152601082019050919050565b600061074f600183610b3c565b91507f29000000000000000000000000000000000000000000000000000000000000006000830152600182019050919050565b600061078f600f83610b3c565b91507f737472696e672076657273696f6e2c00000000000000000000000000000000006000830152600f82019050919050565b60006107cf601183610b3c565b91507f75696e74323536206475726174696f6e2c0000000000000000000000000000006000830152601182019050919050565b600061080f600e83610b3c565b91507f61646472657373206d616b65722c0000000000000000000000000000000000006000830152600e82019050919050565b600061084f601283610b3c565b91507f75696e74323536207072696e636970616c2c00000000000000000000000000006000830152601282019050919050565b600061088f600f83610b3c565b91507f75696e74323536206578706972792c00000000000000000000000000000000006000830152600f82019050919050565b60006108cf600d83610b3c565b91507f454950373132446f6d61696e28000000000000000000000000000000000000006000830152600d82019050919050565b600061090f600d83610b3c565b91507f626f6f6c20666c6f6174696e67000000000000000000000000000000000000006000830152600d82019050919050565b61095361094e82610b8f565b610be8565b82525050565b6000610965828c61056b565b602082019150610975828b61056b565b602082019150610985828a61052e565b601482019150610995828961052e565b6014820191506109a58288610545565b6001820191506109b58287610942565b6020820191506109c58286610942565b6020820191506109d58285610942565b6020820191506109e58284610942565b6020820191508190509a9950505050505050505050565b6000610a0782610642565b9150610a12826105c2565b9150610a1d82610802565b9150610a2882610682565b9150610a3382610902565b9150610a3e82610842565b9150610a4982610602565b9150610a54826107c2565b9150610a5f82610882565b9150610a6a82610742565b9150819050919050565b6000610a7f826108c2565b9150610a8a82610582565b9150610a9582610782565b9150610aa082610702565b9150610aab826106c2565b9150610ab682610742565b9150819050919050565b6000602082019050610ad5600083018461055c565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610b0257610b01610c04565b5b8060405250919050565b600067ffffffffffffffff821115610b2757610b26610c04565b5b601f19601f8301169050602081019050919050565b600081905092915050565b6000610b5282610b6f565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b6000610bb382610bd6565b9050919050565b6000610bc582610bf2565b9050919050565b6000819050919050565b6000610be182610c40565b9050919050565b6000819050919050565b6000610bfd82610c33565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008160f81b9050919050565b60008160601b9050919050565b610c5681610b47565b8114610c6157600080fd5b50565b610c6d81610b59565b8114610c7857600080fd5b50565b610c8481610b65565b8114610c8f57600080fd5b50565b610c9b81610b8f565b8114610ca657600080fd5b5056fea264697066735822122065a28b05749d51bf6f485e7dc0966ae3f04e462f51378357199698403398394d64736f6c63430008000033"

// DeployHashFake deploys a new Ethereum contract, binding an instance of HashFake to it.
func DeployHashFake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashFake, error) {
	parsed, err := abi.JSON(strings.NewReader(HashFakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HashFakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashFake{HashFakeCaller: HashFakeCaller{contract: contract}, HashFakeTransactor: HashFakeTransactor{contract: contract}, HashFakeFilterer: HashFakeFilterer{contract: contract}}, nil
}

// HashFake is an auto generated Go binding around an Ethereum contract.
type HashFake struct {
	HashFakeCaller     // Read-only binding to the contract
	HashFakeTransactor // Write-only binding to the contract
	HashFakeFilterer   // Log filterer for contract events
}

// HashFakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashFakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashFakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashFakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashFakeSession struct {
	Contract     *HashFake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashFakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashFakeCallerSession struct {
	Contract *HashFakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HashFakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashFakeTransactorSession struct {
	Contract     *HashFakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HashFakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashFakeRaw struct {
	Contract *HashFake // Generic contract binding to access the raw methods on
}

// HashFakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashFakeCallerRaw struct {
	Contract *HashFakeCaller // Generic read-only contract binding to access the raw methods on
}

// HashFakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashFakeTransactorRaw struct {
	Contract *HashFakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashFake creates a new instance of HashFake, bound to a specific deployed contract.
func NewHashFake(address common.Address, backend bind.ContractBackend) (*HashFake, error) {
	contract, err := bindHashFake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashFake{HashFakeCaller: HashFakeCaller{contract: contract}, HashFakeTransactor: HashFakeTransactor{contract: contract}, HashFakeFilterer: HashFakeFilterer{contract: contract}}, nil
}

// NewHashFakeCaller creates a new read-only instance of HashFake, bound to a specific deployed contract.
func NewHashFakeCaller(address common.Address, caller bind.ContractCaller) (*HashFakeCaller, error) {
	contract, err := bindHashFake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashFakeCaller{contract: contract}, nil
}

// NewHashFakeTransactor creates a new write-only instance of HashFake, bound to a specific deployed contract.
func NewHashFakeTransactor(address common.Address, transactor bind.ContractTransactor) (*HashFakeTransactor, error) {
	contract, err := bindHashFake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashFakeTransactor{contract: contract}, nil
}

// NewHashFakeFilterer creates a new log filterer instance of HashFake, bound to a specific deployed contract.
func NewHashFakeFilterer(address common.Address, filterer bind.ContractFilterer) (*HashFakeFilterer, error) {
	contract, err := bindHashFake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashFakeFilterer{contract: contract}, nil
}

// bindHashFake binds a generic wrapper to an already deployed contract.
func bindHashFake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashFakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashFake *HashFakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashFake.Contract.HashFakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashFake *HashFakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashFake.Contract.HashFakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashFake *HashFakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashFake.Contract.HashFakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashFake *HashFakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashFake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashFake *HashFakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashFake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashFake *HashFakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashFake.Contract.contract.Transact(opts, method, params...)
}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashFake *HashFakeCaller) DomainTest(opts *bind.CallOpts, n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "domainTest", n, version, c, verifier)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashFake *HashFakeSession) DomainTest(n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	return _HashFake.Contract.DomainTest(&_HashFake.CallOpts, n, version, c, verifier)
}

// DomainTest is a free data retrieval call binding the contract method 0xaf1ef90b.
//
// Solidity: function domainTest(string n, string version, uint256 c, address verifier) pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) DomainTest(n string, version string, c *big.Int, verifier common.Address) ([32]byte, error) {
	return _HashFake.Contract.DomainTest(&_HashFake.CallOpts, n, version, c, verifier)
}

// DomainTypeHash is a free data retrieval call binding the contract method 0x97995ecb.
//
// Solidity: function domainTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCaller) DomainTypeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "domainTypeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainTypeHash is a free data retrieval call binding the contract method 0x97995ecb.
//
// Solidity: function domainTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeSession) DomainTypeHash() ([32]byte, error) {
	return _HashFake.Contract.DomainTypeHash(&_HashFake.CallOpts)
}

// DomainTypeHash is a free data retrieval call binding the contract method 0x97995ecb.
//
// Solidity: function domainTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) DomainTypeHash() ([32]byte, error) {
	return _HashFake.Contract.DomainTypeHash(&_HashFake.CallOpts)
}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashFake *HashFakeCaller) MessageTest(opts *bind.CallOpts, d [32]byte, h [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "messageTest", d, h)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashFake *HashFakeSession) MessageTest(d [32]byte, h [32]byte) ([32]byte, error) {
	return _HashFake.Contract.MessageTest(&_HashFake.CallOpts, d, h)
}

// MessageTest is a free data retrieval call binding the contract method 0x4c05b2a1.
//
// Solidity: function messageTest(bytes32 d, bytes32 h) pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) MessageTest(d [32]byte, h [32]byte) ([32]byte, error) {
	return _HashFake.Contract.MessageTest(&_HashFake.CallOpts, d, h)
}

// OrderTest is a free data retrieval call binding the contract method 0x207d220c.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeCaller) OrderTest(opts *bind.CallOpts, o HashOrder) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "orderTest", o)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTest is a free data retrieval call binding the contract method 0x207d220c.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashFake.Contract.OrderTest(&_HashFake.CallOpts, o)
}

// OrderTest is a free data retrieval call binding the contract method 0x207d220c.
//
// Solidity: function orderTest((bytes32,address,address,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashFake.Contract.OrderTest(&_HashFake.CallOpts, o)
}

// OrderTypeHash is a free data retrieval call binding the contract method 0x91c2600f.
//
// Solidity: function orderTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCaller) OrderTypeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "orderTypeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTypeHash is a free data retrieval call binding the contract method 0x91c2600f.
//
// Solidity: function orderTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeSession) OrderTypeHash() ([32]byte, error) {
	return _HashFake.Contract.OrderTypeHash(&_HashFake.CallOpts)
}

// OrderTypeHash is a free data retrieval call binding the contract method 0x91c2600f.
//
// Solidity: function orderTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) OrderTypeHash() ([32]byte, error) {
	return _HashFake.Contract.OrderTypeHash(&_HashFake.CallOpts)
}
