// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creator

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

// CreatorMetaData contains all meta data concerning the Creator contract.
var CreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Exception\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"p\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sw\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"d\",\"type\":\"uint8\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketPlace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"setMarketPlace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055611e57806100326000396000f3fe60806040523480156200001157600080fd5b50600436106200006f5760003560e01c80636363e86711620000565780636363e86714620000e7578063704b6c02146200012c578063f851a440146200014357600080fd5b80632e25d2a6146200007457806330568a8d14620000bf575b600080fd5b600154620000959073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b620000d6620000d036600462000611565b62000164565b6040519015158152602001620000b6565b620000fe620000f836600462000694565b620002c7565b6040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015201620000b6565b620000d66200013d36600462000611565b620004f7565b600054620000959073ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff16338114620001f0576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a4015b60405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff16156200027d576001546040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152602160048201526000602482018190526044820181905273ffffffffffffffffffffffffffffffffffffffff9092166064820152608481019190915260a401620001e7565b6001805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155915050919050565b600154600090819073ffffffffffffffffffffffffffffffffffffffff1633811462000353576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a401620001e7565b60015473ffffffffffffffffffffffffffffffffffffffff16620003df576001546040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152602260048201526000602482018190526044820181905273ffffffffffffffffffffffffffffffffffffffff9092166064820152608481019190915260a401620001e7565b60008d8d8d8d600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168d8d8d8d8d6040516200041b90620005cb565b620004309a99989796959493929190620007bb565b604051809103906000f0801580156200044d573d6000803e3d6000fd5b50905060008e8d8d8d600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200048790620005d9565b60ff9095168552602085019390935273ffffffffffffffffffffffffffffffffffffffff91821660408501528116606084015216608082015260a001604051809103906000f080158015620004e0573d6000803e3d6000fd5b50919f919e50909c50505050505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163381146200057f576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a401620001e7565b6000805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556001915050919050565b610ced806200084283390190565b6108f3806200152f83390190565b803573ffffffffffffffffffffffffffffffffffffffff811681146200060c57600080fd5b919050565b6000602082840312156200062457600080fd5b6200062f82620005e7565b9392505050565b803560ff811681146200060c57600080fd5b60008083601f8401126200065b57600080fd5b50813567ffffffffffffffff8111156200067457600080fd5b6020830191508360208285010111156200068d57600080fd5b9250929050565b6000806000806000806000806000806101008b8d031215620006b557600080fd5b620006c08b62000636565b9950620006d060208c01620005e7565b985060408b01359750620006e760608c01620005e7565b9650620006f760808c01620005e7565b955060a08b013567ffffffffffffffff808211156200071557600080fd5b620007238e838f0162000648565b909750955060c08d01359150808211156200073d57600080fd5b506200074c8d828e0162000648565b909450925062000761905060e08c0162000636565b90509295989b9194979a5092959850565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600061010060ff8d16835273ffffffffffffffffffffffffffffffffffffffff808d1660208501528b6040850152808b166060850152808a166080850152508060a08401526200080f818401888a62000772565b905082810360c08401526200082681868862000772565b91505060ff831660e08301529b9a505050505050505050505056fe60806040523480156200001157600080fd5b5060405162000ced38038062000ced83398101604081905262000034916200027c565b600380546001600160a01b03808a16610100026001600160a81b031990921660ff8c1617919091179091556004879055600580548783166001600160a01b03199182161790915560068054928716929091169190911790558251620000a1906007906020860190620000da565b508151620000b7906008906020850190620000da565b506009805460ff191660ff92909216919091179055506200038a95505050505050565b828054620000e8906200034e565b90600052602060002090601f0160209004810192826200010c576000855562000157565b82601f106200012757805160ff191683800117855562000157565b8280016001018555821562000157579182015b82811115620001575782518255916020019190600101906200013a565b506200016592915062000169565b5090565b5b808211156200016557600081556001016200016a565b805160ff811681146200019257600080fd5b919050565b80516001600160a01b03811681146200019257600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001d757600080fd5b81516001600160401b0380821115620001f457620001f4620001af565b604051601f8301601f19908116603f011681019082821181831017156200021f576200021f620001af565b816040528381526020925086838588010111156200023c57600080fd5b600091505b8382101562000260578582018301518183018401529082019062000241565b83821115620002725760008385830101525b9695505050505050565b600080600080600080600080610100898b0312156200029a57600080fd5b620002a58962000180565b9750620002b560208a0162000197565b965060408901519550620002cc60608a0162000197565b9450620002dc60808a0162000197565b60a08a01519094506001600160401b0380821115620002fa57600080fd5b620003088c838d01620001c5565b945060c08b01519150808211156200031f57600080fd5b506200032e8b828c01620001c5565b9250506200033f60e08a0162000180565b90509295985092959890939650565b600181811c908216806200036357607f821691505b6020821081036200038457634e487b7160e01b600052602260045260246000fd5b50919050565b610953806200039a6000396000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80638ce74426116100cd578063bba0ad3911610081578063e7ba677411610066578063e7ba6774146104c3578063ee4db5701461051d578063fdfe5f4d1461053d57600080fd5b8063bba0ad3914610437578063e541efa21461045757600080fd5b80639dc29fac116100b25780639dc29fac1461039b578063b4c4a4c8146103df578063b9bb928c146103f257600080fd5b80638ce744261461038657806395d89b411461039357600080fd5b806340c10f19116101245780636521b96a116101095780636521b96a146102fc57806369e527da146103435780636f307dc31461036357600080fd5b806340c10f19146102a257806352bc9430146102e757600080fd5b806323b872dd1161015557806323b872dd146101a15780632ba29d381461023e578063313ce5671461028357600080fd5b806306fdde0314610171578063204f83f91461018f575b600080fd5b610179610583565b60405161018691906106e4565b60405180910390f35b6004545b604051908152602001610186565b61022e6101af366004610780565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff93841681526020808201938452948416600090815260029095529320925183547fffffffffffffffffffffffff0000000000000000000000000000000000000000169216919091178255516001909101556009546301000000900460ff1690565b6040519015158152602001610186565b60065461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610186565b6009546102909060ff1681565b60405160ff9091168152602001610186565b61022e6102b03660046107bc565b73ffffffffffffffffffffffffffffffffffffffff9190911660009081526001602052604090205560095462010000900460ff1690565b6102fa6102f53660046107e6565b610611565b005b6102fa61030a366004610854565b600980549115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff909216919091179055565b60055461025e9073ffffffffffffffffffffffffffffffffffffffff1681565b600354610100900473ffffffffffffffffffffffffffffffffffffffff1661025e565b6003546102909060ff1681565b6101796106d7565b61022e6103a93660046107bc565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260208190526040902055600954610100900460ff1690565b6102fa6103ed36600461087d565b600455565b6102fa610400366004610854565b60098054911515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055565b610193610445366004610896565b60006020819052908152604090205481565b610497610465366004610896565b6002602052600090815260409020805460019091015473ffffffffffffffffffffffffffffffffffffffff9091169082565b6040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915201610186565b6102fa6104d1366004610896565b6003805473ffffffffffffffffffffffffffffffffffffffff909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b61019361052b366004610896565b60016020526000908152604090205481565b6102fa61054b366004610854565b6009805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055565b60078054610590906108b1565b80601f01602080910402602001604051908101604052809291908181526020018280546105bc906108b1565b80156106095780601f106105de57610100808354040283529160200191610609565b820191906000526020600020905b8154815290600101906020018083116105ec57829003601f168201915b505050505081565b6006546040517f52bc943000000000000000000000000000000000000000000000000000000000815260ff8816600482015273ffffffffffffffffffffffffffffffffffffffff8781166024830152604482018790528581166064830152848116608483015260a48201849052909116906352bc94309060c4016020604051808303816000875af11580156106aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ce9190610904565b50505050505050565b60088054610590906108b1565b600060208083528351808285015260005b81811015610711578581018301518582016040015282016106f5565b81811115610723576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461077b57600080fd5b919050565b60008060006060848603121561079557600080fd5b61079e84610757565b92506107ac60208501610757565b9150604084013590509250925092565b600080604083850312156107cf57600080fd5b6107d883610757565b946020939093013593505050565b60008060008060008060c087890312156107ff57600080fd5b863560ff8116811461081057600080fd5b955061081e60208801610757565b94506040870135935061083360608801610757565b925061084160808801610757565b915060a087013590509295509295509295565b60006020828403121561086657600080fd5b8135801515811461087657600080fd5b9392505050565b60006020828403121561088f57600080fd5b5035919050565b6000602082840312156108a857600080fd5b61087682610757565b600181811c908216806108c557607f821691505b6020821081036108fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561091657600080fd5b505191905056fea2646970667358221220ab82732e8b4d0f87d0e20dedbcdfd9a4708dd121141f10519c014c70dccfd0f764736f6c634300080d0033608060405234801561001057600080fd5b506040516108f33803806108f383398101604081905261002f916100a9565b6009805460ff191660ff9690961695909517909455600a92909255600480546001600160a01b03199081166001600160a01b039384161790915560058054821693831693909317909255600680549092169216919091179055610110565b80516001600160a01b03811681146100a457600080fd5b919050565b600080600080600060a086880312156100c157600080fd5b855160ff811681146100d257600080fd5b602087015190955093506100e86040870161008d565b92506100f66060870161008d565b91506101046080870161008d565b90509295509295909350565b6107d48061011f6000396000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c806382cac89c116100e3578063b7dd34831161008c578063d6cb2c0d11610066578063d6cb2c0d146105f4578063da3de9e914610607578063e590c3621461064657600080fd5b8063b7dd34831461056d578063bbce23861461058d578063d0b9d032146105ad57600080fd5b8063a701da69116100bd578063a701da69146104cb578063b326258d14610513578063b4c4a4c81461055a57600080fd5b806382cac89c146104485780638ce7442614610468578063a01cfffb1461048757600080fd5b80633cc31443116101455780635dfe12ac1161011f5780635dfe12ac146103a3578063613a28d1146103e957806364ae3c9d1461042e57600080fd5b80633cc31443146103335780633dfa1f411461033c5780635c70b7c11461035c57600080fd5b806319caf46c1161017657806319caf46c146102b2578063204f83f91461030b5780632e25d2a61461031357600080fd5b8063012b264a1461019d5780630aa93b9b146101e75780631779467314610215575b600080fd5b6005546101bd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6102076101f53660046106db565b60016020526000908152604090205481565b6040519081526020016101de565b6102a26102233660046106fd565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff93841681526020808201938452948416600090815260029095529320925183547fffffffffffffffffffffffff000000000000000000000000000000000000000016921691909117825551600190910155600c546301000000900460ff1690565b60405190151581526020016101de565b6102076102c03660046106db565b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055600b5490565b600a54610207565b6006546101bd9073ffffffffffffffffffffffffffffffffffffffff1681565b61020760085481565b61020761034a3660046106db565b60006020819052908152604090205481565b6103a161036a366004610739565b600c8054911515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055565b005b6103a16103b1366004610739565b600c805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055565b6102a26103f736600461075b565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260016020526040902055600c5462010000900460ff1690565b6102a261043c366004610785565b600855600c5460ff1690565b6007546101bd9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104759060ff1681565b60405160ff90911681526020016101de565b6102a261049536600461075b565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260208190526040902055600c54610100900460ff1690565b6103a16104d9366004610739565b600c8054911515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909216919091179055565b6102a261052136600461075b565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902055600c54640100000000900460ff1690565b6103a1610568366004610785565b600a55565b6004546101bd9073ffffffffffffffffffffffffffffffffffffffff1681565b61020761059b3660046106db565b60036020526000908152604090205481565b6103a16105bb366004610739565b600c80549115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff909216919091179055565b6103a1610602366004610785565b600b55565b6103a1610615366004610739565b600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b6106866106543660046106db565b6002602052600090815260409020805460019091015473ffffffffffffffffffffffffffffffffffffffff9091169082565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101de565b803573ffffffffffffffffffffffffffffffffffffffff811681146106d657600080fd5b919050565b6000602082840312156106ed57600080fd5b6106f6826106b2565b9392505050565b60008060006060848603121561071257600080fd5b61071b846106b2565b9250610729602085016106b2565b9150604084013590509250925092565b60006020828403121561074b57600080fd5b813580151581146106f657600080fd5b6000806040838503121561076e57600080fd5b610777836106b2565b946020939093013593505050565b60006020828403121561079757600080fd5b503591905056fea264697066735822122023d48e8ff42392ae9b9c81598a90f68d142f4b6ce71155c3400fcfeb94b6467264736f6c634300080d0033a2646970667358221220a6381d37ef0d6232ef4c7dff2373c17b31791019389453994ff859df12547c3564736f6c634300080d0033",
}

// CreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatorMetaData.ABI instead.
var CreatorABI = CreatorMetaData.ABI

// CreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreatorMetaData.Bin instead.
var CreatorBin = CreatorMetaData.Bin

// DeployCreator deploys a new Ethereum contract, binding an instance of Creator to it.
func DeployCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Creator, error) {
	parsed, err := CreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// Creator is an auto generated Go binding around an Ethereum contract.
type Creator struct {
	CreatorCaller     // Read-only binding to the contract
	CreatorTransactor // Write-only binding to the contract
	CreatorFilterer   // Log filterer for contract events
}

// CreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatorSession struct {
	Contract     *Creator          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatorCallerSession struct {
	Contract *CreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatorTransactorSession struct {
	Contract     *CreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatorRaw struct {
	Contract *Creator // Generic contract binding to access the raw methods on
}

// CreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatorCallerRaw struct {
	Contract *CreatorCaller // Generic read-only contract binding to access the raw methods on
}

// CreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatorTransactorRaw struct {
	Contract *CreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreator creates a new instance of Creator, bound to a specific deployed contract.
func NewCreator(address common.Address, backend bind.ContractBackend) (*Creator, error) {
	contract, err := bindCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Creator{CreatorCaller: CreatorCaller{contract: contract}, CreatorTransactor: CreatorTransactor{contract: contract}, CreatorFilterer: CreatorFilterer{contract: contract}}, nil
}

// NewCreatorCaller creates a new read-only instance of Creator, bound to a specific deployed contract.
func NewCreatorCaller(address common.Address, caller bind.ContractCaller) (*CreatorCaller, error) {
	contract, err := bindCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorCaller{contract: contract}, nil
}

// NewCreatorTransactor creates a new write-only instance of Creator, bound to a specific deployed contract.
func NewCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatorTransactor, error) {
	contract, err := bindCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatorTransactor{contract: contract}, nil
}

// NewCreatorFilterer creates a new log filterer instance of Creator, bound to a specific deployed contract.
func NewCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatorFilterer, error) {
	contract, err := bindCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatorFilterer{contract: contract}, nil
}

// bindCreator binds a generic wrapper to an already deployed contract.
func bindCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.CreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.CreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creator *CreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creator *CreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creator *CreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Creator *CreatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Creator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Creator *CreatorSession) Admin() (common.Address, error) {
	return _Creator.Contract.Admin(&_Creator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Creator *CreatorCallerSession) Admin() (common.Address, error) {
	return _Creator.Contract.Admin(&_Creator.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Creator *CreatorCaller) MarketPlace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Creator.contract.Call(opts, &out, "marketPlace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Creator *CreatorSession) MarketPlace() (common.Address, error) {
	return _Creator.Contract.MarketPlace(&_Creator.CallOpts)
}

// MarketPlace is a free data retrieval call binding the contract method 0x2e25d2a6.
//
// Solidity: function marketPlace() view returns(address)
func (_Creator *CreatorCallerSession) MarketPlace() (common.Address, error) {
	return _Creator.Contract.MarketPlace(&_Creator.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorTransactor) Create(opts *bind.TransactOpts, p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "create", p, u, m, c, sw, n, s, d)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorSession) Create(p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.Contract.Create(&_Creator.TransactOpts, p, u, m, c, sw, n, s, d)
}

// Create is a paid mutator transaction binding the contract method 0x6363e867.
//
// Solidity: function create(uint8 p, address u, uint256 m, address c, address sw, string n, string s, uint8 d) returns(address, address)
func (_Creator *CreatorTransactorSession) Create(p uint8, u common.Address, m *big.Int, c common.Address, sw common.Address, n string, s string, d uint8) (*types.Transaction, error) {
	return _Creator.Contract.Create(&_Creator.TransactOpts, p, u, m, c, sw, n, s, d)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Creator *CreatorTransactor) SetAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "setAdmin", a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Creator *CreatorSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _Creator.Contract.SetAdmin(&_Creator.TransactOpts, a)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address a) returns(bool)
func (_Creator *CreatorTransactorSession) SetAdmin(a common.Address) (*types.Transaction, error) {
	return _Creator.Contract.SetAdmin(&_Creator.TransactOpts, a)
}

// SetMarketPlace is a paid mutator transaction binding the contract method 0x30568a8d.
//
// Solidity: function setMarketPlace(address m) returns(bool)
func (_Creator *CreatorTransactor) SetMarketPlace(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _Creator.contract.Transact(opts, "setMarketPlace", m)
}

// SetMarketPlace is a paid mutator transaction binding the contract method 0x30568a8d.
//
// Solidity: function setMarketPlace(address m) returns(bool)
func (_Creator *CreatorSession) SetMarketPlace(m common.Address) (*types.Transaction, error) {
	return _Creator.Contract.SetMarketPlace(&_Creator.TransactOpts, m)
}

// SetMarketPlace is a paid mutator transaction binding the contract method 0x30568a8d.
//
// Solidity: function setMarketPlace(address m) returns(bool)
func (_Creator *CreatorTransactorSession) SetMarketPlace(m common.Address) (*types.Transaction, error) {
	return _Creator.Contract.SetMarketPlace(&_Creator.TransactOpts, m)
}
