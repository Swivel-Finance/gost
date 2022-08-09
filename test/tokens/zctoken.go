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
	Bin: "0x6101806040523480156200001257600080fd5b5060405162002516380380620025168339810160408190526200003591620002f2565b82828282600090805190602001906200005092919062000150565b5081516200006690600190602085019062000150565b5060ff81166080524660a0526200007c620000b4565b60c05250505060ff9097166101205250506001600160a01b039384166101005260e092909252821661014052166101605250620004a3565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620000e8919062000400565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b8280546200015e90620003c4565b90600052602060002090601f016020900481019282620001825760008555620001cd565b82601f106200019d57805160ff1916838001178555620001cd565b82800160010185558215620001cd579182015b82811115620001cd578251825591602001919060010190620001b0565b50620001db929150620001df565b5090565b5b80821115620001db5760008155600101620001e0565b805160ff811681146200020857600080fd5b919050565b80516001600160a01b03811681146200020857600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200024d57600080fd5b81516001600160401b03808211156200026a576200026a62000225565b604051601f8301601f19908116603f0116810190828211818310171562000295576200029562000225565b81604052838152602092508683858801011115620002b257600080fd5b600091505b83821015620002d65785820183015181830184015290820190620002b7565b83821115620002e85760008385830101525b9695505050505050565b600080600080600080600080610100898b0312156200031057600080fd5b6200031b89620001f6565b97506200032b60208a016200020d565b9650604089015195506200034260608a016200020d565b94506200035260808a016200020d565b60a08a01519094506001600160401b03808211156200037057600080fd5b6200037e8c838d016200023b565b945060c08b01519150808211156200039557600080fd5b50620003a48b828c016200023b565b925050620003b560e08a01620001f6565b90509295985092959890939650565b600181811c90821680620003d957607f821691505b602082108103620003fa57634e487b7160e01b600052602260045260246000fd5b50919050565b600080835481600182811c9150808316806200041d57607f831692505b602080841082036200043d57634e487b7160e01b86526022600452602486fd5b818015620004545760018114620004665762000495565b60ff1986168952848901965062000495565b60008a81526020902060005b868110156200048d5781548b82015290850190830162000472565b505084890196505b509498975050505050505050565b60805160a05160c05160e05161010051610120516101405161016051611ef5620006216000396000818161027b01528181610677015281816107ee01528181610a0d01528181610b3901528181610dbf01528181610fb7015281816111920152818161138b01526114eb0152600061031b0152600081816103a9015281816105e60152818161075d01528181610d1e01528181610f14015281816110f1015281816112e8015261145a015260008181610342015281816106240152818161079b01528181610d5c01528181610f520152818161112f015281816113260152611498015260008181610241015281816105900152818161064c01528181610707015281816107c301528181610a9a01528181610ae801528181610c5201528181610ca001528181610d8401528181610f7a0152818161103201528181611080015281816111570152818161134e01528181611404015281816114c00152611877015260006109e7015260006109b2015260006102c70152611ef56000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806369e527da116100f9578063a9059cbb11610097578063ce96cb7711610071578063ce96cb771461041f578063d505accf14610432578063d905777e14610447578063dd62ed3e1461045a57600080fd5b8063a9059cbb146103e6578063b460af94146103f9578063ba0876521461040c57600080fd5b80637ecebe00116100d35780637ecebe00146103845780638ce74426146103a457806395d89b41146103cb5780639dc29fac146103d357600080fd5b806369e527da146103165780636f307dc31461033d57806370a082311461036457600080fd5b806323b872dd11610166578063313ce56711610140578063313ce567146102c25780633644e515146102fb57806340c10f19146103035780634cdad5061461022957600080fd5b806323b872dd1461026357806325a8d87d146101ff5780632ba29d381461027657600080fd5b806318160ddd1161019757806318160ddd146102205780631dc7f52114610229578063204f83f91461023c57600080fd5b806306fdde03146101be578063095ea7b3146101dc5780630a28a477146101ff575b600080fd5b6101c6610485565b6040516101d39190611a6f565b60405180910390f35b6101ef6101ea366004611b0b565b610513565b60405190151581526020016101d3565b61021261020d366004611b35565b61058c565b6040519081526020016101d3565b61021260025481565b610212610237366004611b35565b610703565b6102127f000000000000000000000000000000000000000000000000000000000000000081565b6101ef610271366004611b4e565b610868565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102e97f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016101d3565b6102126109ae565b6101ef610311366004611b0b565b610a09565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b61029d7f000000000000000000000000000000000000000000000000000000000000000081565b610212610372366004611b8a565b60036020526000908152604090205481565b610212610392366004611b8a565b60056020526000908152604090205481565b6102e97f000000000000000000000000000000000000000000000000000000000000000081565b6101c6610b28565b6101ef6103e1366004611b0b565b610b35565b6101ef6103f4366004611b0b565b610bc9565b610212610407366004611ba5565b610c4e565b61021261041a366004611ba5565b61102e565b61021261042d366004611b8a565b611400565b610445610440366004611be1565b61158d565b005b610212610455366004611b8a565b611873565b610212610468366004611c54565b600460209081526000928352604080842090915290825290205481565b6000805461049290611c87565b80601f01602080910402602001604051908101604052809291908181526020018280546104be90611c87565b801561050b5780601f106104e05761010080835404028352916020019161050b565b820191906000526020600020905b8154815290600101906020018083116104ee57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061057b9086815260200190565b60405180910390a350600192915050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156105be57506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa1580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190611cda565b9092509050806106f18386611d2d565b6106fb9190611d6a565b949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000042101561073557506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa158015610834573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108589190611cda565b9092509050816106f18286611d2d565b73ffffffffffffffffffffffffffffffffffffffff831660009081526004602090815260408083203384529091528120547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146108fc576108ca8382611da5565b73ffffffffffffffffffffffffffffffffffffffff861660009081526004602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff851660009081526003602052604081208054859290610931908490611da5565b909155505073ffffffffffffffffffffffffffffffffffffffff808516600081815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906109999087815260200190565b60405180910390a360019150505b9392505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146109e4576109df6118ce565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610a98576040517fdc84e3a400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000421115610b14576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a8f565b610b1e8484611968565b5060019392505050565b6001805461049290611c87565b60007f00000000000000000000000000000000000000000000000000000000000000003373ffffffffffffffffffffffffffffffffffffffff821614610bbf576040517fdc84e3a400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610a8f565b610b1e84846119e1565b33600090815260036020526040812080548391908390610bea908490611da5565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600081815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061057b9086815260200190565b60007f0000000000000000000000000000000000000000000000000000000000000000421015610ccc576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a8f565b6000610cd78561058c565b90503373ffffffffffffffffffffffffffffffffffffffff841603610e32576040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f00000000000000000000000000000000000000000000000000000000000000006044830152336064830152858116608483015260a482018390527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af1158015610e08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2c9190611dbc565b506106fb565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020908152604080832033845290915290205481811015610ea7576040517f4a1d07160000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610a8f565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915281208054849290610ee7908490611da5565b90915550506040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f000000000000000000000000000000000000000000000000000000000000000060448301528581166064830152868116608483015260a482018490527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af1158015611000573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110249190611dbc565b5050949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156110ac576040517f03a428230000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152602401610a8f565b3373ffffffffffffffffffffffffffffffffffffffff831603611206576040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f00000000000000000000000000000000000000000000000000000000000000006044830152336064830152848116608483015260a482018690527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af11580156111db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ff9190611dbc565b90506109a7565b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602090815260408083203384529091529020548481101561127b576040517f4a1d07160000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610a8f565b73ffffffffffffffffffffffffffffffffffffffff83166000908152600460209081526040808320338452909152812080548792906112bb908490611da5565b90915550506040517f52bc94300000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f000000000000000000000000000000000000000000000000000000000000000060448301528481166064830152858116608483015260a482018790527f000000000000000000000000000000000000000000000000000000000000000016906352bc94309060c4016020604051808303816000875af11580156113d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f89190611dbc565b9150506109a7565b60007f000000000000000000000000000000000000000000000000000000000000000042101561143257506000919050565b6040517ff7de8b1f0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060ff16600482015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660248301527f0000000000000000000000000000000000000000000000000000000000000000604483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063f7de8b1f906064016040805180830381865afa158015611531573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115559190611cda565b73ffffffffffffffffffffffffffffffffffffffff8616600090815260036020526040902054919350915082906106f1908390611d2d565b428410156115d0576040517fb979466100000000000000000000000000000000000000000000000000000000815260048101859052426024820152604401610a8f565b600060016115dc6109ae565b73ffffffffffffffffffffffffffffffffffffffff8a811660008181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e0830190915280519201919091207f190100000000000000000000000000000000000000000000000000000000000061010083015261010282019290925261012281019190915261014201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa15801561172e573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615806117a857508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b156117fd576040517feecd906e00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff89166024820152604401610a8f565b73ffffffffffffffffffffffffffffffffffffffff90811660009081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b60007f00000000000000000000000000000000000000000000000000000000000000004210156118a557506000919050565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60006040516119009190611dd5565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b806002600082825461197a9190611ea7565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054839290611a16908490611da5565b909155505060028054829003905560405181815260009073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016119d5565b600060208083528351808285015260005b81811015611a9c57858101830151858201604001528201611a80565b81811115611aae576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611b0657600080fd5b919050565b60008060408385031215611b1e57600080fd5b611b2783611ae2565b946020939093013593505050565b600060208284031215611b4757600080fd5b5035919050565b600080600060608486031215611b6357600080fd5b611b6c84611ae2565b9250611b7a60208501611ae2565b9150604084013590509250925092565b600060208284031215611b9c57600080fd5b6109a782611ae2565b600080600060608486031215611bba57600080fd5b83359250611bca60208501611ae2565b9150611bd860408501611ae2565b90509250925092565b600080600080600080600060e0888a031215611bfc57600080fd5b611c0588611ae2565b9650611c1360208901611ae2565b95506040880135945060608801359350608088013560ff81168114611c3757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008060408385031215611c6757600080fd5b611c7083611ae2565b9150611c7e60208401611ae2565b90509250929050565b600181811c90821680611c9b57607f821691505b602082108103611cd4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008060408385031215611ced57600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d6557611d65611cfe565b500290565b600082611da0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600082821015611db757611db7611cfe565b500390565b600060208284031215611dce57600080fd5b5051919050565b600080835481600182811c915080831680611df157607f831692505b60208084108203611e29577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b818015611e3d5760018114611e6c57611e99565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00861689528489019650611e99565b60008a81526020902060005b86811015611e915781548b820152908501908301611e78565b505084890196505b509498975050505050505050565b60008219821115611eba57611eba611cfe565b50019056fea2646970667358221220a9dfa34cdb38e357be86c7f468cdb26d0176343c00a885901b191e3f059521ad64736f6c634300080d0033",
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
