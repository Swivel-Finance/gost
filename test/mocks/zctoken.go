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

// ZcTokenMetaData contains all meta data concerning the ZcToken contract.
var ZcTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"maturityReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000ced38038062000ced83398101604081905262000034916200027c565b600380546001600160a01b03808a16610100026001600160a81b031990921660ff8c1617919091179091556004879055600580548783166001600160a01b03199182161790915560068054928716929091169190911790558251620000a1906007906020860190620000da565b508151620000b7906008906020850190620000da565b506009805460ff191660ff92909216919091179055506200038a95505050505050565b828054620000e8906200034e565b90600052602060002090601f0160209004810192826200010c576000855562000157565b82601f106200012757805160ff191683800117855562000157565b8280016001018555821562000157579182015b82811115620001575782518255916020019190600101906200013a565b506200016592915062000169565b5090565b5b808211156200016557600081556001016200016a565b805160ff811681146200019257600080fd5b919050565b80516001600160a01b03811681146200019257600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001d757600080fd5b81516001600160401b0380821115620001f457620001f4620001af565b604051601f8301601f19908116603f011681019082821181831017156200021f576200021f620001af565b816040528381526020925086838588010111156200023c57600080fd5b600091505b8382101562000260578582018301518183018401529082019062000241565b83821115620002725760008385830101525b9695505050505050565b600080600080600080600080610100898b0312156200029a57600080fd5b620002a58962000180565b9750620002b560208a0162000197565b965060408901519550620002cc60608a0162000197565b9450620002dc60808a0162000197565b60a08a01519094506001600160401b0380821115620002fa57600080fd5b620003088c838d01620001c5565b945060c08b01519150808211156200031f57600080fd5b506200032e8b828c01620001c5565b9250506200033f60e08a0162000180565b90509295985092959890939650565b600181811c908216806200036357607f821691505b6020821081036200038457634e487b7160e01b600052602260045260246000fd5b50919050565b610953806200039a6000396000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80638ce74426116100cd578063bba0ad3911610081578063e7ba677411610066578063e7ba6774146104c3578063ee4db5701461051d578063fdfe5f4d1461053d57600080fd5b8063bba0ad3914610437578063e541efa21461045757600080fd5b80639dc29fac116100b25780639dc29fac1461039b578063b4c4a4c8146103df578063b9bb928c146103f257600080fd5b80638ce744261461038657806395d89b411461039357600080fd5b806340c10f19116101245780636521b96a116101095780636521b96a146102fc57806369e527da146103435780636f307dc31461036357600080fd5b806340c10f19146102a257806352bc9430146102e757600080fd5b806323b872dd1161015557806323b872dd146101a15780632ba29d381461023e578063313ce5671461028357600080fd5b806306fdde0314610171578063204f83f91461018f575b600080fd5b610179610583565b60405161018691906106e4565b60405180910390f35b6004545b604051908152602001610186565b61022e6101af366004610780565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff93841681526020808201938452948416600090815260029095529320925183547fffffffffffffffffffffffff0000000000000000000000000000000000000000169216919091178255516001909101556009546301000000900460ff1690565b6040519015158152602001610186565b60065461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610186565b6009546102909060ff1681565b60405160ff9091168152602001610186565b61022e6102b03660046107bc565b73ffffffffffffffffffffffffffffffffffffffff9190911660009081526001602052604090205560095462010000900460ff1690565b6102fa6102f53660046107e6565b610611565b005b6102fa61030a366004610854565b600980549115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff909216919091179055565b60055461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b600354610100900473ffffffffffffffffffffffffffffffffffffffff1661025e565b6003546102909060ff1681565b6101796106d7565b61022e6103a93660046107bc565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260208190526040902055600954610100900460ff1690565b6102fa6103ed36600461087d565b600455565b6102fa610400366004610854565b60098054911515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055565b610193610445366004610896565b60006020819052908152604090205481565b610497610465366004610896565b6002602052600090815260409020805460019091015473ffffffffffffffffffffffffffffffffffffffff9091169082565b6040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915201610186565b6102fa6104d1366004610896565b6003805473ffffffffffffffffffffffffffffffffffffffff909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b61019361052b366004610896565b60016020526000908152604090205481565b6102fa61054b366004610854565b6009805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055565b60078054610590906108b1565b80601f01602080910402602001604051908101604052809291908181526020018280546105bc906108b1565b80156106095780601f106105de57610100808354040283529160200191610609565b820191906000526020600020905b8154815290600101906020018083116105ec57829003601f168201915b505050505081565b6006546040517f52bc943000000000000000000000000000000000000000000000000000000000815260ff8816600482015273ffffffffffffffffffffffffffffffffffffffff8781166024830152604482018790528581166064830152848116608483015260a48201849052909116906352bc94309060c4016020604051808303816000875af11580156106aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ce9190610904565b50505050505050565b60088054610590906108b1565b600060208083528351808285015260005b81811015610711578581018301518582016040015282016106f5565b81811115610723576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461077b57600080fd5b919050565b60008060006060848603121561079557600080fd5b61079e84610757565b92506107ac60208501610757565b9150604084013590509250925092565b600080604083850312156107cf57600080fd5b6107d883610757565b946020939093013593505050565b60008060008060008060c087890312156107ff57600080fd5b863560ff8116811461081057600080fd5b955061081e60208801610757565b94506040870135935061083360608801610757565b925061084160808801610757565b915060a087013590509295509295509295565b60006020828403121561086657600080fd5b8135801515811461087657600080fd5b9392505050565b60006020828403121561088f57600080fd5b5035919050565b6000602082840312156108a857600080fd5b61087682610757565b600181811c908216806108c557607f821691505b6020821081036108fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561091657600080fd5b505191905056fea2646970667358221220ab82732e8b4d0f87d0e20dedbcdfd9a4708dd121141f10519c014c70dccfd0f764736f6c634300080d0033",
}

// ZcTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ZcTokenMetaData.ABI instead.
var ZcTokenABI = ZcTokenMetaData.ABI

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZcTokenMetaData.Bin instead.
var ZcTokenBin = ZcTokenMetaData.Bin

// DeployZcToken deploys a new Ethereum contract, binding an instance of ZcToken to it.
func DeployZcToken(auth *bind.TransactOpts, backend bind.ContractBackend, p uint8, u common.Address, m *big.Int, c common.Address, r common.Address, n string, s string, d uint8) (common.Address, *types.Transaction, *ZcToken, error) {
	parsed, err := ZcTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZcTokenBin), backend, p, u, m, c, r, n, s, d)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// ZcToken is an auto generated Go binding around an Ethereum contract.
type ZcToken struct {
	ZcTokenCaller     // Read-only binding to the contract
	ZcTokenTransactor // Write-only binding to the contract
	ZcTokenFilterer   // Log filterer for contract events
}

// ZcTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZcTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZcTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZcTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZcTokenSession struct {
	Contract     *ZcToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZcTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZcTokenCallerSession struct {
	Contract *ZcTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZcTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZcTokenTransactorSession struct {
	Contract     *ZcTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZcTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZcTokenRaw struct {
	Contract *ZcToken // Generic contract binding to access the raw methods on
}

// ZcTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZcTokenCallerRaw struct {
	Contract *ZcTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZcTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZcTokenTransactorRaw struct {
	Contract *ZcTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZcToken creates a new instance of ZcToken, bound to a specific deployed contract.
func NewZcToken(address common.Address, backend bind.ContractBackend) (*ZcToken, error) {
	contract, err := bindZcToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// NewZcTokenCaller creates a new read-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenCaller(address common.Address, caller bind.ContractCaller) (*ZcTokenCaller, error) {
	contract, err := bindZcToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenCaller{contract: contract}, nil
}

// NewZcTokenTransactor creates a new write-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZcTokenTransactor, error) {
	contract, err := bindZcToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransactor{contract: contract}, nil
}

// NewZcTokenFilterer creates a new log filterer instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZcTokenFilterer, error) {
	contract, err := bindZcToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZcTokenFilterer{contract: contract}, nil
}

// bindZcToken binds a generic wrapper to an already deployed contract.
func bindZcToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.ZcTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transact(opts, method, params...)
}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BurnCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "burnCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) BurnCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BurnCalled(&_ZcToken.CallOpts, arg0)
}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BurnCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BurnCalled(&_ZcToken.CallOpts, arg0)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenSession) CToken() (common.Address, error) {
	return _ZcToken.Contract.CToken(&_ZcToken.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenCallerSession) CToken() (common.Address, error) {
	return _ZcToken.Contract.CToken(&_ZcToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCallerSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) MintCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "mintCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MintCalled(&_ZcToken.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) MintCalled(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MintCalled(&_ZcToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenCaller) Protocol(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenSession) Protocol() (uint8, error) {
	return _ZcToken.Contract.Protocol(&_ZcToken.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenCallerSession) Protocol() (uint8, error) {
	return _ZcToken.Contract.Protocol(&_ZcToken.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenCaller) Redeemer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "redeemer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenSession) Redeemer() (common.Address, error) {
	return _ZcToken.Contract.Redeemer(&_ZcToken.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Redeemer() (common.Address, error) {
	return _ZcToken.Contract.Redeemer(&_ZcToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_ZcToken *ZcTokenCaller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "transferFromCalled", arg0)

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
func (_ZcToken *ZcTokenSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _ZcToken.Contract.TransferFromCalled(&_ZcToken.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_ZcToken *ZcTokenCallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _ZcToken.Contract.TransferFromCalled(&_ZcToken.CallOpts, arg0)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns()
func (_ZcToken *ZcTokenTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "authRedeem", p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns()
func (_ZcToken *ZcTokenSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.AuthRedeem(&_ZcToken.TransactOpts, p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns()
func (_ZcToken *ZcTokenTransactorSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.AuthRedeem(&_ZcToken.TransactOpts, p, u, m, f, t, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Burn(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "burn", f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xb9bb928c.
//
// Solidity: function burnReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactor) BurnReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "burnReturns", b)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xb9bb928c.
//
// Solidity: function burnReturns(bool b) returns()
func (_ZcToken *ZcTokenSession) BurnReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.BurnReturns(&_ZcToken.TransactOpts, b)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xb9bb928c.
//
// Solidity: function burnReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactorSession) BurnReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.BurnReturns(&_ZcToken.TransactOpts, b)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_ZcToken *ZcTokenTransactor) MaturityReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "maturityReturns", n)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_ZcToken *ZcTokenSession) MaturityReturns(n *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.MaturityReturns(&_ZcToken.TransactOpts, n)
}

// MaturityReturns is a paid mutator transaction binding the contract method 0xb4c4a4c8.
//
// Solidity: function maturityReturns(uint256 n) returns()
func (_ZcToken *ZcTokenTransactorSession) MaturityReturns(n *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.MaturityReturns(&_ZcToken.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Mint(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mint", f, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Mint(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, f, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Mint(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, f, a)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactor) MintReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mintReturns", b)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool b) returns()
func (_ZcToken *ZcTokenSession) MintReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.MintReturns(&_ZcToken.TransactOpts, b)
}

// MintReturns is a paid mutator transaction binding the contract method 0xfdfe5f4d.
//
// Solidity: function mintReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactorSession) MintReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.MintReturns(&_ZcToken.TransactOpts, b)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ZcToken *ZcTokenSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFromReturns(&_ZcToken.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_ZcToken *ZcTokenTransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFromReturns(&_ZcToken.TransactOpts, b)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_ZcToken *ZcTokenTransactor) UnderlyingReturns(opts *bind.TransactOpts, u common.Address) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "underlyingReturns", u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_ZcToken *ZcTokenSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.UnderlyingReturns(&_ZcToken.TransactOpts, u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_ZcToken *ZcTokenTransactorSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.UnderlyingReturns(&_ZcToken.TransactOpts, u)
}
