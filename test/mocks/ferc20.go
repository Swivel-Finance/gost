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

// FErc20MetaData contains all meta data concerning the FErc20 contract.
var FErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"allocateTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allocateToCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"allocateToReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowanceCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"allowanceReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approveCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"approveReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"balanceOfReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"}],\"name\":\"decimalsReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"exchangeRateCurrentReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"mintReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"nameReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemUnderlyingCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"supplyRatePerBlockReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"symbolReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferCalled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferFromCalled\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferFromReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"transferReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"underlyingReturns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061146e806100206000396000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80639dd0ff3711610125578063bd6d894d116100ad578063d6bcd7aa1161007c578063d6bcd7aa14610648578063dd62ed3e14610666578063e541efa214610696578063e7a7b9ce146106c7578063e7ba6774146106e357610211565b8063bd6d894d146105c0578063c1d2e9a1146105de578063c6bf15521461060e578063d4e7fdd41461062a57610211565b8063a5b836bf116100f4578063a5b836bf1461050a578063a9059cbb1461053a578063ad64dcc51461056a578063ae9d70b014610586578063af599a29146105a457610211565b80639dd0ff37146104865780639ff9f1d4146104a2578063a0712d68146104be578063a18729a3146104ee57610211565b806339100838116101a85780636581d543116101775780636581d543146103ba5780636f307dc3146103ea57806370a0823114610408578063852a12e31461043857806395d89b411461046857610211565b8063391008381461033657806342b6cdbc14610352578063579296501461036e5780636521b96a1461039e57610211565b806323b872dd116101e457806323b872dd146102b057806325541305146102e057806329d9ce3e146102fc578063313ce5671461031857610211565b806306fdde031461021657806308bca56614610234578063095ea7b3146102645780630ab2519e14610294575b600080fd5b61021e6106ff565b60405161022b91906111a9565b60405180910390f35b61024e60048036038101906102499190610fdd565b610791565b60405161025b919061118e565b60405180910390f35b61027e60048036038101906102799190610fdd565b6107ef565b60405161028b919061118e565b60405180910390f35b6102ae60048036038101906102a99190611042565b61084c565b005b6102ca60048036038101906102c59190610f8e565b610866565b6040516102d7919061118e565b60405180910390f35b6102fa60048036038101906102f59190611019565b610960565b005b61031660048036038101906103119190611083565b61097d565b005b610320610987565b60405161032d91906111e6565b60405180910390f35b610350600480360381019061034b9190611083565b61099e565b005b61036c60048036038101906103679190611019565b6109a8565b005b61038860048036038101906103839190610f29565b6109c5565b60405161039591906111cb565b60405180910390f35b6103b860048036038101906103b39190611019565b6109dd565b005b6103d460048036038101906103cf9190610f29565b6109fa565b6040516103e191906111cb565b60405180910390f35b6103f2610a12565b6040516103ff919061114a565b60405180910390f35b610422600480360381019061041d9190610f29565b610a3c565b60405161042f91906111cb565b60405180910390f35b610452600480360381019061044d9190611083565b610a89565b60405161045f91906111cb565b60405180910390f35b610470610a9c565b60405161047d91906111a9565b60405180910390f35b6104a0600480360381019061049b9190611019565b610b2e565b005b6104bc60048036038101906104b79190611083565b610b4b565b005b6104d860048036038101906104d39190611083565b610b55565b6040516104e591906111cb565b60405180910390f35b610508600480360381019061050391906110ac565b610b68565b005b610524600480360381019061051f9190610f29565b610b86565b604051610531919061114a565b60405180910390f35b610554600480360381019061054f9190610fdd565b610bb9565b604051610561919061118e565b60405180910390f35b610584600480360381019061057f9190611083565b610c17565b005b61058e610c21565b60405161059b91906111cb565b60405180910390f35b6105be60048036038101906105b99190611042565b610c2b565b005b6105c8610c45565b6040516105d591906111cb565b60405180910390f35b6105f860048036038101906105f39190610f29565b610c4f565b60405161060591906111cb565b60405180910390f35b61062860048036038101906106239190611083565b610c67565b005b610632610c71565b60405161063f91906111cb565b60405180910390f35b610650610c77565b60405161065d91906111cb565b60405180910390f35b610680600480360381019061067b9190610f52565b610c7d565b60405161068d91906111cb565b60405180910390f35b6106b060048036038101906106ab9190610f29565b610d08565b6040516106be929190611165565b60405180910390f35b6106e160048036038101906106dc9190611083565b610d4c565b005b6106fd60048036038101906106f89190610f29565b610d56565b005b60606005805461070e9061130a565b80601f016020809104026020016040519081016040528092919081815260200182805461073a9061130a565b80156107875780601f1061075c57610100808354040283529160200191610787565b820191906000526020600020905b81548152906001019060200180831161076a57829003601f168201915b5050505050905090565b600081601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550601360009054906101000a900460ff16905092915050565b6000816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600760019054906101000a900460ff16905092915050565b8060059080519060200190610862929190610d9a565b5050565b6000610870610e20565b83816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508281602001818152505080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050600a60019054906101000a900460ff169150509392505050565b80601360006101000a81548160ff02191690831515021790555050565b80600d8190555050565b6000600760009054906101000a900460ff16905090565b8060088190555050565b80600a60006101000a81548160ff02191690831515021790555050565b60126020528060005260406000206000915090505481565b80600a60016101000a81548160ff02191690831515021790555050565b60006020528060005260406000206000915090505481565b6000600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506008549050919050565b600081600e81905550600d549050919050565b606060068054610aab9061130a565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad79061130a565b8015610b245780601f10610af957610100808354040283529160200191610b24565b820191906000526020600020905b815481529060010190602001808311610b0757829003601f168201915b5050505050905090565b80600760016101000a81548160ff02191690831515021790555050565b8060108190555050565b600081600c81905550600b549050919050565b80600760006101000a81548160ff021916908360ff16021790555050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600a60009054906101000a900460ff16905092915050565b8060098190555050565b6000601154905090565b8060069080519060200190610c41929190610d9a565b5050565b6000601054905090565b60016020528060005260406000206000915090505481565b8060118190555050565b600c5481565b600e5481565b600081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600954905092915050565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b80600b8190555050565b80600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054610da69061130a565b90600052602060002090601f016020900481019282610dc85760008555610e0f565b82601f10610de157805160ff1916838001178555610e0f565b82800160010185558215610e0f579182015b82811115610e0e578251825591602001919060010190610df3565b5b509050610e1c9190610e50565b5090565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b5b80821115610e69576000816000905550600101610e51565b5090565b6000610e80610e7b84611226565b611201565b905082815260208101848484011115610e9857600080fd5b610ea38482856112c8565b509392505050565b600081359050610eba816113dc565b92915050565b600081359050610ecf816113f3565b92915050565b600082601f830112610ee657600080fd5b8135610ef6848260208601610e6d565b91505092915050565b600081359050610f0e8161140a565b92915050565b600081359050610f2381611421565b92915050565b600060208284031215610f3b57600080fd5b6000610f4984828501610eab565b91505092915050565b60008060408385031215610f6557600080fd5b6000610f7385828601610eab565b9250506020610f8485828601610eab565b9150509250929050565b600080600060608486031215610fa357600080fd5b6000610fb186828701610eab565b9350506020610fc286828701610eab565b9250506040610fd386828701610eff565b9150509250925092565b60008060408385031215610ff057600080fd5b6000610ffe85828601610eab565b925050602061100f85828601610eff565b9150509250929050565b60006020828403121561102b57600080fd5b600061103984828501610ec0565b91505092915050565b60006020828403121561105457600080fd5b600082013567ffffffffffffffff81111561106e57600080fd5b61107a84828501610ed5565b91505092915050565b60006020828403121561109557600080fd5b60006110a384828501610eff565b91505092915050565b6000602082840312156110be57600080fd5b60006110cc84828501610f14565b91505092915050565b6110de81611273565b82525050565b6110ed81611285565b82525050565b60006110fe82611257565b6111088185611262565b93506111188185602086016112d7565b611121816113cb565b840191505092915050565b611135816112b1565b82525050565b611144816112bb565b82525050565b600060208201905061115f60008301846110d5565b92915050565b600060408201905061117a60008301856110d5565b611187602083018461112c565b9392505050565b60006020820190506111a360008301846110e4565b92915050565b600060208201905081810360008301526111c381846110f3565b905092915050565b60006020820190506111e0600083018461112c565b92915050565b60006020820190506111fb600083018461113b565b92915050565b600061120b61121c565b9050611217828261133c565b919050565b6000604051905090565b600067ffffffffffffffff8211156112415761124061139c565b5b61124a826113cb565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061127e82611291565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156112f55780820151818401526020810190506112da565b83811115611304576000848401525b50505050565b6000600282049050600182168061132257607f821691505b602082108114156113365761133561136d565b5b50919050565b611345826113cb565b810181811067ffffffffffffffff821117156113645761136361139c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6113e581611273565b81146113f057600080fd5b50565b6113fc81611285565b811461140757600080fd5b50565b611413816112b1565b811461141e57600080fd5b50565b61142a816112bb565b811461143557600080fd5b5056fea2646970667358221220f5338f8588c7944f8a8522d121602d17af383f9d305007d7f60dffe9d675f4c164736f6c63430008040033",
}

// FErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use FErc20MetaData.ABI instead.
var FErc20ABI = FErc20MetaData.ABI

// FErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FErc20MetaData.Bin instead.
var FErc20Bin = FErc20MetaData.Bin

// DeployFErc20 deploys a new Ethereum contract, binding an instance of FErc20 to it.
func DeployFErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FErc20, error) {
	parsed, err := FErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FErc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FErc20{FErc20Caller: FErc20Caller{contract: contract}, FErc20Transactor: FErc20Transactor{contract: contract}, FErc20Filterer: FErc20Filterer{contract: contract}}, nil
}

// FErc20 is an auto generated Go binding around an Ethereum contract.
type FErc20 struct {
	FErc20Caller     // Read-only binding to the contract
	FErc20Transactor // Write-only binding to the contract
	FErc20Filterer   // Log filterer for contract events
}

// FErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type FErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FErc20Session struct {
	Contract     *FErc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FErc20CallerSession struct {
	Contract *FErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FErc20TransactorSession struct {
	Contract     *FErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type FErc20Raw struct {
	Contract *FErc20 // Generic contract binding to access the raw methods on
}

// FErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FErc20CallerRaw struct {
	Contract *FErc20Caller // Generic read-only contract binding to access the raw methods on
}

// FErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FErc20TransactorRaw struct {
	Contract *FErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFErc20 creates a new instance of FErc20, bound to a specific deployed contract.
func NewFErc20(address common.Address, backend bind.ContractBackend) (*FErc20, error) {
	contract, err := bindFErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FErc20{FErc20Caller: FErc20Caller{contract: contract}, FErc20Transactor: FErc20Transactor{contract: contract}, FErc20Filterer: FErc20Filterer{contract: contract}}, nil
}

// NewFErc20Caller creates a new read-only instance of FErc20, bound to a specific deployed contract.
func NewFErc20Caller(address common.Address, caller bind.ContractCaller) (*FErc20Caller, error) {
	contract, err := bindFErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FErc20Caller{contract: contract}, nil
}

// NewFErc20Transactor creates a new write-only instance of FErc20, bound to a specific deployed contract.
func NewFErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*FErc20Transactor, error) {
	contract, err := bindFErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FErc20Transactor{contract: contract}, nil
}

// NewFErc20Filterer creates a new log filterer instance of FErc20, bound to a specific deployed contract.
func NewFErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*FErc20Filterer, error) {
	contract, err := bindFErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FErc20Filterer{contract: contract}, nil
}

// bindFErc20 binds a generic wrapper to an already deployed contract.
func bindFErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FErc20 *FErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FErc20.Contract.FErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FErc20 *FErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FErc20.Contract.FErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FErc20 *FErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FErc20.Contract.FErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FErc20 *FErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FErc20 *FErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FErc20 *FErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FErc20.Contract.contract.Transact(opts, method, params...)
}

// AllocateToCalled is a free data retrieval call binding the contract method 0x57929650.
//
// Solidity: function allocateToCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Caller) AllocateToCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "allocateToCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocateToCalled is a free data retrieval call binding the contract method 0x57929650.
//
// Solidity: function allocateToCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Session) AllocateToCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.AllocateToCalled(&_FErc20.CallOpts, arg0)
}

// AllocateToCalled is a free data retrieval call binding the contract method 0x57929650.
//
// Solidity: function allocateToCalled(address ) view returns(uint256)
func (_FErc20 *FErc20CallerSession) AllocateToCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.AllocateToCalled(&_FErc20.CallOpts, arg0)
}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_FErc20 *FErc20Caller) AllowanceCalled(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "allowanceCalled", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_FErc20 *FErc20Session) AllowanceCalled(arg0 common.Address) (common.Address, error) {
	return _FErc20.Contract.AllowanceCalled(&_FErc20.CallOpts, arg0)
}

// AllowanceCalled is a free data retrieval call binding the contract method 0xa5b836bf.
//
// Solidity: function allowanceCalled(address ) view returns(address)
func (_FErc20 *FErc20CallerSession) AllowanceCalled(arg0 common.Address) (common.Address, error) {
	return _FErc20.Contract.AllowanceCalled(&_FErc20.CallOpts, arg0)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Caller) ApproveCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "approveCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Session) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.ApproveCalled(&_FErc20.CallOpts, arg0)
}

// ApproveCalled is a free data retrieval call binding the contract method 0x6581d543.
//
// Solidity: function approveCalled(address ) view returns(uint256)
func (_FErc20 *FErc20CallerSession) ApproveCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.ApproveCalled(&_FErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FErc20 *FErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FErc20 *FErc20Session) Decimals() (uint8, error) {
	return _FErc20.Contract.Decimals(&_FErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FErc20 *FErc20CallerSession) Decimals() (uint8, error) {
	return _FErc20.Contract.Decimals(&_FErc20.CallOpts)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_FErc20 *FErc20Caller) ExchangeRateCurrent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "exchangeRateCurrent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_FErc20 *FErc20Session) ExchangeRateCurrent() (*big.Int, error) {
	return _FErc20.Contract.ExchangeRateCurrent(&_FErc20.CallOpts)
}

// ExchangeRateCurrent is a free data retrieval call binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() view returns(uint256)
func (_FErc20 *FErc20CallerSession) ExchangeRateCurrent() (*big.Int, error) {
	return _FErc20.Contract.ExchangeRateCurrent(&_FErc20.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_FErc20 *FErc20Caller) MintCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "mintCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_FErc20 *FErc20Session) MintCalled() (*big.Int, error) {
	return _FErc20.Contract.MintCalled(&_FErc20.CallOpts)
}

// MintCalled is a free data retrieval call binding the contract method 0xd4e7fdd4.
//
// Solidity: function mintCalled() view returns(uint256)
func (_FErc20 *FErc20CallerSession) MintCalled() (*big.Int, error) {
	return _FErc20.Contract.MintCalled(&_FErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FErc20 *FErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FErc20 *FErc20Session) Name() (string, error) {
	return _FErc20.Contract.Name(&_FErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FErc20 *FErc20CallerSession) Name() (string, error) {
	return _FErc20.Contract.Name(&_FErc20.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_FErc20 *FErc20Caller) RedeemUnderlyingCalled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "redeemUnderlyingCalled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_FErc20 *FErc20Session) RedeemUnderlyingCalled() (*big.Int, error) {
	return _FErc20.Contract.RedeemUnderlyingCalled(&_FErc20.CallOpts)
}

// RedeemUnderlyingCalled is a free data retrieval call binding the contract method 0xd6bcd7aa.
//
// Solidity: function redeemUnderlyingCalled() view returns(uint256)
func (_FErc20 *FErc20CallerSession) RedeemUnderlyingCalled() (*big.Int, error) {
	return _FErc20.Contract.RedeemUnderlyingCalled(&_FErc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_FErc20 *FErc20Caller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_FErc20 *FErc20Session) SupplyRatePerBlock() (*big.Int, error) {
	return _FErc20.Contract.SupplyRatePerBlock(&_FErc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_FErc20 *FErc20CallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _FErc20.Contract.SupplyRatePerBlock(&_FErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FErc20 *FErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FErc20 *FErc20Session) Symbol() (string, error) {
	return _FErc20.Contract.Symbol(&_FErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FErc20 *FErc20CallerSession) Symbol() (string, error) {
	return _FErc20.Contract.Symbol(&_FErc20.CallOpts)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Caller) TransferCalled(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "transferCalled", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_FErc20 *FErc20Session) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.TransferCalled(&_FErc20.CallOpts, arg0)
}

// TransferCalled is a free data retrieval call binding the contract method 0xc1d2e9a1.
//
// Solidity: function transferCalled(address ) view returns(uint256)
func (_FErc20 *FErc20CallerSession) TransferCalled(arg0 common.Address) (*big.Int, error) {
	return _FErc20.Contract.TransferCalled(&_FErc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_FErc20 *FErc20Caller) TransferFromCalled(opts *bind.CallOpts, arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "transferFromCalled", arg0)

	outstruct := new(struct {
		To     common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_FErc20 *FErc20Session) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _FErc20.Contract.TransferFromCalled(&_FErc20.CallOpts, arg0)
}

// TransferFromCalled is a free data retrieval call binding the contract method 0xe541efa2.
//
// Solidity: function transferFromCalled(address ) view returns(address to, uint256 amount)
func (_FErc20 *FErc20CallerSession) TransferFromCalled(arg0 common.Address) (struct {
	To     common.Address
	Amount *big.Int
}, error) {
	return _FErc20.Contract.TransferFromCalled(&_FErc20.CallOpts, arg0)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_FErc20 *FErc20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FErc20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_FErc20 *FErc20Session) Underlying() (common.Address, error) {
	return _FErc20.Contract.Underlying(&_FErc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_FErc20 *FErc20CallerSession) Underlying() (common.Address, error) {
	return _FErc20.Contract.Underlying(&_FErc20.CallOpts)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address o, uint256 a) returns(bool)
func (_FErc20 *FErc20Transactor) AllocateTo(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "allocateTo", o, a)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address o, uint256 a) returns(bool)
func (_FErc20 *FErc20Session) AllocateTo(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.AllocateTo(&_FErc20.TransactOpts, o, a)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address o, uint256 a) returns(bool)
func (_FErc20 *FErc20TransactorSession) AllocateTo(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.AllocateTo(&_FErc20.TransactOpts, o, a)
}

// AllocateToReturns is a paid mutator transaction binding the contract method 0x25541305.
//
// Solidity: function allocateToReturns(bool b) returns()
func (_FErc20 *FErc20Transactor) AllocateToReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "allocateToReturns", b)
}

// AllocateToReturns is a paid mutator transaction binding the contract method 0x25541305.
//
// Solidity: function allocateToReturns(bool b) returns()
func (_FErc20 *FErc20Session) AllocateToReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.AllocateToReturns(&_FErc20.TransactOpts, b)
}

// AllocateToReturns is a paid mutator transaction binding the contract method 0x25541305.
//
// Solidity: function allocateToReturns(bool b) returns()
func (_FErc20 *FErc20TransactorSession) AllocateToReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.AllocateToReturns(&_FErc20.TransactOpts, b)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_FErc20 *FErc20Transactor) Allowance(opts *bind.TransactOpts, o common.Address, s common.Address) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "allowance", o, s)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_FErc20 *FErc20Session) Allowance(o common.Address, s common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.Allowance(&_FErc20.TransactOpts, o, s)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) returns(uint256)
func (_FErc20 *FErc20TransactorSession) Allowance(o common.Address, s common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.Allowance(&_FErc20.TransactOpts, o, s)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_FErc20 *FErc20Transactor) AllowanceReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "allowanceReturns", n)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_FErc20 *FErc20Session) AllowanceReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.AllowanceReturns(&_FErc20.TransactOpts, n)
}

// AllowanceReturns is a paid mutator transaction binding the contract method 0xad64dcc5.
//
// Solidity: function allowanceReturns(uint256 n) returns()
func (_FErc20 *FErc20TransactorSession) AllowanceReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.AllowanceReturns(&_FErc20.TransactOpts, n)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_FErc20 *FErc20Transactor) Approve(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "approve", s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_FErc20 *FErc20Session) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Approve(&_FErc20.TransactOpts, s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_FErc20 *FErc20TransactorSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Approve(&_FErc20.TransactOpts, s, a)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_FErc20 *FErc20Transactor) ApproveReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "approveReturns", b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_FErc20 *FErc20Session) ApproveReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.ApproveReturns(&_FErc20.TransactOpts, b)
}

// ApproveReturns is a paid mutator transaction binding the contract method 0x9dd0ff37.
//
// Solidity: function approveReturns(bool b) returns()
func (_FErc20 *FErc20TransactorSession) ApproveReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.ApproveReturns(&_FErc20.TransactOpts, b)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_FErc20 *FErc20Transactor) BalanceOf(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "balanceOf", t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_FErc20 *FErc20Session) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.BalanceOf(&_FErc20.TransactOpts, t)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address t) returns(uint256)
func (_FErc20 *FErc20TransactorSession) BalanceOf(t common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.BalanceOf(&_FErc20.TransactOpts, t)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_FErc20 *FErc20Transactor) BalanceOfReturns(opts *bind.TransactOpts, b *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "balanceOfReturns", b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_FErc20 *FErc20Session) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.BalanceOfReturns(&_FErc20.TransactOpts, b)
}

// BalanceOfReturns is a paid mutator transaction binding the contract method 0x39100838.
//
// Solidity: function balanceOfReturns(uint256 b) returns()
func (_FErc20 *FErc20TransactorSession) BalanceOfReturns(b *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.BalanceOfReturns(&_FErc20.TransactOpts, b)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_FErc20 *FErc20Transactor) DecimalsReturns(opts *bind.TransactOpts, n uint8) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "decimalsReturns", n)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_FErc20 *FErc20Session) DecimalsReturns(n uint8) (*types.Transaction, error) {
	return _FErc20.Contract.DecimalsReturns(&_FErc20.TransactOpts, n)
}

// DecimalsReturns is a paid mutator transaction binding the contract method 0xa18729a3.
//
// Solidity: function decimalsReturns(uint8 n) returns()
func (_FErc20 *FErc20TransactorSession) DecimalsReturns(n uint8) (*types.Transaction, error) {
	return _FErc20.Contract.DecimalsReturns(&_FErc20.TransactOpts, n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_FErc20 *FErc20Transactor) ExchangeRateCurrentReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "exchangeRateCurrentReturns", n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_FErc20 *FErc20Session) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.ExchangeRateCurrentReturns(&_FErc20.TransactOpts, n)
}

// ExchangeRateCurrentReturns is a paid mutator transaction binding the contract method 0x9ff9f1d4.
//
// Solidity: function exchangeRateCurrentReturns(uint256 n) returns()
func (_FErc20 *FErc20TransactorSession) ExchangeRateCurrentReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.ExchangeRateCurrentReturns(&_FErc20.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_FErc20 *FErc20Transactor) Mint(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "mint", n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_FErc20 *FErc20Session) Mint(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Mint(&_FErc20.TransactOpts, n)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 n) returns(uint256)
func (_FErc20 *FErc20TransactorSession) Mint(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Mint(&_FErc20.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_FErc20 *FErc20Transactor) MintReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "mintReturns", n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_FErc20 *FErc20Session) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.MintReturns(&_FErc20.TransactOpts, n)
}

// MintReturns is a paid mutator transaction binding the contract method 0xe7a7b9ce.
//
// Solidity: function mintReturns(uint256 n) returns()
func (_FErc20 *FErc20TransactorSession) MintReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.MintReturns(&_FErc20.TransactOpts, n)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_FErc20 *FErc20Transactor) NameReturns(opts *bind.TransactOpts, s string) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "nameReturns", s)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_FErc20 *FErc20Session) NameReturns(s string) (*types.Transaction, error) {
	return _FErc20.Contract.NameReturns(&_FErc20.TransactOpts, s)
}

// NameReturns is a paid mutator transaction binding the contract method 0x0ab2519e.
//
// Solidity: function nameReturns(string s) returns()
func (_FErc20 *FErc20TransactorSession) NameReturns(s string) (*types.Transaction, error) {
	return _FErc20.Contract.NameReturns(&_FErc20.TransactOpts, s)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_FErc20 *FErc20Transactor) RedeemUnderlying(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "redeemUnderlying", n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_FErc20 *FErc20Session) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.RedeemUnderlying(&_FErc20.TransactOpts, n)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 n) returns(uint256)
func (_FErc20 *FErc20TransactorSession) RedeemUnderlying(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.RedeemUnderlying(&_FErc20.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_FErc20 *FErc20Transactor) RedeemUnderlyingReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "redeemUnderlyingReturns", n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_FErc20 *FErc20Session) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.RedeemUnderlyingReturns(&_FErc20.TransactOpts, n)
}

// RedeemUnderlyingReturns is a paid mutator transaction binding the contract method 0x29d9ce3e.
//
// Solidity: function redeemUnderlyingReturns(uint256 n) returns()
func (_FErc20 *FErc20TransactorSession) RedeemUnderlyingReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.RedeemUnderlyingReturns(&_FErc20.TransactOpts, n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_FErc20 *FErc20Transactor) SupplyRatePerBlockReturns(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "supplyRatePerBlockReturns", n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_FErc20 *FErc20Session) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.SupplyRatePerBlockReturns(&_FErc20.TransactOpts, n)
}

// SupplyRatePerBlockReturns is a paid mutator transaction binding the contract method 0xc6bf1552.
//
// Solidity: function supplyRatePerBlockReturns(uint256 n) returns()
func (_FErc20 *FErc20TransactorSession) SupplyRatePerBlockReturns(n *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.SupplyRatePerBlockReturns(&_FErc20.TransactOpts, n)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_FErc20 *FErc20Transactor) SymbolReturns(opts *bind.TransactOpts, s string) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "symbolReturns", s)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_FErc20 *FErc20Session) SymbolReturns(s string) (*types.Transaction, error) {
	return _FErc20.Contract.SymbolReturns(&_FErc20.TransactOpts, s)
}

// SymbolReturns is a paid mutator transaction binding the contract method 0xaf599a29.
//
// Solidity: function symbolReturns(string s) returns()
func (_FErc20 *FErc20TransactorSession) SymbolReturns(s string) (*types.Transaction, error) {
	return _FErc20.Contract.SymbolReturns(&_FErc20.TransactOpts, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_FErc20 *FErc20Transactor) Transfer(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "transfer", t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_FErc20 *FErc20Session) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Transfer(&_FErc20.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_FErc20 *FErc20TransactorSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.Transfer(&_FErc20.TransactOpts, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_FErc20 *FErc20Transactor) TransferFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "transferFrom", f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_FErc20 *FErc20Session) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.TransferFrom(&_FErc20.TransactOpts, f, t, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address f, address t, uint256 a) returns(bool)
func (_FErc20 *FErc20TransactorSession) TransferFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _FErc20.Contract.TransferFrom(&_FErc20.TransactOpts, f, t, a)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_FErc20 *FErc20Transactor) TransferFromReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "transferFromReturns", b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_FErc20 *FErc20Session) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.TransferFromReturns(&_FErc20.TransactOpts, b)
}

// TransferFromReturns is a paid mutator transaction binding the contract method 0x6521b96a.
//
// Solidity: function transferFromReturns(bool b) returns()
func (_FErc20 *FErc20TransactorSession) TransferFromReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.TransferFromReturns(&_FErc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_FErc20 *FErc20Transactor) TransferReturns(opts *bind.TransactOpts, b bool) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "transferReturns", b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_FErc20 *FErc20Session) TransferReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.TransferReturns(&_FErc20.TransactOpts, b)
}

// TransferReturns is a paid mutator transaction binding the contract method 0x42b6cdbc.
//
// Solidity: function transferReturns(bool b) returns()
func (_FErc20 *FErc20TransactorSession) TransferReturns(b bool) (*types.Transaction, error) {
	return _FErc20.Contract.TransferReturns(&_FErc20.TransactOpts, b)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_FErc20 *FErc20Transactor) UnderlyingReturns(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _FErc20.contract.Transact(opts, "underlyingReturns", a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_FErc20 *FErc20Session) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.UnderlyingReturns(&_FErc20.TransactOpts, a)
}

// UnderlyingReturns is a paid mutator transaction binding the contract method 0xe7ba6774.
//
// Solidity: function underlyingReturns(address a) returns()
func (_FErc20 *FErc20TransactorSession) UnderlyingReturns(a common.Address) (*types.Transaction, error) {
	return _FErc20.Contract.UnderlyingReturns(&_FErc20.TransactOpts, a)
}
