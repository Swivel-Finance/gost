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
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055611b95806100326000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80636363e867116100505780636363e867146100d9578063704b6c0214610119578063f851a4401461012c57600080fd5b80632e25d2a61461006c57806330568a8d146100b6575b600080fd5b60015461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100c96100c4366004610533565b61014c565b60405190151581526020016100ad565b6100ec6100e73660046105af565b6102ac565b6040805173ffffffffffffffffffffffffffffffffffffffff9384168152929091166020830152016100ad565b6100c9610127366004610533565b61041c565b60005461008c9073ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff163381146101d7576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a4015b60405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1615610262576001546040517f6d4c6c89000000000000000000000000000000000000000000000000000000008152602160048201526000602482018190526044820181905273ffffffffffffffffffffffffffffffffffffffff9092166064820152608481019190915260a4016101ce565b6001805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155915050919050565b600154600090819073ffffffffffffffffffffffffffffffffffffffff16338114610336576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a4016101ce565b60008d8d8d8d600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168d8d8d8d8d604051610370906104ee565b6103839a999897969594939291906106c5565b604051809103906000f08015801561039f573d6000803e3d6000fd5b50905060008e8d8d8d6040516103b4906104fc565b60ff9094168452602084019290925273ffffffffffffffffffffffffffffffffffffffff9081166040840152166060820152608001604051809103906000f080158015610405573d6000803e3d6000fd5b50919f919e50909c50505050505050505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff163381146104a2576040517f6d4c6c8900000000000000000000000000000000000000000000000000000000815260006004820181905260248201819052604482015233606482015273ffffffffffffffffffffffffffffffffffffffff8216608482015260a4016101ce565b6000805473ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556001915050919050565b610b72806200074883390190565b6108a680620012ba83390190565b803573ffffffffffffffffffffffffffffffffffffffff8116811461052e57600080fd5b919050565b60006020828403121561054557600080fd5b61054e8261050a565b9392505050565b803560ff8116811461052e57600080fd5b60008083601f84011261057857600080fd5b50813567ffffffffffffffff81111561059057600080fd5b6020830191508360208285010111156105a857600080fd5b9250929050565b6000806000806000806000806000806101008b8d0312156105cf57600080fd5b6105d88b610555565b99506105e660208c0161050a565b985060408b013597506105fb60608c0161050a565b965061060960808c0161050a565b955060a08b013567ffffffffffffffff8082111561062657600080fd5b6106328e838f01610566565b909750955060c08d013591508082111561064b57600080fd5b506106588d828e01610566565b909450925061066b905060e08c01610555565b90509295989b9194979a5092959850565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600061010060ff8d16835273ffffffffffffffffffffffffffffffffffffffff808d1660208501528b6040850152808b166060850152808a166080850152508060a0840152610717818401888a61067c565b905082810360c084015261072c81868861067c565b91505060ff831660e08301529b9a505050505050505050505056fe60806040523480156200001157600080fd5b5060405162000b7238038062000b7283398101604081905262000034916200027c565b600380546001600160a01b03808a16610100026001600160a81b031990921660ff8c1617919091179091556004879055600580548783166001600160a01b03199182161790915560068054928716929091169190911790558251620000a1906007906020860190620000da565b508151620000b7906008906020850190620000da565b506009805460ff191660ff92909216919091179055506200038a95505050505050565b828054620000e8906200034e565b90600052602060002090601f0160209004810192826200010c576000855562000157565b82601f106200012757805160ff191683800117855562000157565b8280016001018555821562000157579182015b82811115620001575782518255916020019190600101906200013a565b506200016592915062000169565b5090565b5b808211156200016557600081556001016200016a565b805160ff811681146200019257600080fd5b919050565b80516001600160a01b03811681146200019257600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001d757600080fd5b81516001600160401b0380821115620001f457620001f4620001af565b604051601f8301601f19908116603f011681019082821181831017156200021f576200021f620001af565b816040528381526020925086838588010111156200023c57600080fd5b600091505b8382101562000260578582018301518183018401529082019062000241565b83821115620002725760008385830101525b9695505050505050565b600080600080600080600080610100898b0312156200029a57600080fd5b620002a58962000180565b9750620002b560208a0162000197565b965060408901519550620002cc60608a0162000197565b9450620002dc60808a0162000197565b60a08a01519094506001600160401b0380821115620002fa57600080fd5b620003088c838d01620001c5565b945060c08b01519150808211156200031f57600080fd5b506200032e8b828c01620001c5565b9250506200033f60e08a0162000180565b90509295985092959890939650565b600181811c908216806200036357607f821691505b6020821081036200038457634e487b7160e01b600052602260045260246000fd5b50919050565b6107d8806200039a6000396000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c80638ce74426116100cd578063bba0ad3911610081578063e7ba677411610066578063e7ba677414610495578063ee4db570146104ef578063fdfe5f4d1461050f57600080fd5b8063bba0ad3914610409578063e541efa21461042957600080fd5b80639dc29fac116100b25780639dc29fac1461036d578063b4c4a4c8146103b1578063b9bb928c146103c457600080fd5b80638ce744261461035857806395d89b411461036557600080fd5b8063313ce567116101245780636521b96a116101095780636521b96a146102cc57806369e527da146103155780636f307dc31461033557600080fd5b8063313ce5671461026857806340c10f191461028757600080fd5b806306fdde0314610156578063204f83f91461017457806323b872dd146101865780632ba29d3814610223575b600080fd5b61015e610555565b60405161016b91906105f0565b60405180910390f35b6004545b60405190815260200161016b565b61021361019436600461068c565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff93841681526020808201938452948416600090815260029095529320925183547fffffffffffffffffffffffff0000000000000000000000000000000000000000169216919091178255516001909101556009546301000000900460ff1690565b604051901515815260200161016b565b6006546102439073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016b565b6009546102759060ff1681565b60405160ff909116815260200161016b565b6102136102953660046106c8565b73ffffffffffffffffffffffffffffffffffffffff9190911660009081526001602052604090205560095462010000900460ff1690565b6103136102da3660046106f2565b600980549115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff909216919091179055565b005b6005546102439073ffffffffffffffffffffffffffffffffffffffff1681565b600354610100900473ffffffffffffffffffffffffffffffffffffffff16610243565b6003546102759060ff1681565b61015e6105e3565b61021361037b3660046106c8565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260208190526040902055600954610100900460ff1690565b6103136103bf36600461071b565b600455565b6103136103d23660046106f2565b60098054911515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055565b610178610417366004610734565b60006020819052908152604090205481565b610469610437366004610734565b6002602052600090815260409020805460019091015473ffffffffffffffffffffffffffffffffffffffff9091169082565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835260208301919091520161016b565b6103136104a3366004610734565b6003805473ffffffffffffffffffffffffffffffffffffffff909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b6101786104fd366004610734565b60016020526000908152604090205481565b61031361051d3660046106f2565b6009805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055565b600780546105629061074f565b80601f016020809104026020016040519081016040528092919081815260200182805461058e9061074f565b80156105db5780601f106105b0576101008083540402835291602001916105db565b820191906000526020600020905b8154815290600101906020018083116105be57829003601f168201915b505050505081565b600880546105629061074f565b600060208083528351808285015260005b8181101561061d57858101830151858201604001528201610601565b8181111561062f576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461068757600080fd5b919050565b6000806000606084860312156106a157600080fd5b6106aa84610663565b92506106b860208501610663565b9150604084013590509250925092565b600080604083850312156106db57600080fd5b6106e483610663565b946020939093013593505050565b60006020828403121561070457600080fd5b8135801515811461071457600080fd5b9392505050565b60006020828403121561072d57600080fd5b5035919050565b60006020828403121561074657600080fd5b61071482610663565b600181811c9082168061076357607f821691505b60208210810361079c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea2646970667358221220980d8d4b94b7cc21a6c0d4a0f6cf63be4557820bc7da9dc206b501059f54910764736f6c634300080d0033608060405234801561001057600080fd5b506040516108a63803806108a683398101604081905261002f91610098565b6008805460ff191660ff9590951694909417909355600991909155600480546001600160a01b03199081166001600160a01b039384161790915560058054909116919092161790556100ee565b80516001600160a01b038116811461009357600080fd5b919050565b600080600080608085870312156100ae57600080fd5b845160ff811681146100bf57600080fd5b602086015190945092506100d56040860161007c565b91506100e36060860161007c565b905092959194509250565b6107a9806100fd6000396000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c806382cac89c116100e3578063b7dd34831161008c578063d6cb2c0d11610066578063d6cb2c0d146105c9578063da3de9e9146105dc578063e590c3621461061b57600080fd5b8063b7dd348314610542578063bbce238614610562578063d0b9d0321461058257600080fd5b8063a701da69116100bd578063a701da69146104a0578063b326258d146104e8578063b4c4a4c81461052f57600080fd5b806382cac89c1461041d5780638ce744261461043d578063a01cfffb1461045c57600080fd5b80633cc31443116101455780635dfe12ac1161011f5780635dfe12ac14610378578063613a28d1146103be57806364ae3c9d1461040357600080fd5b80633cc31443146103085780633dfa1f41146103115780635c70b7c11461033157600080fd5b80631779467311610176578063177946731461020a57806319caf46c146102a7578063204f83f91461030057600080fd5b8063012b264a146101925780630aa93b9b146101dc575b600080fd5b6005546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101fc6101ea3660046106b0565b60016020526000908152604090205481565b6040519081526020016101d3565b6102976102183660046106d2565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff93841681526020808201938452948416600090815260029095529320925183547fffffffffffffffffffffffff000000000000000000000000000000000000000016921691909117825551600190910155600b546301000000900460ff1690565b60405190151581526020016101d3565b6101fc6102b53660046106b0565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055600a5490565b6009546101fc565b6101fc60075481565b6101fc61031f3660046106b0565b60006020819052908152604090205481565b61037661033f36600461070e565b600b8054911515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909216919091179055565b005b61037661038636600461070e565b600b805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055565b6102976103cc366004610730565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260016020526040902055600b5462010000900460ff1690565b61029761041136600461075a565b600755600b5460ff1690565b6006546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b60085461044a9060ff1681565b60405160ff90911681526020016101d3565b61029761046a366004610730565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260208190526040902055600b54610100900460ff1690565b6103766104ae36600461070e565b600b8054911515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909216919091179055565b6102976104f6366004610730565b73ffffffffffffffffffffffffffffffffffffffff91909116600090815260036020526040902055600b54640100000000900460ff1690565b61037661053d36600461075a565b600955565b6004546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b6101fc6105703660046106b0565b60036020526000908152604090205481565b61037661059036600461070e565b600b80549115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff909216919091179055565b6103766105d736600461075a565b600a55565b6103766105ea36600461070e565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b61065b6106293660046106b0565b6002602052600090815260409020805460019091015473ffffffffffffffffffffffffffffffffffffffff9091169082565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b803573ffffffffffffffffffffffffffffffffffffffff811681146106ab57600080fd5b919050565b6000602082840312156106c257600080fd5b6106cb82610687565b9392505050565b6000806000606084860312156106e757600080fd5b6106f084610687565b92506106fe60208501610687565b9150604084013590509250925092565b60006020828403121561072057600080fd5b813580151581146106cb57600080fd5b6000806040838503121561074357600080fd5b61074c83610687565b946020939093013593505050565b60006020828403121561076c57600080fd5b503591905056fea2646970667358221220ca290a90e657b0b5f18abc7fdc3160983ccef42276427b67d97b2a13e8e204df64736f6c634300080d0033a264697066735822122071c18b5ed196201506ef294176d31fe6eaff05a68f49a3be52371fb1e7ab9e5364736f6c634300080d0033",
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
