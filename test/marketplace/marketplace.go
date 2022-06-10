// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MarketPlace

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

// MarketPlaceABI is the input ABI used to generate the binding from.
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"buyPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"returned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"buyUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"returned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address[8]\",\"name\":\"t\",\"type\":\"address[8]\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"sellPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"returned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"sellUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"returned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x60a060405234801561001057600080fd5b50604051611ab7380380611ab783398101604081905261002f91610052565b600280546001600160a01b031916331790556001600160a01b0316608052610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611a146100a36000396000818160fc0152610b5f0152611a146000f3fe60806040523480156200001157600080fd5b5060043610620000b05760003560e01c8063ad31b198116200007f578063c9ac53b11162000062578063c9ac53b1146200019c578063cef26d4314620001b3578063f851a44014620001db57600080fd5b8063ad31b198146200016e578063b892d15a146200018557600080fd5b8063125cf47f14620000b55780632ba29d3814620000f657806352289a15146200011e5780635b0cbb051462000157575b600080fd5b620000cc620000c636600462000ef8565b620001fc565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b620000cc7f000000000000000000000000000000000000000000000000000000000000000081565b620001356200012f36600462000f61565b62000245565b6040516fffffffffffffffffffffffffffffffff9091168152602001620000ed565b620000cc6200016836600462000ef8565b620004a8565b620001356200017f36600462000f61565b620004d1565b620001356200019636600462000f61565b620006b2565b62000135620001ad36600462000f61565b620007e0565b620001ca620001c436600462001034565b6200090e565b6040519015158152602001620000ed565b600254620000cc9073ffffffffffffffffffffffffffffffffffffffff1681565b600060205282600052604060002060205281600052604060002081600981106200022557600080fd5b015473ffffffffffffffffffffffffffffffffffffffff16925083915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602090815260408083208584529091528120819060ff8716600981106200028e576200028e6200115b565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050620003448173ffffffffffffffffffffffffffffffffffffffff1663d94073d46040518163ffffffff1660e01b81526004016020604051808303816000875af115801562000304573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200032a91906200118a565b82856fffffffffffffffffffffffffffffffff1662000c2e565b6040517fe463854d0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff82169063c45be1f2903390839063e463854d906024015b6020604051808303816000875af1158015620003ce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003f49190620011b1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526fffffffffffffffffffffffffffffffff1660248201526044015b6020604051808303816000875af115801562000478573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200049e9190620011b1565b9695505050505050565b600160205282600052604060002060205281600052604060002081600981106200022557600080fd5b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602090815260408083208584529091528120819060ff8716600981106200051a576200051a6200115b565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050620005908173ffffffffffffffffffffffffffffffffffffffff1663d94073d46040518163ffffffff1660e01b81526004016020604051808303816000875af115801562000304573d6000803e3d6000fd5b6040517f76a912530000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff82169063e78a9b0c90339083906376a91253906024015b6020604051808303816000875af11580156200061a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620006409190620011b1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526fffffffffffffffffffffffffffffffff90811660248301528616604482015260640162000458565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602090815260408083208584529091528120819060ff871660098110620006fb57620006fb6200115b565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050620007718173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af115801562000304573d6000803e3d6000fd5b6040517f87c09bbc0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff821690636045eae490339083906387c09bbc90602401620005fa565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602090815260408083208584529091528120819060ff8716600981106200082957620008296200115b565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506200089f8173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af115801562000304573d6000803e3d6000fd5b6040517ff332ea5e0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff821690633e232091903390839063f332ea5e90602401620003ae565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633811462000999576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a65640000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a81166000908152602081815260408083208d8452909152902054161562000a33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d61726b65742065786973747300000000000000000000000000000000000000604482015260640162000990565b60008a8a898989898960405162000a4a9062000e19565b62000a5c97969594939291906200121a565b604051809103906000f08015801562000a79573d6000803e3d6000fd5b50604080516101208101825273ffffffffffffffffffffffffffffffffffffffff83811682528c5181166020808401919091528d0151811682840152918c015182166060808301919091528c015182166080808301919091528c0151821660a0808301919091528c0151821660c0808301919091528c0151821660e0808301919091528c01519091166101008201529091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60005b60088160ff16101562000b9a5762000b85838260ff166009811062000b585762000b586200115b565b60200201517f00000000000000000000000000000000000000000000000000000000000000008462000cff565b8062000b91816200127f565b91505062000b2f565b5073ffffffffffffffffffffffffffffffffffffffff8d166000908152602081815260408083208f8452909152902062000bd79083600962000e27565b506040518c9073ffffffffffffffffffffffffffffffffffffffff8f16907f0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee90600090a35060019c9b505050505050505050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af191505062000c918162000dca565b62000cf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015260640162000990565b50505050565b60006040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af191505062000d628162000dca565b62000cf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f617070726f7665206661696c6564000000000000000000000000000000000000604482015260640162000990565b6000803d8362000dde57806000803e806000fd5b806020811462000df957801562000e0b576000925062000e10565b816000803e6000511515925062000e10565b600192505b50909392505050565b61071880620012c783390190565b826009810192821562000e97579160200282015b8281111562000e9757825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019062000e3b565b5062000ea592915062000ea9565b5090565b5b8082111562000ea5576000815560010162000eaa565b73ffffffffffffffffffffffffffffffffffffffff8116811462000ee357600080fd5b50565b803562000ef38162000ec0565b919050565b60008060006060848603121562000f0e57600080fd5b833562000f1b8162000ec0565b95602085013595506040909401359392505050565b803560ff8116811462000ef357600080fd5b6fffffffffffffffffffffffffffffffff8116811462000ee357600080fd5b6000806000806080858703121562000f7857600080fd5b62000f838562000f30565b9350602085013562000f958162000ec0565b925060408501359150606085013562000fae8162000f42565b939692955090935050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083601f84011262000ffb57600080fd5b50813567ffffffffffffffff8111156200101457600080fd5b6020830191508360208285010111156200102d57600080fd5b9250929050565b6000806000806000806000806101a0898b0312156200105257600080fd5b88356200105f8162000ec0565b97506020898101359750605f8a018b136200107957600080fd5b604051610100810167ffffffffffffffff8282108183111715620010a157620010a162000fb9565b816040528291506101408d018e811115620010bb57600080fd5b60408e015b81811015620010e257620010d48162000ee6565b8452928501928501620010c0565b50839a50803594505080841115620010f957600080fd5b620011078e858f0162000fe8565b90995097506101608d01359350889250808411156200112557600080fd5b505050620011368b828c0162000fe8565b90945092506200114c90506101808a0162000f30565b90509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156200119d57600080fd5b8151620011aa8162000ec0565b9392505050565b600060208284031215620011c457600080fd5b8151620011aa8162000f42565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a0604082015260006200125260a083018789620011d1565b828103606084015262001267818688620011d1565b91505060ff8316608083015298975050505050505050565b600060ff821660ff8103620012bd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019291505056fe608060405234801561001057600080fd5b5060405161071838038061071883398101604081905261002f916101fd565b600180546001600160a01b0319166001600160a01b03871617815560028590556003805460ff19169091179055825161006f9060049060208601906100a2565b5081516100839060059060208501906100a2565b506006805460ff191660ff92909216919091179055506102de92505050565b8280546100ae906102a4565b90600052602060002090601f0160209004810192826100d05760008555610116565b82601f106100e957805160ff1916838001178555610116565b82800160010185558215610116579182015b828111156101165782518255916020019190600101906100fb565b50610122929150610126565b5090565b5b808211156101225760008155600101610127565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261016257600080fd5b81516001600160401b038082111561017c5761017c61013b565b604051601f8301601f19908116603f011681019082821181831017156101a4576101a461013b565b816040528381526020925086838588010111156101c057600080fd5b600091505b838210156101e257858201830151818301840152908201906101c5565b838211156101f35760008385830101525b9695505050505050565b600080600080600060a0868803121561021557600080fd5b85516001600160a01b038116811461022c57600080fd5b6020870151604088015191965094506001600160401b038082111561025057600080fd5b61025c89838a01610151565b9450606088015191508082111561027257600080fd5b5061027f88828901610151565b925050608086015160ff8116811461029657600080fd5b809150509295509295909350565b600181811c908216806102b857607f821691505b6020821081036102d857634e487b7160e01b600052602260045260246000fd5b50919050565b61042b806102ed6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a082311461014757806395d89b411461017d5780639dd0ff3714610185578063c0c6e5ab146101c657600080fd5b806306fdde031461008d578063095ea7b3146100ab578063313ce567146100fa5780636581d54314610119575b600080fd5b6100956101fd565b6040516100a29190610298565b60405180910390f35b6100ea6100b9366004610334565b73ffffffffffffffffffffffffffffffffffffffff9190911660009081526020819052604090205560035460ff1690565b60405190151581526020016100a2565b6006546101079060ff1681565b60405160ff90911681526020016100a2565b61013961012736600461035e565b60006020819052908152604090205481565b6040519081526020016100a2565b61013961015536600461035e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526007602052604090205490565b61009561028b565b6101c4610193366004610380565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b6101c46101d4366004610334565b73ffffffffffffffffffffffffffffffffffffffff909116600090815260076020526040902055565b6004805461020a906103a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610236906103a2565b80156102835780601f1061025857610100808354040283529160200191610283565b820191906000526020600020905b81548152906001019060200180831161026657829003601f168201915b505050505081565b6005805461020a906103a2565b600060208083528351808285015260005b818110156102c5578581018301518582016040015282016102a9565b818111156102d7576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461032f57600080fd5b919050565b6000806040838503121561034757600080fd5b6103508361030b565b946020939093013593505050565b60006020828403121561037057600080fd5b6103798261030b565b9392505050565b60006020828403121561039257600080fd5b8135801515811461037957600080fd5b600181811c908216806103b657607f821691505b6020821081036103ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea2646970667358221220582a4df78f7ba2b51135c430746dd3ef4acafd57d5f3a2ba973b51d1c2dedee064736f6c634300080d0033a2646970667358221220b1672f372c5cea7f47e5f2d72e262e8e10f3d4b99cb7f4e2a6957632cf5c0bae64736f6c634300080d0033"

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketPlaceBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// MarketPlace is an auto generated Go binding around an Ethereum contract.
type MarketPlace struct {
	MarketPlaceCaller     // Read-only binding to the contract
	MarketPlaceTransactor // Write-only binding to the contract
	MarketPlaceFilterer   // Log filterer for contract events
}

// MarketPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketPlaceSession struct {
	Contract     *MarketPlace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketPlaceCallerSession struct {
	Contract *MarketPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketPlaceTransactorSession struct {
	Contract     *MarketPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketPlaceRaw struct {
	Contract *MarketPlace // Generic contract binding to access the raw methods on
}

// MarketPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketPlaceCallerRaw struct {
	Contract *MarketPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketPlaceTransactorRaw struct {
	Contract *MarketPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketPlace creates a new instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlace(address common.Address, backend bind.ContractBackend) (*MarketPlace, error) {
	contract, err := bindMarketPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketPlace{MarketPlaceCaller: MarketPlaceCaller{contract: contract}, MarketPlaceTransactor: MarketPlaceTransactor{contract: contract}, MarketPlaceFilterer: MarketPlaceFilterer{contract: contract}}, nil
}

// NewMarketPlaceCaller creates a new read-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceCaller(address common.Address, caller bind.ContractCaller) (*MarketPlaceCaller, error) {
	contract, err := bindMarketPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCaller{contract: contract}, nil
}

// NewMarketPlaceTransactor creates a new write-only instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketPlaceTransactor, error) {
	contract, err := bindMarketPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceTransactor{contract: contract}, nil
}

// NewMarketPlaceFilterer creates a new log filterer instance of MarketPlace, bound to a specific deployed contract.
func NewMarketPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketPlaceFilterer, error) {
	contract, err := bindMarketPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceFilterer{contract: contract}, nil
}

// bindMarketPlace binds a generic wrapper to an already deployed contract.
func bindMarketPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.MarketPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.MarketPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketPlace *MarketPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketPlace *MarketPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketPlace *MarketPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketPlace.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Admin() (common.Address, error) {
	return _MarketPlace.Contract.Admin(&_MarketPlace.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceCaller) Markets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "markets", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Markets is a free data retrieval call binding the contract method 0x125cf47f.
//
// Solidity: function markets(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Markets(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Pools is a free data retrieval call binding the contract method 0x5b0cbb05.
//
// Solidity: function pools(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "pools", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0x5b0cbb05.
//
// Solidity: function pools(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceSession) Pools(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.Pools(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Pools is a free data retrieval call binding the contract method 0x5b0cbb05.
//
// Solidity: function pools(address , uint256 , uint256 ) view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Pools(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (common.Address, error) {
	return _MarketPlace.Contract.Pools(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_MarketPlace *MarketPlaceCaller) Redeemer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "redeemer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_MarketPlace *MarketPlaceSession) Redeemer() (common.Address, error) {
	return _MarketPlace.Contract.Redeemer(&_MarketPlace.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_MarketPlace *MarketPlaceCallerSession) Redeemer() (common.Address, error) {
	return _MarketPlace.Contract.Redeemer(&_MarketPlace.CallOpts)
}

// BuyPT is a paid mutator transaction binding the contract method 0xb892d15a.
//
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactor) BuyPT(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "buyPT", p, u, m, a)
}

// BuyPT is a paid mutator transaction binding the contract method 0xb892d15a.
//
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceSession) BuyPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyPT is a paid mutator transaction binding the contract method 0xb892d15a.
//
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactorSession) BuyPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactor) BuyUnderlying(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "buyUnderlying", p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceSession) BuyUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactorSession) BuyUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcef26d43.
//
// Solidity: function createMarket(address u, uint256 m, address[8] t, string n, string s, uint8 d) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CreateMarket(opts *bind.TransactOpts, u common.Address, m *big.Int, t [8]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "createMarket", u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcef26d43.
//
// Solidity: function createMarket(address u, uint256 m, address[8] t, string n, string s, uint8 d) returns(bool)
func (_MarketPlace *MarketPlaceSession) CreateMarket(u common.Address, m *big.Int, t [8]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, u, m, t, n, s, d)
}

// CreateMarket is a paid mutator transaction binding the contract method 0xcef26d43.
//
// Solidity: function createMarket(address u, uint256 m, address[8] t, string n, string s, uint8 d) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CreateMarket(u common.Address, m *big.Int, t [8]common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, u, m, t, n, s, d)
}

// SellPT is a paid mutator transaction binding the contract method 0x52289a15.
//
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactor) SellPT(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "sellPT", p, u, m, a)
}

// SellPT is a paid mutator transaction binding the contract method 0x52289a15.
//
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceSession) SellPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellPT is a paid mutator transaction binding the contract method 0x52289a15.
//
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactorSession) SellPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactor) SellUnderlying(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "sellUnderlying", p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceSession) SellUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128 returned)
func (_MarketPlace *MarketPlaceTransactorSession) SellUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// MarketPlaceCreateMarketIterator is returned from FilterCreateMarket and is used to iterate over the raw logs and unpacked data for CreateMarket events raised by the MarketPlace contract.
type MarketPlaceCreateMarketIterator struct {
	Event *MarketPlaceCreateMarket // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCreateMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCreateMarket)
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
		it.Event = new(MarketPlaceCreateMarket)
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
func (it *MarketPlaceCreateMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCreateMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCreateMarket represents a CreateMarket event raised by the MarketPlace contract.
type MarketPlaceCreateMarket struct {
	Underlying common.Address
	Maturity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreateMarket is a free log retrieval operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_MarketPlace *MarketPlaceFilterer) FilterCreateMarket(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*MarketPlaceCreateMarketIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCreateMarketIterator{contract: _MarketPlace.contract, event: "CreateMarket", logs: logs, sub: sub}, nil
}

// WatchCreateMarket is a free log subscription operation binding the contract event 0x0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee.
//
// Solidity: event CreateMarket(address indexed underlying, uint256 indexed maturity)
func (_MarketPlace *MarketPlaceFilterer) WatchCreateMarket(opts *bind.WatchOpts, sink chan<- *MarketPlaceCreateMarket, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "CreateMarket", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCreateMarket)
				if err := _MarketPlace.contract.UnpackLog(event, "CreateMarket", log); err != nil {
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
func (_MarketPlace *MarketPlaceFilterer) ParseCreateMarket(log types.Log) (*MarketPlaceCreateMarket, error) {
	event := new(MarketPlaceCreateMarket)
	if err := _MarketPlace.contract.UnpackLog(event, "CreateMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
