package testing

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewAuth() (*ecdsa.PrivateKey, *bind.TransactOpts) {
	pk, _ := crypto.GenerateKey()
	return pk, bind.NewKeyedTransactor(pk)
}

// GenBytes32 is a convenience function for some tests where a "GetHash" method is not available.
// Returns a type compatible with bytes32, given a string 32 chars or less.
func GenBytes32(s string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(s))
	return bytes
}

// NewCallOpts is a function which allows us to more succinctly place Call options
func NewCallOpts(a *bind.TransactOpts) *bind.CallOpts {
	return &bind.CallOpts{
		// TODO: Should this be a settable argument?
		Pending: false,
		From:    a.From,
	}
}

// NewTxOpts is a function which allows us to more succintly place the transaction options into a "send"
// type transaction. We do this in place of wrapping our TX in a global session variable, as, we tend to
// use different values in various places. Returns the assembled Geth TransactOpts
func NewTxOpts(a *bind.TransactOpts, v *big.Int, p *big.Int, l uint64) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     a.From,
		Signer:   a.Signer,
		Value:    v,
		GasPrice: p,
		GasLimit: l,
	}
}

// Commafy will take a big integer and return a string with commas so that logging
// big integers is human readable
func Commafy(n *big.Int) string {
	in := fmt.Sprintf("%d", n)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}
