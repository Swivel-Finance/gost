package lendertesting

import (
	"math/big"

	"github.com/swivel-finance/gost/test/lender"
)

var TEST_MATURITY = big.NewInt(119240)

var ORDERS = []lender.SwivelOrder{
	{
		Vault:     true,
		Exit:      false,
		Principal: big.NewInt(11111),
		Premium:   big.NewInt(22222),
		Maturity:  TEST_MATURITY,
		Expiry:    big.NewInt(44444),
	},
	{
		Vault:     false,
		Exit:      true,
		Principal: big.NewInt(99999),
		Premium:   big.NewInt(88888),
		Maturity:  TEST_MATURITY,
		Expiry:    big.NewInt(66666),
	},
}

var AMOUNTS = []*big.Int{
	big.NewInt(111),
	big.NewInt(222),
}

var COMPONENTS = []lender.SwivelComponents{
	{
		V: 100,
		R: [32]byte{1, 2, 3},
		S: [32]byte{4, 5, 6},
	},
	{
		V: 200,
		R: [32]byte{7, 8, 9},
		S: [32]byte{10, 11, 12},
	},
}
