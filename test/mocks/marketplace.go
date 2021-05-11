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

// MarketPlaceABI is the input ABI used to generate the binding from.
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"marketCTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketCTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"marketCTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFromMarketZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromMarketZcTokenCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromMarketZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50610a86806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806389de80aa1161006657806389de80aa14610130578063915b280114610163578063b5e34f1214610196578063f3d9b984146101c6578063f9e35847146101f657610093565b8063364cc9c2146100985780634bc81e09146100b45780636c83a451146100d057806371e511e714610100575b600080fd5b6100b260048036038101906100ad91906107fb565b610212565b005b6100ce60048036038101906100c991906108d7565b610255565b005b6100ea60048036038101906100e591906107fb565b610272565b6040516100f79190610963565b60405180910390f35b61011a60048036038101906101159190610824565b61028a565b604051610127919061092d565b60405180910390f35b61014a600480360381019061014591906107fb565b6102f9565b60405161015a949392919061097e565b60405180910390f35b61017d600480360381019061017891906107fb565b610369565b60405161018d949392919061097e565b60405180910390f35b6101b060048036038101906101ab9190610860565b6103d9565b6040516101bd9190610948565b60405180910390f35b6101e060048036038101906101db9190610860565b610568565b6040516101ed9190610948565b60405180910390f35b610210600480360381019061020b91906108d7565b6106f7565b005b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600260006101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006103e3610714565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260019054906101000a900460ff1691505095945050505050565b6000610572610768565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260009054906101000a900460ff1691505095945050505050565b80600260016101000a81548160ff02191690831515021790555050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6000813590506107cb81610a0b565b92915050565b6000813590506107e081610a22565b92915050565b6000813590506107f581610a39565b92915050565b60006020828403121561080d57600080fd5b600061081b848285016107bc565b91505092915050565b6000806040838503121561083757600080fd5b6000610845858286016107bc565b9250506020610856858286016107e6565b9150509250929050565b600080600080600060a0868803121561087857600080fd5b6000610886888289016107bc565b9550506020610897888289016107e6565b94505060406108a8888289016107bc565b93505060606108b9888289016107bc565b92505060806108ca888289016107e6565b9150509295509295909350565b6000602082840312156108e957600080fd5b60006108f7848285016107d1565b91505092915050565b610909816109c3565b82525050565b610918816109d5565b82525050565b61092781610a01565b82525050565b60006020820190506109426000830184610900565b92915050565b600060208201905061095d600083018461090f565b92915050565b6000602082019050610978600083018461091e565b92915050565b6000608082019050610993600083018761091e565b6109a06020830186610900565b6109ad6040830185610900565b6109ba606083018461091e565b95945050505050565b60006109ce826109e1565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610a14816109c3565b8114610a1f57600080fd5b50565b610a2b816109d5565b8114610a3657600080fd5b50565b610a4281610a01565b8114610a4d57600080fd5b5056fea264697066735822122039afa61c26246f307d0313eb47c8ebf5447ff31f5b1551cba8bb16ca9384606d64736f6c63430008000033"

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketPlaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// MarketCTokenAddressCalled is a free data retrieval call binding the contract method 0x6c83a451.
//
// Solidity: function marketCTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) MarketCTokenAddressCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "marketCTokenAddressCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCTokenAddressCalled is a free data retrieval call binding the contract method 0x6c83a451.
//
// Solidity: function marketCTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) MarketCTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketCTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// MarketCTokenAddressCalled is a free data retrieval call binding the contract method 0x6c83a451.
//
// Solidity: function marketCTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) MarketCTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.MarketCTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) MintZcTokenAddingNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mintZcTokenAddingNotionalCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		Owner    common.Address
		Sender   common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Owner = out[1].(common.Address)
	outstruct.Sender = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromMarketZcTokenCalled is a free data retrieval call binding the contract method 0x915b2801.
//
// Solidity: function transferFromMarketZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferFromMarketZcTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferFromMarketZcTokenCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		Owner    common.Address
		Sender   common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Owner = out[1].(common.Address)
	outstruct.Sender = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// TransferFromMarketZcTokenCalled is a free data retrieval call binding the contract method 0x915b2801.
//
// Solidity: function transferFromMarketZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferFromMarketZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromMarketZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromMarketZcTokenCalled is a free data retrieval call binding the contract method 0x915b2801.
//
// Solidity: function transferFromMarketZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferFromMarketZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromMarketZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// MarketCTokenAddress is a paid mutator transaction binding the contract method 0x71e511e7.
//
// Solidity: function marketCTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) MarketCTokenAddress(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "marketCTokenAddress", u, m)
}

// MarketCTokenAddress is a paid mutator transaction binding the contract method 0x71e511e7.
//
// Solidity: function marketCTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) MarketCTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketCTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// MarketCTokenAddress is a paid mutator transaction binding the contract method 0x71e511e7.
//
// Solidity: function marketCTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) MarketCTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketCTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// MarketCTokenAddressReturns is a paid mutator transaction binding the contract method 0x364cc9c2.
//
// Solidity: function marketCTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactor) MarketCTokenAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "marketCTokenAddressReturns", a)
}

// MarketCTokenAddressReturns is a paid mutator transaction binding the contract method 0x364cc9c2.
//
// Solidity: function marketCTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceSession) MarketCTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketCTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// MarketCTokenAddressReturns is a paid mutator transaction binding the contract method 0x364cc9c2.
//
// Solidity: function marketCTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MarketCTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketCTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xf3d9b984.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotional", u, m, o, s, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xf3d9b984.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xf3d9b984.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotionalReturns", b)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromMarketZcToken is a paid mutator transaction binding the contract method 0xb5e34f12.
//
// Solidity: function transferFromMarketZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferFromMarketZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromMarketZcToken", u, m, o, s, a)
}

// TransferFromMarketZcToken is a paid mutator transaction binding the contract method 0xb5e34f12.
//
// Solidity: function transferFromMarketZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferFromMarketZcToken(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromMarketZcToken(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromMarketZcToken is a paid mutator transaction binding the contract method 0xb5e34f12.
//
// Solidity: function transferFromMarketZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromMarketZcToken(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromMarketZcToken(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromMarketZcTokenReturns is a paid mutator transaction binding the contract method 0xf9e35847.
//
// Solidity: function transferFromMarketZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferFromMarketZcTokenReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromMarketZcTokenReturns", b)
}

// TransferFromMarketZcTokenReturns is a paid mutator transaction binding the contract method 0xf9e35847.
//
// Solidity: function transferFromMarketZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferFromMarketZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromMarketZcTokenReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromMarketZcTokenReturns is a paid mutator transaction binding the contract method 0xf9e35847.
//
// Solidity: function transferFromMarketZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromMarketZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromMarketZcTokenReturns(&_MarketPlace.TransactOpts, b)
}
