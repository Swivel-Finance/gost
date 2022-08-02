// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaulttracker

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

// VaultTrackerMetaData contains all meta data concerning the VaultTracker contract.
var VaultTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Exception\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002b9238038062002b928339818101604052810190620000389190620006ac565b8460ff166101008160ff16815250508360e081815250508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050604051806060016040528060008152602001600081526020016200011b87866200018860201b620017241760201c565b8152506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505050505050506200085f565b600060016005811115620001a157620001a062000734565b5b60ff168360ff161480620001d0575060026005811115620001c757620001c662000734565b5b60ff168360ff16145b1562000250578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000222573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000248919062000763565b9050620005c3565b6003600581111562000267576200026662000734565b5b60ff168360ff1603620002ee578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002c0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002e6919062000763565b9050620005c3565b6004600581111562000305576200030462000734565b5b60ff168360ff1603620004815760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000363573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000389919062000795565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620003f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000416919062000795565b6040518263ffffffff1660e01b8152600401620004349190620007d8565b602060405180830381865afa15801562000452573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000478919062000763565b915050620005c3565b60058081111562000497576200049662000734565b5b60ff168360ff160362000536578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401620004ea919062000842565b602060405180830381865afa15801562000508573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200052e919062000763565b9050620005c3565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b81526004016200057c919062000842565b602060405180830381865afa1580156200059a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005c0919062000763565b90505b92915050565b600080fd5b600060ff82169050919050565b620005e681620005ce565b8114620005f257600080fd5b50565b6000815190506200060681620005db565b92915050565b6000819050919050565b62000621816200060c565b81146200062d57600080fd5b50565b600081519050620006418162000616565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620006748262000647565b9050919050565b620006868162000667565b81146200069257600080fd5b50565b600081519050620006a6816200067b565b92915050565b600080600080600060a08688031215620006cb57620006ca620005c9565b5b6000620006db88828901620005f5565b9550506020620006ee8882890162000630565b9450506040620007018882890162000695565b9350506060620007148882890162000695565b9250506080620007278882890162000695565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000602082840312156200077c576200077b620005c9565b5b60006200078c8482850162000630565b91505092915050565b600060208284031215620007ae57620007ad620005c9565b5b6000620007be8482850162000695565b91505092915050565b620007d28162000667565b82525050565b6000602082019050620007ef6000830184620007c7565b92915050565b6000819050919050565b6000819050919050565b60006200082a620008246200081e84620007f5565b620007ff565b6200060c565b9050919050565b6200083c8162000809565b82525050565b600060208201905062000859600083018462000831565b92915050565b60805160a05160c05160e05161010051612270620009226000396000818161059b015281816109d401528181610d1101528181610fef015281816110b501526114c601526000610b65015260008181610328015281816114220152611617015260008181610354015281816108c001528181610b8901528181610baf01528181610f430152818161101501526113170152600081816105bc015281816109f501528181610d32015281816110d6015281816114e7015261170201526122706000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636392a51f1161008c578063a01cfffb11610066578063a01cfffb14610276578063a622ee7c146102a6578063b326258d146102d8578063b7dd348314610308576100ea565b80636392a51f146101f757806364ae3c9d146102285780638ce7442614610258576100ea565b806319caf46c116100c857806319caf46c1461015b578063204f83f91461018b5780632e25d2a6146101a9578063613a28d1146101c7576100ea565b8063012b264a146100ef57806311554c431461010d578063177946731461012b575b600080fd5b6100f7610326565b6040516101049190611b73565b60405180910390f35b61011561034a565b6040516101229190611ba7565b60405180910390f35b61014560048036038101906101409190611c1f565b610350565b6040516101529190611c8d565b60405180910390f35b61017560048036038101906101709190611ca8565b6108bc565b6040516101829190611ba7565b60405180910390f35b610193610b63565b6040516101a09190611ba7565b60405180910390f35b6101b1610b87565b6040516101be9190611b73565b60405180910390f35b6101e160048036038101906101dc9190611cd5565b610bab565b6040516101ee9190611c8d565b60405180910390f35b610211600480360381019061020c9190611ca8565b610ebd565b60405161021f929190611d15565b60405180910390f35b610242600480360381019061023d9190611d3e565b610f3f565b60405161024f9190611c8d565b60405180910390f35b610260610fed565b60405161026d9190611d87565b60405180910390f35b610290600480360381019061028b9190611cd5565b611011565b60405161029d9190611c8d565b60405180910390f35b6102c060048036038101906102bb9190611ca8565b6112e9565b6040516102cf93929190611da2565b60405180910390f35b6102f260048036038101906102ed9190611cd5565b611313565b6040516102ff9190611c8d565b60405180910390f35b610310611700565b60405161031d9190611b73565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ed57600080600033846040517f6d4c6c890000000000000000000000000000000000000000000000000000000081526004016103e4959493929190611e4f565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361046757602060008087876040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260040161045e959493929190611edd565b60405180910390fd5b60008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050816000015185111561059457601f8583600001516000806040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260040161058b959493929190611f6b565b60405180910390fd5b60006105e07f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611724565b90506000806001541115610633576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006001546106189190611fed565b6106229190612076565b61062c91906120a7565b9050610672565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000008461065b9190611fed565b6106659190612076565b61066f91906120a7565b90505b60006a52b7d2dcc80cd2e40000008560000151836106909190611fed565b61069a9190612076565b905080856020018181516106ae91906120db565b9150818152505087856000018181516106c791906120a7565b9150818152505082856040018181525050846000808c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506000846000015111156108365760006001541115610795576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e400000060015461077a9190611fed565b6107849190612076565b61078e91906120a7565b91506107d4565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000856107bd9190611fed565b6107c79190612076565b6107d191906120a7565b91505b60006a52b7d2dcc80cd2e40000008560000151846107f29190611fed565b6107fc9190612076565b9050808560200181815161081091906120db565b91508181525050888560000181815161082991906120db565b9150818152505050610841565b878460000181815250505b82846040018181525050836000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600196505050505050509392505050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095957600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610950959493929190611e4f565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816020015190506000610a197f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611724565b90506000806001541115610a6c576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000600154610a519190611fed565b610a5b9190612076565b610a6591906120a7565b9050610aab565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e400000084610a949190611fed565b610a9e9190612076565b610aa891906120a7565b90505b60006a52b7d2dcc80cd2e4000000856000015183610ac99190611fed565b610ad39190612076565b9050828560400181815250506000856020018181525050846000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050508084610b5691906120db565b9650505050505050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c4857600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610c3f959493929190611e4f565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508060000151841115610d0a57601f8482600001516000806040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610d01959493929190611f6b565b60405180910390fd5b6000610d567f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611724565b90506000806001541115610da9576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000600154610d8e9190611fed565b610d989190612076565b610da291906120a7565b9050610de8565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e400000084610dd19190611fed565b610ddb9190612076565b610de591906120a7565b90505b60006a52b7d2dcc80cd2e4000000846000015183610e069190611fed565b610e109190612076565b90508084602001818151610e2491906120db565b915081815250508684600001818151610e3d91906120a7565b9150818152505082846040018181525050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b60008060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015181602001519250925050915091565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fdc57600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610fd3959493929190611e4f565b60405180910390fd5b826001819055506001915050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110ae57600080600033846040517f6d4c6c890000000000000000000000000000000000000000000000000000000081526004016110a5959493929190611e4f565b60405180910390fd5b60006110fa7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611724565b905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816000015111156112675760008060015411156111c5576a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e40000006001546111aa9190611fed565b6111b49190612076565b6111be91906120a7565b9050611204565b6a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e4000000856111ed9190611fed565b6111f79190612076565b61120191906120a7565b90505b60006a52b7d2dcc80cd2e40000008360000151836112229190611fed565b61122c9190612076565b9050808360200181815161124091906120db565b91508181525050868360000181815161125991906120db565b915081815250505050611272565b848160000181815250505b81816040018181525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506001935050505092915050565b60006020528060005260406000206000915090508060000154908060010154908060020154905083565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113b057600080600033846040517f6d4c6c890000000000000000000000000000000000000000000000000000000081526004016113a7959493929190611e4f565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905084826000018181516114b891906120a7565b91508181525050600061150b7f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000611724565b90506000818360400151146115f8576000600154111561156a576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e400000060015461154f9190611fed565b6115599190612076565b61156391906120a7565b90506115a9565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000846115929190611fed565b61159c9190612076565b6115a691906120a7565b90505b60006a52b7d2dcc80cd2e40000008460000151836115c79190611fed565b6115d19190612076565b905080846020018181516115e591906120db565b9150818152505082846040018181525050505b868360000181815161160a91906120db565b91508181525050826000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006001600581111561173a57611739612131565b5b60ff168360ff16148061176557506002600581111561175c5761175b612131565b5b60ff168360ff16145b156117e0578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d99190612175565b9050611b2c565b600360058111156117f4576117f3612131565b5b60ff168360ff1603611876578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa15801561184b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186f9190612175565b9050611b2c565b6004600581111561188a57611889612131565b5b60ff168360ff16036119f95760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190a91906121b7565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611970573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061199491906121b7565b6040518263ffffffff1660e01b81526004016119b09190611b73565b602060405180830381865afa1580156119cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119f19190612175565b915050611b2c565b600580811115611a0c57611a0b612131565b5b60ff168360ff1603611aa4578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401611a5c919061221f565b602060405180830381865afa158015611a79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9d9190612175565b9050611b2c565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401611ae8919061221f565b602060405180830381865afa158015611b05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b299190612175565b90505b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611b5d82611b32565b9050919050565b611b6d81611b52565b82525050565b6000602082019050611b886000830184611b64565b92915050565b6000819050919050565b611ba181611b8e565b82525050565b6000602082019050611bbc6000830184611b98565b92915050565b600080fd5b611bd081611b52565b8114611bdb57600080fd5b50565b600081359050611bed81611bc7565b92915050565b611bfc81611b8e565b8114611c0757600080fd5b50565b600081359050611c1981611bf3565b92915050565b600080600060608486031215611c3857611c37611bc2565b5b6000611c4686828701611bde565b9350506020611c5786828701611bde565b9250506040611c6886828701611c0a565b9150509250925092565b60008115159050919050565b611c8781611c72565b82525050565b6000602082019050611ca26000830184611c7e565b92915050565b600060208284031215611cbe57611cbd611bc2565b5b6000611ccc84828501611bde565b91505092915050565b60008060408385031215611cec57611ceb611bc2565b5b6000611cfa85828601611bde565b9250506020611d0b85828601611c0a565b9150509250929050565b6000604082019050611d2a6000830185611b98565b611d376020830184611b98565b9392505050565b600060208284031215611d5457611d53611bc2565b5b6000611d6284828501611c0a565b91505092915050565b600060ff82169050919050565b611d8181611d6b565b82525050565b6000602082019050611d9c6000830184611d78565b92915050565b6000606082019050611db76000830186611b98565b611dc46020830185611b98565b611dd16040830184611b98565b949350505050565b6000819050919050565b6000819050919050565b6000611e08611e03611dfe84611dd9565b611de3565b611d6b565b9050919050565b611e1881611ded565b82525050565b6000611e39611e34611e2f84611dd9565b611de3565b611b8e565b9050919050565b611e4981611e1e565b82525050565b600060a082019050611e646000830188611e0f565b611e716020830187611e40565b611e7e6040830186611e40565b611e8b6060830185611b64565b611e986080830184611b64565b9695505050505050565b6000819050919050565b6000611ec7611ec2611ebd84611ea2565b611de3565b611d6b565b9050919050565b611ed781611eac565b82525050565b600060a082019050611ef26000830188611ece565b611eff6020830187611e40565b611f0c6040830186611e40565b611f196060830185611b64565b611f266080830184611b64565b9695505050505050565b6000819050919050565b6000611f55611f50611f4b84611f30565b611de3565b611d6b565b9050919050565b611f6581611f3a565b82525050565b600060a082019050611f806000830188611f5c565b611f8d6020830187611b98565b611f9a6040830186611b98565b611fa76060830185611b64565b611fb46080830184611b64565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611ff882611b8e565b915061200383611b8e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561203c5761203b611fbe565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061208182611b8e565b915061208c83611b8e565b92508261209c5761209b612047565b5b828204905092915050565b60006120b282611b8e565b91506120bd83611b8e565b9250828210156120d0576120cf611fbe565b5b828203905092915050565b60006120e682611b8e565b91506120f183611b8e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561212657612125611fbe565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60008151905061216f81611bf3565b92915050565b60006020828403121561218b5761218a611bc2565b5b600061219984828501612160565b91505092915050565b6000815190506121b181611bc7565b92915050565b6000602082840312156121cd576121cc611bc2565b5b60006121db848285016121a2565b91505092915050565b6000819050919050565b60006122096122046121ff846121e4565b611de3565b611b8e565b9050919050565b612219816121ee565b82525050565b60006020820190506122346000830184612210565b9291505056fea26469706673582212201829cfcc449d0177787d5e815c0a0f3959401698dd408f5f7a1a97aa7c6d687b64736f6c634300080d0033",
}

// VaultTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultTrackerMetaData.ABI instead.
var VaultTrackerABI = VaultTrackerMetaData.ABI

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultTrackerMetaData.Bin instead.
var VaultTrackerBin = VaultTrackerMetaData.Bin

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, p uint8, m *big.Int, c common.Address, s common.Address, mp common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := VaultTrackerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultTrackerBin), backend, p, m, c, s, mp)
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

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_VaultTracker *VaultTrackerCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_VaultTracker *VaultTrackerSession) MarketPlace() (common.Address, error) {
	return _VaultTracker.Contract.MarketPlace(&_VaultTracker.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) MarketPlace() (common.Address, error) {
	return _VaultTracker.Contract.MarketPlace(&_VaultTracker.CallOpts)
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

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerCaller) Protocol(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerSession) Protocol() (uint8, error) {
	return _VaultTracker.Contract.Protocol(&_VaultTracker.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_VaultTracker *VaultTrackerCallerSession) Protocol() (uint8, error) {
	return _VaultTracker.Contract.Protocol(&_VaultTracker.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCaller) Swivel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "swivel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
}

// Swivel is a free data retrieval call binding the contract method 0x012b264a.
//
// Solidity: function swivel() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) Swivel() (common.Address, error) {
	return _VaultTracker.Contract.Swivel(&_VaultTracker.CallOpts)
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

	outstruct.Notional = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Redeemable = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExchangeRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

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

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) MatureVault(opts *bind.TransactOpts, c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "matureVault", c)
}

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerSession) MatureVault(c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts, c)
}

// MatureVault is a paid mutator transaction binding the contract method 0x64ae3c9d.
//
// Solidity: function matureVault(uint256 c) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) MatureVault(c *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.MatureVault(&_VaultTracker.TransactOpts, c)
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

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactor) TransferNotionalFee(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.contract.Transact(opts, "transferNotionalFee", f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
}

// TransferNotionalFee is a paid mutator transaction binding the contract method 0xb326258d.
//
// Solidity: function transferNotionalFee(address f, uint256 a) returns(bool)
func (_VaultTracker *VaultTrackerTransactorSession) TransferNotionalFee(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _VaultTracker.Contract.TransferNotionalFee(&_VaultTracker.TransactOpts, f, a)
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
