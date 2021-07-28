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
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// HashFakeABI is the input ABI used to generate the binding from.
const HashFakeABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permitTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HashFakeBin is the compiled bytecode used for deploying new contracts.
var HashFakeBin = "0x608060405234801561001057600080fd5b50610ffe806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806310ce43bd146100675780634c05b2a11461008557806391c2600f146100b557806397995ecb146100d3578063af1ef90b146100f1578063e2711de414610121575b600080fd5b61006f610151565b60405161007c91906109fc565b60405180910390f35b61009f600480360381019061009a919061049f565b61017d565b6040516100ac91906109fc565b60405180910390f35b6100bd610191565b6040516100ca91906109fc565b60405180910390f35b6100db6101bd565b6040516100e891906109fc565b60405180910390f35b61010b600480360381019061010691906104db565b6101e9565b60405161011891906109fc565b60405180910390f35b61013b6004803603810190610136919061056e565b610201565b60405161014891906109fc565b60405180910390f35b600060405160200161016290610959565b60405160208183030381529060405280519060200120905090565b60006101898383610213565b905092915050565b60006040516020016101a2906108d6565b60405160208183030381529060405280519060200120905090565b60006040516020016101ce906109b0565b60405160208183030381529060405280519060200120905090565b60006101f785858585610254565b9050949350505050565b600061020c826102b3565b9050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007f7ddd38ab5ed1c16b61ca90eeb9579e29da1ba821cf42d8cdef8f30a31a6a414660001b82600001358360200160208101906102f1919061044d565b846040016020810190610304919061044d565b8560600160208101906103179190610476565b86608001602081019061032a9190610476565b8760a001358860c001358960e001358a61010001356040516020016103589a99989796959493929190610a17565b604051602081830303815290604052805190602001209050919050565b600061038861038384610ad8565b610ab3565b9050828152602081018484840111156103a057600080fd5b6103ab848285610b66565b509392505050565b6000813590506103c281610f6c565b92915050565b6000813590506103d781610f83565b92915050565b6000813590506103ec81610f9a565b92915050565b600082601f83011261040357600080fd5b8135610413848260208601610375565b91505092915050565b6000610120828403121561042f57600080fd5b81905092915050565b60008135905061044781610fb1565b92915050565b60006020828403121561045f57600080fd5b600061046d848285016103b3565b91505092915050565b60006020828403121561048857600080fd5b6000610496848285016103c8565b91505092915050565b600080604083850312156104b257600080fd5b60006104c0858286016103dd565b92505060206104d1858286016103dd565b9150509250929050565b600080600080608085870312156104f157600080fd5b600085013567ffffffffffffffff81111561050b57600080fd5b610517878288016103f2565b945050602085013567ffffffffffffffff81111561053457600080fd5b610540878288016103f2565b935050604061055187828801610438565b9250506060610562878288016103b3565b91505092959194509250565b6000610120828403121561058157600080fd5b600061058f8482850161041c565b91505092915050565b6105a181610b14565b82525050565b6105b081610b26565b82525050565b6105bf81610b32565b82525050565b60006105d2600c83610b09565b91506105dd82610be6565b600c82019050919050565b60006105f5600c83610b09565b915061060082610c0f565b600c82019050919050565b6000610618600e83610b09565b915061062382610c38565b600e82019050919050565b600061063b600683610b09565b915061064682610c61565b600682019050919050565b600061065e601083610b09565b915061066982610c8a565b601082019050919050565b6000610681601383610b09565b915061068c82610cb3565b601382019050919050565b60006106a4601983610b09565b91506106af82610cdc565b601982019050919050565b60006106c7601183610b09565b91506106d282610d05565b601182019050919050565b60006106ea601083610b09565b91506106f582610d2e565b601082019050919050565b600061070d601183610b09565b915061071882610d57565b601182019050919050565b6000610730600e83610b09565b915061073b82610d80565b600e82019050919050565b6000610753600183610b09565b915061075e82610da9565b600182019050919050565b6000610776600f83610b09565b915061078182610dd2565b600f82019050919050565b6000610799601083610b09565b91506107a482610dfb565b601082019050919050565b60006107bc600783610b09565b91506107c782610e24565b600782019050919050565b60006107df600e83610b09565b91506107ea82610e4d565b600e82019050919050565b6000610802600e83610b09565b915061080d82610e76565b600e82019050919050565b6000610825600b83610b09565b915061083082610e9f565b600b82019050919050565b6000610848601283610b09565b915061085382610ec8565b601282019050919050565b600061086b600a83610b09565b915061087682610ef1565b600a82019050919050565b600061088e600e83610b09565b915061089982610f1a565b600e82019050919050565b60006108b1600d83610b09565b91506108bc82610f43565b600d82019050919050565b6108d081610b5c565b82525050565b60006108e18261062e565b91506108ec826105e8565b91506108f7826107f5565b915061090282610674565b915061090d82610818565b91506109188261085e565b91506109238261083b565b915061092e82610651565b9150610939826106ba565b915061094482610723565b915061094f82610746565b9150819050919050565b6000610964826107af565b915061096f826107d2565b915061097a8261078c565b915061098582610881565b91506109908261060b565b915061099b82610700565b91506109a682610746565b9150819050919050565b60006109bb826108a4565b91506109c6826105c5565b91506109d182610769565b91506109dc826106dd565b91506109e782610697565b91506109f282610746565b9150819050919050565b6000602082019050610a1160008301846105b6565b92915050565b600061014082019050610a2d600083018d6105b6565b610a3a602083018c6105b6565b610a47604083018b610598565b610a54606083018a610598565b610a6160808301896105a7565b610a6e60a08301886105a7565b610a7b60c08301876108c7565b610a8860e08301866108c7565b610a966101008301856108c7565b610aa46101208301846108c7565b9b9a5050505050505050505050565b6000610abd610ace565b9050610ac98282610b75565b919050565b6000604051905090565b600067ffffffffffffffff821115610af357610af2610ba6565b5b610afc82610bd5565b9050602081019050919050565b600081905092915050565b6000610b1f82610b3c565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610b7e82610bd5565b810181811067ffffffffffffffff82111715610b9d57610b9c610ba6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f737472696e67206e616d652c0000000000000000000000000000000000000000600082015250565b7f62797465733332206b65792c0000000000000000000000000000000000000000600082015250565b7f75696e74323536206e6f6e63652c000000000000000000000000000000000000600082015250565b7f4f72646572280000000000000000000000000000000000000000000000000000600082015250565b7f75696e74323536207072656d69756d2c00000000000000000000000000000000600082015250565b7f6164647265737320756e6465726c79696e672c00000000000000000000000000600082015250565b7f6164647265737320766572696679696e67436f6e747261637400000000000000600082015250565b7f75696e74323536206d617475726974792c000000000000000000000000000000600082015250565b7f75696e7432353620636861696e49642c00000000000000000000000000000000600082015250565b7f75696e7432353620646561646c696e652c000000000000000000000000000000600082015250565b7f75696e7432353620657870697279000000000000000000000000000000000000600082015250565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b7f737472696e672076657273696f6e2c0000000000000000000000000000000000600082015250565b7f61646472657373207370656e6465722c00000000000000000000000000000000600082015250565b7f5065726d69742800000000000000000000000000000000000000000000000000600082015250565b7f61646472657373206f776e65722c000000000000000000000000000000000000600082015250565b7f61646472657373206d616b65722c000000000000000000000000000000000000600082015250565b7f626f6f6c207661756c742c000000000000000000000000000000000000000000600082015250565b7f75696e74323536207072696e636970616c2c0000000000000000000000000000600082015250565b7f626f6f6c20657869742c00000000000000000000000000000000000000000000600082015250565b7f75696e743235362076616c75652c000000000000000000000000000000000000600082015250565b7f454950373132446f6d61696e2800000000000000000000000000000000000000600082015250565b610f7581610b14565b8114610f8057600080fd5b50565b610f8c81610b26565b8114610f9757600080fd5b50565b610fa381610b32565b8114610fae57600080fd5b50565b610fba81610b5c565b8114610fc557600080fd5b5056fea2646970667358221220aaa20b1fc8dfa2b1300958ee4888d3fde206355fc3d0ba6ef36febe57c14f19664736f6c63430008040033"

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

// OrderTest is a free data retrieval call binding the contract method 0xe2711de4.
//
// Solidity: function orderTest((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeCaller) OrderTest(opts *bind.CallOpts, o HashOrder) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "orderTest", o)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTest is a free data retrieval call binding the contract method 0xe2711de4.
//
// Solidity: function orderTest((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashFake.Contract.OrderTest(&_HashFake.CallOpts, o)
}

// OrderTest is a free data retrieval call binding the contract method 0xe2711de4.
//
// Solidity: function orderTest((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
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

// PermitTypeHash is a free data retrieval call binding the contract method 0x10ce43bd.
//
// Solidity: function permitTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCaller) PermitTypeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "permitTypeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PermitTypeHash is a free data retrieval call binding the contract method 0x10ce43bd.
//
// Solidity: function permitTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeSession) PermitTypeHash() ([32]byte, error) {
	return _HashFake.Contract.PermitTypeHash(&_HashFake.CallOpts)
}

// PermitTypeHash is a free data retrieval call binding the contract method 0x10ce43bd.
//
// Solidity: function permitTypeHash() pure returns(bytes32)
func (_HashFake *HashFakeCallerSession) PermitTypeHash() ([32]byte, error) {
	return _HashFake.Contract.PermitTypeHash(&_HashFake.CallOpts)
}
