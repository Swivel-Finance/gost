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

// PendleDataABI is the input ABI used to generate the binding from.
const PendleDataABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"isValidOT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidOTId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidOTMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"o\",\"type\":\"bool\"}],\"name\":\"isValidOTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidOTUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"isValidXYT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidXYTId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidXYTMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"x\",\"type\":\"bool\"}],\"name\":\"isValidXYTReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isValidXYTUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"f\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"xytTokens\",\"outputs\":[{\"internalType\":\"contractIPendleYieldToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xytTokensCalledId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xytTokensCalledMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xytTokensCalledUnderlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"xytTokensReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PendleDataBin is the compiled bytecode used for deploying new contracts.
var PendleDataBin = "0x608060405234801561001057600080fd5b5061085a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80634ec8911c11610097578063a5444e0a11610066578063a5444e0a14610274578063b2879c8314610292578063c0e1dcf7146102b0578063f66f77ce146102ce576100f5565b80634ec8911c146101c85780635dc65934146101f857806394c0c1e014610228578063a10cae6014610258576100f5565b806327d82f67116100d357806327d82f67146101525780633815f2a21461017057806338a0fa601461018c578063424e0ed3146101aa576100f5565b80630488f95f146100fa57806305fc12e31461011857806319d39c7a14610134575b600080fd5b6101026102ec565b60405161010f9190610567565b60405180910390f35b610132600480360381019061012d91906105bf565b6102f2565b005b61013c61030f565b6040516101499190610605565b60405180910390f35b61015a610315565b6040516101679190610567565b60405180910390f35b61018a600480360381019061018591906105bf565b61031b565b005b610194610337565b6040516101a19190610605565b60405180910390f35b6101b261033d565b6040516101bf9190610661565b60405180910390f35b6101e260048036038101906101dd9190610700565b610363565b6040516101ef9190610762565b60405180910390f35b610212600480360381019061020d9190610700565b6103cd565b60405161021f91906107dc565b60405180910390f35b610242600480360381019061023d9190610700565b61044a565b60405161024f9190610762565b60405180910390f35b610272600480360381019061026d91906107f7565b6104b2565b005b61027c6104f6565b6040516102899190610605565b60405180910390f35b61029a6104fc565b6040516102a79190610567565b60405180910390f35b6102b8610502565b6040516102c59190610661565b60405180910390f35b6102d6610528565b6040516102e39190610661565b60405180910390f35b60015481565b80600060016101000a81548160ff02191690831515021790555050565b60095481565b60045481565b806000806101000a81548160ff02191690831515021790555050565b60065481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008360018190555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600381905550600060019054906101000a900460ff1690509392505050565b60008360078190555082600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600981905550600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509392505050565b60008360048190555082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160068190555060008054906101000a900460ff1690509392505050565b80600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60035481565b60075481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000819050919050565b6105618161054e565b82525050565b600060208201905061057c6000830184610558565b92915050565b600080fd5b60008115159050919050565b61059c81610587565b81146105a757600080fd5b50565b6000813590506105b981610593565b92915050565b6000602082840312156105d5576105d4610582565b5b60006105e3848285016105aa565b91505092915050565b6000819050919050565b6105ff816105ec565b82525050565b600060208201905061061a60008301846105f6565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061064b82610620565b9050919050565b61065b81610640565b82525050565b60006020820190506106766000830184610652565b92915050565b6106858161054e565b811461069057600080fd5b50565b6000813590506106a28161067c565b92915050565b6106b181610640565b81146106bc57600080fd5b50565b6000813590506106ce816106a8565b92915050565b6106dd816105ec565b81146106e857600080fd5b50565b6000813590506106fa816106d4565b92915050565b60008060006060848603121561071957610718610582565b5b600061072786828701610693565b9350506020610738868287016106bf565b9250506040610749868287016106eb565b9150509250925092565b61075c81610587565b82525050565b60006020820190506107776000830184610753565b92915050565b6000819050919050565b60006107a261079d61079884610620565b61077d565b610620565b9050919050565b60006107b482610787565b9050919050565b60006107c6826107a9565b9050919050565b6107d6816107bb565b82525050565b60006020820190506107f160008301846107cd565b92915050565b60006020828403121561080d5761080c610582565b5b600061081b848285016106bf565b9150509291505056fea26469706673582212202980e98ea2426a5527c4bdce708acf9368f58dc0431ca7059a1d391fd3fbcfff64736f6c634300080d0033"

// DeployPendleData deploys a new Ethereum contract, binding an instance of PendleData to it.
func DeployPendleData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PendleData, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PendleDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PendleData{PendleDataCaller: PendleDataCaller{contract: contract}, PendleDataTransactor: PendleDataTransactor{contract: contract}, PendleDataFilterer: PendleDataFilterer{contract: contract}}, nil
}

// PendleData is an auto generated Go binding around an Ethereum contract.
type PendleData struct {
	PendleDataCaller     // Read-only binding to the contract
	PendleDataTransactor // Write-only binding to the contract
	PendleDataFilterer   // Log filterer for contract events
}

// PendleDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleDataSession struct {
	Contract     *PendleData       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleDataCallerSession struct {
	Contract *PendleDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PendleDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleDataTransactorSession struct {
	Contract     *PendleDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PendleDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleDataRaw struct {
	Contract *PendleData // Generic contract binding to access the raw methods on
}

// PendleDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleDataCallerRaw struct {
	Contract *PendleDataCaller // Generic read-only contract binding to access the raw methods on
}

// PendleDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleDataTransactorRaw struct {
	Contract *PendleDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleData creates a new instance of PendleData, bound to a specific deployed contract.
func NewPendleData(address common.Address, backend bind.ContractBackend) (*PendleData, error) {
	contract, err := bindPendleData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleData{PendleDataCaller: PendleDataCaller{contract: contract}, PendleDataTransactor: PendleDataTransactor{contract: contract}, PendleDataFilterer: PendleDataFilterer{contract: contract}}, nil
}

// NewPendleDataCaller creates a new read-only instance of PendleData, bound to a specific deployed contract.
func NewPendleDataCaller(address common.Address, caller bind.ContractCaller) (*PendleDataCaller, error) {
	contract, err := bindPendleData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleDataCaller{contract: contract}, nil
}

// NewPendleDataTransactor creates a new write-only instance of PendleData, bound to a specific deployed contract.
func NewPendleDataTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleDataTransactor, error) {
	contract, err := bindPendleData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleDataTransactor{contract: contract}, nil
}

// NewPendleDataFilterer creates a new log filterer instance of PendleData, bound to a specific deployed contract.
func NewPendleDataFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleDataFilterer, error) {
	contract, err := bindPendleData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleDataFilterer{contract: contract}, nil
}

// bindPendleData binds a generic wrapper to an already deployed contract.
func bindPendleData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PendleDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleData *PendleDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleData.Contract.PendleDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleData *PendleDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleData.Contract.PendleDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleData *PendleDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleData.Contract.PendleDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleData *PendleDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleData *PendleDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleData *PendleDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleData.Contract.contract.Transact(opts, method, params...)
}

// IsValidOTId is a free data retrieval call binding the contract method 0x0488f95f.
//
// Solidity: function isValidOTId() view returns(bytes32)
func (_PendleData *PendleDataCaller) IsValidOTId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidOTId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidOTId is a free data retrieval call binding the contract method 0x0488f95f.
//
// Solidity: function isValidOTId() view returns(bytes32)
func (_PendleData *PendleDataSession) IsValidOTId() ([32]byte, error) {
	return _PendleData.Contract.IsValidOTId(&_PendleData.CallOpts)
}

// IsValidOTId is a free data retrieval call binding the contract method 0x0488f95f.
//
// Solidity: function isValidOTId() view returns(bytes32)
func (_PendleData *PendleDataCallerSession) IsValidOTId() ([32]byte, error) {
	return _PendleData.Contract.IsValidOTId(&_PendleData.CallOpts)
}

// IsValidOTMaturity is a free data retrieval call binding the contract method 0xa5444e0a.
//
// Solidity: function isValidOTMaturity() view returns(uint256)
func (_PendleData *PendleDataCaller) IsValidOTMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidOTMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsValidOTMaturity is a free data retrieval call binding the contract method 0xa5444e0a.
//
// Solidity: function isValidOTMaturity() view returns(uint256)
func (_PendleData *PendleDataSession) IsValidOTMaturity() (*big.Int, error) {
	return _PendleData.Contract.IsValidOTMaturity(&_PendleData.CallOpts)
}

// IsValidOTMaturity is a free data retrieval call binding the contract method 0xa5444e0a.
//
// Solidity: function isValidOTMaturity() view returns(uint256)
func (_PendleData *PendleDataCallerSession) IsValidOTMaturity() (*big.Int, error) {
	return _PendleData.Contract.IsValidOTMaturity(&_PendleData.CallOpts)
}

// IsValidOTUnderlying is a free data retrieval call binding the contract method 0xc0e1dcf7.
//
// Solidity: function isValidOTUnderlying() view returns(address)
func (_PendleData *PendleDataCaller) IsValidOTUnderlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidOTUnderlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsValidOTUnderlying is a free data retrieval call binding the contract method 0xc0e1dcf7.
//
// Solidity: function isValidOTUnderlying() view returns(address)
func (_PendleData *PendleDataSession) IsValidOTUnderlying() (common.Address, error) {
	return _PendleData.Contract.IsValidOTUnderlying(&_PendleData.CallOpts)
}

// IsValidOTUnderlying is a free data retrieval call binding the contract method 0xc0e1dcf7.
//
// Solidity: function isValidOTUnderlying() view returns(address)
func (_PendleData *PendleDataCallerSession) IsValidOTUnderlying() (common.Address, error) {
	return _PendleData.Contract.IsValidOTUnderlying(&_PendleData.CallOpts)
}

// IsValidXYTId is a free data retrieval call binding the contract method 0x27d82f67.
//
// Solidity: function isValidXYTId() view returns(bytes32)
func (_PendleData *PendleDataCaller) IsValidXYTId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidXYTId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidXYTId is a free data retrieval call binding the contract method 0x27d82f67.
//
// Solidity: function isValidXYTId() view returns(bytes32)
func (_PendleData *PendleDataSession) IsValidXYTId() ([32]byte, error) {
	return _PendleData.Contract.IsValidXYTId(&_PendleData.CallOpts)
}

// IsValidXYTId is a free data retrieval call binding the contract method 0x27d82f67.
//
// Solidity: function isValidXYTId() view returns(bytes32)
func (_PendleData *PendleDataCallerSession) IsValidXYTId() ([32]byte, error) {
	return _PendleData.Contract.IsValidXYTId(&_PendleData.CallOpts)
}

// IsValidXYTMaturity is a free data retrieval call binding the contract method 0x38a0fa60.
//
// Solidity: function isValidXYTMaturity() view returns(uint256)
func (_PendleData *PendleDataCaller) IsValidXYTMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidXYTMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsValidXYTMaturity is a free data retrieval call binding the contract method 0x38a0fa60.
//
// Solidity: function isValidXYTMaturity() view returns(uint256)
func (_PendleData *PendleDataSession) IsValidXYTMaturity() (*big.Int, error) {
	return _PendleData.Contract.IsValidXYTMaturity(&_PendleData.CallOpts)
}

// IsValidXYTMaturity is a free data retrieval call binding the contract method 0x38a0fa60.
//
// Solidity: function isValidXYTMaturity() view returns(uint256)
func (_PendleData *PendleDataCallerSession) IsValidXYTMaturity() (*big.Int, error) {
	return _PendleData.Contract.IsValidXYTMaturity(&_PendleData.CallOpts)
}

// IsValidXYTUnderlying is a free data retrieval call binding the contract method 0xf66f77ce.
//
// Solidity: function isValidXYTUnderlying() view returns(address)
func (_PendleData *PendleDataCaller) IsValidXYTUnderlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "isValidXYTUnderlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsValidXYTUnderlying is a free data retrieval call binding the contract method 0xf66f77ce.
//
// Solidity: function isValidXYTUnderlying() view returns(address)
func (_PendleData *PendleDataSession) IsValidXYTUnderlying() (common.Address, error) {
	return _PendleData.Contract.IsValidXYTUnderlying(&_PendleData.CallOpts)
}

// IsValidXYTUnderlying is a free data retrieval call binding the contract method 0xf66f77ce.
//
// Solidity: function isValidXYTUnderlying() view returns(address)
func (_PendleData *PendleDataCallerSession) IsValidXYTUnderlying() (common.Address, error) {
	return _PendleData.Contract.IsValidXYTUnderlying(&_PendleData.CallOpts)
}

// XytTokensCalledId is a free data retrieval call binding the contract method 0xb2879c83.
//
// Solidity: function xytTokensCalledId() view returns(bytes32)
func (_PendleData *PendleDataCaller) XytTokensCalledId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "xytTokensCalledId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// XytTokensCalledId is a free data retrieval call binding the contract method 0xb2879c83.
//
// Solidity: function xytTokensCalledId() view returns(bytes32)
func (_PendleData *PendleDataSession) XytTokensCalledId() ([32]byte, error) {
	return _PendleData.Contract.XytTokensCalledId(&_PendleData.CallOpts)
}

// XytTokensCalledId is a free data retrieval call binding the contract method 0xb2879c83.
//
// Solidity: function xytTokensCalledId() view returns(bytes32)
func (_PendleData *PendleDataCallerSession) XytTokensCalledId() ([32]byte, error) {
	return _PendleData.Contract.XytTokensCalledId(&_PendleData.CallOpts)
}

// XytTokensCalledMaturity is a free data retrieval call binding the contract method 0x19d39c7a.
//
// Solidity: function xytTokensCalledMaturity() view returns(uint256)
func (_PendleData *PendleDataCaller) XytTokensCalledMaturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "xytTokensCalledMaturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XytTokensCalledMaturity is a free data retrieval call binding the contract method 0x19d39c7a.
//
// Solidity: function xytTokensCalledMaturity() view returns(uint256)
func (_PendleData *PendleDataSession) XytTokensCalledMaturity() (*big.Int, error) {
	return _PendleData.Contract.XytTokensCalledMaturity(&_PendleData.CallOpts)
}

// XytTokensCalledMaturity is a free data retrieval call binding the contract method 0x19d39c7a.
//
// Solidity: function xytTokensCalledMaturity() view returns(uint256)
func (_PendleData *PendleDataCallerSession) XytTokensCalledMaturity() (*big.Int, error) {
	return _PendleData.Contract.XytTokensCalledMaturity(&_PendleData.CallOpts)
}

// XytTokensCalledUnderlying is a free data retrieval call binding the contract method 0x424e0ed3.
//
// Solidity: function xytTokensCalledUnderlying() view returns(address)
func (_PendleData *PendleDataCaller) XytTokensCalledUnderlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleData.contract.Call(opts, &out, "xytTokensCalledUnderlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XytTokensCalledUnderlying is a free data retrieval call binding the contract method 0x424e0ed3.
//
// Solidity: function xytTokensCalledUnderlying() view returns(address)
func (_PendleData *PendleDataSession) XytTokensCalledUnderlying() (common.Address, error) {
	return _PendleData.Contract.XytTokensCalledUnderlying(&_PendleData.CallOpts)
}

// XytTokensCalledUnderlying is a free data retrieval call binding the contract method 0x424e0ed3.
//
// Solidity: function xytTokensCalledUnderlying() view returns(address)
func (_PendleData *PendleDataCallerSession) XytTokensCalledUnderlying() (common.Address, error) {
	return _PendleData.Contract.XytTokensCalledUnderlying(&_PendleData.CallOpts)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactor) IsValidOT(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidOT", f, u, m)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataSession) IsValidOT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidOT is a paid mutator transaction binding the contract method 0x4ec8911c.
//
// Solidity: function isValidOT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactorSession) IsValidOT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataTransactor) IsValidOTReturns(opts *bind.TransactOpts, o bool) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidOTReturns", o)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataSession) IsValidOTReturns(o bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOTReturns(&_PendleData.TransactOpts, o)
}

// IsValidOTReturns is a paid mutator transaction binding the contract method 0x05fc12e3.
//
// Solidity: function isValidOTReturns(bool o) returns()
func (_PendleData *PendleDataTransactorSession) IsValidOTReturns(o bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidOTReturns(&_PendleData.TransactOpts, o)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactor) IsValidXYT(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidXYT", f, u, m)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataSession) IsValidXYT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidXYT is a paid mutator transaction binding the contract method 0x94c0c1e0.
//
// Solidity: function isValidXYT(bytes32 f, address u, uint256 m) returns(bool)
func (_PendleData *PendleDataTransactorSession) IsValidXYT(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYT(&_PendleData.TransactOpts, f, u, m)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataTransactor) IsValidXYTReturns(opts *bind.TransactOpts, x bool) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "isValidXYTReturns", x)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataSession) IsValidXYTReturns(x bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYTReturns(&_PendleData.TransactOpts, x)
}

// IsValidXYTReturns is a paid mutator transaction binding the contract method 0x3815f2a2.
//
// Solidity: function isValidXYTReturns(bool x) returns()
func (_PendleData *PendleDataTransactorSession) IsValidXYTReturns(x bool) (*types.Transaction, error) {
	return _PendleData.Contract.IsValidXYTReturns(&_PendleData.TransactOpts, x)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataTransactor) XytTokens(opts *bind.TransactOpts, f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "xytTokens", f, u, m)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataSession) XytTokens(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokens(&_PendleData.TransactOpts, f, u, m)
}

// XytTokens is a paid mutator transaction binding the contract method 0x5dc65934.
//
// Solidity: function xytTokens(bytes32 f, address u, uint256 m) returns(address)
func (_PendleData *PendleDataTransactorSession) XytTokens(f [32]byte, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokens(&_PendleData.TransactOpts, f, u, m)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataTransactor) XytTokensReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _PendleData.contract.Transact(opts, "xytTokensReturns", a)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataSession) XytTokensReturns(a common.Address) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokensReturns(&_PendleData.TransactOpts, a)
}

// XytTokensReturns is a paid mutator transaction binding the contract method 0xa10cae60.
//
// Solidity: function xytTokensReturns(address a) returns()
func (_PendleData *PendleDataTransactorSession) XytTokensReturns(a common.Address) (*types.Transaction, error) {
	return _PendleData.Contract.XytTokensReturns(&_PendleData.TransactOpts, a)
}
