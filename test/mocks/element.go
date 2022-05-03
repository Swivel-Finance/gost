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
const ElementABI = "[{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundManagementSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"return_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"singleSwapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractAny\",\"name\":\"assetIn\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"assetOut\",\"type\":\"address\"}],\"internalType\":\"structSingleSwap\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structFundManagement\",\"name\":\"f\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ElementBin is the compiled bytecode used for deploying new contracts.
var ElementBin = "0x608060405234801561001057600080fd5b506106ff806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806328386b831461005c57806369bbd8cd1461007a57806391d51144146100aa578063a4b15a5d146100c8578063aa8c217c146100e6575b600080fd5b610064610104565b60405161007191906101f8565b60405180910390f35b610094600480360381019061008f919061061c565b61012e565b6040516100a191906106ae565b60405180910390f35b6100b261019a565b6040516100bf91906106ae565b60405180910390f35b6100d06101a4565b6040516100dd91906106ae565b60405180910390f35b6100ee6101ae565b6040516100fb91906106ae565b60405180910390f35b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600082600081905550816001819055508360000151600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846040015160038190555060009050949350505050565b6000600154905090565b6000600354905090565b60008054905090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101e2826101b7565b9050919050565b6101f2816101d7565b82525050565b600060208201905061020d60008301846101e9565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102758261022c565b810181811067ffffffffffffffff821117156102945761029361023d565b5b80604052505050565b60006102a7610213565b90506102b3828261026c565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff8211156102e2576102e161023d565b5b6102eb8261022c565b9050602081019050919050565b82818337600083830152505050565b600061031a610315846102c7565b61029d565b905082815260208101848484011115610336576103356102c2565b5b6103418482856102f8565b509392505050565b600082601f83011261035e5761035d6102bd565b5b813561036e848260208601610307565b91505092915050565b6000819050919050565b61038a81610377565b811461039557600080fd5b50565b6000813590506103a781610381565b92915050565b6000819050919050565b6103c0816103ad565b81146103cb57600080fd5b50565b6000813590506103dd816103b7565b92915050565b600281106103f057600080fd5b50565b600081359050610402816103e3565b92915050565b6000610413826101d7565b9050919050565b61042381610408565b811461042e57600080fd5b50565b6000813590506104408161041a565b92915050565b600060c0828403121561045c5761045b610227565b5b61046660c061029d565b9050600082013567ffffffffffffffff811115610486576104856102b8565b5b61049284828501610349565b60008301525060206104a684828501610398565b60208301525060406104ba848285016103ce565b60408301525060606104ce848285016103f3565b60608301525060806104e284828501610431565b60808301525060a06104f684828501610431565b60a08301525092915050565b61050b816101d7565b811461051657600080fd5b50565b60008135905061052881610502565b92915050565b6000610539826101b7565b9050919050565b6105498161052e565b811461055457600080fd5b50565b60008135905061056681610540565b92915050565b60008115159050919050565b6105818161056c565b811461058c57600080fd5b50565b60008135905061059e81610578565b92915050565b6000608082840312156105ba576105b9610227565b5b6105c4608061029d565b905060006105d484828501610519565b60008301525060206105e884828501610557565b60208301525060406105fc8482850161058f565b60408301525060606106108482850161058f565b60608301525092915050565b60008060008060e085870312156106365761063561021d565b5b600085013567ffffffffffffffff81111561065457610653610222565b5b61066087828801610446565b9450506020610671878288016105a4565b93505060a0610682878288016103ce565b92505060c0610693878288016103ce565b91505092959194509250565b6106a8816103ad565b82525050565b60006020820190506106c3600083018461069f565b9291505056fea26469706673582212209d24e24a12b535ca80926692c5265bc017d4391b83422cc83335c37129edbd1c64736f6c634300080d0033"

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

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Element *ElementCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Element.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Element *ElementSession) Amount() (*big.Int, error) {
	return _Element.Contract.Amount(&_Element.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Element *ElementCallerSession) Amount() (*big.Int, error) {
	return _Element.Contract.Amount(&_Element.CallOpts)
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
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 a, uint256 r) returns(uint256)
func (_Element *ElementTransactor) Swap(opts *bind.TransactOpts, s SingleSwap, f FundManagement, a *big.Int, r *big.Int) (*types.Transaction, error) {
	return _Element.contract.Transact(opts, "swap", s, f, a, r)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 a, uint256 r) returns(uint256)
func (_Element *ElementSession) Swap(s SingleSwap, f FundManagement, a *big.Int, r *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, a, r)
}

// Swap is a paid mutator transaction binding the contract method 0x69bbd8cd.
//
// Solidity: function swap((bytes,bytes32,uint256,uint8,address,address) s, (address,address,bool,bool) f, uint256 a, uint256 r) returns(uint256)
func (_Element *ElementTransactorSession) Swap(s SingleSwap, f FundManagement, a *big.Int, r *big.Int) (*types.Transaction, error) {
	return _Element.Contract.Swap(&_Element.TransactOpts, s, f, a, r)
}
