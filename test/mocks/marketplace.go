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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"authRedeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeemReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"exchangeRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultTracker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"redeemVaultInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemVaultInterestCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemVaultInterestReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemZcTokenCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"transferVaultNotionalFeeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferVaultNotionalFeeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612bc4806100206000396000f3fe608060405234801561001057600080fd5b50600436106102275760003560e01c806352bc9430116101305780639f6eddc4116100b8578063db8509011161007c578063db85090114610763578063dc098af914610793578063de2ad3eb146107af578063e3e43c58146107e3578063fcbaab2e1461081757610227565b80639f6eddc414610697578063beaff41e146106c7578063c06760c7146106e3578063d557ee8514610713578063d721b5d51461072f57610227565b80636dd086cc116100ff5780636dd086cc146105af57806376a94483146105e35780637d04cd151461061757806387e157c11461063357806397c2c6911461066357610227565b806352bc9430146105175780635755d76314610547578063589aeec7146105775780636ac910331461059357610227565b806327d3aae2116101b35780633a660bd8116101825780633a660bd81461045f5780634337ec161461048f57806349be053c146104ab5780634bc81e09146104c757806350e3fe2c146104e357610227565b806327d3aae2146103ac578063295b25f2146103e0578063305a21bf146103fc57806335bdafab1461042f57610227565b80630f5ea3c7116101fa5780630f5ea3c7146102dc57806315042ddf146103105780631cd9be91146103405780632634de191461035c57806326817c771461037857610227565b806301cc64481461022c57806304aa1dfd1461025c5780630e337462146102785780630f0016b6146102ac575b600080fd5b610246600480360381019061024191906126e8565b610847565b604051610253919061277e565b60405180910390f35b610276600480360381019061027191906127c5565b6109f7565b005b610292600480360381019061028d91906127f2565b610a14565b6040516102a395949392919061283d565b60405180910390f35b6102c660048036038101906102c19190612890565b610aaa565b6040516102d3919061277e565b60405180910390f35b6102f660048036038101906102f191906127f2565b610c93565b60405161030795949392919061283d565b60405180910390f35b61032a60048036038101906103259190612890565b610d29565b604051610337919061277e565b60405180910390f35b61035a600480360381019061035591906127c5565b610f12565b005b610376600480360381019061037191906127c5565b610f2f565b005b610392600480360381019061038d91906127f2565b610f4c565b6040516103a395949392919061283d565b60405180910390f35b6103c660048036038101906103c191906127f2565b610fe2565b6040516103d795949392919061283d565b60405180910390f35b6103fa60048036038101906103f5919061291d565b611078565b005b6104166004803603810190610411919061294a565b611082565b604051610426949392919061299d565b60405180910390f35b6104496004803603810190610444919061294a565b61112c565b60405161045691906129e2565b60405180910390f35b610479600480360381019061047491906129fd565b6112ab565b6040516104869190612a64565b60405180910390f35b6104a960048036038101906104a49190612a7f565b611443565b005b6104c560048036038101906104c0919061291d565b6115f7565b005b6104e160048036038101906104dc91906127c5565b611601565b005b6104fd60048036038101906104f891906127f2565b61161e565b60405161050e95949392919061283d565b60405180910390f35b610531600480360381019061052c9190612890565b6116b4565b60405161053e9190612a64565b60405180910390f35b610561600480360381019061055c9190612b21565b611890565b60405161056e9190612a64565b60405180910390f35b610591600480360381019061058c919061291d565b61189d565b005b6105ad60048036038101906105a8919061291d565b6118a7565b005b6105c960048036038101906105c491906127f2565b6118b1565b6040516105da95949392919061283d565b60405180910390f35b6105fd60048036038101906105f891906127f2565b611947565b60405161060e95949392919061283d565b60405180910390f35b610631600480360381019061062c91906127c5565b6119dd565b005b61064d600480360381019061064891906126e8565b6119fa565b60405161065a919061277e565b60405180910390f35b61067d600480360381019061067891906127f2565b611baa565b60405161068e95949392919061283d565b60405180910390f35b6106b160048036038101906106ac91906126e8565b611c40565b6040516106be9190612a64565b60405180910390f35b6106e160048036038101906106dc91906127c5565b611de3565b005b6106fd60048036038101906106f89190612890565b611e00565b60405161070a919061277e565b60405180910390f35b61072d60048036038101906107289190612b61565b611fe9565b005b610749600480360381019061074491906127f2565b61202d565b60405161075a95949392919061283d565b60405180910390f35b61077d600480360381019061077891906126e8565b6120c3565b60405161078a919061277e565b60405180910390f35b6107ad60048036038101906107a891906127c5565b612273565b005b6107c960048036038101906107c491906127f2565b612290565b6040516107da95949392919061283d565b60405180910390f35b6107fd60048036038101906107f891906127f2565b612326565b60405161080e95949392919061283d565b60405180910390f35b610831600480360381019061082c9190612890565b6123bc565b60405161083e919061277e565b60405180910390f35b60006108516125a5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600660008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60189054906101000a900460ff1691505095945050505050565b80600c601a6101000a81548160ff02191690831515021790555050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000610ab46125a5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600360008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60159054906101000a900460ff169150509695505050505050565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000610d336125a5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600560008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60179054906101000a900460ff169150509695505050505050565b80600c60156101000a81548160ff02191690831515021790555050565b80600c60146101000a81548160ff02191690831515021790555050565b600b6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60086020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600f8190555050565b600060205282600052604060002060205281600052604060002060205280600052604060002060009250925050508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60006111366125a5565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600160008760ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150509392505050565b60006112b56125a5565b84816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381602001818152505082816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600b60008860ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600f54915050949350505050565b60405180608001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152506000808960ff1660ff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505050505050505050565b80600e8190555050565b80600c60186101000a81548160ff02191690831515021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60006116be6125a5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600a60008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600e549150509695505050505050565b6000601054905092915050565b8060108190555050565b80600d8190555050565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600c60176101000a81548160ff02191690831515021790555050565b6000611a046125a5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600760008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60199054906101000a900460ff1691505095945050505050565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000611c4a6125a5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600960008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600d5491505095945050505050565b80600c60166101000a81548160ff02191690831515021790555050565b6000611e0a6125a5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600260008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60149054906101000a900460ff169150509695505050505050565b80600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60006120cd6125a5565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600860008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c601a9054906101000a900460ff1691505095945050505050565b80600c60196101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60006123c66125a5565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600460008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60169054906101000a900460ff169150509695505050505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600060ff82169050919050565b6126318161261b565b811461263c57600080fd5b50565b60008135905061264e81612628565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061267f82612654565b9050919050565b61268f81612674565b811461269a57600080fd5b50565b6000813590506126ac81612686565b92915050565b6000819050919050565b6126c5816126b2565b81146126d057600080fd5b50565b6000813590506126e2816126bc565b92915050565b600080600080600060a0868803121561270457612703612616565b5b60006127128882890161263f565b95505060206127238882890161269d565b9450506040612734888289016126d3565b93505060606127458882890161269d565b9250506080612756888289016126d3565b9150509295509295909350565b60008115159050919050565b61277881612763565b82525050565b6000602082019050612793600083018461276f565b92915050565b6127a281612763565b81146127ad57600080fd5b50565b6000813590506127bf81612799565b92915050565b6000602082840312156127db576127da612616565b5b60006127e9848285016127b0565b91505092915050565b60006020828403121561280857612807612616565b5b60006128168482850161263f565b91505092915050565b61282881612674565b82525050565b612837816126b2565b82525050565b600060a082019050612852600083018861281f565b61285f602083018761282e565b61286c604083018661281f565b612879606083018561281f565b612886608083018461282e565b9695505050505050565b60008060008060008060c087890312156128ad576128ac612616565b5b60006128bb89828a0161263f565b96505060206128cc89828a0161269d565b95505060406128dd89828a016126d3565b94505060606128ee89828a0161269d565b93505060806128ff89828a0161269d565b92505060a061291089828a016126d3565b9150509295509295509295565b60006020828403121561293357612932612616565b5b6000612941848285016126d3565b91505092915050565b60008060006060848603121561296357612962612616565b5b60006129718682870161263f565b93505060206129828682870161269d565b9250506040612993868287016126d3565b9150509250925092565b60006080820190506129b2600083018761281f565b6129bf602083018661281f565b6129cc604083018561281f565b6129d9606083018461282e565b95945050505050565b60006020820190506129f7600083018461281f565b92915050565b60008060008060808587031215612a1757612a16612616565b5b6000612a258782880161263f565b9450506020612a368782880161269d565b9350506040612a47878288016126d3565b9250506060612a588782880161269d565b91505092959194509250565b6000602082019050612a79600083018461282e565b92915050565b600080600080600080600060e0888a031215612a9e57612a9d612616565b5b6000612aac8a828b0161263f565b9750506020612abd8a828b0161269d565b9650506040612ace8a828b016126d3565b9550506060612adf8a828b0161269d565b9450506080612af08a828b0161269d565b93505060a0612b018a828b0161269d565b92505060c0612b128a828b016126d3565b91505092959891949750929550565b60008060408385031215612b3857612b37612616565b5b6000612b468582860161263f565b9250506020612b578582860161269d565b9150509250929050565b600060208284031215612b7757612b76612616565b5b6000612b858482850161269d565b9150509291505056fea2646970667358221220070074494bf52dd9a4e6986199b8ed60181531f6a7b111acb85cfa73684fd0a164736f6c634300080d0033",
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

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) AuthRedeemCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "authRedeemCalled", arg0)

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

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) AuthRedeemCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.AuthRedeemCalled(&_MarketPlace.CallOpts, arg0)
}

// AuthRedeemCalled is a free data retrieval call binding the contract method 0x6dd086cc.
//
// Solidity: function authRedeemCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) AuthRedeemCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.AuthRedeemCalled(&_MarketPlace.CallOpts, arg0)
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

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x76a94483.
//
// Solidity: function cTokenAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCaller) CTokenAddressCalled(opts *bind.CallOpts, arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "cTokenAddressCalled", arg0)

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

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x76a94483.
//
// Solidity: function cTokenAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceSession) CTokenAddressCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
}

// CTokenAddressCalled is a free data retrieval call binding the contract method 0x76a94483.
//
// Solidity: function cTokenAddressCalled(uint8 ) view returns(address underlying, uint256 maturity, address one, address two, uint256 amount)
func (_MarketPlace *MarketPlaceCallerSession) CTokenAddressCalled(arg0 uint8) (struct {
	Underlying common.Address
	Maturity   *big.Int
	One        common.Address
	Two        common.Address
	Amount     *big.Int
}, error) {
	return _MarketPlace.Contract.CTokenAddressCalled(&_MarketPlace.CallOpts, arg0)
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

// ExchangeRate is a free data retrieval call binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 , address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCaller) ExchangeRate(opts *bind.CallOpts, arg0 uint8, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "exchangeRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 , address ) view returns(uint256)
func (_MarketPlace *MarketPlaceSession) ExchangeRate(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.ExchangeRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x5755d763.
//
// Solidity: function exchangeRate(uint8 , address ) view returns(uint256)
func (_MarketPlace *MarketPlaceCallerSession) ExchangeRate(arg0 uint8, arg1 common.Address) (*big.Int, error) {
	return _MarketPlace.Contract.ExchangeRate(&_MarketPlace.CallOpts, arg0, arg1)
}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceCaller) Markets(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "markets", arg0, arg1, arg2)

	outstruct := new(struct {
		CTokenAddr   common.Address
		ZcToken      common.Address
		VaultTracker common.Address
		MaturityRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CTokenAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ZcToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VaultTracker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MaturityRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceSession) Markets(arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Markets is a free data retrieval call binding the contract method 0x305a21bf.
//
// Solidity: function markets(uint8 , address , uint256 ) view returns(address cTokenAddr, address zcToken, address vaultTracker, uint256 maturityRate)
func (_MarketPlace *MarketPlaceCallerSession) Markets(arg0 uint8, arg1 common.Address, arg2 *big.Int) (struct {
	CTokenAddr   common.Address
	ZcToken      common.Address
	VaultTracker common.Address
	MaturityRate *big.Int
}, error) {
	return _MarketPlace.Contract.Markets(&_MarketPlace.CallOpts, arg0, arg1, arg2)
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

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactor) AuthRedeem(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "authRedeem", p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeem(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// AuthRedeem is a paid mutator transaction binding the contract method 0x52bc9430.
//
// Solidity: function authRedeem(uint8 p, address u, uint256 m, address f, address t, uint256 a) returns(uint256)
func (_MarketPlace *MarketPlaceTransactorSession) AuthRedeem(p uint8, u common.Address, m *big.Int, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeem(&_MarketPlace.TransactOpts, p, u, m, f, t, a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactor) AuthRedeemReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "authRedeemReturns", a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceSession) AuthRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeemReturns(&_MarketPlace.TransactOpts, a)
}

// AuthRedeemReturns is a paid mutator transaction binding the contract method 0x49be053c.
//
// Solidity: function authRedeemReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) AuthRedeemReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.AuthRedeemReturns(&_MarketPlace.TransactOpts, a)
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

// CTokenAddress is a paid mutator transaction binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactor) CTokenAddress(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddress", p, u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceSession) CTokenAddress(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, p, u, m)
}

// CTokenAddress is a paid mutator transaction binding the contract method 0x35bdafab.
//
// Solidity: function cTokenAddress(uint8 p, address u, uint256 m) returns(address)
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddress(p uint8, u common.Address, m *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddress(&_MarketPlace.TransactOpts, p, u, m)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address c) returns()
func (_MarketPlace *MarketPlaceTransactor) CTokenAddressReturns(opts *bind.TransactOpts, c common.Address) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "cTokenAddressReturns", c)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address c) returns()
func (_MarketPlace *MarketPlaceSession) CTokenAddressReturns(c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, c)
}

// CTokenAddressReturns is a paid mutator transaction binding the contract method 0xd557ee85.
//
// Solidity: function cTokenAddressReturns(address c) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CTokenAddressReturns(c common.Address) (*types.Transaction, error) {
	return _MarketPlace.Contract.CTokenAddressReturns(&_MarketPlace.TransactOpts, c)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x4337ec16.
//
// Solidity: function createMarket(uint8 p, address u, uint256 m, address c, address z, address v, uint256 r) returns()
func (_MarketPlace *MarketPlaceTransactor) CreateMarket(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, c common.Address, z common.Address, v common.Address, r *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "createMarket", p, u, m, c, z, v, r)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x4337ec16.
//
// Solidity: function createMarket(uint8 p, address u, uint256 m, address c, address z, address v, uint256 r) returns()
func (_MarketPlace *MarketPlaceSession) CreateMarket(p uint8, u common.Address, m *big.Int, c common.Address, z common.Address, v common.Address, r *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, p, u, m, c, z, v, r)
}

// CreateMarket is a paid mutator transaction binding the contract method 0x4337ec16.
//
// Solidity: function createMarket(uint8 p, address u, uint256 m, address c, address z, address v, uint256 r) returns()
func (_MarketPlace *MarketPlaceTransactorSession) CreateMarket(p uint8, u common.Address, m *big.Int, c common.Address, z common.Address, v common.Address, r *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.CreateMarket(&_MarketPlace.TransactOpts, p, u, m, c, z, v, r)
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

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactor) ExchangeRateReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "exchangeRateReturns", a)
}

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceSession) ExchangeRateReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExchangeRateReturns(&_MarketPlace.TransactOpts, a)
}

// ExchangeRateReturns is a paid mutator transaction binding the contract method 0x589aeec7.
//
// Solidity: function exchangeRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) ExchangeRateReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.ExchangeRateReturns(&_MarketPlace.TransactOpts, a)
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
