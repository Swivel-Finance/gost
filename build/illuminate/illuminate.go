// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package illuminate

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

// IlluminateABI is the input ABI used to generate the binding from.
const IlluminateABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address[7]\",\"name\":\"t\",\"type\":\"address[7]\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IlluminateBin is the compiled bytecode used for deploying new contracts.
var IlluminateBin = "0x60a060405234801561001057600080fd5b5060405161236b38038061236b83398101604081905261002f91610052565b600180546001600160a01b031916331790556001600160a01b0316608052610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b6080516122c86100a360003960008181609e015261039401526122c86000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c8063125cf47f14620000575780632ba29d3814620000985780635f1f94e614620000c0578063f851a44014620000e8575b600080fd5b6200006e6200006836600462000654565b62000109565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6200006e7f000000000000000000000000000000000000000000000000000000000000000081565b620000d7620000d136600462000717565b62000152565b60405190151581526020016200008f565b6001546200006e9073ffffffffffffffffffffffffffffffffffffffff1681565b600060205282600052604060002060205281600052604060002081600881106200013257600080fd5b015473ffffffffffffffffffffffffffffffffffffffff16925083915050565b60015460009073ffffffffffffffffffffffffffffffffffffffff16338114620001dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a81166000908152602081815260408083208d8452909152902054161562000277576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d61726b657420657869737473000000000000000000000000000000000000006044820152606401620001d4565b60008a8a89898989896040516200028e9062000583565b620002a09796959493929190620008b3565b604051809103906000f080158015620002bd573d6000803e3d6000fd5b50604080516101008101825273ffffffffffffffffffffffffffffffffffffffff83811682528c5181166020808401919091528d0151811682840152918c015182166060808301919091528c015182166080808301919091528c0151821660a0808301919091528c0151821660c0808301919091528c015190911660e08201529091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60005b60088160ff161015620003cf57620003ba838260ff16600881106200038d576200038d6200083b565b60200201517f00000000000000000000000000000000000000000000000000000000000000008462000463565b80620003c68162000918565b91505062000364565b5073ffffffffffffffffffffffffffffffffffffffff8d166000908152602081815260408083208f845290915290206200040c9083600862000591565b506040518c9073ffffffffffffffffffffffffffffffffffffffff8f16907f0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee90600090a35060019c9b505050505050505050505050565b60006040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050620004c68162000534565b6200052e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f617070726f7665206661696c65640000000000000000000000000000000000006044820152606401620001d4565b50505050565b6000803d836200054857806000803e806000fd5b8060208114620005635780156200057557600092506200057a565b816000803e600051151592506200057a565b600192505b50909392505050565b611933806200096083390190565b826008810192821562000601579160200282015b828111156200060157825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190620005a5565b506200060f92915062000613565b5090565b5b808211156200060f576000815560010162000614565b803573ffffffffffffffffffffffffffffffffffffffff811681146200064f57600080fd5b919050565b6000806000606084860312156200066a57600080fd5b62000675846200062a565b95602085013595506040909401359392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083601f840112620006cc57600080fd5b50813567ffffffffffffffff811115620006e557600080fd5b602083019150836020828501011115620006fe57600080fd5b9250929050565b803560ff811681146200064f57600080fd5b600080600080600080600080610180898b0312156200073557600080fd5b62000740896200062a565b97506020808a013597508a605f8b01126200075a57600080fd5b60405160e0810167ffffffffffffffff82821081831117156200078157620007816200068a565b816040528291506101208d018e8111156200079b57600080fd5b60408e015b81811015620007c257620007b4816200062a565b8452928501928501620007a0565b50839a50803594505080841115620007d957600080fd5b620007e78e858f01620006b9565b90995097506101408d01359350889250808411156200080557600080fd5b505050620008168b828c01620006b9565b90945092506200082c90506101608a0162000705565b90509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a060408201526000620008eb60a0830187896200086a565b8281036060840152620009008186886200086a565b91505060ff8316608083015298975050505050505050565b600060ff821660ff810362000956577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019291505056fe6101006040523480156200001257600080fd5b5060405162001933380380620019338339810160408190526200003591620002a9565b82828282828282600490805190602001906200005392919062000136565b5081516200006990600590602085019062000136565b5080600260006101000a81548160ff021916908360ff160217905550505050620000ba83604051806040016040528060018152602001603160f81b8152504630620000df60201b62000ab61760201c565b60805250503360a052505050506001600160a01b039190911660c05260e05262000395565b8351602094850120835193850193909320604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815295860194909452928401929092526060830152608082015260a0902090565b828054620001449062000359565b90600052602060002090601f016020900481019282620001685760008555620001b3565b82601f106200018357805160ff1916838001178555620001b3565b82800160010185558215620001b3579182015b82811115620001b357825182559160200191906001019062000196565b50620001c1929150620001c5565b5090565b5b80821115620001c15760008155600101620001c6565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200020457600080fd5b81516001600160401b0380821115620002215762000221620001dc565b604051601f8301601f19908116603f011681019082821181831017156200024c576200024c620001dc565b816040528381526020925086838588010111156200026957600080fd5b600091505b838210156200028d57858201830151818301840152908201906200026e565b838211156200029f5760008385830101525b9695505050505050565b600080600080600060a08688031215620002c257600080fd5b85516001600160a01b0381168114620002da57600080fd5b6020870151604088015191965094506001600160401b0380821115620002ff57600080fd5b6200030d89838a01620001f2565b945060608801519150808211156200032457600080fd5b506200033388828901620001f2565b925050608086015160ff811681146200034b57600080fd5b809150509295509295909350565b600181811c908216806200036e57607f821691505b6020821081036200038f57634e487b7160e01b600052602260045260246000fd5b50919050565b60805160a05160c05160e05161154f620003e460003960006101b30152600061023201526000818161039701528181610597015261065a0152600081816103150152610931015261154f6000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c806370a08231116100cd578063a9059cbb11610081578063d505accf11610066578063d505accf14610337578063dd62ed3e1461034c578063f851a4401461039257600080fd5b8063a9059cbb146102fd578063c2fb26a61461031057600080fd5b806395d89b41116100b257806395d89b41146102cf5780639dc29fac146102d7578063a457c2d7146102ea57600080fd5b806370a08231146102795780637ecebe00146102af57600080fd5b806323b872dd116101245780633950935111610109578063395093511461020757806340c10f191461021a5780636f307dc31461022d57600080fd5b806323b872dd146101d5578063313ce567146101e857600080fd5b806306fdde0314610156578063095ea7b31461017457806318160ddd14610197578063204f83f9146101ae575b600080fd5b61015e6103b9565b60405161016b9190611266565b60405180910390f35b610187610182366004611302565b610447565b604051901515815260200161016b565b6101a060035481565b60405190815260200161016b565b6101a07f000000000000000000000000000000000000000000000000000000000000000081565b6101876101e336600461132c565b61045d565b6002546101f59060ff1681565b60405160ff909116815260200161016b565b610187610215366004611302565b61054f565b610187610228366004611302565b610593565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016b565b6101a0610287366004611368565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101a06102bd366004611368565b60066020526000908152604090205481565b61015e610649565b6101876102e5366004611302565b610656565b6101876102f8366004611302565b610702565b61018761030b366004611302565b6107d1565b6101a07f000000000000000000000000000000000000000000000000000000000000000081565b61034a61034536600461138a565b6107de565b005b6101a061035a3660046113fd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6102547f000000000000000000000000000000000000000000000000000000000000000081565b600480546103c690611430565b80601f01602080910402602001604051908101604052809291908181526020018280546103f290611430565b801561043f5780601f106104145761010080835404028352916020019161043f565b820191906000526020600020905b81548152906001019060200180831161042257829003601f168201915b505050505081565b6000610454338484610b0d565b50600192915050565b600061046a848484610cc2565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260016020908152604080832033845290915290205482811015610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f6572633230207472616e7366657220616d6f756e74206578636565647320616c60448201527f6c6f77616e63650000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610544853361053f86856114b2565b610b0d565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909161045491859061053f9086906114c9565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f73656e646572206d7573742062652061646d696e0000000000000000000000006044820152606401610527565b61063f8484610f7e565b5060019392505050565b600580546103c690611430565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff8216146106f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f73656e646572206d7573742062652061646d696e0000000000000000000000006044820152606401610527565b61063f848461109e565b33600090815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152812054828110156107c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f65726332302064656372656173656420616c6c6f77616e63652062656c6f772060448201527f7a65726f000000000000000000000000000000000000000000000000000000006064820152608401610527565b61063f338561053f86856114b2565b6000610454338484610cc2565b42841015610848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f65726332363132206578706972656420646561646c696e6500000000000000006044820152606401610527565b73ffffffffffffffffffffffffffffffffffffffff871660009081526006602052604081208054610909918a918a918a919086610884836114e1565b90915550604080517f80772249b4aef1688b30651778f4249b05cb73b517d98482439b9d8999b3060260208083019190915273ffffffffffffffffffffffffffffffffffffffff96871682840152949095166060860152608085019290925260a084015260c08084018a90528151808503909101815260e09093019052815191012090565b6040517f190100000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600282015260228101829052604290209091506000906040805160008082526020820180845284905260ff89169282019290925260608101879052608081018690529192509060019060a0016020604051602081039080840390855afa1580156109be573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610a3957508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610a9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6572633236313220696e76616c6964207369676e6174757265000000000000006044820152606401610527565b610aaa8a8a8a610b0d565b50505050505050505050565b8351602094850120835193850193909320604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815295860194909452928401929092526060830152608082015260a0902090565b73ffffffffffffffffffffffffffffffffffffffff8316610bb0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f657263323020617070726f76652066726f6d20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610527565b73ffffffffffffffffffffffffffffffffffffffff8216610c53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f657263323020617070726f766520746f20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610527565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610d64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6572633230207472616e736665722066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610527565b73ffffffffffffffffffffffffffffffffffffffff8216610e07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f6572633230207472616e7366657220746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610527565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ebd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f6572633230207472616e7366657220616d6f756e74206578636565647320626160448201527f6c616e63650000000000000000000000000000000000000000000000000000006064820152608401610527565b610ec782826114b2565b73ffffffffffffffffffffffffffffffffffffffff8086166000908152602081905260408082209390935590851681529081208054849290610f0a9084906114c9565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f7091815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8216610ffb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f6572633230206d696e7420746f20746865207a65726f206164647265737300006044820152606401610527565b806003600082825461100d91906114c9565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040812080548392906110479084906114c9565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821661111b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f6572633230206275726e2066726f6d20746865207a65726f20616464726573736044820152606401610527565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156111d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f6572633230206275726e20616d6f756e7420657863656564732062616c616e6360448201527f65000000000000000000000000000000000000000000000000000000000000006064820152608401610527565b6111db82826114b2565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260208190526040812091909155600380548492906112169084906114b2565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610cb5565b600060208083528351808285015260005b8181101561129357858101830151858201604001528201611277565b818111156112a5576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146112fd57600080fd5b919050565b6000806040838503121561131557600080fd5b61131e836112d9565b946020939093013593505050565b60008060006060848603121561134157600080fd5b61134a846112d9565b9250611358602085016112d9565b9150604084013590509250925092565b60006020828403121561137a57600080fd5b611383826112d9565b9392505050565b600080600080600080600060e0888a0312156113a557600080fd5b6113ae886112d9565b96506113bc602089016112d9565b95506040880135945060608801359350608088013560ff811681146113e057600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561141057600080fd5b611419836112d9565b9150611427602084016112d9565b90509250929050565b600181811c9082168061144457607f821691505b60208210810361147d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156114c4576114c4611483565b500390565b600082198211156114dc576114dc611483565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361151257611512611483565b506001019056fea2646970667358221220850441112195bd30946b863989c92ccac5ccfa6eb0c62d4735fa76b2a62abae964736f6c634300080d0033a2646970667358221220f02967ccca5d12a85bcb0b2fb95122f4bae74903854229b495f309ba1ba2561764736f6c634300080d0033"

// DeployIlluminate deploys a new Ethereum contract, binding an instance of Illuminate to it.
func DeployIlluminate(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Illuminate, error) {
	parsed, err := abi.JSON(strings.NewReader(IlluminateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IlluminateBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// Illuminate is an auto generated Go binding around an Ethereum contract.
type Illuminate struct {
	IlluminateCaller     // Read-only binding to the contract
	IlluminateTransactor // Write-only binding to the contract
	IlluminateFilterer   // Log filterer for contract events
}

// IlluminateCaller is an auto generated read-only Go binding around an Ethereum contract.
type IlluminateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IlluminateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IlluminateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlluminateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IlluminateSession struct {
	Contract     *Illuminate       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IlluminateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IlluminateCallerSession struct {
	Contract *IlluminateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IlluminateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IlluminateTransactorSession struct {
	Contract     *IlluminateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IlluminateRaw is an auto generated low-level Go binding around an Ethereum contract.
type IlluminateRaw struct {
	Contract *Illuminate // Generic contract binding to access the raw methods on
}

// IlluminateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IlluminateCallerRaw struct {
	Contract *IlluminateCaller // Generic read-only contract binding to access the raw methods on
}

// IlluminateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IlluminateTransactorRaw struct {
	Contract *IlluminateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIlluminate creates a new instance of Illuminate, bound to a specific deployed contract.
func NewIlluminate(address common.Address, backend bind.ContractBackend) (*Illuminate, error) {
	contract, err := bindIlluminate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Illuminate{IlluminateCaller: IlluminateCaller{contract: contract}, IlluminateTransactor: IlluminateTransactor{contract: contract}, IlluminateFilterer: IlluminateFilterer{contract: contract}}, nil
}

// NewIlluminateCaller creates a new read-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateCaller(address common.Address, caller bind.ContractCaller) (*IlluminateCaller, error) {
	contract, err := bindIlluminate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateCaller{contract: contract}, nil
}

// NewIlluminateTransactor creates a new write-only instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateTransactor(address common.Address, transactor bind.ContractTransactor) (*IlluminateTransactor, error) {
	contract, err := bindIlluminate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IlluminateTransactor{contract: contract}, nil
}

// NewIlluminateFilterer creates a new log filterer instance of Illuminate, bound to a specific deployed contract.
func NewIlluminateFilterer(address common.Address, filterer bind.ContractFilterer) (*IlluminateFilterer, error) {
	contract, err := bindIlluminate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IlluminateFilterer{contract: contract}, nil
}

// bindIlluminate binds a generic wrapper to an already deployed contract.
func bindIlluminate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IlluminateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.IlluminateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.IlluminateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Illuminate *IlluminateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Illuminate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Illuminate *IlluminateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Illuminate *IlluminateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Illuminate.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateSession) Admin() (common.Address, error) {
	return _Illuminate.Contract.Admin(&_Illuminate.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Illuminate *IlluminateCallerSession) Admin() (common.Address, error) {
	return _Illuminate.Contract.Admin(&_Illuminate.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateCaller) Markets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "markets", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.CallOpts, arg0, arg1, arg2)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_Illuminate *IlluminateCallerSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _Illuminate.Contract.Markets(&_Illuminate.CallOpts, arg0, arg1, arg2)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateCaller) Redeemer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Illuminate.contract.Call(opts, &out, "redeemer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateSession) Redeemer() (common.Address, error) {
	return _Illuminate.Contract.Redeemer(&_Illuminate.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_Illuminate *IlluminateCallerSession) Redeemer() (common.Address, error) {
	return _Illuminate.Contract.Redeemer(&_Illuminate.CallOpts)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateTransactor) CreateMarket(opts *bind.TransactOpts, u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.contract.Transact(opts, "createMarket", u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateSession) CreateMarket(u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.Contract.CreateMarket(&_Illuminate.TransactOpts, u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x5f1f94e6.
//
// Solidity: function createMarket(address u, uint256 m, address[7] t, string n, string s, uint8 d) returns(bool)
func (_Illuminate *IlluminateTransactorSession) CreateMarket(u common.Address, m *big.Int, t [7]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Illuminate.Contract.CreateMarket(&_Illuminate.TransactOpts, u, m, t, n, s, d)
}

// IlluminateCreateMarketIterator is returned from FilterCreateMarket and is used to iterate over the raw logs and unpacked data for CreateMarket events raised by the Illuminate contract.
type IlluminateCreateMarketIterator struct {
	Event *IlluminateCreateMarket // Event containing the contract specifics and raw log

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
func (it *IlluminateCreateMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IlluminateCreateMarket)
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
		it.Event = new(IlluminateCreateMarket)
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
func (it *IlluminateCreateMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IlluminateCreateMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IlluminateCreateMarket represents a CreateMarket event raised by the Illuminate contract.
type IlluminateCreateMarket struct {
	Underlying common.Address
	Maturity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateMarket is a free log retrieval operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) FilterCreateMarket(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*IlluminateCreateMarketIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Illuminate.contract.FilterLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &IlluminateCreateMarketIterator{contract: _Illuminate.contract, event: "CreateMarket", logs: logs, sub: sub}, nil
}

// WatchCreateMarket is a free log subscription operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) WatchCreateMarket(opts *bind.WatchOpts, sink chan<- *IlluminateCreateMarket, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _Illuminate.contract.WatchLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IlluminateCreateMarket)
				if err := _Illuminate.contract.UnpackLog(event, "CreateMarket", log); err != nil {
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

// ParseCreateMarket is a log parse operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_Illuminate *IlluminateFilterer) ParseCreateMarket(log types.Log) (*IlluminateCreateMarket, error) {
	event := new(IlluminateCreateMarket)
	if err := _Illuminate.contract.UnpackLog(event, "CreateMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
