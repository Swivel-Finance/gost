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

// SwivelMetaData contains all meta data concerning the Swivel contract.
var SwivelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"authRedeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeemReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506105ea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806349be053c146100465780636dd086cc146100625780638bfa66be14610095575b600080fd5b610060600480360381019061005b91906103ca565b6100c5565b005b61007c60048036038101906100779190610430565b6100cf565b60405161008c94939291906104ad565b60405180910390f35b6100af60048036038101906100aa919061051e565b61015f565b6040516100bc9190610599565b60405180910390f35b8060018190555050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610169610325565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082816060018181525050806000808960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505060015491505095945050505050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b6000819050919050565b6103a781610394565b81146103b257600080fd5b50565b6000813590506103c48161039e565b92915050565b6000602082840312156103e0576103df61038f565b5b60006103ee848285016103b5565b91505092915050565b600060ff82169050919050565b61040d816103f7565b811461041857600080fd5b50565b60008135905061042a81610404565b92915050565b6000602082840312156104465761044561038f565b5b60006104548482850161041b565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104888261045d565b9050919050565b6104988161047d565b82525050565b6104a781610394565b82525050565b60006080820190506104c2600083018761048f565b6104cf602083018661048f565b6104dc604083018561048f565b6104e9606083018461049e565b95945050505050565b6104fb8161047d565b811461050657600080fd5b50565b600081359050610518816104f2565b92915050565b600080600080600060a0868803121561053a5761053961038f565b5b60006105488882890161041b565b955050602061055988828901610509565b945050604061056a88828901610509565b935050606061057b88828901610509565b925050608061058c888289016103b5565b9150509295509295909350565b60006020820190506105ae600083018461049e565b9291505056fea2646970667358221220f2aab749b56e0893530d9eb12d29694f2797b876835d3e19c504ca01fa4b8fde64736f6c634300080d0033",
}

// SwivelABI is the input ABI used to generate the binding from.
// Deprecated: Use SwivelMetaData.ABI instead.
var SwivelABI = SwivelMetaData.ABI

// SwivelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwivelMetaData.Bin instead.
var SwivelBin = SwivelMetaData.Bin

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := SwivelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwivelBin), backend)
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

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, address one, address two, uint256 amount)
func (_Swivel *SwivelCaller) AuthRedeemCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "authRedeemCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.One = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, address one, address two, uint256 amount)
func (_Swivel *SwivelSession) AuthRedeemCalled(arg0 uint8) (struct {
	Underlying common.Address
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _Swivel.Contract.AuthRedeemCalled(&_Swivel.CallOpts, arg0)
}

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, address one, address two, uint256 amount)
func (_Swivel *SwivelCallerSession) AuthRedeemCalled(arg0 uint8) (struct {
	Underlying common.Address
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _Swivel.Contract.AuthRedeemCalled(&_Swivel.CallOpts, arg0)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(uint256)
func (_Swivel *SwivelTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "authRedeem", p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(uint256)
func (_Swivel *SwivelSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(uint256)
func (_Swivel *SwivelTransactorSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_Swivel *SwivelTransactor) AuthRedeemReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "authRedeemReturns", a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_Swivel *SwivelSession) AuthRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeemReturns(&_Swivel.TransactOpts, a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_Swivel *SwivelTransactorSession) AuthRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeemReturns(&_Swivel.TransactOpts, a)
}
