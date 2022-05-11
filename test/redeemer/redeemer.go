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
const RedeemerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"i\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"illuminate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"illuminateRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendleRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedeemerBin is the compiled bytecode used for deploying new contracts.
var RedeemerBin = "0x60806040523480156200001157600080fd5b5060405162001177380380620011778339810160408190526200003491620000be565b60008054336001600160a01b0319918216179091556001805482166001600160a01b0397881617905560028054821695871695909517909455600380548516938616939093179092556004805484169185169190911790556005805490921692169190911790556200012e565b80516001600160a01b0381168114620000b957600080fd5b919050565b600080600080600060a08688031215620000d757600080fd5b620000e286620000a1565b9450620000f260208701620000a1565b93506200010260408701620000a1565b92506200011260608701620000a1565b91506200012260808701620000a1565b90509295509295909350565b611039806200013e6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063cbf7e67011610076578063d2ba87881161005b578063d2ba8788146101c3578063e36e4b56146101d6578063f851a440146101e957600080fd5b8063cbf7e67014610183578063cee50138146101a357600080fd5b8063502a93c3116100a7578063502a93c31461012d578063a1b1138c1461014d578063ba24ea6c1461017057600080fd5b80631bc90469146100c3578063206aeab31461010d575b600080fd5b6004546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6002546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6005546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b61016061015b366004610da0565b610209565b6040519015158152602001610104565b61016061017e366004610ddf565b6104ee565b6001546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6003546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6101606101d1366004610e30565b610892565b6101606101e4366004610e74565b610af7565b6000546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526024820184905260009283929116906317b3bba790604401610100604051808303816000875af1158015610287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ab9190610ee5565b8560ff16600881106102bf576102bf610f92565b60200201516001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af115801561033c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103609190610fc1565b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018390529192508316906323b872dd906064016020604051808303816000875af11580156103e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104059190610fda565b506040517f1e9a69500000000000000000000000000000000000000000000000000000000081523060048201526024810182905273ffffffffffffffffffffffffffffffffffffffff831690631e9a695090604401600060405180830381600087803b15801561047457600080fd5b505af1158015610488573d6000803e3d6000fd5b50506040805160ff8a1681526020810185905287935073ffffffffffffffffffffffffffffffffffffffff891692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600195945050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af115801561056c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105909190610ee5565b8660ff16600881106105a4576105a4610f92565b6020020151905060007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff8816016105f6575060045473ffffffffffffffffffffffffffffffffffffffff1661065a565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff88160161063f575060035473ffffffffffffffffffffffffffffffffffffffff1661065a565b5060025473ffffffffffffffffffffffffffffffffffffffff165b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000918416906370a08231906024016020604051808303816000875af11580156106ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f29190610fc1565b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018390529192508416906323b872dd906064016020604051808303816000875af1158015610773573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107979190610fda565b506040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820183905260006044830152306064830152831690636c8d4fa1906084015b600060405180830381600087803b15801561081657600080fd5b505af115801561082a573d6000803e3d6000fd5b50506040805160ff8c1681526020810185905289935073ffffffffffffffffffffffffffffffffffffffff8b1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a3506001979650505050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af1158015610910573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109349190610ee5565b8660ff166008811061094857610948610f92565b60200201516002546001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152929350908116916000918416906370a08231906024016020604051808303816000875af11580156109cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ef9190610fc1565b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018390529192508416906323b872dd906064016020604051808303816000875af1158015610a70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a949190610fda565b506040517fd230566c0000000000000000000000000000000000000000000000000000000081526004810186905273ffffffffffffffffffffffffffffffffffffffff88811660248301526044820188905283169063d230566c906064016107fc565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260009283929116906317b3bba790604401610100604051808303816000875af1158015610b75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b999190610ee5565b8760ff1660088110610bad57610bad610f92565b60200201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a08231906024016020604051808303816000875af1158015610c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c489190610fc1565b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018390529192508316906323b872dd906064016020604051808303816000875af1158015610cc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ced9190610fda565b506040517f2b83cccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820188905260448201839052861690632b83cccd90606401600060405180830381600087803b15801561081657600080fd5b803560ff81168114610d7657600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610d9d57600080fd5b50565b600080600060608486031215610db557600080fd5b610dbe84610d65565b92506020840135610dce81610d7b565b929592945050506040919091013590565b60008060008060808587031215610df557600080fd5b610dfe85610d65565b93506020850135610e0e81610d7b565b9250604085013591506060850135610e2581610d7b565b939692955090935050565b60008060008060808587031215610e4657600080fd5b610e4f85610d65565b93506020850135610e5f81610d7b565b93969395505050506040820135916060013590565b600080600080600060a08688031215610e8c57600080fd5b610e9586610d65565b94506020860135610ea581610d7b565b9350604086013592506060860135610ebc81610d7b565b91506080860135610ecc81610d7b565b809150509295509295909350565b8051610d7681610d7b565b6000610100808385031215610ef957600080fd5b83601f840112610f0857600080fd5b60405181810181811067ffffffffffffffff82111715610f51577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052908301908085831115610f6657600080fd5b845b83811015610f8757610f7981610eda565b825260209182019101610f68565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610fd357600080fd5b5051919050565b600060208284031215610fec57600080fd5b81518015158114610ffc57600080fd5b939250505056fea26469706673582212201b13090eeba9b3c65e080bbca312937bc93866881c7ae293cb2c9587d3ed211064736f6c634300080d0033"

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, p common.Address, t common.Address, a common.Address, i common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedeemerBin), backend, m, p, t, a, i)
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

// ApRouter is a free data retrieval call binding the contract method 0x1bc90469.
//
// Solidity: function apRouter() view returns(address)
func (_Redeemer *RedeemerCaller) ApRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "apRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApRouter is a free data retrieval call binding the contract method 0x1bc90469.
//
// Solidity: function apRouter() view returns(address)
func (_Redeemer *RedeemerSession) ApRouter() (common.Address, error) {
	return _Redeemer.Contract.ApRouter(&_Redeemer.CallOpts)
}

// ApRouter is a free data retrieval call binding the contract method 0x1bc90469.
//
// Solidity: function apRouter() view returns(address)
func (_Redeemer *RedeemerCallerSession) ApRouter() (common.Address, error) {
	return _Redeemer.Contract.ApRouter(&_Redeemer.CallOpts)
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

// IlluminateRouter is a free data retrieval call binding the contract method 0x502a93c3.
//
// Solidity: function illuminateRouter() view returns(address)
func (_Redeemer *RedeemerCaller) IlluminateRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "illuminateRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IlluminateRouter is a free data retrieval call binding the contract method 0x502a93c3.
//
// Solidity: function illuminateRouter() view returns(address)
func (_Redeemer *RedeemerSession) IlluminateRouter() (common.Address, error) {
	return _Redeemer.Contract.IlluminateRouter(&_Redeemer.CallOpts)
}

// IlluminateRouter is a free data retrieval call binding the contract method 0x502a93c3.
//
// Solidity: function illuminateRouter() view returns(address)
func (_Redeemer *RedeemerCallerSession) IlluminateRouter() (common.Address, error) {
	return _Redeemer.Contract.IlluminateRouter(&_Redeemer.CallOpts)
}

// PendleRouter is a free data retrieval call binding the contract method 0x206aeab3.
//
// Solidity: function pendleRouter() view returns(address)
func (_Redeemer *RedeemerCaller) PendleRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "pendleRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendleRouter is a free data retrieval call binding the contract method 0x206aeab3.
//
// Solidity: function pendleRouter() view returns(address)
func (_Redeemer *RedeemerSession) PendleRouter() (common.Address, error) {
	return _Redeemer.Contract.PendleRouter(&_Redeemer.CallOpts)
}

// PendleRouter is a free data retrieval call binding the contract method 0x206aeab3.
//
// Solidity: function pendleRouter() view returns(address)
func (_Redeemer *RedeemerCallerSession) PendleRouter() (common.Address, error) {
	return _Redeemer.Contract.PendleRouter(&_Redeemer.CallOpts)
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
