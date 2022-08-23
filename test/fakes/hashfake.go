// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fakes

import (
	"errors"
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
	_ = errors.New
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
	Protocol   uint8
	Maker      common.Address
	Underlying common.Address
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// HashFakeMetaData contains all meta data concerning the HashFake contract.
var HashFakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"protocol\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permitTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611139806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806310ce43bd146100675780634c05b2a11461008557806351c45ca8146100b557806391c2600f146100e557806397995ecb14610103578063af1ef90b14610121575b600080fd5b61006f610151565b60405161007c91906103ad565b60405180910390f35b61009f600480360381019061009a9190610408565b61017d565b6040516100ac91906103ad565b60405180910390f35b6100cf60048036038101906100ca919061046d565b610191565b6040516100dc91906103ad565b60405180910390f35b6100ed6101a3565b6040516100fa91906103ad565b60405180910390f35b61010b6101cf565b60405161011891906103ad565b60405180910390f35b61013b60048036038101906101369190610675565b6101fb565b60405161014891906103ad565b60405180910390f35b600060405160200161016290610933565b60405160208183030381529060405280519060200120905090565b60006101898383610213565b905092915050565b600061019c82610259565b9050919050565b60006040516020016101b490610cce565b60405160208183030381529060405280519060200120905090565b60006040516020016101e090610ed8565b60405160208183030381529060405280519060200120905090565b600061020985858585610330565b9050949350505050565b6000806040517f19010000000000000000000000000000000000000000000000000000000000008152846002820152836022820152604281209150508091505092915050565b60007fbc200cfe92556575f801f821f26e6d54f6421fa132e4b2d65319cac1c687d8e660001b82600001358360200160208101906102979190610f5d565b8460400160208101906102aa9190610f8a565b8560600160208101906102bd9190610f8a565b8660800160208101906102d09190610fef565b8760a00160208101906102e39190610fef565b8860c001358960e001358a61010001358b61012001356040516020016103139b9a99989796959493929190611058565b604051602081830303815290604052805190602001209050919050565b60008085516020870120855160208701206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015286606082015285608082015260a08120935050505080915050949350505050565b6000819050919050565b6103a781610394565b82525050565b60006020820190506103c2600083018461039e565b92915050565b6000604051905090565b600080fd5b600080fd5b6103e581610394565b81146103f057600080fd5b50565b600081359050610402816103dc565b92915050565b6000806040838503121561041f5761041e6103d2565b5b600061042d858286016103f3565b925050602061043e858286016103f3565b9150509250929050565b600080fd5b6000610140828403121561046457610463610448565b5b81905092915050565b60006101408284031215610484576104836103d2565b5b60006104928482850161044d565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104ee826104a5565b810181811067ffffffffffffffff8211171561050d5761050c6104b6565b5b80604052505050565b60006105206103c8565b905061052c82826104e5565b919050565b600067ffffffffffffffff82111561054c5761054b6104b6565b5b610555826104a5565b9050602081019050919050565b82818337600083830152505050565b600061058461057f84610531565b610516565b9050828152602081018484840111156105a05761059f6104a0565b5b6105ab848285610562565b509392505050565b600082601f8301126105c8576105c761049b565b5b81356105d8848260208601610571565b91505092915050565b6000819050919050565b6105f4816105e1565b81146105ff57600080fd5b50565b600081359050610611816105eb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061064282610617565b9050919050565b61065281610637565b811461065d57600080fd5b50565b60008135905061066f81610649565b92915050565b6000806000806080858703121561068f5761068e6103d2565b5b600085013567ffffffffffffffff8111156106ad576106ac6103d7565b5b6106b9878288016105b3565b945050602085013567ffffffffffffffff8111156106da576106d96103d7565b5b6106e6878288016105b3565b93505060406106f787828801610602565b925050606061070887828801610660565b91505092959194509250565b600081905092915050565b7f5065726d69742800000000000000000000000000000000000000000000000000600082015250565b6000610755600783610714565b91506107608261071f565b600782019050919050565b7f61646472657373206f776e65722c000000000000000000000000000000000000600082015250565b60006107a1600e83610714565b91506107ac8261076b565b600e82019050919050565b7f61646472657373207370656e6465722c00000000000000000000000000000000600082015250565b60006107ed601083610714565b91506107f8826107b7565b601082019050919050565b7f75696e743235362076616c75652c000000000000000000000000000000000000600082015250565b6000610839600e83610714565b915061084482610803565b600e82019050919050565b7f75696e74323536206e6f6e63652c000000000000000000000000000000000000600082015250565b6000610885600e83610714565b91506108908261084f565b600e82019050919050565b7f75696e7432353620646561646c696e652c000000000000000000000000000000600082015250565b60006108d1601183610714565b91506108dc8261089b565b601182019050919050565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b600061091d600183610714565b9150610928826108e7565b600182019050919050565b600061093e82610748565b915061094982610794565b9150610954826107e0565b915061095f8261082c565b915061096a82610878565b9150610975826108c4565b915061098082610910565b9150819050919050565b7f4f72646572280000000000000000000000000000000000000000000000000000600082015250565b60006109c0600683610714565b91506109cb8261098a565b600682019050919050565b7f62797465733332206b65792c0000000000000000000000000000000000000000600082015250565b6000610a0c600c83610714565b9150610a17826109d6565b600c82019050919050565b7f75696e74382070726f746f636f6c2c0000000000000000000000000000000000600082015250565b6000610a58600f83610714565b9150610a6382610a22565b600f82019050919050565b7f61646472657373206d616b65722c000000000000000000000000000000000000600082015250565b6000610aa4600e83610714565b9150610aaf82610a6e565b600e82019050919050565b7f6164647265737320756e6465726c79696e672c00000000000000000000000000600082015250565b6000610af0601383610714565b9150610afb82610aba565b601382019050919050565b7f626f6f6c207661756c742c000000000000000000000000000000000000000000600082015250565b6000610b3c600b83610714565b9150610b4782610b06565b600b82019050919050565b7f626f6f6c20657869742c00000000000000000000000000000000000000000000600082015250565b6000610b88600a83610714565b9150610b9382610b52565b600a82019050919050565b7f75696e74323536207072696e636970616c2c0000000000000000000000000000600082015250565b6000610bd4601283610714565b9150610bdf82610b9e565b601282019050919050565b7f75696e74323536207072656d69756d2c00000000000000000000000000000000600082015250565b6000610c20601083610714565b9150610c2b82610bea565b601082019050919050565b7f75696e74323536206d617475726974792c000000000000000000000000000000600082015250565b6000610c6c601183610714565b9150610c7782610c36565b601182019050919050565b7f75696e7432353620657870697279000000000000000000000000000000000000600082015250565b6000610cb8600e83610714565b9150610cc382610c82565b600e82019050919050565b6000610cd9826109b3565b9150610ce4826109ff565b9150610cef82610a4b565b9150610cfa82610a97565b9150610d0582610ae3565b9150610d1082610b2f565b9150610d1b82610b7b565b9150610d2682610bc7565b9150610d3182610c13565b9150610d3c82610c5f565b9150610d4782610cab565b9150610d5282610910565b9150819050919050565b7f454950373132446f6d61696e2800000000000000000000000000000000000000600082015250565b6000610d92600d83610714565b9150610d9d82610d5c565b600d82019050919050565b7f737472696e67206e616d652c0000000000000000000000000000000000000000600082015250565b6000610dde600c83610714565b9150610de982610da8565b600c82019050919050565b7f737472696e672076657273696f6e2c0000000000000000000000000000000000600082015250565b6000610e2a600f83610714565b9150610e3582610df4565b600f82019050919050565b7f75696e7432353620636861696e49642c00000000000000000000000000000000600082015250565b6000610e76601083610714565b9150610e8182610e40565b601082019050919050565b7f6164647265737320766572696679696e67436f6e747261637400000000000000600082015250565b6000610ec2601983610714565b9150610ecd82610e8c565b601982019050919050565b6000610ee382610d85565b9150610eee82610dd1565b9150610ef982610e1d565b9150610f0482610e69565b9150610f0f82610eb5565b9150610f1a82610910565b9150819050919050565b600060ff82169050919050565b610f3a81610f24565b8114610f4557600080fd5b50565b600081359050610f5781610f31565b92915050565b600060208284031215610f7357610f726103d2565b5b6000610f8184828501610f48565b91505092915050565b600060208284031215610fa057610f9f6103d2565b5b6000610fae84828501610660565b91505092915050565b60008115159050919050565b610fcc81610fb7565b8114610fd757600080fd5b50565b600081359050610fe981610fc3565b92915050565b600060208284031215611005576110046103d2565b5b600061101384828501610fda565b91505092915050565b61102581610f24565b82525050565b61103481610637565b82525050565b61104381610fb7565b82525050565b611052816105e1565b82525050565b60006101608201905061106e600083018e61039e565b61107b602083018d61039e565b611088604083018c61101c565b611095606083018b61102b565b6110a2608083018a61102b565b6110af60a083018961103a565b6110bc60c083018861103a565b6110c960e0830187611049565b6110d7610100830186611049565b6110e5610120830185611049565b6110f3610140830184611049565b9c9b50505050505050505050505056fea26469706673582212209375268d60af4cb7425e017053355bf5c3f464b6417b15a5c41c63ca3245608e64736f6c63430008100033",
}

// HashFakeABI is the input ABI used to generate the binding from.
// Deprecated: Use HashFakeMetaData.ABI instead.
var HashFakeABI = HashFakeMetaData.ABI

// HashFakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashFakeMetaData.Bin instead.
var HashFakeBin = HashFakeMetaData.Bin

// DeployHashFake deploys a new Ethereum contract, binding an instance of HashFake to it.
func DeployHashFake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashFake, error) {
	parsed, err := HashFakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashFakeBin), backend)
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

// OrderTest is a free data retrieval call binding the contract method 0x51c45ca8.
//
// Solidity: function orderTest((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeCaller) OrderTest(opts *bind.CallOpts, o HashOrder) ([32]byte, error) {
	var out []interface{}
	err := _HashFake.contract.Call(opts, &out, "orderTest", o)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrderTest is a free data retrieval call binding the contract method 0x51c45ca8.
//
// Solidity: function orderTest((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
func (_HashFake *HashFakeSession) OrderTest(o HashOrder) ([32]byte, error) {
	return _HashFake.Contract.OrderTest(&_HashFake.CallOpts, o)
}

// OrderTest is a free data retrieval call binding the contract method 0x51c45ca8.
//
// Solidity: function orderTest((bytes32,uint8,address,address,bool,bool,uint256,uint256,uint256,uint256) o) pure returns(bytes32)
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
