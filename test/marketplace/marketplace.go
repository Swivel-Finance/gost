// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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
const MarketPlaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matured\",\"type\":\"uint256\"}],\"name\":\"Mature\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"matureMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"reddemVaultInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561005057600080fd5b50611df3806100606000396000f3fe60806040523480156200001157600080fd5b5060043610620000945760003560e01c80635db0ae5811620000635780635db0ae5814620001735780636ce327ec14620001a9578063c86adf7c14620001df578063f851a44014620002155762000094565b8063154e0f2e146200009957806317b3bba714620000cf57806327ee93be14620001075780633271e5f3146200013d575b600080fd5b620000b76004803603810190620000b1919062000d2d565b62000237565b604051620000c6919062000ff7565b60405180910390f35b620000ed6004803603810190620000e7919062000c38565b62000244565b604051620000fe9392919062000f5f565b60405180910390f35b6200012560048036038101906200011f919062000c38565b620002db565b60405162000134919062000ff7565b60405180910390f35b6200015b600480360381019062000155919062000c38565b6200030a565b6040516200016a919062000ff7565b60405180910390f35b6200019160048036038101906200018b919062000c38565b62000316565b604051620001a0919062000ff7565b60405180910390f35b620001c76004803603810190620001c1919062000c79565b62000806565b604051620001d6919062000ff7565b60405180910390f35b620001fd6004803603810190620001f7919062000c38565b62000b05565b6040516200020c91906200107a565b60405180910390f35b6200021f62000b2a565b6040516200022e919062000f15565b60405180910390f35b6000600190509392505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006001905092915050565b6000801515600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060009054906101000a900460ff16151514620003be576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003b59062001058565b60405180910390fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663204f83f96040518163ffffffff1660e01b815260040160206040518083038186803b1580156200047857600080fd5b505afa1580156200048d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004b3919062000daf565b421015620004f8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004ef9062001036565b60405180910390fd5b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015620005b657600080fd5b505af1158015620005cb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005f1919062000daf565b905080600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002081905550600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636b868d516040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156200070457600080fd5b505af115801562000719573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200073f919062000d83565b506001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060006101000a81548160ff021916908315150217905550828473ffffffffffffffffffffffffffffffffffffffff167e80e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a44284604051620007f3929190620010c4565b60405180910390a3600191505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200089b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008929062001014565b60405180910390fd5b600087878686604051620008af9062000b4e565b620008be949392919062000f9c565b604051809103906000f080158015620008db573d6000803e3d6000fd5b50905060008787604051620008f09062000b5c565b620008fd92919062001097565b604051809103906000f0801580156200091a573d6000803e3d6000fd5b50905060405180606001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815250600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050878973ffffffffffffffffffffffffffffffffffffffff167f5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22898560405162000aed92919062000f32565b60405180910390a36001935050505095945050505050565b6003602052816000526040600020602052806000526040600020600091509150505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106cd806200129283390190565b61045f806200195f83390190565b600062000b8162000b7b8462001125565b620010f1565b90508281526020810184848401111562000b9a57600080fd5b62000ba7848285620011be565b509392505050565b60008135905062000bc08162001243565b92915050565b60008151905062000bd7816200125d565b92915050565b600082601f83011262000bef57600080fd5b813562000c0184826020860162000b6a565b91505092915050565b60008135905062000c1b8162001277565b92915050565b60008151905062000c328162001277565b92915050565b6000806040838503121562000c4c57600080fd5b600062000c5c8582860162000baf565b925050602062000c6f8582860162000c0a565b9150509250929050565b600080600080600060a0868803121562000c9257600080fd5b600062000ca28882890162000baf565b955050602062000cb58882890162000c0a565b945050604062000cc88882890162000baf565b935050606086013567ffffffffffffffff81111562000ce657600080fd5b62000cf48882890162000bdd565b925050608086013567ffffffffffffffff81111562000d1257600080fd5b62000d208882890162000bdd565b9150509295509295909350565b60008060006060848603121562000d4357600080fd5b600062000d538682870162000baf565b935050602062000d668682870162000c0a565b925050604062000d798682870162000c0a565b9150509250925092565b60006020828403121562000d9657600080fd5b600062000da68482850162000bc6565b91505092915050565b60006020828403121562000dc257600080fd5b600062000dd28482850162000c21565b91505092915050565b62000de68162001174565b82525050565b62000df78162001188565b82525050565b600062000e0a8262001158565b62000e16818562001163565b935062000e28818560208601620011cd565b62000e338162001232565b840191505092915050565b600062000e4d60148362001163565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b600062000e8f60148362001163565b91507f4d61747572697479206e6f7420726561636865640000000000000000000000006000830152602082019050919050565b600062000ed160168362001163565b91507f4d61726b657420616c7265616479206d617475726564000000000000000000006000830152602082019050919050565b62000f0f81620011b4565b82525050565b600060208201905062000f2c600083018462000ddb565b92915050565b600060408201905062000f49600083018562000ddb565b62000f58602083018462000ddb565b9392505050565b600060608201905062000f76600083018662000ddb565b62000f85602083018562000ddb565b62000f94604083018462000ddb565b949350505050565b600060808201905062000fb3600083018762000ddb565b62000fc2602083018662000f04565b818103604083015262000fd6818562000dfd565b9050818103606083015262000fec818462000dfd565b905095945050505050565b60006020820190506200100e600083018462000dec565b92915050565b600060208201905081810360008301526200102f8162000e3e565b9050919050565b60006020820190508181036000830152620010518162000e80565b9050919050565b60006020820190508181036000830152620010738162000ec2565b9050919050565b600060208201905062001091600083018462000f04565b92915050565b6000604082019050620010ae600083018562000f04565b620010bd602083018462000ddb565b9392505050565b6000604082019050620010db600083018562000f04565b620010ea602083018462000f04565b9392505050565b6000604051905081810181811067ffffffffffffffff821117156200111b576200111a62001203565b5b8060405250919050565b600067ffffffffffffffff82111562001143576200114262001203565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000620011818262001194565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015620011ed578082015181840152602081019050620011d0565b83811115620011fd576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6200124e8162001174565b81146200125a57600080fd5b50565b620012688162001188565b81146200127457600080fd5b50565b6200128281620011b4565b81146200128e57600080fd5b5056fe608060405234801561001057600080fd5b506040516106cd3803806106cd833981810160405281019061003291906101e6565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260038190555081600090805190602001906100909291906100b1565b5080600190805190602001906100a79291906100b1565b5050505050610407565b8280546100bd90610349565b90600052602060002090601f0160209004810192826100df5760008555610126565b82601f106100f857805160ff1916838001178555610126565b82800160010185558215610126579182015b8281111561012557825182559160200191906001019061010a565b5b5090506101339190610137565b5090565b5b80821115610150576000816000905550600101610138565b5090565b6000610167610162846102aa565b610279565b90508281526020810184848401111561017f57600080fd5b61018a848285610316565b509392505050565b6000815190506101a1816103d9565b92915050565b600082601f8301126101b857600080fd5b81516101c8848260208601610154565b91505092915050565b6000815190506101e0816103f0565b92915050565b600080600080608085870312156101fc57600080fd5b600061020a87828801610192565b945050602061021b878288016101d1565b935050604085015167ffffffffffffffff81111561023857600080fd5b610244878288016101a7565b925050606085015167ffffffffffffffff81111561026157600080fd5b61026d878288016101a7565b91505092959194509250565b6000604051905081810181811067ffffffffffffffff821117156102a05761029f6103aa565b5b8060405250919050565b600067ffffffffffffffff8211156102c5576102c46103aa565b5b601f19601f8301169050602081019050919050565b60006102e5826102ec565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015610334578082015181840152602081019050610319565b83811115610343576000848401525b50505050565b6000600282049050600182168061036157607f821691505b602082108114156103755761037461037b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103e2816102da565b81146103ed57600080fd5b50565b6103f98161030c565b811461040457600080fd5b50565b6102b7806104166000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063204f83f9146100515780636f307dc31461006f578063b4c4a4c81461008d578063e7ba6774146100a9575b600080fd5b6100596100c5565b60405161006691906101fc565b60405180910390f35b6100776100cf565b60405161008491906101e1565b60405180910390f35b6100a760048036038101906100a2919061019a565b6100f9565b005b6100c360048036038101906100be9190610171565b610103565b005b6000600354905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060038190555050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008135905061015681610253565b92915050565b60008135905061016b8161026a565b92915050565b60006020828403121561018357600080fd5b600061019184828501610147565b91505092915050565b6000602082840312156101ac57600080fd5b60006101ba8482850161015c565b91505092915050565b6101cc81610217565b82525050565b6101db81610249565b82525050565b60006020820190506101f660008301846101c3565b92915050565b600060208201905061021160008301846101d2565b92915050565b600061022282610229565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61025c81610217565b811461026757600080fd5b50565b61027381610249565b811461027e57600080fd5b5056fea264697066735822122008320ef92ff3de52c01825f41e370ac2420c3d8756015d7136e385bf3c29b66264736f6c63430008000033608060405234801561001057600080fd5b5060405161045f38038061045f833981810160405281019061003291906100aa565b81600181905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610150565b60008151905061008f81610122565b92915050565b6000815190506100a481610139565b92915050565b600080604083850312156100bd57600080fd5b60006100cb85828601610095565b92505060206100dc85828601610080565b9150509250929050565b60006100f1826100f8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61012b816100e6565b811461013657600080fd5b50565b61014281610118565b811461014d57600080fd5b50565b6103008061015f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063204f83f91461005c5780636b868d511461007a578063b4c4a4c814610098578063b7dd3483146100b4578063da3de9e9146100d2575b600080fd5b6100646100ee565b6040516100719190610239565b60405180910390f35b6100826100f8565b60405161008f919061021e565b60405180910390f35b6100b260048036038101906100ad91906101ad565b61010f565b005b6100bc610119565b6040516100c99190610203565b60405180910390f35b6100ec60048036038101906100e79190610184565b61013d565b005b6000600154905090565b6000600260009054906101000a900460ff16905090565b8060018190555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600260006101000a81548160ff02191690831515021790555050565b6000813590506101698161029c565b92915050565b60008135905061017e816102b3565b92915050565b60006020828403121561019657600080fd5b60006101a48482850161015a565b91505092915050565b6000602082840312156101bf57600080fd5b60006101cd8482850161016f565b91505092915050565b6101df81610254565b82525050565b6101ee81610266565b82525050565b6101fd81610292565b82525050565b600060208201905061021860008301846101d6565b92915050565b600060208201905061023360008301846101e5565b92915050565b600060208201905061024e60008301846101f4565b92915050565b600061025f82610272565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6102a581610266565b81146102b057600080fd5b50565b6102bc81610292565b81146102c757600080fd5b5056fea26469706673582212203faaa4f773f2f03631496dc0888e233d63d3dc020a2ba8d170fcd8b8fd0ca7a564736f6c63430008000033a264697066735822122078a328e2cbd689d3326a15af649944cb4080c77b3ac642a7708b18a62fd9047764736f6c63430008000033"

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketPlaceBin), backend)
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

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceCaller) Markets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "markets", arg0, arg1)

	outstruct := new(struct {
		CTokenAddr  common.Address
		ZcTokenAddr common.Address
		VaultAddr   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CTokenAddr = out[0].(common.Address)
	outstruct.ZcTokenAddr = out[1].(common.Address)
	outstruct.VaultAddr = out[2].(common.Address)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceSession) Markets(arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x17b3bba7.
//
// Solidity: function markets(address , uint256 ) view returns(address cTokenAddr, address zcTokenAddr, address vaultAddr)
func (_MarketPlace *MarketPlaceCallerSession) Markets(arg0 common.Address, arg1 *big.Int) (struct {
	CTokenAddr  common.Address
	ZcTokenAddr common.Address
	VaultAddr   common.Address
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1)
}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceCaller) Mature(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mature", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceSession) Mature(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MarketPlace.Contract.Mature(&_MarketPlace.CallOpts, arg0, arg1)
}

// Mature is a free data retrieval call binding the contract method 0x27ee93be.
//
// Solidity: function mature(address , uint256 ) view returns(bool)
func (_MarketPlace *MarketPlaceCallerSession) Mature(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _MarketPlace.Contract.Mature(&_MarketPlace.CallOpts, arg0, arg1)
}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) MaturityRate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "maturityRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) MaturityRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketPlace.Contract.MaturityRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// MaturityRate is a free data retrieval call binding the contract method 0xc86adf7c.
//
// Solidity: function maturityRate(address , uint256 ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) MaturityRate(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _MarketPlace.Contract.MaturityRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x6ce327ec.
//
// Solidity: function createMarket(address u, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CreateMarket(opts *bind.TransactOpts, u common.Address, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "createMarket", u, m, c, n, s)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x6ce327ec.
//
// Solidity: function createMarket(address u, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceSession) CreateMarket(u common.Address, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, u, m, c, n, s)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x6ce327ec.
//
// Solidity: function createMarket(address u, uint256 m, address c, string n, string s) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CreateMarket(u common.Address, m *big.Int, c common.Address, n string, s string) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, u, m, c, n, s)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MatureMarket(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "matureMarket", u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceSession) MatureMarket(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, u, m)
}

// MatureMarket is a paid mutator transaction binding the contract method 0x5db0ae58.
//
// Solidity: function matureMarket(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MatureMarket(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MatureMarket(&_MarketPlace.TransactOpts, u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) ReddemVaultInterest(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "reddemVaultInterest", u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceSession) ReddemVaultInterest(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ReddemVaultInterest(&_MarketPlace.TransactOpts, u, m)
}

// ReddemVaultInterest is a paid mutator transaction binding the contract method 0x3271e5f3.
//
// Solidity: function reddemVaultInterest(address u, uint256 m) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) ReddemVaultInterest(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ReddemVaultInterest(&_MarketPlace.TransactOpts, u, m)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) RedeemZcToken(opts *bind.TransactOpts, u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemZcToken", u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, u, m, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x154e0f2e.
//
// Solidity: function redeemZcToken(address u, uint256 m, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemZcToken(u common.Address, m *big.Int, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, u, m, a)
}

// MarketPlaceCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the MarketPlace contract.
type MarketPlaceCreateIterator struct {
	Event *MarketPlaceCreate // Event containing the contract specifics and raw log

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
func (it *MarketPlaceCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceCreate)
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
		it.Event = new(MarketPlaceCreate)
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
func (it *MarketPlaceCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceCreate represents a Create event raised by the MarketPlace contract.
type MarketPlaceCreate struct {
	Underlying common.Address
	Maturity   *big.Int
	CToken     common.Address
	ZcToken    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address indexed underlying, uint256 indexed maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) FilterCreate(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*MarketPlaceCreateIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Create", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceCreateIterator{contract: _MarketPlace.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address indexed underlying, uint256 indexed maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *MarketPlaceCreate, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Create", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceCreate)
				if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d22.
//
// Solidity: event Create(address indexed underlying, uint256 indexed maturity, address cToken, address zcToken)
func (_MarketPlace *MarketPlaceFilterer) ParseCreate(log types.Log) (*MarketPlaceCreate, error) {
	event := new(MarketPlaceCreate)
	if err := _MarketPlace.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPlaceMatureIterator is returned from FilterMature and is used to iterate over the raw logs and unpacked data for Mature events raised by the MarketPlace contract.
type MarketPlaceMatureIterator struct {
	Event *MarketPlaceMature // Event containing the contract specifics and raw log

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
func (it *MarketPlaceMatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPlaceMature)
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
		it.Event = new(MarketPlaceMature)
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
func (it *MarketPlaceMatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPlaceMatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPlaceMature represents a Mature event raised by the MarketPlace contract.
type MarketPlaceMature struct {
	Underlying   common.Address
	Maturity     *big.Int
	MaturityRate *big.Int
	Matured      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMature is a free log retrieval operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) FilterMature(opts *bind.FilterOpts, underlying []common.Address, maturity []*big.Int) (*MarketPlaceMatureIterator, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Mature", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return &MarketPlaceMatureIterator{contract: _MarketPlace.contract, event: "Mature", logs: logs, sub: sub}, nil
}

// WatchMature is a free log subscription operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) WatchMature(opts *bind.WatchOpts, sink chan<- *MarketPlaceMature, underlying []common.Address, maturity []*big.Int) (event.Subscription, error) {

	var underlyingRule []interface{}
	for _, underlyingItem := range underlying {
		underlyingRule = append(underlyingRule, underlyingItem)
	}
	var maturityRule []interface{}
	for _, maturityItem := range maturity {
		maturityRule = append(maturityRule, maturityItem)
	}

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Mature", underlyingRule, maturityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPlaceMature)
				if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
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

// ParseMature is a log parse operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address indexed underlying, uint256 indexed maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) ParseMature(log types.Log) (*MarketPlaceMature, error) {
	event := new(MarketPlaceMature)
	if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
