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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"authRedeemCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"authRedeemReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burnZcTokenRemovingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"burnZcTokenRemovingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"burnZcTokenRemovingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"}],\"name\":\"cTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"cTokenAddressCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"name\":\"cTokenAddressReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"createMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialExitCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialExitReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"custodialInitiate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"custodialInitiateCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"custodialInitiateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"exchangeRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultTracker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturityRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"maturityRateReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mintZcTokenAddingNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"mintZcTokenAddingNotionalCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"mintZcTokenAddingNotionalReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pVaultExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pVaultExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pVaultExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"p2pZcTokenExchange\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"p2pZcTokenExchangeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"p2pZcTokenExchangeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"redeemVaultInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemVaultInterestCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemVaultInterestReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"redeemZcTokenCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"redeemZcTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferVaultNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"transferVaultNotionalFeeCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"one\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"two\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferVaultNotionalFeeReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612c6f806100206000396000f3fe608060405234801561001057600080fd5b506004361061023d5760003560e01c80635755d7631161013b578063beaff41e116100b8578063dc098af91161007c578063dc098af9146107c5578063de2ad3eb146107e1578063e3e43c5814610815578063f7de8b1f14610849578063fcbaab2e1461087a5761023d565b8063beaff41e146106f9578063c06760c714610715578063d557ee8514610745578063d721b5d514610761578063db850901146107955761023d565b806376a94483116100ff57806376a94483146106155780637d04cd151461064957806387e157c11461066557806397c2c691146106955780639f6eddc4146106c95761023d565b80635755d7631461055d578063589aeec71461058d5780635d425f86146105a95780636ac91033146105c55780636dd086cc146105e15761023d565b806327d3aae2116101c95780634337ec161161018d5780634337ec16146104a557806349be053c146104c15780634bc81e09146104dd57806350e3fe2c146104f957806352bc94301461052d5761023d565b806327d3aae2146103c2578063295b25f2146103f6578063305a21bf1461041257806335bdafab146104455780633a660bd8146104755761023d565b80630f5ea3c7116102105780630f5ea3c7146102f257806315042ddf146103265780631cd9be91146103565780632634de191461037257806326817c771461038e5761023d565b806301cc64481461024257806304aa1dfd146102725780630e3374621461028e5780630f0016b6146102c2575b600080fd5b61025c6004803603810190610257919061276a565b6108aa565b6040516102699190612800565b60405180910390f35b61028c60048036038101906102879190612847565b610a5a565b005b6102a860048036038101906102a39190612874565b610a77565b6040516102b99594939291906128bf565b60405180910390f35b6102dc60048036038101906102d79190612912565b610b0d565b6040516102e99190612800565b60405180910390f35b61030c60048036038101906103079190612874565b610cf6565b60405161031d9594939291906128bf565b60405180910390f35b610340600480360381019061033b9190612912565b610d8c565b60405161034d9190612800565b60405180910390f35b610370600480360381019061036b9190612847565b610f75565b005b61038c60048036038101906103879190612847565b610f92565b005b6103a860048036038101906103a39190612874565b610faf565b6040516103b99594939291906128bf565b60405180910390f35b6103dc60048036038101906103d79190612874565b611045565b6040516103ed9594939291906128bf565b60405180910390f35b610410600480360381019061040b919061299f565b6110db565b005b61042c600480360381019061042791906129cc565b6110e5565b60405161043c9493929190612a1f565b60405180910390f35b61045f600480360381019061045a91906129cc565b61118f565b60405161046c9190612a64565b60405180910390f35b61048f600480360381019061048a9190612a7f565b61130e565b60405161049c9190612ae6565b60405180910390f35b6104bf60048036038101906104ba9190612b01565b6114a6565b005b6104db60048036038101906104d6919061299f565b61165a565b005b6104f760048036038101906104f29190612847565b611664565b005b610513600480360381019061050e9190612874565b611681565b6040516105249594939291906128bf565b60405180910390f35b61054760048036038101906105429190612912565b611717565b6040516105549190612ae6565b60405180910390f35b61057760048036038101906105729190612ba3565b6118f3565b6040516105849190612ae6565b60405180910390f35b6105a760048036038101906105a2919061299f565b611900565b005b6105c360048036038101906105be919061299f565b61190a565b005b6105df60048036038101906105da919061299f565b611914565b005b6105fb60048036038101906105f69190612874565b61191e565b60405161060c9594939291906128bf565b60405180910390f35b61062f600480360381019061062a9190612874565b6119b4565b6040516106409594939291906128bf565b60405180910390f35b610663600480360381019061065e9190612847565b611a4a565b005b61067f600480360381019061067a919061276a565b611a67565b60405161068c9190612800565b60405180910390f35b6106af60048036038101906106aa9190612874565b611c17565b6040516106c09594939291906128bf565b60405180910390f35b6106e360048036038101906106de919061276a565b611cad565b6040516106f09190612ae6565b60405180910390f35b610713600480360381019061070e9190612847565b611e50565b005b61072f600480360381019061072a9190612912565b611e6d565b60405161073c9190612800565b60405180910390f35b61075f600480360381019061075a9190612be3565b612056565b005b61077b60048036038101906107769190612874565b61209a565b60405161078c9594939291906128bf565b60405180910390f35b6107af60048036038101906107aa919061276a565b612130565b6040516107bc9190612800565b60405180910390f35b6107df60048036038101906107da9190612847565b6122e0565b005b6107fb60048036038101906107f69190612874565b6122fd565b60405161080c9594939291906128bf565b60405180910390f35b61082f600480360381019061082a9190612874565b612393565b6040516108409594939291906128bf565b60405180910390f35b610863600480360381019061085e91906129cc565b612429565b604051610871929190612c10565b60405180910390f35b610894600480360381019061088f9190612912565b61243e565b6040516108a19190612800565b60405180910390f35b60006108b4612627565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600660008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60189054906101000a900460ff1691505095945050505050565b80600c601a6101000a81548160ff02191690831515021790555050565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000610b17612627565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600360008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60159054906101000a900460ff169150509695505050505050565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000610d96612627565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600560008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60179054906101000a900460ff169150509695505050505050565b80600c60156101000a81548160ff02191690831515021790555050565b80600c60146101000a81548160ff02191690831515021790555050565b600b6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60086020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600f8190555050565b600060205282600052604060002060205281600052604060002060205280600052604060002060009250925050508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b6000611199612627565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600160008760ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169150509392505050565b6000611318612627565b84816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381602001818152505082816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080600b60008860ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600f54915050949350505050565b60405180608001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001828152506000808960ff1660ff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600087815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015590505050505050505050565b80600e8190555050565b80600c60186101000a81548160ff02191690831515021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000611721612627565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600a60008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600e549150509695505050505050565b6000601054905092915050565b8060108190555050565b8060118190555050565b80600d8190555050565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b80600c60176101000a81548160ff02191690831515021790555050565b6000611a71612627565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600760008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60199054906101000a900460ff1691505095945050505050565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b6000611cb7612627565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600960008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600d5491505095945050505050565b80600c60166101000a81548160ff02191690831515021790555050565b6000611e77612627565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600260008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60149054906101000a900460ff169150509695505050505050565b80600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60046020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b600061213a612627565b85816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508481602001818152505083816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600860008960ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c601a9054906101000a900460ff1691505095945050505050565b80600c60196101000a81548160ff02191690831515021790555050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b60008060115460105491509150935093915050565b6000612448612627565b86816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581602001818152505084816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281608001818152505080600460008a60ff1660ff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040155905050600c60169054906101000a900460ff169150509695505050505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600060ff82169050919050565b6126b38161269d565b81146126be57600080fd5b50565b6000813590506126d0816126aa565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000612701826126d6565b9050919050565b612711816126f6565b811461271c57600080fd5b50565b60008135905061272e81612708565b92915050565b6000819050919050565b61274781612734565b811461275257600080fd5b50565b6000813590506127648161273e565b92915050565b600080600080600060a0868803121561278657612785612698565b5b6000612794888289016126c1565b95505060206127a58882890161271f565b94505060406127b688828901612755565b93505060606127c78882890161271f565b92505060806127d888828901612755565b9150509295509295909350565b60008115159050919050565b6127fa816127e5565b82525050565b600060208201905061281560008301846127f1565b92915050565b612824816127e5565b811461282f57600080fd5b50565b6000813590506128418161281b565b92915050565b60006020828403121561285d5761285c612698565b5b600061286b84828501612832565b91505092915050565b60006020828403121561288a57612889612698565b5b6000612898848285016126c1565b91505092915050565b6128aa816126f6565b82525050565b6128b981612734565b82525050565b600060a0820190506128d460008301886128a1565b6128e160208301876128b0565b6128ee60408301866128a1565b6128fb60608301856128a1565b61290860808301846128b0565b9695505050505050565b60008060008060008060c0878903121561292f5761292e612698565b5b600061293d89828a016126c1565b965050602061294e89828a0161271f565b955050604061295f89828a01612755565b945050606061297089828a0161271f565b935050608061298189828a0161271f565b92505060a061299289828a01612755565b9150509295509295509295565b6000602082840312156129b5576129b4612698565b5b60006129c384828501612755565b91505092915050565b6000806000606084860312156129e5576129e4612698565b5b60006129f3868287016126c1565b9350506020612a048682870161271f565b9250506040612a1586828701612755565b9150509250925092565b6000608082019050612a3460008301876128a1565b612a4160208301866128a1565b612a4e60408301856128a1565b612a5b60608301846128b0565b95945050505050565b6000602082019050612a7960008301846128a1565b92915050565b60008060008060808587031215612a9957612a98612698565b5b6000612aa7878288016126c1565b9450506020612ab88782880161271f565b9350506040612ac987828801612755565b9250506060612ada8782880161271f565b91505092959194509250565b6000602082019050612afb60008301846128b0565b92915050565b600080600080600080600060e0888a031215612b2057612b1f612698565b5b6000612b2e8a828b016126c1565b9750506020612b3f8a828b0161271f565b9650506040612b508a828b01612755565b9550506060612b618a828b0161271f565b9450506080612b728a828b0161271f565b93505060a0612b838a828b0161271f565b92505060c0612b948a828b01612755565b91505092959891949750929550565b60008060408385031215612bba57612bb9612698565b5b6000612bc8858286016126c1565b9250506020612bd98582860161271f565b9150509250929050565b600060208284031215612bf957612bf8612698565b5b6000612c078482850161271f565b91505092915050565b6000604082019050612c2560008301856128b0565b612c3260208301846128b0565b939250505056fea264697066735822122025a531fed757d41747f2160a781b2bcfbccaf71f3fa0a434bbdc6823966b5a3864736f6c63430008100033",
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

// Rates is a free data retrieval call binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 , address , uint256 ) view returns(uint256, uint256)
func (_MarketPlace *MarketPlaceCaller) Rates(opts *bind.CallOpts, arg0 uint8, arg1 common.Address, arg2 *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MarketPlace.contract.Call(opts, &out, "rates", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Rates is a free data retrieval call binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 , address , uint256 ) view returns(uint256, uint256)
func (_MarketPlace *MarketPlaceSession) Rates(arg0 uint8, arg1 common.Address, arg2 *big.Int) (*big.Int, *big.Int, error) {
	return _MarketPlace.Contract.Rates(&_MarketPlace.CallOpts, arg0, arg1, arg2)
}

// Rates is a free data retrieval call binding the contract method 0xf7de8b1f.
//
// Solidity: function rates(uint8 , address , uint256 ) view returns(uint256, uint256)
func (_MarketPlace *MarketPlaceCallerSession) Rates(arg0 uint8, arg1 common.Address, arg2 *big.Int) (*big.Int, *big.Int, error) {
	return _MarketPlace.Contract.Rates(&_MarketPlace.CallOpts, arg0, arg1, arg2)
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

// MaturityRateReturns is a paid mutator transaction binding the contract method 0x5d425f86.
//
// Solidity: function maturityRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactor) MaturityRateReturns(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.contract.Transact(opts, "maturityRateReturns", a)
}

// MaturityRateReturns is a paid mutator transaction binding the contract method 0x5d425f86.
//
// Solidity: function maturityRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceSession) MaturityRateReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MaturityRateReturns(&_MarketPlace.TransactOpts, a)
}

// MaturityRateReturns is a paid mutator transaction binding the contract method 0x5d425f86.
//
// Solidity: function maturityRateReturns(uint256 a) returns()
func (_MarketPlace *MarketPlaceTransactorSession) MaturityRateReturns(a *big.Int) (*types.Transaction, error) {
	return _MarketPlace.Contract.MaturityRateReturns(&_MarketPlace.TransactOpts, a)
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
