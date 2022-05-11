// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeemer

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

// RedeemerABI is the input ABI used to generate the binding from.
const RedeemerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apwineRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"illuminate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedeemerBin is the compiled bytecode used for deploying new contracts.
var RedeemerBin = "0x608060405234801561001057600080fd5b506040516109fc3803806109fc83398101604081905261002f91610098565b600080546001600160a01b03199081163317909155600180546001600160a01b039586169083161790556002805493851693821693909317909255600380549190931691161790556100db565b80516001600160a01b038116811461009357600080fd5b919050565b6000806000606084860312156100ad57600080fd5b6100b68461007c565b92506100c46020850161007c565b91506100d26040850161007c565b90509250925092565b610912806100ea6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063cee501381161005b578063cee5013814610131578063d2ba878814610151578063e36e4b5614610169578063f851a4401461018257600080fd5b80637de923711461008d578063a1b1138c146100d7578063ba24ea6c146100fe578063cbf7e67014610111575b600080fd5b6002546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100ee6100e53660046106a2565b60019392505050565b60405190151581526020016100ce565b6100ee61010c3660046106e1565b6101a2565b6001546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6003546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6100ee61015f366004610732565b6001949350505050565b6100ee610177366004610776565b600195945050505050565b6000546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af1158015610220573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024491906107e7565b8660ff166008811061025857610258610894565b60200201516040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529192508291600091908316906370a08231906024016020604051808303816000875af11580156102d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f891906108c3565b90507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff8916016103b6576002546040517ff3fef3a300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8781166004830152602482018490529091169063f3fef3a3906044015b600060405180830381600087803b15801561039957600080fd5b505af11580156103ad573d6000803e3d6000fd5b505050506104df565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff89160161044e576003546040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820189905260448201849052898116606483015290911690636c8d4fa19060840161037f565b60ff88166104df576040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201839052841690639dc29fac90604401600060405180830381600087803b1580156104c657600080fd5b505af11580156104da573d6000803e3d6000fd5b505050505b6104ea87868361054c565b604080516000815260208101839052879173ffffffffffffffffffffffffffffffffffffffff8a16917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a3506001979650505050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af19150506105ad8161061d565b610617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015260640160405180910390fd5b50505050565b6000803d8361063057806000803e806000fd5b8060208114610648578015610659576000925061065e565b816000803e6000511515925061065e565b600192505b50909392505050565b803560ff8116811461067857600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461069f57600080fd5b50565b6000806000606084860312156106b757600080fd5b6106c084610667565b925060208401356106d08161067d565b929592945050506040919091013590565b600080600080608085870312156106f757600080fd5b61070085610667565b935060208501356107108161067d565b92506040850135915060608501356107278161067d565b939692955090935050565b6000806000806080858703121561074857600080fd5b61075185610667565b935060208501356107618161067d565b93969395505050506040820135916060013590565b600080600080600060a0868803121561078e57600080fd5b61079786610667565b945060208601356107a78161067d565b93506040860135925060608601356107be8161067d565b915060808601356107ce8161067d565b809150509295509295909350565b80516106788161067d565b60006101008083850312156107fb57600080fd5b83601f84011261080a57600080fd5b60405181810181811067ffffffffffffffff82111715610853577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290830190808583111561086857600080fd5b845b838110156108895761087b816107dc565b82526020918201910161086a565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156108d557600080fd5b505191905056fea2646970667358221220196b50729437202da6420531509f54d3fba37dc6bb0e5f4dbd1d750477e5f33a64736f6c634300080d0033"

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, a common.Address, t common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedeemerBin), backend, m, a, t)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// Redeemer is an auto generated Go binding around an Ethereum contract.
type Redeemer struct {
	RedeemerCaller     // Read-only binding to the contract
	RedeemerTransactor // Write-only binding to the contract
	RedeemerFilterer   // Log filterer for contract events
}

// RedeemerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemerSession struct {
	Contract     *Redeemer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemerCallerSession struct {
	Contract *RedeemerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RedeemerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemerTransactorSession struct {
	Contract     *RedeemerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RedeemerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemerRaw struct {
	Contract *Redeemer // Generic contract binding to access the raw methods on
}

// RedeemerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemerCallerRaw struct {
	Contract *RedeemerCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemerTransactorRaw struct {
	Contract *RedeemerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemer creates a new instance of Redeemer, bound to a specific deployed contract.
func NewRedeemer(address common.Address, backend bind.ContractBackend) (*Redeemer, error) {
	contract, err := bindRedeemer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// NewRedeemerCaller creates a new read-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerCaller(address common.Address, caller bind.ContractCaller) (*RedeemerCaller, error) {
	contract, err := bindRedeemer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerCaller{contract: contract}, nil
}

// NewRedeemerTransactor creates a new write-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemerTransactor, error) {
	contract, err := bindRedeemer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerTransactor{contract: contract}, nil
}

// NewRedeemerFilterer creates a new log filterer instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemerFilterer, error) {
	contract, err := bindRedeemer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemerFilterer{contract: contract}, nil
}

// bindRedeemer binds a generic wrapper to an already deployed contract.
func bindRedeemer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.RedeemerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCallerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// ApwineRouter is a free data retrieval call binding the contract method 0x7de92371.
//
// Solidity: function apwineRouter() view returns(address)
func (_Redeemer *RedeemerCaller) ApwineRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "apwineRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApwineRouter is a free data retrieval call binding the contract method 0x7de92371.
//
// Solidity: function apwineRouter() view returns(address)
func (_Redeemer *RedeemerSession) ApwineRouter() (common.Address, error) {
	return _Redeemer.Contract.ApwineRouter(&_Redeemer.CallOpts)
}

// ApwineRouter is a free data retrieval call binding the contract method 0x7de92371.
//
// Solidity: function apwineRouter() view returns(address)
func (_Redeemer *RedeemerCallerSession) ApwineRouter() (common.Address, error) {
	return _Redeemer.Contract.ApwineRouter(&_Redeemer.CallOpts)
}

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Redeemer *RedeemerCaller) Illuminate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "illuminate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Redeemer *RedeemerSession) Illuminate() (common.Address, error) {
	return _Redeemer.Contract.Illuminate(&_Redeemer.CallOpts)
}

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Redeemer *RedeemerCallerSession) Illuminate() (common.Address, error) {
	return _Redeemer.Contract.Illuminate(&_Redeemer.CallOpts)
}

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Redeemer *RedeemerCaller) TempusRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "tempusRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Redeemer *RedeemerSession) TempusRouter() (common.Address, error) {
	return _Redeemer.Contract.TempusRouter(&_Redeemer.CallOpts)
}

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Redeemer *RedeemerCallerSession) TempusRouter() (common.Address, error) {
	return _Redeemer.Contract.TempusRouter(&_Redeemer.CallOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem", p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem0", p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem1(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem1", p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address a) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem2(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, d common.Address, a common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem2", p, u, m, d, a)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address a) returns(bool)
func (_Redeemer *RedeemerSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, a common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, a)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address a) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, a common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, a)
}

// RedeemerRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Redeemer contract.
type RedeemerRedeemIterator struct {
	Event *RedeemerRedeem // Event containing the contract specifics and raw log

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
func (it *RedeemerRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemerRedeem)
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
		it.Event = new(RedeemerRedeem)
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
func (it *RedeemerRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemerRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemerRedeem represents a Redeem event raised by the Redeemer contract.
type RedeemerRedeem struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) FilterRedeem(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*RedeemerRedeemIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.FilterLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &RedeemerRedeemIterator{contract: _Redeemer.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *RedeemerRedeem, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.WatchLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemerRedeem)
				if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) ParseRedeem(log types.Log) (*RedeemerRedeem, error) {
	event := new(RedeemerRedeem)
	if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
