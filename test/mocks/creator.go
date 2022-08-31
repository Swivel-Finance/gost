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
	Bin: "0x608060405234801561001057600080fd5b50610df9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636363e86714610046578063ca359fea14610077578063d219cdc714610093575b600080fd5b610060600480360381019061005b9190610838565b6100c9565b60405161006e929190610935565b60405180910390f35b610091600480360381019061008c919061095e565b61034a565b005b6100ad60048036038101906100a8919061099e565b6103d0565b6040516100c09796959493929190610a68565b60405180910390f35b6000806100d461058f565b89816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508881602001818152505087816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505086816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050858160800181905250848160a00181905250838160c0019060ff16908160ff1681525050806000808d60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040190816102b79190610cf1565b5060a08201518160050190816102cd9190610cf1565b5060c08201518160060160006101000a81548160ff021916908360ff160217905550905050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692509250509850989650505050505050565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600401805461046b90610b14565b80601f016020809104026020016040519081016040528092919081815260200182805461049790610b14565b80156104e45780601f106104b9576101008083540402835291602001916104e4565b820191906000526020600020905b8154815290600101906020018083116104c757829003601f168201915b5050505050908060050180546104f990610b14565b80601f016020809104026020016040519081016040528092919081815260200182805461052590610b14565b80156105725780601f1061054757610100808354040283529160200191610572565b820191906000526020600020905b81548152906001019060200180831161055557829003601f168201915b5050505050908060060160009054906101000a900460ff16905087565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001600060ff1681525090565b6000604051905090565b600080fd5b600080fd5b600060ff82169050919050565b61063b81610625565b811461064657600080fd5b50565b60008135905061065881610632565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106898261065e565b9050919050565b6106998161067e565b81146106a457600080fd5b50565b6000813590506106b681610690565b92915050565b6000819050919050565b6106cf816106bc565b81146106da57600080fd5b50565b6000813590506106ec816106c6565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610745826106fc565b810181811067ffffffffffffffff821117156107645761076361070d565b5b80604052505050565b6000610777610611565b9050610783828261073c565b919050565b600067ffffffffffffffff8211156107a3576107a261070d565b5b6107ac826106fc565b9050602081019050919050565b82818337600083830152505050565b60006107db6107d684610788565b61076d565b9050828152602081018484840111156107f7576107f66106f7565b5b6108028482856107b9565b509392505050565b600082601f83011261081f5761081e6106f2565b5b813561082f8482602086016107c8565b91505092915050565b600080600080600080600080610100898b0312156108595761085861061b565b5b60006108678b828c01610649565b98505060206108788b828c016106a7565b97505060406108898b828c016106dd565b965050606061089a8b828c016106a7565b95505060806108ab8b828c016106a7565b94505060a089013567ffffffffffffffff8111156108cc576108cb610620565b5b6108d88b828c0161080a565b93505060c089013567ffffffffffffffff8111156108f9576108f8610620565b5b6109058b828c0161080a565b92505060e06109168b828c01610649565b9150509295985092959890939650565b61092f8161067e565b82525050565b600060408201905061094a6000830185610926565b6109576020830184610926565b9392505050565b600080604083850312156109755761097461061b565b5b6000610983858286016106a7565b9250506020610994858286016106a7565b9150509250929050565b6000602082840312156109b4576109b361061b565b5b60006109c284828501610649565b91505092915050565b6109d4816106bc565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a145780820151818401526020810190506109f9565b60008484015250505050565b6000610a2b826109da565b610a3581856109e5565b9350610a458185602086016109f6565b610a4e816106fc565b840191505092915050565b610a6281610625565b82525050565b600060e082019050610a7d600083018a610926565b610a8a60208301896109cb565b610a976040830188610926565b610aa46060830187610926565b8181036080830152610ab68186610a20565b905081810360a0830152610aca8185610a20565b9050610ad960c0830184610a59565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b2c57607f821691505b602082108103610b3f57610b3e610ae5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ba77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b6a565b610bb18683610b6a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610bee610be9610be4846106bc565b610bc9565b6106bc565b9050919050565b6000819050919050565b610c0883610bd3565b610c1c610c1482610bf5565b848454610b77565b825550505050565b600090565b610c31610c24565b610c3c818484610bff565b505050565b5b81811015610c6057610c55600082610c29565b600181019050610c42565b5050565b601f821115610ca557610c7681610b45565b610c7f84610b5a565b81016020851015610c8e578190505b610ca2610c9a85610b5a565b830182610c41565b50505b505050565b600082821c905092915050565b6000610cc860001984600802610caa565b1980831691505092915050565b6000610ce18383610cb7565b9150826002028217905092915050565b610cfa826109da565b67ffffffffffffffff811115610d1357610d1261070d565b5b610d1d8254610b14565b610d28828285610c64565b600060209050601f831160018114610d5b5760008415610d49578287015190505b610d538582610cd5565b865550610dbb565b601f198416610d6986610b45565b60005b82811015610d9157848901518255600182019150602085019450602081019050610d6c565b86831015610dae5784890151610daa601f891682610cb7565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220eb13174110973a89d5ab24679a31a442f18d5415623dd287dc7156264ceffd8864736f6c63430008100033",
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
