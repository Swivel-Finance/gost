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

// TempusABI is the input ABI used to generate the binding from.
const TempusABI = "[{\"inputs\":[{\"internalType\":\"contractAny\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"depositAndFix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositAndFixCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"amm\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minimumReturned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"depositAndFixReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityTimeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"redeemToBacking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemToBackingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldBearingToken\",\"outputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"yieldBearingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TempusBin is the compiled bytecode used for deploying new contracts.
var TempusBin = "0x608060405234801561001057600080fd5b506109b0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635d61761f116100665780635d61761f146101385780636c8d4fa1146101545780639683982814610170578063d08dd24b1461018c578063e77638cb146101a857610093565b806308123d86146100985780632d81f838146100cc5780634e8bfdaa146100ea5780634f255a1f14610108575b600080fd5b6100b260048036038101906100ad9190610580565b6101da565b6040516100c3959493929190610656565b60405180910390f35b6100d461025d565b6040516100e191906106ca565b60405180910390f35b6100f2610287565b6040516100ff91906106e5565b60405180910390f35b610122600480360381019061011d919061077c565b610291565b60405161012f91906106e5565b60405180910390f35b610152600480360381019061014d9190610580565b6103cf565b005b61016e60048036038101906101699190610835565b6103d9565b005b61018a60048036038101906101859190610580565b6104ad565b005b6101a660048036038101906101a191906108da565b6104b7565b005b6101c260048036038101906101bd9190610907565b6104fb565b6040516101d193929190610943565b60405180910390f35b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060020154908060030154905085565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b60006040518060a001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018515158152602001848152602001838152506004600087815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548160ff021916908315150217905550606082015181600201556080820151816003015590505060005490509695505050505050565b8060018190555050565b60405180606001604052808381526020018481526020018273ffffffffffffffffffffffffffffffffffffffff16815250600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505050505050565b8060008190555050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60036020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b600080fd5b6000819050919050565b61055d8161054a565b811461056857600080fd5b50565b60008135905061057a81610554565b92915050565b60006020828403121561059657610595610545565b5b60006105a48482850161056b565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006105f26105ed6105e8846105ad565b6105cd565b6105ad565b9050919050565b6000610604826105d7565b9050919050565b6000610616826105f9565b9050919050565b6106268161060b565b82525050565b60008115159050919050565b6106418161062c565b82525050565b6106508161054a565b82525050565b600060a08201905061066b600083018861061d565b610678602083018761061d565b6106856040830186610638565b6106926060830185610647565b61069f6080830184610647565b9695505050505050565b60006106b4826105f9565b9050919050565b6106c4816106a9565b82525050565b60006020820190506106df60008301846106bb565b92915050565b60006020820190506106fa6000830184610647565b92915050565b600061070b826105ad565b9050919050565b600061071d82610700565b9050919050565b61072d81610712565b811461073857600080fd5b50565b60008135905061074a81610724565b92915050565b6107598161062c565b811461076457600080fd5b50565b60008135905061077681610750565b92915050565b60008060008060008060c0878903121561079957610798610545565b5b60006107a789828a0161073b565b96505060206107b889828a0161073b565b95505060406107c989828a0161056b565b94505060606107da89828a01610767565b93505060806107eb89828a0161056b565b92505060a06107fc89828a0161056b565b9150509295509295509295565b61081281610700565b811461081d57600080fd5b50565b60008135905061082f81610809565b92915050565b6000806000806080858703121561084f5761084e610545565b5b600061085d87828801610820565b945050602061086e8782880161056b565b935050604061087f8782880161056b565b925050606061089087828801610820565b91505092959194509250565b60006108a782610700565b9050919050565b6108b78161089c565b81146108c257600080fd5b50565b6000813590506108d4816108ae565b92915050565b6000602082840312156108f0576108ef610545565b5b60006108fe848285016108c5565b91505092915050565b60006020828403121561091d5761091c610545565b5b600061092b84828501610820565b91505092915050565b61093d81610700565b82525050565b60006060820190506109586000830186610647565b6109656020830185610647565b6109726040830184610934565b94935050505056fea264697066735822122066d96f647973fe1de654661e31d5594becc4559508cc6b1db9b931d57eda920064736f6c634300080d0033"

// DeployTempus deploys a new Ethereum contract, binding an instance of Tempus to it.
func DeployTempus(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tempus, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TempusBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tempus{TempusCaller: TempusCaller{contract: contract}, TempusTransactor: TempusTransactor{contract: contract}, TempusFilterer: TempusFilterer{contract: contract}}, nil
}

// Tempus is an auto generated Go binding around an Ethereum contract.
type Tempus struct {
	TempusCaller     // Read-only binding to the contract
	TempusTransactor // Write-only binding to the contract
	TempusFilterer   // Log filterer for contract events
}

// TempusCaller is an auto generated read-only Go binding around an Ethereum contract.
type TempusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TempusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TempusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TempusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TempusSession struct {
	Contract     *Tempus           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TempusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TempusCallerSession struct {
	Contract *TempusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TempusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TempusTransactorSession struct {
	Contract     *TempusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TempusRaw is an auto generated low-level Go binding around an Ethereum contract.
type TempusRaw struct {
	Contract *Tempus // Generic contract binding to access the raw methods on
}

// TempusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TempusCallerRaw struct {
	Contract *TempusCaller // Generic read-only contract binding to access the raw methods on
}

// TempusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TempusTransactorRaw struct {
	Contract *TempusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTempus creates a new instance of Tempus, bound to a specific deployed contract.
func NewTempus(address common.Address, backend bind.ContractBackend) (*Tempus, error) {
	contract, err := bindTempus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tempus{TempusCaller: TempusCaller{contract: contract}, TempusTransactor: TempusTransactor{contract: contract}, TempusFilterer: TempusFilterer{contract: contract}}, nil
}

// NewTempusCaller creates a new read-only instance of Tempus, bound to a specific deployed contract.
func NewTempusCaller(address common.Address, caller bind.ContractCaller) (*TempusCaller, error) {
	contract, err := bindTempus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TempusCaller{contract: contract}, nil
}

// NewTempusTransactor creates a new write-only instance of Tempus, bound to a specific deployed contract.
func NewTempusTransactor(address common.Address, transactor bind.ContractTransactor) (*TempusTransactor, error) {
	contract, err := bindTempus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TempusTransactor{contract: contract}, nil
}

// NewTempusFilterer creates a new log filterer instance of Tempus, bound to a specific deployed contract.
func NewTempusFilterer(address common.Address, filterer bind.ContractFilterer) (*TempusFilterer, error) {
	contract, err := bindTempus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TempusFilterer{contract: contract}, nil
}

// bindTempus binds a generic wrapper to an already deployed contract.
func bindTempus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TempusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tempus *TempusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tempus.Contract.TempusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tempus *TempusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tempus.Contract.TempusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tempus *TempusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tempus.Contract.TempusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tempus *TempusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tempus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tempus *TempusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tempus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tempus *TempusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tempus.Contract.contract.Transact(opts, method, params...)
}

// DepositAndFixCalled is a free data retrieval call binding the contract method 0x08123d86.
//
// Solidity: function depositAndFixCalled(uint256 ) view returns(address amm, address pool, bool bt, uint256 minimumReturned, uint256 deadline)
func (_Tempus *TempusCaller) DepositAndFixCalled(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amm             common.Address
	Pool            common.Address
	Bt              bool
	MinimumReturned *big.Int
	Deadline        *big.Int
}, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "depositAndFixCalled", arg0)

	outstruct := new(struct {
		Amm             common.Address
		Pool            common.Address
		Bt              bool
		MinimumReturned *big.Int
		Deadline        *big.Int
	})

	outstruct.Amm = out[0].(common.Address)
	outstruct.Pool = out[1].(common.Address)
	outstruct.Bt = out[2].(bool)
	outstruct.MinimumReturned = out[3].(*big.Int)
	outstruct.Deadline = out[4].(*big.Int)

	return *outstruct, err

}

// DepositAndFixCalled is a free data retrieval call binding the contract method 0x08123d86.
//
// Solidity: function depositAndFixCalled(uint256 ) view returns(address amm, address pool, bool bt, uint256 minimumReturned, uint256 deadline)
func (_Tempus *TempusSession) DepositAndFixCalled(arg0 *big.Int) (struct {
	Amm             common.Address
	Pool            common.Address
	Bt              bool
	MinimumReturned *big.Int
	Deadline        *big.Int
}, error) {
	return _Tempus.Contract.DepositAndFixCalled(&_Tempus.CallOpts, arg0)
}

// DepositAndFixCalled is a free data retrieval call binding the contract method 0x08123d86.
//
// Solidity: function depositAndFixCalled(uint256 ) view returns(address amm, address pool, bool bt, uint256 minimumReturned, uint256 deadline)
func (_Tempus *TempusCallerSession) DepositAndFixCalled(arg0 *big.Int) (struct {
	Amm             common.Address
	Pool            common.Address
	Bt              bool
	MinimumReturned *big.Int
	Deadline        *big.Int
}, error) {
	return _Tempus.Contract.DepositAndFixCalled(&_Tempus.CallOpts, arg0)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusSession) MaturityTime() (*big.Int, error) {
	return _Tempus.Contract.MaturityTime(&_Tempus.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_Tempus *TempusCallerSession) MaturityTime() (*big.Int, error) {
	return _Tempus.Contract.MaturityTime(&_Tempus.CallOpts)
}

// RedeemToBackingCalled is a free data retrieval call binding the contract method 0xe77638cb.
//
// Solidity: function redeemToBackingCalled(address ) view returns(uint256 amount, uint256 maturity, address underlying)
func (_Tempus *TempusCaller) RedeemToBackingCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	Maturity   *big.Int
	Underlying common.Address
}, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "redeemToBackingCalled", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		Maturity   *big.Int
		Underlying common.Address
	})

	outstruct.Amount = out[0].(*big.Int)
	outstruct.Maturity = out[1].(*big.Int)
	outstruct.Underlying = out[2].(common.Address)

	return *outstruct, err

}

// RedeemToBackingCalled is a free data retrieval call binding the contract method 0xe77638cb.
//
// Solidity: function redeemToBackingCalled(address ) view returns(uint256 amount, uint256 maturity, address underlying)
func (_Tempus *TempusSession) RedeemToBackingCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	Maturity   *big.Int
	Underlying common.Address
}, error) {
	return _Tempus.Contract.RedeemToBackingCalled(&_Tempus.CallOpts, arg0)
}

// RedeemToBackingCalled is a free data retrieval call binding the contract method 0xe77638cb.
//
// Solidity: function redeemToBackingCalled(address ) view returns(uint256 amount, uint256 maturity, address underlying)
func (_Tempus *TempusCallerSession) RedeemToBackingCalled(arg0 common.Address) (struct {
	Amount     *big.Int
	Maturity   *big.Int
	Underlying common.Address
}, error) {
	return _Tempus.Contract.RedeemToBackingCalled(&_Tempus.CallOpts, arg0)
}

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusCaller) YieldBearingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "yieldBearingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusSession) YieldBearingToken() (common.Address, error) {
	return _Tempus.Contract.YieldBearingToken(&_Tempus.CallOpts)
}

// YieldBearingToken is a free data retrieval call binding the contract method 0x2d81f838.
//
// Solidity: function yieldBearingToken() view returns(address)
func (_Tempus *TempusCallerSession) YieldBearingToken() (common.Address, error) {
	return _Tempus.Contract.YieldBearingToken(&_Tempus.CallOpts)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusTransactor) DepositAndFix(opts *bind.TransactOpts, x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "depositAndFix", x, p, a, bt, mr, d)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusSession) DepositAndFix(x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFix(&_Tempus.TransactOpts, x, p, a, bt, mr, d)
}

// DepositAndFix is a paid mutator transaction binding the contract method 0x4f255a1f.
//
// Solidity: function depositAndFix(address x, address p, uint256 a, bool bt, uint256 mr, uint256 d) returns(uint256)
func (_Tempus *TempusTransactorSession) DepositAndFix(x common.Address, p common.Address, a *big.Int, bt bool, mr *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFix(&_Tempus.TransactOpts, x, p, a, bt, mr, d)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusTransactor) DepositAndFixReturns(opts *bind.TransactOpts, r *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "depositAndFixReturns", r)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusSession) DepositAndFixReturns(r *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFixReturns(&_Tempus.TransactOpts, r)
}

// DepositAndFixReturns is a paid mutator transaction binding the contract method 0x96839828.
//
// Solidity: function depositAndFixReturns(uint256 r) returns()
func (_Tempus *TempusTransactorSession) DepositAndFixReturns(r *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.DepositAndFixReturns(&_Tempus.TransactOpts, r)
}

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusTransactor) MaturityTimeReturns(opts *bind.TransactOpts, m *big.Int) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "maturityTimeReturns", m)
}

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusSession) MaturityTimeReturns(m *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.MaturityTimeReturns(&_Tempus.TransactOpts, m)
}

// MaturityTimeReturns is a paid mutator transaction binding the contract method 0x5d61761f.
//
// Solidity: function maturityTimeReturns(uint256 m) returns()
func (_Tempus *TempusTransactorSession) MaturityTimeReturns(m *big.Int) (*types.Transaction, error) {
	return _Tempus.Contract.MaturityTimeReturns(&_Tempus.TransactOpts, m)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 m, uint256 a, address u) returns()
func (_Tempus *TempusTransactor) RedeemToBacking(opts *bind.TransactOpts, o common.Address, m *big.Int, a *big.Int, u common.Address) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "redeemToBacking", o, m, a, u)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 m, uint256 a, address u) returns()
func (_Tempus *TempusSession) RedeemToBacking(o common.Address, m *big.Int, a *big.Int, u common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.RedeemToBacking(&_Tempus.TransactOpts, o, m, a, u)
}

// RedeemToBacking is a paid mutator transaction binding the contract method 0x6c8d4fa1.
//
// Solidity: function redeemToBacking(address o, uint256 m, uint256 a, address u) returns()
func (_Tempus *TempusTransactorSession) RedeemToBacking(o common.Address, m *big.Int, a *big.Int, u common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.RedeemToBacking(&_Tempus.TransactOpts, o, m, a, u)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusTransactor) YieldBearingTokenReturns(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Tempus.contract.Transact(opts, "yieldBearingTokenReturns", t)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusSession) YieldBearingTokenReturns(t common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.YieldBearingTokenReturns(&_Tempus.TransactOpts, t)
}

// YieldBearingTokenReturns is a paid mutator transaction binding the contract method 0xd08dd24b.
//
// Solidity: function yieldBearingTokenReturns(address t) returns()
func (_Tempus *TempusTransactorSession) YieldBearingTokenReturns(t common.Address) (*types.Transaction, error) {
	return _Tempus.Contract.YieldBearingTokenReturns(&_Tempus.TransactOpts, t)
}
