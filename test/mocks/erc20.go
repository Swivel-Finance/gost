// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowanceCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"allowanceReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approveCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"approveReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"}],\"name\":\"decimalsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"nameReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"symbolReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061105b806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806395d89b41116100ad578063ad64dcc511610071578063ad64dcc514610353578063af599a291461036f578063c1d2e9a11461038b578063dd62ed3e146103bb578063e541efa2146103eb5761012c565b806395d89b411461029d5780639dd0ff37146102bb578063a18729a3146102d7578063a5b836bf146102f3578063a9059cbb146103235761012c565b806339100838116100f457806339100838146101e957806342b6cdbc146102055780636521b96a146102215780636581d5431461023d57806370a082311461026d5761012c565b806306fdde0314610131578063095ea7b31461014f5780630ab2519e1461017f57806323b872dd1461019b578063313ce567146101cb575b600080fd5b61013961041c565b6040516101469190610ab4565b60405180910390f35b61016960048036038101906101649190610b7e565b6104ae565b6040516101769190610bd9565b60405180910390f35b61019960048036038101906101949190610d29565b61050b565b005b6101b560048036038101906101b09190610d72565b610525565b6040516101c29190610bd9565b60405180910390f35b6101d361061f565b6040516101e09190610de1565b60405180910390f35b61020360048036038101906101fe9190610dfc565b610636565b005b61021f600480360381019061021a9190610e55565b610640565b005b61023b60048036038101906102369190610e55565b61065d565b005b61025760048036038101906102529190610e82565b61067a565b6040516102649190610ebe565b60405180910390f35b61028760048036038101906102829190610e82565b610692565b6040516102949190610ebe565b60405180910390f35b6102a56106df565b6040516102b29190610ab4565b60405180910390f35b6102d560048036038101906102d09190610e55565b610771565b005b6102f160048036038101906102ec9190610f05565b61078e565b005b61030d60048036038101906103089190610e82565b6107ac565b60405161031a9190610f41565b60405180910390f35b61033d60048036038101906103389190610b7e565b6107df565b60405161034a9190610bd9565b60405180910390f35b61036d60048036038101906103689190610dfc565b61083d565b005b61038960048036038101906103849190610d29565b610847565b005b6103a560048036038101906103a09190610e82565b610861565b6040516103b29190610ebe565b60405180910390f35b6103d560048036038101906103d09190610f5c565b610879565b6040516103e29190610ebe565b60405180910390f35b61040560048036038101906104009190610e82565b610904565b604051610413929190610f9c565b60405180910390f35b60606005805461042b90610ff4565b80601f016020809104026020016040519081016040528092919081815260200182805461045790610ff4565b80156104a45780601f10610479576101008083540402835291602001916104a4565b820191906000526020600020905b81548152906001019060200180831161048757829003601f168201915b5050505050905090565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600760019054906101000a900460ff16905092915050565b8060059080519060200190610521929190610948565b5050565b600061052f6109ce565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600a60019054906101000a900460ff169150509392505050565b6000600760009054906101000a900460ff16905090565b8060088190555050565b80600a60006101000a81548160ff02191690831515021790555050565b80600a60016101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915090505481565b600081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008549050919050565b6060600680546106ee90610ff4565b80601f016020809104026020016040519081016040528092919081815260200182805461071a90610ff4565b80156107675780601f1061073c57610100808354040283529160200191610767565b820191906000526020600020905b81548152906001019060200180831161074a57829003601f168201915b5050505050905090565b80600760016101000a81548160ff02191690831515021790555050565b80600760006101000a81548160ff021916908360ff16021790555050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600a60009054906101000a900460ff16905092915050565b8060098190555050565b806006908051906020019061085d929190610948565b5050565b60016020528060005260406000206000915090505481565b600081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600954905092915050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b82805461095490610ff4565b90600052602060002090601f01602090048101928261097657600085556109bd565b82601f1061098f57805160ff19168380011785556109bd565b828001600101855582156109bd579182015b828111156109bc5782518255916020019190600101906109a1565b5b5090506109ca91906109fe565b5090565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b5b80821115610a175760008160009055506001016109ff565b5090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a55578082015181840152602081019050610a3a565b83811115610a64576000848401525b50505050565b6000601f19601f8301169050919050565b6000610a8682610a1b565b610a908185610a26565b9350610aa0818560208601610a37565b610aa981610a6a565b840191505092915050565b60006020820190508181036000830152610ace8184610a7b565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b1582610aea565b9050919050565b610b2581610b0a565b8114610b3057600080fd5b50565b600081359050610b4281610b1c565b92915050565b6000819050919050565b610b5b81610b48565b8114610b6657600080fd5b50565b600081359050610b7881610b52565b92915050565b60008060408385031215610b9557610b94610ae0565b5b6000610ba385828601610b33565b9250506020610bb485828601610b69565b9150509250929050565b60008115159050919050565b610bd381610bbe565b82525050565b6000602082019050610bee6000830184610bca565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c3682610a6a565b810181811067ffffffffffffffff82111715610c5557610c54610bfe565b5b80604052505050565b6000610c68610ad6565b9050610c748282610c2d565b919050565b600067ffffffffffffffff821115610c9457610c93610bfe565b5b610c9d82610a6a565b9050602081019050919050565b82818337600083830152505050565b6000610ccc610cc784610c79565b610c5e565b905082815260208101848484011115610ce857610ce7610bf9565b5b610cf3848285610caa565b509392505050565b600082601f830112610d1057610d0f610bf4565b5b8135610d20848260208601610cb9565b91505092915050565b600060208284031215610d3f57610d3e610ae0565b5b600082013567ffffffffffffffff811115610d5d57610d5c610ae5565b5b610d6984828501610cfb565b91505092915050565b600080600060608486031215610d8b57610d8a610ae0565b5b6000610d9986828701610b33565b9350506020610daa86828701610b33565b9250506040610dbb86828701610b69565b9150509250925092565b600060ff82169050919050565b610ddb81610dc5565b82525050565b6000602082019050610df66000830184610dd2565b92915050565b600060208284031215610e1257610e11610ae0565b5b6000610e2084828501610b69565b91505092915050565b610e3281610bbe565b8114610e3d57600080fd5b50565b600081359050610e4f81610e29565b92915050565b600060208284031215610e6b57610e6a610ae0565b5b6000610e7984828501610e40565b91505092915050565b600060208284031215610e9857610e97610ae0565b5b6000610ea684828501610b33565b91505092915050565b610eb881610b48565b82525050565b6000602082019050610ed36000830184610eaf565b92915050565b610ee281610dc5565b8114610eed57600080fd5b50565b600081359050610eff81610ed9565b92915050565b600060208284031215610f1b57610f1a610ae0565b5b6000610f2984828501610ef0565b91505092915050565b610f3b81610b0a565b82525050565b6000602082019050610f566000830184610f32565b92915050565b60008060408385031215610f7357610f72610ae0565b5b6000610f8185828601610b33565b9250506020610f9285828601610b33565b9150509250929050565b6000604082019050610fb16000830185610f32565b610fbe6020830184610eaf565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061100c57607f821691505b60208210810361101f5761101e610fc5565b5b5091905056fea264697066735822122017803ac68a9be4ddab7600e9f1cdbafb3ddbeea829815c52ae1497ae905e0d0364736f6c634300080d0033",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Erc20MetaData.Bin instead.
var Erc20Bin = Erc20MetaData.Bin

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Erc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_Erc20 *Erc20Caller) AllowanceCalled(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowanceCalled", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_Erc20 *Erc20Session) AllowanceCalled(arg0 common.Address) (common.Address, error) {
	return _Erc20.Contract.AllowanceCalled(&_Erc20.CallOpts, arg0)
}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_Erc20 *Erc20CallerSession) AllowanceCalled(arg0 common.Address) (common.Address, error) {
	return _Erc20.Contract.AllowanceCalled(&_Erc20.CallOpts, arg0)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) ApproveCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "approveCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Session) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.ApproveCalled(&_Erc20.CallOpts, arg0)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.ApproveCalled(&_Erc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) TransferCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20Session) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.TransferCalled(&_Erc20.CallOpts, arg0)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.TransferCalled(&_Erc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20Caller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20Session) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Erc20.Contract.TransferFromCalled(&_Erc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_Erc20 *Erc20CallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _Erc20.Contract.TransferFromCalled(&_Erc20.CallOpts, arg0)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_Erc20 *Erc20Transactor) Allowance(opts *bind.TransactOpts, o common.Address, s common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "allowance", o, s)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_Erc20 *Erc20Session) Allowance(o common.Address, s common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.Allowance(&_Erc20.TransactOpts, o, s)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_Erc20 *Erc20TransactorSession) Allowance(o common.Address, s common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.Allowance(&_Erc20.TransactOpts, o, s)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_Erc20 *Erc20Transactor) AllowanceReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "allowanceReturns", n)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_Erc20 *Erc20Session) AllowanceReturns(n *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.AllowanceReturns(&_Erc20.TransactOpts, n)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_Erc20 *Erc20TransactorSession) AllowanceReturns(n *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.AllowanceReturns(&_Erc20.TransactOpts, n)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, s, a)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) ApproveReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approveReturns", b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20Session) ApproveReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.ApproveReturns(&_Erc20.TransactOpts, b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) ApproveReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.ApproveReturns(&_Erc20.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_Erc20 *Erc20Transactor) BalanceOf(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "balanceOf", t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.TransactOpts, t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_Erc20 *Erc20TransactorSession) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.TransactOpts, t)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Erc20 *Erc20Transactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Erc20 *Erc20Session) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.BalanceOfReturns(&_Erc20.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_Erc20 *Erc20TransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.BalanceOfReturns(&_Erc20.TransactOpts, b)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_Erc20 *Erc20Transactor) DecimalsReturns(opts *bind.TransactOpts, n uint8) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "decimalsReturns", n)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_Erc20 *Erc20Session) DecimalsReturns(n uint8) (*types.Transaction, error) {
	return _Erc20.Contract.DecimalsReturns(&_Erc20.TransactOpts, n)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_Erc20 *Erc20TransactorSession) DecimalsReturns(n uint8) (*types.Transaction, error) {
	return _Erc20.Contract.DecimalsReturns(&_Erc20.TransactOpts, n)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_Erc20 *Erc20Transactor) NameReturns(opts *bind.TransactOpts, s string) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "nameReturns", s)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_Erc20 *Erc20Session) NameReturns(s string) (*types.Transaction, error) {
	return _Erc20.Contract.NameReturns(&_Erc20.TransactOpts, s)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_Erc20 *Erc20TransactorSession) NameReturns(s string) (*types.Transaction, error) {
	return _Erc20.Contract.NameReturns(&_Erc20.TransactOpts, s)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_Erc20 *Erc20Transactor) SymbolReturns(opts *bind.TransactOpts, s string) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "symbolReturns", s)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_Erc20 *Erc20Session) SymbolReturns(s string) (*types.Transaction, error) {
	return _Erc20.Contract.SymbolReturns(&_Erc20.TransactOpts, s)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_Erc20 *Erc20TransactorSession) SymbolReturns(s string) (*types.Transaction, error) {
	return _Erc20.Contract.SymbolReturns(&_Erc20.TransactOpts, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20Session) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFromReturns(&_Erc20.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFromReturns(&_Erc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20Transactor) TransferReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferReturns", b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20Session) TransferReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferReturns(&_Erc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_Erc20 *Erc20TransactorSession) TransferReturns(b bool) (*types.Transaction, error) {
	return _Erc20.Contract.TransferReturns(&_Erc20.TransactOpts, b)
}
