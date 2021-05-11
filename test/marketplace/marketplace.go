// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// MarketPlaceABI is the input ABI used to generate the binding from.
const MarketPlaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturedOn\",\"type\":\"uint256\"}],\"name\":\"Mature\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"matureMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"reddemVaultInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561005057600080fd5b50610806806100606000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80633271e5f31161005b5780633271e5f31461014f5780635db0ae581461017f578063c86adf7c146101af578063f851a440146101df57610088565b8063023b4d721461008d578063154e0f2e146100bd57806317b3bba7146100ed57806327ee93be1461011f575b600080fd5b6100a760048036038101906100a291906104ee565b6101fd565b6040516100b49190610654565b60405180910390f35b6100d760048036038101906100d2919061049f565b61029d565b6040516100e49190610654565b60405180910390f35b61010760048036038101906101029190610463565b6102aa565b6040516101169392919061061d565b60405180910390f35b61013960048036038101906101349190610463565b610341565b6040516101469190610654565b60405180910390f35b61016960048036038101906101649190610463565b610370565b6040516101769190610654565b60405180910390f35b61019960048036038101906101949190610463565b61037c565b6040516101a69190610654565b60405180910390f35b6101c960048036038101906101c49190610463565b610388565b6040516101d6919061068f565b60405180910390f35b6101e76103ad565b6040516101f49190610602565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102869061066f565b60405180910390fd5b600191505095945050505050565b6000600190509392505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006001905092915050565b60006001905092915050565b6003602052816000526040600020602052806000526040600020600091509150505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006103e46103df846106db565b6106aa565b9050828152602081018484840111156103fc57600080fd5b610407848285610764565b509392505050565b60008135905061041e816107a2565b92915050565b600082601f83011261043557600080fd5b81356104458482602086016103d1565b91505092915050565b60008135905061045d816107b9565b92915050565b6000806040838503121561047657600080fd5b60006104848582860161040f565b92505060206104958582860161044e565b9150509250929050565b6000806000606084860312156104b457600080fd5b60006104c28682870161040f565b93505060206104d38682870161044e565b92505060406104e48682870161044e565b9150509250925092565b600080600080600060a0868803121561050657600080fd5b600086013567ffffffffffffffff81111561052057600080fd5b61052c88828901610424565b955050602086013567ffffffffffffffff81111561054957600080fd5b61055588828901610424565b94505060406105668882890161040f565b93505060606105778882890161044e565b92505060806105888882890161040f565b9150509295509295909350565b61059e8161071c565b82525050565b6105ad8161072e565b82525050565b60006105c060148361070b565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b6105fc8161075a565b82525050565b60006020820190506106176000830184610595565b92915050565b60006060820190506106326000830186610595565b61063f6020830185610595565b61064c6040830184610595565b949350505050565b600060208201905061066960008301846105a4565b92915050565b60006020820190508181036000830152610688816105b3565b9050919050565b60006020820190506106a460008301846105f3565b92915050565b6000604051905081810181811067ffffffffffffffff821117156106d1576106d0610773565b5b8060405250919050565b600067ffffffffffffffff8211156106f6576106f5610773565b5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b60006107278261073a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107ab8161071c565b81146107b657600080fd5b50565b6107c28161075a565b81146107cd57600080fd5b5056fea264697066735822122072a95a9bfd7aacaed519c3753ac3d4bc535905c06c0401f411f3d0ec3b21b21264736f6c63430008000033"

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketPlaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceCaller) Markets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "markets", arg0, arg1)

	outstruct := new(struct {
		CTokenAddr  common.Address
		ZcTokenAddr common.Address
		VaultAddr   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CTokenAddr = out[0].(common.Address)
	outstruct.ZcTokenAddr = out[1].(common.Address)
	outstruct.VaultAddr = out[2].(common.Address)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceSession) Markets(arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceCallerSession) Markets(arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1)
}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceCaller) Mature(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mature", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceSession) Mature(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MarketPlace.Contract.Mature(&_MarketPlace.CallOpts, arg0, arg1)
}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceCallerSession) Mature(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MarketPlace.Contract.Mature(&_MarketPlace.CallOpts, arg0, arg1)
}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) MaturityRate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "maturityRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) MaturityRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketPlace.Contract.MaturityRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) MaturityRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketPlace.Contract.MaturityRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x023b4d72.
//
// Solidity: function createMarket(string n, string s, address u, uint256 m, address c) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CreateMarket(opts *bind.TransactOpts, n string, s string, u common.Address, m *big.Int, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "createMarket", n, s, u, m, c)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x023b4d72.
//
// Solidity: function createMarket(string n, string s, address u, uint256 m, address c) returns(bool)
func (_MarketPlace *MarketPlaceSession) CreateMarket(n string, s string, u common.Address, m *big.Int, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, n, s, u, m, c)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x023b4d72.
//
// Solidity: function createMarket(string n, string s, address u, uint256 m, address c) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CreateMarket(n string, s string, u common.Address, m *big.Int, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, n, s, u, m, c)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MatureMarket(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "matureMarket", u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceSession) MatureMarket(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MatureMarket(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) ReddemVaultInterest(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "reddemVaultInterest", u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceSession) ReddemVaultInterest(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ReddemVaultInterest(&_MarketPlace.TransactOpts, u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) ReddemVaultInterest(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ReddemVaultInterest(&_MarketPlace.TransactOpts, u, m)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) RedeemZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemZcToken", u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, u, m, a)
}

// MarketPlaceCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the MarketPlace contract.
type MarketPlaceCreateIterator struct {
	Event *MarketPlaceCreate // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCreate)
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
		it.Event = new(MarketPlaceCreate)
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
func (it *MarketPlaceCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCreate represents a Create event raised by the MarketPlace contract.
type MarketPlaceCreate struct {
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	ZcToken    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address underlying, uint256 maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) FilterCreate(opts *bind.FilterOpts) (*MarketPlaceCreateIterator, error) {

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCreateIterator{contract: _MarketPlace.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address underlying, uint256 maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *MarketPlaceCreate) (event.Subscription, error) {

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCreate)
				if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address underlying, uint256 maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) ParseCreate(log types.Log) (*MarketPlaceCreate, error) {
	event := new(MarketPlaceCreate)
	if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceMatureIterator is returned from FilterMature and is used to iterate over the raw logs and unpacked data for Mature events raised by the MarketPlace contract.
type MarketPlaceMatureIterator struct {
	Event *MarketPlaceMature // Event containing the contract specifics and raw log

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
func (it *MarketPlaceMatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceMature)
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
		it.Event = new(MarketPlaceMature)
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
func (it *MarketPlaceMatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceMatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceMature represents a Mature event raised by the MarketPlace contract.
type MarketPlaceMature struct {
	Underlying   common.Address
	Maturity     *big.Int
	MaturityRate *big.Int
	MaturedOn    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMature is a free log retrieval operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 maturedOn)
func (_MarketPlace *MarketPlaceFilterer) FilterMature(opts *bind.FilterOpts) (*MarketPlaceMatureIterator, error) {

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Mature")
	if err != nil {
		return nil, err
	}
	return &MarketPlaceMatureIterator{contract: _MarketPlace.contract, event: "Mature", logs: logs, sub: sub}, nil
}

// WatchMature is a free log subscription operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 maturedOn)
func (_MarketPlace *MarketPlaceFilterer) WatchMature(opts *bind.WatchOpts, sink chan<- *MarketPlaceMature) (event.Subscription, error) {

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Mature")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceMature)
				if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
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

// ParseMature is a log parse operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 maturedOn)
func (_MarketPlace *MarketPlaceFilterer) ParseMature(log types.Log) (*MarketPlaceMature, error) {
	event := new(MarketPlaceMature)
	if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
