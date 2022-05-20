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

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structComponents[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"i\",\"type\":\"bool\"}],\"name\":\"initiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemZcTokenCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"a\",\"type\":\"bool\"}],\"name\":\"redeemZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x608060405234801561001057600080fd5b5061096e806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063bef14a321161005b578063bef14a32146100fe578063c268dc2a1461011a578063d2144f581461014a578063ddcfbbda1461017a5761007d565b8063154e0f2e1461008257806338599fc8146100b25780634578b3ce146100e2575b600080fd5b61009c600480360381019061009791906104da565b6101ab565b6040516100a99190610548565b60405180910390f35b6100cc60048036038101906100c79190610563565b610231565b6040516100d991906105ac565b60405180910390f35b6100fc60048036038101906100f791906105f3565b610251565b005b610118600480360381019061011391906105f3565b61026d565b005b610134600480360381019061012f9190610563565b61028a565b604051610141919061062f565b60405180910390f35b610164600480360381019061015f919061075c565b6102a2565b6040516101719190610548565b60405180910390f35b610194600480360381019061018f9190610563565b610418565b6040516101a2929190610810565b60405180910390f35b6000604051806040016040528083815260200184815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050600060019054906101000a900460ff1690509392505050565b60026020528060005260406000206000915054906101000a900460ff1681565b806000806101000a81548160ff02191690831515021790555050565b80600060016101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600080600090505b878790508110156103fd578585828181106102c8576102c7610839565b5b90506020020135600160008a8a858181106102e6576102e5610839565b5b9050610120020160200160208101906102ff9190610563565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508383828181106103505761034f610839565b5b90506060020160000160208101906103689190610894565b600260008a8a8581811061037f5761037e610839565b5b9050610120020160200160208101906103989190610563565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555080806103f5906108f0565b9150506102aa565b5060008054906101000a900460ff1690509695505050505050565b60036020528060005260406000206000915090508060000154908060010154905082565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061047182610446565b9050919050565b61048181610466565b811461048c57600080fd5b50565b60008135905061049e81610478565b92915050565b6000819050919050565b6104b7816104a4565b81146104c257600080fd5b50565b6000813590506104d4816104ae565b92915050565b6000806000606084860312156104f3576104f261043c565b5b60006105018682870161048f565b9350506020610512868287016104c5565b9250506040610523868287016104c5565b9150509250925092565b60008115159050919050565b6105428161052d565b82525050565b600060208201905061055d6000830184610539565b92915050565b6000602082840312156105795761057861043c565b5b60006105878482850161048f565b91505092915050565b600060ff82169050919050565b6105a681610590565b82525050565b60006020820190506105c1600083018461059d565b92915050565b6105d08161052d565b81146105db57600080fd5b50565b6000813590506105ed816105c7565b92915050565b6000602082840312156106095761060861043c565b5b6000610617848285016105de565b91505092915050565b610629816104a4565b82525050565b60006020820190506106446000830184610620565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f84011261066f5761066e61064a565b5b8235905067ffffffffffffffff81111561068c5761068b61064f565b5b602083019150836101208202830111156106a9576106a8610654565b5b9250929050565b60008083601f8401126106c6576106c561064a565b5b8235905067ffffffffffffffff8111156106e3576106e261064f565b5b6020830191508360208202830111156106ff576106fe610654565b5b9250929050565b60008083601f84011261071c5761071b61064a565b5b8235905067ffffffffffffffff8111156107395761073861064f565b5b60208301915083606082028301111561075557610754610654565b5b9250929050565b600080600080600080606087890312156107795761077861043c565b5b600087013567ffffffffffffffff81111561079757610796610441565b5b6107a389828a01610659565b9650965050602087013567ffffffffffffffff8111156107c6576107c5610441565b5b6107d289828a016106b0565b9450945050604087013567ffffffffffffffff8111156107f5576107f4610441565b5b61080189828a01610706565b92509250509295509295509295565b60006040820190506108256000830185610620565b6108326020830184610620565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b61087181610590565b811461087c57600080fd5b50565b60008135905061088e81610868565b92915050565b6000602082840312156108aa576108a961043c565b5b60006108b88482850161087f565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108fb826104a4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361092d5761092c6108c1565b5b60018201905091905056fea2646970667358221220abfa8e5bab37b78614bd35ca4cd2a2a8f096d5a340209109f16a84c0c863816764736f6c634300080d0033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// Swivel is an auto generated Go binding around an Ethereum contract.
type Swivel struct {
	SwivelCaller     // Read-only binding to the contract
	SwivelTransactor // Write-only binding to the contract
	SwivelFilterer   // Log filterer for contract events
}

// SwivelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelSession struct {
	Contract     *Swivel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelCallerSession struct {
	Contract *SwivelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwivelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTransactorSession struct {
	Contract     *SwivelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelRaw struct {
	Contract *Swivel // Generic contract binding to access the raw methods on
}

// SwivelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelCallerRaw struct {
	Contract *SwivelCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTransactorRaw struct {
	Contract *SwivelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivel creates a new instance of Swivel, bound to a specific deployed contract.
func NewSwivel(address common.Address, backend bind.ContractBackend) (*Swivel, error) {
	contract, err := bindSwivel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// NewSwivelCaller creates a new read-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelCaller(address common.Address, caller bind.ContractCaller) (*SwivelCaller, error) {
	contract, err := bindSwivel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelCaller{contract: contract}, nil
}

// NewSwivelTransactor creates a new write-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTransactor, error) {
	contract, err := bindSwivel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTransactor{contract: contract}, nil
}

// NewSwivelFilterer creates a new log filterer instance of Swivel, bound to a specific deployed contract.
func NewSwivelFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelFilterer, error) {
	contract, err := bindSwivel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelFilterer{contract: contract}, nil
}

// bindSwivel binds a generic wrapper to an already deployed contract.
func bindSwivel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.SwivelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transact(opts, method, params...)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCaller) InitiateCalledAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCallerSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCaller) InitiateCalledSignature(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledSignature", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCallerSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 amount, uint256 maturity)
func (_Swivel *SwivelCaller) RedeemZcTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount   *big.Int
	Maturity *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "redeemZcTokenCalled", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Maturity *big.Int
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.Maturity = out[1].(*big.Int)

	return *outstruct, err

}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 amount, uint256 maturity)
func (_Swivel *SwivelSession) RedeemZcTokenCalled(arg0 common.Address) (struct {
	Amount   *big.Int
	Maturity *big.Int
}, error) {
	return _Swivel.Contract.RedeemZcTokenCalled(&_Swivel.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 amount, uint256 maturity)
func (_Swivel *SwivelCallerSession) RedeemZcTokenCalled(arg0 common.Address) (struct {
	Amount   *big.Int
	Maturity *big.Int
}, error) {
	return _Swivel.Contract.RedeemZcTokenCalled(&_Swivel.CallOpts, arg0)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactor) Initiate(opts *bind.TransactOpts, o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiate", o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactorSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactor) InitiateReturns(opts *bind.TransactOpts, i bool) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiateReturns", i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactorSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) RedeemZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemZcToken", u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, u, m, a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool a) returns()
func (_Swivel *SwivelTransactor) RedeemZcTokenReturns(opts *bind.TransactOpts, a bool) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemZcTokenReturns", a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool a) returns()
func (_Swivel *SwivelSession) RedeemZcTokenReturns(a bool) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcTokenReturns(&_Swivel.TransactOpts, a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0xbef14a32.
//
// Solidity: function redeemZcTokenReturns(bool a) returns()
func (_Swivel *SwivelTransactorSession) RedeemZcTokenReturns(a bool) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcTokenReturns(&_Swivel.TransactOpts, a)
}
