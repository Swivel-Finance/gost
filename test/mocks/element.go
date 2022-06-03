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
const ElementABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"f\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"l\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"}],\"name\":\"withdrawPrincipal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawPrincipalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ElementBin is the compiled bytecode used for deploying new contracts.
var ElementBin = "0x608060405234801561001057600080fd5b506108e6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633b1988dc1461005c5780635658724f1461008c57806369bbd8cd146100bf578063884e17f3146100ef578063d94108181461010b575b600080fd5b6100766004803603810190610071919061034c565b610127565b6040516100839190610392565b60405180910390f35b6100a660048036038101906100a1919061034c565b61013f565b6040516100b694939291906103bc565b60405180910390f35b6100d960048036038101906100d491906107c0565b61018f565b6040516100e69190610392565b60405180910390f35b61010960048036038101906101049190610843565b610288565b005b61012560048036038101906101209190610883565b6102d0565b005b60026020528060005260406000206000915090505481565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b60006040518060800160405280856020015173ffffffffffffffffffffffffffffffffffffffff168152602001866040015181526020018481526020018381525060016000866000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301559050506000549050949350505050565b81600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b8060008190555050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610319826102ee565b9050919050565b6103298161030e565b811461033457600080fd5b50565b60008135905061034681610320565b92915050565b600060208284031215610362576103616102e4565b5b600061037084828501610337565b91505092915050565b6000819050919050565b61038c81610379565b82525050565b60006020820190506103a76000830184610383565b92915050565b6103b68161030e565b82525050565b60006080820190506103d160008301876103ad565b6103de6020830186610383565b6103eb6040830185610383565b6103f86060830184610383565b95945050505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61044f82610406565b810181811067ffffffffffffffff8211171561046e5761046d610417565b5b80604052505050565b60006104816102da565b905061048d8282610446565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff8211156104bc576104bb610417565b5b6104c582610406565b9050602081019050919050565b82818337600083830152505050565b60006104f46104ef846104a1565b610477565b9050828152602081018484840111156105105761050f61049c565b5b61051b8482856104d2565b509392505050565b600082601f83011261053857610537610497565b5b81356105488482602086016104e1565b91505092915050565b6000819050919050565b61056481610551565b811461056f57600080fd5b50565b6000813590506105818161055b565b92915050565b61059081610379565b811461059b57600080fd5b50565b6000813590506105ad81610587565b92915050565b600281106105c057600080fd5b50565b6000813590506105d2816105b3565b92915050565b60006105e38261030e565b9050919050565b6105f3816105d8565b81146105fe57600080fd5b50565b600081359050610610816105ea565b92915050565b600060c0828403121561062c5761062b610401565b5b61063660c0610477565b9050600082013567ffffffffffffffff81111561065657610655610492565b5b61066284828501610523565b600083015250602061067684828501610572565b602083015250604061068a8482850161059e565b604083015250606061069e848285016105c3565b60608301525060806106b284828501610601565b60808301525060a06106c684828501610601565b60a08301525092915050565b60006106dd826102ee565b9050919050565b6106ed816106d2565b81146106f857600080fd5b50565b60008135905061070a816106e4565b92915050565b60008115159050919050565b61072581610710565b811461073057600080fd5b50565b6000813590506107428161071c565b92915050565b60006080828403121561075e5761075d610401565b5b6107686080610477565b9050600061077884828501610337565b600083015250602061078c848285016106fb565b60208301525060406107a084828501610733565b60408301525060606107b484828501610733565b60608301525092915050565b60008060008060e085870312156107da576107d96102e4565b5b600085013567ffffffffffffffff8111156107f8576107f76102e9565b5b61080487828801610616565b945050602061081587828801610748565b93505060a06108268782880161059e565b92505060c06108378782880161059e565b91505092959194509250565b6000806040838503121561085a576108596102e4565b5b60006108688582860161059e565b925050602061087985828601610337565b9150509250929050565b600060208284031215610899576108986102e4565b5b60006108a78482850161059e565b9150509291505056fea264697066735822122067c154bd14ebe43cac9977741cc66888aeb731f8756464c916aa96db0901361664736f6c634300080d0033"

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

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns(address recipient, uint256 swapAmount, uint256 limit, uint256 deadline)
func (_Element *ElementCaller) SwapCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Recipient  common.Address
	SwapAmount *big.Int
	Limit      *big.Int
	Deadline   *big.Int
}, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "swapCalled", arg0)

	outstruct := new(struct {
		Recipient  common.Address
		SwapAmount *big.Int
		Limit      *big.Int
		Deadline   *big.Int
	})

	outstruct.Recipient = out[0].(common.Address)
	outstruct.SwapAmount = out[1].(*big.Int)
	outstruct.Limit = out[2].(*big.Int)
	outstruct.Deadline = out[3].(*big.Int)

	return *outstruct, err

}

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns(address recipient, uint256 swapAmount, uint256 limit, uint256 deadline)
func (_Element *ElementSession) SwapCalled(arg0 common.Address) (struct {
	Recipient  common.Address
	SwapAmount *big.Int
	Limit      *big.Int
	Deadline   *big.Int
}, error) {
	return _Element.Contract.SwapCalled(&_Element.CallOpts, arg0)
}

// SwapCalled is a free data retrieval call binding the contract method 0x5658724f.
//
// Solidity: function swapCalled(address ) view returns(address recipient, uint256 swapAmount, uint256 limit, uint256 deadline)
func (_Element *ElementCallerSession) SwapCalled(arg0 common.Address) (struct {
	Recipient  common.Address
	SwapAmount *big.Int
	Limit      *big.Int
	Deadline   *big.Int
}, error) {
	return _Element.Contract.SwapCalled(&_Element.CallOpts, arg0)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_Element *ElementCaller) WithdrawPrincipalCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "withdrawPrincipalCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_Element *ElementSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _Element.Contract.WithdrawPrincipalCalled(&_Element.CallOpts, arg0)
}

// WithdrawPrincipalCalled is a free data retrieval call binding the contract method 0x3b1988dc.
//
// Solidity: function withdrawPrincipalCalled(address ) view returns(uint256)
func (_Element *ElementCallerSession) WithdrawPrincipalCalled(arg0 common.Address) (*big.Int, error) {
	return _Element.Contract.WithdrawPrincipalCalled(&_Element.CallOpts, arg0)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 l, uint256 d) returns(uint256)
func (_Element *ElementTransactor) Swap(opts *bind.TransactOpts, s SingleSwap, f FundManagement, l *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "swap", s, f, l, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 l, uint256 d) returns(uint256)
func (_Element *ElementSession) Swap(s SingleSwap, f FundManagement, l *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, l, d)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 l, uint256 d) returns(uint256)
func (_Element *ElementTransactorSession) Swap(s SingleSwap, f FundManagement, l *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, l, d)
}

// SwapReturns is a paid mutator transaction binding the contract method 0xd9410818.
//
// Solidity: function swapReturns(uint256 s) returns()
func (_Element *ElementTransactor) SwapReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "swapReturns", s)
}

// SwapReturns is a paid mutator transaction binding the contract method 0xd9410818.
//
// Solidity: function swapReturns(uint256 s) returns()
func (_Element *ElementSession) SwapReturns(s *big.Int) (*types.Transaction, error) {
	return _Element.Contract.SwapReturns(&_Element.TransactOpts, s)
}

// SwapReturns is a paid mutator transaction binding the contract method 0xd9410818.
//
// Solidity: function swapReturns(uint256 s) returns()
func (_Element *ElementTransactorSession) SwapReturns(s *big.Int) (*types.Transaction, error) {
	return _Element.Contract.SwapReturns(&_Element.TransactOpts, s)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_Element *ElementTransactor) WithdrawPrincipal(opts *bind.TransactOpts, a *big.Int, d common.Address) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "withdrawPrincipal", a, d)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_Element *ElementSession) WithdrawPrincipal(a *big.Int, d common.Address) (*types.Transaction, error) {
	return _Element.Contract.WithdrawPrincipal(&_Element.TransactOpts, a, d)
}

// WithdrawPrincipal is a paid mutator transaction binding the contract method 0x884e17f3.
//
// Solidity: function withdrawPrincipal(uint256 a, address d) returns()
func (_Element *ElementTransactorSession) WithdrawPrincipal(a *big.Int, d common.Address) (*types.Transaction, error) {
	return _Element.Contract.WithdrawPrincipal(&_Element.TransactOpts, a, d)
}
