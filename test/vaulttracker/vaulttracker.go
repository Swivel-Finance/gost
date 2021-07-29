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
const VaultTrackerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
var VaultTrackerBin = "0x60e06040523480156200001157600080fd5b50604051620022df380380620022df8339818101604052810190620000379190620000e3565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508160c081815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b81525050505062000196565b600081519050620000c68162000162565b92915050565b600081519050620000dd816200017c565b92915050565b60008060408385031215620000f757600080fd5b60006200010785828601620000cc565b92505060206200011a85828601620000b5565b9150509250929050565b6000620001318262000138565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200016d8162000124565b81146200017957600080fd5b50565b620001878162000158565b81146200019357600080fd5b50565b60805160601c60a05160601c60c0516120cd6200021260003960008181610efd01526114630152600081816103b201528181610bd60152818161104e015281816114e1015281816116210152611aad01526000818161029a01528181610afc01528181610f36015281816115900152611ad101526120cd6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80636392a51f116100715780636392a51f146101a35780636b868d51146101d4578063a01cfffb146101f2578063a622ee7c14610222578063b7dd348314610254578063f851a44014610272576100b4565b806311554c43146100b957806317794673146100d757806319caf46c14610107578063204f83f914610137578063454c87b314610155578063613a28d114610173575b600080fd5b6100c1610290565b6040516100ce9190611d7e565b60405180910390f35b6100f160048036038101906100ec9190611b5b565b610296565b6040516100fe9190611ce3565b60405180910390f35b610121600480360381019061011c9190611b32565b610af8565b60405161012e9190611d7e565b60405180910390f35b61013f610efb565b60405161014c9190611d7e565b60405180910390f35b61015d610f1f565b60405161016a9190611ce3565b60405180910390f35b61018d60048036038101906101889190611baa565b610f32565b60405161019a9190611ce3565b60405180910390f35b6101bd60048036038101906101b89190611b32565b6113cf565b6040516101cb929190611d99565b60405180910390f35b6101dc61145f565b6040516101e99190611ce3565b60405180910390f35b61020c60048036038101906102079190611baa565b61158c565b6040516102199190611ce3565b60405180910390f35b61023c60048036038101906102379190611b32565b611a81565b60405161024b93929190611dc2565b60405180910390f35b61025c611aab565b6040516102699190611cc8565b60405180910390f35b61027a611acf565b6040516102879190611cc8565b60405180910390f35b60015481565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610327576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031e90611d1e565b60405180910390fd5b826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156103ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a290611d3e565b60405180910390fd5b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561041857600080fd5b505af115801561042c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104509190611be6565b9050600260009054906101000a900460ff161561054e576a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006001546104ce9190611e91565b6104d89190611e60565b6104e29190611eeb565b92506a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548461053d9190611e91565b6105479190611e60565b915061062f565b6a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000836105b39190611e91565b6105bd9190611e60565b6105c79190611eeb565b92506a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846106229190611e91565b61062c9190611e60565b91505b816000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546106809190611e0a565b92505081905550856000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546106d89190611eeb565b92505081905550806000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555060008060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541115610a4a576000600260009054906101000a900460ff161561086d576a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006001546107ed9190611e91565b6107f79190611e60565b6108019190611eeb565b93506a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548561085c9190611e91565b6108669190611e60565b905061094e565b6a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000846108d29190611e91565b6108dc9190611e60565b6108e69190611eeb565b93506a52b7d2dcc80cd2e40000006000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154856109419190611e91565b61094b9190611e60565b90505b806000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461099f9190611e0a565b92505081905550866000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546109f79190611e0a565b92505081905550816000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555050610ae9565b856000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254610a9b9190611e0a565b92505081905550806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b60019450505050509392505050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8090611d1e565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905060008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610c3c57600080fd5b505af1158015610c50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c749190611be6565b9050600260009054906101000a900460ff1615610d72576a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000600154610cf29190611e91565b610cfc9190611e60565b610d069190611eeb565b92506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015484610d619190611e91565b610d6b9190611e60565b9150610e53565b6a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000083610dd79190611e91565b610de19190611e60565b610deb9190611eeb565b92506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015484610e469190611e91565b610e509190611e60565b91505b8184610e5f9190611e0a565b9350806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555060008060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508395505050505050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600260009054906101000a900460ff1681565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fc3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fba90611d1e565b60405180910390fd5b826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015611047576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103e90611d5e565b60405180910390fd5b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156110b457600080fd5b505af11580156110c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ec9190611be6565b9050600260009054906101000a900460ff16156111ea576a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000060015461116a9190611e91565b6111749190611e60565b61117e9190611eeb565b92506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846111d99190611e91565b6111e39190611e60565b91506112cb565b6a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000008361124f9190611e91565b6112599190611e60565b6112639190611eeb565b92506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846112be9190611e91565b6112c89190611e60565b91505b816000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461131c9190611e0a565b92505081905550856000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546113749190611eeb565b92505081905550806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600194505050505092915050565b6000806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b60007f00000000000000000000000000000000000000000000000000000000000000004210156114c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114bb90611cfe565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561154757600080fd5b505af115801561155b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061157f9190611be6565b6001819055506001905090565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461161d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161161490611d1e565b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561168757600080fd5b505af115801561169b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bf9190611be6565b905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015411156119e857600080600260009054906101000a900460ff161561180a576a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000060015461178a9190611e91565b6117949190611e60565b61179e9190611eeb565b91506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154836117f99190611e91565b6118039190611e60565b90506118eb565b6a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000008561186f9190611e91565b6118799190611e60565b6118839190611eeb565b91506a52b7d2dcc80cd2e40000006000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154836118de9190611e91565b6118e89190611e60565b90505b806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461193c9190611e0a565b92505081905550856000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546119949190611e0a565b92505081905550826000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505050611a75565b836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550806000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b60019250505092915050565b60006020528060005260406000206000915090508060000154908060010154908060020154905083565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600081359050611b0281612069565b92915050565b600081359050611b1781612080565b92915050565b600081519050611b2c81612080565b92915050565b600060208284031215611b4457600080fd5b6000611b5284828501611af3565b91505092915050565b600080600060608486031215611b7057600080fd5b6000611b7e86828701611af3565b9350506020611b8f86828701611af3565b9250506040611ba086828701611b08565b9150509250925092565b60008060408385031215611bbd57600080fd5b6000611bcb85828601611af3565b9250506020611bdc85828601611b08565b9150509250929050565b600060208284031215611bf857600080fd5b6000611c0684828501611b1d565b91505092915050565b611c1881611f1f565b82525050565b611c2781611f31565b82525050565b6000611c3a601d83611df9565b9150611c4582611fc5565b602082019050919050565b6000611c5d601483611df9565b9150611c6882611fee565b602082019050919050565b6000611c80602083611df9565b9150611c8b82612017565b602082019050919050565b6000611ca3601c83611df9565b9150611cae82612040565b602082019050919050565b611cc281611f5d565b82525050565b6000602082019050611cdd6000830184611c0f565b92915050565b6000602082019050611cf86000830184611c1e565b92915050565b60006020820190508181036000830152611d1781611c2d565b9050919050565b60006020820190508181036000830152611d3781611c50565b9050919050565b60006020820190508181036000830152611d5781611c73565b9050919050565b60006020820190508181036000830152611d7781611c96565b9050919050565b6000602082019050611d936000830184611cb9565b92915050565b6000604082019050611dae6000830185611cb9565b611dbb6020830184611cb9565b9392505050565b6000606082019050611dd76000830186611cb9565b611de46020830185611cb9565b611df16040830184611cb9565b949350505050565b600082825260208201905092915050565b6000611e1582611f5d565b9150611e2083611f5d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611e5557611e54611f67565b5b828201905092915050565b6000611e6b82611f5d565b9150611e7683611f5d565b925082611e8657611e85611f96565b5b828204905092915050565b6000611e9c82611f5d565b9150611ea783611f5d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ee057611edf611f67565b5b828202905092915050565b6000611ef682611f5d565b9150611f0183611f5d565b925082821015611f1457611f13611f67565b5b828203905092915050565b6000611f2a82611f3d565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f6d6174757269747920686173206e6f74206265656e2072656163686564000000600082015250565b7f73656e646572206d7573742062652061646d696e000000000000000000000000600082015250565b7f616d6f756e74206578636565647320617661696c61626c652062616c616e6365600082015250565b7f616d6f756e742065786365656473207661756c742062616c616e636500000000600082015250565b61207281611f1f565b811461207d57600080fd5b50565b61208981611f5d565b811461209457600080fd5b5056fea2646970667358221220156e9f132406657630fbfcb945aec6605fc28b867f2201d31c25710fbaa471a464736f6c63430008040033"

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

// Matured is a free data retrieval call binding the contract method 0x454c87b3.
//
// Solidity: function matured() view returns(bool)
func (_VaultTracker *VaultTrackerCaller) Matured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "matured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Matured is a free data retrieval call binding the contract method 0x454c87b3.
//
// Solidity: function matured() view returns(bool)
func (_VaultTracker *VaultTrackerSession) Matured() (bool, error) {
	return _VaultTracker.Contract.Matured(&_VaultTracker.CallOpts)
}

// Matured is a free data retrieval call binding the contract method 0x454c87b3.
//
// Solidity: function matured() view returns(bool)
func (_VaultTracker *VaultTrackerCallerSession) Matured() (bool, error) {
	return _VaultTracker.Contract.Matured(&_VaultTracker.CallOpts)
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

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerCaller) MaturityRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "maturityRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerSession) MaturityRate() (*big.Int, error) {
	return _VaultTracker.Contract.MaturityRate(&_VaultTracker.CallOpts)
}

// MaturityRate is a free data retrieval call binding the contract method 0x11554c43.
//
// Solidity: function maturityRate() view returns(uint256)
func (_VaultTracker *VaultTrackerCallerSession) MaturityRate() (*big.Int, error) {
	return _VaultTracker.Contract.MaturityRate(&_VaultTracker.CallOpts)
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
