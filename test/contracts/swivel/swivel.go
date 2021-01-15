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

// SwivelABI is the input ABI used to generate the binding from.
const SwivelABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Initiate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agreementKey\",\"type\":\"bytes32\"}],\"name\":\"Release\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"agreements\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"floating\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"released\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"initialRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"release\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"o\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"},{\"internalType\":\"uint256[6]\",\"name\":\"p\",\"type\":\"uint256[6]\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fillFixed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwivelBin is the compiled bytecode used for deploying new contracts.
var SwivelBin = "0x60806040523480156200001157600080fd5b50604051620021f9380380620021f9833981810160405281019062000037919062000197565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161462000074578162000076565b305b9050620000fb6040518060400160405280600e81526020017f53776976656c2046696e616e63650000000000000000000000000000000000008152506040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525085846200010a60201b62000d5e1760201c565b6000819055505050506200024a565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b6000815190506200017a8162000216565b92915050565b600081519050620001918162000230565b92915050565b60008060408385031215620001ab57600080fd5b6000620001bb8582860162000180565b9250506020620001ce8582860162000169565b9150509250929050565b6000620001e582620001ec565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200022181620001d8565b81146200022d57600080fd5b50565b6200023b816200020c565b81146200024757600080fd5b50565b611f9f806200025a6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063288cdc911461005c5780632ac126221461008c578063331132b5146100bc57806352a9674b146100ec578063be3b52ec1461010a575b600080fd5b6100766004803603810190610071919061153a565b610144565b6040516100839190611cf8565b60405180910390f35b6100a660048036038101906100a1919061153a565b61015c565b6040516100b39190611a69565b60405180910390f35b6100d660048036038101906100d19190611563565b61017c565b6040516100e39190611a69565b60405180910390f35b6100f4610c77565b6040516101019190611a84565b60405180910390f35b610124600480360381019061011f919061162e565b610c7d565b60405161013b9b9a9998979695949392919061195e565b60405180910390f35b60026020528060005260406000206000915090505481565b60016020528060005260406000206000915054906101000a900460ff1681565b6000801515600160008c815260200190815260200160002060009054906101000a900460ff161515146101e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101db90611bf8565b60405180910390fd5b428860046005811115610220577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610257577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020135101561029d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029490611bb8565b60405180910390fd5b6102c16102b96000546102b48d8d8a60008f610dbd565b6110eb565b85858561112c565b73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff161461032e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032590611bd8565b60405180910390fd5b600260008b815260200190815260200160002054886002600581111561037d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106103b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201356103c39190611e05565b871115610405576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fc90611c18565b60405180910390fd5b61040d6113d5565b670de0b6b3a76400008960026005811115610451577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610488577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020135670de0b6b3a76400008a6104a19190611dab565b6104ab9190611d7a565b8a600160058111156104e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061051d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002013561052c9190611dab565b6105369190611d7a565b8160e0018181525050878161010001818152505060008690508073ffffffffffffffffffffffffffffffffffffffff166323b872dd8c308560e001516040518463ffffffff1660e01b815260040161059093929190611a09565b602060405180830381600087803b1580156105aa57600080fd5b505af11580156105be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e29190611511565b610621576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061890611c78565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd33308561010001516040518463ffffffff1660e01b815260040161066393929190611a09565b602060405180830381600087803b15801561067d57600080fd5b505af1158015610691573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b59190611511565b6106f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106eb90611c58565b60405180910390fd5b60006107f7888c60026005811115610735577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061076c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358d600160058111156107ac577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106107e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201356107f29190611d24565b611241565b11610837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082e90611c98565b60405180910390fd5b600073822397d9a55d0fefd20f5c4bcab33c5f65bd28eb90508b836000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505033836020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505087836040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060008360600190151590811515815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561095157600080fd5b505af1158015610965573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610989919061166a565b8360a00181815250508a600060058111156109cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610a04577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358360c00181815250508a60036005811115610a4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610a84577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358361012001818152505042836101200151610aa49190611d24565b8361014001818152505082600360008f815260200190815260200160002060008b815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a0820151816003015560c0820151816004015560e082015181600501556101008201518160060155610120820151816007015561014082015181600801559050507fa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb8d8a604051610c5b929190611a9f565b60405180910390a1600193505050509998505050505050505050565b60005481565b6003602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160149054906101000a900460ff16908060020160159054906101000a900460ff1690806003015490806004015490806005015490806006015490806007015490806008015490508b565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b60007f982d366ee870e6c64d27e7b149dff6bf737fd1c909c2392b1b6dda92d31a24e260001b868686868660006005811115610e22577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610e59577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358760016005811115610e99577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610ed0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358860026005811115610f10577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610f47577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358960036005811115610f87577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110610fbe577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358a60046005811115610ffe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110611035577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201358b600580811115611074577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600681106110ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201356040516020016110ca9b9a99989796959493929190611ac8565b60405160208183030381529060405280519060200120905095945050505050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115611194576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118b90611cd8565b60405180910390fd5b601b8460ff1614806111a95750601c8460ff16145b6111e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111df90611c38565b60405180910390fd5b6001858585856040516000815260200160405260405161120b9493929190611b73565b6020604051602081039080840390855afa15801561122d573d6000803e3d6000fd5b505050602060405103519050949350505050565b6000808390508073ffffffffffffffffffffffffffffffffffffffff1663095ea7b373822397d9a55d0fefd20f5c4bcab33c5f65bd28eb856040518363ffffffff1660e01b8152600401611296929190611a40565b602060405180830381600087803b1580156112b057600080fd5b505af11580156112c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e89190611511565b611327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131e90611cb8565b60405180910390fd5b600073822397d9a55d0fefd20f5c4bcab33c5f65bd28eb90508073ffffffffffffffffffffffffffffffffffffffff1663a0712d68856040518263ffffffff1660e01b81526004016113799190611cf8565b602060405180830381600087803b15801561139357600080fd5b505af11580156113a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113cb919061166a565b9250505092915050565b604051806101600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581526020016000151581526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b60008135905061148481611ef6565b92915050565b6000819050826020600602820111156114a257600080fd5b92915050565b6000815190506114b781611f0d565b92915050565b6000813590506114cc81611f24565b92915050565b6000813590506114e181611f3b565b92915050565b6000815190506114f681611f3b565b92915050565b60008135905061150b81611f52565b92915050565b60006020828403121561152357600080fd5b6000611531848285016114a8565b91505092915050565b60006020828403121561154c57600080fd5b600061155a848285016114bd565b91505092915050565b60008060008060008060008060006101c08a8c03121561158257600080fd5b60006115908c828d016114bd565b99505060206115a18c828d01611475565b98505060406115b28c828d0161148a565b9750506101006115c48c828d016114d2565b9650506101206115d68c828d016114bd565b9550506101406115e88c828d01611475565b9450506101606115fa8c828d016114fc565b93505061018061160c8c828d016114bd565b9250506101a061161e8c828d016114bd565b9150509295985092959850929598565b6000806040838503121561164157600080fd5b600061164f858286016114bd565b9250506020611660858286016114bd565b9150509250929050565b60006020828403121561167c57600080fd5b600061168a848285016114e7565b91505092915050565b61169c81611e39565b82525050565b6116ab81611e4b565b82525050565b6116ba81611e57565b82525050565b60006116cd601183611d13565b91507f6f726465722068617320657870697265640000000000000000000000000000006000830152602082019050919050565b600061170d601183611d13565b91507f696e76616c6964207369676e61747572650000000000000000000000000000006000830152602082019050919050565b600061174d601883611d13565b91507f6f7264657220686173206265656e2063616e63656c6c656400000000000000006000830152602082019050919050565b600061178d601f83611d13565b91507f74616b657220616d6f756e74203e20617661696c61626c6520766f6c756d65006000830152602082019050919050565b60006117cd601b83611d13565b91507f696e76616c6964207369676e6174757265202276222076616c756500000000006000830152602082019050919050565b600061180d601a83611d13565b91507f7472616e736665722066726f6d2074616b6572206661696c65640000000000006000830152602082019050919050565b600061184d601a83611d13565b91507f7472616e736665722066726f6d206d616b6572206661696c65640000000000006000830152602082019050919050565b600061188d601583611d13565b91507f43546f6b656e206d696e74696e67206661696c656400000000000000000000006000830152602082019050919050565b60006118cd601a83611d13565b91507f756e6465726c79696e6720617070726f76616c206661696c65640000000000006000830152602082019050919050565b600061190d601b83611d13565b91507f696e76616c6964207369676e6174757265202273222076616c756500000000006000830152602082019050919050565b61194981611e81565b82525050565b61195881611e8b565b82525050565b600061016082019050611974600083018e611693565b611981602083018d611693565b61198e604083018c611693565b61199b606083018b6116a2565b6119a8608083018a6116a2565b6119b560a0830189611940565b6119c260c0830188611940565b6119cf60e0830187611940565b6119dd610100830186611940565b6119eb610120830185611940565b6119f9610140830184611940565b9c9b505050505050505050505050565b6000606082019050611a1e6000830186611693565b611a2b6020830185611693565b611a386040830184611940565b949350505050565b6000604082019050611a556000830185611693565b611a626020830184611940565b9392505050565b6000602082019050611a7e60008301846116a2565b92915050565b6000602082019050611a9960008301846116b1565b92915050565b6000604082019050611ab460008301856116b1565b611ac160208301846116b1565b9392505050565b600061016082019050611ade600083018e6116b1565b611aeb602083018d6116b1565b611af8604083018c611693565b611b05606083018b611693565b611b12608083018a6116a2565b611b1f60a0830189611940565b611b2c60c0830188611940565b611b3960e0830187611940565b611b47610100830186611940565b611b55610120830185611940565b611b63610140830184611940565b9c9b505050505050505050505050565b6000608082019050611b8860008301876116b1565b611b95602083018661194f565b611ba260408301856116b1565b611baf60608301846116b1565b95945050505050565b60006020820190508181036000830152611bd1816116c0565b9050919050565b60006020820190508181036000830152611bf181611700565b9050919050565b60006020820190508181036000830152611c1181611740565b9050919050565b60006020820190508181036000830152611c3181611780565b9050919050565b60006020820190508181036000830152611c51816117c0565b9050919050565b60006020820190508181036000830152611c7181611800565b9050919050565b60006020820190508181036000830152611c9181611840565b9050919050565b60006020820190508181036000830152611cb181611880565b9050919050565b60006020820190508181036000830152611cd1816118c0565b9050919050565b60006020820190508181036000830152611cf181611900565b9050919050565b6000602082019050611d0d6000830184611940565b92915050565b600082825260208201905092915050565b6000611d2f82611e81565b9150611d3a83611e81565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611d6f57611d6e611e98565b5b828201905092915050565b6000611d8582611e81565b9150611d9083611e81565b925082611da057611d9f611ec7565b5b828204905092915050565b6000611db682611e81565b9150611dc183611e81565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611dfa57611df9611e98565b5b828202905092915050565b6000611e1082611e81565b9150611e1b83611e81565b925082821015611e2e57611e2d611e98565b5b828203905092915050565b6000611e4482611e61565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b611eff81611e39565b8114611f0a57600080fd5b50565b611f1681611e4b565b8114611f2157600080fd5b50565b611f2d81611e57565b8114611f3857600080fd5b50565b611f4481611e81565b8114611f4f57600080fd5b50565b611f5b81611e8b565b8114611f6657600080fd5b5056fea264697066735822122025a28121ccebd3fef6af52f1b5ed3342d183246c8b2d7a73b5e9a1ee4bd5a76f64736f6c63430008000033"

// DeploySwivel deploys a new Ethereum contract, binding an instance of Swivel to it.
func DeploySwivel(auth *bind.TransactOpts, backend bind.ContractBackend, c *big.Int, v common.Address) (common.Address, *types.Transaction, *Swivel, error) {
	parsed, err := abi.JSON(strings.NewReader(SwivelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SwivelBin), backend, c, v)
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

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 initialRate, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCaller) Agreements(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	Maker       common.Address
	Taker       common.Address
	Underlying  common.Address
	Floating    bool
	Released    bool
	InitialRate *big.Int
	Rate        *big.Int
	Principal   *big.Int
	Interest    *big.Int
	Duration    *big.Int
	Release     *big.Int
}, error) {
	var out []interface{}
	err := _Swivel.contract.Call(opts, &out, "agreements", arg0, arg1)

	outstruct := new(struct {
		Maker       common.Address
		Taker       common.Address
		Underlying  common.Address
		Floating    bool
		Released    bool
		InitialRate *big.Int
		Rate        *big.Int
		Principal   *big.Int
		Interest    *big.Int
		Duration    *big.Int
		Release     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maker = out[0].(common.Address)
	outstruct.Taker = out[1].(common.Address)
	outstruct.Underlying = out[2].(common.Address)
	outstruct.Floating = out[3].(bool)
	outstruct.Released = out[4].(bool)
	outstruct.InitialRate = out[5].(*big.Int)
	outstruct.Rate = out[6].(*big.Int)
	outstruct.Principal = out[7].(*big.Int)
	outstruct.Interest = out[8].(*big.Int)
	outstruct.Duration = out[9].(*big.Int)
	outstruct.Release = out[10].(*big.Int)

	return *outstruct, err

}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 initialRate, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker       common.Address
	Taker       common.Address
	Underlying  common.Address
	Floating    bool
	Released    bool
	InitialRate *big.Int
	Rate        *big.Int
	Principal   *big.Int
	Interest    *big.Int
	Duration    *big.Int
	Release     *big.Int
}, error) {
	return _Swivel.Contract.Agreements(&_Swivel.CallOpts, arg0, arg1)
}

// Agreements is a free data retrieval call binding the contract method 0xbe3b52ec.
//
// Solidity: function agreements(bytes32 , bytes32 ) view returns(address maker, address taker, address underlying, bool floating, bool released, uint256 initialRate, uint256 rate, uint256 principal, uint256 interest, uint256 duration, uint256 release)
func (_Swivel *SwivelCallerSession) Agreements(arg0 [32]byte, arg1 [32]byte) (struct {
	Maker       common.Address
	Taker       common.Address
	Underlying  common.Address
	Floating    bool
	Released    bool
	InitialRate *big.Int
	Rate        *big.Int
	Principal   *big.Int
	Interest    *big.Int
	Duration    *big.Int
	Release     *big.Int
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

// FillFixed is a paid mutator transaction binding the contract method 0x331132b5.
//
// Solidity: function fillFixed(bytes32 o, address m, uint256[6] p, uint256 a, bytes32 k, address u, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Swivel *SwivelTransactor) FillFixed(opts *bind.TransactOpts, o [32]byte, m common.Address, p [6]*big.Int, a *big.Int, k [32]byte, u common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Swivel.contract.Transact(opts, "fillFixed", o, m, p, a, k, u, v, r, s)
}

// FillFixed is a paid mutator transaction binding the contract method 0x331132b5.
//
// Solidity: function fillFixed(bytes32 o, address m, uint256[6] p, uint256 a, bytes32 k, address u, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Swivel *SwivelSession) FillFixed(o [32]byte, m common.Address, p [6]*big.Int, a *big.Int, k [32]byte, u common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, m, p, a, k, u, v, r, s)
}

// FillFixed is a paid mutator transaction binding the contract method 0x331132b5.
//
// Solidity: function fillFixed(bytes32 o, address m, uint256[6] p, uint256 a, bytes32 k, address u, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Swivel *SwivelTransactorSession) FillFixed(o [32]byte, m common.Address, p [6]*big.Int, a *big.Int, k [32]byte, u common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Swivel.Contract.FillFixed(&_Swivel.TransactOpts, o, m, p, a, k, u, v, r, s)
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
// Solidity: event Cancel(bytes32 key)
func (_Swivel *SwivelFilterer) FilterCancel(opts *bind.FilterOpts) (*SwivelCancelIterator, error) {

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return &SwivelCancelIterator{contract: _Swivel.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0xe8d9861dbc9c663ed3accd261bbe2fe01e0d3d9e5f51fa38523b265c7757a93a.
//
// Solidity: event Cancel(bytes32 key)
func (_Swivel *SwivelFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *SwivelCancel) (event.Subscription, error) {

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Cancel")
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
// Solidity: event Cancel(bytes32 key)
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
// Solidity: event Initiate(bytes32 orderKey, bytes32 agreementKey)
func (_Swivel *SwivelFilterer) FilterInitiate(opts *bind.FilterOpts) (*SwivelInitiateIterator, error) {

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Initiate")
	if err != nil {
		return nil, err
	}
	return &SwivelInitiateIterator{contract: _Swivel.contract, event: "Initiate", logs: logs, sub: sub}, nil
}

// WatchInitiate is a free log subscription operation binding the contract event 0xa93e646db02470b4aa881817da6191d55ffba8bd3377a40de3d62abb38fc7ceb.
//
// Solidity: event Initiate(bytes32 orderKey, bytes32 agreementKey)
func (_Swivel *SwivelFilterer) WatchInitiate(opts *bind.WatchOpts, sink chan<- *SwivelInitiate) (event.Subscription, error) {

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Initiate")
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
// Solidity: event Initiate(bytes32 orderKey, bytes32 agreementKey)
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
// Solidity: event Release(bytes32 orderKey, bytes32 agreementKey)
func (_Swivel *SwivelFilterer) FilterRelease(opts *bind.FilterOpts) (*SwivelReleaseIterator, error) {

	logs, sub, err := _Swivel.contract.FilterLogs(opts, "Release")
	if err != nil {
		return nil, err
	}
	return &SwivelReleaseIterator{contract: _Swivel.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0xa26be6031ace8c10be363a05e05ce508eae43c02ff806b3ff75af1d96dd90294.
//
// Solidity: event Release(bytes32 orderKey, bytes32 agreementKey)
func (_Swivel *SwivelFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *SwivelRelease) (event.Subscription, error) {

	logs, sub, err := _Swivel.contract.WatchLogs(opts, "Release")
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
// Solidity: event Release(bytes32 orderKey, bytes32 agreementKey)
func (_Swivel *SwivelFilterer) ParseRelease(log types.Log) (*SwivelRelease, error) {
	event := new(SwivelRelease)
	if err := _Swivel.contract.UnpackLog(event, "Release", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
