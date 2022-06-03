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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"domainTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"}],\"name\":\"messageTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"orderTest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permitTypeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611049806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806310ce43bd146100675780634c05b2a11461008557806391c2600f146100b557806397995ecb146100d3578063af1ef90b146100f1578063e2711de414610121575b600080fd5b61006f610151565b60405161007c9190610398565b60405180910390f35b61009f600480360381019061009a91906103f3565b61017d565b6040516100ac9190610398565b60405180910390f35b6100bd610191565b6040516100ca9190610398565b60405180910390f35b6100db6101bd565b6040516100e89190610398565b60405180910390f35b61010b6004803603810190610106919061060d565b6101e9565b6040516101189190610398565b60405180910390f35b61013b600480360381019061013691906106d1565b610201565b6040516101489190610398565b60405180910390f35b60006040516020016101629061091e565b60405160208183030381529060405280519060200120905090565b60006101898383610213565b905092915050565b60006040516020016101a290610c6d565b60405160208183030381529060405280519060200120905090565b60006040516020016101ce90610e6c565b60405160208183030381529060405280519060200120905090565b60006101f785858585610259565b9050949350505050565b600061020c826102bd565b9050919050565b6000806040517f19010000000000000000000000000000000000000000000000000000000000008152846002820152836022820152604281209150508091505092915050565b60008085516020870120855160208701206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015286606082015285608082015260a08120935050505080915050949350505050565b60007f7ddd38ab5ed1c16b61ca90eeb9579e29da1ba821cf42d8cdef8f30a31a6a414660001b82600001358360200160208101906102fb9190610eb8565b84604001602081019061030e9190610eb8565b8560600160208101906103219190610f1d565b8660800160208101906103349190610f1d565b8760a001358860c001358960e001358a61010001356040516020016103629a99989796959493929190610f77565b604051602081830303815290604052805190602001209050919050565b6000819050919050565b6103928161037f565b82525050565b60006020820190506103ad6000830184610389565b92915050565b6000604051905090565b600080fd5b600080fd5b6103d08161037f565b81146103db57600080fd5b50565b6000813590506103ed816103c7565b92915050565b6000806040838503121561040a576104096103bd565b5b6000610418858286016103de565b9250506020610429858286016103de565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104868261043d565b810181811067ffffffffffffffff821117156104a5576104a461044e565b5b80604052505050565b60006104b86103b3565b90506104c4828261047d565b919050565b600067ffffffffffffffff8211156104e4576104e361044e565b5b6104ed8261043d565b9050602081019050919050565b82818337600083830152505050565b600061051c610517846104c9565b6104ae565b90508281526020810184848401111561053857610537610438565b5b6105438482856104fa565b509392505050565b600082601f8301126105605761055f610433565b5b8135610570848260208601610509565b91505092915050565b6000819050919050565b61058c81610579565b811461059757600080fd5b50565b6000813590506105a981610583565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105da826105af565b9050919050565b6105ea816105cf565b81146105f557600080fd5b50565b600081359050610607816105e1565b92915050565b60008060008060808587031215610627576106266103bd565b5b600085013567ffffffffffffffff811115610645576106446103c2565b5b6106518782880161054b565b945050602085013567ffffffffffffffff811115610672576106716103c2565b5b61067e8782880161054b565b935050604061068f8782880161059a565b92505060606106a0878288016105f8565b91505092959194509250565b600080fd5b600061012082840312156106c8576106c76106ac565b5b81905092915050565b600061012082840312156106e8576106e76103bd565b5b60006106f6848285016106b1565b91505092915050565b600081905092915050565b7f5065726d69742800000000000000000000000000000000000000000000000000600082015250565b60006107406007836106ff565b915061074b8261070a565b600782019050919050565b7f61646472657373206f776e65722c000000000000000000000000000000000000600082015250565b600061078c600e836106ff565b915061079782610756565b600e82019050919050565b7f61646472657373207370656e6465722c00000000000000000000000000000000600082015250565b60006107d86010836106ff565b91506107e3826107a2565b601082019050919050565b7f75696e743235362076616c75652c000000000000000000000000000000000000600082015250565b6000610824600e836106ff565b915061082f826107ee565b600e82019050919050565b7f75696e74323536206e6f6e63652c000000000000000000000000000000000000600082015250565b6000610870600e836106ff565b915061087b8261083a565b600e82019050919050565b7f75696e7432353620646561646c696e652c000000000000000000000000000000600082015250565b60006108bc6011836106ff565b91506108c782610886565b601182019050919050565b7f2900000000000000000000000000000000000000000000000000000000000000600082015250565b60006109086001836106ff565b9150610913826108d2565b600182019050919050565b600061092982610733565b91506109348261077f565b915061093f826107cb565b915061094a82610817565b915061095582610863565b9150610960826108af565b915061096b826108fb565b9150819050919050565b7f4f72646572280000000000000000000000000000000000000000000000000000600082015250565b60006109ab6006836106ff565b91506109b682610975565b600682019050919050565b7f62797465733332206b65792c0000000000000000000000000000000000000000600082015250565b60006109f7600c836106ff565b9150610a02826109c1565b600c82019050919050565b7f61646472657373206d616b65722c000000000000000000000000000000000000600082015250565b6000610a43600e836106ff565b9150610a4e82610a0d565b600e82019050919050565b7f6164647265737320756e6465726c79696e672c00000000000000000000000000600082015250565b6000610a8f6013836106ff565b9150610a9a82610a59565b601382019050919050565b7f626f6f6c207661756c742c000000000000000000000000000000000000000000600082015250565b6000610adb600b836106ff565b9150610ae682610aa5565b600b82019050919050565b7f626f6f6c20657869742c00000000000000000000000000000000000000000000600082015250565b6000610b27600a836106ff565b9150610b3282610af1565b600a82019050919050565b7f75696e74323536207072696e636970616c2c0000000000000000000000000000600082015250565b6000610b736012836106ff565b9150610b7e82610b3d565b601282019050919050565b7f75696e74323536207072656d69756d2c00000000000000000000000000000000600082015250565b6000610bbf6010836106ff565b9150610bca82610b89565b601082019050919050565b7f75696e74323536206d617475726974792c000000000000000000000000000000600082015250565b6000610c0b6011836106ff565b9150610c1682610bd5565b601182019050919050565b7f75696e7432353620657870697279000000000000000000000000000000000000600082015250565b6000610c57600e836106ff565b9150610c6282610c21565b600e82019050919050565b6000610c788261099e565b9150610c83826109ea565b9150610c8e82610a36565b9150610c9982610a82565b9150610ca482610ace565b9150610caf82610b1a565b9150610cba82610b66565b9150610cc582610bb2565b9150610cd082610bfe565b9150610cdb82610c4a565b9150610ce6826108fb565b9150819050919050565b7f454950373132446f6d61696e2800000000000000000000000000000000000000600082015250565b6000610d26600d836106ff565b9150610d3182610cf0565b600d82019050919050565b7f737472696e67206e616d652c0000000000000000000000000000000000000000600082015250565b6000610d72600c836106ff565b9150610d7d82610d3c565b600c82019050919050565b7f737472696e672076657273696f6e2c0000000000000000000000000000000000600082015250565b6000610dbe600f836106ff565b9150610dc982610d88565b600f82019050919050565b7f75696e7432353620636861696e49642c00000000000000000000000000000000600082015250565b6000610e0a6010836106ff565b9150610e1582610dd4565b601082019050919050565b7f6164647265737320766572696679696e67436f6e747261637400000000000000600082015250565b6000610e566019836106ff565b9150610e6182610e20565b601982019050919050565b6000610e7782610d19565b9150610e8282610d65565b9150610e8d82610db1565b9150610e9882610dfd565b9150610ea382610e49565b9150610eae826108fb565b9150819050919050565b600060208284031215610ece57610ecd6103bd565b5b6000610edc848285016105f8565b91505092915050565b60008115159050919050565b610efa81610ee5565b8114610f0557600080fd5b50565b600081359050610f1781610ef1565b92915050565b600060208284031215610f3357610f326103bd565b5b6000610f4184828501610f08565b91505092915050565b610f53816105cf565b82525050565b610f6281610ee5565b82525050565b610f7181610579565b82525050565b600061014082019050610f8d600083018d610389565b610f9a602083018c610389565b610fa7604083018b610f4a565b610fb4606083018a610f4a565b610fc16080830189610f59565b610fce60a0830188610f59565b610fdb60c0830187610f68565b610fe860e0830186610f68565b610ff6610100830185610f68565b611004610120830184610f68565b9b9a505050505050505050505056fea2646970667358221220d2ec69c52211f2f0b401e7edc21e1ca71da73bd550a545c9778492328b38b6ad64736f6c634300080d0033",
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
