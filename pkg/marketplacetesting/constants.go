package marketplacetesting

import "math/big"

const ONE_ETH = 1000000000000000000
const ONE_GWEI = 1000000000
const ONE_WEI = 1

var ZERO = big.NewInt(0)

// MATURITY is one day, in seconds
const MATURITY = 86400

// MATURE_EVENT_SIG = crypto.Keccak256Hash([]byte("Mature(uint8,address,uint256,uint256,uint256))").Hex()
const MATURE_EVENT_SIG = "0xa43c0392e4bc23fcadd5a4c4d6d69a1148b6bcec3ac53d7654921bcc33f5addf"

var CHAIN_ID = big.NewInt(1337)
