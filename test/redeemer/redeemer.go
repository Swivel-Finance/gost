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
const RedeemerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apwineAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"illuminate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedeemerBin is the compiled bytecode used for deploying new contracts.
var RedeemerBin = "0x60806040523480156200001157600080fd5b5060405162001320380380620013208339810160408190526200003491620000be565b60008054336001600160a01b0319918216179091556001805482166001600160a01b0397881617905560028054821695871695909517909455600380548516938616939093179092556004805484169185169190911790556005805490921692169190911790556200012e565b80516001600160a01b0381168114620000b957600080fd5b919050565b600080600080600060a08688031215620000d757600080fd5b620000e286620000a1565b9450620000f260208701620000a1565b93506200010260408701620000a1565b92506200011260608701620000a1565b91506200012260808701620000a1565b90509295509295909350565b6111e2806200013e6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063de1d3cb511610076578063ea08c0311161005b578063ea08c031146101a9578063ef603569146101c9578063f851a440146101e957600080fd5b8063de1d3cb514610176578063e36e4b561461019657600080fd5b8063ba24ea6c116100a7578063ba24ea6c14610130578063cbf7e67014610143578063d2ba87881461016357600080fd5b806337a3eeec146100c3578063a1b1138c1461010d575b600080fd5b6002546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012061011b366004610f49565b610209565b6040519015158152602001610104565b61012061013e366004610f88565b610490565b6001546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b610120610171366004610fd9565b610837565b6003546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6101206101a436600461101d565b610a9b565b6005546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6004546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6000546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526024820184905260009283929116906317b3bba790604401610100604051808303816000875af1158015610287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ab919061108e565b8560ff16600881106102bf576102bf61113b565b60200201516001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af115801561033c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610360919061116a565b60015490915061038990839073ffffffffffffffffffffffffffffffffffffffff163084610d08565b6005546040517f154e0f2e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201879052604482018490529091169063154e0f2e906064016020604051808303816000875af1158015610409573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042d9190611183565b506040805160ff8816815260208101839052859173ffffffffffffffffffffffffffffffffffffffff8816917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600195945050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af115801561050e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610532919061108e565b8660ff16600881106105465761054661113b565b60200201516040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529192506000918316906370a08231906024016020604051808303816000875af11580156105bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e3919061116a565b90507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff8816016106a1576002546040517ff3fef3a300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490529091169063f3fef3a3906044015b600060405180830381600087803b15801561068457600080fd5b505af1158015610698573d6000803e3d6000fd5b505050506107ca565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff881601610739576003546040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820188905260448201849052888116606483015290911690636c8d4fa19060840161066a565b60ff87166107ca576040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201839052831690639dc29fac90604401600060405180830381600087803b1580156107b157600080fd5b505af11580156107c5573d6000803e3d6000fd5b505050505b6107d5868583610df7565b604080516000815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe991015b60405180910390a35060019695505050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af11580156108b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d9919061108e565b8660ff16600881106108ed576108ed61113b565b60200201516001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af115801561096a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098e919061116a565b6001549091506109b790839073ffffffffffffffffffffffffffffffffffffffff163084610d08565b600480546040517fd230566c00000000000000000000000000000000000000000000000000000000815291820186905273ffffffffffffffffffffffffffffffffffffffff888116602484015260448301889052169063d230566c90606401600060405180830381600087803b158015610a3057600080fd5b505af1158015610a44573d6000803e3d6000fd5b50506040805160ff8b1681526020810185905288935073ffffffffffffffffffffffffffffffffffffffff8a1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe99101610822565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260009283929116906317b3bba790604401610100604051808303816000875af1158015610b19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3d919061108e565b8760ff1660088110610b5157610b5161113b565b60200201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a08231906024016020604051808303816000875af1158015610bc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bec919061116a565b600154909150610c1590839073ffffffffffffffffffffffffffffffffffffffff163084610d08565b6040517f2b83cccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820188905260448201839052861690632b83cccd90606401600060405180830381600087803b158015610c8c57600080fd5b505af1158015610ca0573d6000803e3d6000fd5b50506040805160ff8c1681526020810185905289935073ffffffffffffffffffffffffffffffffffffffff8b1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a3506001979650505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af1915050610d8581610ec4565b610df0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064015b60405180910390fd5b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050610e5881610ec4565b610ebe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c656400000000000000000000000000000000006044820152606401610de7565b50505050565b6000803d83610ed757806000803e806000fd5b8060208114610eef578015610f005760009250610f05565b816000803e60005115159250610f05565b600192505b50909392505050565b803560ff81168114610f1f57600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f4657600080fd5b50565b600080600060608486031215610f5e57600080fd5b610f6784610f0e565b92506020840135610f7781610f24565b929592945050506040919091013590565b60008060008060808587031215610f9e57600080fd5b610fa785610f0e565b93506020850135610fb781610f24565b9250604085013591506060850135610fce81610f24565b939692955090935050565b60008060008060808587031215610fef57600080fd5b610ff885610f0e565b9350602085013561100881610f24565b93969395505050506040820135916060013590565b600080600080600060a0868803121561103557600080fd5b61103e86610f0e565b9450602086013561104e81610f24565b935060408601359250606086013561106581610f24565b9150608086013561107581610f24565b809150509295509295909350565b8051610f1f81610f24565b60006101008083850312156110a257600080fd5b83601f8401126110b157600080fd5b60405181810181811067ffffffffffffffff821117156110fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290830190808583111561110f57600080fd5b845b838110156111305761112281611083565b825260209182019101611111565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561117c57600080fd5b5051919050565b60006020828403121561119557600080fd5b815180151581146111a557600080fd5b939250505056fea2646970667358221220f3dcbc9b5c327c03e1865d9718e69dda3b4bce47356d330ca726a4ca1556da9864736f6c634300080d0033"

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, a common.Address, t common.Address, p common.Address, s common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedeemerBin), backend, m, a, t, p, s)
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

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCaller) ApwineAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "apwineAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
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

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCaller) PendleAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "pendleAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCaller) SwivelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "swivelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCaller) TempusAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "tempusAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
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
