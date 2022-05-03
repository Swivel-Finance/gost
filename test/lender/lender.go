// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lender

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

// SwivelComponents is an auto generated low-level Go binding around an user-defined struct.
type SwivelComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelOrder is an auto generated low-level Go binding around an user-defined struct.
type SwivelOrder struct {
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

// LenderABI is the input ABI used to generate the binding from.
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x608060405234801561001057600080fd5b5060405161122b38038061122b83398101604081905261002f91610062565b60008054336001600160a01b031991821617909155600180549091166001600160a01b0392909216919091179055610092565b60006020828403121561007457600080fd5b81516001600160a01b038116811461008b57600080fd5b9392505050565b61118a806100a16000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063ea08c03111610050578063ea08c031146100d7578063f1aa654e146100f7578063f851a4401461010a57600080fd5b80632e25d2a61461006c5780635b1e5fc2146100b6575b600080fd5b60015461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100c96100c4366004610b3f565b61012a565b6040519081526020016100ad565b60025461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b6100c9610105366004610c1b565b610365565b60005461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b60008080805b888110156102a15760008a8a8381811061014c5761014c610c74565b905061012002018036038101906101639190610d15565b90508c8160e00151146101b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1614610224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016101a8565b88888381811061023657610236610c74565b90506020020135846102489190610ddb565b93508060a001518160c001518a8a8581811061026657610266610c74565b905060200201356102779190610df3565b6102819190610e30565b61028b9084610ddb565b925050808061029990610e6b565b915050610130565b508b6102af813330866107c7565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f589061030f908d908d908d908d908d908d90600401610f47565b6020604051808303816000875af115801561032e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103529190611050565b50909d9c50505050505050505050505050565b600080839050848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af11580156103b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103dd9190611074565b63ffffffff161461044a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6d6174757269747920646f6573206e6f74206d6174636800000000000000000060448201526064016101a8565b8530610458823383886107c7565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c61047f886108b1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af11580156104e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050a919061109a565b905061051783888861090b565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af115801561059d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c1919061109a565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b1601610741576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af1158015610667573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068b91906110cc565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af115801561071a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073e9190611050565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af1915050610844816109d8565b6108aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064016101a8565b5050505050565b60006fffffffffffffffffffffffffffffffff821115610907576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016101a8565b5090565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af191505061096c816109d8565b6109d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c6564000000000000000000000000000000000060448201526064016101a8565b50505050565b6000803d836109eb57806000803e806000fd5b8060208114610a03578015610a145760009250610a19565b816000803e60005115159250610a19565b600192505b50909392505050565b803560ff81168114610a3357600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a5a57600080fd5b50565b8035610a3381610a38565b60008083601f840112610a7a57600080fd5b50813567ffffffffffffffff811115610a9257600080fd5b60208301915083602061012083028501011115610aae57600080fd5b9250929050565b60008083601f840112610ac757600080fd5b50813567ffffffffffffffff811115610adf57600080fd5b6020830191508360208260051b8501011115610aae57600080fd5b60008083601f840112610b0c57600080fd5b50813567ffffffffffffffff811115610b2457600080fd5b602083019150836020606083028501011115610aae57600080fd5b60008060008060008060008060008060e08b8d031215610b5e57600080fd5b610b678b610a22565b995060208b0135610b7781610a38565b985060408b0135975060608b0135610b8e81610a38565b965060808b013567ffffffffffffffff80821115610bab57600080fd5b610bb78e838f01610a68565b909850965060a08d0135915080821115610bd057600080fd5b610bdc8e838f01610ab5565b909650945060c08d0135915080821115610bf557600080fd5b50610c028d828e01610afa565b915080935050809150509295989b9194979a5092959850565b600080600080600060a08688031215610c3357600080fd5b610c3c86610a22565b94506020860135610c4c81610a38565b9350604086013592506060860135610c6381610a38565b949793965091946080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610cf657610cf6610ca3565b60405290565b8015158114610a5a57600080fd5b8035610a3381610cfc565b60006101208284031215610d2857600080fd5b610d30610cd2565b82358152610d4060208401610a5d565b6020820152610d5160408401610a5d565b6040820152610d6260608401610d0a565b6060820152610d7360808401610d0a565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610dee57610dee610dac565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e2b57610e2b610dac565b500290565b600082610e66577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e9c57610e9c610dac565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610ed557600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b85811015610f3c5760ff610f1583610a22565b16875281830135838801526040808301359088015260609687019690910190600101610f02565b509495945050505050565b606080825281810187905260009060808084018a845b8b8110156110175781358352602080830135610f7881610a38565b73ffffffffffffffffffffffffffffffffffffffff16908401526040610f9f838201610a5d565b73ffffffffffffffffffffffffffffffffffffffff1690840152610fc4828601610d0a565b151585840152610fd5828501610d0a565b15158484015260a0828101359084015260c0808301359084015260e0808301359084015261010080830135908401526101209283019290910190600101610f5d565b5050848103602086015261102c81898b610ea3565b925050508281036040840152611043818587610ef2565b9998505050505050505050565b60006020828403121561106257600080fd5b815161106d81610cfc565b9392505050565b60006020828403121561108657600080fd5b815163ffffffff8116811461106d57600080fd5b6000602082840312156110ac57600080fd5b81516fffffffffffffffffffffffffffffffff8116811461106d57600080fd5b60006101008083850312156110e057600080fd5b83601f8401126110ef57600080fd5b60405181810181811067ffffffffffffffff8211171561111157611111610ca3565b60405290830190808583111561112657600080fd5b845b8381101561114957805161113b81610a38565b825260209182019101611128565b50909594505050505056fea26469706673582212206b496851955df7126a58d1e07cf06e3c56dd79ed94ee3f15ec183d761c0f412064736f6c634300080d0033"

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LenderBin), backend, m)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// Lender is an auto generated Go binding around an Ethereum contract.
type Lender struct {
	LenderCaller     // Read-only binding to the contract
	LenderTransactor // Write-only binding to the contract
	LenderFilterer   // Log filterer for contract events
}

// LenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LenderSession struct {
	Contract     *Lender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LenderCallerSession struct {
	Contract *LenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LenderTransactorSession struct {
	Contract     *LenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LenderRaw struct {
	Contract *Lender // Generic contract binding to access the raw methods on
}

// LenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LenderCallerRaw struct {
	Contract *LenderCaller // Generic read-only contract binding to access the raw methods on
}

// LenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LenderTransactorRaw struct {
	Contract *LenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLender creates a new instance of Lender, bound to a specific deployed contract.
func NewLender(address common.Address, backend bind.ContractBackend) (*Lender, error) {
	contract, err := bindLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// NewLenderCaller creates a new read-only instance of Lender, bound to a specific deployed contract.
func NewLenderCaller(address common.Address, caller bind.ContractCaller) (*LenderCaller, error) {
	contract, err := bindLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LenderCaller{contract: contract}, nil
}

// NewLenderTransactor creates a new write-only instance of Lender, bound to a specific deployed contract.
func NewLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*LenderTransactor, error) {
	contract, err := bindLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LenderTransactor{contract: contract}, nil
}

// NewLenderFilterer creates a new log filterer instance of Lender, bound to a specific deployed contract.
func NewLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*LenderFilterer, error) {
	contract, err := bindLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LenderFilterer{contract: contract}, nil
}

// bindLender binds a generic wrapper to an already deployed contract.
func bindLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.LenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderSession) Admin() (common.Address, error) {
	return _Lender.Contract.Admin(&_Lender.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderCallerSession) Admin() (common.Address, error) {
	return _Lender.Contract.Admin(&_Lender.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderSession) MarketPlace() (common.Address, error) {
	return _Lender.Contract.MarketPlace(&_Lender.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderCallerSession) MarketPlace() (common.Address, error) {
	return _Lender.Contract.MarketPlace(&_Lender.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderCaller) SwivelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "swivelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderSession) SwivelAddr() (common.Address, error) {
	return _Lender.Contract.SwivelAddr(&_Lender.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderCallerSession) SwivelAddr() (common.Address, error) {
	return _Lender.Contract.SwivelAddr(&_Lender.CallOpts)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactor) Lend(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend", p, u, m, y, o, a, s)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderSession) Lend(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactorSession) Lend(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend0 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactor) Lend0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend0", p, u, m, y, a)
}

// Lend0 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderSession) Lend0(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, y, a)
}

// Lend0 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactorSession) Lend0(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, y, a)
}

// LenderLendIterator is returned from FilterLend and is used to iterate over the raw logs and unpacked data for Lend events raised by the Lender contract.
type LenderLendIterator struct {
	Event *LenderLend // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LenderLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LenderLend)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LenderLend)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LenderLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LenderLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LenderLend represents a Lend event raised by the Lender contract.
type LenderLend struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Returned   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLend is a free log retrieval operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) FilterLend(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*LenderLendIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.FilterLogs(opts, "Lend", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &LenderLendIterator{contract: _Lender.contract, event: "Lend", logs: logs, sub: sub}, nil
}

// WatchLend is a free log subscription operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) WatchLend(opts *bind.WatchOpts, sink chan<- *LenderLend, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.WatchLogs(opts, "Lend", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LenderLend)
				if err := _Lender.contract.UnpackLog(event, "Lend", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLend is a log parse operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) ParseLend(log types.Log) (*LenderLend, error) {
	event := new(LenderLend)
	if err := _Lender.contract.UnpackLog(event, "Lend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
