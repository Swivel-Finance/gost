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
const TempusABI = "[{\"inputs\":[],\"name\":\"amountCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlineCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"contractAny\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"bt\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"depositAndFix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"depositAndFixReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBackingTokenCalled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"maturityTimeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumReturnCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"redeemToBacking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemToBackingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAMMCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusPoolCalled\",\"outputs\":[{\"internalType\":\"contractAny\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldBearingToken\",\"outputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIErc20Metadata\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"yieldBearingTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TempusBin is the compiled bytecode used for deploying new contracts.
var TempusBin = "0x608060405234801561001057600080fd5b506109dc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80635d61761f1161008c57806396839828116100665780639683982814610229578063c203bd6b14610245578063d08dd24b14610263578063e77638cb1461027f576100ea565b80635d61761f146101d35780636c8d4fa1146101ef5780637822212d1461020b576100ea565b80633ef943bb116100c85780633ef943bb146101495780634cf1cca7146101675780634e8bfdaa146101855780634f255a1f146101a3576100ea565b806302d103c0146100ef578063229204d91461010d5780632d81f8381461012b575b600080fd5b6100f76102b1565b60405161010491906105a9565b60405180910390f35b6101156102c4565b60405161012291906105dd565b60405180910390f35b6101336102ca565b6040516101409190610677565b60405180910390f35b6101516102f4565b60405161015e91906105dd565b60405180910390f35b61016f6102fa565b60405161017c91906105dd565b60405180910390f35b61018d610300565b60405161019a91906105dd565b60405180910390f35b6101bd60048036038101906101b8919061073f565b61030a565b6040516101ca91906105dd565b60405180910390f35b6101ed60048036038101906101e891906107cc565b6103cc565b005b61020960048036038101906102049190610825565b6103d6565b005b6102136104aa565b60405161022091906108ad565b60405180910390f35b610243600480360381019061023e91906107cc565b6104d0565b005b61024d6104da565b60405161025a91906108ad565b60405180910390f35b61027d60048036038101906102789190610906565b610500565b005b61029960048036038101906102949190610933565b610544565b6040516102a89392919061096f565b60405180910390f35b600660009054906101000a900460ff1681565b60075481565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60055481565b60085481565b6000600154905090565b600086600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508460058190555083600660006101000a81548160ff021916908315150217905550826007819055508160088190555060005490509695505050505050565b8060018190555050565b60405180606001604052808381526020018481526020018273ffffffffffffffffffffffffffffffffffffffff16815250600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060008190555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60096020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60008115159050919050565b6105a38161058e565b82525050565b60006020820190506105be600083018461059a565b92915050565b6000819050919050565b6105d7816105c4565b82525050565b60006020820190506105f260008301846105ce565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061063d610638610633846105f8565b610618565b6105f8565b9050919050565b600061064f82610622565b9050919050565b600061066182610644565b9050919050565b61067181610656565b82525050565b600060208201905061068c6000830184610668565b92915050565b600080fd5b60006106a2826105f8565b9050919050565b60006106b482610697565b9050919050565b6106c4816106a9565b81146106cf57600080fd5b50565b6000813590506106e1816106bb565b92915050565b6106f0816105c4565b81146106fb57600080fd5b50565b60008135905061070d816106e7565b92915050565b61071c8161058e565b811461072757600080fd5b50565b60008135905061073981610713565b92915050565b60008060008060008060c0878903121561075c5761075b610692565b5b600061076a89828a016106d2565b965050602061077b89828a016106d2565b955050604061078c89828a016106fe565b945050606061079d89828a0161072a565b93505060806107ae89828a016106fe565b92505060a06107bf89828a016106fe565b9150509295509295509295565b6000602082840312156107e2576107e1610692565b5b60006107f0848285016106fe565b91505092915050565b61080281610697565b811461080d57600080fd5b50565b60008135905061081f816107f9565b92915050565b6000806000806080858703121561083f5761083e610692565b5b600061084d87828801610810565b945050602061085e878288016106fe565b935050604061086f878288016106fe565b925050606061088087828801610810565b91505092959194509250565b600061089782610644565b9050919050565b6108a78161088c565b82525050565b60006020820190506108c2600083018461089e565b92915050565b60006108d382610697565b9050919050565b6108e3816108c8565b81146108ee57600080fd5b50565b600081359050610900816108da565b92915050565b60006020828403121561091c5761091b610692565b5b600061092a848285016108f1565b91505092915050565b60006020828403121561094957610948610692565b5b600061095784828501610810565b91505092915050565b61096981610697565b82525050565b600060608201905061098460008301866105ce565b61099160208301856105ce565b61099e6040830184610960565b94935050505056fea26469706673582212207a0ee4279f85f111738e5272093212ab03f4bd35a575a406eb6f68d07f62884364736f6c634300080d0033"

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
