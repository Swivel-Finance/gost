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
const HashFakeABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// HashFakeBin is the compiled bytecode used for deploying new contracts.
var HashFakeBin = "0x608060405234801561001057600080fd5b50610d8a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80634c05b2a11461005c57806391c2600f1461008c57806397995ecb146100aa578063af1ef90b146100c8578063e2711de4146100f8575b600080fd5b6100766004803603810190610071919061044a565b610128565b604051610083919061087e565b60405180910390f35b61009461013c565b6040516100a1919061087e565b60405180910390f35b6100b2610168565b6040516100bf919061087e565b60405180910390f35b6100e260048036038101906100dd9190610486565b610194565b6040516100ef919061087e565b60405180910390f35b610112600480360381019061010d9190610519565b6101ac565b60405161011f919061087e565b60405180910390f35b600061013483836101be565b905092915050565b600060405160200161014d906107af565b60405160208183030381529060405280519060200120905090565b600060405160200161017990610832565b60405160208183030381529060405280519060200120905090565b60006101a2858585856101ff565b9050949350505050565b60006101b78261025e565b9050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007f7ddd38ab5ed1c16b61ca90eeb9579e29da1ba821cf42d8cdef8f30a31a6a414660001b826000013583602001602081019061029c91906103f8565b8460400160208101906102af91906103f8565b8560600160208101906102c29190610421565b8660800160208101906102d59190610421565b8760a001358860c001358960e001358a61010001356040516020016103039a99989796959493929190610899565b604051602081830303815290604052805190602001209050919050565b600061033361032e8461095a565b610935565b90508281526020810184848401111561034b57600080fd5b6103568482856109e8565b509392505050565b60008135905061036d81610cf8565b92915050565b60008135905061038281610d0f565b92915050565b60008135905061039781610d26565b92915050565b600082601f8301126103ae57600080fd5b81356103be848260208601610320565b91505092915050565b600061012082840312156103da57600080fd5b81905092915050565b6000813590506103f281610d3d565b92915050565b60006020828403121561040a57600080fd5b60006104188482850161035e565b91505092915050565b60006020828403121561043357600080fd5b600061044184828501610373565b91505092915050565b6000806040838503121561045d57600080fd5b600061046b85828601610388565b925050602061047c85828601610388565b9150509250929050565b6000806000806080858703121561049c57600080fd5b600085013567ffffffffffffffff8111156104b657600080fd5b6104c28782880161039d565b945050602085013567ffffffffffffffff8111156104df57600080fd5b6104eb8782880161039d565b93505060406104fc878288016103e3565b925050606061050d8782880161035e565b91505092959194509250565b6000610120828403121561052c57600080fd5b600061053a848285016103c7565b91505092915050565b61054c81610996565b82525050565b61055b816109a8565b82525050565b61056a816109b4565b82525050565b600061057d600c8361098b565b915061058882610a68565b600c82019050919050565b60006105a0600c8361098b565b91506105ab82610a91565b600c82019050919050565b60006105c360068361098b565b91506105ce82610aba565b600682019050919050565b60006105e660108361098b565b91506105f182610ae3565b601082019050919050565b600061060960138361098b565b915061061482610b0c565b601382019050919050565b600061062c60198361098b565b915061063782610b35565b601982019050919050565b600061064f60118361098b565b915061065a82610b5e565b601182019050919050565b600061067260108361098b565b915061067d82610b87565b601082019050919050565b6000610695600e8361098b565b91506106a082610bb0565b600e82019050919050565b60006106b860018361098b565b91506106c382610bd9565b600182019050919050565b60006106db600f8361098b565b91506106e682610c02565b600f82019050919050565b60006106fe600e8361098b565b915061070982610c2b565b600e82019050919050565b6000610721600b8361098b565b915061072c82610c54565b600b82019050919050565b600061074460128361098b565b915061074f82610c7d565b601282019050919050565b6000610767600a8361098b565b915061077282610ca6565b600a82019050919050565b600061078a600d8361098b565b915061079582610ccf565b600d82019050919050565b6107a9816109de565b82525050565b60006107ba826105b6565b91506107c582610593565b91506107d0826106f1565b91506107db826105fc565b91506107e682610714565b91506107f18261075a565b91506107fc82610737565b9150610807826105d9565b915061081282610642565b915061081d82610688565b9150610828826106ab565b9150819050919050565b600061083d8261077d565b915061084882610570565b9150610853826106ce565b915061085e82610665565b91506108698261061f565b9150610874826106ab565b9150819050919050565b60006020820190506108936000830184610561565b92915050565b6000610140820190506108af600083018d610561565b6108bc602083018c610561565b6108c9604083018b610543565b6108d6606083018a610543565b6108e36080830189610552565b6108f060a0830188610552565b6108fd60c08301876107a0565b61090a60e08301866107a0565b6109186101008301856107a0565b6109266101208301846107a0565b9b9a5050505050505050505050565b600061093f610950565b905061094b82826109f7565b919050565b6000604051905090565b600067ffffffffffffffff82111561097557610974610a28565b5b61097e82610a57565b9050602081019050919050565b600081905092915050565b60006109a1826109be565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b610a0082610a57565b810181811067ffffffffffffffff82111715610a1f57610a1e610a28565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f737472696e67206e616d652c0000000000000000000000000000000000000000600082015250565b7f62797465733332206b65792c0000000000000000000000000000000000000000600082015250565b7f4f72646572280000000000000000000000000000000000000000000000000000600082015250565b7f75696e74323536207072656d69756d2c00000000000000000000000000000000600082015250565b7f6164647265737320756e6465726c79696e672c00000000000000000000000000600082015250565b7f6164647265737320766572696679696e67436f6e747261637400000000000000600082015250565b7f75696e74323536206d617475726974792c000000000000000000000000000000600082015250565b7f75696e7432353620636861696e49642c00000000000000000000000000000000600082015250565b7f75696e7432353620657870697279000000000000000000000000000000000000600082015250565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b7f737472696e672076657273696f6e2c0000000000000000000000000000000000600082015250565b7f61646472657373206d616b65722c000000000000000000000000000000000000600082015250565b7f626f6f6c207661756c742c000000000000000000000000000000000000000000600082015250565b7f75696e74323536207072696e636970616c2c0000000000000000000000000000600082015250565b7f626f6f6c20657869742c00000000000000000000000000000000000000000000600082015250565b7f454950373132446f6d61696e2800000000000000000000000000000000000000600082015250565b610d0181610996565b8114610d0c57600080fd5b50565b610d18816109a8565b8114610d2357600080fd5b50565b610d2f816109b4565b8114610d3a57600080fd5b50565b610d46816109de565b8114610d5157600080fd5b5056fea2646970667358221220a6abcdb8b38bb87a6613ec1487f9f352356fab73a211235a29cb8e5c6b720aa264736f6c63430008040033"

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
