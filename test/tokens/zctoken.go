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
const ZcTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"m\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"n\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"o\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
var ZcTokenBin = "0x60a06040523480156200001157600080fd5b50604051620027b7380380620027b7833981810160405281019062000037919062000308565b81818181816003908051906020019062000053929190620001b8565b5080600490805190602001906200006c929190620001b8565b505050620000bd826040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525046306200015960201b62000c811760201c565b60808181525050505083600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260088190555033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000588565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b828054620001c69062000479565b90600052602060002090601f016020900481019282620001ea576000855562000236565b82601f106200020557805160ff191683800117855562000236565b8280016001018555821562000236579182015b828111156200023557825182559160200191906001019062000218565b5b50905062000245919062000249565b5090565b5b80821115620002645760008160009055506001016200024a565b5090565b60006200027f6200027984620003cf565b620003a6565b9050828152602081018484840111156200029857600080fd5b620002a584828562000443565b509392505050565b600081519050620002be8162000554565b92915050565b600082601f830112620002d657600080fd5b8151620002e884826020860162000268565b91505092915050565b60008151905062000302816200056e565b92915050565b600080600080608085870312156200031f57600080fd5b60006200032f87828801620002ad565b94505060206200034287828801620002f1565b935050604085015167ffffffffffffffff8111156200036057600080fd5b6200036e87828801620002c4565b925050606085015167ffffffffffffffff8111156200038c57600080fd5b6200039a87828801620002c4565b91505092959194509250565b6000620003b2620003c5565b9050620003c08282620004af565b919050565b6000604051905090565b600067ffffffffffffffff821115620003ed57620003ec62000514565b5b620003f88262000543565b9050602081019050919050565b6000620004128262000419565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b838110156200046357808201518184015260208101905062000446565b8381111562000473576000848401525b50505050565b600060028204905060018216806200049257607f821691505b60208210811415620004a957620004a8620004e5565b5b50919050565b620004ba8262000543565b810181811067ffffffffffffffff82111715620004dc57620004db62000514565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6200055f8162000405565b81146200056b57600080fd5b50565b620005798162000439565b81146200058557600080fd5b50565b60805161220c620005ab6000396000818161070f0152610a9c015261220c6000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80636f307dc3116100ad578063a457c2d711610071578063a457c2d714610348578063a9059cbb14610378578063d505accf146103a8578063dd62ed3e146103c4578063f851a440146103f457610121565b80636f307dc31461027c57806370a082311461029a5780637ecebe00146102ca57806395d89b41146102fa5780639dc29fac1461031857610121565b806323b872dd116100f457806323b872dd146101b0578063313ce567146101e057806339509351146101fe57806340c10f191461022e57806352a9674b1461025e57610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd14610174578063204f83f914610192575b600080fd5b61012e610412565b60405161013b91906119f5565b60405180910390f35b61015e60048036038101906101599190611677565b6104a4565b60405161016b9190611919565b60405180910390f35b61017c6104bb565b6040516101899190611bb7565b60405180910390f35b61019a6104c5565b6040516101a79190611bb7565b60405180910390f35b6101ca60048036038101906101c5919061158a565b6104cb565b6040516101d79190611919565b60405180910390f35b6101e86105be565b6040516101f59190611bd2565b60405180910390f35b61021860048036038101906102139190611677565b6105c7565b6040516102259190611919565b60405180910390f35b61024860048036038101906102439190611677565b610665565b6040516102559190611919565b60405180910390f35b61026661070d565b6040516102739190611934565b60405180910390f35b610284610731565b60405161029191906118fe565b60405180910390f35b6102b460048036038101906102af9190611525565b610757565b6040516102c19190611bb7565b60405180910390f35b6102e460048036038101906102df9190611525565b61079f565b6040516102f19190611bb7565b60405180910390f35b6103026107b7565b60405161030f91906119f5565b60405180910390f35b610332600480360381019061032d9190611677565b610849565b60405161033f9190611919565b60405180910390f35b610362600480360381019061035d9190611677565b6108f1565b60405161036f9190611919565b60405180910390f35b610392600480360381019061038d9190611677565b6109d7565b60405161039f9190611919565b60405180910390f35b6103c260048036038101906103bd91906115d9565b6109ee565b005b6103de60048036038101906103d9919061154e565b610bd4565b6040516103eb9190611bb7565b60405180910390f35b6103fc610c5b565b60405161040991906118fe565b60405180910390f35b60606003805461042190611d25565b80601f016020809104026020016040519081016040528092919081815260200182805461044d90611d25565b801561049a5780601f1061046f5761010080835404028352916020019161049a565b820191906000526020600020905b81548152906001019060200180831161047d57829003601f168201915b5050505050905090565b60006104b1338484610ce0565b6001905092915050565b6000600254905090565b60085481565b60006104d8848484610eab565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561059c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059390611af7565b60405180910390fd5b6105b2853385846105ad9190611c5f565b610ce0565b60019150509392505050565b60006012905090565b600061065b338484600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546106569190611c09565b610ce0565b6001905092915050565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ef90611a57565b60405180910390fd5b610702848461111f565b600191505092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60056020528060005260406000206000915090505481565b6060600480546107c690611d25565b80601f01602080910402602001604051908101604052809291908181526020018280546107f290611d25565b801561083f5780601f106108145761010080835404028352916020019161083f565b820191906000526020600020905b81548152906001019060200180831161082257829003601f168201915b5050505050905090565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d390611a57565b60405180910390fd5b6108e68484611267565b600191505092915050565b600080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156109b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ad90611b77565b60405180910390fd5b6109cc338585846109c79190611c5f565b610ce0565b600191505092915050565b60006109e4338484610eab565b6001905092915050565b42841015610a31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2890611ab7565b60405180910390fd5b6000610a93888888600560008d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190610a8990611d57565b919050558961142f565b90506000610ac17f000000000000000000000000000000000000000000000000000000000000000083611490565b9050600060018287878760405160008152602001604052604051610ae894939291906119b0565b6020604051602081039080840390855afa158015610b0a573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610b7e57508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610bbd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb490611ad7565b60405180910390fd5b610bc88a8a8a610ce0565b50505050505050505050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600084516020860120845160208601206040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f815282602082015281604082015285606082015284608082015260a081209350505050949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610d50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4790611b57565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db790611a77565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610e9e9190611bb7565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610f1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1290611b37565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8290611a17565b60405180910390fd5b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611011576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100890611a97565b60405180910390fd5b818161101d9190611c5f565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110ad9190611c09565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111119190611bb7565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561118f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118690611b97565b60405180910390fd5b80600260008282546111a19190611c09565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546111f69190611c09565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161125b9190611bb7565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156112d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ce90611b17565b60405180910390fd5b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561135d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161135490611a37565b60405180910390fd5b81816113699190611c5f565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546113bd9190611c5f565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114229190611bb7565b60405180910390a3505050565b60007f80772249b4aef1688b30651778f4249b05cb73b517d98482439b9d8999b3060260001b868686868660405160200161146f9695949392919061194f565b60405160208183030381529060405280519060200120905095945050505050565b60006040517f190100000000000000000000000000000000000000000000000000000000000081528360028201528260228201526042812091505092915050565b6000813590506114e08161217a565b92915050565b6000813590506114f581612191565b92915050565b60008135905061150a816121a8565b92915050565b60008135905061151f816121bf565b92915050565b60006020828403121561153757600080fd5b6000611545848285016114d1565b91505092915050565b6000806040838503121561156157600080fd5b600061156f858286016114d1565b9250506020611580858286016114d1565b9150509250929050565b60008060006060848603121561159f57600080fd5b60006115ad868287016114d1565b93505060206115be868287016114d1565b92505060406115cf868287016114fb565b9150509250925092565b600080600080600080600060e0888a0312156115f457600080fd5b60006116028a828b016114d1565b97505060206116138a828b016114d1565b96505060406116248a828b016114fb565b95505060606116358a828b016114fb565b94505060806116468a828b01611510565b93505060a06116578a828b016114e6565b92505060c06116688a828b016114e6565b91505092959891949750929550565b6000806040838503121561168a57600080fd5b6000611698858286016114d1565b92505060206116a9858286016114fb565b9150509250929050565b6116bc81611c93565b82525050565b6116cb81611ca5565b82525050565b6116da81611cb1565b82525050565b60006116eb82611bed565b6116f58185611bf8565b9350611705818560208601611cf2565b61170e81611dfe565b840191505092915050565b6000611726602383611bf8565b915061173182611e0f565b604082019050919050565b6000611749602283611bf8565b915061175482611e5e565b604082019050919050565b600061176c601483611bf8565b915061177782611ead565b602082019050919050565b600061178f602283611bf8565b915061179a82611ed6565b604082019050919050565b60006117b2602683611bf8565b91506117bd82611f25565b604082019050919050565b60006117d5601983611bf8565b91506117e082611f74565b602082019050919050565b60006117f8601a83611bf8565b915061180382611f9d565b602082019050919050565b600061181b602883611bf8565b915061182682611fc6565b604082019050919050565b600061183e602183611bf8565b915061184982612015565b604082019050919050565b6000611861602583611bf8565b915061186c82612064565b604082019050919050565b6000611884602483611bf8565b915061188f826120b3565b604082019050919050565b60006118a7602583611bf8565b91506118b282612102565b604082019050919050565b60006118ca601f83611bf8565b91506118d582612151565b602082019050919050565b6118e981611cdb565b82525050565b6118f881611ce5565b82525050565b600060208201905061191360008301846116b3565b92915050565b600060208201905061192e60008301846116c2565b92915050565b600060208201905061194960008301846116d1565b92915050565b600060c08201905061196460008301896116d1565b61197160208301886116b3565b61197e60408301876116b3565b61198b60608301866118e0565b61199860808301856118e0565b6119a560a08301846118e0565b979650505050505050565b60006080820190506119c560008301876116d1565b6119d260208301866118ef565b6119df60408301856116d1565b6119ec60608301846116d1565b95945050505050565b60006020820190508181036000830152611a0f81846116e0565b905092915050565b60006020820190508181036000830152611a3081611719565b9050919050565b60006020820190508181036000830152611a508161173c565b9050919050565b60006020820190508181036000830152611a708161175f565b9050919050565b60006020820190508181036000830152611a9081611782565b9050919050565b60006020820190508181036000830152611ab0816117a5565b9050919050565b60006020820190508181036000830152611ad0816117c8565b9050919050565b60006020820190508181036000830152611af0816117eb565b9050919050565b60006020820190508181036000830152611b108161180e565b9050919050565b60006020820190508181036000830152611b3081611831565b9050919050565b60006020820190508181036000830152611b5081611854565b9050919050565b60006020820190508181036000830152611b7081611877565b9050919050565b60006020820190508181036000830152611b908161189a565b9050919050565b60006020820190508181036000830152611bb0816118bd565b9050919050565b6000602082019050611bcc60008301846118e0565b92915050565b6000602082019050611be760008301846118ef565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611c1482611cdb565b9150611c1f83611cdb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c5457611c53611da0565b5b828201905092915050565b6000611c6a82611cdb565b9150611c7583611cdb565b925082821015611c8857611c87611da0565b5b828203905092915050565b6000611c9e82611cbb565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611d10578082015181840152602081019050611cf5565b83811115611d1f576000848401525b50505050565b60006002820490506001821680611d3d57607f821691505b60208210811415611d5157611d50611dcf565b5b50919050565b6000611d6282611cdb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d9557611d94611da0565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f73656e646572206d7573742062652061646d696e000000000000000000000000600082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f457263323631323a206578706972656420646561646c696e6500000000000000600082015250565b7f457263323631323a20696e76616c6964207369676e6174757265000000000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b61218381611c93565b811461218e57600080fd5b50565b61219a81611cb1565b81146121a557600080fd5b50565b6121b181611cdb565b81146121bc57600080fd5b50565b6121c881611ce5565b81146121d357600080fd5b5056fea2646970667358221220fe40a0f7bb45bdc274142f21a3dfd2240f326dac5ea6812bda76af5795f44c9f64736f6c63430008040033"

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
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, account)
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Approve(&_ZcToken.TransactOpts, spender, amount)
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
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.DecreaseAllowance(&_ZcToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.IncreaseAllowance(&_ZcToken.TransactOpts, spender, addedValue)
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
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, sender, recipient, amount)
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
