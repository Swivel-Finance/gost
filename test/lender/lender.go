// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lender

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

// LenderABI is the input ABI used to generate the binding from.
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"i\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"illuminate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x60806040523480156200001157600080fd5b5060405162001d5638038062001d5683398101604081905262000034916200008e565b600080546001600160a01b03199081163317909155600180546001600160a01b0394851690831617905560028054929093169116179055620000c6565b80516001600160a01b03811681146200008957600080fd5b919050565b60008060408385031215620000a257600080fd5b620000ad8362000071565b9150620000bd6020840162000071565b90509250929050565b611c8080620000d66000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063dc4c7ca91161005b578063dc4c7ca914610100578063ea08c03114610123578063f1aa654e14610143578063f851a4401461015657600080fd5b80635b1e5fc214610082578063cbf7e670146100a8578063d5907c36146100ed575b600080fd5b6100956100903660046113a7565b610176565b6040519081526020015b60405180910390f35b6001546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161009f565b6100956100fb366004611483565b610464565b61011361010e3660046114f8565b6108fc565b604051901515815260200161009f565b6002546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b61009561015136600461153c565b610ac8565b6000546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b60008080805b888110156103395760008a8a8381811061019857610198611595565b905061012002018036038101906101af9190611636565b90508c8160e0015114610223576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f73776976656c206d6174757269747920213d206d61747572697479000000000060448201526064015b60405180910390fd5b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16146102bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f73776976656c20756e6465726c79696e6720213d20756e6465726c79696e6700604482015260640161021a565b8888838181106102ce576102ce611595565b90506020020135846102e091906116fc565b93508060a001518160c001518a8a858181106102fe576102fe611595565b9050602002013561030f9190611714565b6103199190611751565b61032390846116fc565b92505080806103319061178c565b91505061017c565b506103468c33308561102f565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f58906103a6908c908c908c908c908c908c90600401611868565b6020604051808303816000875af11580156103c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e99190611971565b506103fe8c8b6103f98585611995565b611119565b6040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b60006104728833308761102f565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905260009216906317b3bba790604401610100604051808303816000875af11580156104ed573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051191906119ac565b8a60ff166008811061052557610525611595565b602002015190508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610590573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b49190611a34565b73ffffffffffffffffffffffffffffffffffffffff161461060b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161021a565b878173ffffffffffffffffffffffffffffffffffffffff1663aa082a9d6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610659573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067d9190611a51565b146106be576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161021a565b6040805160e081018252600060c08201818152825260208201899052918101879052606081018281526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152509050600060405180608001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152509050898b73ffffffffffffffffffffffffffffffffffffffff167f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d98e8c73ffffffffffffffffffffffffffffffffffffffff166369bbd8cd87878d8d6040518563ffffffff1660e01b81526004016107f29493929190611aa5565b6020604051808303816000875af1158015610811573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108359190611a51565b6040805160ff909316835260208301919091520160405180910390a36040517f69bbd8cd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906369bbd8cd906108a990859085908b908b90600401611aa5565b6020604051808303816000875af11580156108c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ec9190611a51565b9c9b505050505050505050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260009283929116906317b3bba790604401610100604051808303816000875af115801561097a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099e91906119ac565b90506109c5818760ff16600881106109b8576109b8611595565b602002015133308661102f565b80600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810185905273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015610a41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a659190611971565b506040805160ff8816815260208101859052859173ffffffffffffffffffffffffffffffffffffffff8816917f309b03ba657e17f1beadbc6eb3c06ba79b38084eb8d0e5452cc222462a17f1f6910160405180910390a350600195945050505050565b6000808390508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635001f3b56040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610b32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b569190611a34565b73ffffffffffffffffffffffffffffffffffffffff1614610bd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7969656c64206261736520213d20756e6465726c79696e670000000000000000604482015260640161021a565b848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610c21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c459190611bf2565b63ffffffff1614610cb2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7969656c64206d6174757269747920213d206d61747572697479000000000000604482015260640161021a565b8530610cc08233838861102f565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c610ce7886111e6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af1158015610d4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d729190611c18565b9050610d7f838888611119565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af1158015610e05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e299190611c18565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b1601610fa9576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af1158015610ecf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef391906119ac565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015610f82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa69190611971565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af19150506110ac81611240565b611112576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c6564000000000000000000000000604482015260640161021a565b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af191505061117a81611240565b6111e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015260640161021a565b50505050565b60006fffffffffffffffffffffffffffffffff82111561123c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161021a565b5090565b6000803d8361125357806000803e806000fd5b806020811461126b57801561127c5760009250611281565b816000803e60005115159250611281565b600192505b50909392505050565b803560ff8116811461129b57600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146112c257600080fd5b50565b803561129b816112a0565b60008083601f8401126112e257600080fd5b50813567ffffffffffffffff8111156112fa57600080fd5b6020830191508360206101208302850101111561131657600080fd5b9250929050565b60008083601f84011261132f57600080fd5b50813567ffffffffffffffff81111561134757600080fd5b6020830191508360208260051b850101111561131657600080fd5b60008083601f84011261137457600080fd5b50813567ffffffffffffffff81111561138c57600080fd5b60208301915083602060608302850101111561131657600080fd5b60008060008060008060008060008060e08b8d0312156113c657600080fd5b6113cf8b61128a565b995060208b01356113df816112a0565b985060408b0135975060608b01356113f6816112a0565b965060808b013567ffffffffffffffff8082111561141357600080fd5b61141f8e838f016112d0565b909850965060a08d013591508082111561143857600080fd5b6114448e838f0161131d565b909650945060c08d013591508082111561145d57600080fd5b5061146a8d828e01611362565b915080935050809150509295989b9194979a5092959850565b600080600080600080600080610100898b0312156114a057600080fd5b6114a98961128a565b975060208901356114b9816112a0565b96506040890135955060608901356114d0816112a0565b979a969950949760808101359660a0820135965060c0820135955060e0909101359350915050565b6000806000806080858703121561150e57600080fd5b6115178561128a565b93506020850135611527816112a0565b93969395505050506040820135916060013590565b600080600080600060a0868803121561155457600080fd5b61155d8661128a565b9450602086013561156d816112a0565b9350604086013592506060860135611584816112a0565b949793965091946080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715611617576116176115c4565b60405290565b80151581146112c257600080fd5b803561129b8161161d565b6000610120828403121561164957600080fd5b6116516115f3565b82358152611661602084016112c5565b6020820152611672604084016112c5565b60408201526116836060840161162b565b60608201526116946080840161162b565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561170f5761170f6116cd565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561174c5761174c6116cd565b500290565b600082611787577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117bd576117bd6116cd565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156117f657600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b8581101561185d5760ff6118368361128a565b16875281830135838801526040808301359088015260609687019690910190600101611823565b509495945050505050565b606080825281810187905260009060808084018a845b8b8110156119385781358352602080830135611899816112a0565b73ffffffffffffffffffffffffffffffffffffffff169084015260406118c08382016112c5565b73ffffffffffffffffffffffffffffffffffffffff16908401526118e582860161162b565b1515858401526118f682850161162b565b15158484015260a0828101359084015260c0808301359084015260e080830135908401526101008083013590840152610120928301929091019060010161187e565b5050848103602086015261194d81898b6117c4565b925050508281036040840152611964818587611813565b9998505050505050505050565b60006020828403121561198357600080fd5b815161198e8161161d565b9392505050565b6000828210156119a7576119a76116cd565b500390565b60006101008083850312156119c057600080fd5b83601f8401126119cf57600080fd5b60405181810181811067ffffffffffffffff821117156119f1576119f16115c4565b604052908301908085831115611a0657600080fd5b845b83811015611a29578051611a1b816112a0565b825260209182019101611a08565b509095945050505050565b600060208284031215611a4657600080fd5b815161198e816112a0565b600060208284031215611a6357600080fd5b5051919050565b60028110611aa1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60e081526000855160c060e08401528051806101a085015260005b81811015611ade5760208184018101516101c0878401015201611ac0565b81811115611af15760006101c083870101525b506020880151610100850152604088015161012085015260608801519150611b1d610140850183611a6a565b608088015173ffffffffffffffffffffffffffffffffffffffff8116610160860152915060a088015173ffffffffffffffffffffffffffffffffffffffff81166101808601529150601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683016101c0019150611be09050602083018673ffffffffffffffffffffffffffffffffffffffff808251168352806020830151166020840152506040810151151560408301526060810151151560608301525050565b60a082019390935260c0015292915050565b600060208284031215611c0457600080fd5b815163ffffffff8116811461198e57600080fd5b600060208284031215611c2a57600080fd5b81516fffffffffffffffffffffffffffffffff8116811461198e57600080fdfea2646970667358221220018f62ef6d4c62eca7b618107b573c899ebaf437f9e9096c59fb630610909f9764736f6c634300080d0033"

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, i common.Address, s common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LenderBin), backend, i, s)
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

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Lender *LenderCaller) Illuminate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "illuminate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Lender *LenderSession) Illuminate() (common.Address, error) {
	return _Lender.Contract.Illuminate(&_Lender.CallOpts)
}

// Illuminate is a free data retrieval call binding the contract method 0xcbf7e670.
//
// Solidity: function illuminate() view returns(address)
func (_Lender *LenderCallerSession) Illuminate() (common.Address, error) {
	return _Lender.Contract.Illuminate(&_Lender.CallOpts)
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

// Mint is a paid mutator transaction binding the contract method 0xdc4c7ca9.
//
// Solidity: function mint(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Lender *LenderTransactor) Mint(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "mint", p, u, m, a)
}

// Mint is a paid mutator transaction binding the contract method 0xdc4c7ca9.
//
// Solidity: function mint(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Lender *LenderSession) Mint(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Mint(&_Lender.TransactOpts, p, u, m, a)
}

// Mint is a paid mutator transaction binding the contract method 0xdc4c7ca9.
//
// Solidity: function mint(uint8 p, address u, uint256 m, uint256 a) returns(bool)
func (_Lender *LenderTransactorSession) Mint(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Mint(&_Lender.TransactOpts, p, u, m, a)
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

// LenderMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Lender contract.
type LenderMintIterator struct {
	Event *LenderMint // Event containing the contract specifics and raw log

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
func (it *LenderMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LenderMint)
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
		it.Event = new(LenderMint)
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
func (it *LenderMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LenderMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LenderMint represents a Mint event raised by the Lender contract.
type LenderMint struct {
	Principal  uint8
	Underlying common.Address
	Maturity   *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x309b03ba657e17f1beadbc6eb3c06ba79b38084eb8d0e5452cc222462a17f1f6.
//
// Solidity: event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Lender *LenderFilterer) FilterMint(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*LenderMintIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.FilterLogs(opts, "Mint", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &LenderMintIterator{contract: _Lender.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x309b03ba657e17f1beadbc6eb3c06ba79b38084eb8d0e5452cc222462a17f1f6.
//
// Solidity: event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Lender *LenderFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LenderMint, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Lender.contract.WatchLogs(opts, "Mint", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LenderMint)
				if err := _Lender.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x309b03ba657e17f1beadbc6eb3c06ba79b38084eb8d0e5452cc222462a17f1f6.
//
// Solidity: event Mint(uint8 principal, address indexed underlying, uint256 indexed maturity, uint256 amount)
func (_Lender *LenderFilterer) ParseMint(log types.Log) (*LenderMint, error) {
	event := new(LenderMint)
	if err := _Lender.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
