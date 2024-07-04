package uniswap_v2

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	univ2pair "github.com/viaweb3/onchain-quoter/bindings/uniswap_v2/pair"
	"github.com/viaweb3/onchain-quoter/token"
	"math/big"
)

var (
	ErrWrongToken = errors.New("uniswap v2 pool: PriceOf: token not in pool")
)

type PairOpts struct {
	Token0 token.Token
	Token1 token.Token
}

type PairState struct {
	BlockTimestampLast uint32
	Reserve0           *big.Int
	Reserve1           *big.Int
}

type Pair struct {
	Name       string
	Address    common.Address
	caller     *univ2pair.PairCaller
	Immutables PairOpts
	State      PairState
}

func NewPair(client *ethclient.Client, name string, pairAddress common.Address, immutables PairOpts) (*Pair, error) {
	caller, err := univ2pair.NewPairCaller(pairAddress, client)
	if err != nil {
		return nil, err
	}

	return &Pair{
		Name:       name,
		Address:    pairAddress,
		caller:     caller,
		Immutables: immutables,
		State:      PairState{},
	}, nil
}

func (p *Pair) UpdateState(ctx context.Context, client *ethclient.Client) error {
	opts := &bind.CallOpts{Context: ctx}

	caller, err := univ2pair.NewPairCaller(p.Address, client)
	if err != nil {
		return err
	}

	reserves, err := caller.GetReserves(opts)
	if err != nil {
		return err
	}

	p.State.BlockTimestampLast = reserves.BlockTimestampLast
	p.State.Reserve0 = reserves.Reserve0
	p.State.Reserve1 = reserves.Reserve1

	return nil
}

func (p *Pair) PriceOf(token common.Address) (float64, error) {
	var (
		token0Multiplier = new(big.Int).Exp(big.NewInt(10), big.NewInt(p.Immutables.Token0.Decimals), nil)
		token1Multiplier = new(big.Int).Exp(big.NewInt(10), big.NewInt(p.Immutables.Token1.Decimals), nil)
	)

	if token == p.Immutables.Token0.Address {
		numerator := new(big.Float).Quo(new(big.Float).SetInt(p.State.Reserve1), new(big.Float).SetInt(token1Multiplier))
		denominator := new(big.Float).Quo(new(big.Float).SetInt(p.State.Reserve0), new(big.Float).SetInt(token0Multiplier))
		priceF, _ := numerator.Quo(numerator, denominator).Float64()
		return priceF, nil
	} else if token == p.Immutables.Token1.Address {
		numerator := new(big.Float).Quo(new(big.Float).SetInt(p.State.Reserve0), new(big.Float).SetInt(token0Multiplier))
		denominator := new(big.Float).Quo(new(big.Float).SetInt(p.State.Reserve1), new(big.Float).SetInt(token1Multiplier))
		priceF, _ := numerator.Quo(numerator, denominator).Float64()
		return priceF, nil
	}

	return 0, ErrWrongToken
}

func Decode(poolBytes []byte) (*Pair, error) {
	buf := bytes.NewBuffer(poolBytes)
	dec := gob.NewDecoder(buf)

	var pair Pair
	if err := dec.Decode(&pair); err != nil {
		return nil, err
	}

	return &pair, nil
}

func (p Pair) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
