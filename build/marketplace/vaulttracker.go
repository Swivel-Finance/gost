// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulttracker

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

// VaultTrackerABI is the input ABI used to generate the binding from.
const VaultTrackerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
var VaultTrackerBin = "0x60806040523480156200001157600080fd5b50604051620023d9380380620023d98339818101604052810190620000379190620000f5565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001a8565b600081519050620000d88162000174565b92915050565b600081519050620000ef816200018e565b92915050565b600080604083850312156200010957600080fd5b60006200011985828601620000de565b92505060206200012c85828601620000c7565b9150509250929050565b600062000143826200014a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200017f8162000136565b81146200018b57600080fd5b50565b62000199816200016a565b8114620001a557600080fd5b50565b61222180620001b86000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80636b868d51116100715780636b868d511461018d578063a01cfffb146101ab578063a622ee7c146101db578063b7dd34831461020d578063beabacc81461022b578063f851a4401461025b576100a9565b806317794673146100ae57806319caf46c146100de578063204f83f91461010e578063613a28d11461012c5780636392a51f1461015c575b600080fd5b6100c860048036038101906100c39190611c7f565b610279565b6040516100d59190611ebb565b60405180910390f35b6100f860048036038101906100f39190611c56565b6103a9565b6040516101059190611f76565b60405180910390f35b6101166107bd565b6040516101239190611f76565b60405180910390f35b61014660048036038101906101419190611cce565b6107c3565b6040516101539190611ebb565b60405180910390f35b61017660048036038101906101719190611c56565b610c72565b604051610184929190611f91565b60405180910390f35b610195610d04565b6040516101a29190611ebb565b60405180910390f35b6101c560048036038101906101c09190611cce565b610e15565b6040516101d29190611ebb565b60405180910390f35b6101f560048036038101906101f09190611c56565b61131e565b60405161020493929190611fba565b60405180910390f35b610215611348565b6040516102229190611ea0565b60405180910390f35b61024560048036038101906102409190611c7f565b61136e565b6040516102529190611ebb565b60405180910390f35b610263611bf3565b6040516102709190611ea0565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290611ef6565b60405180910390fd5b61031585846107c3565b610354576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034b90611f36565b60405180910390fd5b61035e8484610e15565b61039d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039490611f16565b60405180910390fd5b60019150509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043290611ef6565b60405180910390fd5b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490506000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104f157600080fd5b505af1158015610505573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105299190611d0a565b905060011515600260009054906101000a900460ff1615151415610630576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546105af9190612089565b6105b99190612058565b6105c391906120e3565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548461061f9190612089565b6106299190612058565b9150610713565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000836106969190612089565b6106a09190612058565b6106aa91906120e3565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846107069190612089565b6107109190612058565b91505b818461071f9190612002565b935080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055506000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508395505050505050919050565b60015481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084c90611ef6565b60405180910390fd5b82600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d190611f56565b60405180910390fd5b6000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561094957600080fd5b505af115801561095d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109819190611d0a565b905060011515600260009054906101000a900460ff1615151415610a88576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000600354610a079190612089565b610a119190612058565b610a1b91906120e3565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015484610a779190612089565b610a819190612058565b9150610b6b565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000083610aee9190612089565b610af89190612058565b610b0291906120e3565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015484610b5e9190612089565b610b689190612058565b91505b81600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254610bbd9190612002565b9250508190555085600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254610c1691906120e3565b9250508190555080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600194505050505092915050565b600080600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b6000600154421015610d4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4290611ed6565b60405180910390fd5b6001600260006101000a81548160ff021916908315150217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610dd057600080fd5b505af1158015610de4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e089190611d0a565b6003819055506001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ea7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9e90611ef6565b60405180910390fd5b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610f1357600080fd5b505af1158015610f27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4b9190611d0a565b90506000600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015411156112835760008060011515600260009054906101000a900460ff16151514156110a0576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000060035461101f9190612089565b6110299190612058565b61103391906120e3565b91506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548361108f9190612089565b6110999190612058565b9050611183565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000856111069190612089565b6111109190612058565b61111a91906120e3565b91506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154836111769190612089565b6111809190612058565b90505b80600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546111d59190612002565b9250508190555085600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082825461122e9190612002565b9250508190555082600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505050611312565b83600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555080600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b60019250505092915050565b60056020528060005260406000206000915090508060000154908060010154908060020154905083565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611400576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f790611ef6565b60405180910390fd5b82600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015611485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147c90611f56565b60405180910390fd5b6000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156114f457600080fd5b505af1158015611508573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152c9190611d0a565b905060011515600260009054906101000a900460ff1615151415611633576a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546115b29190612089565b6115bc9190612058565b6115c691906120e3565b92506a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846116229190612089565b61162c9190612058565b9150611716565b6a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000836116999190612089565b6116a39190612058565b6116ad91906120e3565b92506a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846117099190612089565b6117139190612058565b91505b81600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546117689190612002565b9250508190555085600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546117c191906120e3565b9250508190555080600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055506000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541115611b4357600060011515600260009054906101000a900460ff1615151415611961576a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546118e09190612089565b6118ea9190612058565b6118f491906120e3565b93506a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154856119509190612089565b61195a9190612058565b9050611a44565b6a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000846119c79190612089565b6119d19190612058565b6119db91906120e3565b93506a52b7d2dcc80cd2e4000000600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015485611a379190612089565b611a419190612058565b90505b80600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254611a969190612002565b9250508190555086600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254611aef9190612002565b9250508190555081600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555050611be4565b85600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254611b959190612002565b9250508190555080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b60019450505050509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081359050611c26816121bd565b92915050565b600081359050611c3b816121d4565b92915050565b600081519050611c50816121d4565b92915050565b600060208284031215611c6857600080fd5b6000611c7684828501611c17565b91505092915050565b600080600060608486031215611c9457600080fd5b6000611ca286828701611c17565b9350506020611cb386828701611c17565b9250506040611cc486828701611c2c565b9150509250925092565b60008060408385031215611ce157600080fd5b6000611cef85828601611c17565b9250506020611d0085828601611c2c565b9150509250929050565b600060208284031215611d1c57600080fd5b6000611d2a84828501611c41565b91505092915050565b611d3c81612117565b82525050565b611d4b81612129565b82525050565b6000611d5e601d83611ff1565b91507f6d6174757269747920686173206e6f74206265656e20726561636865640000006000830152602082019050919050565b6000611d9e601483611ff1565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b6000611dde601383611ff1565b91507f616464206e6f74696f6e616c206661696c6564000000000000000000000000006000830152602082019050919050565b6000611e1e601683611ff1565b91507f72656d6f7665206e6f74696f6e616c206661696c6564000000000000000000006000830152602082019050919050565b6000611e5e601c83611ff1565b91507f616d6f756e742065786365656473207661756c742062616c616e6365000000006000830152602082019050919050565b611e9a81612155565b82525050565b6000602082019050611eb56000830184611d33565b92915050565b6000602082019050611ed06000830184611d42565b92915050565b60006020820190508181036000830152611eef81611d51565b9050919050565b60006020820190508181036000830152611f0f81611d91565b9050919050565b60006020820190508181036000830152611f2f81611dd1565b9050919050565b60006020820190508181036000830152611f4f81611e11565b9050919050565b60006020820190508181036000830152611f6f81611e51565b9050919050565b6000602082019050611f8b6000830184611e91565b92915050565b6000604082019050611fa66000830185611e91565b611fb36020830184611e91565b9392505050565b6000606082019050611fcf6000830186611e91565b611fdc6020830185611e91565b611fe96040830184611e91565b949350505050565b600082825260208201905092915050565b600061200d82612155565b915061201883612155565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561204d5761204c61215f565b5b828201905092915050565b600061206382612155565b915061206e83612155565b92508261207e5761207d61218e565b5b828204905092915050565b600061209482612155565b915061209f83612155565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156120d8576120d761215f565b5b828202905092915050565b60006120ee82612155565b91506120f983612155565b92508282101561210c5761210b61215f565b5b828203905092915050565b600061212282612135565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6121c681612117565b81146121d157600080fd5b50565b6121dd81612155565b81146121e857600080fd5b5056fea26469706673582212208ff2de384cc4c321a84ec3da56007b2d7f4f6a6dfc01ab9f312a037d84ec481864736f6c63430008000033"

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, m *big.Int, c common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultTrackerBin), backend, m, c)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// VaultTracker is an auto generated Go binding around an Ethereum contract.
type VaultTracker struct {
	VaultTrackerCaller     // Read-only binding to the contract
	VaultTrackerTransactor // Write-only binding to the contract
	VaultTrackerFilterer   // Log filterer for contract events
}

// VaultTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultTrackerSession struct {
	Contract     *VaultTracker     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultTrackerCallerSession struct {
	Contract *VaultTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTrackerTransactorSession struct {
	Contract     *VaultTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultTrackerRaw struct {
	Contract *VaultTracker // Generic contract binding to access the raw methods on
}

// VaultTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultTrackerCallerRaw struct {
	Contract *VaultTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTrackerTransactorRaw struct {
	Contract *VaultTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultTracker creates a new instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTracker(address common.Address, backend bind.ContractBackend) (*VaultTracker, error) {
	contract, err := bindVaultTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultTracker{VaultTrackerCaller: VaultTrackerCaller{contract: contract}, VaultTrackerTransactor: VaultTrackerTransactor{contract: contract}, VaultTrackerFilterer: VaultTrackerFilterer{contract: contract}}, nil
}

// NewVaultTrackerCaller creates a new read-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerCaller(address common.Address, caller bind.ContractCaller) (*VaultTrackerCaller, error) {
	contract, err := bindVaultTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerCaller{contract: contract}, nil
}

// NewVaultTrackerTransactor creates a new write-only instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTrackerTransactor, error) {
	contract, err := bindVaultTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerTransactor{contract: contract}, nil
}

// NewVaultTrackerFilterer creates a new log filterer instance of VaultTracker, bound to a specific deployed contract.
func NewVaultTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultTrackerFilterer, error) {
	contract, err := bindVaultTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerFilterer{contract: contract}, nil
}

// bindVaultTracker binds a generic wrapper to an already deployed contract.
func bindVaultTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.VaultTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.VaultTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultTracker *VaultTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultTracker *VaultTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultTracker *VaultTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultTracker.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerSession) Admin() (common.Address, error) {
	return _VaultTracker.Contract.Admin(&_VaultTracker.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) Admin() (common.Address, error) {
	return _VaultTracker.Contract.Admin(&_VaultTracker.CallOpts)
}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCaller) BalancesOf(opts *bind.CallOpts, o common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "balancesOf", o)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerSession) BalancesOf(o common.Address) (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.BalancesOf(&_VaultTracker.CallOpts, o)
}

// BalancesOf is a free data retrieval call binding the contract method 0x6392a51f.
//
// Solidity: function balancesOf(address o) view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCallerSession) BalancesOf(o common.Address) (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.BalancesOf(&_VaultTracker.CallOpts, o)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCaller) CTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "cTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// CTokenAddr is a free data retrieval call binding the contract method 0xb7dd3483.
//
// Solidity: function cTokenAddr() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) CTokenAddr() (common.Address, error) {
	return _VaultTracker.Contract.CTokenAddr(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) Maturity() (*big.Int, error) {
	return _VaultTracker.Contract.Maturity(&_VaultTracker.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerCaller) Vaults(opts *bind.CallOpts, arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Notional     *big.Int
		Redeemable   *big.Int
		ExchangeRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Notional = out[0].(*big.Int)
	outstruct.Redeemable = out[1].(*big.Int)
	outstruct.ExchangeRate = out[2].(*big.Int)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerSession) Vaults(arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	return _VaultTracker.Contract.Vaults(&_VaultTracker.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults(address ) view returns(uint256 notional, uint256 redeemable, uint256 exchangeRate)
func (_VaultTracker *VaultTrackerCallerSession) Vaults(arg0 common.Address) (struct {
	Notional     *big.Int
	Redeemable   *big.Int
	ExchangeRate *big.Int
}, error) {
	return _VaultTracker.Contract.Vaults(&_VaultTracker.CallOpts, arg0)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) AddNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "addNotional", o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// AddNotional is a paid mutator transaction binding the contract method 0xa01cfffb.
//
// Solidity: function addNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) AddNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.AddNotional(&_VaultTracker.TransactOpts, o, a)
}

// MatureVault is a paid mutator transaction binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() returns(bool)
func (_VaultTracker *VaultTrackerTransactor) MatureVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "matureVault")
}

// MatureVault is a paid mutator transaction binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() returns(bool)
func (_VaultTracker *VaultTrackerSession) MatureVault() (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts)
}

// MatureVault is a paid mutator transaction binding the contract method 0x6b868d51.
//
// Solidity: function matureVault() returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) MatureVault() (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactor) RedeemInterest(opts *bind.TransactOpts, o common.Address) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "redeemInterest", o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RedeemInterest is a paid mutator transaction binding the contract method 0x19caf46c.
//
// Solidity: function redeemInterest(address o) returns(uint256)
func (_VaultTracker *VaultTrackerTransactorSession) RedeemInterest(o common.Address) (*types.Transaction, error) {
	return _VaultTracker.Contract.RedeemInterest(&_VaultTracker.TransactOpts, o)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) RemoveNotional(opts *bind.TransactOpts, o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "removeNotional", o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// RemoveNotional is a paid mutator transaction binding the contract method 0x613a28d1.
//
// Solidity: function removeNotional(address o, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) RemoveNotional(o common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.RemoveNotional(&_VaultTracker.TransactOpts, o, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) Transfer(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transfer", f, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) Transfer(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.Transfer(&_VaultTracker.TransactOpts, f, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) Transfer(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.Transfer(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFrom(opts *bind.TransactOpts, f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFrom", f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}

// TransferNotionalFrom is a paid mutator transaction binding the contract method 0x17794673.
//
// Solidity: function transferNotionalFrom(address f, address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFrom(f common.Address, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFrom(&_VaultTracker.TransactOpts, f, t, a)
}
