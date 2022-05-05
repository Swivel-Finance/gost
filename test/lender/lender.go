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
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"su\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"tempusLend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x60806040523480156200001157600080fd5b5060405162002814380380620028148339810160408190526200003491620000af565b60008054336001600160a01b0319918216179091556001805482166001600160a01b0396871617905560028054821694861694909417909355600380548416928516929092179091556004805490921692169190911790556200010c565b80516001600160a01b0381168114620000aa57600080fd5b919050565b60008060008060808587031215620000c657600080fd5b620000d18562000092565b9350620000e16020860162000092565b9250620000f16040860162000092565b9150620001016060860162000092565b905092959194509250565b6126f8806200011c6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063cee5013811610076578063ea08c0311161005b578063ea08c031146101a7578063f1aa654e146101c7578063f851a440146101da57600080fd5b8063cee5013814610174578063d5907c361461019457600080fd5b80634135c9d1116100a75780634135c9d11461012e5780635b1e5fc2146101415780636d13582c1461015457600080fd5b806321401e9e146100c35780632e25d2a6146100e9575b600080fd5b6100d66100d1366004611b32565b6101fa565b6040519081526020015b60405180910390f35b6001546101099073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e0565b6100d661013c366004611bb5565b6105d7565b6100d661014f366004611ce3565b610ba9565b6003546101099073ffffffffffffffffffffffffffffffffffffffff1681565b6004546101099073ffffffffffffffffffffffffffffffffffffffff1681565b6100d66101a2366004611dbf565b610e92565b6002546101099073ffffffffffffffffffffffffffffffffffffffff1681565b6100d66101d5366004611e34565b61132a565b6000546101099073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89811660048301526024820189905260009283929116906317b3bba790604401610100604051808303816000875af1158015610278573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029c9190611f35565b8a60ff16600881106102b0576102b0611fbd565b602002015160045490915073ffffffffffffffffffffffffffffffffffffffff16896102de8133308c611891565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8516906370a08231906024016020604051808303816000875af115801561034d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103719190611fec565b6040517f4f255a1f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a811660048301528981166024830152604482018d905260016064830152608482018c905260a48201899052851690634f255a1f9060c4016020604051808303816000875af1158015610403573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104279190611fec565b6104319190612034565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8f81166004830152602482018f90529293509116906317b3bba790604401610100604051808303816000875af11580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d19190611f35565b600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810183905273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af115801561054c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105709190612059565b506040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260009283929116906317b3bba790604401610100604051808303816000875af1158015610655573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106799190611f35565b8860ff166008811061068d5761068d611fbd565b6020020151905060008190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af11580156106fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610721919061207d565b73ffffffffffffffffffffffffffffffffffffffff16146107a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f70656e646c6520756e6465726c79696e6720213d20756e6465726c79696e670060448201526064015b60405180910390fd5b868173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af11580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108159190611fec565b1461087c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f70656e646c65206d6174757269747920213d206d617475726974790000000000604482015260640161079a565b61088888333089611891565b60408051600280825260608201835260009260208301908036833701905050905088816000815181106108bd576108bd611fbd565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160018151811061090b5761090b611fbd565b73ffffffffffffffffffffffffffffffffffffffff92831660209182029290920101526003546040517f38ed173900000000000000000000000000000000000000000000000000000000815260009291909116906338ed17399061097b908b908b90879030908d9060040161209a565b6000604051808303816000875af115801561099a573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526109e09190810190612125565b6000815181106109f2576109f2611fbd565b60209081029190910101516001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d81166004830152602482018d9052929350600092909116906317b3bba790604401610100604051808303816000875af1158015610a7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa19190611f35565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810184905273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015610b1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b439190612059565b506040805160ff8e168152602081018490528b9173ffffffffffffffffffffffffffffffffffffffff8e16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a3509a9950505050505050505050565b60008080805b88811015610d675760008a8a83818110610bcb57610bcb611fbd565b90506101200201803603810190610be291906121d6565b90508c8160e0015114610c51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f73776976656c206d6174757269747920213d206d617475726974790000000000604482015260640161079a565b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1614610cea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f73776976656c20756e6465726c79696e6720213d20756e6465726c79696e6700604482015260640161079a565b888883818110610cfc57610cfc611fbd565b9050602002013584610d0e919061226d565b93508060a001518160c001518a8a85818110610d2c57610d2c611fbd565b90506020020135610d3d9190612285565b610d4791906122c2565b610d51908461226d565b9250508080610d5f906122fd565b915050610baf565b50610d748c333085611891565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f5890610dd4908c908c908c908c908c908c906004016123d9565b6020604051808303816000875af1158015610df3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e179190612059565b50610e2c8c8b610e278585612034565b61197b565b6040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b6000610ea088333087611891565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905260009216906317b3bba790604401610100604051808303816000875af1158015610f1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3f9190611f35565b8a60ff1660088110610f5357610f53611fbd565b602002015190508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610fbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe2919061207d565b73ffffffffffffffffffffffffffffffffffffffff1614611039576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161079a565b878173ffffffffffffffffffffffffffffffffffffffff1663aa082a9d6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015611087573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ab9190611fec565b146110ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161079a565b6040805160e081018252600060c08201818152825260208201899052918101879052606081018281526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152509050600060405180608001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152509050898b73ffffffffffffffffffffffffffffffffffffffff167f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d98e8c73ffffffffffffffffffffffffffffffffffffffff166369bbd8cd87878d8d6040518563ffffffff1660e01b8152600401611220949392919061251d565b6020604051808303816000875af115801561123f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112639190611fec565b6040805160ff909316835260208301919091520160405180910390a36040517f69bbd8cd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906369bbd8cd906112d790859085908b908b9060040161251d565b6020604051808303816000875af11580156112f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131a9190611fec565b9c9b505050505050505050505050565b6000808390508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635001f3b56040518163ffffffff1660e01b81526004016020604051808303816000875af1158015611394573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b8919061207d565b73ffffffffffffffffffffffffffffffffffffffff1614611435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7969656c64206261736520213d20756e6465726c79696e670000000000000000604482015260640161079a565b848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015611483573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a7919061266a565b63ffffffff1614611514576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7969656c64206d6174757269747920213d206d61747572697479000000000000604482015260640161079a565b853061152282338388611891565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c61154988611a48565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af11580156115b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d49190612690565b90506115e183888861197b565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af1158015611667573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061168b9190612690565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b160161180b576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af1158015611731573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117559190611f35565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af11580156117e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118089190612059565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af191505061190e81611aa2565b611974576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c6564000000000000000000000000604482015260640161079a565b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af19150506119dc81611aa2565b611a42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015260640161079a565b50505050565b60006fffffffffffffffffffffffffffffffff821115611a9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526000602482015260440161079a565b5090565b6000803d83611ab557806000803e806000fd5b8060208114611acd578015611ade5760009250611ae3565b816000803e60005115159250611ae3565b600192505b50909392505050565b803560ff81168114611afd57600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114611b2457600080fd5b50565b8035611afd81611b02565b600080600080600080600080610100898b031215611b4f57600080fd5b611b5889611aec565b97506020890135611b6881611b02565b965060408901359550606089013594506080890135935060a0890135611b8d81611b02565b925060c0890135611b9d81611b02565b8092505060e089013590509295985092959890939650565b60008060008060008060c08789031215611bce57600080fd5b611bd787611aec565b95506020870135611be781611b02565b95989597505050506040840135936060810135936080820135935060a0909101359150565b60008083601f840112611c1e57600080fd5b50813567ffffffffffffffff811115611c3657600080fd5b60208301915083602061012083028501011115611c5257600080fd5b9250929050565b60008083601f840112611c6b57600080fd5b50813567ffffffffffffffff811115611c8357600080fd5b6020830191508360208260051b8501011115611c5257600080fd5b60008083601f840112611cb057600080fd5b50813567ffffffffffffffff811115611cc857600080fd5b602083019150836020606083028501011115611c5257600080fd5b60008060008060008060008060008060e08b8d031215611d0257600080fd5b611d0b8b611aec565b995060208b0135611d1b81611b02565b985060408b0135975060608b0135611d3281611b02565b965060808b013567ffffffffffffffff80821115611d4f57600080fd5b611d5b8e838f01611c0c565b909850965060a08d0135915080821115611d7457600080fd5b611d808e838f01611c59565b909650945060c08d0135915080821115611d9957600080fd5b50611da68d828e01611c9e565b915080935050809150509295989b9194979a5092959850565b600080600080600080600080610100898b031215611ddc57600080fd5b611de589611aec565b97506020890135611df581611b02565b9650604089013595506060890135611e0c81611b02565b979a969950949760808101359660a0820135965060c0820135955060e0909101359350915050565b600080600080600060a08688031215611e4c57600080fd5b611e5586611aec565b94506020860135611e6581611b02565b9350604086013592506060860135611e7c81611b02565b949793965091946080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715611ee057611ee0611e8d565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611f2d57611f2d611e8d565b604052919050565b6000610100808385031215611f4957600080fd5b83601f840112611f5857600080fd5b60405181810181811067ffffffffffffffff82111715611f7a57611f7a611e8d565b604052908301908085831115611f8f57600080fd5b845b83811015611fb2578051611fa481611b02565b825260209182019101611f91565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611ffe57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561204657612046612005565b500390565b8015158114611b2457600080fd5b60006020828403121561206b57600080fd5b81516120768161204b565b9392505050565b60006020828403121561208f57600080fd5b815161207681611b02565b600060a082018783526020878185015260a0604085015281875180845260c086019150828901935060005b818110156120f757845173ffffffffffffffffffffffffffffffffffffffff16835293830193918301916001016120c5565b505073ffffffffffffffffffffffffffffffffffffffff969096166060850152505050608001529392505050565b6000602080838503121561213857600080fd5b825167ffffffffffffffff8082111561215057600080fd5b818501915085601f83011261216457600080fd5b81518181111561217657612176611e8d565b8060051b9150612187848301611ee6565b81815291830184019184810190888411156121a157600080fd5b938501935b838510156121bf578451825293850193908501906121a6565b98975050505050505050565b8035611afd8161204b565b600061012082840312156121e957600080fd5b6121f1611ebc565b8235815261220160208401611b27565b602082015261221260408401611b27565b6040820152612223606084016121cb565b6060820152612234608084016121cb565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b6000821982111561228057612280612005565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156122bd576122bd612005565b500290565b6000826122f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361232e5761232e612005565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561236757600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b858110156123ce5760ff6123a783611aec565b16875281830135838801526040808301359088015260609687019690910190600101612394565b509495945050505050565b606080825281810187905260009060808084018a845b8b8110156124a9578135835260208083013561240a81611b02565b73ffffffffffffffffffffffffffffffffffffffff16908401526040612431838201611b27565b73ffffffffffffffffffffffffffffffffffffffff16908401526124568286016121cb565b1515858401526124678285016121cb565b15158484015260a0828101359084015260c0808301359084015260e08083013590840152610100808301359084015261012092830192909101906001016123ef565b505084810360208601526124be81898b612335565b9250505082810360408401526124d5818587612384565b9998505050505050505050565b60028110612519577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60e081526000855160c060e08401528051806101a085015260005b818110156125565760208184018101516101c0878401015201612538565b818111156125695760006101c083870101525b5060208801516101008501526040880151610120850152606088015191506125956101408501836124e2565b608088015173ffffffffffffffffffffffffffffffffffffffff8116610160860152915060a088015173ffffffffffffffffffffffffffffffffffffffff81166101808601529150601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683016101c00191506126589050602083018673ffffffffffffffffffffffffffffffffffffffff808251168352806020830151166020840152506040810151151560408301526060810151151560608301525050565b60a082019390935260c0015292915050565b60006020828403121561267c57600080fd5b815163ffffffff8116811461207657600080fd5b6000602082840312156126a257600080fd5b81516fffffffffffffffffffffffffffffffff8116811461207657600080fdfea2646970667358221220953d6b5a3274f411621b5742043e836b248b1aba3cc3f3e519db85ce7635820764736f6c634300080d0033"

// DeployLender deploys a new Ethereum contract, binding an instance of Lender to it.
func DeployLender(auth *bind.TransactOpts, backend bind.ContractBackend, m common.Address, s common.Address, su common.Address, t common.Address) (common.Address, *types.Transaction, *Lender, error) {
	parsed, err := abi.JSON(strings.NewReader(LenderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LenderBin), backend, m, s, su, t)
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

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Lender *LenderCaller) TempusRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lender.contract.Call(opts, &out, "tempusRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Lender *LenderSession) TempusRouter() (common.Address, error) {
	return _Lender.Contract.TempusRouter(&_Lender.CallOpts)
}

// TempusRouter is a free data retrieval call binding the contract method 0xcee50138.
//
// Solidity: function tempusRouter() view returns(address)
func (_Lender *LenderCallerSession) TempusRouter() (common.Address, error) {
	return _Lender.Contract.TempusRouter(&_Lender.CallOpts)
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

// TempusLend is a paid mutator transaction binding the contract method 0x21401e9e.
//
// Solidity: function tempusLend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderTransactor) TempusLend(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "tempusLend", p, u, m, a, r, x, t, d)
}

// TempusLend is a paid mutator transaction binding the contract method 0x21401e9e.
//
// Solidity: function tempusLend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderSession) TempusLend(p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.TempusLend(&_Lender.TransactOpts, p, u, m, a, r, x, t, d)
}

// TempusLend is a paid mutator transaction binding the contract method 0x21401e9e.
//
// Solidity: function tempusLend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderTransactorSession) TempusLend(p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.TempusLend(&_Lender.TransactOpts, p, u, m, a, r, x, t, d)
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
