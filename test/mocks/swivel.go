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

// Components is an auto generated low-level Go binding around an user-defined struct.
type Components struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Key        [32]byte
	Maker      common.Address
	Underlying common.Address
	Vault      bool
	Exit       bool
	Principal  *big.Int
	Premium    *big.Int
	Maturity   *big.Int
	Expiry     *big.Int
}

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structOrder[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structComponents[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"initiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initiateCalledSignature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"i\",\"type\":\"bool\"}],\"name\":\"initiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"redeemZcTokenCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"underlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x608060405234801561001057600080fd5b506108f0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063154e0f2e1461006757806338599fc8146100835780634578b3ce146100b3578063c268dc2a146100cf578063d2144f58146100ff578063ddcfbbda1461012f575b600080fd5b610081600480360381019061007c919061045c565b610160565b005b61009d600480360381019061009891906104af565b6101d0565b6040516100aa91906104f8565b60405180910390f35b6100cd60048036038101906100c8919061054b565b6101f0565b005b6100e960048036038101906100e491906104af565b61020c565b6040516100f69190610587565b60405180910390f35b610119600480360381019061011491906106b4565b610224565b6040516101269190610777565b60405180910390f35b610149600480360381019061014491906104af565b61039a565b604051610157929190610792565b60405180910390f35b604051806040016040528083815260200182815250600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050505050565b60026020528060005260406000206000915054906101000a900460ff1681565b806000806101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090505481565b600080600090505b8787905081101561037f5785858281811061024a576102496107bb565b5b90506020020135600160008a8a85818110610268576102676107bb565b5b90506101200201602001602081019061028191906104af565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508383828181106102d2576102d16107bb565b5b90506060020160000160208101906102ea9190610816565b600260008a8a85818110610301576103006107bb565b5b90506101200201602001602081019061031a91906104af565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff160217905550808061037790610872565b91505061022c565b5060008054906101000a900460ff1690509695505050505050565b60036020528060005260406000206000915090508060000154908060010154905082565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103f3826103c8565b9050919050565b610403816103e8565b811461040e57600080fd5b50565b600081359050610420816103fa565b92915050565b6000819050919050565b61043981610426565b811461044457600080fd5b50565b60008135905061045681610430565b92915050565b600080600060608486031215610475576104746103be565b5b600061048386828701610411565b935050602061049486828701610447565b92505060406104a586828701610447565b9150509250925092565b6000602082840312156104c5576104c46103be565b5b60006104d384828501610411565b91505092915050565b600060ff82169050919050565b6104f2816104dc565b82525050565b600060208201905061050d60008301846104e9565b92915050565b60008115159050919050565b61052881610513565b811461053357600080fd5b50565b6000813590506105458161051f565b92915050565b600060208284031215610561576105606103be565b5b600061056f84828501610536565b91505092915050565b61058181610426565b82525050565b600060208201905061059c6000830184610578565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126105c7576105c66105a2565b5b8235905067ffffffffffffffff8111156105e4576105e36105a7565b5b60208301915083610120820283011115610601576106006105ac565b5b9250929050565b60008083601f84011261061e5761061d6105a2565b5b8235905067ffffffffffffffff81111561063b5761063a6105a7565b5b602083019150836020820283011115610657576106566105ac565b5b9250929050565b60008083601f840112610674576106736105a2565b5b8235905067ffffffffffffffff811115610691576106906105a7565b5b6020830191508360608202830111156106ad576106ac6105ac565b5b9250929050565b600080600080600080606087890312156106d1576106d06103be565b5b600087013567ffffffffffffffff8111156106ef576106ee6103c3565b5b6106fb89828a016105b1565b9650965050602087013567ffffffffffffffff81111561071e5761071d6103c3565b5b61072a89828a01610608565b9450945050604087013567ffffffffffffffff81111561074d5761074c6103c3565b5b61075989828a0161065e565b92509250509295509295509295565b61077181610513565b82525050565b600060208201905061078c6000830184610768565b92915050565b60006040820190506107a76000830185610578565b6107b46020830184610578565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6107f3816104dc565b81146107fe57600080fd5b50565b600081359050610810816107ea565b92915050565b60006020828403121561082c5761082b6103be565b5b600061083a84828501610801565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061087d82610426565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108af576108ae610843565b5b60018201905091905056fea264697066735822122091d076a6dd5f20c2c4308e97ea684c63aaa23946a80e04ad3a807a822a77200e64736f6c634300080d0033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// Swivel is an auto generated Go binding around an Ethereum contract.
type Swivel struct {
	SwivelCaller     // Read-only binding to the contract
	SwivelTransactor // Write-only binding to the contract
	SwivelFilterer   // Log filterer for contract events
}

// SwivelCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwivelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwivelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwivelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwivelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwivelSession struct {
	Contract     *Swivel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwivelCallerSession struct {
	Contract *SwivelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwivelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwivelTransactorSession struct {
	Contract     *SwivelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwivelRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwivelRaw struct {
	Contract *Swivel // Generic contract binding to access the raw methods on
}

// SwivelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwivelCallerRaw struct {
	Contract *SwivelCaller // Generic read-only contract binding to access the raw methods on
}

// SwivelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwivelTransactorRaw struct {
	Contract *SwivelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwivel creates a new instance of Swivel, bound to a specific deployed contract.
func NewSwivel(address common.Address, backend bind.ContractBackend) (*Swivel, error) {
	contract, err := bindSwivel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swivel{SwivelCaller: SwivelCaller{contract: contract}, SwivelTransactor: SwivelTransactor{contract: contract}, SwivelFilterer: SwivelFilterer{contract: contract}}, nil
}

// NewSwivelCaller creates a new read-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelCaller(address common.Address, caller bind.ContractCaller) (*SwivelCaller, error) {
	contract, err := bindSwivel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelCaller{contract: contract}, nil
}

// NewSwivelTransactor creates a new write-only instance of Swivel, bound to a specific deployed contract.
func NewSwivelTransactor(address common.Address, transactor bind.ContractTransactor) (*SwivelTransactor, error) {
	contract, err := bindSwivel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwivelTransactor{contract: contract}, nil
}

// NewSwivelFilterer creates a new log filterer instance of Swivel, bound to a specific deployed contract.
func NewSwivelFilterer(address common.Address, filterer bind.ContractFilterer) (*SwivelFilterer, error) {
	contract, err := bindSwivel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwivelFilterer{contract: contract}, nil
}

// bindSwivel binds a generic wrapper to an already deployed contract.
func bindSwivel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.SwivelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.SwivelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swivel *SwivelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swivel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swivel *SwivelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swivel *SwivelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swivel.Contract.contract.Transact(opts, method, params...)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCaller) InitiateCalledAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledAmount is a free data retrieval call binding the contract method 0xc268dc2a.
//
// Solidity: function initiateCalledAmount(address ) view returns(uint256)
func (_Swivel *SwivelCallerSession) InitiateCalledAmount(arg0 common.Address) (*big.Int, error) {
	return _Swivel.Contract.InitiateCalledAmount(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCaller) InitiateCalledSignature(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "initiateCalledSignature", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// InitiateCalledSignature is a free data retrieval call binding the contract method 0x38599fc8.
//
// Solidity: function initiateCalledSignature(address ) view returns(uint8)
func (_Swivel *SwivelCallerSession) InitiateCalledSignature(arg0 common.Address) (uint8, error) {
	return _Swivel.Contract.InitiateCalledSignature(&_Swivel.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 underlying, uint256 maturity)
func (_Swivel *SwivelCaller) RedeemZcTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Underlying *big.Int
	Maturity   *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "redeemZcTokenCalled", arg0)

	outstruct := new(struct {
		Underlying *big.Int
		Maturity   *big.Int
	})

	outstruct.Underlying = out[0].(*big.Int)
	outstruct.Maturity = out[1].(*big.Int)

	return *outstruct, err

}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 underlying, uint256 maturity)
func (_Swivel *SwivelSession) RedeemZcTokenCalled(arg0 common.Address) (struct {
	Underlying *big.Int
	Maturity   *big.Int
}, error) {
	return _Swivel.Contract.RedeemZcTokenCalled(&_Swivel.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0xddcfbbda.
//
// Solidity: function redeemZcTokenCalled(address ) view returns(uint256 underlying, uint256 maturity)
func (_Swivel *SwivelCallerSession) RedeemZcTokenCalled(arg0 common.Address) (struct {
	Underlying *big.Int
	Maturity   *big.Int
}, error) {
	return _Swivel.Contract.RedeemZcTokenCalled(&_Swivel.CallOpts, arg0)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactor) Initiate(opts *bind.TransactOpts, o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiate", o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// Initiate is a paid mutator transaction binding the contract method 0xd2144f58.
//
// Solidity: function initiate((bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(bool)
func (_Swivel *SwivelTransactorSession) Initiate(o []Order, a []*big.Int, s []Components) (*types.Transaction, error) {
	return _Swivel.Contract.Initiate(&_Swivel.TransactOpts, o, a, s)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactor) InitiateReturns(opts *bind.TransactOpts, i bool) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "initiateReturns", i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}

// InitiateReturns is a paid mutator transaction binding the contract method 0x4578b3ce.
//
// Solidity: function initiateReturns(bool i) returns()
func (_Swivel *SwivelTransactorSession) InitiateReturns(i bool) (*types.Transaction, error) {
	return _Swivel.Contract.InitiateReturns(&_Swivel.TransactOpts, i)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns()
func (_Swivel *SwivelTransactor) RedeemZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "redeemZcToken", u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns()
func (_Swivel *SwivelSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns()
func (_Swivel *SwivelTransactorSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Swivel.Contract.RedeemZcToken(&_Swivel.TransactOpts, u, m, a)
}
