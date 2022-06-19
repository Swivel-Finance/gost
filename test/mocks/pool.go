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

// PoolBurnOutputs is an auto generated low-level Go binding around an user-defined struct.
type PoolBurnOutputs struct {
	TokensBurned     *big.Int
	TokensReturned   *big.Int
	PtTokensReturned *big.Int
}

// PoolMintOutputs is an auto generated low-level Go binding around an user-defined struct.
type PoolMintOutputs struct {
	BaseIn       *big.Int
	PtIn         *big.Int
	TokensMinted *big.Int
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remained\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"PTTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"name\":\"burnForUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burnForUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensMinted\",\"type\":\"uint256\"}],\"internalType\":\"structPool.MintOutputs\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"burnForUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokensBurned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensReturned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ptTokensReturned\",\"type\":\"uint256\"}],\"internalType\":\"structPool.BurnOutputs\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"burnReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"p\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalToken\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyPrincipalTokenCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"PrincipalTokenOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyPrincipalTokenPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"o\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyPrincipalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"buyUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"min\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"p\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"b\",\"type\":\"uint128\"}],\"name\":\"buyUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remained\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"remainder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensMinted\",\"type\":\"uint256\"}],\"internalType\":\"structPool.MintOutputs\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remained\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ptToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"name\":\"mintWithUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintWithUnderlyingCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"remainder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PTToBuy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"baseIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensMinted\",\"type\":\"uint256\"}],\"internalType\":\"structPool.MintOutputs\",\"name\":\"o\",\"type\":\"tuple\"}],\"name\":\"mintWithUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"principalToken\",\"outputs\":[{\"internalType\":\"contractIErc5095\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"principalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"ptReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalToken\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellPrincipalTokenCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellPrincipalTokenPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"i\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellPrincipalTokenReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"m\",\"type\":\"uint128\"}],\"name\":\"sellUnderlying\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sellUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"u\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingPreview\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellUnderlyingPreviewCalled\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingPreviewReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"s\",\"type\":\"uint128\"}],\"name\":\"sellUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e6a806100206000396000f3fe608060405234801561001057600080fd5b506004361061025e5760003560e01c8063b3f1c93d11610146578063db02cd1b116100c3578063e7ba677411610087578063e7ba6774146107c3578063ec92245a146107df578063ee4db57014610810578063f038a50f14610842578063f332ea5e1461085e578063f3ec0d2c1461088e5761025e565b8063db02cd1b1461070a578063dbc162de14610726578063e42b16cb14610744578063e525edce14610762578063e78a9b0c146107935761025e565b8063c59964b31161010a578063c59964b31461063d578063c9c917d714610659578063d7020d0a14610689578063d94073d4146106bb578063d9b03f88146106d95761025e565b8063b3f1c93d1461056e578063b45d8801146105a0578063bba0ad39146105bc578063bfdf149f146105ee578063c100a7fd1461060a5761025e565b8063556dd440116101df57806376a91253116101a357806376a912531461048a578063814c83d7146104ba5780639b9c9fde146104d65780639d92bc2f146104f2578063a57c378214610522578063b03ceb151461053e5761025e565b8063556dd440146103e5578063584c0192146104015780636f307dc3146104325780636fccc08714610450578063757b2b761461046e5761025e565b80633e232091116102265780633e23209114610319578063435bf33114610349578063463a861e146103675780634b266e281461039757806350ed7074146103b35761025e565b80630326b5ed1461026357806308e1cff51461027f5780630a7e546e1461029b578063134df58b146102cb57806315c49330146102fb575b600080fd5b61027d600480360381019061027891906117ff565b6108aa565b005b61029960048036038101906102949190611874565b6108ee565b005b6102b560048036038101906102b09190611874565b61092a565b6040516102c291906118b0565b60405180910390f35b6102e560048036038101906102e091906118cb565b61098b565b6040516102f291906118b0565b60405180910390f35b610303610a2a565b60405161031091906118b0565b60405180910390f35b610333600480360381019061032e91906118cb565b610a4c565b60405161034091906118b0565b60405180910390f35b610351610aeb565b60405161035e91906118b0565b60405180910390f35b610381600480360381019061037c91906117ff565b610b0d565b60405161038e91906118b0565b60405180910390f35b6103b160048036038101906103ac9190611874565b610b3c565b005b6103cd60048036038101906103c89190611941565b610b78565b6040516103dc939291906119cb565b60405180910390f35b6103ff60048036038101906103fa9190611874565b610c7e565b005b61041b60048036038101906104169190611a02565b610cba565b604051610429929190611a55565b60405180910390f35b61043a610d40565b6040516104479190611a8d565b60405180910390f35b610458610d6a565b60405161046591906118b0565b60405180910390f35b61048860048036038101906104839190611874565b610d8c565b005b6104a4600480360381019061049f9190611874565b610dc8565b6040516104b191906118b0565b60405180910390f35b6104d460048036038101906104cf9190611b9d565b610e29565b005b6104f060048036038101906104eb9190611874565b610e50565b005b61050c600480360381019061050791906117ff565b610e8c565b60405161051991906118b0565b60405180910390f35b61053c60048036038101906105379190611c2e565b610ebb565b005b61055860048036038101906105539190611c5b565b610ee2565b60405161056591906118b0565b60405180910390f35b61058860048036038101906105839190611cae565b611005565b604051610597939291906119cb565b60405180910390f35b6105ba60048036038101906105b59190611874565b6110fa565b005b6105d660048036038101906105d191906117ff565b611136565b6040516105e593929190611d15565b60405180910390f35b61060860048036038101906106039190611b9d565b611180565b005b610624600480360381019061061f91906117ff565b6111a7565b6040516106349493929190611d4c565b60405180910390f35b61065760048036038101906106529190611874565b6111f7565b005b610673600480360381019061066e9190611874565b611233565b60405161068091906118b0565b60405180910390f35b6106a3600480360381019061069e9190611cae565b611294565b6040516106b2939291906119cb565b60405180910390f35b6106c3611389565b6040516106d09190611a8d565b60405180910390f35b6106f360048036038101906106ee91906117ff565b6113b3565b604051610701929190611d91565b60405180910390f35b610724600480360381019061071f9190611b9d565b61140f565b005b61072e611436565b60405161073b9190611e19565b60405180910390f35b61074c61145f565b60405161075991906118b0565b60405180910390f35b61077c600480360381019061077791906117ff565b611481565b60405161078a929190611d91565b60405180910390f35b6107ad60048036038101906107a89190611c5b565b6114dd565b6040516107ba91906118b0565b60405180910390f35b6107dd60048036038101906107d891906117ff565b611600565b005b6107f960048036038101906107f491906117ff565b611644565b604051610807929190611a55565b60405180910390f35b61082a600480360381019061082591906117ff565b611668565b60405161083993929190611d15565b60405180910390f35b61085c60048036038101906108579190611874565b6116b2565b005b61087860048036038101906108739190611874565b6116ee565b60405161088591906118b0565b60405180910390f35b6108a860048036038101906108a391906117ff565b61174f565b005b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80600260106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081601760006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600560009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600081601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600360009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b601760009054906101000a90046fffffffffffffffffffffffffffffffff1681565b600081601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600260009054906101000a90046fffffffffffffffffffffffffffffffff16905092915050565b601760109054906101000a90046fffffffffffffffffffffffffffffffff1681565b60126020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600460006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600080600060405180608001604052808873ffffffffffffffffffffffffffffffffffffffff16815260200187815260200186815260200185815250601960008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030155905050600a60000154600a60010154600a60020154925092509250955095509592505050565b80600560106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600080604051806040016040528085815260200184815250601b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015590505060106000015460106001015491509150935093915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601660109054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600260006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081601660106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600460109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b80600a60008201518160000155602082015181600101556040820151816002015590505050565b80600560006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b60146020528060005260406000206000915054906101000a90046fffffffffffffffffffffffffffffffff1681565b80600d60008201518160000155602082015181600101556040820151816002015590505050565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250601560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600360109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b600080600060405180606001604052808773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815250601860008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050506007600001546007600101546007600201549250925092509450945094915050565b80600360106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b601a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b80600760008201518160000155602082015181600101556040820151816002015590505050565b60196020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b80600360006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081601760106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600560109054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b600080600060405180606001604052808773ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815250601a60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050600d60000154600d60010154600d600201549250925092509450945094915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60136020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b80600a60008201518160000155602082015181600101556040820151816002015590505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b601660009054906101000a90046fffffffffffffffffffffffffffffffff1681565b60156020528060005260406000206000915090508060000160009054906101000a90046fffffffffffffffffffffffffffffffff16908060000160109054906101000a90046fffffffffffffffffffffffffffffffff16905082565b60006040518060400160405280846fffffffffffffffffffffffffffffffff168152602001836fffffffffffffffffffffffffffffffff16815250601360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555060208201518160000160106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550905050600260109054906101000a90046fffffffffffffffffffffffffffffffff1690509392505050565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b601b6020528060005260406000206000915090508060000154908060010154905082565b60186020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b80600460106101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555050565b600081601660006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550600460009054906101000a90046fffffffffffffffffffffffffffffffff169050919050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000604051905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006117cc826117a1565b9050919050565b6117dc816117c1565b81146117e757600080fd5b50565b6000813590506117f9816117d3565b92915050565b6000602082840312156118155761181461179c565b5b6000611823848285016117ea565b91505092915050565b60006fffffffffffffffffffffffffffffffff82169050919050565b6118518161182c565b811461185c57600080fd5b50565b60008135905061186e81611848565b92915050565b60006020828403121561188a5761188961179c565b5b60006118988482850161185f565b91505092915050565b6118aa8161182c565b82525050565b60006020820190506118c560008301846118a1565b92915050565b600080604083850312156118e2576118e161179c565b5b60006118f0858286016117ea565b92505060206119018582860161185f565b9150509250929050565b6000819050919050565b61191e8161190b565b811461192957600080fd5b50565b60008135905061193b81611915565b92915050565b600080600080600060a0868803121561195d5761195c61179c565b5b600061196b888289016117ea565b955050602061197c888289016117ea565b945050604061198d8882890161192c565b935050606061199e8882890161192c565b92505060806119af8882890161192c565b9150509295509295909350565b6119c58161190b565b82525050565b60006060820190506119e060008301866119bc565b6119ed60208301856119bc565b6119fa60408301846119bc565b949350505050565b600080600060608486031215611a1b57611a1a61179c565b5b6000611a29868287016117ea565b9350506020611a3a8682870161192c565b9250506040611a4b8682870161192c565b9150509250925092565b6000604082019050611a6a60008301856119bc565b611a7760208301846119bc565b9392505050565b611a87816117c1565b82525050565b6000602082019050611aa26000830184611a7e565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611af682611aad565b810181811067ffffffffffffffff82111715611b1557611b14611abe565b5b80604052505050565b6000611b28611792565b9050611b348282611aed565b919050565b600060608284031215611b4f57611b4e611aa8565b5b611b596060611b1e565b90506000611b698482850161192c565b6000830152506020611b7d8482850161192c565b6020830152506040611b918482850161192c565b60408301525092915050565b600060608284031215611bb357611bb261179c565b5b6000611bc184828501611b39565b91505092915050565b600060608284031215611be057611bdf611aa8565b5b611bea6060611b1e565b90506000611bfa8482850161192c565b6000830152506020611c0e8482850161192c565b6020830152506040611c228482850161192c565b60408301525092915050565b600060608284031215611c4457611c4361179c565b5b6000611c5284828501611bca565b91505092915050565b600080600060608486031215611c7457611c7361179c565b5b6000611c82868287016117ea565b9350506020611c938682870161185f565b9250506040611ca48682870161185f565b9150509250925092565b60008060008060808587031215611cc857611cc761179c565b5b6000611cd6878288016117ea565b9450506020611ce7878288016117ea565b9350506040611cf88782880161192c565b9250506060611d098782880161192c565b91505092959194509250565b6000606082019050611d2a6000830186611a7e565b611d3760208301856119bc565b611d4460408301846119bc565b949350505050565b6000608082019050611d616000830187611a7e565b611d6e60208301866119bc565b611d7b60408301856119bc565b611d8860608301846119bc565b95945050505050565b6000604082019050611da660008301856118a1565b611db360208301846118a1565b9392505050565b6000819050919050565b6000611ddf611dda611dd5846117a1565b611dba565b6117a1565b9050919050565b6000611df182611dc4565b9050919050565b6000611e0382611de6565b9050919050565b611e1381611df8565b82525050565b6000602082019050611e2e6000830184611e0a565b9291505056fea26469706673582212208f23b8f4b51c97e55f13dfcd39f5d58486923273c95dc823bd71a150aa071e6364736f6c634300080d0033",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// PoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolMetaData.Bin instead.
var PoolBin = PoolMetaData.Bin

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolCaller) PT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "PT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolSession) PT() (common.Address, error) {
	return _Pool.Contract.PT(&_Pool.CallOpts)
}

// PT is a free data retrieval call binding the contract method 0xd94073d4.
//
// Solidity: function PT() view returns(address)
func (_Pool *PoolCallerSession) PT() (common.Address, error) {
	return _Pool.Contract.PT(&_Pool.CallOpts)
}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(address PTTo, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCaller) BurnCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	PTTo     common.Address
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "burnCalled", arg0)

	outstruct := new(struct {
		PTTo     common.Address
		MinRatio *big.Int
		MaxRatio *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PTTo = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(address PTTo, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolSession) BurnCalled(arg0 common.Address) (struct {
	PTTo     common.Address
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	return _Pool.Contract.BurnCalled(&_Pool.CallOpts, arg0)
}

// BurnCalled is a free data retrieval call binding the contract method 0xbba0ad39.
//
// Solidity: function burnCalled(address ) view returns(address PTTo, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCallerSession) BurnCalled(arg0 common.Address) (struct {
	PTTo     common.Address
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	return _Pool.Contract.BurnCalled(&_Pool.CallOpts, arg0)
}

// BurnForUnderlyingCalled is a free data retrieval call binding the contract method 0xec92245a.
//
// Solidity: function burnForUnderlyingCalled(address ) view returns(uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCaller) BurnForUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "burnForUnderlyingCalled", arg0)

	outstruct := new(struct {
		MinRatio *big.Int
		MaxRatio *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinRatio = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BurnForUnderlyingCalled is a free data retrieval call binding the contract method 0xec92245a.
//
// Solidity: function burnForUnderlyingCalled(address ) view returns(uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolSession) BurnForUnderlyingCalled(arg0 common.Address) (struct {
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	return _Pool.Contract.BurnForUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BurnForUnderlyingCalled is a free data retrieval call binding the contract method 0xec92245a.
//
// Solidity: function burnForUnderlyingCalled(address ) view returns(uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCallerSession) BurnForUnderlyingCalled(arg0 common.Address) (struct {
	MinRatio *big.Int
	MaxRatio *big.Int
}, error) {
	return _Pool.Contract.BurnForUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolCaller) BuyPrincipalTokenCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPrincipalTokenCalled", arg0)

	outstruct := new(struct {
		PrincipalTokenOut *big.Int
		Min               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrincipalTokenOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Min = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolSession) BuyPrincipalTokenCalled(arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	return _Pool.Contract.BuyPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// BuyPrincipalTokenCalled is a free data retrieval call binding the contract method 0xe525edce.
//
// Solidity: function buyPrincipalTokenCalled(address ) view returns(uint128 PrincipalTokenOut, uint128 min)
func (_Pool *PoolCallerSession) BuyPrincipalTokenCalled(arg0 common.Address) (struct {
	PrincipalTokenOut *big.Int
	Min               *big.Int
}, error) {
	return _Pool.Contract.BuyPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) BuyPrincipalTokenPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyPrincipalTokenPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) BuyPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// BuyPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x435bf331.
//
// Solidity: function buyPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) BuyPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolCaller) BuyUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyUnderlyingCalled", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Min    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Min = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolSession) BuyUnderlyingCalled(arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	return _Pool.Contract.BuyUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BuyUnderlyingCalled is a free data retrieval call binding the contract method 0xd9b03f88.
//
// Solidity: function buyUnderlyingCalled(address ) view returns(uint128 amount, uint128 min)
func (_Pool *PoolCallerSession) BuyUnderlyingCalled(arg0 common.Address) (struct {
	Amount *big.Int
	Min    *big.Int
}, error) {
	return _Pool.Contract.BuyUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) BuyUnderlyingPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "buyUnderlyingPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) BuyUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// BuyUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0x6fccc087.
//
// Solidity: function buyUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) BuyUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.BuyUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(address remainder, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCaller) MintCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Remainder common.Address
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "mintCalled", arg0)

	outstruct := new(struct {
		Remainder common.Address
		MinRatio  *big.Int
		MaxRatio  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remainder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(address remainder, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolSession) MintCalled(arg0 common.Address) (struct {
	Remainder common.Address
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	return _Pool.Contract.MintCalled(&_Pool.CallOpts, arg0)
}

// MintCalled is a free data retrieval call binding the contract method 0xee4db570.
//
// Solidity: function mintCalled(address ) view returns(address remainder, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCallerSession) MintCalled(arg0 common.Address) (struct {
	Remainder common.Address
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	return _Pool.Contract.MintCalled(&_Pool.CallOpts, arg0)
}

// MintWithUnderlyingCalled is a free data retrieval call binding the contract method 0xc100a7fd.
//
// Solidity: function mintWithUnderlyingCalled(address ) view returns(address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCaller) MintWithUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	Remainder common.Address
	PTToBuy   *big.Int
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "mintWithUnderlyingCalled", arg0)

	outstruct := new(struct {
		Remainder common.Address
		PTToBuy   *big.Int
		MinRatio  *big.Int
		MaxRatio  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remainder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PTToBuy = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinRatio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxRatio = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintWithUnderlyingCalled is a free data retrieval call binding the contract method 0xc100a7fd.
//
// Solidity: function mintWithUnderlyingCalled(address ) view returns(address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolSession) MintWithUnderlyingCalled(arg0 common.Address) (struct {
	Remainder common.Address
	PTToBuy   *big.Int
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	return _Pool.Contract.MintWithUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// MintWithUnderlyingCalled is a free data retrieval call binding the contract method 0xc100a7fd.
//
// Solidity: function mintWithUnderlyingCalled(address ) view returns(address remainder, uint256 PTToBuy, uint256 minRatio, uint256 maxRatio)
func (_Pool *PoolCallerSession) MintWithUnderlyingCalled(arg0 common.Address) (struct {
	Remainder common.Address
	PTToBuy   *big.Int
	MinRatio  *big.Int
	MaxRatio  *big.Int
}, error) {
	return _Pool.Contract.MintWithUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolCaller) PrincipalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "principalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolSession) PrincipalToken() (common.Address, error) {
	return _Pool.Contract.PrincipalToken(&_Pool.CallOpts)
}

// PrincipalToken is a free data retrieval call binding the contract method 0xdbc162de.
//
// Solidity: function principalToken() view returns(address)
func (_Pool *PoolCallerSession) PrincipalToken() (common.Address, error) {
	return _Pool.Contract.PrincipalToken(&_Pool.CallOpts)
}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolCaller) SellPrincipalTokenCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPrincipalTokenCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// SellPrincipalTokenCalled is a free data retrieval call binding the contract method 0x9d92bc2f.
//
// Solidity: function sellPrincipalTokenCalled(address ) view returns(uint128)
func (_Pool *PoolCallerSession) SellPrincipalTokenCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenCalled(&_Pool.CallOpts, arg0)
}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) SellPrincipalTokenPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellPrincipalTokenPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// SellPrincipalTokenPreviewCalled is a free data retrieval call binding the contract method 0x15c49330.
//
// Solidity: function sellPrincipalTokenPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) SellPrincipalTokenPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewCalled(&_Pool.CallOpts)
}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolCaller) SellUnderlyingCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellUnderlyingCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolSession) SellUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// SellUnderlyingCalled is a free data retrieval call binding the contract method 0x463a861e.
//
// Solidity: function sellUnderlyingCalled(address ) view returns(uint128)
func (_Pool *PoolCallerSession) SellUnderlyingCalled(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingCalled(&_Pool.CallOpts, arg0)
}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCaller) SellUnderlyingPreviewCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sellUnderlyingPreviewCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolSession) SellUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// SellUnderlyingPreviewCalled is a free data retrieval call binding the contract method 0xe42b16cb.
//
// Solidity: function sellUnderlyingPreviewCalled() view returns(uint128)
func (_Pool *PoolCallerSession) SellUnderlyingPreviewCalled() (*big.Int, error) {
	return _Pool.Contract.SellUnderlyingPreviewCalled(&_Pool.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolSession) Underlying() (common.Address, error) {
	return _Pool.Contract.Underlying(&_Pool.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Pool *PoolCallerSession) Underlying() (common.Address, error) {
	return _Pool.Contract.Underlying(&_Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactor) Burn(opts *bind.TransactOpts, t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burn", t, remained, minRatio, maxRatio)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolSession) Burn(t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, t, remained, minRatio, maxRatio)
}

// Burn is a paid mutator transaction binding the contract method 0xd7020d0a.
//
// Solidity: function burn(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactorSession) Burn(t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, t, remained, minRatio, maxRatio)
}

// BurnForUnderlying is a paid mutator transaction binding the contract method 0x584c0192.
//
// Solidity: function burnForUnderlying(address t, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256)
func (_Pool *PoolTransactor) BurnForUnderlying(opts *bind.TransactOpts, t common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burnForUnderlying", t, minRatio, maxRatio)
}

// BurnForUnderlying is a paid mutator transaction binding the contract method 0x584c0192.
//
// Solidity: function burnForUnderlying(address t, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256)
func (_Pool *PoolSession) BurnForUnderlying(t common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BurnForUnderlying(&_Pool.TransactOpts, t, minRatio, maxRatio)
}

// BurnForUnderlying is a paid mutator transaction binding the contract method 0x584c0192.
//
// Solidity: function burnForUnderlying(address t, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256)
func (_Pool *PoolTransactorSession) BurnForUnderlying(t common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BurnForUnderlying(&_Pool.TransactOpts, t, minRatio, maxRatio)
}

// BurnForUnderlyingReturns is a paid mutator transaction binding the contract method 0xdb02cd1b.
//
// Solidity: function burnForUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactor) BurnForUnderlyingReturns(opts *bind.TransactOpts, o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burnForUnderlyingReturns", o)
}

// BurnForUnderlyingReturns is a paid mutator transaction binding the contract method 0xdb02cd1b.
//
// Solidity: function burnForUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolSession) BurnForUnderlyingReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.BurnForUnderlyingReturns(&_Pool.TransactOpts, o)
}

// BurnForUnderlyingReturns is a paid mutator transaction binding the contract method 0xdb02cd1b.
//
// Solidity: function burnForUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactorSession) BurnForUnderlyingReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.BurnForUnderlyingReturns(&_Pool.TransactOpts, o)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xa57c3782.
//
// Solidity: function burnReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactor) BurnReturns(opts *bind.TransactOpts, o PoolBurnOutputs) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burnReturns", o)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xa57c3782.
//
// Solidity: function burnReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolSession) BurnReturns(o PoolBurnOutputs) (*types.Transaction, error) {
	return _Pool.Contract.BurnReturns(&_Pool.TransactOpts, o)
}

// BurnReturns is a paid mutator transaction binding the contract method 0xa57c3782.
//
// Solidity: function burnReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactorSession) BurnReturns(o PoolBurnOutputs) (*types.Transaction, error) {
	return _Pool.Contract.BurnReturns(&_Pool.TransactOpts, o)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) BuyPrincipalToken(opts *bind.TransactOpts, t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalToken", t, p, m)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolSession) BuyPrincipalToken(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalToken(&_Pool.TransactOpts, t, p, m)
}

// BuyPrincipalToken is a paid mutator transaction binding the contract method 0xb03ceb15.
//
// Solidity: function buyPrincipalToken(address t, uint128 p, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPrincipalToken(t common.Address, p *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalToken(&_Pool.TransactOpts, t, p, m)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactor) BuyPrincipalTokenPreview(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenPreview", o)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolSession) BuyPrincipalTokenPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreview(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreview is a paid mutator transaction binding the contract method 0xc9c917d7.
//
// Solidity: function buyPrincipalTokenPreview(uint128 o) returns(uint128)
func (_Pool *PoolTransactorSession) BuyPrincipalTokenPreview(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreview(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolTransactor) BuyPrincipalTokenPreviewReturns(opts *bind.TransactOpts, o *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenPreviewReturns", o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolSession) BuyPrincipalTokenPreviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewReturns(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x556dd440.
//
// Solidity: function buyPrincipalTokenPreviewReturns(uint128 o) returns()
func (_Pool *PoolTransactorSession) BuyPrincipalTokenPreviewReturns(o *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenPreviewReturns(&_Pool.TransactOpts, o)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolTransactor) BuyPrincipalTokenReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyPrincipalTokenReturns", b)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolSession) BuyPrincipalTokenReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenReturns(&_Pool.TransactOpts, b)
}

// BuyPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xb45d8801.
//
// Solidity: function buyPrincipalTokenReturns(uint128 b) returns()
func (_Pool *PoolTransactorSession) BuyPrincipalTokenReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyPrincipalTokenReturns(&_Pool.TransactOpts, b)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) BuyUnderlying(opts *bind.TransactOpts, t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlying", t, u, m)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolSession) BuyUnderlying(t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlying(&_Pool.TransactOpts, t, u, m)
}

// BuyUnderlying is a paid mutator transaction binding the contract method 0xe78a9b0c.
//
// Solidity: function buyUnderlying(address t, uint128 u, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) BuyUnderlying(t common.Address, u *big.Int, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlying(&_Pool.TransactOpts, t, u, m)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactor) BuyUnderlyingPreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingPreview", u)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolSession) BuyUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreview(&_Pool.TransactOpts, u)
}

// BuyUnderlyingPreview is a paid mutator transaction binding the contract method 0x76a91253.
//
// Solidity: function buyUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactorSession) BuyUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreview(&_Pool.TransactOpts, u)
}

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolTransactor) BuyUnderlyingPreviewReturns(opts *bind.TransactOpts, p *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingPreviewReturns", p)
}

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolSession) BuyUnderlyingPreviewReturns(p *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreviewReturns(&_Pool.TransactOpts, p)
}

// BuyUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0xf038a50f.
//
// Solidity: function buyUnderlyingPreviewReturns(uint128 p) returns()
func (_Pool *PoolTransactorSession) BuyUnderlyingPreviewReturns(p *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingPreviewReturns(&_Pool.TransactOpts, p)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolTransactor) BuyUnderlyingReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "buyUnderlyingReturns", b)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolSession) BuyUnderlyingReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingReturns(&_Pool.TransactOpts, b)
}

// BuyUnderlyingReturns is a paid mutator transaction binding the contract method 0x08e1cff5.
//
// Solidity: function buyUnderlyingReturns(uint128 b) returns()
func (_Pool *PoolTransactorSession) BuyUnderlyingReturns(b *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.BuyUnderlyingReturns(&_Pool.TransactOpts, b)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", t, remained, minRatio, maxRatio)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolSession) Mint(t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, t, remained, minRatio, maxRatio)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address t, address remained, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactorSession) Mint(t common.Address, remained common.Address, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, t, remained, minRatio, maxRatio)
}

// MintReturns is a paid mutator transaction binding the contract method 0xbfdf149f.
//
// Solidity: function mintReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactor) MintReturns(opts *bind.TransactOpts, o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mintReturns", o)
}

// MintReturns is a paid mutator transaction binding the contract method 0xbfdf149f.
//
// Solidity: function mintReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolSession) MintReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.MintReturns(&_Pool.TransactOpts, o)
}

// MintReturns is a paid mutator transaction binding the contract method 0xbfdf149f.
//
// Solidity: function mintReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactorSession) MintReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.MintReturns(&_Pool.TransactOpts, o)
}

// MintWithUnderlying is a paid mutator transaction binding the contract method 0x50ed7074.
//
// Solidity: function mintWithUnderlying(address t, address remained, uint256 ptToBuy, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactor) MintWithUnderlying(opts *bind.TransactOpts, t common.Address, remained common.Address, ptToBuy *big.Int, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mintWithUnderlying", t, remained, ptToBuy, minRatio, maxRatio)
}

// MintWithUnderlying is a paid mutator transaction binding the contract method 0x50ed7074.
//
// Solidity: function mintWithUnderlying(address t, address remained, uint256 ptToBuy, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolSession) MintWithUnderlying(t common.Address, remained common.Address, ptToBuy *big.Int, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.MintWithUnderlying(&_Pool.TransactOpts, t, remained, ptToBuy, minRatio, maxRatio)
}

// MintWithUnderlying is a paid mutator transaction binding the contract method 0x50ed7074.
//
// Solidity: function mintWithUnderlying(address t, address remained, uint256 ptToBuy, uint256 minRatio, uint256 maxRatio) returns(uint256, uint256, uint256)
func (_Pool *PoolTransactorSession) MintWithUnderlying(t common.Address, remained common.Address, ptToBuy *big.Int, minRatio *big.Int, maxRatio *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.MintWithUnderlying(&_Pool.TransactOpts, t, remained, ptToBuy, minRatio, maxRatio)
}

// MintWithUnderlyingReturns is a paid mutator transaction binding the contract method 0x814c83d7.
//
// Solidity: function mintWithUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactor) MintWithUnderlyingReturns(opts *bind.TransactOpts, o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mintWithUnderlyingReturns", o)
}

// MintWithUnderlyingReturns is a paid mutator transaction binding the contract method 0x814c83d7.
//
// Solidity: function mintWithUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolSession) MintWithUnderlyingReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.MintWithUnderlyingReturns(&_Pool.TransactOpts, o)
}

// MintWithUnderlyingReturns is a paid mutator transaction binding the contract method 0x814c83d7.
//
// Solidity: function mintWithUnderlyingReturns((uint256,uint256,uint256) o) returns()
func (_Pool *PoolTransactorSession) MintWithUnderlyingReturns(o PoolMintOutputs) (*types.Transaction, error) {
	return _Pool.Contract.MintWithUnderlyingReturns(&_Pool.TransactOpts, o)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolTransactor) PrincipalTokenReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "principalTokenReturns", p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PrincipalTokenReturns(&_Pool.TransactOpts, p)
}

// PrincipalTokenReturns is a paid mutator transaction binding the contract method 0xf3ec0d2c.
//
// Solidity: function principalTokenReturns(address p) returns()
func (_Pool *PoolTransactorSession) PrincipalTokenReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PrincipalTokenReturns(&_Pool.TransactOpts, p)
}

// PtReturns is a paid mutator transaction binding the contract method 0x0326b5ed.
//
// Solidity: function ptReturns(address p) returns()
func (_Pool *PoolTransactor) PtReturns(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "ptReturns", p)
}

// PtReturns is a paid mutator transaction binding the contract method 0x0326b5ed.
//
// Solidity: function ptReturns(address p) returns()
func (_Pool *PoolSession) PtReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PtReturns(&_Pool.TransactOpts, p)
}

// PtReturns is a paid mutator transaction binding the contract method 0x0326b5ed.
//
// Solidity: function ptReturns(address p) returns()
func (_Pool *PoolTransactorSession) PtReturns(p common.Address) (*types.Transaction, error) {
	return _Pool.Contract.PtReturns(&_Pool.TransactOpts, p)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) SellPrincipalToken(opts *bind.TransactOpts, t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalToken", t, m)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolSession) SellPrincipalToken(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalToken(&_Pool.TransactOpts, t, m)
}

// SellPrincipalToken is a paid mutator transaction binding the contract method 0x134df58b.
//
// Solidity: function sellPrincipalToken(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) SellPrincipalToken(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalToken(&_Pool.TransactOpts, t, m)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactor) SellPrincipalTokenPreview(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenPreview", i)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolSession) SellPrincipalTokenPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreview(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreview is a paid mutator transaction binding the contract method 0x0a7e546e.
//
// Solidity: function sellPrincipalTokenPreview(uint128 i) returns(uint128)
func (_Pool *PoolTransactorSession) SellPrincipalTokenPreview(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreview(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolTransactor) SellPrincipalTokenPreviewReturns(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenPreviewReturns", i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolSession) SellPrincipalTokenPreviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewReturns(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenPreviewReturns is a paid mutator transaction binding the contract method 0x9b9c9fde.
//
// Solidity: function sellPrincipalTokenPreviewReturns(uint128 i) returns()
func (_Pool *PoolTransactorSession) SellPrincipalTokenPreviewReturns(i *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenPreviewReturns(&_Pool.TransactOpts, i)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellPrincipalTokenReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellPrincipalTokenReturns", s)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolSession) SellPrincipalTokenReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenReturns(&_Pool.TransactOpts, s)
}

// SellPrincipalTokenReturns is a paid mutator transaction binding the contract method 0xc59964b3.
//
// Solidity: function sellPrincipalTokenReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellPrincipalTokenReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellPrincipalTokenReturns(&_Pool.TransactOpts, s)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactor) SellUnderlying(opts *bind.TransactOpts, t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlying", t, m)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolSession) SellUnderlying(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlying(&_Pool.TransactOpts, t, m)
}

// SellUnderlying is a paid mutator transaction binding the contract method 0x3e232091.
//
// Solidity: function sellUnderlying(address t, uint128 m) returns(uint128)
func (_Pool *PoolTransactorSession) SellUnderlying(t common.Address, m *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlying(&_Pool.TransactOpts, t, m)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactor) SellUnderlyingPreview(opts *bind.TransactOpts, u *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingPreview", u)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolSession) SellUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreview(&_Pool.TransactOpts, u)
}

// SellUnderlyingPreview is a paid mutator transaction binding the contract method 0xf332ea5e.
//
// Solidity: function sellUnderlyingPreview(uint128 u) returns(uint128)
func (_Pool *PoolTransactorSession) SellUnderlyingPreview(u *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreview(&_Pool.TransactOpts, u)
}

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellUnderlyingPreviewReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingPreviewReturns", s)
}

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolSession) SellUnderlyingPreviewReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreviewReturns(&_Pool.TransactOpts, s)
}

// SellUnderlyingPreviewReturns is a paid mutator transaction binding the contract method 0x4b266e28.
//
// Solidity: function sellUnderlyingPreviewReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellUnderlyingPreviewReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingPreviewReturns(&_Pool.TransactOpts, s)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolTransactor) SellUnderlyingReturns(opts *bind.TransactOpts, s *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sellUnderlyingReturns", s)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolSession) SellUnderlyingReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingReturns(&_Pool.TransactOpts, s)
}

// SellUnderlyingReturns is a paid mutator transaction binding the contract method 0x757b2b76.
//
// Solidity: function sellUnderlyingReturns(uint128 s) returns()
func (_Pool *PoolTransactorSession) SellUnderlyingReturns(s *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SellUnderlyingReturns(&_Pool.TransactOpts, s)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolTransactor) UnderlyingReturns(opts *bind.TransactOpts, u common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "underlyingReturns", u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UnderlyingReturns(&_Pool.TransactOpts, u)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address u) returns()
func (_Pool *PoolTransactorSession) UnderlyingReturns(u common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UnderlyingReturns(&_Pool.TransactOpts, u)
}
