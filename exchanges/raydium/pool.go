package raydium

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
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

type Pool struct {
	Name       string
	Address    solana.PublicKey
	client     *rpc.Client
	Immutables PoolOpts
	State      PoolState
}

func NewPool(client *rpc.Client, name string, poolAddress string, immutables PoolOpts) (*Pool, error) {
	return &Pool{
		Name:       name,
		Address:    solana.MustPublicKeyFromBase58(poolAddress),
		client:     client,
		Immutables: immutables,
		State:      PoolState{},
	}, nil
}

func (p *Pool) UpdateState(ctx context.Context) error {
	err := p.client.GetAccountDataBorshInto(ctx, p.Address, &p.State)
	if err != nil {
		return err
	}
	return nil
}

func (p *Pool) PriceOf(token string) (float64, error) {
	var (
		token0Multiplier = new(big.Int).Exp(big.NewInt(10), big.NewInt(p.Immutables.Token0.Decimals), nil)
		token1Multiplier = new(big.Int).Exp(big.NewInt(10), big.NewInt(p.Immutables.Token1.Decimals), nil)
	)

	if token == p.Immutables.Token0.Address {
		numerator := new(big.Int).Exp(p.State.SqrtPriceX64.BigInt(), big.NewInt(2), nil)
		// multiply by token decimals
		numerator = numerator.Mul(numerator, token0Multiplier)
		n := new(big.Float).SetInt(numerator)

		denominator := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
		d := new(big.Float).SetInt(denominator)

		res := n.Quo(n, d)
		price, _ := res.Quo(res, new(big.Float).SetInt(token1Multiplier)).Float64()
		return price, nil
	} else if token == p.Immutables.Token1.Address {
		numerator := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)
		numerator = numerator.Mul(numerator, token1Multiplier)
		n := new(big.Float).SetInt(numerator)

		denominator := new(big.Int).Exp(p.State.SqrtPriceX64.BigInt(), big.NewInt(2), nil)
		d := new(big.Float).SetInt(denominator)

		res := n.Quo(n, d)
		price, _ := res.Quo(res, new(big.Float).SetInt(token0Multiplier)).Float64()

		return price, nil
	}

	return 0, ErrWrongToken
}

func Decode(poolBytes []byte) (*Pool, error) {
	buf := bytes.NewBuffer(poolBytes)
	dec := gob.NewDecoder(buf)

	var pool Pool
	if err := dec.Decode(&pool); err != nil {
		return nil, err
	}

	return &pool, nil
}

func (p Pool) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
