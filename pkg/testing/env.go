package testing

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

// Env, holds Auth objects capable of signing transactions.
// Also holds the Geth simulated backend.
type Env struct {
	Alloc      core.GenesisAlloc
	Owner      *bind.TransactOpts
	Backend    *bind.TransactOpts
	User1      *bind.TransactOpts
	User2      *bind.TransactOpts
	Blockchain *backends.SimulatedBackend
}

// NewEnv returns a hydrated Env struct, ready for use.
// Given a balance argument, it assigns this as the wallet balance for
// each authorization object in the Ctx
func NewEnv(b *big.Int) *Env {
	owner := NewAuthObject()
	backend := NewAuthObject()
	u1 := NewAuthObject()
	u2 := NewAuthObject()
	alloc := make(core.GenesisAlloc)
	alloc[owner.From] = core.GenesisAccount{Balance: b}
	alloc[backend.From] = core.GenesisAccount{Balance: b}
	alloc[u1.From] = core.GenesisAccount{Balance: b}
	alloc[u2.From] = core.GenesisAccount{Balance: b}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &Env{
		Alloc:      alloc,
		Owner:      owner,
		Backend:    backend,
		User1:      u1,
		User2:      u2,
		Blockchain: bc,
	}
}
