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

// SushiABI is the input ABI used to generate the binding from.
const SushiABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"o\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"p\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"r\",\"type\":\"uint256[]\"}],\"name\":\"swapExactTokensForTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SushiBin is the compiled bytecode used for deploying new contracts.
var SushiBin = "0x608060405234801561001057600080fd5b506106c7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806338ed17391461003b578063b2f2260b1461006b575b600080fd5b6100556004803603810190610050919061037f565b610087565b60405161006291906104d7565b60405180910390f35b61008560048036038101906100809190610648565b61014e565b005b606086600181905550856002819055508484600391906100a8929190610168565b5082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600581905550600080548060200260200160405190810160405280929190818152602001828054801561013d57602002820191906000526020600020905b815481526020019060010190808311610129575b505050505090509695505050505050565b8060009080519060200190610164929190610208565b5050565b8280548282559060005260206000209081019282156101f7579160200282015b828111156101f657823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610188565b5b5090506102049190610255565b5090565b828054828255906000526020600020908101928215610244579160200282015b82811115610243578251825591602001919060010190610228565b5b5090506102519190610255565b5090565b5b8082111561026e576000816000905550600101610256565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61029981610286565b81146102a457600080fd5b50565b6000813590506102b681610290565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126102e1576102e06102bc565b5b8235905067ffffffffffffffff8111156102fe576102fd6102c1565b5b60208301915083602082028301111561031a576103196102c6565b5b9250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061034c82610321565b9050919050565b61035c81610341565b811461036757600080fd5b50565b60008135905061037981610353565b92915050565b60008060008060008060a0878903121561039c5761039b61027c565b5b60006103aa89828a016102a7565b96505060206103bb89828a016102a7565b955050604087013567ffffffffffffffff8111156103dc576103db610281565b5b6103e889828a016102cb565b945094505060606103fb89828a0161036a565b925050608061040c89828a016102a7565b9150509295509295509295565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61044e81610286565b82525050565b60006104608383610445565b60208301905092915050565b6000602082019050919050565b600061048482610419565b61048e8185610424565b935061049983610435565b8060005b838110156104ca5781516104b18882610454565b97506104bc8361046c565b92505060018101905061049d565b5085935050505092915050565b600060208201905081810360008301526104f18184610479565b905092915050565b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610542826104f9565b810181811067ffffffffffffffff821117156105615761056061050a565b5b80604052505050565b6000610574610272565b90506105808282610539565b919050565b600067ffffffffffffffff8211156105a05761059f61050a565b5b602082029050602081019050919050565b60006105c46105bf84610585565b61056a565b905080838252602082019050602084028301858111156105e7576105e66102c6565b5b835b8181101561061057806105fc88826102a7565b8452602084019350506020810190506105e9565b5050509392505050565b600082601f83011261062f5761062e6102bc565b5b813561063f8482602086016105b1565b91505092915050565b60006020828403121561065e5761065d61027c565b5b600082013567ffffffffffffffff81111561067c5761067b610281565b5b6106888482850161061a565b9150509291505056fea2646970667358221220a346561c2345c63e7bcc42982feb5f4f3f21c686e92852448b6f3be17dff207564736f6c634300080d0033"

// DeploySushi deploys a new Ethereum contract, binding an instance of Sushi to it.
func DeploySushi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sushi, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SushiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sushi{SushiCaller: SushiCaller{contract: contract}, SushiTransactor: SushiTransactor{contract: contract}, SushiFilterer: SushiFilterer{contract: contract}}, nil
}

// Sushi is an auto generated Go binding around an Ethereum contract.
type Sushi struct {
	SushiCaller     // Read-only binding to the contract
	SushiTransactor // Write-only binding to the contract
	SushiFilterer   // Log filterer for contract events
}

// SushiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiSession struct {
	Contract     *Sushi            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiCallerSession struct {
	Contract *SushiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SushiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiTransactorSession struct {
	Contract     *SushiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiRaw struct {
	Contract *Sushi // Generic contract binding to access the raw methods on
}

// SushiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiCallerRaw struct {
	Contract *SushiCaller // Generic read-only contract binding to access the raw methods on
}

// SushiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiTransactorRaw struct {
	Contract *SushiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushi creates a new instance of Sushi, bound to a specific deployed contract.
func NewSushi(address common.Address, backend bind.ContractBackend) (*Sushi, error) {
	contract, err := bindSushi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sushi{SushiCaller: SushiCaller{contract: contract}, SushiTransactor: SushiTransactor{contract: contract}, SushiFilterer: SushiFilterer{contract: contract}}, nil
}

// NewSushiCaller creates a new read-only instance of Sushi, bound to a specific deployed contract.
func NewSushiCaller(address common.Address, caller bind.ContractCaller) (*SushiCaller, error) {
	contract, err := bindSushi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiCaller{contract: contract}, nil
}

// NewSushiTransactor creates a new write-only instance of Sushi, bound to a specific deployed contract.
func NewSushiTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiTransactor, error) {
	contract, err := bindSushi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiTransactor{contract: contract}, nil
}

// NewSushiFilterer creates a new log filterer instance of Sushi, bound to a specific deployed contract.
func NewSushiFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiFilterer, error) {
	contract, err := bindSushi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiFilterer{contract: contract}, nil
}

// bindSushi binds a generic wrapper to an already deployed contract.
func bindSushi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushi *SushiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushi.Contract.SushiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushi *SushiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushi.Contract.SushiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushi *SushiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushi.Contract.SushiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushi *SushiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushi *SushiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushi *SushiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushi.Contract.contract.Transact(opts, method, params...)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.contract.Transact(opts, "swapExactTokensForTokens", i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokens(&_Sushi.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 i, uint256 o, address[] p, address t, uint256 d) returns(uint256[])
func (_Sushi *SushiTransactorSession) SwapExactTokensForTokens(i *big.Int, o *big.Int, p []common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokens(&_Sushi.TransactOpts, i, o, p, t, d)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiTransactor) SwapExactTokensForTokensReturns(opts *bind.TransactOpts, r []*big.Int) (*types.Transaction, error) {
	return _Sushi.contract.Transact(opts, "swapExactTokensForTokensReturns", r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokensReturns(&_Sushi.TransactOpts, r)
}

// SwapExactTokensForTokensReturns is a paid mutator transaction binding the contract method 0xb2f2260b.
//
// Solidity: function swapExactTokensForTokensReturns(uint256[] r) returns()
func (_Sushi *SushiTransactorSession) SwapExactTokensForTokensReturns(r []*big.Int) (*types.Transaction, error) {
	return _Sushi.Contract.SwapExactTokensForTokensReturns(&_Sushi.TransactOpts, r)
}
