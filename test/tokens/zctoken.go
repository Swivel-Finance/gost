// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// ZcTokenABI is the input ABI used to generate the binding from.
const ZcTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"s\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
var ZcTokenBin = "0x6101006040523480156200001257600080fd5b50604051620027ca380380620027ca8339818101604052810190620000389190620002fb565b81818181816003908051906020019062000054929190620001ab565b5080600490805190602001906200006d929190620001ab565b505050620000be826040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525046306200014760201b62000c8b1760201c565b6080818152505050503373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508373ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff1660601b815250508260e08181525050505050506200057b565b60008085516020870120855160208701206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015286606082015285608082015260a08120935050505080915050949350505050565b828054620001b9906200046c565b90600052602060002090601f016020900481019282620001dd576000855562000229565b82601f10620001f857805160ff191683800117855562000229565b8280016001018555821562000229579182015b82811115620002285782518255916020019190600101906200020b565b5b5090506200023891906200023c565b5090565b5b80821115620002575760008160009055506001016200023d565b5090565b6000620002726200026c84620003c2565b62000399565b9050828152602081018484840111156200028b57600080fd5b6200029884828562000436565b509392505050565b600081519050620002b18162000547565b92915050565b600082601f830112620002c957600080fd5b8151620002db8482602086016200025b565b91505092915050565b600081519050620002f58162000561565b92915050565b600080600080608085870312156200031257600080fd5b60006200032287828801620002a0565b94505060206200033587828801620002e4565b935050604085015167ffffffffffffffff8111156200035357600080fd5b6200036187828801620002b7565b925050606085015167ffffffffffffffff8111156200037f57600080fd5b6200038d87828801620002b7565b91505092959194509250565b6000620003a5620003b8565b9050620003b38282620004a2565b919050565b6000604051905090565b600067ffffffffffffffff821115620003e057620003df62000507565b5b620003eb8262000536565b9050602081019050919050565b600062000405826200040c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b838110156200045657808201518184015260208101905062000439565b8381111562000466576000848401525b50505050565b600060028204905060018216806200048557607f821691505b602082108114156200049c576200049b620004d8565b5b50919050565b620004ad8262000536565b810181811067ffffffffffffffff82111715620004cf57620004ce62000507565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6200055281620003f8565b81146200055e57600080fd5b50565b6200056c816200042c565b81146200057857600080fd5b50565b60805160a05160601c60c05160601c60e0516121fa620005d060003960006104bf0152600061074701526000818161067f0152818161085b0152610c690152600081816107230152610aa801526121fa6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80636f307dc3116100ad578063a457c2d711610071578063a457c2d714610348578063a9059cbb14610378578063d505accf146103a8578063dd62ed3e146103c4578063f851a440146103f457610121565b80636f307dc31461027c57806370a082311461029a5780637ecebe00146102ca57806395d89b41146102fa5780639dc29fac1461031857610121565b806323b872dd116100f457806323b872dd146101b0578063313ce567146101e057806339509351146101fe57806340c10f191461022e57806352a9674b1461025e57610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd14610174578063204f83f914610192575b600080fd5b61012e610412565b60405161013b9190611a09565b60405180910390f35b61015e6004803603810190610159919061168b565b6104a0565b60405161016b919061192d565b60405180910390f35b61017c6104b7565b6040516101899190611bcb565b60405180910390f35b61019a6104bd565b6040516101a79190611bcb565b60405180910390f35b6101ca60048036038101906101c5919061159e565b6104e1565b6040516101d7919061192d565b60405180910390f35b6101e86105d4565b6040516101f59190611be6565b60405180910390f35b6102186004803603810190610213919061168b565b6105dd565b604051610225919061192d565b60405180910390f35b6102486004803603810190610243919061168b565b61067b565b604051610255919061192d565b60405180910390f35b610266610721565b6040516102739190611948565b60405180910390f35b610284610745565b6040516102919190611912565b60405180910390f35b6102b460048036038101906102af9190611539565b610769565b6040516102c19190611bcb565b60405180910390f35b6102e460048036038101906102df9190611539565b6107b1565b6040516102f19190611bcb565b60405180910390f35b6103026107c9565b60405161030f9190611a09565b60405180910390f35b610332600480360381019061032d919061168b565b610857565b60405161033f919061192d565b60405180910390f35b610362600480360381019061035d919061168b565b6108fd565b60405161036f919061192d565b60405180910390f35b610392600480360381019061038d919061168b565b6109e3565b60405161039f919061192d565b60405180910390f35b6103c260048036038101906103bd91906115ed565b6109fa565b005b6103de60048036038101906103d99190611562565b610be0565b6040516103eb9190611bcb565b60405180910390f35b6103fc610c67565b6040516104099190611912565b60405180910390f35b6003805461041f90611d39565b80601f016020809104026020016040519081016040528092919081815260200182805461044b90611d39565b80156104985780601f1061046d57610100808354040283529160200191610498565b820191906000526020600020905b81548152906001019060200180831161047b57829003601f168201915b505050505081565b60006104ad338484610cef565b6001905092915050565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006104ee848484610eba565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156105b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a990611b6b565b60405180910390fd5b6105c8853385846105c39190611c73565b610cef565b60019150509392505050565b60006012905090565b6000610671338484600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461066c9190611c1d565b610cef565b6001905092915050565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461070c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070390611a2b565b60405180910390fd5b610716848461112e565b600191505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60056020528060005260406000206000915090505481565b600480546107d690611d39565b80601f016020809104026020016040519081016040528092919081815260200182805461080290611d39565b801561084f5780601f106108245761010080835404028352916020019161084f565b820191906000526020600020905b81548152906001019060200180831161083257829003601f168201915b505050505081565b60007f00000000000000000000000000000000000000000000000000000000000000008073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108df90611a2b565b60405180910390fd5b6108f28484611276565b600191505092915050565b600080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156109c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b990611bab565b60405180910390fd5b6109d8338585846109d39190611c73565b610cef565b600191505092915050565b60006109f0338484610eba565b6001905092915050565b42841015610a3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3490611b0b565b60405180910390fd5b6000610a9f888888600560008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610a9590611d6b565b919050558961143e565b90506000610acd7f00000000000000000000000000000000000000000000000000000000000000008361149f565b9050600060018287878760405160008152602001604052604051610af494939291906119c4565b6020604051602081039080840390855afa158015610b16573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610b8a57508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610bc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc090611b2b565b60405180910390fd5b610bd48a8a8a610cef565b50505050505050505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008085516020870120855160208701206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015286606082015285608082015260a08120935050505080915050949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610d5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5690611b8b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc690611a8b565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610ead9190611bcb565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2190611acb565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9190611a4b565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611020576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101790611aeb565b60405180910390fd5b818161102c9190611c73565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110bc9190611c1d565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111209190611bcb565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561119e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119590611aab565b60405180910390fd5b80600260008282546111b09190611c1d565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112059190611c1d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161126a9190611bcb565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156112e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112dd90611b4b565b60405180910390fd5b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561136c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136390611a6b565b60405180910390fd5b81816113789190611c73565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546113cc9190611c73565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114319190611bcb565b60405180910390a3505050565b60007f80772249b4aef1688b30651778f4249b05cb73b517d98482439b9d8999b3060260001b868686868660405160200161147e96959493929190611963565b60405160208183030381529060405280519060200120905095945050505050565b6000806040517f19010000000000000000000000000000000000000000000000000000000000008152846002820152836022820152604281209150508091505092915050565b6000813590506114f481612168565b92915050565b6000813590506115098161217f565b92915050565b60008135905061151e81612196565b92915050565b600081359050611533816121ad565b92915050565b60006020828403121561154b57600080fd5b6000611559848285016114e5565b91505092915050565b6000806040838503121561157557600080fd5b6000611583858286016114e5565b9250506020611594858286016114e5565b9150509250929050565b6000806000606084860312156115b357600080fd5b60006115c1868287016114e5565b93505060206115d2868287016114e5565b92505060406115e38682870161150f565b9150509250925092565b600080600080600080600060e0888a03121561160857600080fd5b60006116168a828b016114e5565b97505060206116278a828b016114e5565b96505060406116388a828b0161150f565b95505060606116498a828b0161150f565b945050608061165a8a828b01611524565b93505060a061166b8a828b016114fa565b92505060c061167c8a828b016114fa565b91505092959891949750929550565b6000806040838503121561169e57600080fd5b60006116ac858286016114e5565b92505060206116bd8582860161150f565b9150509250929050565b6116d081611ca7565b82525050565b6116df81611cb9565b82525050565b6116ee81611cc5565b82525050565b60006116ff82611c01565b6117098185611c0c565b9350611719818560208601611d06565b61172281611e12565b840191505092915050565b600061173a601483611c0c565b915061174582611e23565b602082019050919050565b600061175d602283611c0c565b915061176882611e4c565b604082019050919050565b6000611780602183611c0c565b915061178b82611e9b565b604082019050919050565b60006117a3602183611c0c565b91506117ae82611eea565b604082019050919050565b60006117c6601e83611c0c565b91506117d182611f39565b602082019050919050565b60006117e9602483611c0c565b91506117f482611f62565b604082019050919050565b600061180c602583611c0c565b915061181782611fb1565b604082019050919050565b600061182f601883611c0c565b915061183a82612000565b602082019050919050565b6000611852601983611c0c565b915061185d82612029565b602082019050919050565b6000611875602083611c0c565b915061188082612052565b602082019050919050565b6000611898602783611c0c565b91506118a38261207b565b604082019050919050565b60006118bb602383611c0c565b91506118c6826120ca565b604082019050919050565b60006118de602483611c0c565b91506118e982612119565b604082019050919050565b6118fd81611cef565b82525050565b61190c81611cf9565b82525050565b600060208201905061192760008301846116c7565b92915050565b600060208201905061194260008301846116d6565b92915050565b600060208201905061195d60008301846116e5565b92915050565b600060c08201905061197860008301896116e5565b61198560208301886116c7565b61199260408301876116c7565b61199f60608301866118f4565b6119ac60808301856118f4565b6119b960a08301846118f4565b979650505050505050565b60006080820190506119d960008301876116e5565b6119e66020830186611903565b6119f360408301856116e5565b611a0060608301846116e5565b95945050505050565b60006020820190508181036000830152611a2381846116f4565b905092915050565b60006020820190508181036000830152611a448161172d565b9050919050565b60006020820190508181036000830152611a6481611750565b9050919050565b60006020820190508181036000830152611a8481611773565b9050919050565b60006020820190508181036000830152611aa481611796565b9050919050565b60006020820190508181036000830152611ac4816117b9565b9050919050565b60006020820190508181036000830152611ae4816117dc565b9050919050565b60006020820190508181036000830152611b04816117ff565b9050919050565b60006020820190508181036000830152611b2481611822565b9050919050565b60006020820190508181036000830152611b4481611845565b9050919050565b60006020820190508181036000830152611b6481611868565b9050919050565b60006020820190508181036000830152611b848161188b565b9050919050565b60006020820190508181036000830152611ba4816118ae565b9050919050565b60006020820190508181036000830152611bc4816118d1565b9050919050565b6000602082019050611be060008301846118f4565b92915050565b6000602082019050611bfb6000830184611903565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611c2882611cef565b9150611c3383611cef565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c6857611c67611db4565b5b828201905092915050565b6000611c7e82611cef565b9150611c8983611cef565b925082821015611c9c57611c9b611db4565b5b828203905092915050565b6000611cb282611ccf565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611d24578082015181840152602081019050611d09565b83811115611d33576000848401525b50505050565b60006002820490506001821680611d5157607f821691505b60208210811415611d6557611d64611de3565b5b50919050565b6000611d7682611cef565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611da957611da8611db4565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f73656e646572206d7573742062652061646d696e000000000000000000000000600082015250565b7f6572633230207472616e7366657220746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f6572633230206275726e20616d6f756e7420657863656564732062616c616e6360008201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b7f657263323020617070726f766520746f20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f6572633230206d696e7420746f20746865207a65726f20616464726573730000600082015250565b7f6572633230207472616e736665722066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f6572633230207472616e7366657220616d6f756e74206578636565647320626160008201527f6c616e6365000000000000000000000000000000000000000000000000000000602082015250565b7f65726332363132206578706972656420646561646c696e650000000000000000600082015250565b7f6572633236313220696e76616c6964207369676e617475726500000000000000600082015250565b7f6572633230206275726e2066726f6d20746865207a65726f2061646472657373600082015250565b7f6572633230207472616e7366657220616d6f756e74206578636565647320616c60008201527f6c6f77616e636500000000000000000000000000000000000000000000000000602082015250565b7f657263323020617070726f76652066726f6d20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f65726332302064656372656173656420616c6c6f77616e63652062656c6f772060008201527f7a65726f00000000000000000000000000000000000000000000000000000000602082015250565b61217181611ca7565b811461217c57600080fd5b50565b61218881611cc5565b811461219357600080fd5b50565b61219f81611cef565b81146121aa57600080fd5b50565b6121b681611cf9565b81146121c157600080fd5b5056fea2646970667358221220a015fd8055f023d03763735a46501af80a2474be3f778bd1f3131abb4afaf74364736f6c63430008040033"

// DeployZcToken deploys a new Ethereum contract, binding an instance of ZcToken to it.
func DeployZcToken(auth *bind.TransactOpts, backend bind.ContractBackend, u common.Address, m *big.Int, n string, s string) (common.Address, *types.Transaction, *ZcToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZcTokenBin), backend, u, m, n, s)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// ZcToken is an auto generated Go binding around an Ethereum contract.
type ZcToken struct {
	ZcTokenCaller     // Read-only binding to the contract
	ZcTokenTransactor // Write-only binding to the contract
	ZcTokenFilterer   // Log filterer for contract events
}

// ZcTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZcTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZcTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZcTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZcTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZcTokenSession struct {
	Contract     *ZcToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZcTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZcTokenCallerSession struct {
	Contract *ZcTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZcTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZcTokenTransactorSession struct {
	Contract     *ZcTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZcTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZcTokenRaw struct {
	Contract *ZcToken // Generic contract binding to access the raw methods on
}

// ZcTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZcTokenCallerRaw struct {
	Contract *ZcTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZcTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZcTokenTransactorRaw struct {
	Contract *ZcTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZcToken creates a new instance of ZcToken, bound to a specific deployed contract.
func NewZcToken(address common.Address, backend bind.ContractBackend) (*ZcToken, error) {
	contract, err := bindZcToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZcToken{ZcTokenCaller: ZcTokenCaller{contract: contract}, ZcTokenTransactor: ZcTokenTransactor{contract: contract}, ZcTokenFilterer: ZcTokenFilterer{contract: contract}}, nil
}

// NewZcTokenCaller creates a new read-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenCaller(address common.Address, caller bind.ContractCaller) (*ZcTokenCaller, error) {
	contract, err := bindZcToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenCaller{contract: contract}, nil
}

// NewZcTokenTransactor creates a new write-only instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZcTokenTransactor, error) {
	contract, err := bindZcToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransactor{contract: contract}, nil
}

// NewZcTokenFilterer creates a new log filterer instance of ZcToken, bound to a specific deployed contract.
func NewZcTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZcTokenFilterer, error) {
	contract, err := bindZcToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZcTokenFilterer{contract: contract}, nil
}

// bindZcToken binds a generic wrapper to an already deployed contract.
func bindZcToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZcTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.ZcTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.ZcTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZcToken *ZcTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZcToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZcToken *ZcTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZcToken *ZcTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZcToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_ZcToken *ZcTokenCaller) DOMAIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "DOMAIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_ZcToken *ZcTokenSession) DOMAIN() ([32]byte, error) {
	return _ZcToken.Contract.DOMAIN(&_ZcToken.CallOpts)
}

// DOMAIN is a free data retrieval call binding the contract method 0x52a9674b.
//
// Solidity: function DOMAIN() view returns(bytes32)
func (_ZcToken *ZcTokenCallerSession) DOMAIN() ([32]byte, error) {
	return _ZcToken.Contract.DOMAIN(&_ZcToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenSession) Admin() (common.Address, error) {
	return _ZcToken.Contract.Admin(&_ZcToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Admin() (common.Address, error) {
	return _ZcToken.Contract.Admin(&_ZcToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Allowance(opts *bind.CallOpts, o common.Address, s common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "allowance", o, s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_ZcToken *ZcTokenSession) Allowance(o common.Address, s common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, o, s)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address o, address s) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Allowance(o common.Address, s common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, o, s)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balanceOf", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address a) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, a)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZcToken *ZcTokenCallerSession) Decimals() (uint8, error) {
	return _ZcToken.Contract.Decimals(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCaller) Maturity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "maturity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// Maturity is a free data retrieval call binding the contract method 0x204f83f9.
//
// Solidity: function maturity() view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Maturity() (*big.Int, error) {
	return _ZcToken.Contract.Maturity(&_ZcToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Name() (string, error) {
	return _ZcToken.Contract.Name(&_ZcToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Nonces(&_ZcToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Nonces(&_ZcToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZcToken *ZcTokenCallerSession) Symbol() (string, error) {
	return _ZcToken.Contract.Symbol(&_ZcToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenSession) TotalSupply() (*big.Int, error) {
	return _ZcToken.Contract.TotalSupply(&_ZcToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ZcToken.Contract.TotalSupply(&_ZcToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Underlying() (common.Address, error) {
	return _ZcToken.Contract.Underlying(&_ZcToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Approve(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "approve", s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, s, a)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Approve(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, s, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Burn(opts *bind.TransactOpts, f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "burn", f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address f, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Burn(f common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Burn(&_ZcToken.TransactOpts, f, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "decreaseAllowance", s, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) DecreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, s, a)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) DecreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "increaseAllowance", s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) IncreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, s, a)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address s, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) IncreaseAllowance(s common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, s, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Mint(opts *bind.TransactOpts, t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "mint", t, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Mint(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, t, a)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address t, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Mint(t common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Mint(&_ZcToken.TransactOpts, t, a)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address o, address spender, uint256 a, uint256 d, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenTransactor) Permit(opts *bind.TransactOpts, o common.Address, spender common.Address, a *big.Int, d *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "permit", o, spender, a, d, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address o, address spender, uint256 a, uint256 d, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenSession) Permit(o common.Address, spender common.Address, a *big.Int, d *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.Contract.Permit(&_ZcToken.TransactOpts, o, spender, a, d, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address o, address spender, uint256 a, uint256 d, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenTransactorSession) Permit(o common.Address, spender common.Address, a *big.Int, d *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.Contract.Permit(&_ZcToken.TransactOpts, o, spender, a, d, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) Transfer(opts *bind.TransactOpts, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transfer", r, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) Transfer(r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, r, a)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Transfer(r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactor) TransferFrom(opts *bind.TransactOpts, s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFrom", s, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenSession) TransferFrom(s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, s, r, a)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address s, address r, uint256 a) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) TransferFrom(s common.Address, r common.Address, a *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, s, r, a)
}

// ZcTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZcToken contract.
type ZcTokenApprovalIterator struct {
	Event *ZcTokenApproval // Event containing the contract specifics and raw log

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
func (it *ZcTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZcTokenApproval)
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
		it.Event = new(ZcTokenApproval)
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
func (it *ZcTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZcTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZcTokenApproval represents a Approval event raised by the ZcToken contract.
type ZcTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZcTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZcToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZcTokenApprovalIterator{contract: _ZcToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZcTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZcToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZcTokenApproval)
				if err := _ZcToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZcToken *ZcTokenFilterer) ParseApproval(log types.Log) (*ZcTokenApproval, error) {
	event := new(ZcTokenApproval)
	if err := _ZcToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZcTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZcToken contract.
type ZcTokenTransferIterator struct {
	Event *ZcTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ZcTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZcTokenTransfer)
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
		it.Event = new(ZcTokenTransfer)
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
func (it *ZcTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZcTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZcTokenTransfer represents a Transfer event raised by the ZcToken contract.
type ZcTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZcTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZcTokenTransferIterator{contract: _ZcToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZcTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZcTokenTransfer)
				if err := _ZcToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZcToken *ZcTokenFilterer) ParseTransfer(log types.Log) (*ZcTokenTransfer, error) {
	event := new(ZcTokenTransfer)
	if err := _ZcToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
