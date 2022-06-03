// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redeemer

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

// RedeemerABI is the input ABI used to generate the binding from.
const RedeemerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apwineAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"contractMarketPlace\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"setMarketPlaceAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedeemerBin is the compiled bytecode used for deploying new contracts.
var RedeemerBin = "0x60806040523480156200001157600080fd5b5060405162001487380380620014878339810160408190526200003491620000af565b60008054336001600160a01b0319918216179091556002805482166001600160a01b0396871617905560038054821694861694909417909355600480548416928516929092179091556005805490921692169190911790556200010c565b80516001600160a01b0381168114620000aa57600080fd5b919050565b60008060008060808587031215620000c657600080fd5b620000d18562000092565b9350620000e16020860162000092565b9250620000f16040860162000092565b9150620001016060860162000092565b905092959194509250565b61136b806200011c6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063d2ba878811610081578063ea08c0311161005b578063ea08c031146101c7578063ef603569146101e7578063f851a4401461020757600080fd5b8063d2ba878814610181578063de1d3cb514610194578063e36e4b56146101b457600080fd5b8063a1b1138c116100b2578063a1b1138c14610138578063ba24ea6c1461015b578063cc98ef961461016e57600080fd5b80632e25d2a6146100ce57806337a3eeec14610118575b600080fd5b6001546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6005546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b61014b61014636600461117f565b610227565b604051901515815260200161010f565b61014b6101693660046111be565b6106d5565b61014b61017c36600461120f565b610ab1565b61014b61018f366004611233565b610ba6565b6004546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b61014b6101c2366004611277565b610db3565b6002546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b6003546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b6000546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f125cf47f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526024820184905260ff86166044830152600092839291169063125cf47f90606401602060405180830381865afa1580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf91906112dd565b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af1158015610347573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036b91906112fa565b60015490915061039490839073ffffffffffffffffffffffffffffffffffffffff163084611010565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff871601610470576002546040517f154e0f2e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201879052604482018490529091169063154e0f2e906064016020604051808303816000875af115801561043e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104629190611313565b61046b57600080fd5b610673565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd60ff87160161052c576001546040517f884e17f30000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff91821660248201529083169063884e17f3906044015b600060405180830381600087803b15801561050f57600080fd5b505af1158015610523573d6000803e3d6000fd5b50505050610673565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8716016105b6576040517f0e6dfcd5000000000000000000000000000000000000000000000000000000008152306004820181905260248201526044810182905273ffffffffffffffffffffffffffffffffffffffff831690630e6dfcd5906064016104f5565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff860ff871601610673576040517fd905777e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83169063d905777e906024016020604051808303816000875af115801561064c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067091906112fa565b90505b6040805160ff8816815260208101839052859173ffffffffffffffffffffffffffffffffffffffff8816917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600195945050505050565b6001546040517f125cf47f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260ff87166044830152600092839291169063125cf47f90606401602060405180830381865afa158015610759573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077d91906112dd565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529192506000918316906370a08231906024016020604051808303816000875af11580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081591906112fa565b905060ff8716156108465760015461084690879073ffffffffffffffffffffffffffffffffffffffff163084611010565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff881601610902576005546040517ff3fef3a300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490529091169063f3fef3a3906044015b600060405180830381600087803b1580156108e557600080fd5b505af11580156108f9573d6000803e3d6000fd5b50505050610a50565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff88160161099a57600480546040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116938201939093526024810184905260006044820152306064820152911690636c8d4fa1906084016108cb565b60ff8716610a50576040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201839052831690639dc29fac90604401600060405180830381600087803b158015610a1257600080fd5b505af1158015610a26573d6000803e3d6000fd5b5050600154610a50925088915073ffffffffffffffffffffffffffffffffffffffff163084611010565b604080516000815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a35060019695505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114610b39576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1615610b5c57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155915050919050565b6001546040517f125cf47f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260ff87166044830152600092839291169063125cf47f90606401602060405180830381865afa158015610c2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4e91906112dd565b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af1158015610cc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cea91906112fa565b600154909150610d1390839073ffffffffffffffffffffffffffffffffffffffff163084611010565b6003546040517fd230566c0000000000000000000000000000000000000000000000000000000081526004810186905273ffffffffffffffffffffffffffffffffffffffff8881166024830152604482018890529091169063d230566c90606401600060405180830381600087803b158015610d8e57600080fd5b505af1158015610da2573d6000803e3d6000fd5b5060019a9950505050505050505050565b6001546040517f125cf47f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260ff88166044830152600092839291169063125cf47f90606401602060405180830381865afa158015610e37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5b91906112dd565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820181905291925060009073ffffffffffffffffffffffffffffffffffffffff8416906370a08231906024016020604051808303816000875af1158015610ecf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef391906112fa565b600154909150610f1c90849073ffffffffffffffffffffffffffffffffffffffff168484611010565b6040517f2b83cccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820189905260448201839052871690632b83cccd90606401600060405180830381600087803b158015610f9357600080fd5b505af1158015610fa7573d6000803e3d6000fd5b50506040805160ff8d168152602081018590528a935073ffffffffffffffffffffffffffffffffffffffff8c1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600198975050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af191505061108d816110fa565b6110f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c65640000000000000000000000006044820152606401610b30565b5050505050565b6000803d8361110d57806000803e806000fd5b8060208114611125578015611136576000925061113b565b816000803e6000511515925061113b565b600192505b50909392505050565b803560ff8116811461115557600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461117c57600080fd5b50565b60008060006060848603121561119457600080fd5b61119d84611144565b925060208401356111ad8161115a565b929592945050506040919091013590565b600080600080608085870312156111d457600080fd5b6111dd85611144565b935060208501356111ed8161115a565b92506040850135915060608501356112048161115a565b939692955090935050565b60006020828403121561122157600080fd5b813561122c8161115a565b9392505050565b6000806000806080858703121561124957600080fd5b61125285611144565b935060208501356112628161115a565b93969395505050506040820135916060013590565b600080600080600060a0868803121561128f57600080fd5b61129886611144565b945060208601356112a88161115a565b93506040860135925060608601356112bf8161115a565b915060808601356112cf8161115a565b809150509295509295909350565b6000602082840312156112ef57600080fd5b815161122c8161115a565b60006020828403121561130c57600080fd5b5051919050565b60006020828403121561132557600080fd5b8151801515811461122c57600080fdfea26469706673582212201255ddc3c4b25a1aacdd813256fa975c6b5a7bba808a7ccc315a71262e2d7a9764736f6c634300080d0033"

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, s common.Address, p common.Address, t common.Address, a common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedeemerBin), backend, s, p, t, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// Redeemer is an auto generated Go binding around an Ethereum contract.
type Redeemer struct {
	RedeemerCaller     // Read-only binding to the contract
	RedeemerTransactor // Write-only binding to the contract
	RedeemerFilterer   // Log filterer for contract events
}

// RedeemerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedeemerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedeemerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedeemerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedeemerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedeemerSession struct {
	Contract     *Redeemer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedeemerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedeemerCallerSession struct {
	Contract *RedeemerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RedeemerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedeemerTransactorSession struct {
	Contract     *RedeemerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RedeemerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedeemerRaw struct {
	Contract *Redeemer // Generic contract binding to access the raw methods on
}

// RedeemerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedeemerCallerRaw struct {
	Contract *RedeemerCaller // Generic read-only contract binding to access the raw methods on
}

// RedeemerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedeemerTransactorRaw struct {
	Contract *RedeemerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedeemer creates a new instance of Redeemer, bound to a specific deployed contract.
func NewRedeemer(address common.Address, backend bind.ContractBackend) (*Redeemer, error) {
	contract, err := bindRedeemer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redeemer{RedeemerCaller: RedeemerCaller{contract: contract}, RedeemerTransactor: RedeemerTransactor{contract: contract}, RedeemerFilterer: RedeemerFilterer{contract: contract}}, nil
}

// NewRedeemerCaller creates a new read-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerCaller(address common.Address, caller bind.ContractCaller) (*RedeemerCaller, error) {
	contract, err := bindRedeemer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerCaller{contract: contract}, nil
}

// NewRedeemerTransactor creates a new write-only instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerTransactor(address common.Address, transactor bind.ContractTransactor) (*RedeemerTransactor, error) {
	contract, err := bindRedeemer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedeemerTransactor{contract: contract}, nil
}

// NewRedeemerFilterer creates a new log filterer instance of Redeemer, bound to a specific deployed contract.
func NewRedeemerFilterer(address common.Address, filterer bind.ContractFilterer) (*RedeemerFilterer, error) {
	contract, err := bindRedeemer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedeemerFilterer{contract: contract}, nil
}

// bindRedeemer binds a generic wrapper to an already deployed contract.
func bindRedeemer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.RedeemerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.RedeemerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redeemer *RedeemerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redeemer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redeemer *RedeemerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redeemer *RedeemerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redeemer.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Redeemer *RedeemerCallerSession) Admin() (common.Address, error) {
	return _Redeemer.Contract.Admin(&_Redeemer.CallOpts)
}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCaller) ApwineAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "apwineAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
}

// ApwineAddr is a free data retrieval call binding the contract method 0x37a3eeec.
//
// Solidity: function apwineAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) ApwineAddr() (common.Address, error) {
	return _Redeemer.Contract.ApwineAddr(&_Redeemer.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerSession) MarketPlace() (common.Address, error) {
	return _Redeemer.Contract.MarketPlace(&_Redeemer.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Redeemer *RedeemerCallerSession) MarketPlace() (common.Address, error) {
	return _Redeemer.Contract.MarketPlace(&_Redeemer.CallOpts)
}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCaller) PendleAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "pendleAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// PendleAddr is a free data retrieval call binding the contract method 0xef603569.
//
// Solidity: function pendleAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) PendleAddr() (common.Address, error) {
	return _Redeemer.Contract.PendleAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCaller) SwivelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "swivelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) SwivelAddr() (common.Address, error) {
	return _Redeemer.Contract.SwivelAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCaller) TempusAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "tempusAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
}

// TempusAddr is a free data retrieval call binding the contract method 0xde1d3cb5.
//
// Solidity: function tempusAddr() view returns(address)
func (_Redeemer *RedeemerCallerSession) TempusAddr() (common.Address, error) {
	return _Redeemer.Contract.TempusAddr(&_Redeemer.CallOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem", p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem is a paid mutator transaction binding the contract method 0xa1b1138c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem(&_Redeemer.TransactOpts, p, u, m)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem0", p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba24ea6c.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address o) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem0(p uint8, u common.Address, m *big.Int, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem0(&_Redeemer.TransactOpts, p, u, m, o)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem1(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem1", p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem1 is a paid mutator transaction binding the contract method 0xd2ba8788.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, bytes32 i) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem1(p uint8, u common.Address, m *big.Int, i [32]byte) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem1(&_Redeemer.TransactOpts, p, u, m, i)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerTransactor) Redeem2(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "redeem2", p, u, m, d, o)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, o)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xe36e4b56.
//
// Solidity: function redeem(uint8 p, address u, uint256 m, address d, address o) returns(bool)
func (_Redeemer *RedeemerTransactorSession) Redeem2(p uint8, u common.Address, m *big.Int, d common.Address, o common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.Redeem2(&_Redeemer.TransactOpts, p, u, m, d, o)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerTransactor) SetMarketPlaceAddress(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "setMarketPlaceAddress", m)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerSession) SetMarketPlaceAddress(m common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetMarketPlaceAddress(&_Redeemer.TransactOpts, m)
}

// SetMarketPlaceAddress is a paid mutator transaction binding the contract method 0xcc98ef96.
//
// Solidity: function setMarketPlaceAddress(address m) returns(bool)
func (_Redeemer *RedeemerTransactorSession) SetMarketPlaceAddress(m common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetMarketPlaceAddress(&_Redeemer.TransactOpts, m)
}

// RedeemerRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Redeemer contract.
type RedeemerRedeemIterator struct {
	Event *RedeemerRedeem // Event containing the contract specifics and raw log

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
func (it *RedeemerRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedeemerRedeem)
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
		it.Event = new(RedeemerRedeem)
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
func (it *RedeemerRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedeemerRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedeemerRedeem represents a Redeem event raised by the Redeemer contract.
type RedeemerRedeem struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) FilterRedeem(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*RedeemerRedeemIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.FilterLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &RedeemerRedeemIterator{contract: _Redeemer.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *RedeemerRedeem, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Redeemer.contract.WatchLogs(opts, "Redeem", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedeemerRedeem)
				if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xfd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9.
//
// Solidity: event Redeem(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Redeemer *RedeemerFilterer) ParseRedeem(log types.Log) (*RedeemerRedeem, error) {
	event := new(RedeemerRedeem)
	if err := _Redeemer.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
