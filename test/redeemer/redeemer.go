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
const RedeemerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apwineAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"d\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l\",\"type\":\"address\"}],\"name\":\"setLenderAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"setMarketPlaceAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RedeemerBin is the compiled bytecode used for deploying new contracts.
var RedeemerBin = "0x60806040523480156200001157600080fd5b50604051620019a9380380620019a98339810160408190526200003491620000be565b60008054336001600160a01b0319918216179091556002805482166001600160a01b0397881617905560038054821695871695909517909455600480548516938616939093179092556005805484169185169190911790556006805490921692169190911790556200012e565b80516001600160a01b0381168114620000b957600080fd5b919050565b600080600080600060a08688031215620000d757600080fd5b620000e286620000a1565b9450620000f260208701620000a1565b93506200010260408701620000a1565b92506200011260608701620000a1565b91506200012260808701620000a1565b90509295509295909350565b61186b806200013e6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063cc98ef961161008c578063e36e4b5611610066578063e36e4b561461021b578063ea08c0311461022e578063ef6035691461024e578063f851a4401461026e57600080fd5b8063cc98ef96146101d5578063d2ba8788146101e8578063de1d3cb5146101fb57600080fd5b8063a1b1138c116100c8578063a1b1138c1461017c578063b4d7813d1461018f578063ba24ea6c146101a2578063bcead63e146101b557600080fd5b80632e25d2a6146100ef57806337a3eeec1461013957806370a03ced14610159575b600080fd5b60015461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b60065461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c61016736600461160e565b61028e565b6040519015158152602001610130565b61016c61018a36600461167f565b6104c1565b61016c61019d3660046116be565b610a06565b61016c6101b03660046116e2565b610af8565b60025461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c6101e33660046116be565b610f3b565b61016c6101f6366004611733565b61102b565b60055461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b61016c610229366004611777565b611247565b60035461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60045461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b60005461010f9073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260006044830181905292169063ca1695f0906064016020604051808303816000875af115801561030f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033391906117dd565b3373ffffffffffffffffffffffffffffffffffffffff8216146103b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89811660048301526024820189905260006044830181905292169063ca1695f0906064016020604051808303816000875af1158015610438573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045c91906117dd565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f676f74206865726500000000000000000000000000000000000000000000000060448201529091506064016103ae565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301526024820184905260ff86166044830152600092839291169063ca1695f0906064016020604051808303816000875af1158015610547573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056b91906117dd565b905060ff851660011480610582575060ff85166003145b80610590575060ff85166002145b8061059e575060ff85166008145b610604576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207072696e636970616c00000000000000000000000000000060448201526064016103ae565b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526000918316906370a08231906024016020604051808303816000875af1158015610678573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069c91906117fa565b6002549091506106c590839073ffffffffffffffffffffffffffffffffffffffff1630846114b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff8716016107a1576003546040517f154e0f2e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015260248201879052604482018490529091169063154e0f2e906064016020604051808303816000875af115801561076f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107939190611813565b61079c57600080fd5b6109a4565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd60ff87160161085d576001546040517f884e17f30000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff91821660248201529083169063884e17f3906044015b600060405180830381600087803b15801561084057600080fd5b505af1158015610854573d6000803e3d6000fd5b505050506109a4565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8716016108e7576040517f0e6dfcd5000000000000000000000000000000000000000000000000000000008152306004820181905260248201526044810182905273ffffffffffffffffffffffffffffffffffffffff831690630e6dfcd590606401610826565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff860ff8716016109a4576040517fd905777e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83169063d905777e906024016020604051808303816000875af115801561097d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a191906117fa565b90505b6040805160ff8816815260208101839052859173ffffffffffffffffffffffffffffffffffffffff8816917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600195945050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114610a89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064016103ae565b60025473ffffffffffffffffffffffffffffffffffffffff1615610aac57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556001915050919050565b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260ff87166044830152600092839291169063ca1695f0906064016020604051808303816000875af1158015610b7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba291906117dd565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301529192506000918316906370a08231906024016020604051808303816000875af1158015610c16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3a91906117fa565b905060ff871615610c6b57600254610c6b90879073ffffffffffffffffffffffffffffffffffffffff1630846114b5565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff960ff881601610d27576006546040517ff3fef3a300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490529091169063f3fef3a3906044015b600060405180830381600087803b158015610d0a57600080fd5b505af1158015610d1e573d6000803e3d6000fd5b50505050610eda565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb60ff881601610dbd576005546040517f6c8d4fa100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018490526000604483015230606483015290911690636c8d4fa190608401610cf0565b60ff8716610e78576040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201839052831690639dc29fac90604401600060405180830381600087803b158015610e3557600080fd5b505af1158015610e49573d6000803e3d6000fd5b5050600254610e73925088915073ffffffffffffffffffffffffffffffffffffffff1630846114b5565b610eda565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c6964207072696e636970616c00000000000000000000000000000060448201526064016103ae565b604080516000815260208101839052869173ffffffffffffffffffffffffffffffffffffffff8916917ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a35060019695505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114610fbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064016103ae565b60015473ffffffffffffffffffffffffffffffffffffffff1615610fe157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155915050919050565b600060ff851660041461103d57600080fd5b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260ff88166044830152600092169063ca1695f0906064016020604051808303816000875af11580156110c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e491906117dd565b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152919250600091908316906370a08231906024016020604051808303816000875af115801561115c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118091906117fa565b6002549091506111a990839073ffffffffffffffffffffffffffffffffffffffff1630846114b5565b600480546040517fd230566c00000000000000000000000000000000000000000000000000000000815291820186905273ffffffffffffffffffffffffffffffffffffffff888116602484015260448301889052169063d230566c90606401600060405180830381600087803b15801561122257600080fd5b505af1158015611236573d6000803e3d6000fd5b5060019a9950505050505050505050565b600060ff861660061461125957600080fd5b6001546040517fca1695f000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260ff89166044830152600092169063ca1695f0906064016020604051808303816000875af11580156112dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130091906117dd565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820181905291925060009073ffffffffffffffffffffffffffffffffffffffff8416906370a08231906024016020604051808303816000875af1158015611374573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139891906117fa565b6002549091506113c190849073ffffffffffffffffffffffffffffffffffffffff1684846114b5565b6040517f2b83cccd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820189905260448201839052871690632b83cccd90606401600060405180830381600087803b15801561143857600080fd5b505af115801561144c573d6000803e3d6000fd5b50506040805160ff8d168152602081018590528a935073ffffffffffffffffffffffffffffffffffffffff8c1692507ffd23960ad64d756515d736fc98d5f473529ef08157d573f8a13c1723cc632fe9910160405180910390a350600198975050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af19150506115328161159f565b611598576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064016103ae565b5050505050565b6000803d836115b257806000803e806000fd5b80602081146115ca5780156115db57600092506115e0565b816000803e600051151592506115e0565b600192505b50909392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461160b57600080fd5b50565b600080600080600060a0868803121561162657600080fd5b8535611631816115e9565b9450602086013593506040860135611648816115e9565b92506060860135611658816115e9565b949793965091946080013592915050565b803560ff8116811461167a57600080fd5b919050565b60008060006060848603121561169457600080fd5b61169d84611669565b925060208401356116ad816115e9565b929592945050506040919091013590565b6000602082840312156116d057600080fd5b81356116db816115e9565b9392505050565b600080600080608085870312156116f857600080fd5b61170185611669565b93506020850135611711816115e9565b9250604085013591506060850135611728816115e9565b939692955090935050565b6000806000806080858703121561174957600080fd5b61175285611669565b93506020850135611762816115e9565b93969395505050506040820135916060013590565b600080600080600060a0868803121561178f57600080fd5b61179886611669565b945060208601356117a8816115e9565b93506040860135925060608601356117bf816115e9565b915060808601356117cf816115e9565b809150509295509295909350565b6000602082840312156117ef57600080fd5b81516116db816115e9565b60006020828403121561180c57600080fd5b5051919050565b60006020828403121561182557600080fd5b815180151581146116db57600080fdfea26469706673582212202675c49095156e837193b76bd96e8050bdaedd8cfc0e7affdfcc43675309345364736f6c634300080d0033"

// DeployRedeemer deploys a new Ethereum contract, binding an instance of Redeemer to it.
func DeployRedeemer(auth *bind.TransactOpts, backend bind.ContractBackend, l common.Address, s common.Address, p common.Address, t common.Address, a common.Address) (common.Address, *types.Transaction, *Redeemer, error) {
	parsed, err := abi.JSON(strings.NewReader(RedeemerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RedeemerBin), backend, l, s, p, t, a)
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

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redeemer.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerSession) Lender() (common.Address, error) {
	return _Redeemer.Contract.Lender(&_Redeemer.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Redeemer *RedeemerCallerSession) Lender() (common.Address, error) {
	return _Redeemer.Contract.Lender(&_Redeemer.CallOpts)
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

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerTransactor) AuthRedeem(opts *bind.TransactOpts, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "authRedeem", u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerSession) AuthRedeem(u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.AuthRedeem(&_Redeemer.TransactOpts, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x70a03ced.
//
// Solidity: function authRedeem(address u, uint256 m, address f, address t, uint256 a) returns(bool)
func (_Redeemer *RedeemerTransactorSession) AuthRedeem(u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _Redeemer.Contract.AuthRedeem(&_Redeemer.TransactOpts, u, m, f, t, a)
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

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerTransactor) SetLenderAddress(opts *bind.TransactOpts, l common.Address) (*types.Transaction, error) {
	return _Redeemer.contract.Transact(opts, "setLenderAddress", l)
}

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerSession) SetLenderAddress(l common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetLenderAddress(&_Redeemer.TransactOpts, l)
}

// SetLenderAddress is a paid mutator transaction binding the contract method 0xb4d7813d.
//
// Solidity: function setLenderAddress(address l) returns(bool)
func (_Redeemer *RedeemerTransactorSession) SetLenderAddress(l common.Address) (*types.Transaction, error) {
	return _Redeemer.Contract.SetLenderAddress(&_Redeemer.TransactOpts, l)
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
