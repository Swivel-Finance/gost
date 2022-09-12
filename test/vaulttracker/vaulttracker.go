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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Exception\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002bcc38038062002bcc83398181016040528101906200003891906200075f565b8460ff166101008160ff16815250508360e081815250508273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050604051806060016040528060008152602001600081526020016200011b87866200018860201b620015a61760201c565b8152506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050505050505062000972565b600060016006811115620001a157620001a0620007e7565b5b60ff168360ff161480620001d0575060026006811115620001c757620001c6620007e7565b5b60ff168360ff16145b1562000250578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000222573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000248919062000816565b905062000676565b60036006811115620002675762000266620007e7565b5b60ff168360ff1603620002ee578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002c0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002e6919062000816565b905062000676565b60046006811115620003055762000304620007e7565b5b60ff168360ff1603620004815760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000363573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000389919062000848565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620003f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000416919062000848565b6040518263ffffffff1660e01b81526004016200043491906200088b565b602060405180830381865afa15801562000452573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000478919062000816565b91505062000676565b60056006811115620004985762000497620007e7565b5b60ff168360ff160362000537578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401620004eb9190620008f5565b602060405180830381865afa15801562000509573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200052f919062000816565b905062000676565b6006808111156200054d576200054c620007e7565b5b60ff168360ff1603620005e9578173ffffffffffffffffffffffffffffffffffffffff1663bb2952fc670de0b6b3a76400006040518263ffffffff1660e01b81526004016200059d919062000955565b602060405180830381865afa158015620005bb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620005e1919062000816565b905062000676565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b81526004016200062f9190620008f5565b602060405180830381865afa1580156200064d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000673919062000816565b90505b92915050565b600080fd5b600060ff82169050919050565b620006998162000681565b8114620006a557600080fd5b50565b600081519050620006b9816200068e565b92915050565b6000819050919050565b620006d481620006bf565b8114620006e057600080fd5b50565b600081519050620006f481620006c9565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200072782620006fa565b9050919050565b62000739816200071a565b81146200074557600080fd5b50565b60008151905062000759816200072e565b92915050565b600080600080600060a086880312156200077e576200077d6200067c565b5b60006200078e88828901620006a8565b9550506020620007a188828901620006e3565b9450506040620007b48882890162000748565b9350506060620007c78882890162000748565b9250506080620007da8882890162000748565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000602082840312156200082f576200082e6200067c565b5b60006200083f84828501620006e3565b91505092915050565b6000602082840312156200086157620008606200067c565b5b6000620008718482850162000748565b91505092915050565b62000885816200071a565b82525050565b6000602082019050620008a260008301846200087a565b92915050565b6000819050919050565b6000819050919050565b6000620008dd620008d7620008d184620008a8565b620008b2565b620006bf565b9050919050565b620008ef81620008bc565b82525050565b60006020820190506200090c6000830184620008e4565b92915050565b6000819050919050565b60006200093d62000937620009318462000912565b620008b2565b620006bf565b9050919050565b6200094f816200091c565b82525050565b60006020820190506200096c600083018462000944565b92915050565b60805160a05160c05160e051610100516121cf620009fd60003960008181610ac30152610efd01526000610a730152600081816103520152818161133e015261149901526000818161037e0152818161084101528181610a9701528181610b2c01528181610e5101528181610f2301526111b4015260008181610ae4015261158401526121cf6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063613a28d111610097578063a01cfffb11610066578063a01cfffb146102a0578063a622ee7c146102d0578063b326258d14610302578063b7dd348314610332576100f5565b8063613a28d1146101f15780636392a51f1461022157806364ae3c9d146102525780638ce7442614610282576100f5565b806319caf46c116100d357806319caf46c14610166578063204f83f9146101965780632e25d2a6146101b457806343f48fbd146101d2576100f5565b8063012b264a146100fa57806311554c43146101185780631779467314610136575b600080fd5b610102610350565b60405161010f9190611a9e565b60405180910390f35b610120610374565b60405161012d9190611ad2565b60405180910390f35b610150600480360381019061014b9190611b4a565b61037a565b60405161015d9190611bb8565b60405180910390f35b610180600480360381019061017b9190611bd3565b61083d565b60405161018d9190611ad2565b60405180910390f35b61019e610a71565b6040516101ab9190611ad2565b60405180910390f35b6101bc610a95565b6040516101c99190611a9e565b60405180910390f35b6101da610ab9565b6040516101e8929190611c00565b60405180910390f35b61020b60048036038101906102069190611c29565b610b28565b6040516102189190611bb8565b60405180910390f35b61023b60048036038101906102369190611bd3565b610dcb565b604051610249929190611c00565b60405180910390f35b61026c60048036038101906102679190611c69565b610e4d565b6040516102799190611bb8565b60405180910390f35b61028a610efb565b6040516102979190611cb2565b60405180910390f35b6102ba60048036038101906102b59190611c29565b610f1f565b6040516102c79190611bb8565b60405180910390f35b6102ea60048036038101906102e59190611bd3565b611186565b6040516102f993929190611ccd565b60405180910390f35b61031c60048036038101906103179190611c29565b6111b0565b6040516103299190611bb8565b60405180910390f35b61033a611582565b6040516103479190611a9e565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461041757600080600033846040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260040161040e959493929190611d7a565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361049157602060008087876040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610488959493929190611e08565b60405180910390fd5b60008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015184111561055257601f84826000015188886040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610549959493929190611e96565b60405180910390fd5b60008061055d610ab9565b9150915060006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000008561058b9190611f18565b6105959190611fa1565b61059f9190611fd2565b905060006a52b7d2dcc80cd2e4000000856020015186600001516105c39190612006565b836105ce9190611f18565b6105d89190611fa1565b90508085602001516105ea9190612006565b8560200181815250508785600001516106039190611fd2565b856000018181525050828410610619578261061b565b835b856040018181525050846000808c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060008060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816000015111156107a8576a52b7d2dcc80cd2e400000081604001516a52b7d2dcc80cd2e4000000876107249190611f18565b61072e9190611fa1565b6107389190611fd2565b92506a52b7d2dcc80cd2e40000008160200151826000015161075a9190612006565b846107659190611f18565b61076f9190611fa1565b91508181602001516107819190612006565b81602001818152505088816000015161079a9190612006565b8160000181815250506107b3565b888160000181815250505b8385106107c057836107c2565b845b816040018181525050806000808c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019750505050505050509392505050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108da57600080600033846040517f6d4c6c890000000000000000000000000000000000000000000000000000000081526004016108d1959493929190611d7a565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050600081602001519050600080610959610ab9565b9150915060006a52b7d2dcc80cd2e400000085604001516a52b7d2dcc80cd2e4000000856109879190611f18565b6109919190611fa1565b61099b9190611fd2565b905060006a52b7d2dcc80cd2e4000000866020015187600001516109bf9190612006565b836109ca9190611f18565b6109d49190611fa1565b90508284106109e357826109e5565b835b866040018181525050856020016000815250856000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050508085610a639190612006565b975050505050505050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000806000610b087f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006115a6565b9050600060015411610b1a5780610b1e565b6001545b8192509250509091565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bc557600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610bbc959493929190611d7a565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508060000151841115610c8757601f8482600001518760006040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610c7e959493929190611e96565b60405180910390fd5b600080610c92610ab9565b9150915060006a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e400000085610cc09190611f18565b610cca9190611fa1565b610cd49190611fd2565b905060006a52b7d2dcc80cd2e400000085602001518660000151610cf89190612006565b83610d039190611f18565b610d0d9190611fa1565b9050808560200151610d1f9190612006565b856020018181525050878560000151610d389190611fd2565b8560000181815250508260015410610d505782610d52565b835b856040018181525050846000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506001965050505050505092915050565b60008060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015181602001519250925050915091565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610eea57600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610ee1959493929190611d7a565b60405180910390fd5b826001819055506001915050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fbc57600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401610fb3959493929190611d7a565b60405180910390fd5b600080610fc7610ab9565b9150915060008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816000015111156110f55760006a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e40000008661106d9190611f18565b6110779190611fa1565b6110819190611fd2565b905060006a52b7d2dcc80cd2e4000000836020015184600001516110a59190612006565b836110b09190611f18565b6110ba9190611fa1565b90508083602001516110cc9190612006565b8360200181815250508783600001516110e59190612006565b8360000181815250505050611100565b858160000181815250505b81831061110d578161110f565b825b816040018181525050806000808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600194505050505092915050565b60006020528060005260406000206000915090508060000154908060010154908060020154905083565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461124d57600080600033846040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401611244959493929190611d7a565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015184111561130f57601f8482600001518760006040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152600401611306959493929190611e96565b60405180910390fd5b83816000015161131f9190611fd2565b816000018181525050600080611333610ab9565b9150915060008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508181604001511461147a5760006a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e4000000866113f79190611f18565b6114019190611fa1565b61140b9190611fd2565b905060006a52b7d2dcc80cd2e40000008360200151846000015161142f9190612006565b8361143a9190611f18565b6114449190611fa1565b90508083602001516114569190612006565b83602001818152505083851061146c578361146e565b845b83604001818152505050505b86816000015161148a9190612006565b816000018181525050806000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000600160068111156115bc576115bb61203a565b5b60ff168360ff1614806115e75750600260068111156115de576115dd61203a565b5b60ff168360ff16145b15611662578173ffffffffffffffffffffffffffffffffffffffff1663bd6d894d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061165b919061207e565b9050611a57565b600360068111156116765761167561203a565b5b60ff168360ff16036116f8578173ffffffffffffffffffffffffffffffffffffffff166399530b066040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f1919061207e565b9050611a57565b6004600681111561170c5761170b61203a565b5b60ff168360ff160361187b5760008290508073ffffffffffffffffffffffffffffffffffffffff16637535d2466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611768573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061178c91906120c0565b73ffffffffffffffffffffffffffffffffffffffff1663d15e00538273ffffffffffffffffffffffffffffffffffffffff1663b16a19de6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181691906120c0565b6040518263ffffffff1660e01b81526004016118329190611a9e565b602060405180830381865afa15801561184f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611873919061207e565b915050611a57565b6005600681111561188f5761188e61203a565b5b60ff168360ff1603611927578173ffffffffffffffffffffffffffffffffffffffff1663010ad6d16a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b81526004016118df9190612128565b602060405180830381865afa1580156118fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611920919061207e565b9050611a57565b60068081111561193a5761193961203a565b5b60ff168360ff16036119cf578173ffffffffffffffffffffffffffffffffffffffff1663bb2952fc670de0b6b3a76400006040518263ffffffff1660e01b8152600401611987919061217e565b602060405180830381865afa1580156119a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c8919061207e565b9050611a57565b8173ffffffffffffffffffffffffffffffffffffffff166307a2d13a6a52b7d2dcc80cd2e40000006040518263ffffffff1660e01b8152600401611a139190612128565b602060405180830381865afa158015611a30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a54919061207e565b90505b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611a8882611a5d565b9050919050565b611a9881611a7d565b82525050565b6000602082019050611ab36000830184611a8f565b92915050565b6000819050919050565b611acc81611ab9565b82525050565b6000602082019050611ae76000830184611ac3565b92915050565b600080fd5b611afb81611a7d565b8114611b0657600080fd5b50565b600081359050611b1881611af2565b92915050565b611b2781611ab9565b8114611b3257600080fd5b50565b600081359050611b4481611b1e565b92915050565b600080600060608486031215611b6357611b62611aed565b5b6000611b7186828701611b09565b9350506020611b8286828701611b09565b9250506040611b9386828701611b35565b9150509250925092565b60008115159050919050565b611bb281611b9d565b82525050565b6000602082019050611bcd6000830184611ba9565b92915050565b600060208284031215611be957611be8611aed565b5b6000611bf784828501611b09565b91505092915050565b6000604082019050611c156000830185611ac3565b611c226020830184611ac3565b9392505050565b60008060408385031215611c4057611c3f611aed565b5b6000611c4e85828601611b09565b9250506020611c5f85828601611b35565b9150509250929050565b600060208284031215611c7f57611c7e611aed565b5b6000611c8d84828501611b35565b91505092915050565b600060ff82169050919050565b611cac81611c96565b82525050565b6000602082019050611cc76000830184611ca3565b92915050565b6000606082019050611ce26000830186611ac3565b611cef6020830185611ac3565b611cfc6040830184611ac3565b949350505050565b6000819050919050565b6000819050919050565b6000611d33611d2e611d2984611d04565b611d0e565b611c96565b9050919050565b611d4381611d18565b82525050565b6000611d64611d5f611d5a84611d04565b611d0e565b611ab9565b9050919050565b611d7481611d49565b82525050565b600060a082019050611d8f6000830188611d3a565b611d9c6020830187611d6b565b611da96040830186611d6b565b611db66060830185611a8f565b611dc36080830184611a8f565b9695505050505050565b6000819050919050565b6000611df2611ded611de884611dcd565b611d0e565b611c96565b9050919050565b611e0281611dd7565b82525050565b600060a082019050611e1d6000830188611df9565b611e2a6020830187611d6b565b611e376040830186611d6b565b611e446060830185611a8f565b611e516080830184611a8f565b9695505050505050565b6000819050919050565b6000611e80611e7b611e7684611e5b565b611d0e565b611c96565b9050919050565b611e9081611e65565b82525050565b600060a082019050611eab6000830188611e87565b611eb86020830187611ac3565b611ec56040830186611ac3565b611ed26060830185611a8f565b611edf6080830184611a8f565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611f2382611ab9565b9150611f2e83611ab9565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611f6757611f66611ee9565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611fac82611ab9565b9150611fb783611ab9565b925082611fc757611fc6611f72565b5b828204905092915050565b6000611fdd82611ab9565b9150611fe883611ab9565b925082820390508181111561200057611fff611ee9565b5b92915050565b600061201182611ab9565b915061201c83611ab9565b925082820190508082111561203457612033611ee9565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60008151905061207881611b1e565b92915050565b60006020828403121561209457612093611aed565b5b60006120a284828501612069565b91505092915050565b6000815190506120ba81611af2565b92915050565b6000602082840312156120d6576120d5611aed565b5b60006120e4848285016120ab565b91505092915050565b6000819050919050565b600061211261210d612108846120ed565b611d0e565b611ab9565b9050919050565b612122816120f7565b82525050565b600060208201905061213d6000830184612119565b92915050565b6000819050919050565b600061216861216361215e84612143565b611d0e565b611ab9565b9050919050565b6121788161214d565b82525050565b6000602082019050612193600083018461216f565b9291505056fea26469706673582212206d7c3ea5375dd814c9f1f2364d71e3b570fc96848cdf73c01570bd068058985964736f6c63430008100033",
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

// Rates is a free data retrieval call binding the contract method 0x43f48fbd.
//
// Solidity: function rates() view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCaller) Rates(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "rates")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Rates is a free data retrieval call binding the contract method 0x43f48fbd.
//
// Solidity: function rates() view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerSession) Rates() (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.Rates(&_VaultTracker.CallOpts)
}

// Rates is a free data retrieval call binding the contract method 0x43f48fbd.
//
// Solidity: function rates() view returns(uint256, uint256)
func (_VaultTracker *VaultTrackerCallerSession) Rates() (*big.Int, *big.Int, error) {
	return _VaultTracker.Contract.Rates(&_VaultTracker.CallOpts)
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
