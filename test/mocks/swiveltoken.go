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

// Components is an auto generated low-level Go binding around an user-defined struct.
type Components struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
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

// SwivelTokenABI is the input ABI used to generate the binding from.
const SwivelTokenABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structComponents[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"i\",\"type\":\"bool\"}],\"name\":\"initiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemZcTokenCalledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemZcTokenCalledToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"r\",\"type\":\"bool\"}],\"name\":\"redeemZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwivelTokenBin is the compiled bytecode used for deploying new contracts.
var SwivelTokenBin = "0x608060405234801561001057600080fd5b506109a8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639371d5781161005b5780639371d57814610139578063bef14a3214610169578063c268dc2a14610185578063d2144f58146101b557610088565b8063154e0f2e1461008d57806338599fc8146100bd5780634578b3ce146100ed57806384f5544714610109575b600080fd5b6100a760048036038101906100a2919061053d565b6101e5565b6040516100b491906105ab565b60405180910390f35b6100d760048036038101906100d291906105c6565b610288565b6040516100e4919061060f565b60405180910390f35b61010760048036038101906101029190610656565b6102a8565b005b610123600480360381019061011e91906105c6565b6102c4565b6040516101309190610692565b60405180910390f35b610153600480360381019061014e91906105c6565b6102dc565b6040516101609190610692565b60405180910390f35b610183600480360381019061017e9190610656565b6102f4565b005b61019f600480360381019061019a91906105c6565b610311565b6040516101ac9190610692565b60405180910390f35b6101cf60048036038101906101ca91906107bf565b610329565b6040516101dc91906105ab565b60405180910390f35b600082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600060019054906101000a900460ff1690509392505050565b60026020528060005260406000206000915054906101000a900460ff1681565b806000806101000a81548160ff02191690831515021790555050565b60046020528060005260406000206000915090505481565b60036020528060005260406000206000915090505481565b80600060016101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600080600090505b878790508110156104845785858281811061034f5761034e610873565b5b90506020020135600160008a8a8581811061036d5761036c610873565b5b90506101200201602001602081019061038691906105c6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508383828181106103d7576103d6610873565b5b90506060020160000160208101906103ef91906108ce565b600260008a8a8581811061040657610405610873565b5b90506101200201602001602081019061041f91906105c6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550808061047c9061092a565b915050610331565b5060008054906101000a900460ff1690509695505050505050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104d4826104a9565b9050919050565b6104e4816104c9565b81146104ef57600080fd5b50565b600081359050610501816104db565b92915050565b6000819050919050565b61051a81610507565b811461052557600080fd5b50565b60008135905061053781610511565b92915050565b6000806000606084860312156105565761055561049f565b5b6000610564868287016104f2565b935050602061057586828701610528565b925050604061058686828701610528565b9150509250925092565b60008115159050919050565b6105a581610590565b82525050565b60006020820190506105c0600083018461059c565b92915050565b6000602082840312156105dc576105db61049f565b5b60006105ea848285016104f2565b91505092915050565b600060ff82169050919050565b610609816105f3565b82525050565b60006020820190506106246000830184610600565b92915050565b61063381610590565b811461063e57600080fd5b50565b6000813590506106508161062a565b92915050565b60006020828403121561066c5761066b61049f565b5b600061067a84828501610641565b91505092915050565b61068c81610507565b82525050565b60006020820190506106a76000830184610683565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126106d2576106d16106ad565b5b8235905067ffffffffffffffff8111156106ef576106ee6106b2565b5b6020830191508361012082028301111561070c5761070b6106b7565b5b9250929050565b60008083601f840112610729576107286106ad565b5b8235905067ffffffffffffffff811115610746576107456106b2565b5b602083019150836020820283011115610762576107616106b7565b5b9250929050565b60008083601f84011261077f5761077e6106ad565b5b8235905067ffffffffffffffff81111561079c5761079b6106b2565b5b6020830191508360608202830111156107b8576107b76106b7565b5b9250929050565b600080600080600080606087890312156107dc576107db61049f565b5b600087013567ffffffffffffffff8111156107fa576107f96104a4565b5b61080689828a016106bc565b9650965050602087013567ffffffffffffffff811115610829576108286104a4565b5b61083589828a01610713565b9450945050604087013567ffffffffffffffff811115610858576108576104a4565b5b61086489828a01610769565b92509250509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6108ab816105f3565b81146108b657600080fd5b50565b6000813590506108c8816108a2565b92915050565b6000602082840312156108e4576108e361049f565b5b60006108f2848285016108b9565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061093582610507565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610967576109666108fb565b5b60018201905091905056fea2646970667358221220de036c5ccf502f39f0a0bbc3df9adc1351d371e7f33234b0f40bb66446c9f5a264736f6c634300080d0033"

// DeploySwivelToken deploys a new Ethereum contract, binding an instance of SwivelToken to it.
func DeploySwivelToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwivelToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwivelToken{SwivelTokenCaller: SwivelTokenCaller{contract: contract}, SwivelTokenTransactor: SwivelTokenTransactor{contract: contract}, SwivelTokenFilterer: SwivelTokenFilterer{contract: contract}}, nil
}

// SwivelToken is an auto generated Go binding around an Ethereum contract.
type SwivelToken struct {
	SwivelTokenCaller     // Read-only binding to the contract
	SwivelTokenTransactor // Write-only binding to the contract
	SwivelTokenFilterer   // Log filterer for contract events
}

// SwivelTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelTokenSession struct {
	Contract     *SwivelToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelTokenCallerSession struct {
	Contract *SwivelTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwivelTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTokenTransactorSession struct {
	Contract     *SwivelTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwivelTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelTokenRaw struct {
	Contract *SwivelToken // Generic contract binding to access the raw methods on
}

// SwivelTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelTokenCallerRaw struct {
	Contract *SwivelTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTokenTransactorRaw struct {
	Contract *SwivelTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivelToken creates a new instance of SwivelToken, bound to a specific deployed contract.
func NewSwivelToken(address common.Address, backend bind.ContractBackend) (*SwivelToken, error) {
	contract, err := bindSwivelToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwivelToken{SwivelTokenCaller: SwivelTokenCaller{contract: contract}, SwivelTokenTransactor: SwivelTokenTransactor{contract: contract}, SwivelTokenFilterer: SwivelTokenFilterer{contract: contract}}, nil
}

// NewSwivelTokenCaller creates a new read-only instance of SwivelToken, bound to a specific deployed contract.
func NewSwivelTokenCaller(address common.Address, caller bind.ContractCaller) (*SwivelTokenCaller, error) {
	contract, err := bindSwivelToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTokenCaller{contract: contract}, nil
}

// NewSwivelTokenTransactor creates a new write-only instance of SwivelToken, bound to a specific deployed contract.
func NewSwivelTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTokenTransactor, error) {
	contract, err := bindSwivelToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTokenTransactor{contract: contract}, nil
}

// NewSwivelTokenFilterer creates a new log filterer instance of SwivelToken, bound to a specific deployed contract.
func NewSwivelTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelTokenFilterer, error) {
	contract, err := bindSwivelToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelTokenFilterer{contract: contract}, nil
}

// bindSwivelToken binds a generic wrapper to an already deployed contract.
func bindSwivelToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwivelToken *SwivelTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwivelToken.Contract.SwivelTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwivelToken *SwivelTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwivelToken.Contract.SwivelTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwivelToken *SwivelTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwivelToken.Contract.SwivelTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwivelToken *SwivelTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwivelToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwivelToken *SwivelTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwivelToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwivelToken *SwivelTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwivelToken.Contract.contract.Transact(opts, method, params...)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCaller) InitiateCalledAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwivelToken.contract.Call(opts, &out, "initiateCalledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.InitiateCalledAmount(&_SwivelToken.CallOpts, arg0)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCallerSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.InitiateCalledAmount(&_SwivelToken.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_SwivelToken *SwivelTokenCaller) InitiateCalledSignature(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _SwivelToken.contract.Call(opts, &out, "initiateCalledSignature", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_SwivelToken *SwivelTokenSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _SwivelToken.Contract.InitiateCalledSignature(&_SwivelToken.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_SwivelToken *SwivelTokenCallerSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _SwivelToken.Contract.InitiateCalledSignature(&_SwivelToken.CallOpts, arg0)
}

// RedeemZcTokenCalledAmount is a free data retrieval call binding the contract method 0x9371d578.
//
// Solidity: function redeemZcTokenCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCaller) RedeemZcTokenCalledAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwivelToken.contract.Call(opts, &out, "redeemZcTokenCalledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemZcTokenCalledAmount is a free data retrieval call binding the contract method 0x9371d578.
//
// Solidity: function redeemZcTokenCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenSession) RedeemZcTokenCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.RedeemZcTokenCalledAmount(&_SwivelToken.CallOpts, arg0)
}

// RedeemZcTokenCalledAmount is a free data retrieval call binding the contract method 0x9371d578.
//
// Solidity: function redeemZcTokenCalledAmount(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCallerSession) RedeemZcTokenCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.RedeemZcTokenCalledAmount(&_SwivelToken.CallOpts, arg0)
}

// RedeemZcTokenCalledToken is a free data retrieval call binding the contract method 0x84f55447.
//
// Solidity: function redeemZcTokenCalledToken(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCaller) RedeemZcTokenCalledToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwivelToken.contract.Call(opts, &out, "redeemZcTokenCalledToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemZcTokenCalledToken is a free data retrieval call binding the contract method 0x84f55447.
//
// Solidity: function redeemZcTokenCalledToken(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenSession) RedeemZcTokenCalledToken(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.RedeemZcTokenCalledToken(&_SwivelToken.CallOpts, arg0)
}

// RedeemZcTokenCalledToken is a free data retrieval call binding the contract method 0x84f55447.
//
// Solidity: function redeemZcTokenCalledToken(address ) view returns(uint256)
func (_SwivelToken *SwivelTokenCallerSession) RedeemZcTokenCalledToken(arg0 common.Address) (*big.Int, error) {
	return _SwivelToken.Contract.RedeemZcTokenCalledToken(&_SwivelToken.CallOpts, arg0)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_SwivelToken *SwivelTokenTransactor) Initiate(opts *bind.TransactOpts, o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _SwivelToken.contract.Transact(opts, "initiate", o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_SwivelToken *SwivelTokenSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _SwivelToken.Contract.Initiate(&_SwivelToken.TransactOpts, o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_SwivelToken *SwivelTokenTransactorSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _SwivelToken.Contract.Initiate(&_SwivelToken.TransactOpts, o, a, s)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_SwivelToken *SwivelTokenTransactor) InitiateReturns(opts *bind.TransactOpts, i bool) (*types.Transaction, error) {
	return _SwivelToken.contract.Transact(opts, "initiateReturns", i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_SwivelToken *SwivelTokenSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _SwivelToken.Contract.InitiateReturns(&_SwivelToken.TransactOpts, i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_SwivelToken *SwivelTokenTransactorSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _SwivelToken.Contract.InitiateReturns(&_SwivelToken.TransactOpts, i)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address a, uint256 u, uint256 t) returns(bool)
func (_SwivelToken *SwivelTokenTransactor) RedeemZcToken(opts *bind.TransactOpts, a common.Address, u *big.Int, t *big.Int) (*types.Transaction, error) {
	return _SwivelToken.contract.Transact(opts, "redeemZcToken", a, u, t)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address a, uint256 u, uint256 t) returns(bool)
func (_SwivelToken *SwivelTokenSession) RedeemZcToken(a common.Address, u *big.Int, t *big.Int) (*types.Transaction, error) {
	return _SwivelToken.Contract.RedeemZcToken(&_SwivelToken.TransactOpts, a, u, t)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address a, uint256 u, uint256 t) returns(bool)
func (_SwivelToken *SwivelTokenTransactorSession) RedeemZcToken(a common.Address, u *big.Int, t *big.Int) (*types.Transaction, error) {
	return _SwivelToken.Contract.RedeemZcToken(&_SwivelToken.TransactOpts, a, u, t)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool r) returns()
func (_SwivelToken *SwivelTokenTransactor) RedeemZcTokenReturns(opts *bind.TransactOpts, r bool) (*types.Transaction, error) {
	return _SwivelToken.contract.Transact(opts, "redeemZcTokenReturns", r)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool r) returns()
func (_SwivelToken *SwivelTokenSession) RedeemZcTokenReturns(r bool) (*types.Transaction, error) {
	return _SwivelToken.Contract.RedeemZcTokenReturns(&_SwivelToken.TransactOpts, r)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool r) returns()
func (_SwivelToken *SwivelTokenTransactorSession) RedeemZcTokenReturns(r bool) (*types.Transaction, error) {
	return _SwivelToken.Contract.RedeemZcTokenReturns(&_SwivelToken.TransactOpts, r)
}
