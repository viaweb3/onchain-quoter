package uniswap_v2

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	univ2factory "github.com/viaweb3/onchain-quoter/bindings/uniswap_v2/factory"
	"github.com/viaweb3/onchain-quoter/token"
	"github.com/viaweb3/onchain-quoter/utils"
	"log"
	"time"
)

var (
	ErrPoolNotFound = errors.New("uniswap v2 factory: pair not found")
)

type UniswapV2 struct {
	Client         *ethclient.Client
	tokenManager   token.TokenManager
	Factory        *univ2factory.FactoryCaller
	defaultTimeout time.Duration
}

func New(client *ethclient.Client, tokenManager token.TokenManager, factoryAddress common.Address) *UniswapV2 {
	factory, err := univ2factory.NewFactoryCaller(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &UniswapV2{
		Client:         client,
		tokenManager:   tokenManager,
		Factory:        factory,
		defaultTimeout: time.Second * 10,
	}
}

func (v2 *UniswapV2) GetPrice(ctx context.Context, token0, token1 common.Address) (float64, error) {
	pair, err := v2.GetPair(ctx, token0, token1)
	if err != nil {
		return 0, err
	}

	return pair.PriceOf(token0)
}

func (v2 *UniswapV2) GetPairAddress(ctx context.Context, token0, token1 common.Address) (common.Address, error) {
	token0, token1 = utils.SortTokens(token0, token1)
	zeroAddress := [20]byte{}
	pair, err := v2.Factory.GetPair(&bind.CallOpts{Context: ctx}, token0, token1)
	if err != nil {
		return zeroAddress, err
	}

	if pair == zeroAddress {
		return zeroAddress, ErrPoolNotFound
	}

	return pair, nil
}

func (v2 *UniswapV2) GetPair(ctx context.Context, token0, token1 common.Address) (*Pair, error) {
	token0, token1 = utils.SortTokens(token0, token1)
	pairAddr, err := v2.GetPairAddress(ctx, token0, token1)
	if err != nil {
		return nil, err
	}

	t0, err := v2.tokenManager.GetToken(ctx, token0)
	if err != nil {
		return nil, err
	}

	t1, err := v2.tokenManager.GetToken(ctx, token1)
	if err != nil {
		return nil, err
	}

	immutables := PairOpts{
		Token0: t0,
		Token1: t1,
	}

	pairName := t0.Symbol + t1.Symbol
	pool, err := NewPair(v2.Client, pairName, pairAddr, immutables)
	if err != nil {
		return nil, err
	}

	err = pool.UpdateState(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
