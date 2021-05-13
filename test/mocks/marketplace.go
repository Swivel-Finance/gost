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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFromNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFromZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromZcTokenCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b506110c3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80634bc81e0911610097578063d557ee8511610066578063d557ee85146102da578063dc098af9146102f6578063f3d9b98414610312578063fb1a0f5814610342576100f5565b80634bc81e091461023f57806389de80aa1461025b578063b80ce9311461028e578063c129a1c2146102aa576100f5565b806328782788116100d3578063287827881461017957806330e5b1b9146101ac5780633a5e2b75146101df5780633f73df2c1461020f576100f5565b806303f5e400146100fa57806305e1dc251461012d57806319178e711461015d575b600080fd5b610114600480360381019061010f9190610e38565b610372565b6040516101249493929190610fbb565b60405180910390f35b61014760048036038101906101429190610e61565b6103e2565b6040516101549190610f6a565b60405180910390f35b61017760048036038101906101729190610f14565b610451565b005b610193600480360381019061018e9190610e38565b61046e565b6040516101a39493929190610fbb565b60405180910390f35b6101c660048036038101906101c19190610e38565b6104de565b6040516101d69493929190610fbb565b60405180910390f35b6101f960048036038101906101f49190610e9d565b61054e565b6040516102069190610f85565b60405180910390f35b61022960048036038101906102249190610e38565b6106dd565b6040516102369190610fa0565b60405180910390f35b61025960048036038101906102549190610f14565b6106f5565b005b61027560048036038101906102709190610e38565b610712565b6040516102859493929190610fbb565b60405180910390f35b6102a860048036038101906102a39190610f14565b610782565b005b6102c460048036038101906102bf9190610e9d565b61079e565b6040516102d19190610f85565b60405180910390f35b6102f460048036038101906102ef9190610e38565b61092b565b005b610310600480360381019061030b9190610f14565b61096e565b005b61032c60048036038101906103279190610e9d565b61098b565b6040516103399190610f85565b60405180910390f35b61035c60048036038101906103579190610e9d565b610b1a565b6040516103699190610f85565b60405180910390f35b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b80600260036101000a81548160ff02191690831515021790555050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610558610ca9565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260019054906101000a900460ff1691505095945050505050565b60016020528060005260406000206000915090505481565b80600260006101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b806002806101000a81548160ff02191690831515021790555050565b60006107a8610cfd565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505060028054906101000a900460ff1691505095945050505050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600260016101000a81548160ff02191690831515021790555050565b6000610995610d51565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260009054906101000a900460ff1691505095945050505050565b6000610b24610da5565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600260039054906101000a900460ff1691505095945050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050610e0881611048565b92915050565b600081359050610e1d8161105f565b92915050565b600081359050610e3281611076565b92915050565b600060208284031215610e4a57600080fd5b6000610e5884828501610df9565b91505092915050565b60008060408385031215610e7457600080fd5b6000610e8285828601610df9565b9250506020610e9385828601610e23565b9150509250929050565b600080600080600060a08688031215610eb557600080fd5b6000610ec388828901610df9565b9550506020610ed488828901610e23565b9450506040610ee588828901610df9565b9350506060610ef688828901610df9565b9250506080610f0788828901610e23565b9150509295509295909350565b600060208284031215610f2657600080fd5b6000610f3484828501610e0e565b91505092915050565b610f4681611000565b82525050565b610f5581611012565b82525050565b610f648161103e565b82525050565b6000602082019050610f7f6000830184610f3d565b92915050565b6000602082019050610f9a6000830184610f4c565b92915050565b6000602082019050610fb56000830184610f5b565b92915050565b6000608082019050610fd06000830187610f5b565b610fdd6020830186610f3d565b610fea6040830185610f3d565b610ff76060830184610f5b565b95945050505050565b600061100b8261101e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61105181611000565b811461105c57600080fd5b50565b61106881611012565b811461107357600080fd5b50565b61107f8161103e565b811461108a57600080fd5b5056fea2646970667358221220dbe72adca4aa7772be981380c2437c260cda4b9b637fa5e34b6bbab5c895cb0064736f6c63430008000033"

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

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) BurnZcTokenRemovingNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "burnZcTokenRemovingNotionalCalled", arg0)

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

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) BurnZcTokenRemovingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) CTokenAddressCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAddressCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
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

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferFromNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferFromNotionalCalled", arg0)

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

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferFromNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromNotionalCalled is a free data retrieval call binding the contract method 0x30e5b1b9.
//
// Solidity: function transferFromNotionalCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferFromNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferFromZcTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferFromZcTokenCalled", arg0)

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

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferFromZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferFromZcTokenCalled is a free data retrieval call binding the contract method 0x28782788.
//
// Solidity: function transferFromZcTokenCalled(address ) view returns(uint256 maturity, address owner, address sender, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferFromZcTokenCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	Owner    common.Address
	Sender   common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferFromZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x3a5e2b75.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotional", u, m, o, s, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x3a5e2b75.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x3a5e2b75.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotionalReturns", b)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) CTokenAddress(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddress", u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactor) CTokenAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddressReturns", a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
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

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferFromNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromNotional", u, m, o, s, a)
}

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferFromNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromNotional is a paid mutator transaction binding the contract method 0xfb1a0f58.
//
// Solidity: function transferFromNotional(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromNotional(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotional(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferFromNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromNotionalReturns", b)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferFromNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromNotionalReturns is a paid mutator transaction binding the contract method 0x19178e71.
//
// Solidity: function transferFromNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferFromZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromZcToken", u, m, o, s, a)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferFromZcToken(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcToken(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromZcToken is a paid mutator transaction binding the contract method 0xc129a1c2.
//
// Solidity: function transferFromZcToken(address u, uint256 m, address o, address s, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromZcToken(u common.Address, m *big.Int, o common.Address, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcToken(&_MarketPlace.TransactOpts, u, m, o, s, a)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferFromZcTokenReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferFromZcTokenReturns", b)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferFromZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcTokenReturns(&_MarketPlace.TransactOpts, b)
}

// TransferFromZcTokenReturns is a paid mutator transaction binding the contract method 0xb80ce931.
//
// Solidity: function transferFromZcTokenReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferFromZcTokenReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferFromZcTokenReturns(&_MarketPlace.TransactOpts, b)
}
