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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"}],\"name\":\"CreateMarket\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"buyPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"buyUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address[8]\",\"name\":\"t\",\"type\":\"address[8]\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"sellPT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"a\",\"type\":\"uint128\"}],\"name\":\"sellUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x60a060405234801561001057600080fd5b50604051611f73380380611f7383398101604081905261002f91610052565b600280546001600160a01b031916331790556001600160a01b0316608052610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611ecf6100a4600039600081816101250152610fcd0152611ecf6000f3fe60806040523480156200001157600080fd5b5060043610620000d95760003560e01c8063ad31b198116200008b578063c9ac53b11162000062578063c9ac53b11462000213578063cef26d43146200022a578063f851a440146200024157600080fd5b8063ad31b19814620001ce578063b46f1efa14620001e5578063b892d15a14620001fc57600080fd5b806352289a1511620000c057806352289a1514620001475780635b0cbb0514620001805780635c975abb146200019757600080fd5b8063125cf47f14620000de5780632ba29d38146200011f575b600080fd5b620000f5620000ef36600462001366565b62000262565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b620000f57f000000000000000000000000000000000000000000000000000000000000000081565b6200015e62000158366004620013cf565b620002ab565b6040516fffffffffffffffffffffffffffffffff909116815260200162000116565b620000f56200019136600462001366565b6200059a565b600254620001bd9074010000000000000000000000000000000000000000900460ff1681565b604051901515815260200162000116565b6200015e620001df366004620013cf565b620005c3565b620001bd620001f636600462001427565b6200082c565b6200015e6200020d366004620013cf565b62000a14565b6200015e62000224366004620013cf565b62000bca565b620001bd6200023b366004620014ef565b62000d80565b600254620000f59073ffffffffffffffffffffffffffffffffffffffff1681565b600060205282600052604060002060205281600052604060002081600981106200028b57600080fd5b015473ffffffffffffffffffffffffffffffffffffffff16925083915050565b60025460009074010000000000000000000000000000000000000000900460ff161562000339576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d61726b6574732061726520706175736564000000000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84166000908152600160209081526040808320868452909152812060ff87166009811062000380576200038062001616565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050620004368173ffffffffffffffffffffffffffffffffffffffff1663d94073d46040518163ffffffff1660e01b81526004016020604051808303816000875af1158015620003f6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200041c919062001645565b82856fffffffffffffffffffffffffffffffff166200109c565b6040517fe463854d0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff82169063c45be1f2903390839063e463854d906024015b6020604051808303816000875af1158015620004c0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004e691906200166c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526fffffffffffffffffffffffffffffffff1660248201526044015b6020604051808303816000875af11580156200056a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200059091906200166c565b9695505050505050565b600160205282600052604060002060205281600052604060002081600981106200028b57600080fd5b60025460009074010000000000000000000000000000000000000000900460ff16156200064d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d61726b65747320617265207061757365640000000000000000000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600160209081526040808320868452909152812060ff87166009811062000694576200069462001616565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506200070a8173ffffffffffffffffffffffffffffffffffffffff1663d94073d46040518163ffffffff1660e01b81526004016020604051808303816000875af1158015620003f6573d6000803e3d6000fd5b6040517f76a912530000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff82169063e78a9b0c90339083906376a91253906024015b6020604051808303816000875af115801562000794573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007ba91906200166c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526fffffffffffffffffffffffffffffffff9081166024830152861660448201526064016200054a565b60025460009073ffffffffffffffffffffffffffffffffffffffff16338114620008b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a656400000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600160209081526040808320878452909152812060ff881660098110620008fa57620008fa62001616565b015473ffffffffffffffffffffffffffffffffffffffff16146200097b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f706f6f6c20657869737473000000000000000000000000000000000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff851660009081526001602090815260408083208784529091529020839060ff881660098110620009c457620009c462001616565b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550600195945050505050565b60025460009074010000000000000000000000000000000000000000900460ff161562000a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d61726b65747320617265207061757365640000000000000000000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600160209081526040808320868452909152812060ff87166009811062000ae55762000ae562001616565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905062000b5b8173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015620003f6573d6000803e3d6000fd5b6040517f87c09bbc0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff821690636045eae490339083906387c09bbc9060240162000774565b60025460009074010000000000000000000000000000000000000000900460ff161562000c54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d61726b65747320617265207061757365640000000000000000000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600160209081526040808320868452909152812060ff87166009811062000c9b5762000c9b62001616565b0160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905062000d118173ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b81526004016020604051808303816000875af1158015620003f6573d6000803e3d6000fd5b6040517ff332ea5e0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff8416600482015273ffffffffffffffffffffffffffffffffffffffff821690633e232091903390839063f332ea5e90602401620004a0565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633811462000e07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f73656e646572206d75737420626520617574686f72697a656400000000000000604482015260640162000330565b73ffffffffffffffffffffffffffffffffffffffff8a81166000908152602081815260408083208d8452909152902054161562000ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d61726b65742065786973747300000000000000000000000000000000000000604482015260640162000330565b60008a8a898989898960405162000eb89062001287565b62000eca9796959493929190620016d5565b604051809103906000f08015801562000ee7573d6000803e3d6000fd5b50604080516101208101825273ffffffffffffffffffffffffffffffffffffffff83811682528c5181166020808401919091528d0151811682840152918c015182166060808301919091528c015182166080808301919091528c0151821660a0808301919091528c0151821660c0808301919091528c0151821660e0808301919091528c01519091166101008201529091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60005b60088160ff161015620010085762000ff3838260ff166009811062000fc65762000fc662001616565b60200201517f0000000000000000000000000000000000000000000000000000000000000000846200116d565b8062000fff816200173a565b91505062000f9d565b5073ffffffffffffffffffffffffffffffffffffffff8d166000908152602081815260408083208f84529091529020620010459083600962001295565b506040518c9073ffffffffffffffffffffffffffffffffffffffff8f16907f0ce205d5fda43f489af3d143ec11073757d0079e07852a6c1799f1e29e72e1ee90600090a35060019c9b505050505050505050505050565b60006040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050620010ff8162001238565b62001167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f7472616e73666572206661696c65640000000000000000000000000000000000604482015260640162000330565b50505050565b60006040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201528260248201526000806044836000895af1915050620011d08162001238565b62001167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f617070726f7665206661696c6564000000000000000000000000000000000000604482015260640162000330565b6000803d836200124c57806000803e806000fd5b8060208114620012675780156200127957600092506200127e565b816000803e600051151592506200127e565b600192505b50909392505050565b610718806200178283390190565b826009810192821562001305579160200282015b828111156200130557825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190620012a9565b506200131392915062001317565b5090565b5b8082111562001313576000815560010162001318565b73ffffffffffffffffffffffffffffffffffffffff811681146200135157600080fd5b50565b803562001361816200132e565b919050565b6000806000606084860312156200137c57600080fd5b833562001389816200132e565b95602085013595506040909401359392505050565b803560ff811681146200136157600080fd5b6fffffffffffffffffffffffffffffffff811681146200135157600080fd5b60008060008060808587031215620013e657600080fd5b620013f1856200139e565b9350602085013562001403816200132e565b92506040850135915060608501356200141c81620013b0565b939692955090935050565b600080600080608085870312156200143e57600080fd5b62001449856200139e565b935060208501356200145b816200132e565b92506040850135915060608501356200141c816200132e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008083601f840112620014b657600080fd5b50813567ffffffffffffffff811115620014cf57600080fd5b602083019150836020828501011115620014e857600080fd5b9250929050565b6000806000806000806000806101a0898b0312156200150d57600080fd5b88356200151a816200132e565b97506020898101359750605f8a018b136200153457600080fd5b604051610100810167ffffffffffffffff82821081831117156200155c576200155c62001474565b816040528291506101408d018e8111156200157657600080fd5b60408e015b818110156200159d576200158f8162001354565b84529285019285016200157b565b50839a50803594505080841115620015b457600080fd5b620015c28e858f01620014a3565b90995097506101608d0135935088925080841115620015e057600080fd5b505050620015f18b828c01620014a3565b90945092506200160790506101808a016200139e565b90509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156200165857600080fd5b815162001665816200132e565b9392505050565b6000602082840312156200167f57600080fd5b81516200166581620013b0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a0604082015260006200170d60a0830187896200168c565b8281036060840152620017228186886200168c565b91505060ff8316608083015298975050505050505050565b600060ff821660ff810362001778577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001019291505056fe608060405234801561001057600080fd5b5060405161071838038061071883398101604081905261002f916101fd565b600180546001600160a01b0319166001600160a01b03871617815560028590556003805460ff19169091179055825161006f9060049060208601906100a2565b5081516100839060059060208501906100a2565b506006805460ff191660ff92909216919091179055506102de92505050565b8280546100ae906102a4565b90600052602060002090601f0160209004810192826100d05760008555610116565b82601f106100e957805160ff1916838001178555610116565b82800160010185558215610116579182015b828111156101165782518255916020019190600101906100fb565b50610122929150610126565b5090565b5b808211156101225760008155600101610127565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261016257600080fd5b81516001600160401b038082111561017c5761017c61013b565b604051601f8301601f19908116603f011681019082821181831017156101a4576101a461013b565b816040528381526020925086838588010111156101c057600080fd5b600091505b838210156101e257858201830151818301840152908201906101c5565b838211156101f35760008385830101525b9695505050505050565b600080600080600060a0868803121561021557600080fd5b85516001600160a01b038116811461022c57600080fd5b6020870151604088015191965094506001600160401b038082111561025057600080fd5b61025c89838a01610151565b9450606088015191508082111561027257600080fd5b5061027f88828901610151565b925050608086015160ff8116811461029657600080fd5b809150509295509295909350565b600181811c908216806102b857607f821691505b6020821081036102d857634e487b7160e01b600052602260045260246000fd5b50919050565b61042b806102ed6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a082311461014757806395d89b411461017d5780639dd0ff3714610185578063c0c6e5ab146101c657600080fd5b806306fdde031461008d578063095ea7b3146100ab578063313ce567146100fa5780636581d54314610119575b600080fd5b6100956101fd565b6040516100a29190610298565b60405180910390f35b6100ea6100b9366004610334565b73ffffffffffffffffffffffffffffffffffffffff9190911660009081526020819052604090205560035460ff1690565b60405190151581526020016100a2565b6006546101079060ff1681565b60405160ff90911681526020016100a2565b61013961012736600461035e565b60006020819052908152604090205481565b6040519081526020016100a2565b61013961015536600461035e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526007602052604090205490565b61009561028b565b6101c4610193366004610380565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b6101c46101d4366004610334565b73ffffffffffffffffffffffffffffffffffffffff909116600090815260076020526040902055565b6004805461020a906103a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610236906103a2565b80156102835780601f1061025857610100808354040283529160200191610283565b820191906000526020600020905b81548152906001019060200180831161026657829003601f168201915b505050505081565b6005805461020a906103a2565b600060208083528351808285015260005b818110156102c5578581018301518582016040015282016102a9565b818111156102d7576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461032f57600080fd5b919050565b6000806040838503121561034757600080fd5b6103508361030b565b946020939093013593505050565b60006020828403121561037057600080fd5b6103798261030b565b9392505050565b60006020828403121561039257600080fd5b8135801515811461037957600080fd5b600181811c908216806103b657607f821691505b6020821081036103ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea2646970667358221220582a4df78f7ba2b51135c430746dd3ef4acafd57d5f3a2ba973b51d1c2dedee064736f6c634300080d0033a264697066735822122035ec697601fcc476392d3c698e790298088bd8aba46ece2532901f1a6b61762064736f6c634300080d0033"

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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketPlace *MarketPlaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketPlace *MarketPlaceSession) Paused() (bool, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketPlace *MarketPlaceCallerSession) Paused() (bool, error) {
	return _MarketPlace.Contract.Paused(&_MarketPlace.CallOpts)
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
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactor) BuyPT(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "buyPT", p, u, m, a)
}

// BuyPT is a paid mutator transaction binding the contract method 0xb892d15a.
//
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceSession) BuyPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyPT is a paid mutator transaction binding the contract method 0xb892d15a.
//
// Solidity: function buyPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactorSession) BuyPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactor) BuyUnderlying(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "buyUnderlying", p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceSession) BuyUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BuyUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xad31b198.
//
// Solidity: function buyUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
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
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactor) SellPT(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "sellPT", p, u, m, a)
}

// SellPT is a paid mutator transaction binding the contract method 0x52289a15.
//
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceSession) SellPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellPT is a paid mutator transaction binding the contract method 0x52289a15.
//
// Solidity: function sellPT(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactorSession) SellPT(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellPT(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactor) SellUnderlying(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "sellUnderlying", p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceSession) SellUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0xc9ac53b1.
//
// Solidity: function sellUnderlying(uint8 p, address u, uint256 m, uint128 a) returns(uint128)
func (_MarketPlace *MarketPlaceTransactorSession) SellUnderlying(p uint8, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.SellUnderlying(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SetPool is a paid mutator transaction binding the contract method 0xb46f1efa.
//
// Solidity: function setPool(uint8 p, address u, uint256 m, address a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) SetPool(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "setPool", p, u, m, a)
}

// SetPool is a paid mutator transaction binding the contract method 0xb46f1efa.
//
// Solidity: function setPool(uint8 p, address u, uint256 m, address a) returns(bool)
func (_MarketPlace *MarketPlaceSession) SetPool(p uint8, u common.Address, m *big.Int, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetPool(&_MarketPlace.TransactOpts, p, u, m, a)
}

// SetPool is a paid mutator transaction binding the contract method 0xb46f1efa.
//
// Solidity: function setPool(uint8 p, address u, uint256 m, address a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) SetPool(p uint8, u common.Address, m *big.Int, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.SetPool(&_MarketPlace.TransactOpts, p, u, m, a)
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
