// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lender

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

// SwivelComponents is an auto generated low-level Go binding around an user-defined struct.
type SwivelComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelOrder is an auto generated low-level Go binding around an user-defined struct.
type SwivelOrder struct {
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

// LenderMetaData contains all meta data concerning the Lender contract.
var LenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001b1838038062001b1883398101604081905262000034916200008e565b600080546001600160a01b03199081163317909155600180546001600160a01b0394851690831617905560028054929093169116179055620000c6565b80516001600160a01b03811681146200008957600080fd5b919050565b60008060408385031215620000a257600080fd5b620000ad8362000071565b9150620000bd6020840162000071565b90509250929050565b611a4280620000d66000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063ea08c03111610050578063ea08c031146100f5578063f1aa654e14610115578063f851a4401461012857600080fd5b80632e25d2a6146100775780635b1e5fc2146100c1578063d5907c36146100e2575b600080fd5b6001546100979073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d46100cf3660046111ad565b610148565b6040519081526020016100b8565b6100d46100f0366004611289565b610436565b6002546100979073ffffffffffffffffffffffffffffffffffffffff1681565b6100d46101233660046112fe565b6108ce565b6000546100979073ffffffffffffffffffffffffffffffffffffffff1681565b60008080805b8881101561030b5760008a8a8381811061016a5761016a611357565b9050610120020180360381019061018191906113f8565b90508c8160e00151146101f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f73776976656c206d6174757269747920213d206d61747572697479000000000060448201526064015b60405180910390fd5b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff161461028e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f73776976656c20756e6465726c79696e6720213d20756e6465726c79696e670060448201526064016101ec565b8888838181106102a0576102a0611357565b90506020020135846102b291906114be565b93508060a001518160c001518a8a858181106102d0576102d0611357565b905060200201356102e191906114d6565b6102eb9190611513565b6102f590846114be565b92505080806103039061154e565b91505061014e565b506103188c333085610e35565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f5890610378908c908c908c908c908c908c9060040161162a565b6020604051808303816000875af1158015610397573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103bb9190611733565b506103d08c8b6103cb8585611757565b610f1f565b6040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b600061044488333087610e35565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905260009216906317b3bba790604401610100604051808303816000875af11580156104bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e3919061176e565b8a60ff16600881106104f7576104f7611357565b602002015190508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610562573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058691906117f6565b73ffffffffffffffffffffffffffffffffffffffff16146105dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016101ec565b878173ffffffffffffffffffffffffffffffffffffffff1663aa082a9d6040518163ffffffff1660e01b81526004016020604051808303816000875af115801561062b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064f9190611813565b14610690576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016101ec565b6040805160e081018252600060c08201818152825260208201899052918101879052606081018281526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152509050600060405180608001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152509050898b73ffffffffffffffffffffffffffffffffffffffff167f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d98e8c73ffffffffffffffffffffffffffffffffffffffff166369bbd8cd87878d8d6040518563ffffffff1660e01b81526004016107c49493929190611867565b6020604051808303816000875af11580156107e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108079190611813565b6040805160ff909316835260208301919091520160405180910390a36040517f69bbd8cd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906369bbd8cd9061087b90859085908b908b90600401611867565b6020604051808303816000875af115801561089a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108be9190611813565b9c9b505050505050505050505050565b6000808390508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635001f3b56040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095c91906117f6565b73ffffffffffffffffffffffffffffffffffffffff16146109d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7969656c64206261736520213d20756e6465726c79696e67000000000000000060448201526064016101ec565b848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610a27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4b91906119b4565b63ffffffff1614610ab8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7969656c64206d6174757269747920213d206d6174757269747900000000000060448201526064016101ec565b8530610ac682338388610e35565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c610aed88610fec565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af1158015610b54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7891906119da565b9050610b85838888610f1f565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af1158015610c0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2f91906119da565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b1601610daf576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af1158015610cd5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf9919061176e565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015610d88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dac9190611733565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af1915050610eb281611046565b610f18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064016101ec565b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050610f8081611046565b610fe6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c6564000000000000000000000000000000000060448201526064016101ec565b50505050565b60006fffffffffffffffffffffffffffffffff821115611042576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016101ec565b5090565b6000803d8361105957806000803e806000fd5b80602081146110715780156110825760009250611087565b816000803e60005115159250611087565b600192505b50909392505050565b803560ff811681146110a157600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146110c857600080fd5b50565b80356110a1816110a6565b60008083601f8401126110e857600080fd5b50813567ffffffffffffffff81111561110057600080fd5b6020830191508360206101208302850101111561111c57600080fd5b9250929050565b60008083601f84011261113557600080fd5b50813567ffffffffffffffff81111561114d57600080fd5b6020830191508360208260051b850101111561111c57600080fd5b60008083601f84011261117a57600080fd5b50813567ffffffffffffffff81111561119257600080fd5b60208301915083602060608302850101111561111c57600080fd5b60008060008060008060008060008060e08b8d0312156111cc57600080fd5b6111d58b611090565b995060208b01356111e5816110a6565b985060408b0135975060608b01356111fc816110a6565b965060808b013567ffffffffffffffff8082111561121957600080fd5b6112258e838f016110d6565b909850965060a08d013591508082111561123e57600080fd5b61124a8e838f01611123565b909650945060c08d013591508082111561126357600080fd5b506112708d828e01611168565b915080935050809150509295989b9194979a5092959850565b600080600080600080600080610100898b0312156112a657600080fd5b6112af89611090565b975060208901356112bf816110a6565b96506040890135955060608901356112d6816110a6565b979a969950949760808101359660a0820135965060c0820135955060e0909101359350915050565b600080600080600060a0868803121561131657600080fd5b61131f86611090565b9450602086013561132f816110a6565b9350604086013592506060860135611346816110a6565b949793965091946080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff811182821017156113d9576113d9611386565b60405290565b80151581146110c857600080fd5b80356110a1816113df565b6000610120828403121561140b57600080fd5b6114136113b5565b82358152611423602084016110cb565b6020820152611434604084016110cb565b6040820152611445606084016113ed565b6060820152611456608084016113ed565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156114d1576114d161148f565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561150e5761150e61148f565b500290565b600082611549577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361157f5761157f61148f565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156115b857600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b8581101561161f5760ff6115f883611090565b168752818301358388015260408083013590880152606096870196909101906001016115e5565b509495945050505050565b606080825281810187905260009060808084018a845b8b8110156116fa578135835260208083013561165b816110a6565b73ffffffffffffffffffffffffffffffffffffffff169084015260406116828382016110cb565b73ffffffffffffffffffffffffffffffffffffffff16908401526116a78286016113ed565b1515858401526116b88285016113ed565b15158484015260a0828101359084015260c0808301359084015260e0808301359084015261010080830135908401526101209283019290910190600101611640565b5050848103602086015261170f81898b611586565b9250505082810360408401526117268185876115d5565b9998505050505050505050565b60006020828403121561174557600080fd5b8151611750816113df565b9392505050565b6000828210156117695761176961148f565b500390565b600061010080838503121561178257600080fd5b83601f84011261179157600080fd5b60405181810181811067ffffffffffffffff821117156117b3576117b3611386565b6040529083019080858311156117c857600080fd5b845b838110156117eb5780516117dd816110a6565b8252602091820191016117ca565b509095945050505050565b60006020828403121561180857600080fd5b8151611750816110a6565b60006020828403121561182557600080fd5b5051919050565b60028110611863577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60e081526000855160c060e08401528051806101a085015260005b818110156118a05760208184018101516101c0878401015201611882565b818111156118b35760006101c083870101525b5060208801516101008501526040880151610120850152606088015191506118df61014085018361182c565b608088015173ffffffffffffffffffffffffffffffffffffffff8116610160860152915060a088015173ffffffffffffffffffffffffffffffffffffffff81166101808601529150601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683016101c00191506119a29050602083018673ffffffffffffffffffffffffffffffffffffffff808251168352806020830151166020840152506040810151151560408301526060810151151560608301525050565b60a082019390935260c0015292915050565b6000602082840312156119c657600080fd5b815163ffffffff8116811461175057600080fd5b6000602082840312156119ec57600080fd5b81516fffffffffffffffffffffffffffffffff8116811461175057600080fdfea2646970667358221220c41ca530ea29e0e17f25cdcede8eee6034625f3370a391027866a3d5351cc5c064736f6c634300080d0033",
}

// LenderABI is the input ABI used to generate the binding from.
// Deprecated: Use LenderMetaData.ABI instead.
var LenderABI = LenderMetaData.ABI

// LenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LenderMetaData.Bin instead.
var LenderBin = LenderMetaData.Bin

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, s common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := LenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LenderBin), backend, m, s)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// Lender is an auto generated Go binding around an Ethereum contract.
type Lender struct {
	LenderCaller     // Read-only binding to the contract
	LenderTransactor // Write-only binding to the contract
	LenderFilterer   // Log filterer for contract events
}

// LenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LenderSession struct {
	Contract     *Lender           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LenderCallerSession struct {
	Contract *LenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LenderTransactorSession struct {
	Contract     *LenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LenderRaw struct {
	Contract *Lender // Generic contract binding to access the raw methods on
}

// LenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LenderCallerRaw struct {
	Contract *LenderCaller // Generic read-only contract binding to access the raw methods on
}

// LenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LenderTransactorRaw struct {
	Contract *LenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLender creates a new instance of Lender, bound to a specific deployed contract.
func NewLender(address common.Address, backend bind.ContractBackend) (*Lender, error) {
	contract, err := bindLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lender{LenderCaller: LenderCaller{contract: contract}, LenderTransactor: LenderTransactor{contract: contract}, LenderFilterer: LenderFilterer{contract: contract}}, nil
}

// NewLenderCaller creates a new read-only instance of Lender, bound to a specific deployed contract.
func NewLenderCaller(address common.Address, caller bind.ContractCaller) (*LenderCaller, error) {
	contract, err := bindLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LenderCaller{contract: contract}, nil
}

// NewLenderTransactor creates a new write-only instance of Lender, bound to a specific deployed contract.
func NewLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*LenderTransactor, error) {
	contract, err := bindLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LenderTransactor{contract: contract}, nil
}

// NewLenderFilterer creates a new log filterer instance of Lender, bound to a specific deployed contract.
func NewLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*LenderFilterer, error) {
	contract, err := bindLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LenderFilterer{contract: contract}, nil
}

// bindLender binds a generic wrapper to an already deployed contract.
func bindLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.LenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.LenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lender *LenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lender *LenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lender *LenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lender.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderSession) Admin() (common.Address, error) {
	return _Lender.Contract.Admin(&_Lender.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Lender *LenderCallerSession) Admin() (common.Address, error) {
	return _Lender.Contract.Admin(&_Lender.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderSession) MarketPlace() (common.Address, error) {
	return _Lender.Contract.MarketPlace(&_Lender.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Lender *LenderCallerSession) MarketPlace() (common.Address, error) {
	return _Lender.Contract.MarketPlace(&_Lender.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderCaller) SwivelAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "swivelAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderSession) SwivelAddr() (common.Address, error) {
	return _Lender.Contract.SwivelAddr(&_Lender.CallOpts)
}

// SwivelAddr is a free data retrieval call binding the contract method 0xea08c031.
//
// Solidity: function swivelAddr() view returns(address)
func (_Lender *LenderCallerSession) SwivelAddr() (common.Address, error) {
	return _Lender.Contract.SwivelAddr(&_Lender.CallOpts)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactor) Lend(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend", p, u, m, y, o, a, s)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderSession) Lend(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactorSession) Lend(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend0 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderTransactor) Lend0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend0", p, u, m, e, i, a, r, d)
}

// Lend0 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderSession) Lend0(p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, e, i, a, r, d)
}

// Lend0 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderTransactorSession) Lend0(p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, e, i, a, r, d)
}

// Lend1 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactor) Lend1(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend1", p, u, m, y, a)
}

// Lend1 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderSession) Lend1(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend1(&_Lender.TransactOpts, p, u, m, y, a)
}

// Lend1 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactorSession) Lend1(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend1(&_Lender.TransactOpts, p, u, m, y, a)
}

// LenderLendIterator is returned from FilterLend and is used to iterate over the raw logs and unpacked data for Lend events raised by the Lender contract.
type LenderLendIterator struct {
	Event *LenderLend // Event containing the contract specifics and raw log

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
func (it *LenderLendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LenderLend)
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
		it.Event = new(LenderLend)
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
func (it *LenderLendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LenderLendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LenderLend represents a Lend event raised by the Lender contract.
type LenderLend struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Returned   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLend is a free log retrieval operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) FilterLend(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*LenderLendIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.FilterLogs(opts, "Lend", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &LenderLendIterator{contract: _Lender.contract, event: "Lend", logs: logs, sub: sub}, nil
}

// WatchLend is a free log subscription operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) WatchLend(opts *bind.WatchOpts, sink chan<- *LenderLend, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.WatchLogs(opts, "Lend", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LenderLend)
				if err := _Lender.contract.UnpackLog(event, "Lend", log); err != nil {
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

// ParseLend is a log parse operation binding the contract event 0x26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9.
//
// Solidity: event Lend(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 returned)
func (_Lender *LenderFilterer) ParseLend(log types.Log) (*LenderLend, error) {
	event := new(LenderLend)
	if err := _Lender.contract.UnpackLog(event, "Lend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
