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
const VaultTrackerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RedeemInterest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
var VaultTrackerBin = "0x60806040523480156200001157600080fd5b5060405162002111380380620021118339818101604052810190620000379190620000f5565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001a8565b600081519050620000d88162000174565b92915050565b600081519050620000ef816200018e565b92915050565b600080604083850312156200010957600080fd5b60006200011985828601620000de565b92505060206200012c85828601620000c7565b9150509250929050565b600062000143826200014a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200017f8162000136565b81146200018b57600080fd5b50565b62000199816200016a565b8114620001a557600080fd5b50565b611f5980620001b86000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a01cfffb11610066578063a01cfffb14610170578063a622ee7c146101a0578063a9059cbb146101d2578063b7dd348314610202578063f851a440146102205761009e565b806319caf46c146100a3578063204f83f9146100d3578063613a28d1146100f15780636392a51f146101215780636b868d5114610152575b600080fd5b6100bd60048036038101906100b89190611a9d565b61023e565b6040516100ca9190611cae565b60405180910390f35b6100db610696565b6040516100e89190611cae565b60405180910390f35b61010b60048036038101906101069190611ac6565b61069c565b6040516101189190611c33565b60405180910390f35b61013b60048036038101906101369190611a9d565b610b4b565b604051610149929190611cc9565b60405180910390f35b61015a610bdd565b6040516101679190611c33565b60405180910390f35b61018a60048036038101906101859190611ac6565b610cee565b6040516101979190611c33565b60405180910390f35b6101ba60048036038101906101b59190611a9d565b6111f7565b6040516101c993929190611cf2565b60405180910390f35b6101ec60048036038101906101e79190611ac6565b611221565b6040516101f99190611c33565b60405180910390f35b61020a611a14565b6040516102179190611c18565b60405180910390f35b610228611a3a565b6040516102359190611c18565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c790611c8e565b60405180910390fd5b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490506000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561038657600080fd5b505af115801561039a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103be9190611b02565b905060011515600260009054906101000a900460ff16151514156104c5576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546104449190611dc1565b61044e9190611d90565b6104589190611e1b565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846104b49190611dc1565b6104be9190611d90565b91506105a8565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000008361052b9190611dc1565b6105359190611d90565b61053f9190611e1b565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548461059b9190611dc1565b6105a59190611d90565b91505b81846105b49190611d3a565b935080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055506000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550838773ffffffffffffffffffffffffffffffffffffffff167f83a945bd12c713615b59a6e48a3467c05d1a7442350600d6f7fce6af9f7190e960405160405180910390a38395505050505050919050565b60015481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461072e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072590611c8e565b60405180910390fd5b82600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156107b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107aa90611c4e565b60405180910390fd5b6000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561082257600080fd5b505af1158015610836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085a9190611b02565b905060011515600260009054906101000a900460ff1615151415610961576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546108e09190611dc1565b6108ea9190611d90565b6108f49190611e1b565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846109509190611dc1565b61095a9190611d90565b9150610a44565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000836109c79190611dc1565b6109d19190611d90565b6109db9190611e1b565b92506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015484610a379190611dc1565b610a419190611d90565b91505b81600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016000828254610a969190611d3a565b9250508190555085600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254610aef9190611e1b565b9250508190555080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020181905550600194505050505092915050565b600080600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b6000600154421015610c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1b90611c6e565b60405180910390fd5b6001600260006101000a81548160ff021916908315150217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610ca957600080fd5b505af1158015610cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce19190611b02565b6003819055506001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7790611c8e565b60405180910390fd5b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610dec57600080fd5b505af1158015610e00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e249190611b02565b90506000600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154111561115c5760008060011515600260009054906101000a900460ff1615151415610f79576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000600354610ef89190611dc1565b610f029190611d90565b610f0c9190611e1b565b91506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015483610f689190611dc1565b610f729190611d90565b905061105c565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e400000085610fdf9190611dc1565b610fe99190611d90565b610ff39190611e1b565b91506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548361104f9190611dc1565b6110599190611d90565b90505b80600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546110ae9190611d3a565b9250508190555085600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546111079190611d3a565b9250508190555082600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555050506111eb565b83600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555080600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b60019250505092915050565b60056020528060005260406000206000915090508060000154908060010154908060020154905083565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015410156112a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129f90611c4e565b60405180910390fd5b6000806000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561131757600080fd5b505af115801561132b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134f9190611b02565b905060011515600260009054906101000a900460ff1615151415611456576a52b7d2dcc80cd2e4000000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546113d59190611dc1565b6113df9190611d90565b6113e99190611e1b565b92506a52b7d2dcc80cd2e4000000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154846114459190611dc1565b61144f9190611d90565b9150611539565b6a52b7d2dcc80cd2e4000000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000836114bc9190611dc1565b6114c69190611d90565b6114d09190611e1b565b92506a52b7d2dcc80cd2e4000000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548461152c9190611dc1565b6115369190611d90565b91505b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461158b9190611d3a565b9250508190555084600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546115e49190611e1b565b9250508190555080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055506000600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154111561196657600060011515600260009054906101000a900460ff1615151415611784576a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e40000006003546117039190611dc1565b61170d9190611d90565b6117179190611e1b565b93506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154856117739190611dc1565b61177d9190611d90565b9050611867565b6a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546a52b7d2dcc80cd2e4000000846117ea9190611dc1565b6117f49190611d90565b6117fe9190611e1b565b93506a52b7d2dcc80cd2e4000000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001548561185a9190611dc1565b6118649190611d90565b90505b80600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546118b99190611d3a565b9250508190555085600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546119129190611d3a565b9250508190555081600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002018190555050611a07565b84600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546119b89190611d3a565b9250508190555080600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505b6001935050505092915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081359050611a6d81611ef5565b92915050565b600081359050611a8281611f0c565b92915050565b600081519050611a9781611f0c565b92915050565b600060208284031215611aaf57600080fd5b6000611abd84828501611a5e565b91505092915050565b60008060408385031215611ad957600080fd5b6000611ae785828601611a5e565b9250506020611af885828601611a73565b9150509250929050565b600060208284031215611b1457600080fd5b6000611b2284828501611a88565b91505092915050565b611b3481611e4f565b82525050565b611b4381611e61565b82525050565b6000611b56601c83611d29565b91507f416d6f756e742065786365656473207661756c742062616c616e6365000000006000830152602082019050919050565b6000611b96601d83611d29565b91507f6d6174757269747920686173206e6f74206265656e20726561636865640000006000830152602082019050919050565b6000611bd6601483611d29565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b611c1281611e8d565b82525050565b6000602082019050611c2d6000830184611b2b565b92915050565b6000602082019050611c486000830184611b3a565b92915050565b60006020820190508181036000830152611c6781611b49565b9050919050565b60006020820190508181036000830152611c8781611b89565b9050919050565b60006020820190508181036000830152611ca781611bc9565b9050919050565b6000602082019050611cc36000830184611c09565b92915050565b6000604082019050611cde6000830185611c09565b611ceb6020830184611c09565b9392505050565b6000606082019050611d076000830186611c09565b611d146020830185611c09565b611d216040830184611c09565b949350505050565b600082825260208201905092915050565b6000611d4582611e8d565b9150611d5083611e8d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611d8557611d84611e97565b5b828201905092915050565b6000611d9b82611e8d565b9150611da683611e8d565b925082611db657611db5611ec6565b5b828204905092915050565b6000611dcc82611e8d565b9150611dd783611e8d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e1057611e0f611e97565b5b828202905092915050565b6000611e2682611e8d565b9150611e3183611e8d565b925082821015611e4457611e43611e97565b5b828203905092915050565b6000611e5a82611e6d565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b611efe81611e4f565b8114611f0957600080fd5b50565b611f1581611e8d565b8114611f2057600080fd5b5056fea26469706673582212202fc44590c2c51171ab37baed5d2d9368c5acbd989ac13d7ade03242d724b8efd64736f6c63430008000033"

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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) Transfer(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transfer", t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.Transfer(&_VaultTracker.TransactOpts, t, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address t, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) Transfer(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.Transfer(&_VaultTracker.TransactOpts, t, a)
}

// VaultTrackerRedeemInterestIterator is returned from FilterRedeemInterest and is used to iterate over the raw logs and unpacked data for RedeemInterest events raised by the VaultTracker contract.
type VaultTrackerRedeemInterestIterator struct {
	Event *VaultTrackerRedeemInterest // Event containing the contract specifics and raw log

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
func (it *VaultTrackerRedeemInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTrackerRedeemInterest)
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
		it.Event = new(VaultTrackerRedeemInterest)
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
func (it *VaultTrackerRedeemInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTrackerRedeemInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTrackerRedeemInterest represents a RedeemInterest event raised by the VaultTracker contract.
type VaultTrackerRedeemInterest struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeemInterest is a free log retrieval operation binding the contract event 0x83a945bd12c713615b59a6e48a3467c05d1a7442350600d6f7fce6af9f7190e9.
//
// Solidity: event RedeemInterest(address indexed owner, uint256 indexed amount)
func (_VaultTracker *VaultTrackerFilterer) FilterRedeemInterest(opts *bind.FilterOpts, owner []common.Address, amount []*big.Int) (*VaultTrackerRedeemInterestIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _VaultTracker.contract.FilterLogs(opts, "RedeemInterest", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &VaultTrackerRedeemInterestIterator{contract: _VaultTracker.contract, event: "RedeemInterest", logs: logs, sub: sub}, nil
}

// WatchRedeemInterest is a free log subscription operation binding the contract event 0x83a945bd12c713615b59a6e48a3467c05d1a7442350600d6f7fce6af9f7190e9.
//
// Solidity: event RedeemInterest(address indexed owner, uint256 indexed amount)
func (_VaultTracker *VaultTrackerFilterer) WatchRedeemInterest(opts *bind.WatchOpts, sink chan<- *VaultTrackerRedeemInterest, owner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _VaultTracker.contract.WatchLogs(opts, "RedeemInterest", ownerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTrackerRedeemInterest)
				if err := _VaultTracker.contract.UnpackLog(event, "RedeemInterest", log); err != nil {
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

// ParseRedeemInterest is a log parse operation binding the contract event 0x83a945bd12c713615b59a6e48a3467c05d1a7442350600d6f7fce6af9f7190e9.
//
// Solidity: event RedeemInterest(address indexed owner, uint256 indexed amount)
func (_VaultTracker *VaultTrackerFilterer) ParseRedeemInterest(log types.Log) (*VaultTrackerRedeemInterest, error) {
	event := new(VaultTrackerRedeemInterest)
	if err := _VaultTracker.contract.UnpackLog(event, "RedeemInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
