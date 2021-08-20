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
const MarketPlaceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferVaultNotionalFeeCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferVaultNotionalFeeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
var MarketPlaceBin = "0x608060405234801561001057600080fd5b506117b8806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80637d04cd15116100c3578063c4c257261161007c578063c4c2572614610454578063d557ee8514610487578063dc098af9146104a3578063eee6e814146104bf578063ef267f2c146104f2578063f8e51bcb1461052257610158565b80637d04cd151461035957806389de80aa146103755780638c6b9b41146103a8578063b50a66f7146103d8578063bddbfbe414610408578063beaff41e1461043857610158565b80633cf9a4e3116101155780633cf9a4e3146102475780633f25be9d146102775780633f73df2c146102aa5780634521d303146102da5780634bc81e091461030d57806365a963aa1461032957610158565b806303f5e4001461015d57806304aa1dfd1461019057806305e1dc25146101ac5780631cd9be91146101dc5780632634de19146101f857806333fe997a14610214575b600080fd5b610177600480360381019061017291906114ca565b610552565b60405161018794939291906116b0565b60405180910390f35b6101aa60048036038101906101a59190611609565b6105c2565b005b6101c660048036038101906101c191906114f3565b6105df565b6040516101d3919061165f565b60405180910390f35b6101f660048036038101906101f19190611609565b61064f565b005b610212600480360381019061020d9190611609565b61066c565b005b61022e600480360381019061022991906114ca565b610689565b60405161023e94939291906116b0565b60405180910390f35b610261600480360381019061025c91906115a6565b6106f9565b60405161026e919061167a565b60405180910390f35b610291600480360381019061028c91906114ca565b61084f565b6040516102a194939291906116b0565b60405180910390f35b6102c460048036038101906102bf91906114ca565b6108bf565b6040516102d19190611695565b60405180910390f35b6102f460048036038101906102ef91906114ca565b6108d7565b60405161030494939291906116b0565b60405180910390f35b61032760048036038101906103229190611609565b610947565b005b610343600480360381019061033e919061152f565b610964565b604051610350919061167a565b60405180910390f35b610373600480360381019061036e9190611609565b610af3565b005b61038f600480360381019061038a91906114ca565b610b10565b60405161039f94939291906116b0565b60405180910390f35b6103c260048036038101906103bd919061152f565b610b80565b6040516103cf919061167a565b60405180910390f35b6103f260048036038101906103ed91906115a6565b610d0f565b6040516103ff919061167a565b60405180910390f35b610422600480360381019061041d919061152f565b610e65565b60405161042f919061167a565b60405180910390f35b610452600480360381019061044d9190611609565b610ff4565b005b61046e600480360381019061046991906114ca565b611011565b60405161047e94939291906116b0565b60405180910390f35b6104a1600480360381019061049c91906114ca565b611081565b005b6104bd60048036038101906104b89190611609565b6110c5565b005b6104d960048036038101906104d491906114ca565b6110e2565b6040516104e994939291906116b0565b60405180910390f35b61050c600480360381019061050791906115a6565b611152565b604051610519919061167a565b60405180910390f35b61053c6004803603810190610537919061152f565b6112a8565b604051610549919061167a565b60405180910390f35b60066020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b806008601a6101000a81548160ff02191690831515021790555050565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b80600860156101000a81548160ff02191690831515021790555050565b80600860146101000a81548160ff02191690831515021790555050565b60076020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610703611437565b8481600001818152505083816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600760008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301559050506008601a9054906101000a900460ff16915050949350505050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006020528060005260406000206000915090505481565b60026020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b80600860186101000a81548160ff02191690831515021790555050565b600061096e611437565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860169054906101000a900460ff1691505095945050505050565b80600860176101000a81548160ff02191690831515021790555050565b60056020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000610b8a611437565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860159054906101000a900460ff1691505095945050505050565b6000610d19611437565b8481600001818152505083816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860199054906101000a900460ff16915050949350505050565b6000610e6f611437565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860179054906101000a900460ff1691505095945050505050565b80600860166101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b80600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600860196101000a81548160ff02191690831515021790555050565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b600061115c611437565b8481600001818152505083816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860189054906101000a900460ff16915050949350505050565b60006112b2611437565b8581600001818152505084816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281606001818152505080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155905050600860149054906101000a900460ff1691505095945050505050565b604051806080016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b60008135905061149a8161173d565b92915050565b6000813590506114af81611754565b92915050565b6000813590506114c48161176b565b92915050565b6000602082840312156114dc57600080fd5b60006114ea8482850161148b565b91505092915050565b6000806040838503121561150657600080fd5b60006115148582860161148b565b9250506020611525858286016114b5565b9150509250929050565b600080600080600060a0868803121561154757600080fd5b60006115558882890161148b565b9550506020611566888289016114b5565b94505060406115778882890161148b565b93505060606115888882890161148b565b9250506080611599888289016114b5565b9150509295509295909350565b600080600080608085870312156115bc57600080fd5b60006115ca8782880161148b565b94505060206115db878288016114b5565b93505060406115ec8782880161148b565b92505060606115fd878288016114b5565b91505092959194509250565b60006020828403121561161b57600080fd5b6000611629848285016114a0565b91505092915050565b61163b816116f5565b82525050565b61164a81611707565b82525050565b61165981611733565b82525050565b60006020820190506116746000830184611632565b92915050565b600060208201905061168f6000830184611641565b92915050565b60006020820190506116aa6000830184611650565b92915050565b60006080820190506116c56000830187611650565b6116d26020830186611632565b6116df6040830185611632565b6116ec6060830184611650565b95945050505050565b600061170082611713565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b611746816116f5565b811461175157600080fd5b50565b61175d81611707565b811461176857600080fd5b50565b61177481611733565b811461177f57600080fd5b5056fea26469706673582212203d3f10c8b4bb20851fab313d6be2bfbd7e36de39a4a0dec2a25bfaf4533a36b864736f6c63430008040033"

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

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x33fe997a.
//
// Solidity: function transferVaultNotionalFeeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferVaultNotionalFeeCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferVaultNotionalFeeCalled", arg0)

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

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x33fe997a.
//
// Solidity: function transferVaultNotionalFeeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFeeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x33fe997a.
//
// Solidity: function transferVaultNotionalFeeCalled(address ) view returns(uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferVaultNotionalFeeCalled(arg0 common.Address) (struct {
	Maturity *big.Int
	One      common.Address
	Two      common.Address
	Amount   *big.Int
}, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeCalled(&_MarketPlace.CallOpts, arg0)
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

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0x3cf9a4e3.
//
// Solidity: function transferVaultNotionalFee(address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferVaultNotionalFee(opts *bind.TransactOpts, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferVaultNotionalFee", u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0x3cf9a4e3.
//
// Solidity: function transferVaultNotionalFee(address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFee(u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0x3cf9a4e3.
//
// Solidity: function transferVaultNotionalFee(address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferVaultNotionalFee(u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, u, m, f, a)
}

// TransferVaultNotionalFeeReturns is a paid mutator transaction binding the contract method 0x04aa1dfd.
//
// Solidity: function transferVaultNotionalFeeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactor) TransferVaultNotionalFeeReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferVaultNotionalFeeReturns", b)
}

// TransferVaultNotionalFeeReturns is a paid mutator transaction binding the contract method 0x04aa1dfd.
//
// Solidity: function transferVaultNotionalFeeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFeeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeReturns(&_MarketPlace.TransactOpts, b)
}

// TransferVaultNotionalFeeReturns is a paid mutator transaction binding the contract method 0x04aa1dfd.
//
// Solidity: function transferVaultNotionalFeeReturns(bool b) returns()
func (_MarketPlace *MarketPlaceTransactorSession) TransferVaultNotionalFeeReturns(b bool) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeReturns(&_MarketPlace.TransactOpts, b)
}
