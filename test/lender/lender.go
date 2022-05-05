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
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"su\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x60806040523480156200001157600080fd5b50604051620023293803806200232983398101604081905262000034916200009f565b600080546001600160a01b03199081163317909155600180546001600160a01b03958616908316179055600280549385169382169390931790925560038054919093169116179055620000e9565b80516001600160a01b03811681146200009a57600080fd5b919050565b600080600060608486031215620000b557600080fd5b620000c08462000082565b9250620000d06020850162000082565b9150620000e06040850162000082565b90509250925092565b61223080620000f96000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063d5907c361161005b578063d5907c361461012b578063ea08c0311461013e578063f1aa654e1461015e578063f851a4401461017157600080fd5b80632e25d2a61461008d5780634135c9d1146100d75780635b1e5fc2146100f85780636d13582c1461010b575b600080fd5b6001546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100ea6100e53660046116ed565b610191565b6040519081526020016100ce565b6100ea61010636600461181b565b610764565b6003546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6100ea6101393660046118f7565b610a4d565b6002546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6100ea61016c36600461196c565b610ee5565b6000546100ad9073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260009283929116906317b3bba790604401610100604051808303816000875af115801561020f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102339190611a6d565b8860ff166008811061024757610247611af5565b6020020151905060008190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af11580156102b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102db9190611b24565b73ffffffffffffffffffffffffffffffffffffffff161461035d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f70656e646c6520756e6465726c79696e6720213d20756e6465726c79696e670060448201526064015b60405180910390fd5b868173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af11580156103ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cf9190611b48565b14610436576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f70656e646c65206d6174757269747920213d206d6174757269747900000000006044820152606401610354565b6104428833308961144c565b604080516002808252606082018352600092602083019080368337019050509050888160008151811061047757610477611af5565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082816001815181106104c5576104c5611af5565b73ffffffffffffffffffffffffffffffffffffffff92831660209182029290920101526003546040517f38ed173900000000000000000000000000000000000000000000000000000000815291169060009082906338ed173990610535908c908c90889030908e90600401611b61565b6000604051808303816000875af1158015610554573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261059a9190810190611bec565b6000815181106105ac576105ac611af5565b60209081029190910101516001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8e81166004830152602482018e9052929350600092909116906317b3bba790604401610100604051808303816000875af1158015610637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065b9190611a6d565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810184905273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af11580156106d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106fd9190611ca0565b506040805160ff8f168152602081018490528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a3509b9a5050505050505050505050565b60008080805b888110156109225760008a8a8381811061078657610786611af5565b9050610120020180360381019061079d9190611cc8565b90508c8160e001511461080c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f73776976656c206d6174757269747920213d206d6174757269747900000000006044820152606401610354565b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16146108a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f73776976656c20756e6465726c79696e6720213d20756e6465726c79696e67006044820152606401610354565b8888838181106108b7576108b7611af5565b90506020020135846108c99190611d8e565b93508060a001518160c001518a8a858181106108e7576108e7611af5565b905060200201356108f89190611da6565b6109029190611de3565b61090c9084611d8e565b925050808061091a90611e1e565b91505061076a565b5061092f8c33308561144c565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f589061098f908c908c908c908c908c908c90600401611efa565b6020604051808303816000875af11580156109ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d29190611ca0565b506109e78c8b6109e28585612003565b611536565b6040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b6000610a5b8833308761144c565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905260009216906317b3bba790604401610100604051808303816000875af1158015610ad6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afa9190611a6d565b8a60ff1660088110610b0e57610b0e611af5565b602002015190508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610b79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9d9190611b24565b73ffffffffffffffffffffffffffffffffffffffff1614610bf4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610354565b878173ffffffffffffffffffffffffffffffffffffffff1663aa082a9d6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610c42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c669190611b48565b14610ca7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610354565b6040805160e081018252600060c08201818152825260208201899052918101879052606081018281526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152509050600060405180608001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152509050898b73ffffffffffffffffffffffffffffffffffffffff167f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d98e8c73ffffffffffffffffffffffffffffffffffffffff166369bbd8cd87878d8d6040518563ffffffff1660e01b8152600401610ddb9493929190612055565b6020604051808303816000875af1158015610dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1e9190611b48565b6040805160ff909316835260208301919091520160405180910390a36040517f69bbd8cd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906369bbd8cd90610e9290859085908b908b90600401612055565b6020604051808303816000875af1158015610eb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed59190611b48565b9c9b505050505050505050505050565b6000808390508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635001f3b56040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610f4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f739190611b24565b73ffffffffffffffffffffffffffffffffffffffff1614610ff0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7969656c64206261736520213d20756e6465726c79696e6700000000000000006044820152606401610354565b848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af115801561103e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061106291906121a2565b63ffffffff16146110cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7969656c64206d6174757269747920213d206d617475726974790000000000006044820152606401610354565b85306110dd8233838861144c565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c61110488611603565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af115801561116b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118f91906121c8565b905061119c838888611536565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af1158015611222573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124691906121c8565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b16016113c6576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af11580156112ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113109190611a6d565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af115801561139f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c39190611ca0565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af19150506114c98161165d565b61152f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c65640000000000000000000000006044820152606401610354565b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af19150506115978161165d565b6115fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c656400000000000000000000000000000000006044820152606401610354565b50505050565b60006fffffffffffffffffffffffffffffffff821115611659576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610354565b5090565b6000803d8361167057806000803e806000fd5b8060208114611688578015611699576000925061169e565b816000803e6000511515925061169e565b600192505b50909392505050565b803560ff811681146116b857600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff811681146116df57600080fd5b50565b80356116b8816116bd565b60008060008060008060c0878903121561170657600080fd5b61170f876116a7565b9550602087013561171f816116bd565b95989597505050506040840135936060810135936080820135935060a0909101359150565b60008083601f84011261175657600080fd5b50813567ffffffffffffffff81111561176e57600080fd5b6020830191508360206101208302850101111561178a57600080fd5b9250929050565b60008083601f8401126117a357600080fd5b50813567ffffffffffffffff8111156117bb57600080fd5b6020830191508360208260051b850101111561178a57600080fd5b60008083601f8401126117e857600080fd5b50813567ffffffffffffffff81111561180057600080fd5b60208301915083602060608302850101111561178a57600080fd5b60008060008060008060008060008060e08b8d03121561183a57600080fd5b6118438b6116a7565b995060208b0135611853816116bd565b985060408b0135975060608b013561186a816116bd565b965060808b013567ffffffffffffffff8082111561188757600080fd5b6118938e838f01611744565b909850965060a08d01359150808211156118ac57600080fd5b6118b88e838f01611791565b909650945060c08d01359150808211156118d157600080fd5b506118de8d828e016117d6565b915080935050809150509295989b9194979a5092959850565b600080600080600080600080610100898b03121561191457600080fd5b61191d896116a7565b9750602089013561192d816116bd565b9650604089013595506060890135611944816116bd565b979a969950949760808101359660a0820135965060c0820135955060e0909101359350915050565b600080600080600060a0868803121561198457600080fd5b61198d866116a7565b9450602086013561199d816116bd565b93506040860135925060608601356119b4816116bd565b949793965091946080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715611a1857611a186119c5565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611a6557611a656119c5565b604052919050565b6000610100808385031215611a8157600080fd5b83601f840112611a9057600080fd5b60405181810181811067ffffffffffffffff82111715611ab257611ab26119c5565b604052908301908085831115611ac757600080fd5b845b83811015611aea578051611adc816116bd565b825260209182019101611ac9565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611b3657600080fd5b8151611b41816116bd565b9392505050565b600060208284031215611b5a57600080fd5b5051919050565b600060a082018783526020878185015260a0604085015281875180845260c086019150828901935060005b81811015611bbe57845173ffffffffffffffffffffffffffffffffffffffff1683529383019391830191600101611b8c565b505073ffffffffffffffffffffffffffffffffffffffff969096166060850152505050608001529392505050565b60006020808385031215611bff57600080fd5b825167ffffffffffffffff80821115611c1757600080fd5b818501915085601f830112611c2b57600080fd5b815181811115611c3d57611c3d6119c5565b8060051b9150611c4e848301611a1e565b8181529183018401918481019088841115611c6857600080fd5b938501935b83851015611c8657845182529385019390850190611c6d565b98975050505050505050565b80151581146116df57600080fd5b600060208284031215611cb257600080fd5b8151611b4181611c92565b80356116b881611c92565b60006101208284031215611cdb57600080fd5b611ce36119f4565b82358152611cf3602084016116e2565b6020820152611d04604084016116e2565b6040820152611d1560608401611cbd565b6060820152611d2660808401611cbd565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611da157611da1611d5f565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611dde57611dde611d5f565b500290565b600082611e19577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e4f57611e4f611d5f565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611e8857600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b85811015611eef5760ff611ec8836116a7565b16875281830135838801526040808301359088015260609687019690910190600101611eb5565b509495945050505050565b606080825281810187905260009060808084018a845b8b811015611fca5781358352602080830135611f2b816116bd565b73ffffffffffffffffffffffffffffffffffffffff16908401526040611f528382016116e2565b73ffffffffffffffffffffffffffffffffffffffff1690840152611f77828601611cbd565b151585840152611f88828501611cbd565b15158484015260a0828101359084015260c0808301359084015260e0808301359084015261010080830135908401526101209283019290910190600101611f10565b50508481036020860152611fdf81898b611e56565b925050508281036040840152611ff6818587611ea5565b9998505050505050505050565b60008282101561201557612015611d5f565b500390565b60028110612051577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60e081526000855160c060e08401528051806101a085015260005b8181101561208e5760208184018101516101c0878401015201612070565b818111156120a15760006101c083870101525b5060208801516101008501526040880151610120850152606088015191506120cd61014085018361201a565b608088015173ffffffffffffffffffffffffffffffffffffffff8116610160860152915060a088015173ffffffffffffffffffffffffffffffffffffffff81166101808601529150601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683016101c00191506121909050602083018673ffffffffffffffffffffffffffffffffffffffff808251168352806020830151166020840152506040810151151560408301526060810151151560608301525050565b60a082019390935260c0015292915050565b6000602082840312156121b457600080fd5b815163ffffffff81168114611b4157600080fd5b6000602082840312156121da57600080fd5b81516fffffffffffffffffffffffffffffffff81168114611b4157600080fdfea264697066735822122098ad129bf865b7ea88adbb31b2a18482a015540e2729102c5d70047ae4e48cc164736f6c634300080d0033"

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, s common.Address, su common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LenderBin), backend, m, s, su)
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

// SushiRouter is a free data retrieval call binding the contract method 0x6d13582c.
//
// Solidity: function sushiRouter() view returns(address)
func (_Lender *LenderCaller) SushiRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "sushiRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SushiRouter is a free data retrieval call binding the contract method 0x6d13582c.
//
// Solidity: function sushiRouter() view returns(address)
func (_Lender *LenderSession) SushiRouter() (common.Address, error) {
	return _Lender.Contract.SushiRouter(&_Lender.CallOpts)
}

// SushiRouter is a free data retrieval call binding the contract method 0x6d13582c.
//
// Solidity: function sushiRouter() view returns(address)
func (_Lender *LenderCallerSession) SushiRouter() (common.Address, error) {
	return _Lender.Contract.SushiRouter(&_Lender.CallOpts)
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

// Lend is a paid mutator transaction binding the contract method 0x4135c9d1.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 mb, uint256 d) returns(uint256)
func (_Lender *LenderTransactor) Lend(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int, mb *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend", p, u, m, a, mb, d)
}

// Lend is a paid mutator transaction binding the contract method 0x4135c9d1.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 mb, uint256 d) returns(uint256)
func (_Lender *LenderSession) Lend(p uint8, u common.Address, m *big.Int, a *big.Int, mb *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, a, mb, d)
}

// Lend is a paid mutator transaction binding the contract method 0x4135c9d1.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 mb, uint256 d) returns(uint256)
func (_Lender *LenderTransactorSession) Lend(p uint8, u common.Address, m *big.Int, a *big.Int, mb *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend(&_Lender.TransactOpts, p, u, m, a, mb, d)
}

// Lend0 is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactor) Lend0(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend0", p, u, m, y, o, a, s)
}

// Lend0 is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderSession) Lend0(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend0 is a paid mutator transaction binding the contract method 0x5b1e5fc2.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, (bytes32,address,address,bool,bool,uint256,uint256,uint256,uint256)[] o, uint256[] a, (uint8,bytes32,bytes32)[] s) returns(uint256)
func (_Lender *LenderTransactorSession) Lend0(p uint8, u common.Address, m *big.Int, y common.Address, o []SwivelOrder, a []*big.Int, s []SwivelComponents) (*types.Transaction, error) {
	return _Lender.Contract.Lend0(&_Lender.TransactOpts, p, u, m, y, o, a, s)
}

// Lend1 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderTransactor) Lend1(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend1", p, u, m, e, i, a, r, d)
}

// Lend1 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderSession) Lend1(p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend1(&_Lender.TransactOpts, p, u, m, e, i, a, r, d)
}

// Lend1 is a paid mutator transaction binding the contract method 0xd5907c36.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address e, bytes32 i, uint256 a, uint256 r, uint256 d) returns(uint256)
func (_Lender *LenderTransactorSession) Lend1(p uint8, u common.Address, m *big.Int, e common.Address, i [32]byte, a *big.Int, r *big.Int, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend1(&_Lender.TransactOpts, p, u, m, e, i, a, r, d)
}

// Lend2 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactor) Lend2(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend2", p, u, m, y, a)
}

// Lend2 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderSession) Lend2(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend2(&_Lender.TransactOpts, p, u, m, y, a)
}

// Lend2 is a paid mutator transaction binding the contract method 0xf1aa654e.
//
// Solidity: function lend(uint8 p, address u, uint256 m, address y, uint256 a) returns(uint256)
func (_Lender *LenderTransactorSession) Lend2(p uint8, u common.Address, m *big.Int, y common.Address, a *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend2(&_Lender.TransactOpts, p, u, m, y, a)
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
