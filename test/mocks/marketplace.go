// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b50611535806100206000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806389de80aa116100b8578063c4c257261161007c578063c4c25726146103b4578063d557ee85146103e7578063dc098af914610403578063eee6e8141461041f578063ef267f2c14610452578063f8e51bcb1461048257610137565b806389de80aa146102d55780638c6b9b4114610308578063b50a66f714610338578063bddbfbe414610368578063beaff41e1461039857610137565b80633f73df2c116100ff5780633f73df2c1461020a5780634521d3031461023a5780634bc81e091461026d57806365a963aa146102895780637d04cd15146102b957610137565b806303f5e4001461013c57806305e1dc251461016f5780631cd9be911461019f5780632634de19146101bb5780633f25be9d146101d7575b600080fd5b61015660048036038101906101519190611247565b6104b2565b604051610166949392919061142d565b60405180910390f35b61018960048036038101906101849190611270565b610522565b60405161019691906113dc565b60405180910390f35b6101b960048036038101906101b49190611386565b610592565b005b6101d560048036038101906101d09190611386565b6105af565b005b6101f160048036038101906101ec9190611247565b6105cc565b604051610201949392919061142d565b60405180910390f35b610224600480360381019061021f9190611247565b61063c565b6040516102319190611412565b60405180910390f35b610254600480360381019061024f9190611247565b610654565b604051610264949392919061142d565b60405180910390f35b61028760048036038101906102829190611386565b6106c4565b005b6102a3600480360381019061029e91906112ac565b6106e1565b6040516102b091906113f7565b60405180910390f35b6102d360048036038101906102ce9190611386565b610870565b005b6102ef60048036038101906102ea9190611247565b61088d565b6040516102ff949392919061142d565b60405180910390f35b610322600480360381019061031d91906112ac565b6108fd565b60405161032f91906113f7565b60405180910390f35b610352600480360381019061034d9190611323565b610a8c565b60405161035f91906113f7565b60405180910390f35b610382600480360381019061037d91906112ac565b610be2565b60405161038f91906113f7565b60405180910390f35b6103b260048036038101906103ad9190611386565b610d71565b005b6103ce60048036038101906103c99190611247565b610d8e565b6040516103de949392919061142d565b60405180910390f35b61040160048036038101906103fc9190611247565b610dfe565b005b61041d60048036038101906104189190611386565b610e42565b005b61043960048036038101906104349190611247565b610e5f565b604051610449949392919061142d565b60405180910390f35b61046c60048036038101906104679190611323565b610ecf565b60405161047991906113f7565b60405180910390f35b61049c600480360381019061049791906112ac565b611025565b6040516104a991906113f7565b60405180910390f35b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b80600760156101000a81548160ff02191690831515021790555050565b80600760146101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006020528060005260406000206000915090505481565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b80600760186101000a81548160ff02191690831515021790555050565b60006106eb6111b4565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760169054906101000a900460ff1691505095945050505050565b80600760176101000a81548160ff02191690831515021790555050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006109076111b4565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760159054906101000a900460ff1691505095945050505050565b6000610a966111b4565b8481600001818152505083816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760199054906101000a900460ff16915050949350505050565b6000610bec6111b4565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760179054906101000a900460ff1691505095945050505050565b80600760166101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600760196101000a81548160ff02191690831515021790555050565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610ed96111b4565b8481600001818152505083816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760189054906101000a900460ff16915050949350505050565b600061102f6111b4565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600760149054906101000a900460ff1691505095945050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600081359050611217816114ba565b92915050565b60008135905061122c816114d1565b92915050565b600081359050611241816114e8565b92915050565b60006020828403121561125957600080fd5b600061126784828501611208565b91505092915050565b6000806040838503121561128357600080fd5b600061129185828601611208565b92505060206112a285828601611232565b9150509250929050565b600080600080600060a086880312156112c457600080fd5b60006112d288828901611208565b95505060206112e388828901611232565b94505060406112f488828901611208565b935050606061130588828901611208565b925050608061131688828901611232565b9150509295509295909350565b6000806000806080858703121561133957600080fd5b600061134787828801611208565b945050602061135887828801611232565b935050604061136987828801611208565b925050606061137a87828801611232565b91505092959194509250565b60006020828403121561139857600080fd5b60006113a68482850161121d565b91505092915050565b6113b881611472565b82525050565b6113c781611484565b82525050565b6113d6816114b0565b82525050565b60006020820190506113f160008301846113af565b92915050565b600060208201905061140c60008301846113be565b92915050565b600060208201905061142760008301846113cd565b92915050565b600060808201905061144260008301876113cd565b61144f60208301866113af565b61145c60408301856113af565b61146960608301846113cd565b95945050505050565b600061147d82611490565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6114c381611472565b81146114ce57600080fd5b50565b6114da81611484565b81146114e557600080fd5b50565b6114f1816114b0565b81146114fc57600080fd5b5056fea26469706673582212201881de46b28df98d7e166ed5fc130a3895b51972f6299dc9657ea66e7cfcc20464736f6c63430008040033"

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

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) BurnZcTokenRemovingNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "burnZcTokenRemovingNotionalCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x03f5e400.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) BurnZcTokenRemovingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) CTokenAddressCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAddressCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x3f73df2c.
//
// Solidity: function cTokenAddressCalled(address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAddressCalled(arg0 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialExitCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialExitCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialExitCalled is a free data retrieval call binding the contract method 0x4521d303.
//
// Solidity: function custodialExitCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialExitCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialInitiateCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialInitiateCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x3f25be9d.
//
// Solidity: function custodialInitiateCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialInitiateCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) MintZcTokenAddingNotionalCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mintZcTokenAddingNotionalCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x89de80aa.
//
// Solidity: function mintZcTokenAddingNotionalCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) MintZcTokenAddingNotionalCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pVaultExchangeCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pVaultExchangeCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xeee6e814.
//
// Solidity: function p2pVaultExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pVaultExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pZcTokenExchangeCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pZcTokenExchangeCalled", arg0)

	outstruct := new(struct {
		Maturity *big.Int
		One      common.Address
		Two      common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Maturity = out[0].(*big.Int)
	outstruct.One = out[1].(common.Address)
	outstruct.Two = out[2].(common.Address)
	outstruct.Amount = out[3].(*big.Int)

	return *outstruct, err

}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xc4c25726.
//
// Solidity: function p2pZcTokenExchangeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pZcTokenExchangeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0xb50a66f7.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotional", u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0xb50a66f7.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotional(u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0xb50a66f7.
//
// Solidity: function burnZcTokenRemovingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotional(u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, u, m, t, a)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotionalReturns", b)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// BurnZcTokenRemovingNotionalReturns is a paid mutator transaction binding the contract method 0xdc098af9.
//
// Solidity: function burnZcTokenRemovingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) CTokenAddress(opts *bind.TransactOpts, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddress", u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x05e1dc25.
//
// Solidity: function cTokenAddress(address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddress(u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, u, m)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactor) CTokenAddressReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddressReturns", a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddressReturns(a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialExit(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExit", u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x8c6b9b41.
//
// Solidity: function custodialExit(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExit(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) CustodialExitReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExitReturns", b)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) CustodialExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExitReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialExitReturns is a paid mutator transaction binding the contract method 0x1cd9be91.
//
// Solidity: function custodialExitReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExitReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExitReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiate(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiate", u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xf8e51bcb.
//
// Solidity: function custodialInitiate(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiate(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiateReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiateReturns", b)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) CustodialInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// CustodialInitiateReturns is a paid mutator transaction binding the contract method 0x2634de19.
//
// Solidity: function custodialInitiateReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiateReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiateReturns(&_MarketPlace.TransactOpts, b)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xef267f2c.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotional(opts *bind.TransactOpts, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotional", u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xef267f2c.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0xef267f2c.
//
// Solidity: function mintZcTokenAddingNotional(address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotional(u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, u, m, t, a)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotionalReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotionalReturns", b)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// MintZcTokenAddingNotionalReturns is a paid mutator transaction binding the contract method 0x4bc81e09.
//
// Solidity: function mintZcTokenAddingNotionalReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotionalReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalReturns(&_MarketPlace.TransactOpts, b)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchange(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchange", u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0xbddbfbe4.
//
// Solidity: function p2pVaultExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchangeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchangeReturns", b)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) P2pVaultExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pVaultExchangeReturns is a paid mutator transaction binding the contract method 0x7d04cd15.
//
// Solidity: function p2pVaultExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchange(opts *bind.TransactOpts, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchange", u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0x65a963aa.
//
// Solidity: function p2pZcTokenExchange(address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchange(u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, u, m, o, t, a)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchangeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchangeReturns", b)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeReturns(&_MarketPlace.TransactOpts, b)
}

// P2pZcTokenExchangeReturns is a paid mutator transaction binding the contract method 0xbeaff41e.
//
// Solidity: function p2pZcTokenExchangeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchangeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeReturns(&_MarketPlace.TransactOpts, b)
}
