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
var VaultTrackerBin = "0x60806040523480156200001157600080fd5b5060405162001f7038038062001f708339818101604052810190620000379190620000f5565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160018190555080600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001a8565b600081519050620000d88162000174565b92915050565b600081519050620000ef816200018e565b92915050565b600080604083850312156200010957600080fd5b60006200011985828601620000de565b92505060206200012c85828601620000c7565b9150509250929050565b600062000143826200014a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6200017f8162000136565b81146200018b57600080fd5b50565b62000199816200016a565b8114620001a557600080fd5b50565b611db880620001b86000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a01cfffb11610066578063a01cfffb14610170578063a622ee7c146101a0578063a9059cbb146101d2578063b7dd348314610202578063f851a440146102205761009e565b806319caf46c146100a3578063204f83f9146100d3578063613a28d1146100f15780636392a51f146101215780636b868d5114610152575b600080fd5b6100bd60048036038101906100b891906118fc565b61023e565b6040516100ca9190611b0d565b60405180910390f35b6100db610636565b6040516100e89190611b0d565b60405180910390f35b61010b60048036038101906101069190611925565b61063c565b6040516101189190611a92565b60405180910390f35b61013b600480360381019061013691906118fc565b610a48565b604051610149929190611b28565b60405180910390f35b61015a610ada565b6040516101679190611a92565b60405180910390f35b61018a60048036038101906101859190611925565b610beb565b6040516101979190611a92565b60405180910390f35b6101ba60048036038101906101b591906118fc565b6110b8565b6040516101c993929190611b51565b60405180910390f35b6101ec60048036038101906101e79190611925565b6110e2565b6040516101f99190611a92565b60405180910390f35b61020a611873565b6040516102179190611a77565b60405180910390f35b610228611899565b6040516102359190611a77565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c790611aed565b60405180910390fd5b6000600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050600081602001519050600080600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060011515600260009054906101000a900460ff16151514156103f95760006a52b7d2dcc80cd2e400000085604001516a52b7d2dcc80cd2e40000006003546103b59190611c20565b6103bf9190611bef565b6103c99190611c7a565b90506a52b7d2dcc80cd2e40000008560000151826103e79190611c20565b6103f19190611bef565b9250506104e2565b60006a52b7d2dcc80cd2e400000085604001516a52b7d2dcc80cd2e40000008473ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561046057600080fd5b505af1158015610474573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104989190611961565b6104a29190611c20565b6104ac9190611bef565b6104b69190611c7a565b90506a52b7d2dcc80cd2e40000008560000151826104d49190611c20565b6104de9190611bef565b9250505b81836104ee9190611b99565b92508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561053857600080fd5b505af115801561054c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105709190611961565b846040018181525050600084602001818152505083600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050828773ffffffffffffffffffffffffffffffffffffffff167f83a945bd12c713615b59a6e48a3467c05d1a7442350600d6f7fce6af9f7190e960405160405180910390a38295505050505050919050565b60015481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c590611aed565b60405180910390fd5b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508381600001511015610781576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077890611aad565b60405180910390fd5b600080600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060011515600260009054906101000a900460ff16151514156108355760006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006003546107f19190611c20565b6107fb9190611bef565b6108059190611c7a565b90506a52b7d2dcc80cd2e40000008460000151826108239190611c20565b61082d9190611bef565b92505061091e565b60006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000008473ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561089c57600080fd5b505af11580156108b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d49190611961565b6108de9190611c20565b6108e89190611bef565b6108f29190611c7a565b90506a52b7d2dcc80cd2e40000008460000151826109109190611c20565b61091a9190611bef565b9250505b81836020018181516109309190611b99565b9150818152505085836000018181516109499190611c7a565b915081815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561099857600080fd5b505af11580156109ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d09190611961565b83604001818152505082600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600194505050505092915050565b600080600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015491509150915091565b6000600154421015610b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1890611acd565b60405180910390fd5b6001600260006101000a81548160ff021916908315150217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610ba657600080fd5b505af1158015610bba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bde9190611961565b6003819055506001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7490611aed565b60405180910390fd5b6000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600082600001511115610fb657600060011515600260009054906101000a900460ff1615151415610dab5760006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000600354610d679190611c20565b610d719190611bef565b610d7b9190611c7a565b90506a52b7d2dcc80cd2e4000000846000015182610d999190611c20565b610da39190611bef565b915050610e94565b60006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000008573ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610e1257600080fd5b505af1158015610e26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4a9190611961565b610e549190611c20565b610e5e9190611bef565b610e689190611c7a565b90506a52b7d2dcc80cd2e4000000846000015182610e869190611c20565b610e909190611bef565b9150505b8083602001818151610ea69190611b99565b915081815250508583600001818151610ebf9190611b99565b915081815250508173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610f0e57600080fd5b505af1158015610f22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f469190611961565b83604001818152505082600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050506110ab565b848260000181815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561100857600080fd5b505af115801561101c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110409190611961565b82604001818152505081600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050505b6001935050505092915050565b60056020528060005260406000206000915090508060000154908060010154908060020154905083565b600080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508281600001511015611196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118d90611aad565b60405180910390fd5b600080600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060011515600260009054906101000a900460ff161515141561124a5760006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006003546112069190611c20565b6112109190611bef565b61121a9190611c7a565b90506a52b7d2dcc80cd2e40000008460000151826112389190611c20565b6112429190611bef565b925050611333565b60006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000008473ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156112b157600080fd5b505af11580156112c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e99190611961565b6112f39190611c20565b6112fd9190611bef565b6113079190611c7a565b90506a52b7d2dcc80cd2e40000008460000151826113259190611c20565b61132f9190611bef565b9250505b81836020018181516113459190611b99565b91508181525050848360000181815161135e9190611c7a565b915081815250508073ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156113ad57600080fd5b505af11580156113c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e59190611961565b83604001818152505082600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506000600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008160000151111561176157600060011515600260009054906101000a900460ff16151514156115565760006a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e40000006003546115129190611c20565b61151c9190611bef565b6115269190611c7a565b90506a52b7d2dcc80cd2e40000008360000151826115449190611c20565b61154e9190611bef565b91505061163f565b60006a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e40000008673ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156115bd57600080fd5b505af11580156115d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115f59190611961565b6115ff9190611c20565b6116099190611bef565b6116139190611c7a565b90506a52b7d2dcc80cd2e40000008360000151826116319190611c20565b61163b9190611bef565b9150505b80826020018181516116519190611b99565b91508181525050868260000181815161166a9190611b99565b915081815250508273ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156116b957600080fd5b505af11580156116cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f19190611961565b82604001818152505081600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505050611865565b85816000018181516117739190611b99565b915081815250508173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156117c257600080fd5b505af11580156117d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117fa9190611961565b81604001818152505080600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050505b600194505050505092915050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000813590506118cc81611d54565b92915050565b6000813590506118e181611d6b565b92915050565b6000815190506118f681611d6b565b92915050565b60006020828403121561190e57600080fd5b600061191c848285016118bd565b91505092915050565b6000806040838503121561193857600080fd5b6000611946858286016118bd565b9250506020611957858286016118d2565b9150509250929050565b60006020828403121561197357600080fd5b6000611981848285016118e7565b91505092915050565b61199381611cae565b82525050565b6119a281611cc0565b82525050565b60006119b5601c83611b88565b91507f416d6f756e742065786365656473207661756c742062616c616e6365000000006000830152602082019050919050565b60006119f5601d83611b88565b91507f6d6174757269747920686173206e6f74206265656e20726561636865640000006000830152602082019050919050565b6000611a35601483611b88565b91507f73656e646572206d7573742062652061646d696e0000000000000000000000006000830152602082019050919050565b611a7181611cec565b82525050565b6000602082019050611a8c600083018461198a565b92915050565b6000602082019050611aa76000830184611999565b92915050565b60006020820190508181036000830152611ac6816119a8565b9050919050565b60006020820190508181036000830152611ae6816119e8565b9050919050565b60006020820190508181036000830152611b0681611a28565b9050919050565b6000602082019050611b226000830184611a68565b92915050565b6000604082019050611b3d6000830185611a68565b611b4a6020830184611a68565b9392505050565b6000606082019050611b666000830186611a68565b611b736020830185611a68565b611b806040830184611a68565b949350505050565b600082825260208201905092915050565b6000611ba482611cec565b9150611baf83611cec565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611be457611be3611cf6565b5b828201905092915050565b6000611bfa82611cec565b9150611c0583611cec565b925082611c1557611c14611d25565b5b828204905092915050565b6000611c2b82611cec565b9150611c3683611cec565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c6f57611c6e611cf6565b5b828202905092915050565b6000611c8582611cec565b9150611c9083611cec565b925082821015611ca357611ca2611cf6565b5b828203905092915050565b6000611cb982611ccc565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b611d5d81611cae565b8114611d6857600080fd5b50565b611d7481611cec565b8114611d7f57600080fd5b5056fea2646970667358221220a3dcc767f36e1baf1d96a25b5495481e992acb2d1ffc57595ed95f5a569ceafa64736f6c63430008000033"

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
