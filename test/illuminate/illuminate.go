// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package illuminate

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

// IlluminateMetaData contains all meta data concerning the Illuminate contract.
var IlluminateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address[7]\",\"name\":\"t\",\"type\":\"address[7]\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610ee4380380610ee483398101604081905261002f91610052565b600180546001600160a01b031916331790556001600160a01b0316608052610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051610e416100a360003960008181609301526103750152610e416000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063125cf47f146100515780632ba29d381461008e5780635f1f94e6146100b5578063f851a440146100d8575b600080fd5b61006461005f36600461061e565b6100f8565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100647f000000000000000000000000000000000000000000000000000000000000000081565b6100c86100c33660046106da565b610140565b6040519015158152602001610085565b6001546100649073ffffffffffffffffffffffffffffffffffffffff1681565b6000602052826000526040600020602052816000526040600020816008811061012057600080fd5b015473ffffffffffffffffffffffffffffffffffffffff16925083915050565b60015460009073ffffffffffffffffffffffffffffffffffffffff163381146101ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a81166000908152602081815260408083208d84529091529020541615610262576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d61726b6574206578697374730000000000000000000000000000000000000060448201526064016101c1565b60008a8a898989898960405161027790610555565b6102879796959493929190610862565b604051809103906000f0801580156102a3573d6000803e3d6000fd5b50604080516101008101825273ffffffffffffffffffffffffffffffffffffffff83811682528c5181166020808401919091528d0151811682840152918c015182166060808301919091528c015182166080808301919091528c0151821660a0808301919091528c0151821660c0808301919091528c015190911660e08201529091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60005b60088160ff1610156103ac5761039a8c8260ff166007811061036e5761036e6107ea565b60200201517f00000000000000000000000000000000000000000000000000000000000000008461043e565b806103a4816108c3565b91505061034a565b5073ffffffffffffffffffffffffffffffffffffffff8d166000908152602081815260408083208f845290915290206103e790836008610563565b506040518c9073ffffffffffffffffffffffffffffffffffffffff8f16907f0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee90600090a35060019c9b505050505050505050505050565b60006040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af191505061049f8161050b565b610505576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f617070726f7665206661696c656400000000000000000000000000000000000060448201526064016101c1565b50505050565b6000803d8361051e57806000803e806000fd5b8060208114610536578015610547576000925061054c565b816000803e6000511515925061054c565b600192505b50909392505050565b610502806200090a83390190565b82600881019282156105d0579160200282015b828111156105d057825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610576565b506105dc9291506105e0565b5090565b5b808211156105dc57600081556001016105e1565b803573ffffffffffffffffffffffffffffffffffffffff8116811461061957600080fd5b919050565b60008060006060848603121561063357600080fd5b61063c846105f5565b95602085013595506040909401359392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083601f84011261069257600080fd5b50813567ffffffffffffffff8111156106aa57600080fd5b6020830191508360208285010111156106c257600080fd5b9250929050565b803560ff8116811461061957600080fd5b600080600080600080600080610180898b0312156106f757600080fd5b610700896105f5565b97506020808a013597508a605f8b011261071957600080fd5b60405160e0810167ffffffffffffffff828210818311171561073d5761073d610651565b816040528291506101208d018e81111561075657600080fd5b60408e015b818110156107795761076c816105f5565b845292850192850161075b565b50839a5080359450508084111561078f57600080fd5b61079b8e858f01610680565b90995097506101408d01359350889250808411156107b857600080fd5b5050506107c78b828c01610680565b90945092506107db90506101608a016106c9565b90509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a06040820152600061089860a083018789610819565b82810360608401526108ab818688610819565b91505060ff8316608083015298975050505050505050565b600060ff821660ff8103610900577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019291505056fe608060405234801561001057600080fd5b5060405161050238038061050283398101604081905261002f916101f0565b600080546001600160a01b0319166001600160a01b03871617905560018490558251610062906002906020860190610095565b508151610076906003906020850190610095565b506004805460ff191660ff92909216919091179055506102d192505050565b8280546100a190610297565b90600052602060002090601f0160209004810192826100c35760008555610109565b82601f106100dc57805160ff1916838001178555610109565b82800160010185558215610109579182015b828111156101095782518255916020019190600101906100ee565b50610115929150610119565b5090565b5b80821115610115576000815560010161011a565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261015557600080fd5b81516001600160401b038082111561016f5761016f61012e565b604051601f8301601f19908116603f011681019082821181831017156101975761019761012e565b816040528381526020925086838588010111156101b357600080fd5b600091505b838210156101d557858201830151818301840152908201906101b8565b838211156101e65760008385830101525b9695505050505050565b600080600080600060a0868803121561020857600080fd5b85516001600160a01b038116811461021f57600080fd5b6020870151604088015191965094506001600160401b038082111561024357600080fd5b61024f89838a01610144565b9450606088015191508082111561026557600080fd5b5061027288828901610144565b925050608086015160ff8116811461028957600080fd5b809150509295509295909350565b600181811c908216806102ab57607f821691505b6020821081036102cb57634e487b7160e01b600052602260045260246000fd5b50919050565b610222806102e06000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306fdde0314610046578063313ce5671461006457806395d89b4114610083575b600080fd5b61004e61008b565b60405161005b9190610126565b60405180910390f35b6004546100719060ff1681565b60405160ff909116815260200161005b565b61004e610119565b6002805461009890610199565b80601f01602080910402602001604051908101604052809291908181526020018280546100c490610199565b80156101115780601f106100e657610100808354040283529160200191610111565b820191906000526020600020905b8154815290600101906020018083116100f457829003601f168201915b505050505081565b6003805461009890610199565b600060208083528351808285015260005b8181101561015357858101830151858201604001528201610137565b81811115610165576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b600181811c908216806101ad57607f821691505b6020821081036101e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea264697066735822122078ae9f332b5c9c1452cdad76ff212ca51016fcc6c26886c24a1429f0fd0e1b2664736f6c634300080d0033a264697066735822122097fa1f525e4bcfe5b4cb7acd86d9d92f46440079706a142cd25b026d052b0f7064736f6c634300080d0033",
}

// IlluminateABI is the input ABI used to generate the binding from.
// Deprecated: Use IlluminateMetaData.ABI instead.
var IlluminateABI = IlluminateMetaData.ABI

// IlluminateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IlluminateMetaData.Bin instead.
var IlluminateBin = IlluminateMetaData.Bin

// DeployIlluminate deploys a new Ethereum contract, binding an instance of Illuminate to it.
func DeployIlluminate(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Illuminate, error) {
	parsed, err := IlluminateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IlluminateBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// Illuminate is an auto generated Go binding around an Ethereum contract.
type Illuminate struct {
	IlluminateCaller     // Read-only binding to the contract
	IlluminateTransactor // Write-only binding to the contract
	IlluminateFilterer   // Log filterer for contract events
}

// IlluminateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IlluminateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IlluminateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IlluminateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IlluminateSession struct {
	Contract     *Illuminate       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IlluminateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IlluminateCallerSession struct {
	Contract *IlluminateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IlluminateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IlluminateTransactorSession struct {
	Contract     *IlluminateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IlluminateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IlluminateRaw struct {
	Contract *Illuminate // Generic contract binding to access the raw methods on
}

// IlluminateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IlluminateCallerRaw struct {
	Contract *IlluminateCaller // Generic read-only contract binding to access the raw methods on
}

// IlluminateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IlluminateTransactorRaw struct {
	Contract *IlluminateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIlluminate creates a new instance of Illuminate, bound to a specific deployed contract.
func NewIlluminate(address common.Address, backend bind.ContractBackend) (*Illuminate, error) {
	contract, err := bindIlluminate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// NewIlluminateCaller creates a new read-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateCaller(address common.Address, caller bind.ContractCaller) (*IlluminateCaller, error) {
	contract, err := bindIlluminate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateCaller{contract: contract}, nil
}

// NewIlluminateTransactor creates a new write-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateTransactor(address common.Address, transactor bind.ContractTransactor) (*IlluminateTransactor, error) {
	contract, err := bindIlluminate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateTransactor{contract: contract}, nil
}

// NewIlluminateFilterer creates a new log filterer instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateFilterer(address common.Address, filterer bind.ContractFilterer) (*IlluminateFilterer, error) {
	contract, err := bindIlluminate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IlluminateFilterer{contract: contract}, nil
}

// bindIlluminate binds a generic wrapper to an already deployed contract.
func bindIlluminate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IlluminateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.IlluminateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateSession) Admin() (common.Address, error) {
	return _Illuminate.Contract.Admin(&_Illuminate.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateCallerSession) Admin() (common.Address, error) {
	return _Illuminate.Contract.Admin(&_Illuminate.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateCaller) Markets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "markets", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.CallOpts, arg0, arg1, arg2)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateCallerSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.CallOpts, arg0, arg1, arg2)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateCaller) Redeemer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "redeemer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateSession) Redeemer() (common.Address, error) {
	return _Illuminate.Contract.Redeemer(&_Illuminate.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateCallerSession) Redeemer() (common.Address, error) {
	return _Illuminate.Contract.Redeemer(&_Illuminate.CallOpts)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateTransactor) CreateMarket(opts *bind.TransactOpts, u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.contract.Transact(opts, "createMarket", u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateSession) CreateMarket(u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.Contract.CreateMarket(&_Illuminate.TransactOpts, u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateTransactorSession) CreateMarket(u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.Contract.CreateMarket(&_Illuminate.TransactOpts, u, m, t, n, s, d)
}

// IlluminateCreateMarketIterator is returned from FilterCreateMarket and is used to iterate over the raw logs and unpacked data for CreateMarket events raised by the Illuminate contract.
type IlluminateCreateMarketIterator struct {
	Event *IlluminateCreateMarket // Event containing the contract specifics and raw log

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
func (it *IlluminateCreateMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlluminateCreateMarket)
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
		it.Event = new(IlluminateCreateMarket)
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
func (it *IlluminateCreateMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlluminateCreateMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlluminateCreateMarket represents a CreateMarket event raised by the Illuminate contract.
type IlluminateCreateMarket struct {
	Underlying common.Address
	Maturity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateMarket is a free log retrieval operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) FilterCreateMarket(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*IlluminateCreateMarketIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Illuminate.contract.FilterLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &IlluminateCreateMarketIterator{contract: _Illuminate.contract, event: "CreateMarket", logs: logs, sub: sub}, nil
}

// WatchCreateMarket is a free log subscription operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) WatchCreateMarket(opts *bind.WatchOpts, sink chan<- *IlluminateCreateMarket, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Illuminate.contract.WatchLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlluminateCreateMarket)
				if err := _Illuminate.contract.UnpackLog(event, "CreateMarket", log); err != nil {
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

// ParseCreateMarket is a log parse operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) ParseCreateMarket(log types.Log) (*IlluminateCreateMarket, error) {
	event := new(IlluminateCreateMarket)
	if err := _Illuminate.contract.UnpackLog(event, "CreateMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
