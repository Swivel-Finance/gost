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
const MarketPlaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"matured\",\"type\":\"uint256\"}],\"name\":\"Mature\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"createMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"matureMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"reddemVaultInterest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561005057600080fd5b5061166b806100606000396000f3fe60806040523480156200001157600080fd5b5060043610620000945760003560e01c80635db0ae5811620000635780635db0ae5814620001735780636ce327ec14620001a9578063c86adf7c14620001df578063f851a44014620002155762000094565b8063154e0f2e146200009957806317b3bba714620000cf57806327ee93be14620001075780633271e5f3146200013d575b600080fd5b620000b76004803603810190620000b191906200081b565b62000237565b604051620000c6919062000a09565b60405180910390f35b620000ed6004803603810190620000e7919062000726565b62000244565b604051620000fe9392919062000971565b60405180910390f35b6200012560048036038101906200011f919062000726565b620002db565b60405162000134919062000a09565b60405180910390f35b6200015b600480360381019062000155919062000726565b6200030a565b6040516200016a919062000a09565b60405180910390f35b6200019160048036038101906200018b919062000726565b62000316565b604051620001a0919062000a09565b60405180910390f35b620001c76004803603810190620001c1919062000767565b62000322565b604051620001d6919062000a09565b60405180910390f35b620001fd6004803603810190620001f7919062000726565b62000621565b6040516200020c919062000a48565b60405180910390f35b6200021f62000646565b6040516200022e919062000927565b60405180910390f35b6000600190509392505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083565b60026020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006001905092915050565b60006001905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620003b7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003ae9062000a26565b60405180910390fd5b600087878686604051620003cb906200066a565b620003da9493929190620009ae565b604051809103906000f080158015620003f7573d6000803e3d6000fd5b509050600087876040516200040c9062000678565b6200041992919062000a65565b604051809103906000f08015801562000436573d6000803e3d6000fd5b50905060405180606001604052808873ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff16815250600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050878973ffffffffffffffffffffffffffffffffffffffff167f5c8943d544aa04815bd907fde89a7a9156cd14eb9ffe4f5c3b1277d7557a2d2289856040516200060992919062000944565b60405180910390a36001935050505095945050505050565b6003602052816000526040600020602052806000526040600020600091509150505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106cd8062000c1983390190565b61035080620012e683390190565b60006200069d620006978462000ac6565b62000a92565b905082815260208101848484011115620006b657600080fd5b620006c384828562000b5f565b509392505050565b600081359050620006dc8162000be4565b92915050565b600082601f830112620006f457600080fd5b81356200070684826020860162000686565b91505092915050565b600081359050620007208162000bfe565b92915050565b600080604083850312156200073a57600080fd5b60006200074a85828601620006cb565b92505060206200075d858286016200070f565b9150509250929050565b600080600080600060a086880312156200078057600080fd5b60006200079088828901620006cb565b9550506020620007a3888289016200070f565b9450506040620007b688828901620006cb565b935050606086013567ffffffffffffffff811115620007d457600080fd5b620007e288828901620006e2565b925050608086013567ffffffffffffffff8111156200080057600080fd5b6200080e88828901620006e2565b9150509295509295909350565b6000806000606084860312156200083157600080fd5b60006200084186828701620006cb565b935050602062000854868287016200070f565b925050604062000867868287016200070f565b9150509250925092565b6200087c8162000b15565b82525050565b6200088d8162000b29565b82525050565b6000620008a08262000af9565b620008ac818562000b04565b9350620008be81856020860162000b6e565b620008c98162000bd3565b840191505092915050565b6000620008e360148362000b04565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b620009218162000b55565b82525050565b60006020820190506200093e600083018462000871565b92915050565b60006040820190506200095b600083018562000871565b6200096a602083018462000871565b9392505050565b600060608201905062000988600083018662000871565b62000997602083018562000871565b620009a6604083018462000871565b949350505050565b6000608082019050620009c5600083018762000871565b620009d4602083018662000916565b8181036040830152620009e8818562000893565b90508181036060830152620009fe818462000893565b905095945050505050565b600060208201905062000a20600083018462000882565b92915050565b6000602082019050818103600083015262000a4181620008d4565b9050919050565b600060208201905062000a5f600083018462000916565b92915050565b600060408201905062000a7c600083018562000916565b62000a8b602083018462000871565b9392505050565b6000604051905081810181811067ffffffffffffffff8211171562000abc5762000abb62000ba4565b5b8060405250919050565b600067ffffffffffffffff82111562000ae45762000ae362000ba4565b5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600062000b228262000b35565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101562000b8e57808201518184015260208101905062000b71565b8381111562000b9e576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b62000bef8162000b15565b811462000bfb57600080fd5b50565b62000c098162000b55565b811462000c1557600080fd5b5056fe608060405234801561001057600080fd5b506040516106cd3803806106cd833981810160405281019061003291906101e6565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260038190555081600090805190602001906100909291906100b1565b5080600190805190602001906100a79291906100b1565b5050505050610407565b8280546100bd90610349565b90600052602060002090601f0160209004810192826100df5760008555610126565b82601f106100f857805160ff1916838001178555610126565b82800160010185558215610126579182015b8281111561012557825182559160200191906001019061010a565b5b5090506101339190610137565b5090565b5b80821115610150576000816000905550600101610138565b5090565b6000610167610162846102aa565b610279565b90508281526020810184848401111561017f57600080fd5b61018a848285610316565b509392505050565b6000815190506101a1816103d9565b92915050565b600082601f8301126101b857600080fd5b81516101c8848260208601610154565b91505092915050565b6000815190506101e0816103f0565b92915050565b600080600080608085870312156101fc57600080fd5b600061020a87828801610192565b945050602061021b878288016101d1565b935050604085015167ffffffffffffffff81111561023857600080fd5b610244878288016101a7565b925050606085015167ffffffffffffffff81111561026157600080fd5b61026d878288016101a7565b91505092959194509250565b6000604051905081810181811067ffffffffffffffff821117156102a05761029f6103aa565b5b8060405250919050565b600067ffffffffffffffff8211156102c5576102c46103aa565b5b601f19601f8301169050602081019050919050565b60006102e5826102ec565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015610334578082015181840152602081019050610319565b83811115610343576000848401525b50505050565b6000600282049050600182168061036157607f821691505b602082108114156103755761037461037b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103e2816102da565b81146103ed57600080fd5b50565b6103f98161030c565b811461040457600080fd5b50565b6102b7806104166000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063204f83f9146100515780636f307dc31461006f578063b4c4a4c81461008d578063e7ba6774146100a9575b600080fd5b6100596100c5565b60405161006691906101fc565b60405180910390f35b6100776100cf565b60405161008491906101e1565b60405180910390f35b6100a760048036038101906100a2919061019a565b6100f9565b005b6100c360048036038101906100be9190610171565b610103565b005b6000600354905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060038190555050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008135905061015681610253565b92915050565b60008135905061016b8161026a565b92915050565b60006020828403121561018357600080fd5b600061019184828501610147565b91505092915050565b6000602082840312156101ac57600080fd5b60006101ba8482850161015c565b91505092915050565b6101cc81610217565b82525050565b6101db81610249565b82525050565b60006020820190506101f660008301846101c3565b92915050565b600060208201905061021160008301846101d2565b92915050565b600061022282610229565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61025c81610217565b811461026757600080fd5b50565b61027381610249565b811461027e57600080fd5b5056fea264697066735822122008320ef92ff3de52c01825f41e370ac2420c3d8756015d7136e385bf3c29b66264736f6c63430008000033608060405234801561001057600080fd5b50604051610350380380610350833981810160405281019061003291906100aa565b81600181905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050610150565b60008151905061008f81610122565b92915050565b6000815190506100a481610139565b92915050565b600080604083850312156100bd57600080fd5b60006100cb85828601610095565b92505060206100dc85828601610080565b9150509250929050565b60006100f1826100f8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61012b816100e6565b811461013657600080fd5b50565b61014281610118565b811461014d57600080fd5b50565b6101f18061015f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063204f83f914610046578063b4c4a4c814610064578063b7dd348314610080575b600080fd5b61004e61009e565b60405161005b919061014d565b60405180910390f35b61007e600480360381019061007991906100eb565b6100a8565b005b6100886100b2565b6040516100959190610132565b60405180910390f35b6000600154905090565b8060018190555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000813590506100e5816101a4565b92915050565b6000602082840312156100fd57600080fd5b600061010b848285016100d6565b91505092915050565b61011d81610168565b82525050565b61012c8161019a565b82525050565b60006020820190506101476000830184610114565b92915050565b60006020820190506101626000830184610123565b92915050565b60006101738261017a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6101ad8161019a565b81146101b857600080fd5b5056fea2646970667358221220d399eddb5731edc38c50a05ba183d0321606fac3d537ddb86120d6fbfd84def364736f6c63430008000033a264697066735822122032b6ac6567aef2fd31c4394eeb7d37b34f5d4d4a91f47ca517be437d68d3e1c664736f6c63430008000033"

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
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) FilterMature(opts *bind.FilterOpts) (*MarketPlaceMatureIterator, error) {

	logs, sub, err := _MarketPlace.contract.FilterLogs(opts, "Mature")
	if err != nil {
		return nil, err
	}
	return &MarketPlaceMatureIterator{contract: _MarketPlace.contract, event: "Mature", logs: logs, sub: sub}, nil
}

// WatchMature is a free log subscription operation binding the contract event 0x0080e09d7b4544aa5a923873be1df3e31945593d40cb1c874d99259ec3ac43a4.
//
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) WatchMature(opts *bind.WatchOpts, sink chan<- *MarketPlaceMature) (event.Subscription, error) {

	logs, sub, err := _MarketPlace.contract.WatchLogs(opts, "Mature")
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
// Solidity: event Mature(address underlying, uint256 maturity, uint256 maturityRate, uint256 matured)
func (_MarketPlace *MarketPlaceFilterer) ParseMature(log types.Log) (*MarketPlaceMature, error) {
	event := new(MarketPlaceMature)
	if err := _MarketPlace.contract.UnpackLog(event, "Mature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
