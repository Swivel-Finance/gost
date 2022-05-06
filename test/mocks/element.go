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

// FundManagement is an auto generated low-level Go binding around an user-defined struct.
type FundManagement struct {
	Sender              common.Address
	Recipient           common.Address
	FromInternalBalance bool
	ToInternalBalance   bool
}

// SingleSwap is an auto generated low-level Go binding around an user-defined struct.
type SingleSwap struct {
	UserData []byte
	PoolId   [32]byte
	Amount   *big.Int
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
}

// ElementABI is the input ABI used to generate the binding from.
const ElementABI = "[{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundManagementSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"return_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"singleSwapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"f\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ElementBin is the compiled bytecode used for deploying new contracts.
var ElementBin = "0x608060405234801561001057600080fd5b506106ff806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806328386b831461005c57806329dcb0cf1461007a57806369bbd8cd1461009857806391d51144146100c8578063a4b15a5d146100e6575b600080fd5b610064610104565b60405161007191906101f8565b60405180910390f35b61008261012e565b60405161008f919061022c565b60405180910390f35b6100b260048036038101906100ad9190610646565b610137565b6040516100bf919061022c565b60405180910390f35b6100d06101a3565b6040516100dd919061022c565b60405180910390f35b6100ee6101ad565b6040516100fb919061022c565b60405180910390f35b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054905090565b600082600181905550816000819055508360000151600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846040015160038190555060009050949350505050565b6000600154905090565b6000600354905090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101e2826101b7565b9050919050565b6101f2816101d7565b82525050565b600060208201905061020d60008301846101e9565b92915050565b6000819050919050565b61022681610213565b82525050565b6000602082019050610241600083018461021d565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102a982610260565b810181811067ffffffffffffffff821117156102c8576102c7610271565b5b80604052505050565b60006102db610247565b90506102e782826102a0565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82111561031657610315610271565b5b61031f82610260565b9050602081019050919050565b82818337600083830152505050565b600061034e610349846102fb565b6102d1565b90508281526020810184848401111561036a576103696102f6565b5b61037584828561032c565b509392505050565b600082601f830112610392576103916102f1565b5b81356103a284826020860161033b565b91505092915050565b6000819050919050565b6103be816103ab565b81146103c957600080fd5b50565b6000813590506103db816103b5565b92915050565b6103ea81610213565b81146103f557600080fd5b50565b600081359050610407816103e1565b92915050565b6002811061041a57600080fd5b50565b60008135905061042c8161040d565b92915050565b600061043d826101d7565b9050919050565b61044d81610432565b811461045857600080fd5b50565b60008135905061046a81610444565b92915050565b600060c082840312156104865761048561025b565b5b61049060c06102d1565b9050600082013567ffffffffffffffff8111156104b0576104af6102ec565b5b6104bc8482850161037d565b60008301525060206104d0848285016103cc565b60208301525060406104e4848285016103f8565b60408301525060606104f88482850161041d565b606083015250608061050c8482850161045b565b60808301525060a06105208482850161045b565b60a08301525092915050565b610535816101d7565b811461054057600080fd5b50565b6000813590506105528161052c565b92915050565b6000610563826101b7565b9050919050565b61057381610558565b811461057e57600080fd5b50565b6000813590506105908161056a565b92915050565b60008115159050919050565b6105ab81610596565b81146105b657600080fd5b50565b6000813590506105c8816105a2565b92915050565b6000608082840312156105e4576105e361025b565b5b6105ee60806102d1565b905060006105fe84828501610543565b600083015250602061061284828501610581565b6020830152506040610626848285016105b9565b604083015250606061063a848285016105b9565b60608301525092915050565b60008060008060e085870312156106605761065f610251565b5b600085013567ffffffffffffffff81111561067e5761067d610256565b5b61068a87828801610470565b945050602061069b878288016105ce565b93505060a06106ac878288016103f8565b92505060c06106bd878288016103f8565b9150509295919450925056fea26469706673582212205a9ceaf596980dd58f9d9486f0aa3f1b4f5d1a3f8eb54b05283de554586f6e7564736f6c634300080d0033"

// DeployElement deploys a new Ethereum contract, binding an instance of Element to it.
func DeployElement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Element, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ElementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Element{ElementCaller: ElementCaller{contract: contract}, ElementTransactor: ElementTransactor{contract: contract}, ElementFilterer: ElementFilterer{contract: contract}}, nil
}

// Element is an auto generated Go binding around an Ethereum contract.
type Element struct {
	ElementCaller     // Read-only binding to the contract
	ElementTransactor // Write-only binding to the contract
	ElementFilterer   // Log filterer for contract events
}

// ElementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElementSession struct {
	Contract     *Element          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElementCallerSession struct {
	Contract *ElementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ElementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElementTransactorSession struct {
	Contract     *ElementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ElementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElementRaw struct {
	Contract *Element // Generic contract binding to access the raw methods on
}

// ElementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElementCallerRaw struct {
	Contract *ElementCaller // Generic read-only contract binding to access the raw methods on
}

// ElementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElementTransactorRaw struct {
	Contract *ElementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElement creates a new instance of Element, bound to a specific deployed contract.
func NewElement(address common.Address, backend bind.ContractBackend) (*Element, error) {
	contract, err := bindElement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Element{ElementCaller: ElementCaller{contract: contract}, ElementTransactor: ElementTransactor{contract: contract}, ElementFilterer: ElementFilterer{contract: contract}}, nil
}

// NewElementCaller creates a new read-only instance of Element, bound to a specific deployed contract.
func NewElementCaller(address common.Address, caller bind.ContractCaller) (*ElementCaller, error) {
	contract, err := bindElement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElementCaller{contract: contract}, nil
}

// NewElementTransactor creates a new write-only instance of Element, bound to a specific deployed contract.
func NewElementTransactor(address common.Address, transactor bind.ContractTransactor) (*ElementTransactor, error) {
	contract, err := bindElement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElementTransactor{contract: contract}, nil
}

// NewElementFilterer creates a new log filterer instance of Element, bound to a specific deployed contract.
func NewElementFilterer(address common.Address, filterer bind.ContractFilterer) (*ElementFilterer, error) {
	contract, err := bindElement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElementFilterer{contract: contract}, nil
}

// bindElement binds a generic wrapper to an already deployed contract.
func bindElement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Element *ElementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Element.Contract.ElementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Element *ElementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Element.Contract.ElementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Element *ElementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Element.Contract.ElementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Element *ElementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Element.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Element *ElementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Element.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Element *ElementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Element.Contract.contract.Transact(opts, method, params...)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementSession) Deadline() (*big.Int, error) {
	return _Element.Contract.Deadline(&_Element.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_Element *ElementCallerSession) Deadline() (*big.Int, error) {
	return _Element.Contract.Deadline(&_Element.CallOpts)
}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementCaller) FundManagementSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "fundManagementSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementSession) FundManagementSender() (common.Address, error) {
	return _Element.Contract.FundManagementSender(&_Element.CallOpts)
}

// FundManagementSender is a free data retrieval call binding the contract method 0x28386b83.
//
// Solidity: function fundManagementSender() view returns(address)
func (_Element *ElementCallerSession) FundManagementSender() (common.Address, error) {
	return _Element.Contract.FundManagementSender(&_Element.CallOpts)
}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementCaller) Return(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "return_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementSession) Return() (*big.Int, error) {
	return _Element.Contract.Return(&_Element.CallOpts)
}

// Return is a free data retrieval call binding the contract method 0x91d51144.
//
// Solidity: function return_() view returns(uint256)
func (_Element *ElementCallerSession) Return() (*big.Int, error) {
	return _Element.Contract.Return(&_Element.CallOpts)
}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementCaller) SingleSwapAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "singleSwapAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementSession) SingleSwapAmount() (*big.Int, error) {
	return _Element.Contract.SingleSwapAmount(&_Element.CallOpts)
}

// SingleSwapAmount is a free data retrieval call binding the contract method 0xa4b15a5d.
//
// Solidity: function singleSwapAmount() view returns(uint256)
func (_Element *ElementCallerSession) SingleSwapAmount() (*big.Int, error) {
	return _Element.Contract.SingleSwapAmount(&_Element.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementTransactor) Swap(opts *bind.TransactOpts, s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "swap", s, f, r, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementSession) Swap(s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, r, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 r, uint256 d) returns(uint256)
func (_Element *ElementTransactorSession) Swap(s SingleSwap, f FundManagement, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, r, d)
}
