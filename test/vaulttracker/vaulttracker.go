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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"adapterAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"addNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"balancesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"matureVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"}],\"name\":\"redeemInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"removeNotional\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swivel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferNotionalFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"notional\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162002357380380620023578339818101604052810190620000389190620002bc565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508361010081815250508273ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff1681525050604051806060016040528060008152602001600081526020018373ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b866040518263ffffffff1660e01b81526004016200016591906200033f565b6020604051808303816000875af115801562000185573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001ab91906200035c565b8152506000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050505050506200038e565b600080fd5b6000819050919050565b62000231816200021c565b81146200023d57600080fd5b50565b600081519050620002518162000226565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002848262000257565b9050919050565b620002968162000277565b8114620002a257600080fd5b50565b600081519050620002b6816200028b565b92915050565b60008060008060808587031215620002d957620002d862000217565b5b6000620002e98782880162000240565b9450506020620002fc87828801620002a5565b93505060406200030f87828801620002a5565b92505060606200032287828801620002a5565b91505092959194509250565b620003398162000277565b82525050565b60006020820190506200035660008301846200032e565b92915050565b60006020828403121562000375576200037462000217565b5b6000620003858482850162000240565b91505092915050565b60805160a05160c05160e05161010051611f06620004516000396000610c150152600081816103280152818161152e015261179501526000818161057101528181610a1001528181610d7f01528181611158015281816115cf01526118c80152600081816105ad01528181610a4c01528181610dbb015281816111940152818161160b01526118800152600081816103540152818161090a01528181610c3b01528181611025015281816110c70152818161142f01526118a40152611f066000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806364ae3c9d1161008c578063b326258d11610066578063b326258d1461029c578063b7dd3483146102cc578063f851a440146102ea578063f9d3af2b14610308576100ea565b806364ae3c9d1461020a578063a01cfffb1461023a578063a622ee7c1461026a576100ea565b806319caf46c116100c857806319caf46c1461015b578063204f83f91461018b578063613a28d1146101a95780636392a51f146101d9576100ea565b8063012b264a146100ef57806311554c431461010d578063177946731461012b575b600080fd5b6100f7610326565b604051610104919061192b565b60405180910390f35b61011561034a565b604051610122919061195f565b60405180910390f35b610145600480360381019061014091906119d7565b610350565b6040516101529190611a45565b60405180910390f35b61017560048036038101906101709190611a60565b610906565b604051610182919061195f565b60405180910390f35b610193610c13565b6040516101a0919061195f565b60405180910390f35b6101c360048036038101906101be9190611a8d565b610c37565b6040516101d09190611a45565b60405180910390f35b6101f360048036038101906101ee9190611a60565b610f9f565b604051610201929190611acd565b60405180910390f35b610224600480360381019061021f9190611af6565b611021565b6040516102319190611a45565b60405180910390f35b610254600480360381019061024f9190611a8d565b6110c3565b6040516102619190611a45565b60405180910390f35b610284600480360381019061027f9190611a60565b611401565b60405161029393929190611b23565b60405180910390f35b6102b660048036038101906102b19190611a8d565b61142b565b6040516102c39190611a45565b60405180910390f35b6102d461187e565b6040516102e1919061192b565b60405180910390f35b6102f26118a2565b6040516102ff919061192b565b60405180910390f35b6103106118c6565b60405161031d919061192b565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60015481565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d890611bb7565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff160361044f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044690611c23565b60405180910390fd5b60008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050848260000151101561056c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056390611c8f565b60405180910390fd5b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b81526004016105e8919061192b565b6020604051808303816000875af1158015610607573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061062b9190611cc4565b90506000600154111561067d576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006001546106629190611d20565b61066c9190611da9565b6106769190611dda565b91506106bc565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000836106a59190611d20565b6106af9190611da9565b6106b99190611dda565b91505b60006a52b7d2dcc80cd2e40000008560000151846106da9190611d20565b6106e49190611da9565b905080856020018181516106f89190611e0e565b9150818152505087856000018181516107119190611dda565b9150818152505081856040018181525050846000808c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060008460000151111561088057600060015411156107df576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e40000006001546107c49190611d20565b6107ce9190611da9565b6107d89190611dda565b925061081e565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000846108079190611d20565b6108119190611da9565b61081b9190611dda565b92505b60006a52b7d2dcc80cd2e400000085600001518561083c9190611d20565b6108469190611da9565b9050808560200181815161085a9190611e0e565b9150818152505088856000018181516108739190611e0e565b915081815250505061088b565b878460000181815250505b81846040018181525050836000808b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050600196505050505050509392505050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098e90611bb7565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090506000816020015190506000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610a87919061192b565b6020604051808303816000875af1158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca9190611cc4565b905060006001541115610b1c576a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e4000000600154610b019190611d20565b610b0b9190611da9565b610b159190611dda565b9150610b5b565b6a52b7d2dcc80cd2e400000084604001516a52b7d2dcc80cd2e400000083610b449190611d20565b610b4e9190611da9565b610b589190611dda565b91505b60006a52b7d2dcc80cd2e4000000856000015184610b799190611d20565b610b839190611da9565b9050818560400181815250506000856020018181525050846000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050508084610c069190611e0e565b9650505050505050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbf90611bb7565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082015481526020016001820154815260200160028201548152505090508381600001511015610d7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7190611eb0565b60405180910390fd5b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401610df6919061192b565b6020604051808303816000875af1158015610e15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e399190611cc4565b905060006001541115610e8b576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000600154610e709190611d20565b610e7a9190611da9565b610e849190611dda565b9150610eca565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e400000083610eb39190611d20565b610ebd9190611da9565b610ec79190611dda565b91505b60006a52b7d2dcc80cd2e4000000846000015184610ee89190611d20565b610ef29190611da9565b90508084602001818151610f069190611e0e565b915081815250508684600001818151610f1f9190611dda565b9150818152505081846040018181525050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b60008060008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180606001604052908160008201548152602001600182015481526020016002820154815250509050806000015181602001519250925050915091565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a990611bb7565b60405180910390fd5b826001819055506001915050919050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611154576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114b90611bb7565b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b81526004016111cf919061192b565b6020604051808303816000875af11580156111ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112129190611cc4565b905060008060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008160000151111561137f5760008060015411156112dd576a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e40000006001546112c29190611d20565b6112cc9190611da9565b6112d69190611dda565b905061131c565b6a52b7d2dcc80cd2e400000082604001516a52b7d2dcc80cd2e4000000856113059190611d20565b61130f9190611da9565b6113199190611dda565b90505b60006a52b7d2dcc80cd2e400000083600001518361133a9190611d20565b6113449190611da9565b905080836020018181516113589190611e0e565b9150818152505086836000018181516113719190611e0e565b91508181525050505061138a565b848160000181815250505b81816040018181525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155604082015181600201559050506001935050505092915050565b60006020528060005260406000206000915090508060000154908060010154908060020154905083565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b390611bb7565b60405180910390fd5b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905060008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050905084826000018181516115c49190611dda565b9150818152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dc3b7c8b7f00000000000000000000000000000000000000000000000000000000000000006040518263ffffffff1660e01b8152600401611646919061192b565b6020604051808303816000875af1158015611665573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116899190611cc4565b905060008183604001511461177657600060015411156116e8576a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e40000006001546116cd9190611d20565b6116d79190611da9565b6116e19190611dda565b9050611727565b6a52b7d2dcc80cd2e400000083604001516a52b7d2dcc80cd2e4000000846117109190611d20565b61171a9190611da9565b6117249190611dda565b90505b60006a52b7d2dcc80cd2e40000008460000151836117459190611d20565b61174f9190611da9565b905080846020018181516117639190611e0e565b9150818152505082846040018181525050505b86836000018181516117889190611e0e565b91508181525050826000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155905050836000808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101556040820151816002015590505060019550505050505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611915826118ea565b9050919050565b6119258161190a565b82525050565b6000602082019050611940600083018461191c565b92915050565b6000819050919050565b61195981611946565b82525050565b60006020820190506119746000830184611950565b92915050565b600080fd5b6119888161190a565b811461199357600080fd5b50565b6000813590506119a58161197f565b92915050565b6119b481611946565b81146119bf57600080fd5b50565b6000813590506119d1816119ab565b92915050565b6000806000606084860312156119f0576119ef61197a565b5b60006119fe86828701611996565b9350506020611a0f86828701611996565b9250506040611a20868287016119c2565b9150509250925092565b60008115159050919050565b611a3f81611a2a565b82525050565b6000602082019050611a5a6000830184611a36565b92915050565b600060208284031215611a7657611a7561197a565b5b6000611a8484828501611996565b91505092915050565b60008060408385031215611aa457611aa361197a565b5b6000611ab285828601611996565b9250506020611ac3858286016119c2565b9150509250929050565b6000604082019050611ae26000830185611950565b611aef6020830184611950565b9392505050565b600060208284031215611b0c57611b0b61197a565b5b6000611b1a848285016119c2565b91505092915050565b6000606082019050611b386000830186611950565b611b456020830185611950565b611b526040830184611950565b949350505050565b600082825260208201905092915050565b7f73656e646572206d75737420626520617574686f72697a656400000000000000600082015250565b6000611ba1601983611b5a565b9150611bac82611b6b565b602082019050919050565b60006020820190508181036000830152611bd081611b94565b9050919050565b7f63616e6e6f74207472616e73666572206e6f74696f6e616c20746f2073656c66600082015250565b6000611c0d602083611b5a565b9150611c1882611bd7565b602082019050919050565b60006020820190508181036000830152611c3c81611c00565b9050919050565b7f616d6f756e74206578636565647320617661696c61626c652062616c616e6365600082015250565b6000611c79602083611b5a565b9150611c8482611c43565b602082019050919050565b60006020820190508181036000830152611ca881611c6c565b9050919050565b600081519050611cbe816119ab565b92915050565b600060208284031215611cda57611cd961197a565b5b6000611ce884828501611caf565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611d2b82611946565b9150611d3683611946565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d6f57611d6e611cf1565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611db482611946565b9150611dbf83611946565b925082611dcf57611dce611d7a565b5b828204905092915050565b6000611de582611946565b9150611df083611946565b925082821015611e0357611e02611cf1565b5b828203905092915050565b6000611e1982611946565b9150611e2483611946565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611e5957611e58611cf1565b5b828201905092915050565b7f616d6f756e742065786365656473207661756c742062616c616e636500000000600082015250565b6000611e9a601c83611b5a565b9150611ea582611e64565b602082019050919050565b60006020820190508181036000830152611ec981611e8d565b905091905056fea2646970667358221220b9d91e4c393ce4854bc238c9417ac35366844caeb466654f654d5d44422ce2bf64736f6c634300080d0033",
}

// VaultTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultTrackerMetaData.ABI instead.
var VaultTrackerABI = VaultTrackerMetaData.ABI

// VaultTrackerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultTrackerMetaData.Bin instead.
var VaultTrackerBin = VaultTrackerMetaData.Bin

// DeployVaultTracker deploys a new Ethereum contract, binding an instance of VaultTracker to it.
func DeployVaultTracker(auth *bind.TransactOpts, backend bind.ContractBackend, m *big.Int, c common.Address, a common.Address, s common.Address) (common.Address, *types.Transaction, *VaultTracker, error) {
	parsed, err := VaultTrackerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultTrackerBin), backend, m, c, a, s)
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

// AdapterAddr is a free data retrieval call binding the contract method 0xf9d3af2b.
//
// Solidity: function adapterAddr() view returns(address)
func (_VaultTracker *VaultTrackerCaller) AdapterAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultTracker.contract.Call(opts, &out, "adapterAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdapterAddr is a free data retrieval call binding the contract method 0xf9d3af2b.
//
// Solidity: function adapterAddr() view returns(address)
func (_VaultTracker *VaultTrackerSession) AdapterAddr() (common.Address, error) {
	return _VaultTracker.Contract.AdapterAddr(&_VaultTracker.CallOpts)
}

// AdapterAddr is a free data retrieval call binding the contract method 0xf9d3af2b.
//
// Solidity: function adapterAddr() view returns(address)
func (_VaultTracker *VaultTrackerCallerSession) AdapterAddr() (common.Address, error) {
	return _VaultTracker.Contract.AdapterAddr(&_VaultTracker.CallOpts)
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
