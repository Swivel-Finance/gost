// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// ZcTokenMetaData contains all meta data concerning the ZcToken contract.
var ZcTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_protocol\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_redeemer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"approved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approvals\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Invalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Maturity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"convertToPrincipal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principalAmount\",\"type\":\"uint256\"}],\"name\":\"convertToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principalAmount\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocol\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b506040516200257238038062002572833981016040819052620000359162000233565b828282600062000046848262000394565b50600162000055838262000394565b5060ff81166080524660a0526200006b620000a3565b60c05250505060ff9097166101205250506001600160a01b039384166101005260e092909252821661014052166101605250620004de565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620000d7919062000460565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b805160ff811681146200015157600080fd5b919050565b80516001600160a01b03811681146200015157600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200019657600080fd5b81516001600160401b0380821115620001b357620001b36200016e565b604051601f8301601f19908116603f01168101908282118183101715620001de57620001de6200016e565b81604052838152602092508683858801011115620001fb57600080fd5b600091505b838210156200021f578582018301518183018401529082019062000200565b600093810190920192909252949350505050565b600080600080600080600080610100898b0312156200025157600080fd5b6200025c896200013f565b97506200026c60208a0162000156565b9650604089015195506200028360608a0162000156565b94506200029360808a0162000156565b60a08a01519094506001600160401b0380821115620002b157600080fd5b620002bf8c838d0162000184565b945060c08b0151915080821115620002d657600080fd5b50620002e58b828c0162000184565b925050620002f660e08a016200013f565b90509295985092959890939650565b600181811c908216806200031a57607f821691505b6020821081036200033b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200038f57600081815260208120601f850160051c810160208610156200036a5750805b601f850160051c820191505b818110156200038b5782815560010162000376565b5050505b505050565b81516001600160401b03811115620003b057620003b06200016e565b620003c881620003c1845462000305565b8462000341565b602080601f831160018114620004005760008415620003e75750858301515b600019600386901b1c1916600185901b1785556200038b565b600085815260208120601f198616915b82811015620004315788860151825594840194600190910190840162000410565b5085821015620004505787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000808354620004708162000305565b600182811680156200048b5760018114620004a157620004d2565b60ff1984168752821515830287019450620004d2565b8760005260208060002060005b85811015620004c95781548a820152908401908201620004ae565b50505082870194505b50929695505050505050565b60805160a05160c05160e05161010051610120516101405161016051611f166200065c6000396000818161027b01528181610678015281816107ef01528181610a0e01528181610b3a01528181610dc001528181610fce015281816111a9015281816113b801526115180152600061031b0152600081816103a9015281816105e70152818161075e01528181610d1f01528181610f43015281816111080152818161132d0152611487015260008181610342015281816106250152818161079c01528181610d5d01528181610f6c015281816111460152818161135601526114c5015260008181610241015281816105910152818161064d01528181610708015281816107c401528181610a9b01528181610ae901528181610c5301528181610ca101528181610d8501528181610f9401528181611049015281816110970152818161116e0152818161137e01528181611431015281816114ed01526118a4015260006109e8015260006109b3015260006102c70152611f166000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806369e527da116100f9578063a9059cbb11610097578063ce96cb7711610071578063ce96cb771461041f578063d505accf14610432578063d905777e14610447578063dd62ed3e1461045a57600080fd5b8063a9059cbb146103e6578063b460af94146103f9578063ba0876521461040c57600080fd5b80637ecebe00116100d35780637ecebe00146103845780638ce74426146103a457806395d89b41146103cb5780639dc29fac146103d357600080fd5b806369e527da146103165780636f307dc31461033d57806370a082311461036457600080fd5b806323b872dd11610166578063313ce56711610140578063313ce567146102c25780633644e515146102fb57806340c10f19146103035780634cdad5061461022957600080fd5b806323b872dd1461026357806325a8d87d146101ff5780632ba29d381461027657600080fd5b806318160ddd1161019757806318160ddd146102205780631dc7f52114610229578063204f83f91461023c57600080fd5b806306fdde03146101be578063095ea7b3146101dc5780630a28a477146101ff575b600080fd5b6101c6610485565b6040516101d39190611a9c565b60405180910390f35b6101ef6101ea366004611b31565b610513565b60405190151581526020016101d3565b61021261020d366004611b5b565b61058d565b6040519081526020016101d3565b61021260025481565b610212610237366004611b5b565b610704565b6102127f000000000000000000000000000000000000000000000000000000000000000081565b6101ef610271366004611b74565b610869565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102e97f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101d3565b6102126109af565b6101ef610311366004611b31565b610a0a565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b610212610372366004611bb0565b60036020526000908152604090205481565b610212610392366004611bb0565b60056020526000908152604090205481565b6102e97f000000000000000000000000000000000000000000000000000000000000000081565b6101c6610b29565b6101ef6103e1366004611b31565b610b36565b6101ef6103f4366004611b31565b610bca565b610212610407366004611bcb565b610c4f565b61021261041a366004611bcb565b611045565b61021261042d366004611bb0565b61142d565b610445610440366004611c07565b6115ba565b005b610212610455366004611bb0565b6118a0565b610212610468366004611c7a565b600460209081526000928352604080842090915290825290205481565b6000805461049290611cad565b80601f01602080910402602001604051908101604052809291908181526020018280546104be90611cad565b801561050b5780601f106104e05761010080835404028352916020019161050b565b820191906000526020600020905b8154815290600101906020018083116104ee57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061057b9086815260200190565b60405180910390a35060015b92915050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156105bf57506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa1580156106be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e29190611d00565b9092509050806106f28386611d53565b6106fc9190611d90565b949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000042101561073657506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa158015610835573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108599190611d00565b9092509050816106f28286611d53565b73ffffffffffffffffffffffffffffffffffffffff831660009081526004602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146108fd576108cb8382611dcb565b73ffffffffffffffffffffffffffffffffffffffff861660009081526004602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604081208054859290610932908490611dcb565b909155505073ffffffffffffffffffffffffffffffffffffffff808516600081815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061099a9087815260200190565b60405180910390a360019150505b9392505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146109e5576109e06118fb565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610a99576040517fdc84e3a400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000421115610b15576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a90565b610b1f8484611995565b5060019392505050565b6001805461049290611cad565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610bc0576040517fdc84e3a400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610a90565b610b1f8484611a0e565b33600090815260036020526040812080548391908390610beb908490611dcb565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600081815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061057b9086815260200190565b60007f0000000000000000000000000000000000000000000000000000000000000000421015610ccd576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a90565b6000610cd88561058d565b90503373ffffffffffffffffffffffffffffffffffffffff841603610e33576040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f00000000000000000000000000000000000000000000000000000000000000006044830152336064830152858116608483015260a482018390527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af1158015610e09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2d9190611dde565b506106fc565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020908152604080832033845290915290205481811015610ea8576040517f4a1d07160000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610a90565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054610ee4908390611dcb565b73ffffffffffffffffffffffffffffffffffffffff85811660008181526004602081815260408084203385529091529182902094909455517f52bc943000000000000000000000000000000000000000000000000000000000815260ff7f000000000000000000000000000000000000000000000000000000000000000016938101939093527f0000000000000000000000000000000000000000000000000000000000000000821660248401527f000000000000000000000000000000000000000000000000000000000000000060448401526064830152868116608483015260a482018490527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af1158015611017573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103b9190611dde565b5050949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156110c3576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a90565b3373ffffffffffffffffffffffffffffffffffffffff83160361121d576040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f00000000000000000000000000000000000000000000000000000000000000006044830152336064830152848116608483015260a482018690527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af11580156111f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112169190611dde565b90506109a8565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260046020908152604080832033845290915290205484811015611292576040517f4a1d07160000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610a90565b73ffffffffffffffffffffffffffffffffffffffff831660009081526004602090815260408083203384529091529020546112ce908690611dcb565b73ffffffffffffffffffffffffffffffffffffffff84811660008181526004602081815260408084203385529091529182902094909455517f52bc943000000000000000000000000000000000000000000000000000000000815260ff7f000000000000000000000000000000000000000000000000000000000000000016938101939093527f0000000000000000000000000000000000000000000000000000000000000000821660248401527f000000000000000000000000000000000000000000000000000000000000000060448401526064830152858116608483015260a482018790527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af1158015611401573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114259190611dde565b9150506109a8565b60007f000000000000000000000000000000000000000000000000000000000000000042101561145f57506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa15801561155e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115829190611d00565b73ffffffffffffffffffffffffffffffffffffffff8616600090815260036020526040902054919350915082906106f2908390611d53565b428410156115fd576040517fb979466100000000000000000000000000000000000000000000000000000000815260048101859052426024820152604401610a90565b600060016116096109af565b73ffffffffffffffffffffffffffffffffffffffff8a811660008181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e0830190915280519201919091207f190100000000000000000000000000000000000000000000000000000000000061010083015261010282019290925261012281019190915261014201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa15801561175b573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615806117d557508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b1561182a576040517feecd906e00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff89166024820152604401610a90565b73ffffffffffffffffffffffffffffffffffffffff90811660009081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156118d257506000919050565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f600060405161192d9190611df7565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b80600260008282546119a79190611ecd565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054839290611a43908490611dcb565b909155505060028054829003905560405181815260009073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001611a02565b600060208083528351808285015260005b81811015611ac957858101830151858201604001528201611aad565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611b2c57600080fd5b919050565b60008060408385031215611b4457600080fd5b611b4d83611b08565b946020939093013593505050565b600060208284031215611b6d57600080fd5b5035919050565b600080600060608486031215611b8957600080fd5b611b9284611b08565b9250611ba060208501611b08565b9150604084013590509250925092565b600060208284031215611bc257600080fd5b6109a882611b08565b600080600060608486031215611be057600080fd5b83359250611bf060208501611b08565b9150611bfe60408501611b08565b90509250925092565b600080600080600080600060e0888a031215611c2257600080fd5b611c2b88611b08565b9650611c3960208901611b08565b95506040880135945060608801359350608088013560ff81168114611c5d57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611c8d57600080fd5b611c9683611b08565b9150611ca460208401611b08565b90509250929050565b600181811c90821680611cc157607f821691505b602082108103611cfa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008060408385031215611d1357600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d8b57611d8b611d24565b500290565b600082611dc6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8181038181111561058757610587611d24565b600060208284031215611df057600080fd5b5051919050565b600080835481600182811c915080831680611e1357607f831692505b60208084108203611e4b577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b818015611e5f5760018114611e9257611ebf565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0086168952841515850289019650611ebf565b60008a81526020902060005b86811015611eb75781548b820152908501908301611e9e565b505084890196505b509498975050505050505050565b8082018082111561058757610587611d2456fea26469706673582212203d3f7e3edf545a87c5cff7d74ab37b50750166e26991bdee5ad266be6e02fabc64736f6c63430008100033",
}

// ZcTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ZcTokenMetaData.ABI instead.
var ZcTokenABI = ZcTokenMetaData.ABI

// ZcTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZcTokenMetaData.Bin instead.
var ZcTokenBin = ZcTokenMetaData.Bin

// DeployZcToken deploys a new Ethereum contract, binding an instance of ZcToken to it.
func DeployZcToken(auth *bind.TransactOpts, backend bind.ContractBackend, _protocol uint8, _underlying common.Address, _maturity *big.Int, _cToken common.Address, _redeemer common.Address, _name string, _symbol string, _decimals uint8) (common.Address, *types.Transaction, *ZcToken, error) {
	parsed, err := ZcTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZcTokenBin), backend, _protocol, _underlying, _maturity, _cToken, _redeemer, _name, _symbol, _decimals)
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZcToken *ZcTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZcToken *ZcTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ZcToken.Contract.DOMAINSEPARATOR(&_ZcToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ZcToken *ZcTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ZcToken.Contract.DOMAINSEPARATOR(&_ZcToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.Allowance(&_ZcToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ZcToken *ZcTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ZcToken *ZcTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ZcToken.Contract.BalanceOf(&_ZcToken.CallOpts, arg0)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenSession) CToken() (common.Address, error) {
	return _ZcToken.Contract.CToken(&_ZcToken.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_ZcToken *ZcTokenCallerSession) CToken() (common.Address, error) {
	return _ZcToken.Contract.CToken(&_ZcToken.CallOpts)
}

// ConvertToPrincipal is a free data retrieval call binding the contract method 0x25a8d87d.
//
// Solidity: function convertToPrincipal(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenCaller) ConvertToPrincipal(opts *bind.CallOpts, underlyingAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "convertToPrincipal", underlyingAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToPrincipal is a free data retrieval call binding the contract method 0x25a8d87d.
//
// Solidity: function convertToPrincipal(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenSession) ConvertToPrincipal(underlyingAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.ConvertToPrincipal(&_ZcToken.CallOpts, underlyingAmount)
}

// ConvertToPrincipal is a free data retrieval call binding the contract method 0x25a8d87d.
//
// Solidity: function convertToPrincipal(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) ConvertToPrincipal(underlyingAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.ConvertToPrincipal(&_ZcToken.CallOpts, underlyingAmount)
}

// ConvertToUnderlying is a free data retrieval call binding the contract method 0x1dc7f521.
//
// Solidity: function convertToUnderlying(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenCaller) ConvertToUnderlying(opts *bind.CallOpts, principalAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "convertToUnderlying", principalAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToUnderlying is a free data retrieval call binding the contract method 0x1dc7f521.
//
// Solidity: function convertToUnderlying(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenSession) ConvertToUnderlying(principalAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.ConvertToUnderlying(&_ZcToken.CallOpts, principalAmount)
}

// ConvertToUnderlying is a free data retrieval call binding the contract method 0x1dc7f521.
//
// Solidity: function convertToUnderlying(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) ConvertToUnderlying(principalAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.ConvertToUnderlying(&_ZcToken.CallOpts, principalAmount)
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

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ZcToken *ZcTokenCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ZcToken *ZcTokenSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MaxRedeem(&_ZcToken.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MaxRedeem(&_ZcToken.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ZcToken *ZcTokenCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ZcToken *ZcTokenSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MaxWithdraw(&_ZcToken.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ZcToken.Contract.MaxWithdraw(&_ZcToken.CallOpts, owner)
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

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenCaller) PreviewRedeem(opts *bind.CallOpts, principalAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "previewRedeem", principalAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenSession) PreviewRedeem(principalAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.PreviewRedeem(&_ZcToken.CallOpts, principalAmount)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 principalAmount) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) PreviewRedeem(principalAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.PreviewRedeem(&_ZcToken.CallOpts, principalAmount)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenCaller) PreviewWithdraw(opts *bind.CallOpts, underlyingAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "previewWithdraw", underlyingAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenSession) PreviewWithdraw(underlyingAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.PreviewWithdraw(&_ZcToken.CallOpts, underlyingAmount)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 underlyingAmount) view returns(uint256)
func (_ZcToken *ZcTokenCallerSession) PreviewWithdraw(underlyingAmount *big.Int) (*big.Int, error) {
	return _ZcToken.Contract.PreviewWithdraw(&_ZcToken.CallOpts, underlyingAmount)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenCaller) Protocol(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "protocol")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenSession) Protocol() (uint8, error) {
	return _ZcToken.Contract.Protocol(&_ZcToken.CallOpts)
}

// Protocol is a free data retrieval call binding the contract method 0x8ce74426.
//
// Solidity: function protocol() view returns(uint8)
func (_ZcToken *ZcTokenCallerSession) Protocol() (uint8, error) {
	return _ZcToken.Contract.Protocol(&_ZcToken.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenCaller) Redeemer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZcToken.contract.Call(opts, &out, "redeemer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenSession) Redeemer() (common.Address, error) {
	return _ZcToken.Contract.Redeemer(&_ZcToken.CallOpts)
}

// Redeemer is a free data retrieval call binding the contract method 0x2ba29d38.
//
// Solidity: function redeemer() view returns(address)
func (_ZcToken *ZcTokenCallerSession) Redeemer() (common.Address, error) {
	return _ZcToken.Contract.Redeemer(&_ZcToken.CallOpts)
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
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.Contract.Permit(&_ZcToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ZcToken *ZcTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ZcToken.Contract.Permit(&_ZcToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 principalAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenTransactor) Redeem(opts *bind.TransactOpts, principalAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "redeem", principalAmount, receiver, holder)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 principalAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenSession) Redeem(principalAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.Redeem(&_ZcToken.TransactOpts, principalAmount, receiver, holder)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 principalAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenTransactorSession) Redeem(principalAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.Redeem(&_ZcToken.TransactOpts, principalAmount, receiver, holder)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.Transfer(&_ZcToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ZcToken *ZcTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZcToken.Contract.TransferFrom(&_ZcToken.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 underlyingAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenTransactor) Withdraw(opts *bind.TransactOpts, underlyingAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.contract.Transact(opts, "withdraw", underlyingAmount, receiver, holder)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 underlyingAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenSession) Withdraw(underlyingAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.Withdraw(&_ZcToken.TransactOpts, underlyingAmount, receiver, holder)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 underlyingAmount, address receiver, address holder) returns(uint256)
func (_ZcToken *ZcTokenTransactorSession) Withdraw(underlyingAmount *big.Int, receiver common.Address, holder common.Address) (*types.Transaction, error) {
	return _ZcToken.Contract.Withdraw(&_ZcToken.TransactOpts, underlyingAmount, receiver, holder)
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
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ZcToken *ZcTokenFilterer) ParseApproval(log types.Log) (*ZcTokenApproval, error) {
	event := new(ZcTokenApproval)
	if err := _ZcToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZcTokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the ZcToken contract.
type ZcTokenRedeemIterator struct {
	Event *ZcTokenRedeem // Event containing the contract specifics and raw log

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
func (it *ZcTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZcTokenRedeem)
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
		it.Event = new(ZcTokenRedeem)
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
func (it *ZcTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZcTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZcTokenRedeem represents a Redeem event raised by the ZcToken contract.
type ZcTokenRedeem struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_ZcToken *ZcTokenFilterer) FilterRedeem(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZcTokenRedeemIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.FilterLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZcTokenRedeemIterator{contract: _ZcToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_ZcToken *ZcTokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *ZcTokenRedeem, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZcToken.contract.WatchLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZcTokenRedeem)
				if err := _ZcToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_ZcToken *ZcTokenFilterer) ParseRedeem(log types.Log) (*ZcTokenRedeem, error) {
	event := new(ZcTokenRedeem)
	if err := _ZcToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ZcToken *ZcTokenFilterer) ParseTransfer(log types.Log) (*ZcTokenTransfer, error) {
	event := new(ZcTokenTransfer)
	if err := _ZcToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
