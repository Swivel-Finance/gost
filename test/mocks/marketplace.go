// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocks

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MarketPlaceMetaData contains all meta data concerning the MarketPlace contract.
var MarketPlaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAndAdapterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"cTokenAndAdapterAddressCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"cTokenAndAdapterAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"redeemVaultInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemVaultInterestCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemVaultInterestReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemZcTokenCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"transferVaultNotionalFeeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferVaultNotionalFeeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506124ce806100206000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c806350e3fe2c11610104578063beaff41e116100a2578063dc098af911610071578063dc098af91461062c578063de2ad3eb14610648578063e3e43c581461067c578063fcbaab2e146106b0576101da565b8063beaff41e1461057c578063c06760c714610598578063d721b5d5146105c8578063db850901146105fc576101da565b80637d04cd15116100de5780637d04cd15146104cc57806387e157c1146104e857806397c2c691146105185780639f6eddc41461054c576101da565b806350e3fe2c146104485780636ac910331461047c5780636ba14fe714610498576101da565b80631cd9be911161017c57806327d3aae21161014b57806327d3aae2146103ac578063295b25f2146103e05780633a660bd8146103fc5780634bc81e091461042c576101da565b80631cd9be911461030f5780632512ca631461032b5780632634de191461035c57806326817c7714610378576101da565b80630f0016b6116101b85780630f0016b61461025f5780630f5ea3c71461028f57806315042ddf146102c3578063153d4625146102f3576101da565b806301cc6448146101df57806304aa1dfd1461020f5780630e3374621461022b575b600080fd5b6101f960048036038101906101f491906120f8565b6106e0565b604051610206919061218e565b60405180910390f35b610229600480360381019061022491906121d5565b610890565b005b61024560048036038101906102409190612202565b6108ad565b60405161025695949392919061224d565b60405180910390f35b610279600480360381019061027491906122a0565b610943565b604051610286919061218e565b60405180910390f35b6102a960048036038101906102a49190612202565b610b2c565b6040516102ba95949392919061224d565b60405180910390f35b6102dd60048036038101906102d891906122a0565b610bc2565b6040516102ea919061218e565b60405180910390f35b61030d6004803603810190610308919061232d565b610dab565b005b610329600480360381019061032491906121d5565b610e31565b005b6103456004803603810190610340919061236d565b610e4e565b6040516103539291906123c0565b60405180910390f35b610376600480360381019061037191906121d5565b610ff3565b005b610392600480360381019061038d9190612202565b611010565b6040516103a395949392919061224d565b60405180910390f35b6103c660048036038101906103c19190612202565b6110a6565b6040516103d795949392919061224d565b60405180910390f35b6103fa60048036038101906103f591906123e9565b61113c565b005b61041660048036038101906104119190612416565b611146565b604051610423919061247d565b60405180910390f35b610446600480360381019061044191906121d5565b6112de565b005b610462600480360381019061045d9190612202565b6112fb565b60405161047395949392919061224d565b60405180910390f35b610496600480360381019061049191906123e9565b611391565b005b6104b260048036038101906104ad9190612202565b61139b565b6040516104c395949392919061224d565b60405180910390f35b6104e660048036038101906104e191906121d5565b611431565b005b61050260048036038101906104fd91906120f8565b61144e565b60405161050f919061218e565b60405180910390f35b610532600480360381019061052d9190612202565b6115fe565b60405161054395949392919061224d565b60405180910390f35b610566600480360381019061056191906120f8565b611694565b604051610573919061247d565b60405180910390f35b610596600480360381019061059191906121d5565b611837565b005b6105b260048036038101906105ad91906122a0565b611854565b6040516105bf919061218e565b60405180910390f35b6105e260048036038101906105dd9190612202565b611a3d565b6040516105f395949392919061224d565b60405180910390f35b610616600480360381019061061191906120f8565b611ad3565b604051610623919061218e565b60405180910390f35b610646600480360381019061064191906121d5565b611c83565b005b610662600480360381019061065d9190612202565b611ca0565b60405161067395949392919061224d565b60405180910390f35b61069660048036038101906106919190612202565b611d36565b6040516106a795949392919061224d565b60405180910390f35b6106ca60048036038101906106c591906122a0565b611dcc565b6040516106d7919061218e565b60405180910390f35b60006106ea611fb5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600560008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60189054906101000a900460ff1691505095945050505050565b80600b601a6101000a81548160ff02191690831515021790555050565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b600061094d611fb5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600260008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60159054906101000a900460ff169150509695505050505050565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000610bcc611fb5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600460008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60179054906101000a900460ff169150509695505050505050565b81600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b80600b60156101000a81548160ff02191690831515021790555050565b600080610e59611fb5565b84816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816020018181525050806000808860ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250925050935093915050565b80600b60146101000a81548160ff02191690831515021790555050565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600d8190555050565b6000611150611fb5565b84816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381602001818152505082816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600960008860ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600d54915050949350505050565b80600b60186101000a81548160ff02191690831515021790555050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600c8190555050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600b60176101000a81548160ff02191690831515021790555050565b6000611458611fb5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600660008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60199054906101000a900460ff1691505095945050505050565b60086020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b600061169e611fb5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600860008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c5491505095945050505050565b80600b60166101000a81548160ff02191690831515021790555050565b600061185e611fb5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600160008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60149054906101000a900460ff169150509695505050505050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000611add611fb5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600760008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b601a9054906101000a900460ff1691505095945050505050565b80600b60196101000a81548160ff02191690831515021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000611dd6611fb5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600360008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600b60169054906101000a900460ff169150509695505050505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600060ff82169050919050565b6120418161202b565b811461204c57600080fd5b50565b60008135905061205e81612038565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061208f82612064565b9050919050565b61209f81612084565b81146120aa57600080fd5b50565b6000813590506120bc81612096565b92915050565b6000819050919050565b6120d5816120c2565b81146120e057600080fd5b50565b6000813590506120f2816120cc565b92915050565b600080600080600060a0868803121561211457612113612026565b5b60006121228882890161204f565b9550506020612133888289016120ad565b9450506040612144888289016120e3565b9350506060612155888289016120ad565b9250506080612166888289016120e3565b9150509295509295909350565b60008115159050919050565b61218881612173565b82525050565b60006020820190506121a3600083018461217f565b92915050565b6121b281612173565b81146121bd57600080fd5b50565b6000813590506121cf816121a9565b92915050565b6000602082840312156121eb576121ea612026565b5b60006121f9848285016121c0565b91505092915050565b60006020828403121561221857612217612026565b5b60006122268482850161204f565b91505092915050565b61223881612084565b82525050565b612247816120c2565b82525050565b600060a082019050612262600083018861222f565b61226f602083018761223e565b61227c604083018661222f565b612289606083018561222f565b612296608083018461223e565b9695505050505050565b60008060008060008060c087890312156122bd576122bc612026565b5b60006122cb89828a0161204f565b96505060206122dc89828a016120ad565b95505060406122ed89828a016120e3565b94505060606122fe89828a016120ad565b935050608061230f89828a016120ad565b92505060a061232089828a016120e3565b9150509295509295509295565b6000806040838503121561234457612343612026565b5b6000612352858286016120ad565b9250506020612363858286016120ad565b9150509250929050565b60008060006060848603121561238657612385612026565b5b60006123948682870161204f565b93505060206123a5868287016120ad565b92505060406123b6868287016120e3565b9150509250925092565b60006040820190506123d5600083018561222f565b6123e2602083018461222f565b9392505050565b6000602082840312156123ff576123fe612026565b5b600061240d848285016120e3565b91505092915050565b600080600080608085870312156124305761242f612026565b5b600061243e8782880161204f565b945050602061244f878288016120ad565b9350506040612460878288016120e3565b9250506060612471878288016120ad565b91505092959194509250565b6000602082019050612492600083018461223e565b9291505056fea2646970667358221220e39074815d974e0bed90fb5dd6b0df1a2e6ca6e59a5b3423809684bb3a16751964736f6c634300080d0033",
}

// MarketPlaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketPlaceMetaData.ABI instead.
var MarketPlaceABI = MarketPlaceMetaData.ABI

// MarketPlaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketPlaceMetaData.Bin instead.
var MarketPlaceBin = MarketPlaceMetaData.Bin

// DeployMarketPlace deploys a new Ethereum contract, binding an instance of MarketPlace to it.
func DeployMarketPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketPlace, error) {
	parsed, err := MarketPlaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketPlaceBin), backend)
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

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x0e337462.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) BurnZcTokenRemovingNotionalCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "burnZcTokenRemovingNotionalCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x0e337462.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotionalCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotionalCalled is a free data retrieval call binding the contract method 0x0e337462.
//
// Solidity: function burnZcTokenRemovingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) BurnZcTokenRemovingNotionalCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAndAdapterAddressCalled is a free data retrieval call binding the contract method 0x6ba14fe7.
//
// Solidity: function cTokenAndAdapterAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CTokenAndAdapterAddressCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAndAdapterAddressCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CTokenAndAdapterAddressCalled is a free data retrieval call binding the contract method 0x6ba14fe7.
//
// Solidity: function cTokenAndAdapterAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CTokenAndAdapterAddressCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAndAdapterAddressCalled is a free data retrieval call binding the contract method 0x6ba14fe7.
//
// Solidity: function cTokenAndAdapterAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAndAdapterAddressCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialExitCalled is a free data retrieval call binding the contract method 0xde2ad3eb.
//
// Solidity: function custodialExitCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialExitCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialExitCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CustodialExitCalled is a free data retrieval call binding the contract method 0xde2ad3eb.
//
// Solidity: function custodialExitCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialExitCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialExitCalled is a free data retrieval call binding the contract method 0xde2ad3eb.
//
// Solidity: function custodialExitCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialExitCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialExitCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x50e3fe2c.
//
// Solidity: function custodialInitiateCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CustodialInitiateCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "custodialInitiateCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x50e3fe2c.
//
// Solidity: function custodialInitiateCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CustodialInitiateCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// CustodialInitiateCalled is a free data retrieval call binding the contract method 0x50e3fe2c.
//
// Solidity: function custodialInitiateCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CustodialInitiateCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CustodialInitiateCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x0f5ea3c7.
//
// Solidity: function mintZcTokenAddingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) MintZcTokenAddingNotionalCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "mintZcTokenAddingNotionalCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x0f5ea3c7.
//
// Solidity: function mintZcTokenAddingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotionalCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// MintZcTokenAddingNotionalCalled is a free data retrieval call binding the contract method 0x0f5ea3c7.
//
// Solidity: function mintZcTokenAddingNotionalCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) MintZcTokenAddingNotionalCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotionalCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xe3e43c58.
//
// Solidity: function p2pVaultExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pVaultExchangeCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pVaultExchangeCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xe3e43c58.
//
// Solidity: function p2pVaultExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchangeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pVaultExchangeCalled is a free data retrieval call binding the contract method 0xe3e43c58.
//
// Solidity: function p2pVaultExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pVaultExchangeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.P2pVaultExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xd721b5d5.
//
// Solidity: function p2pZcTokenExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) P2pZcTokenExchangeCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "p2pZcTokenExchangeCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xd721b5d5.
//
// Solidity: function p2pZcTokenExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchangeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// P2pZcTokenExchangeCalled is a free data retrieval call binding the contract method 0xd721b5d5.
//
// Solidity: function p2pZcTokenExchangeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) P2pZcTokenExchangeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.P2pZcTokenExchangeCalled(&_MarketPlace.CallOpts, arg0)
}

// RedeemVaultInterestCalled is a free data retrieval call binding the contract method 0x26817c77.
//
// Solidity: function redeemVaultInterestCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) RedeemVaultInterestCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "redeemVaultInterestCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RedeemVaultInterestCalled is a free data retrieval call binding the contract method 0x26817c77.
//
// Solidity: function redeemVaultInterestCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) RedeemVaultInterestCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.RedeemVaultInterestCalled(&_MarketPlace.CallOpts, arg0)
}

// RedeemVaultInterestCalled is a free data retrieval call binding the contract method 0x26817c77.
//
// Solidity: function redeemVaultInterestCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) RedeemVaultInterestCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.RedeemVaultInterestCalled(&_MarketPlace.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0x97c2c691.
//
// Solidity: function redeemZcTokenCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) RedeemZcTokenCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "redeemZcTokenCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0x97c2c691.
//
// Solidity: function redeemZcTokenCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) RedeemZcTokenCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.RedeemZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// RedeemZcTokenCalled is a free data retrieval call binding the contract method 0x97c2c691.
//
// Solidity: function redeemZcTokenCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) RedeemZcTokenCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.RedeemZcTokenCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x27d3aae2.
//
// Solidity: function transferVaultNotionalFeeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) TransferVaultNotionalFeeCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "transferVaultNotionalFeeCalled", arg0)

	outstruct := new(struct {
		Underlying common.Address
		Maturity   *big.Int
		One        common.Address
		Two        common.Address
		Amount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Underlying = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Maturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.One = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Two = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x27d3aae2.
//
// Solidity: function transferVaultNotionalFeeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFeeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeCalled(&_MarketPlace.CallOpts, arg0)
}

// TransferVaultNotionalFeeCalled is a free data retrieval call binding the contract method 0x27d3aae2.
//
// Solidity: function transferVaultNotionalFeeCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) TransferVaultNotionalFeeCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFeeCalled(&_MarketPlace.CallOpts, arg0)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) BurnZcTokenRemovingNotional(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "burnZcTokenRemovingNotional", p, u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) BurnZcTokenRemovingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// BurnZcTokenRemovingNotional is a paid mutator transaction binding the contract method 0x87e157c1.
//
// Solidity: function burnZcTokenRemovingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) BurnZcTokenRemovingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.BurnZcTokenRemovingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
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

// CTokenAndAdapterAddress is a paid mutator transaction binding the contract method 0x2512ca63.
//
// Solidity: function cTokenAndAdapterAddress(uint8 p, address u, uint256 m) returns(address, address)
func (_MarketPlace *MarketPlaceTransactor) CTokenAndAdapterAddress(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAndAdapterAddress", p, u, m)
}

// CTokenAndAdapterAddress is a paid mutator transaction binding the contract method 0x2512ca63.
//
// Solidity: function cTokenAndAdapterAddress(uint8 p, address u, uint256 m) returns(address, address)
func (_MarketPlace *MarketPlaceSession) CTokenAndAdapterAddress(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddress(&_MarketPlace.TransactOpts, p, u, m)
}

// CTokenAndAdapterAddress is a paid mutator transaction binding the contract method 0x2512ca63.
//
// Solidity: function cTokenAndAdapterAddress(uint8 p, address u, uint256 m) returns(address, address)
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAndAdapterAddress(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddress(&_MarketPlace.TransactOpts, p, u, m)
}

// CTokenAndAdapterAddressReturns is a paid mutator transaction binding the contract method 0x153d4625.
//
// Solidity: function cTokenAndAdapterAddressReturns(address c, address a) returns()
func (_MarketPlace *MarketPlaceTransactor) CTokenAndAdapterAddressReturns(opts *bind.TransactOpts, c common.Address, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAndAdapterAddressReturns", c, a)
}

// CTokenAndAdapterAddressReturns is a paid mutator transaction binding the contract method 0x153d4625.
//
// Solidity: function cTokenAndAdapterAddressReturns(address c, address a) returns()
func (_MarketPlace *MarketPlaceSession) CTokenAndAdapterAddressReturns(c common.Address, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddressReturns(&_MarketPlace.TransactOpts, c, a)
}

// CTokenAndAdapterAddressReturns is a paid mutator transaction binding the contract method 0x153d4625.
//
// Solidity: function cTokenAndAdapterAddressReturns(address c, address a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAndAdapterAddressReturns(c common.Address, a common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAndAdapterAddressReturns(&_MarketPlace.TransactOpts, c, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialExit(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialExit", p, u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialExit(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
}

// CustodialExit is a paid mutator transaction binding the contract method 0x0f0016b6.
//
// Solidity: function custodialExit(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialExit(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialExit(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
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

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) CustodialInitiate(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "custodialInitiate", p, u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) CustodialInitiate(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
}

// CustodialInitiate is a paid mutator transaction binding the contract method 0xc06760c7.
//
// Solidity: function custodialInitiate(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) CustodialInitiate(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CustodialInitiate(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
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

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) MintZcTokenAddingNotional(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "mintZcTokenAddingNotional", p, u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) MintZcTokenAddingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// MintZcTokenAddingNotional is a paid mutator transaction binding the contract method 0x01cc6448.
//
// Solidity: function mintZcTokenAddingNotional(uint8 p, address u, uint256 m, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) MintZcTokenAddingNotional(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MintZcTokenAddingNotional(&_MarketPlace.TransactOpts, p, u, m, t, a)
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

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pVaultExchange(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pVaultExchange", p, u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pVaultExchange(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
}

// P2pVaultExchange is a paid mutator transaction binding the contract method 0x15042ddf.
//
// Solidity: function p2pVaultExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pVaultExchange(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pVaultExchange(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
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

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) P2pZcTokenExchange(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "p2pZcTokenExchange", p, u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) P2pZcTokenExchange(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
}

// P2pZcTokenExchange is a paid mutator transaction binding the contract method 0xfcbaab2e.
//
// Solidity: function p2pZcTokenExchange(uint8 p, address u, uint256 m, address o, address t, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) P2pZcTokenExchange(p uint8, u common.Address, m *big.Int, o common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.P2pZcTokenExchange(&_MarketPlace.TransactOpts, p, u, m, o, t, a)
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

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) RedeemVaultInterest(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemVaultInterest", p, u, m, t)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterest(&_MarketPlace.TransactOpts, p, u, m, t)
}

// RedeemVaultInterest is a paid mutator transaction binding the contract method 0x3a660bd8.
//
// Solidity: function redeemVaultInterest(uint8 p, address u, uint256 m, address t) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemVaultInterest(p uint8, u common.Address, m *big.Int, t common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterest(&_MarketPlace.TransactOpts, p, u, m, t)
}

// RedeemVaultInterestReturns is a paid mutator transaction binding the contract method 0x295b25f2.
//
// Solidity: function redeemVaultInterestReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactor) RedeemVaultInterestReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemVaultInterestReturns", a)
}

// RedeemVaultInterestReturns is a paid mutator transaction binding the contract method 0x295b25f2.
//
// Solidity: function redeemVaultInterestReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceSession) RedeemVaultInterestReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterestReturns(&_MarketPlace.TransactOpts, a)
}

// RedeemVaultInterestReturns is a paid mutator transaction binding the contract method 0x295b25f2.
//
// Solidity: function redeemVaultInterestReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) RedeemVaultInterestReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemVaultInterestReturns(&_MarketPlace.TransactOpts, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) RedeemZcToken(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemZcToken", p, u, m, t, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// RedeemZcToken is a paid mutator transaction binding the contract method 0x9f6eddc4.
//
// Solidity: function redeemZcToken(uint8 p, address u, uint256 m, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) RedeemZcToken(p uint8, u common.Address, m *big.Int, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcToken(&_MarketPlace.TransactOpts, p, u, m, t, a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0x6ac91033.
//
// Solidity: function redeemZcTokenReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactor) RedeemZcTokenReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "redeemZcTokenReturns", a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0x6ac91033.
//
// Solidity: function redeemZcTokenReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceSession) RedeemZcTokenReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcTokenReturns(&_MarketPlace.TransactOpts, a)
}

// RedeemZcTokenReturns is a paid mutator transaction binding the contract method 0x6ac91033.
//
// Solidity: function redeemZcTokenReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) RedeemZcTokenReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.RedeemZcTokenReturns(&_MarketPlace.TransactOpts, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactor) TransferVaultNotionalFee(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "transferVaultNotionalFee", p, u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceSession) TransferVaultNotionalFee(p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, p, u, m, f, a)
}

// TransferVaultNotionalFee is a paid mutator transaction binding the contract method 0xdb850901.
//
// Solidity: function transferVaultNotionalFee(uint8 p, address u, uint256 m, address f, uint256 a) returns(bool)
func (_MarketPlace *MarketPlaceTransactorSession) TransferVaultNotionalFee(p uint8, u common.Address, m *big.Int, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.TransferVaultNotionalFee(&_MarketPlace.TransactOpts, p, u, m, f, a)
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
