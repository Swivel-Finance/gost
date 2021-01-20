// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swivel

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

// HashOrder is an auto generated low-level Go binding around an user-defined struct.
type HashOrder struct {
	Key        [32]byte
	Maker      common.Address
	Underlying common.Address
	Floating   bool
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Expiry     *big.Int
	Nonce      *big.Int
}

// SigComponents is an auto generated low-level Go binding around an user-defined struct.
type SigComponents struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Initiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Release\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CTOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"agreements\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"released\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"release\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structHash.Order\",\"name\":\"o\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSig.Components\",\"name\":\"c\",\"type\":\"tuple\"}],\"name\":\"fillFixed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x60806040523480156200001157600080fd5b5060405162001fb338038062001fb383398181016040528101906200003791906200024b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620000aa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a190620002e3565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161462000127578162000129565b305b9050620001ae6040518060400160405280600e81526020017f53776976656c2046696e616e63650000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e300000000000000000000000000000000000000000000000000000008152508684620001be60201b62000b3f1760201c565b6001819055505050505062000388565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b6000815190506200022e8162000354565b92915050565b60008151905062000245816200036e565b92915050565b6000806000606084860312156200026157600080fd5b6000620002718682870162000234565b935050602062000284868287016200021d565b925050604062000297868287016200021d565b9150509250925092565b6000620002b0601f8362000305565b91507f636f6d706f756e6420746f6b656e2061646472657373207265717569726564006000830152602082019050919050565b60006020820190508181036000830152620002fe81620002a1565b9050919050565b600082825260208201905092915050565b600062000323826200032a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200035f8162000316565b81146200036b57600080fd5b50565b62000379816200034a565b81146200038557600080fd5b50565b611c1b80620003986000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806379d5e8671161005b57806379d5e86714610129578063a3f4df7e14610159578063be3b52ec14610177578063ffa1ad74146101b057610088565b8063288cdc911461008d5780632ac12622146100bd57806343295411146100ed57806352a9674b1461010b575b600080fd5b6100a760048036038101906100a29190611173565b6101ce565b6040516100b49190611925565b60405180910390f35b6100d760048036038101906100d29190611173565b6101e6565b6040516100e491906116ac565b60405180910390f35b6100f5610206565b6040516101029190611595565b60405180910390f35b61011361022a565b60405161012091906116c7565b60405180910390f35b610143600480360381019061013e91906111d8565b610230565b60405161015091906116ac565b60405180910390f35b6101616109f2565b60405161016e91906117c3565b60405180910390f35b610191600480360381019061018c919061119c565b610a2b565b6040516101a79a999897969594939291906115b0565b60405180910390f35b6101b8610b06565b6040516101c591906117c3565b60405180910390f35b60036020528060005260406000206000915090505481565b60026020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6000801515600260008760000135815260200190815260200160002060009054906101000a900460ff1615151461029c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029390611825565b60405180910390fd5b428560e0013510156102e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102da906117e5565b60405180910390fd5b6103006102fa6001546102f588610b9e565b610c52565b83610c93565b73ffffffffffffffffffffffffffffffffffffffff1685602001602081019061032991906110f8565b73ffffffffffffffffffffffffffffffffffffffff161461037f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037690611805565b60405180910390fd5b6003600086600001358152602001908152602001600020548560a001356103a69190611a3d565b8411156103e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103df90611845565b60405180910390fd5b6103f0610f95565b670de0b6b3a76400008660a00135670de0b6b3a76400008761041291906119e3565b61041c91906119b2565b876080013561042b91906119e3565b61043591906119b2565b8160c0018181525050848160e0018181525050600086604001602081019061045d91906110f8565b90508073ffffffffffffffffffffffffffffffffffffffff166323b872dd88602001602081019061048e91906110f8565b308560c001516040518463ffffffff1660e01b81526004016104b29392919061164c565b602060405180830381600087803b1580156104cc57600080fd5b505af11580156104e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610504919061114a565b610543576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053a906118a5565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd33308560e001516040518463ffffffff1660e01b81526004016105849392919061164c565b602060405180830381600087803b15801561059e57600080fd5b505af11580156105b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d6919061114a565b610615576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060c90611885565b60405180910390fd5b600061064688604001602081019061062d91906110f8565b8960a001358a60800135610641919061195c565b610de8565b11610686576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067d906118c5565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508760200160208101906106bf91906110f8565b836000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505033836020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505087604001602081019061074191906110f8565b836040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060008360600190151590811515815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156107d157600080fd5b505af11580156107e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610809919061123f565b8360a00181815250508760c001358361010001818152505042836101000151610832919061195c565b8361012001818152505082600460008a600001358152602001908152602001600020600088815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a0820151816003015560c0820151816004015560e08201518160050155610100820151816006015561012082015181600701559050508588600001357fa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb60405160405180910390a360019350505050949350505050565b6040518060400160405280600e81526020017f53776976656c2046696e616e636500000000000000000000000000000000000081525081565b6004602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060020160159054906101000a900460ff1690806003015490806004015490806005015490806006015490806007015490508a565b6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007ff01cc16c244dd394ae954a5b2cd48a4f17007f995776783399d1190f45ada62360001b8260000135836020016020810190610bdc91906110f8565b846040016020810190610bef91906110f8565b856060016020810190610c029190611121565b86608001358760a001358860c001358960e001358a6101000135604051602001610c359a999897969594939291906116e2565b604051602081830303815290604052805190602001209050919050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0826040013560001c1115610cff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf690611905565b60405180910390fd5b601b826000016020810190610d149190611268565b60ff161480610d385750601c826000016020810190610d339190611268565b60ff16145b610d77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6e90611865565b60405180910390fd5b600183836000016020810190610d8d9190611268565b8460200135856040013560405160008152602001604052604051610db4949392919061177e565b6020604051602081039080840390855afa158015610dd6573d6000803e3d6000fd5b50505060206040510351905092915050565b6000808390508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b360008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040518363ffffffff1660e01b8152600401610e49929190611683565b602060405180830381600087803b158015610e6357600080fd5b505af1158015610e77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9b919061114a565b610eda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed1906118e5565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663a0712d68856040518263ffffffff1660e01b8152600401610f399190611925565b602060405180830381600087803b158015610f5357600080fd5b505af1158015610f67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8b919061123f565b9250505092915050565b604051806101400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160001515815260200160001515815260200160008152602001600081526020016000815260200160008152602001600081525090565b60008135905061103d81611b72565b92915050565b60008135905061105281611b89565b92915050565b60008151905061106781611b89565b92915050565b60008135905061107c81611ba0565b92915050565b60006060828403121561109457600080fd5b81905092915050565b600061012082840312156110b057600080fd5b81905092915050565b6000813590506110c881611bb7565b92915050565b6000815190506110dd81611bb7565b92915050565b6000813590506110f281611bce565b92915050565b60006020828403121561110a57600080fd5b60006111188482850161102e565b91505092915050565b60006020828403121561113357600080fd5b600061114184828501611043565b91505092915050565b60006020828403121561115c57600080fd5b600061116a84828501611058565b91505092915050565b60006020828403121561118557600080fd5b60006111938482850161106d565b91505092915050565b600080604083850312156111af57600080fd5b60006111bd8582860161106d565b92505060206111ce8582860161106d565b9150509250929050565b6000806000806101c085870312156111ef57600080fd5b60006111fd8782880161109d565b94505061012061120f878288016110b9565b9350506101406112218782880161106d565b92505061016061123387828801611082565b91505092959194509250565b60006020828403121561125157600080fd5b600061125f848285016110ce565b91505092915050565b60006020828403121561127a57600080fd5b6000611288848285016110e3565b91505092915050565b61129a81611a71565b82525050565b6112a981611a83565b82525050565b6112b881611a8f565b82525050565b60006112c982611940565b6112d3818561194b565b93506112e3818560208601611ad0565b6112ec81611b61565b840191505092915050565b600061130460118361194b565b91507f6f726465722068617320657870697265640000000000000000000000000000006000830152602082019050919050565b600061134460118361194b565b91507f696e76616c6964207369676e61747572650000000000000000000000000000006000830152602082019050919050565b600061138460188361194b565b91507f6f7264657220686173206265656e2063616e63656c6c656400000000000000006000830152602082019050919050565b60006113c4601f8361194b565b91507f74616b657220616d6f756e74203e20617661696c61626c6520766f6c756d65006000830152602082019050919050565b6000611404601b8361194b565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b6000611444601a8361194b565b91507f7472616e736665722066726f6d2074616b6572206661696c65640000000000006000830152602082019050919050565b6000611484601a8361194b565b91507f7472616e736665722066726f6d206d616b6572206661696c65640000000000006000830152602082019050919050565b60006114c460158361194b565b91507f43546f6b656e206d696e74696e67206661696c656400000000000000000000006000830152602082019050919050565b6000611504601a8361194b565b91507f756e6465726c79696e6720617070726f76616c206661696c65640000000000006000830152602082019050919050565b6000611544601b8361194b565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b61158081611ab9565b82525050565b61158f81611ac3565b82525050565b60006020820190506115aa6000830184611291565b92915050565b6000610140820190506115c6600083018d611291565b6115d3602083018c611291565b6115e0604083018b611291565b6115ed606083018a6112a0565b6115fa60808301896112a0565b61160760a0830188611577565b61161460c0830187611577565b61162160e0830186611577565b61162f610100830185611577565b61163d610120830184611577565b9b9a5050505050505050505050565b60006060820190506116616000830186611291565b61166e6020830185611291565b61167b6040830184611577565b949350505050565b60006040820190506116986000830185611291565b6116a56020830184611577565b9392505050565b60006020820190506116c160008301846112a0565b92915050565b60006020820190506116dc60008301846112af565b92915050565b6000610140820190506116f8600083018d6112af565b611705602083018c6112af565b611712604083018b611291565b61171f606083018a611291565b61172c60808301896112a0565b61173960a0830188611577565b61174660c0830187611577565b61175360e0830186611577565b611761610100830185611577565b61176f610120830184611577565b9b9a5050505050505050505050565b600060808201905061179360008301876112af565b6117a06020830186611586565b6117ad60408301856112af565b6117ba60608301846112af565b95945050505050565b600060208201905081810360008301526117dd81846112be565b905092915050565b600060208201905081810360008301526117fe816112f7565b9050919050565b6000602082019050818103600083015261181e81611337565b9050919050565b6000602082019050818103600083015261183e81611377565b9050919050565b6000602082019050818103600083015261185e816113b7565b9050919050565b6000602082019050818103600083015261187e816113f7565b9050919050565b6000602082019050818103600083015261189e81611437565b9050919050565b600060208201905081810360008301526118be81611477565b9050919050565b600060208201905081810360008301526118de816114b7565b9050919050565b600060208201905081810360008301526118fe816114f7565b9050919050565b6000602082019050818103600083015261191e81611537565b9050919050565b600060208201905061193a6000830184611577565b92915050565b600081519050919050565b600082825260208201905092915050565b600061196782611ab9565b915061197283611ab9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156119a7576119a6611b03565b5b828201905092915050565b60006119bd82611ab9565b91506119c883611ab9565b9250826119d8576119d7611b32565b5b828204905092915050565b60006119ee82611ab9565b91506119f983611ab9565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611a3257611a31611b03565b5b828202905092915050565b6000611a4882611ab9565b9150611a5383611ab9565b925082821015611a6657611a65611b03565b5b828203905092915050565b6000611a7c82611a99565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611aee578082015181840152602081019050611ad3565b83811115611afd576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000601f19601f8301169050919050565b611b7b81611a71565b8114611b8657600080fd5b50565b611b9281611a83565b8114611b9d57600080fd5b50565b611ba981611a8f565b8114611bb457600080fd5b50565b611bc081611ab9565b8114611bcb57600080fd5b50565b611bd781611ac3565b8114611be257600080fd5b5056fea2646970667358221220c27c0578e79a75ab523b75058bc5a179ea238cf3de75921c687811ca9da1a3e764736f6c63430008000033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend, i *big.Int, c common.Address, v common.Address) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend, i, c, v)
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

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelCaller) CTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "CTOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelSession) CTOKEN() (common.Address, error) {
	return _Swivel.Contract.CTOKEN(&_Swivel.CallOpts)
}

// CTOKEN is a free data retrieval call binding the contract method 0x43295411.
//
// Solidity: function CTOKEN() view returns(address)
func (_Swivel *SwivelCallerSession) CTOKEN() (common.Address, error) {
	return _Swivel.Contract.CTOKEN(&_Swivel.CallOpts)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelCaller) DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelSession) DOMAIN() ([32]byte, error) {
	return _Swivel.Contract.DOMAIN(&_Swivel.CallOpts)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_Swivel *SwivelCallerSession) DOMAIN() ([32]byte, error) {
	return _Swivel.Contract.DOMAIN(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Swivel *SwivelCallerSession) NAME() (string, error) {
	return _Swivel.Contract.NAME(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Swivel *SwivelCallerSession) VERSION() (string, error) {
	return _Swivel.Contract.VERSION(&_Swivel.CallOpts)
}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCaller) Agreements(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "agreements", arg0, arg1)

	outstruct := new(struct {
		Maker      common.Address
		Taker      common.Address
		Underlying common.Address
		Floating   bool
		Released   bool
		Rate       *big.Int
		Principal  *big.Int
		Interest   *big.Int
		Duration   *big.Int
		Release    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maker = out[0].(common.Address)
	outstruct.Taker = out[1].(common.Address)
	outstruct.Underlying = out[2].(common.Address)
	outstruct.Floating = out[3].(bool)
	outstruct.Released = out[4].(bool)
	outstruct.Rate = out[5].(*big.Int)
	outstruct.Principal = out[6].(*big.Int)
	outstruct.Interest = out[7].(*big.Int)
	outstruct.Duration = out[8].(*big.Int)
	outstruct.Release = out[9].(*big.Int)

	return *outstruct, err

}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	return _Swivel.Contract.Agreements(&_Swivel.CallOpts, arg0, arg1)
}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCallerSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker      common.Address
	Taker      common.Address
	Underlying common.Address
	Floating   bool
	Released   bool
	Rate       *big.Int
	Principal  *big.Int
	Interest   *big.Int
	Duration   *big.Int
	Release    *big.Int
}, error) {
	return _Swivel.Contract.Agreements(&_Swivel.CallOpts, arg0, arg1)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "cancelled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled(bytes32 ) view returns(bool)
func (_Swivel *SwivelCallerSession) Cancelled(arg0 [32]byte) (bool, error) {
	return _Swivel.Contract.Cancelled(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "filled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled(bytes32 ) view returns(uint256)
func (_Swivel *SwivelCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Swivel.Contract.Filled(&_Swivel.CallOpts, arg0)
}

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactor) FillFixed(opts *bind.TransactOpts, o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "fillFixed", o, a, k, c)
}

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelSession) FillFixed(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, a, k, c)
}

// FillFixed is a paid mutator transaction binding the contract method 0x79d5e867.
//
// Solidity: function fillFixed((bytes32,address,address,bool,uint256,uint256,uint256,uint256,uint256) o, uint256 a, bytes32 k, (uint8,bytes32,bytes32) c) returns(bool)
func (_Swivel *SwivelTransactorSession) FillFixed(o HashOrder, a *big.Int, k [32]byte, c SigComponents) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, a, k, c)
}

// SwivelCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the Swivel contract.
type SwivelCancelIterator struct {
	Event *SwivelCancel // Event containing the contract specifics and raw log

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
func (it *SwivelCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelCancel)
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
		it.Event = new(SwivelCancel)
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
func (it *SwivelCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelCancel represents a Cancel event raised by the Swivel contract.
type SwivelCancel struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) FilterCancel(opts *bind.FilterOpts, key [][32]byte) (*SwivelCancelIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelCancelIterator{contract: _Swivel.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *SwivelCancel, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Cancel", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelCancel)
				if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// ParseCancel is a log parse operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 indexed key)
func (_Swivel *SwivelFilterer) ParseCancel(log types.Log) (*SwivelCancel, error) {
	event := new(SwivelCancel)
	if err := _Swivel.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelInitiateIterator is returned from FilterInitiate and is used to iterate over the raw logs and unpacked data for Initiate events raised by the Swivel contract.
type SwivelInitiateIterator struct {
	Event *SwivelInitiate // Event containing the contract specifics and raw log

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
func (it *SwivelInitiateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelInitiate)
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
		it.Event = new(SwivelInitiate)
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
func (it *SwivelInitiateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelInitiateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelInitiate represents a Initiate event raised by the Swivel contract.
type SwivelInitiate struct {
	OrderKey     [32]byte
	AgreementKey [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiate is a free log retrieval operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) FilterInitiate(opts *bind.FilterOpts, orderKey [][32]byte, agreementKey [][32]byte) (*SwivelInitiateIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Initiate", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelInitiateIterator{contract: _Swivel.contract, event: "Initiate", logs: logs, sub: sub}, nil
}

// WatchInitiate is a free log subscription operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) WatchInitiate(opts *bind.WatchOpts, sink chan<- *SwivelInitiate, orderKey [][32]byte, agreementKey [][32]byte) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Initiate", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelInitiate)
				if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
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

// ParseInitiate is a log parse operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) ParseInitiate(log types.Log) (*SwivelInitiate, error) {
	event := new(SwivelInitiate)
	if err := _Swivel.contract.UnpackLog(event, "Initiate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwivelReleaseIterator is returned from FilterRelease and is used to iterate over the raw logs and unpacked data for Release events raised by the Swivel contract.
type SwivelReleaseIterator struct {
	Event *SwivelRelease // Event containing the contract specifics and raw log

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
func (it *SwivelReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwivelRelease)
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
		it.Event = new(SwivelRelease)
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
func (it *SwivelReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwivelReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwivelRelease represents a Release event raised by the Swivel contract.
type SwivelRelease struct {
	OrderKey     [32]byte
	AgreementKey [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelease is a free log retrieval operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) FilterRelease(opts *bind.FilterOpts, orderKey [][32]byte, agreementKey [][32]byte) (*SwivelReleaseIterator, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Release", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return &SwivelReleaseIterator{contract: _Swivel.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *SwivelRelease, orderKey [][32]byte, agreementKey [][32]byte) (event.Subscription, error) {

	var orderKeyRule []interface{}
	for _, orderKeyItem := range orderKey {
		orderKeyRule = append(orderKeyRule, orderKeyItem)
	}
	var agreementKeyRule []interface{}
	for _, agreementKeyItem := range agreementKey {
		agreementKeyRule = append(agreementKeyRule, agreementKeyItem)
	}

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Release", orderKeyRule, agreementKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwivelRelease)
				if err := _Swivel.contract.UnpackLog(event, "Release", log); err != nil {
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

// ParseRelease is a log parse operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 indexed orderKey, bytes32 indexed agreementKey)
func (_Swivel *SwivelFilterer) ParseRelease(log types.Log) (*SwivelRelease, error) {
	event := new(SwivelRelease)
	if err := _Swivel.contract.UnpackLog(event, "Release", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
