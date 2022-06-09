package redeemertesting

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const ONE_ETH = 1000000000000000000
const ONE_GWEI = 1000000000
const ONE_WEI = 1

var ZERO = big.NewInt(0)

var CHAIN_ID = big.NewInt(1337)

var LENDER_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000001")
