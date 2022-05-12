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
const ElementABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"f\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"l\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"swapReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ElementBin is the compiled bytecode used for deploying new contracts.
var ElementBin = "0x608060405234801561001057600080fd5b506107e4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635658724f1461004657806369bbd8cd14610079578063d9410818146100a9575b600080fd5b610060600480360381019061005b919061028a565b6100c5565b60405161007094939291906102df565b60405180910390f35b610093600480360381019061008e91906106e3565b610115565b6040516100a09190610766565b60405180910390f35b6100c360048036038101906100be9190610781565b61020e565b005b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b60006040518060800160405280856020015173ffffffffffffffffffffffffffffffffffffffff168152602001866040015181526020018481526020018381525060016000866000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301559050506000549050949350505050565b8060008190555050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102578261022c565b9050919050565b6102678161024c565b811461027257600080fd5b50565b6000813590506102848161025e565b92915050565b6000602082840312156102a05761029f610222565b5b60006102ae84828501610275565b91505092915050565b6102c08161024c565b82525050565b6000819050919050565b6102d9816102c6565b82525050565b60006080820190506102f460008301876102b7565b61030160208301866102d0565b61030e60408301856102d0565b61031b60608301846102d0565b95945050505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61037282610329565b810181811067ffffffffffffffff821117156103915761039061033a565b5b80604052505050565b60006103a4610218565b90506103b08282610369565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff8211156103df576103de61033a565b5b6103e882610329565b9050602081019050919050565b82818337600083830152505050565b6000610417610412846103c4565b61039a565b905082815260208101848484011115610433576104326103bf565b5b61043e8482856103f5565b509392505050565b600082601f83011261045b5761045a6103ba565b5b813561046b848260208601610404565b91505092915050565b6000819050919050565b61048781610474565b811461049257600080fd5b50565b6000813590506104a48161047e565b92915050565b6104b3816102c6565b81146104be57600080fd5b50565b6000813590506104d0816104aa565b92915050565b600281106104e357600080fd5b50565b6000813590506104f5816104d6565b92915050565b60006105068261024c565b9050919050565b610516816104fb565b811461052157600080fd5b50565b6000813590506105338161050d565b92915050565b600060c0828403121561054f5761054e610324565b5b61055960c061039a565b9050600082013567ffffffffffffffff811115610579576105786103b5565b5b61058584828501610446565b600083015250602061059984828501610495565b60208301525060406105ad848285016104c1565b60408301525060606105c1848285016104e6565b60608301525060806105d584828501610524565b60808301525060a06105e984828501610524565b60a08301525092915050565b60006106008261022c565b9050919050565b610610816105f5565b811461061b57600080fd5b50565b60008135905061062d81610607565b92915050565b60008115159050919050565b61064881610633565b811461065357600080fd5b50565b6000813590506106658161063f565b92915050565b60006080828403121561068157610680610324565b5b61068b608061039a565b9050600061069b84828501610275565b60008301525060206106af8482850161061e565b60208301525060406106c384828501610656565b60408301525060606106d784828501610656565b60608301525092915050565b60008060008060e085870312156106fd576106fc610222565b5b600085013567ffffffffffffffff81111561071b5761071a610227565b5b61072787828801610539565b94505060206107388782880161066b565b93505060a0610749878288016104c1565b92505060c061075a878288016104c1565b91505092959194509250565b600060208201905061077b60008301846102d0565b92915050565b60006020828403121561079757610796610222565b5b60006107a5848285016104c1565b9150509291505056fea26469706673582212202c5293343b3746a207329ea9f6af34d657809d780d81d5596a52d7e31f42ab5964736f6c634300080d0033"

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
