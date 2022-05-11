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
const TempusABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"depositAndFix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"depositAndFixReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBackingTokenCalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityTimeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumReturnCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"redeemToBacking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAMMCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusPoolCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldBearingToken\",\"outputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"yieldBearingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TempusBin is the compiled bytecode used for deploying new contracts.
var TempusBin = "0x608060405234801561001057600080fd5b5061099b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80635d61761f11610097578063a2f428f711610066578063a2f428f714610279578063c203bd6b14610297578063d08dd24b146102b5578063f189b337146102d157610100565b80635d61761f146102075780636c8d4fa1146102235780637822212d1461023f578063968398281461025d57610100565b80634cf1cca7116100d35780634cf1cca71461017d5780634e8bfdaa1461019b5780634f255a1f146101b957806357e5b682146101e957610100565b806302d103c014610105578063229204d9146101235780632d81f838146101415780633ef943bb1461015f575b600080fd5b61010d6102ef565b60405161011a91906105b1565b60405180910390f35b61012b610302565b60405161013891906105e5565b60405180910390f35b610149610308565b604051610156919061067f565b60405180910390f35b610167610332565b60405161017491906105e5565b60405180910390f35b610185610338565b60405161019291906105e5565b60405180910390f35b6101a361033e565b6040516101b091906105e5565b60405180910390f35b6101d360048036038101906101ce9190610747565b610348565b6040516101e091906105e5565b60405180910390f35b6101f161040a565b6040516101fe91906105e5565b60405180910390f35b610221600480360381019061021c91906107d4565b610410565b005b61023d6004803603810190610238919061082d565b61041a565b005b6102476104b0565b60405161025491906108b5565b60405180910390f35b610277600480360381019061027291906107d4565b6104d6565b005b6102816104e0565b60405161028e91906108df565b60405180910390f35b61029f610506565b6040516102ac91906108b5565b60405180910390f35b6102cf60048036038101906102ca9190610938565b61052c565b005b6102d9610570565b6040516102e691906108df565b60405180910390f35b600660009054906101000a900460ff1681565b60075481565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055481565b60085481565b6000600154905090565b600086600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460058190555083600660006101000a81548160ff021916908315150217905550826007819055508160088190555060005490509695505050505050565b600a5481565b8060018190555050565b83600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600a819055508160058190555080600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060008190555050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008115159050919050565b6105ab81610596565b82525050565b60006020820190506105c660008301846105a2565b92915050565b6000819050919050565b6105df816105cc565b82525050565b60006020820190506105fa60008301846105d6565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061064561064061063b84610600565b610620565b610600565b9050919050565b60006106578261062a565b9050919050565b60006106698261064c565b9050919050565b6106798161065e565b82525050565b60006020820190506106946000830184610670565b92915050565b600080fd5b60006106aa82610600565b9050919050565b60006106bc8261069f565b9050919050565b6106cc816106b1565b81146106d757600080fd5b50565b6000813590506106e9816106c3565b92915050565b6106f8816105cc565b811461070357600080fd5b50565b600081359050610715816106ef565b92915050565b61072481610596565b811461072f57600080fd5b50565b6000813590506107418161071b565b92915050565b60008060008060008060c087890312156107645761076361069a565b5b600061077289828a016106da565b965050602061078389828a016106da565b955050604061079489828a01610706565b94505060606107a589828a01610732565b93505060806107b689828a01610706565b92505060a06107c789828a01610706565b9150509295509295509295565b6000602082840312156107ea576107e961069a565b5b60006107f884828501610706565b91505092915050565b61080a8161069f565b811461081557600080fd5b50565b60008135905061082781610801565b92915050565b600080600080608085870312156108475761084661069a565b5b600061085587828801610818565b945050602061086687828801610706565b935050604061087787828801610706565b925050606061088887828801610818565b91505092959194509250565b600061089f8261064c565b9050919050565b6108af81610894565b82525050565b60006020820190506108ca60008301846108a6565b92915050565b6108d98161069f565b82525050565b60006020820190506108f460008301846108d0565b92915050565b60006109058261069f565b9050919050565b610915816108fa565b811461092057600080fd5b50565b6000813590506109328161090c565b92915050565b60006020828403121561094e5761094d61069a565b5b600061095c84828501610923565b9150509291505056fea2646970667358221220afaf1d0d9ca199670e7be997c798180189a39f309fbe421357d5102016a162b464736f6c634300080d0033"

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

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusCaller) AmountCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "amountCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusSession) AmountCalled() (*big.Int, error) {
	return _Tempus.Contract.AmountCalled(&_Tempus.CallOpts)
}

// AmountCalled is a free data retrieval call binding the contract method 0x3ef943bb.
//
// Solidity: function amountCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) AmountCalled() (*big.Int, error) {
	return _Tempus.Contract.AmountCalled(&_Tempus.CallOpts)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusCaller) DeadlineCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "deadlineCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusSession) DeadlineCalled() (*big.Int, error) {
	return _Tempus.Contract.DeadlineCalled(&_Tempus.CallOpts)
}

// DeadlineCalled is a free data retrieval call binding the contract method 0x4cf1cca7.
//
// Solidity: function deadlineCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) DeadlineCalled() (*big.Int, error) {
	return _Tempus.Contract.DeadlineCalled(&_Tempus.CallOpts)
}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusCaller) IsBackingTokenCalled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "isBackingTokenCalled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusSession) IsBackingTokenCalled() (bool, error) {
	return _Tempus.Contract.IsBackingTokenCalled(&_Tempus.CallOpts)
}

// IsBackingTokenCalled is a free data retrieval call binding the contract method 0x02d103c0.
//
// Solidity: function isBackingTokenCalled() view returns(bool)
func (_Tempus *TempusCallerSession) IsBackingTokenCalled() (bool, error) {
	return _Tempus.Contract.IsBackingTokenCalled(&_Tempus.CallOpts)
}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Tempus *TempusCaller) MaturityCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "maturityCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Tempus *TempusSession) MaturityCalled() (*big.Int, error) {
	return _Tempus.Contract.MaturityCalled(&_Tempus.CallOpts)
}

// MaturityCalled is a free data retrieval call binding the contract method 0x57e5b682.
//
// Solidity: function maturityCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) MaturityCalled() (*big.Int, error) {
	return _Tempus.Contract.MaturityCalled(&_Tempus.CallOpts)
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

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusCaller) MinimumReturnCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "minimumReturnCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusSession) MinimumReturnCalled() (*big.Int, error) {
	return _Tempus.Contract.MinimumReturnCalled(&_Tempus.CallOpts)
}

// MinimumReturnCalled is a free data retrieval call binding the contract method 0x229204d9.
//
// Solidity: function minimumReturnCalled() view returns(uint256)
func (_Tempus *TempusCallerSession) MinimumReturnCalled() (*big.Int, error) {
	return _Tempus.Contract.MinimumReturnCalled(&_Tempus.CallOpts)
}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Tempus *TempusCaller) OwnerCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "ownerCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Tempus *TempusSession) OwnerCalled() (common.Address, error) {
	return _Tempus.Contract.OwnerCalled(&_Tempus.CallOpts)
}

// OwnerCalled is a free data retrieval call binding the contract method 0xa2f428f7.
//
// Solidity: function ownerCalled() view returns(address)
func (_Tempus *TempusCallerSession) OwnerCalled() (common.Address, error) {
	return _Tempus.Contract.OwnerCalled(&_Tempus.CallOpts)
}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusCaller) TempusAMMCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "tempusAMMCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusSession) TempusAMMCalled() (common.Address, error) {
	return _Tempus.Contract.TempusAMMCalled(&_Tempus.CallOpts)
}

// TempusAMMCalled is a free data retrieval call binding the contract method 0xc203bd6b.
//
// Solidity: function tempusAMMCalled() view returns(address)
func (_Tempus *TempusCallerSession) TempusAMMCalled() (common.Address, error) {
	return _Tempus.Contract.TempusAMMCalled(&_Tempus.CallOpts)
}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusCaller) TempusPoolCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "tempusPoolCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusSession) TempusPoolCalled() (common.Address, error) {
	return _Tempus.Contract.TempusPoolCalled(&_Tempus.CallOpts)
}

// TempusPoolCalled is a free data retrieval call binding the contract method 0x7822212d.
//
// Solidity: function tempusPoolCalled() view returns(address)
func (_Tempus *TempusCallerSession) TempusPoolCalled() (common.Address, error) {
	return _Tempus.Contract.TempusPoolCalled(&_Tempus.CallOpts)
}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Tempus *TempusCaller) UnderlyingCalled(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tempus.contract.Call(opts, &out, "underlyingCalled")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Tempus *TempusSession) UnderlyingCalled() (common.Address, error) {
	return _Tempus.Contract.UnderlyingCalled(&_Tempus.CallOpts)
}

// UnderlyingCalled is a free data retrieval call binding the contract method 0xf189b337.
//
// Solidity: function underlyingCalled() view returns(address)
func (_Tempus *TempusCallerSession) UnderlyingCalled() (common.Address, error) {
	return _Tempus.Contract.UnderlyingCalled(&_Tempus.CallOpts)
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
