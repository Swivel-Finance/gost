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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"marketCTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketCTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"marketCTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50610680806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063364cc9c2146100675780633cc3e567146100835780634bc81e09146100b35780636c83a451146100cf57806371e511e7146100ff57806389de80aa1461012f575b600080fd5b610081600480360381019061007c9190610417565b610161565b005b61009d6004803603810190610098919061047c565b6101a4565b6040516100aa9190610550565b60405180910390f35b6100cd60048036038101906100c891906104df565b6102b3565b005b6100e960048036038101906100e49190610417565b6102d0565b6040516100f6919061056b565b60405180910390f35b61011960048036038101906101149190610440565b6102e8565b6040516101269190610535565b60405180910390f35b61014960048036038101906101449190610417565b610357565b60405161015893929190610586565b60405180910390f35b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006101ae6103a1565b848160000181815250508381602001818152505082816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050600260009054906101000a900460ff16915050949350505050565b80600260006101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b60036020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60405180606001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b6000813590506103e781610605565b92915050565b6000813590506103fc8161061c565b92915050565b60008135905061041181610633565b92915050565b60006020828403121561042957600080fd5b6000610437848285016103d8565b91505092915050565b6000806040838503121561045357600080fd5b6000610461858286016103d8565b925050602061047285828601610402565b9150509250929050565b6000806000806080858703121561049257600080fd5b60006104a0878288016103d8565b94505060206104b187828801610402565b93505060406104c287828801610402565b92505060606104d3878288016103d8565b91505092959194509250565b6000602082840312156104f157600080fd5b60006104ff848285016103ed565b91505092915050565b610511816105bd565b82525050565b610520816105cf565b82525050565b61052f816105fb565b82525050565b600060208201905061054a6000830184610508565b92915050565b60006020820190506105656000830184610517565b92915050565b60006020820190506105806000830184610526565b92915050565b600060608201905061059b6000830186610526565b6105a86020830185610526565b6105b56040830184610508565b949350505050565b60006105c8826105db565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61060e816105bd565b811461061957600080fd5b50565b610625816105cf565b811461063057600080fd5b50565b61063c816105fb565b811461064757600080fd5b5056fea2646970667358221220c26951d0e4c6863dbdcb9ec4a24ce8f0ed7efcb66d54c3d3bdf9f6d8ac6e892764736f6c63430008000033"

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
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, uint256 amount, address owner)
func (_MarketPlace *MarketPlaceCaller) MintZcTokenAddingNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
	Owner    common.Address
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mintZcTokenAddingNotionalCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		Amount   *big.Int
		Owner    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.Amount = out[1].(*big.Int)
	outstruct.Owner = out[2].(common.Address)

	return *outstruct, err

}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, uint256 amount, address owner)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
	Owner    common.Address
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, uint256 amount, address owner)
func (_MarketPlace *MarketPlaceCallerSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Amount   *big.Int
	Owner    common.Address
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
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

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x3cc3e567.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, uint256 a, address o) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, a *big.Int, o common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotional", u, m, a, o)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x3cc3e567.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, uint256 a, address o) returns(bool)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, a *big.Int, o common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, a, o)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x3cc3e567.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, uint256 a, address o) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, a *big.Int, o common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, a, o)
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
