package uniswap_v3

import (
	"context"
	"errors"
	"github.com/viaweb3/onchain-quoter/utils"
	"log"
	"math/big"
	"time"

	"github.com/viaweb3/onchain-quoter/token"

	univ3factory "github.com/viaweb3/onchain-quoter/bindings/uniswap_v3/factory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrPoolNotFound = errors.New("uniswap v3 factory: pool not found")
)

type UniswapV3 struct {
	Client         *ethclient.Client
	tokenManager   token.TokenManager
	Factory        *univ3factory.Univ3factoryCaller
	defaultTimeout time.Duration
}

func New(client *ethclient.Client, tokenManager token.TokenManager, factoryAddress common.Address) *UniswapV3 {
	factory, err := univ3factory.NewUniv3factoryCaller(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &UniswapV3{
		Client:         client,
		tokenManager:   tokenManager,
		Factory:        factory,
		defaultTimeout: time.Second * 10,
	}
}

func (v3 *UniswapV3) GetPrice(ctx context.Context, token0, token1 common.Address, fee int64) (float64, error) {
	pool, err := v3.GetPool(ctx, token0, token1, fee)
	if err != nil {
		return 0, err
	}

	return pool.PriceOf(token0)
}

func (v3 *UniswapV3) GetPoolAddress(ctx context.Context, token0, token1 common.Address, fee int64) (common.Address, error) {
	token0, token1 = utils.SortTokens(token0, token1)
	zeroAddress := [20]byte{}
	pool, err := v3.Factory.GetPool(&bind.CallOpts{Context: ctx}, token0, token1, big.NewInt(fee))
	if err != nil {
		return zeroAddress, err
	}

	if pool == zeroAddress {
		return zeroAddress, ErrPoolNotFound
	}

	return pool, nil
}

func (v3 *UniswapV3) GetPool(ctx context.Context, token0, token1 common.Address, fee int64) (*Pool, error) {
	token0, token1 = utils.SortTokens(token0, token1)
	poolAddr, err := v3.GetPoolAddress(ctx, token0, token1, fee)
	if err != nil {
		return nil, err
	}

	t0, err := v3.tokenManager.GetToken(ctx, token0)
	if err != nil {
		return nil, err
	}

	t1, err := v3.tokenManager.GetToken(ctx, token1)
	if err != nil {
		return nil, err
	}

	immutables := PoolOpts{
		Token0: t0,
		Token1: t1,
		Fee:    fee,
	}

	poolName := t0.Symbol + t1.Symbol
	pool, err := NewPool(v3.Client, poolName, poolAddr, immutables)
	if err != nil {
		return nil, err
	}

	err = pool.UpdateState(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
