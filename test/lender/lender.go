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
const LenderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"su\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"principal\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returned\",\"type\":\"uint256\"}],\"name\":\"Lend\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"vault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structSwivel.Order[]\",\"name\":\"o\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"a\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSwivel.Components[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"e\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"y\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"x\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivelAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tempusRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LenderBin is the compiled bytecode used for deploying new contracts.
var LenderBin = "0x60806040523480156200001157600080fd5b5060405162002748380380620027488339810160408190526200003491620000af565b60008054336001600160a01b0319918216179091556001805482166001600160a01b0396871617905560028054821694861694909417909355600380548416928516929092179091556004805490921692169190911790556200010c565b80516001600160a01b0381168114620000aa57600080fd5b919050565b60008060008060808587031215620000c657600080fd5b620000d18562000092565b9350620000e16020860162000092565b9250620000f16040860162000092565b9150620001016060860162000092565b905092959194509250565b61262c806200011c6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063d5907c3611610076578063f1aa654e1161005b578063f1aa654e146101b4578063f3cf26f3146101c7578063f851a440146101da57600080fd5b8063d5907c3614610181578063ea08c0311461019457600080fd5b80635b1e5fc2116100a75780635b1e5fc21461012e5780636d13582c14610141578063cee501381461016157600080fd5b80632e25d2a6146100c35780634135c9d11461010d575b600080fd5b6001546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012061011b366004611a66565b6101fa565b604051908152602001610104565b61012061013c366004611b94565b6107cc565b6003546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6004546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b61012061018f366004611c70565b610ab5565b6002546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6101206101c2366004611ce5565b610f4d565b6101206101d5366004611d3e565b6114b4565b6000546100e39073ffffffffffffffffffffffffffffffffffffffff1681565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820187905260009283929116906317b3bba790604401610100604051808303816000875af1158015610278573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029c9190611e69565b8860ff16600881106102b0576102b0611ef1565b6020020151905060008190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610320573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103449190611f20565b73ffffffffffffffffffffffffffffffffffffffff16146103c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f70656e646c6520756e6465726c79696e6720213d20756e6465726c79696e670060448201526064015b60405180910390fd5b868173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610414573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104389190611f44565b1461049f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f70656e646c65206d6174757269747920213d206d61747572697479000000000060448201526064016103bd565b6104ab883330896117c5565b60408051600280825260608201835260009260208301908036833701905050905088816000815181106104e0576104e0611ef1565b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160018151811061052e5761052e611ef1565b73ffffffffffffffffffffffffffffffffffffffff92831660209182029290920101526003546040517f38ed173900000000000000000000000000000000000000000000000000000000815260009291909116906338ed17399061059e908b908b90879030908d90600401611f5d565b6000604051808303816000875af11580156105bd573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526106039190810190611fe8565b60008151811061061557610615611ef1565b60209081029190910101516001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8d81166004830152602482018d9052929350600092909116906317b3bba790604401610100604051808303816000875af11580156106a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c49190611e69565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810184905273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015610742573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610766919061209c565b506040805160ff8e168152602081018490528b9173ffffffffffffffffffffffffffffffffffffffff8e16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a3509a9950505050505050505050565b60008080805b8881101561098a5760008a8a838181106107ee576107ee611ef1565b9050610120020180360381019061080591906120c4565b90508c8160e0015114610874576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f73776976656c206d6174757269747920213d206d61747572697479000000000060448201526064016103bd565b8d73ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff161461090d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f73776976656c20756e6465726c79696e6720213d20756e6465726c79696e670060448201526064016103bd565b88888381811061091f5761091f611ef1565b9050602002013584610931919061218a565b93508060a001518160c001518a8a8581811061094f5761094f611ef1565b9050602002013561096091906121a2565b61096a91906121df565b610974908461218a565b92505080806109829061221a565b9150506107d2565b506109978c3330856117c5565b6002546040517fd2144f5800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063d2144f58906109f7908c908c908c908c908c908c906004016122f6565b6020604051808303816000875af1158015610a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3a919061209c565b50610a4f8c8b610a4a85856123ff565b6118af565b6040805160ff8f168152602081018390528c9173ffffffffffffffffffffffffffffffffffffffff8f16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39c9b505050505050505050505050565b6000610ac3883330876117c5565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905260009216906317b3bba790604401610100604051808303816000875af1158015610b3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b629190611e69565b8a60ff1660088110610b7657610b76611ef1565b602002015190508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610be1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c059190611f20565b73ffffffffffffffffffffffffffffffffffffffff1614610c5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016103bd565b878173ffffffffffffffffffffffffffffffffffffffff1663aa082a9d6040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610caa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cce9190611f44565b14610d0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016103bd565b6040805160e081018252600060c08201818152825260208201899052918101879052606081018281526020018b73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152509050600060405180608001604052803073ffffffffffffffffffffffffffffffffffffffff1681526020013073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152509050898b73ffffffffffffffffffffffffffffffffffffffff167f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d98e8c73ffffffffffffffffffffffffffffffffffffffff166369bbd8cd87878d8d6040518563ffffffff1660e01b8152600401610e439493929190612451565b6020604051808303816000875af1158015610e62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e869190611f44565b6040805160ff909316835260208301919091520160405180910390a36040517f69bbd8cd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906369bbd8cd90610efa90859085908b908b90600401612451565b6020604051808303816000875af1158015610f19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3d9190611f44565b9c9b505050505050505050505050565b6000808390508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16635001f3b56040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610fb7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fdb9190611f20565b73ffffffffffffffffffffffffffffffffffffffff1614611058576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7969656c64206261736520213d20756e6465726c79696e67000000000000000060448201526064016103bd565b848173ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b81526004016020604051808303816000875af11580156110a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ca919061259e565b63ffffffff1614611137576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7969656c64206d6174757269747920213d206d6174757269747900000000000060448201526064016103bd565b8530611145823383886117c5565b60008373ffffffffffffffffffffffffffffffffffffffff166313e7bc8c61116c8861197c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526fffffffffffffffffffffffffffffffff90911660048201526024016020604051808303816000875af11580156111d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f791906125c4565b90506112048388886118af565b6040517fbcc1694f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301526fffffffffffffffffffffffffffffffff8316602483015285169063bcc1694f906044016020604051808303816000875af115801561128a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ae91906125c4565b507ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8b160161142e576001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af1158015611354573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113789190611e69565b905080600060200201516040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526fffffffffffffffffffffffffffffffff8416602482015273ffffffffffffffffffffffffffffffffffffffff909116906340c10f19906044016020604051808303816000875af1158015611407573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142b919061209c565b50505b6040805160ff8c1681526fffffffffffffffffffffffffffffffff83166020820152899173ffffffffffffffffffffffffffffffffffffffff8c16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a36fffffffffffffffffffffffffffffffff169998505050505050505050565b6000876114c38133308a6117c5565b6001546040517f17b3bba700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152602482018b905260009216906317b3bba790604401610100604051808303816000875af115801561153e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115629190611e69565b600060200201516040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915060009073ffffffffffffffffffffffffffffffffffffffff8316906370a08231906024016020604051808303816000875af11580156115db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ff9190611f44565b600480546040517f4f255a1f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b8116938201939093528983166024820152604481018d905260016064820152608481018c905260a48101899052911690634f255a1f9060c4016020604051808303816000875af1158015611697573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bb9190611f44565b6116c591906123ff565b6040517f40c10f190000000000000000000000000000000000000000000000000000000081523360048201526024810182905290915073ffffffffffffffffffffffffffffffffffffffff8316906340c10f19906044016020604051808303816000875af115801561173b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175f919061209c565b506040805160ff8e168152602081018390528b9173ffffffffffffffffffffffffffffffffffffffff8e16917f26e45da718e624ee4c0a2c979b1436c3b050d3ef1197fdb210e20002900457d9910160405180910390a39b9a5050505050505050505050565b60006040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015273ffffffffffffffffffffffffffffffffffffffff8416602482015282604482015260008060648360008a5af1915050611842816119d6565b6118a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7472616e736665722066726f6d206661696c656400000000000000000000000060448201526064016103bd565b5050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050611910816119d6565b611976576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c6564000000000000000000000000000000000060448201526064016103bd565b50505050565b60006fffffffffffffffffffffffffffffffff8211156119d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044016103bd565b5090565b6000803d836119e957806000803e806000fd5b8060208114611a01578015611a125760009250611a17565b816000803e60005115159250611a17565b600192505b50909392505050565b803560ff81168114611a3157600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114611a5857600080fd5b50565b8035611a3181611a36565b60008060008060008060c08789031215611a7f57600080fd5b611a8887611a20565b95506020870135611a9881611a36565b95989597505050506040840135936060810135936080820135935060a0909101359150565b60008083601f840112611acf57600080fd5b50813567ffffffffffffffff811115611ae757600080fd5b60208301915083602061012083028501011115611b0357600080fd5b9250929050565b60008083601f840112611b1c57600080fd5b50813567ffffffffffffffff811115611b3457600080fd5b6020830191508360208260051b8501011115611b0357600080fd5b60008083601f840112611b6157600080fd5b50813567ffffffffffffffff811115611b7957600080fd5b602083019150836020606083028501011115611b0357600080fd5b60008060008060008060008060008060e08b8d031215611bb357600080fd5b611bbc8b611a20565b995060208b0135611bcc81611a36565b985060408b0135975060608b0135611be381611a36565b965060808b013567ffffffffffffffff80821115611c0057600080fd5b611c0c8e838f01611abd565b909850965060a08d0135915080821115611c2557600080fd5b611c318e838f01611b0a565b909650945060c08d0135915080821115611c4a57600080fd5b50611c578d828e01611b4f565b915080935050809150509295989b9194979a5092959850565b600080600080600080600080610100898b031215611c8d57600080fd5b611c9689611a20565b97506020890135611ca681611a36565b9650604089013595506060890135611cbd81611a36565b979a969950949760808101359660a0820135965060c0820135955060e0909101359350915050565b600080600080600060a08688031215611cfd57600080fd5b611d0686611a20565b94506020860135611d1681611a36565b9350604086013592506060860135611d2d81611a36565b949793965091946080013592915050565b600080600080600080600080610100898b031215611d5b57600080fd5b611d6489611a20565b97506020890135611d7481611a36565b965060408901359550606089013594506080890135935060a0890135611d9981611a36565b925060c0890135611da981611a36565b8092505060e089013590509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715611e1457611e14611dc1565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611e6157611e61611dc1565b604052919050565b6000610100808385031215611e7d57600080fd5b83601f840112611e8c57600080fd5b60405181810181811067ffffffffffffffff82111715611eae57611eae611dc1565b604052908301908085831115611ec357600080fd5b845b83811015611ee6578051611ed881611a36565b825260209182019101611ec5565b509095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215611f3257600080fd5b8151611f3d81611a36565b9392505050565b600060208284031215611f5657600080fd5b5051919050565b600060a082018783526020878185015260a0604085015281875180845260c086019150828901935060005b81811015611fba57845173ffffffffffffffffffffffffffffffffffffffff1683529383019391830191600101611f88565b505073ffffffffffffffffffffffffffffffffffffffff969096166060850152505050608001529392505050565b60006020808385031215611ffb57600080fd5b825167ffffffffffffffff8082111561201357600080fd5b818501915085601f83011261202757600080fd5b81518181111561203957612039611dc1565b8060051b915061204a848301611e1a565b818152918301840191848101908884111561206457600080fd5b938501935b8385101561208257845182529385019390850190612069565b98975050505050505050565b8015158114611a5857600080fd5b6000602082840312156120ae57600080fd5b8151611f3d8161208e565b8035611a318161208e565b600061012082840312156120d757600080fd5b6120df611df0565b823581526120ef60208401611a5b565b602082015261210060408401611a5b565b6040820152612111606084016120b9565b6060820152612122608084016120b9565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152508091505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561219d5761219d61215b565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156121da576121da61215b565b500290565b600082612215577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361224b5761224b61215b565b5060010190565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561228457600080fd5b8260051b8083602087013760009401602001938452509192915050565b8183526000602080850194508260005b858110156122eb5760ff6122c483611a20565b168752818301358388015260408083013590880152606096870196909101906001016122b1565b509495945050505050565b606080825281810187905260009060808084018a845b8b8110156123c6578135835260208083013561232781611a36565b73ffffffffffffffffffffffffffffffffffffffff1690840152604061234e838201611a5b565b73ffffffffffffffffffffffffffffffffffffffff16908401526123738286016120b9565b1515858401526123848285016120b9565b15158484015260a0828101359084015260c0808301359084015260e080830135908401526101008083013590840152610120928301929091019060010161230c565b505084810360208601526123db81898b612252565b9250505082810360408401526123f28185876122a1565b9998505050505050505050565b6000828210156124115761241161215b565b500390565b6002811061244d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60e081526000855160c060e08401528051806101a085015260005b8181101561248a5760208184018101516101c087840101520161246c565b8181111561249d5760006101c083870101525b5060208801516101008501526040880151610120850152606088015191506124c9610140850183612416565b608088015173ffffffffffffffffffffffffffffffffffffffff8116610160860152915060a088015173ffffffffffffffffffffffffffffffffffffffff81166101808601529150601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683016101c001915061258c9050602083018673ffffffffffffffffffffffffffffffffffffffff808251168352806020830151166020840152506040810151151560408301526060810151151560608301525050565b60a082019390935260c0015292915050565b6000602082840312156125b057600080fd5b815163ffffffff81168114611f3d57600080fd5b6000602082840312156125d657600080fd5b81516fffffffffffffffffffffffffffffffff81168114611f3d57600080fdfea264697066735822122005e3d77f27738c98130b90ae5aab2d1463bee1cc98c90dd924709d558c92ac5f64736f6c634300080d0033"

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

// Lend3 is a paid mutator transaction binding the contract method 0xf3cf26f3.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderTransactor) Lend3(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.contract.Transact(opts, "lend3", p, u, m, a, r, x, t, d)
}

// Lend3 is a paid mutator transaction binding the contract method 0xf3cf26f3.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderSession) Lend3(p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend3(&_Lender.TransactOpts, p, u, m, a, r, x, t, d)
}

// Lend3 is a paid mutator transaction binding the contract method 0xf3cf26f3.
//
// Solidity: function lend(uint8 p, address u, uint256 m, uint256 a, uint256 r, address x, address t, uint256 d) returns(uint256)
func (_Lender *LenderTransactorSession) Lend3(p uint8, u common.Address, m *big.Int, a *big.Int, r *big.Int, x common.Address, t common.Address, d *big.Int) (*types.Transaction, error) {
	return _Lender.Contract.Lend3(&_Lender.TransactOpts, p, u, m, a, r, x, t, d)
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
