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

// APWineTokenABI is the input ABI used to generate the binding from.
const APWineTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOfCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingOfIBTAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getUnderlyingOfIBTAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// APWineTokenBin is the compiled bytecode used for deploying new contracts.
var APWineTokenBin = "0x608060405234801561001057600080fd5b5061081b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806370a082311161007157806370a082311461014e578063a9059cbb1461017e578063c1d2e9a1146101ae578063c6ec27bf146101de578063dea1a7e2146101fc578063e541efa21461021a576100a9565b806323b872dd146100ae57806339100838146100de57806342b6cdbc146100fa5780636521b96a1461011657806366954b2114610132575b600080fd5b6100c860048036038101906100c391906105ec565b61024b565b6040516100d5919061065a565b60405180910390f35b6100f860048036038101906100f39190610675565b610345565b005b610114600480360381019061010f91906106ce565b61034f565b005b610130600480360381019061012b91906106ce565b61036c565b005b61014c600480360381019061014791906106fb565b610389565b005b610168600480360381019061016391906106fb565b6103cd565b6040516101759190610737565b60405180910390f35b61019860048036038101906101939190610752565b61041a565b6040516101a5919061065a565b60405180910390f35b6101c860048036038101906101c391906106fb565b610477565b6040516101d59190610737565b60405180910390f35b6101e661048f565b6040516101f391906107a1565b60405180910390f35b6102046104b9565b60405161021191906107a1565b60405180910390f35b610234600480360381019061022f91906106fb565b6104df565b6040516102429291906107bc565b60405180910390f35b6000610255610523565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600560019054906101000a900460ff169150509392505050565b8060048190555050565b80600560006101000a81548160ff02191690831515021790555050565b80600560016101000a81548160ff02191690831515021790555050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506004549050919050565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600560009054906101000a900460ff16905092915050565b60006020528060005260406000206000915090505481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061058382610558565b9050919050565b61059381610578565b811461059e57600080fd5b50565b6000813590506105b08161058a565b92915050565b6000819050919050565b6105c9816105b6565b81146105d457600080fd5b50565b6000813590506105e6816105c0565b92915050565b60008060006060848603121561060557610604610553565b5b6000610613868287016105a1565b9350506020610624868287016105a1565b9250506040610635868287016105d7565b9150509250925092565b60008115159050919050565b6106548161063f565b82525050565b600060208201905061066f600083018461064b565b92915050565b60006020828403121561068b5761068a610553565b5b6000610699848285016105d7565b91505092915050565b6106ab8161063f565b81146106b657600080fd5b50565b6000813590506106c8816106a2565b92915050565b6000602082840312156106e4576106e3610553565b5b60006106f2848285016106b9565b91505092915050565b60006020828403121561071157610710610553565b5b600061071f848285016105a1565b91505092915050565b610731816105b6565b82525050565b600060208201905061074c6000830184610728565b92915050565b6000806040838503121561076957610768610553565b5b6000610777858286016105a1565b9250506020610788858286016105d7565b9150509250929050565b61079b81610578565b82525050565b60006020820190506107b66000830184610792565b92915050565b60006040820190506107d16000830185610792565b6107de6020830184610728565b939250505056fea2646970667358221220bb685df62516fd767946a0cbc6e091787e85aff27a27e3c50173ab042a478ecc64736f6c634300080d0033"

// DeployAPWineToken deploys a new Ethereum contract, binding an instance of APWineToken to it.
func DeployAPWineToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *APWineToken, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(APWineTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// APWineToken is an auto generated Go binding around an Ethereum contract.
type APWineToken struct {
	APWineTokenCaller     // Read-only binding to the contract
	APWineTokenTransactor // Write-only binding to the contract
	APWineTokenFilterer   // Log filterer for contract events
}

// APWineTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type APWineTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type APWineTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type APWineTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// APWineTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type APWineTokenSession struct {
	Contract     *APWineToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// APWineTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type APWineTokenCallerSession struct {
	Contract *APWineTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// APWineTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type APWineTokenTransactorSession struct {
	Contract     *APWineTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// APWineTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type APWineTokenRaw struct {
	Contract *APWineToken // Generic contract binding to access the raw methods on
}

// APWineTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type APWineTokenCallerRaw struct {
	Contract *APWineTokenCaller // Generic read-only contract binding to access the raw methods on
}

// APWineTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type APWineTokenTransactorRaw struct {
	Contract *APWineTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAPWineToken creates a new instance of APWineToken, bound to a specific deployed contract.
func NewAPWineToken(address common.Address, backend bind.ContractBackend) (*APWineToken, error) {
	contract, err := bindAPWineToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &APWineToken{APWineTokenCaller: APWineTokenCaller{contract: contract}, APWineTokenTransactor: APWineTokenTransactor{contract: contract}, APWineTokenFilterer: APWineTokenFilterer{contract: contract}}, nil
}

// NewAPWineTokenCaller creates a new read-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenCaller(address common.Address, caller bind.ContractCaller) (*APWineTokenCaller, error) {
	contract, err := bindAPWineToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenCaller{contract: contract}, nil
}

// NewAPWineTokenTransactor creates a new write-only instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*APWineTokenTransactor, error) {
	contract, err := bindAPWineToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &APWineTokenTransactor{contract: contract}, nil
}

// NewAPWineTokenFilterer creates a new log filterer instance of APWineToken, bound to a specific deployed contract.
func NewAPWineTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*APWineTokenFilterer, error) {
	contract, err := bindAPWineToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &APWineTokenFilterer{contract: contract}, nil
}

// bindAPWineToken binds a generic wrapper to an already deployed contract.
func bindAPWineToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(APWineTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.APWineTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.APWineTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_APWineToken *APWineTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _APWineToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_APWineToken *APWineTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_APWineToken *APWineTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _APWineToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_APWineToken *APWineTokenCaller) BalanceOfCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "balanceOfCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_APWineToken *APWineTokenSession) BalanceOfCalled() (common.Address, error) {
	return _APWineToken.Contract.BalanceOfCalled(&_APWineToken.CallOpts)
}

// BalanceOfCalled is a free data retrieval call binding the contract method 0xdea1a7e2.
//
// Solidity: function balanceOfCalled() view returns(address)
func (_APWineToken *APWineTokenCallerSession) BalanceOfCalled() (common.Address, error) {
	return _APWineToken.Contract.BalanceOfCalled(&_APWineToken.CallOpts)
}

// GetUnderlyingOfIBTAddress is a free data retrieval call binding the contract method 0xc6ec27bf.
//
// Solidity: function getUnderlyingOfIBTAddress() view returns(address)
func (_APWineToken *APWineTokenCaller) GetUnderlyingOfIBTAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "getUnderlyingOfIBTAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUnderlyingOfIBTAddress is a free data retrieval call binding the contract method 0xc6ec27bf.
//
// Solidity: function getUnderlyingOfIBTAddress() view returns(address)
func (_APWineToken *APWineTokenSession) GetUnderlyingOfIBTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetUnderlyingOfIBTAddress(&_APWineToken.CallOpts)
}

// GetUnderlyingOfIBTAddress is a free data retrieval call binding the contract method 0xc6ec27bf.
//
// Solidity: function getUnderlyingOfIBTAddress() view returns(address)
func (_APWineToken *APWineTokenCallerSession) GetUnderlyingOfIBTAddress() (common.Address, error) {
	return _APWineToken.Contract.GetUnderlyingOfIBTAddress(&_APWineToken.CallOpts)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_APWineToken *APWineTokenCaller) TransferCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "transferCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_APWineToken *APWineTokenSession) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _APWineToken.Contract.TransferCalled(&_APWineToken.CallOpts, arg0)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_APWineToken *APWineTokenCallerSession) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _APWineToken.Contract.TransferCalled(&_APWineToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _APWineToken.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})

	outstruct.To = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _APWineToken.Contract.TransferFromCalled(&_APWineToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_APWineToken *APWineTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _APWineToken.Contract.TransferFromCalled(&_APWineToken.CallOpts, arg0)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_APWineToken *APWineTokenTransactor) BalanceOf(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "balanceOf", t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_APWineToken *APWineTokenSession) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOf(&_APWineToken.TransactOpts, t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_APWineToken *APWineTokenTransactorSession) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOf(&_APWineToken.TransactOpts, t)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_APWineToken *APWineTokenTransactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_APWineToken *APWineTokenSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOfReturns(&_APWineToken.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_APWineToken *APWineTokenTransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.BalanceOfReturns(&_APWineToken.TransactOpts, b)
}

// GetUnderlyingOfIBTAddressReturns is a paid mutator transaction binding the contract method 0x66954b21.
//
// Solidity: function getUnderlyingOfIBTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactor) GetUnderlyingOfIBTAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "getUnderlyingOfIBTAddressReturns", a)
}

// GetUnderlyingOfIBTAddressReturns is a paid mutator transaction binding the contract method 0x66954b21.
//
// Solidity: function getUnderlyingOfIBTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenSession) GetUnderlyingOfIBTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetUnderlyingOfIBTAddressReturns(&_APWineToken.TransactOpts, a)
}

// GetUnderlyingOfIBTAddressReturns is a paid mutator transaction binding the contract method 0x66954b21.
//
// Solidity: function getUnderlyingOfIBTAddressReturns(address a) returns()
func (_APWineToken *APWineTokenTransactorSession) GetUnderlyingOfIBTAddressReturns(a common.Address) (*types.Transaction, error) {
	return _APWineToken.Contract.GetUnderlyingOfIBTAddressReturns(&_APWineToken.TransactOpts, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactor) Transfer(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transfer", t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.Transfer(&_APWineToken.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactorSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.Transfer(&_APWineToken.TransactOpts, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFrom(&_APWineToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_APWineToken *APWineTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFrom(&_APWineToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFromReturns(&_APWineToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferFromReturns(&_APWineToken.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactor) TransferReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _APWineToken.contract.Transact(opts, "transferReturns", b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_APWineToken *APWineTokenSession) TransferReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferReturns(&_APWineToken.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_APWineToken *APWineTokenTransactorSession) TransferReturns(b bool) (*types.Transaction, error) {
	return _APWineToken.Contract.TransferReturns(&_APWineToken.TransactOpts, b)
}
