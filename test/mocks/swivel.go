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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"authRedeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"authRedeemReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610651806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806343c808e2146100465780636dd086cc146100625780638bfa66be14610095575b600080fd5b610060600480360381019061005b91906103ec565b6100c5565b005b61007c60048036038101906100779190610452565b6100e2565b60405161008c94939291906104d9565b60405180910390f35b6100af60048036038101906100aa9190610576565b610172565b6040516100bc9190610600565b60405180910390f35b80600160006101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b600061017c610345565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082816060018181525050806000808960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600160009054906101000a900460ff1691505095945050505050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b60008115159050919050565b6103c9816103b4565b81146103d457600080fd5b50565b6000813590506103e6816103c0565b92915050565b600060208284031215610402576104016103af565b5b6000610410848285016103d7565b91505092915050565b600060ff82169050919050565b61042f81610419565b811461043a57600080fd5b50565b60008135905061044c81610426565b92915050565b600060208284031215610468576104676103af565b5b60006104768482850161043d565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104aa8261047f565b9050919050565b6104ba8161049f565b82525050565b6000819050919050565b6104d3816104c0565b82525050565b60006080820190506104ee60008301876104b1565b6104fb60208301866104b1565b61050860408301856104b1565b61051560608301846104ca565b95945050505050565b6105278161049f565b811461053257600080fd5b50565b6000813590506105448161051e565b92915050565b610553816104c0565b811461055e57600080fd5b50565b6000813590506105708161054a565b92915050565b600080600080600060a08688031215610592576105916103af565b5b60006105a08882890161043d565b95505060206105b188828901610535565b94505060406105c288828901610535565b93505060606105d388828901610535565b92505060806105e488828901610561565b9150509295509295909350565b6105fa816103b4565b82525050565b600060208201905061061560008301846105f1565b9291505056fea26469706673582212209216bb63f7d4e417ac47b9245668cba6be856e02ce21d1ac699c783e64ce526864736f6c63430008100033",
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
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "authRedeem", p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x8bfa66be.
//
// Solidity: function authRedeem(uint8 p, address u, address c, address t, uint256 a) returns(bool)
func (_Swivel *SwivelTransactorSession) AuthRedeem(p uint8, u common.Address, c common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeem(&_Swivel.TransactOpts, p, u, c, t, a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x43c808e2.
//
// Solidity: function authRedeemReturns(bool b) returns()
func (_Swivel *SwivelTransactor) AuthRedeemReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "authRedeemReturns", b)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x43c808e2.
//
// Solidity: function authRedeemReturns(bool b) returns()
func (_Swivel *SwivelSession) AuthRedeemReturns(b bool) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeemReturns(&_Swivel.TransactOpts, b)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x43c808e2.
//
// Solidity: function authRedeemReturns(bool b) returns()
func (_Swivel *SwivelTransactorSession) AuthRedeemReturns(b bool) (*types.Transaction, error) {
	return _Swivel.Contract.AuthRedeemReturns(&_Swivel.TransactOpts, b)
}
