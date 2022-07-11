// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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

// CreatorMetaData contains all meta data concerning the Creator contract.
var CreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sw\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"createCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swivel\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"createReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c35806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636363e86714610046578063ca359fea14610077578063d219cdc714610093575b600080fd5b610060600480360381019061005b91906108e9565b6100c9565b60405161006e9291906109e6565b60405180910390f35b610091600480360381019061008c9190610a0f565b610358565b005b6100ad60048036038101906100a89190610a4f565b6103de565b6040516100c09796959493929190610b22565b60405180910390f35b6000806100d461059d565b89816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508881602001818152505087816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505086816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050858160800181905250848160a00181905250838160c0019060ff16908160ff1681525050806000808d60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190805190602001906102be92919061061f565b5060a08201518160050190805190602001906102db92919061061f565b5060c08201518160060160006101000a81548160ff021916908360ff160217905550905050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692509250509850989650505050505050565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600401805461047990610bce565b80601f01602080910402602001604051908101604052809291908181526020018280546104a590610bce565b80156104f25780601f106104c7576101008083540402835291602001916104f2565b820191906000526020600020905b8154815290600101906020018083116104d557829003601f168201915b50505050509080600501805461050790610bce565b80601f016020809104026020016040519081016040528092919081815260200182805461053390610bce565b80156105805780601f1061055557610100808354040283529160200191610580565b820191906000526020600020905b81548152906001019060200180831161056357829003601f168201915b5050505050908060060160009054906101000a900460ff16905087565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600060ff1681525090565b82805461062b90610bce565b90600052602060002090601f01602090048101928261064d5760008555610694565b82601f1061066657805160ff1916838001178555610694565b82800160010185558215610694579182015b82811115610693578251825591602001919060010190610678565b5b5090506106a191906106a5565b5090565b5b808211156106be5760008160009055506001016106a6565b5090565b6000604051905090565b600080fd5b600080fd5b600060ff82169050919050565b6106ec816106d6565b81146106f757600080fd5b50565b600081359050610709816106e3565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061073a8261070f565b9050919050565b61074a8161072f565b811461075557600080fd5b50565b60008135905061076781610741565b92915050565b6000819050919050565b6107808161076d565b811461078b57600080fd5b50565b60008135905061079d81610777565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107f6826107ad565b810181811067ffffffffffffffff82111715610815576108146107be565b5b80604052505050565b60006108286106c2565b905061083482826107ed565b919050565b600067ffffffffffffffff821115610854576108536107be565b5b61085d826107ad565b9050602081019050919050565b82818337600083830152505050565b600061088c61088784610839565b61081e565b9050828152602081018484840111156108a8576108a76107a8565b5b6108b384828561086a565b509392505050565b600082601f8301126108d0576108cf6107a3565b5b81356108e0848260208601610879565b91505092915050565b600080600080600080600080610100898b03121561090a576109096106cc565b5b60006109188b828c016106fa565b98505060206109298b828c01610758565b975050604061093a8b828c0161078e565b965050606061094b8b828c01610758565b955050608061095c8b828c01610758565b94505060a089013567ffffffffffffffff81111561097d5761097c6106d1565b5b6109898b828c016108bb565b93505060c089013567ffffffffffffffff8111156109aa576109a96106d1565b5b6109b68b828c016108bb565b92505060e06109c78b828c016106fa565b9150509295985092959890939650565b6109e08161072f565b82525050565b60006040820190506109fb60008301856109d7565b610a0860208301846109d7565b9392505050565b60008060408385031215610a2657610a256106cc565b5b6000610a3485828601610758565b9250506020610a4585828601610758565b9150509250929050565b600060208284031215610a6557610a646106cc565b5b6000610a73848285016106fa565b91505092915050565b610a858161076d565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ac5578082015181840152602081019050610aaa565b83811115610ad4576000848401525b50505050565b6000610ae582610a8b565b610aef8185610a96565b9350610aff818560208601610aa7565b610b08816107ad565b840191505092915050565b610b1c816106d6565b82525050565b600060e082019050610b37600083018a6109d7565b610b446020830189610a7c565b610b5160408301886109d7565b610b5e60608301876109d7565b8181036080830152610b708186610ada565b905081810360a0830152610b848185610ada565b9050610b9360c0830184610b13565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610be657607f821691505b602082108103610bf957610bf8610b9f565b5b5091905056fea264697066735822122041902c5645b01ec0b05f2a940cb094bb2b8713b32326b40c184aaf7731c614ff64736f6c634300080d0033",
}

// CreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorMetaData.ABI instead.
var CreatorABI = CreatorMetaData.ABI

// CreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreatorMetaData.Bin instead.
var CreatorBin = CreatorMetaData.Bin

// DeployCreator deploys a new Ethereum contract, binding an instance of Creator to it.
func DeployCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Creator, error) {
	parsed, err := CreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// Creator is an auto generated Go binding around an Ethereum contract.
type Creator struct {
	CreatorCaller     // Read-only binding to the contract
	CreatorTransactor // Write-only binding to the contract
	CreatorFilterer   // Log filterer for contract events
}

// CreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorSession struct {
	Contract     *Creator          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorCallerSession struct {
	Contract *CreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorTransactorSession struct {
	Contract     *CreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorRaw struct {
	Contract *Creator // Generic contract binding to access the raw methods on
}

// CreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorCallerRaw struct {
	Contract *CreatorCaller // Generic read-only contract binding to access the raw methods on
}

// CreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorTransactorRaw struct {
	Contract *CreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreator creates a new instance of Creator, bound to a specific deployed contract.
func NewCreator(address common.Address, backend bind.ContractBackend) (*Creator, error) {
	contract, err := bindCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// NewCreatorCaller creates a new read-only instance of Creator, bound to a specific deployed contract.
func NewCreatorCaller(address common.Address, caller bind.ContractCaller) (*CreatorCaller, error) {
	contract, err := bindCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCaller{contract: contract}, nil
}

// NewCreatorTransactor creates a new write-only instance of Creator, bound to a specific deployed contract.
func NewCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatorTransactor, error) {
	contract, err := bindCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTransactor{contract: contract}, nil
}

// NewCreatorFilterer creates a new log filterer instance of Creator, bound to a specific deployed contract.
func NewCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatorFilterer, error) {
	contract, err := bindCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorFilterer{contract: contract}, nil
}

// bindCreator binds a generic wrapper to an already deployed contract.
func bindCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.CreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transact(opts, method, params...)
}

// CreateCalled is a free data retrieval call binding the contract method 0xd219cdc7.
//
// Solidity: function createCalled(uint8 ) view returns(address underlying, uint256 maturity, address cToken, address swivel, string name, string symbol, uint8 decimals)
func (_Creator *CreatorCaller) CreateCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	Swivel     common.Address
	Name       string
	Symbol     string
	Decimals   uint8
}, error) {
	var out []interface{}
	err := _Creator.contract.Call(opts, &out, "createCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		CToken     common.Address
		Swivel     common.Address
		Name       string
		Symbol     string
		Decimals   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Swivel = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// CreateCalled is a free data retrieval call binding the contract method 0xd219cdc7.
//
// Solidity: function createCalled(uint8 ) view returns(address underlying, uint256 maturity, address cToken, address swivel, string name, string symbol, uint8 decimals)
func (_Creator *CreatorSession) CreateCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	Swivel     common.Address
	Name       string
	Symbol     string
	Decimals   uint8
}, error) {
	return _Creator.Contract.CreateCalled(&_Creator.CallOpts, arg0)
}

// CreateCalled is a free data retrieval call binding the contract method 0xd219cdc7.
//
// Solidity: function createCalled(uint8 ) view returns(address underlying, uint256 maturity, address cToken, address swivel, string name, string symbol, uint8 decimals)
func (_Creator *CreatorCallerSession) CreateCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	Swivel     common.Address
	Name       string
	Symbol     string
	Decimals   uint8
}, error) {
	return _Creator.Contract.CreateCalled(&_Creator.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorTransactor) Create(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "create", p, u, m, c, sw, n, s, d)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorSession) Create(p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.Contract.Create(&_Creator.TransactOpts, p, u, m, c, sw, n, s, d)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorTransactorSession) Create(p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.Contract.Create(&_Creator.TransactOpts, p, u, m, c, sw, n, s, d)
}

// CreateReturns is a paid mutator transaction binding the contract method 0xca359fea.
//
// Solidity: function createReturns(address z, address v) returns()
func (_Creator *CreatorTransactor) CreateReturns(opts *bind.TransactOpts, z common.Address, v common.Address) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "createReturns", z, v)
}

// CreateReturns is a paid mutator transaction binding the contract method 0xca359fea.
//
// Solidity: function createReturns(address z, address v) returns()
func (_Creator *CreatorSession) CreateReturns(z common.Address, v common.Address) (*types.Transaction, error) {
	return _Creator.Contract.CreateReturns(&_Creator.TransactOpts, z, v)
}

// CreateReturns is a paid mutator transaction binding the contract method 0xca359fea.
//
// Solidity: function createReturns(address z, address v) returns()
func (_Creator *CreatorTransactorSession) CreateReturns(z common.Address, v common.Address) (*types.Transaction, error) {
	return _Creator.Contract.CreateReturns(&_Creator.TransactOpts, z, v)
}
