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
	Bin: "0x608060405234801561001057600080fd5b5061121f806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806395d89b41116100ad578063ad64dcc511610071578063ad64dcc514610353578063af599a291461036f578063c1d2e9a11461038b578063dd62ed3e146103bb578063e541efa2146103eb5761012c565b806395d89b411461029d5780639dd0ff37146102bb578063a18729a3146102d7578063a5b836bf146102f3578063a9059cbb146103235761012c565b806339100838116100f457806339100838146101e957806342b6cdbc146102055780636521b96a146102215780636581d5431461023d57806370a082311461026d5761012c565b806306fdde0314610131578063095ea7b31461014f5780630ab2519e1461017f57806323b872dd1461019b578063313ce567146101cb575b600080fd5b61013961041c565b60405161014691906109fa565b60405180910390f35b61016960048036038101906101649190610ac4565b6104ae565b6040516101769190610b1f565b60405180910390f35b61019960048036038101906101949190610c6f565b61050b565b005b6101b560048036038101906101b09190610cb8565b61051e565b6040516101c29190610b1f565b60405180910390f35b6101d3610618565b6040516101e09190610d27565b60405180910390f35b61020360048036038101906101fe9190610d42565b61062f565b005b61021f600480360381019061021a9190610d9b565b610639565b005b61023b60048036038101906102369190610d9b565b610656565b005b61025760048036038101906102529190610dc8565b610673565b6040516102649190610e04565b60405180910390f35b61028760048036038101906102829190610dc8565b61068b565b6040516102949190610e04565b60405180910390f35b6102a56106d8565b6040516102b291906109fa565b60405180910390f35b6102d560048036038101906102d09190610d9b565b61076a565b005b6102f160048036038101906102ec9190610e4b565b610787565b005b61030d60048036038101906103089190610dc8565b6107a5565b60405161031a9190610e87565b60405180910390f35b61033d60048036038101906103389190610ac4565b6107d8565b60405161034a9190610b1f565b60405180910390f35b61036d60048036038101906103689190610d42565b610836565b005b61038960048036038101906103849190610c6f565b610840565b005b6103a560048036038101906103a09190610dc8565b610853565b6040516103b29190610e04565b60405180910390f35b6103d560048036038101906103d09190610ea2565b61086b565b6040516103e29190610e04565b60405180910390f35b61040560048036038101906104009190610dc8565b6108f6565b604051610413929190610ee2565b60405180910390f35b60606005805461042b90610f3a565b80601f016020809104026020016040519081016040528092919081815260200182805461045790610f3a565b80156104a45780601f10610479576101008083540402835291602001916104a4565b820191906000526020600020905b81548152906001019060200180831161048757829003601f168201915b5050505050905090565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600760019054906101000a900460ff16905092915050565b806005908161051a9190611117565b5050565b600061052861093a565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600a60019054906101000a900460ff169150509392505050565b6000600760009054906101000a900460ff16905090565b8060088190555050565b80600a60006101000a81548160ff02191690831515021790555050565b80600a60016101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915090505481565b600081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008549050919050565b6060600680546106e790610f3a565b80601f016020809104026020016040519081016040528092919081815260200182805461071390610f3a565b80156107605780601f1061073557610100808354040283529160200191610760565b820191906000526020600020905b81548152906001019060200180831161074357829003601f168201915b5050505050905090565b80600760016101000a81548160ff02191690831515021790555050565b80600760006101000a81548160ff021916908360ff16021790555050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600a60009054906101000a900460ff16905092915050565b8060098190555050565b806006908161084f9190611117565b5050565b60016020528060005260406000206000915090505481565b600081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600954905092915050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081519050919050565b600082825260208201905092915050565b60005b838110156109a4578082015181840152602081019050610989565b60008484015250505050565b6000601f19601f8301169050919050565b60006109cc8261096a565b6109d68185610975565b93506109e6818560208601610986565b6109ef816109b0565b840191505092915050565b60006020820190508181036000830152610a1481846109c1565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a5b82610a30565b9050919050565b610a6b81610a50565b8114610a7657600080fd5b50565b600081359050610a8881610a62565b92915050565b6000819050919050565b610aa181610a8e565b8114610aac57600080fd5b50565b600081359050610abe81610a98565b92915050565b60008060408385031215610adb57610ada610a26565b5b6000610ae985828601610a79565b9250506020610afa85828601610aaf565b9150509250929050565b60008115159050919050565b610b1981610b04565b82525050565b6000602082019050610b346000830184610b10565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b7c826109b0565b810181811067ffffffffffffffff82111715610b9b57610b9a610b44565b5b80604052505050565b6000610bae610a1c565b9050610bba8282610b73565b919050565b600067ffffffffffffffff821115610bda57610bd9610b44565b5b610be3826109b0565b9050602081019050919050565b82818337600083830152505050565b6000610c12610c0d84610bbf565b610ba4565b905082815260208101848484011115610c2e57610c2d610b3f565b5b610c39848285610bf0565b509392505050565b600082601f830112610c5657610c55610b3a565b5b8135610c66848260208601610bff565b91505092915050565b600060208284031215610c8557610c84610a26565b5b600082013567ffffffffffffffff811115610ca357610ca2610a2b565b5b610caf84828501610c41565b91505092915050565b600080600060608486031215610cd157610cd0610a26565b5b6000610cdf86828701610a79565b9350506020610cf086828701610a79565b9250506040610d0186828701610aaf565b9150509250925092565b600060ff82169050919050565b610d2181610d0b565b82525050565b6000602082019050610d3c6000830184610d18565b92915050565b600060208284031215610d5857610d57610a26565b5b6000610d6684828501610aaf565b91505092915050565b610d7881610b04565b8114610d8357600080fd5b50565b600081359050610d9581610d6f565b92915050565b600060208284031215610db157610db0610a26565b5b6000610dbf84828501610d86565b91505092915050565b600060208284031215610dde57610ddd610a26565b5b6000610dec84828501610a79565b91505092915050565b610dfe81610a8e565b82525050565b6000602082019050610e196000830184610df5565b92915050565b610e2881610d0b565b8114610e3357600080fd5b50565b600081359050610e4581610e1f565b92915050565b600060208284031215610e6157610e60610a26565b5b6000610e6f84828501610e36565b91505092915050565b610e8181610a50565b82525050565b6000602082019050610e9c6000830184610e78565b92915050565b60008060408385031215610eb957610eb8610a26565b5b6000610ec785828601610a79565b9250506020610ed885828601610a79565b9150509250929050565b6000604082019050610ef76000830185610e78565b610f046020830184610df5565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610f5257607f821691505b602082108103610f6557610f64610f0b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610fcd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f90565b610fd78683610f90565b95508019841693508086168417925050509392505050565b6000819050919050565b600061101461100f61100a84610a8e565b610fef565b610a8e565b9050919050565b6000819050919050565b61102e83610ff9565b61104261103a8261101b565b848454610f9d565b825550505050565b600090565b61105761104a565b611062818484611025565b505050565b5b818110156110865761107b60008261104f565b600181019050611068565b5050565b601f8211156110cb5761109c81610f6b565b6110a584610f80565b810160208510156110b4578190505b6110c86110c085610f80565b830182611067565b50505b505050565b600082821c905092915050565b60006110ee600019846008026110d0565b1980831691505092915050565b600061110783836110dd565b9150826002028217905092915050565b6111208261096a565b67ffffffffffffffff81111561113957611138610b44565b5b6111438254610f3a565b61114e82828561108a565b600060209050601f831160018114611181576000841561116f578287015190505b61117985826110fb565b8655506111e1565b601f19841661118f86610f6b565b60005b828110156111b757848901518255600182019150602085019450602081019050611192565b868310156111d457848901516111d0601f8916826110dd565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220eeeeaad88fa6e330729d031c6fb5a88cfec6923f307d550351f08118ceb50d5864736f6c63430008100033",
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
