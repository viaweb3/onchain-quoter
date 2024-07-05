package raydium

import (
	"context"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/viaweb3/onchain-quoter/token"
	"math/big"
)

var (
	ErrWrongToken = errors.New("raydium pool: PriceOf: token not in pool")
)

type PoolOpts struct {
	Token0 token.Token
	Token1 token.Token
}

type PoolState struct {
	BlockTimestampLast uint32
	Reserve0           *big.Int
	Reserve1           *big.Int
}

type Pool struct {
	Name       string
	Address    solana.PublicKey
	client     *rpc.Client
	Immutables PoolOpts
	State      PoolState
}

func NewPool(client *rpc.Client, name string, poolAddress solana.PublicKey, immutables PoolOpts) (*Pool, error) {
	return &Pool{
		Name:       name,
		Address:    poolAddress,
		client:     client,
		Immutables: immutables,
		State:      PoolState{},
	}, nil
}

func (p *Pool) UpdateState(ctx context.Context) error {
	info, err := p.client.GetAccountInfo(ctx, p.Address)
	if err != nil {
		return err
	}

	spew.Dump(info)

	return nil
}
