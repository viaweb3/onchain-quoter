package uniswap_v3

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/viaweb3/onchain-quoter/token"
)

var (
	wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	usdtAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	fee         = int64(500)
)

func newConnectedUniV3() *UniswapV3 {
	var (
		factoryAddress = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
		rpcApi         = "https://ethereum.blockpi.network/v1/rpc/public"
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	c, err := ethclient.DialContext(ctx, rpcApi)
	if err != nil {
		log.Fatal(err)
	}

	tdb := token.NewTokenDB(c)
	return New(c, tdb, factoryAddress)
}

func TestGetPrice(t *testing.T) {
	uni := newConnectedUniV3()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	price, err := uni.GetPrice(ctx, wethAddress, usdcAddress, fee)
	if err != nil {
		t.Log(err)
	}

	t.Log(wethAddress, usdcAddress, price)

	price, err = uni.GetPrice(ctx, usdtAddress, usdcAddress, fee)
	if err != nil {
		t.Log(err)
	}

	t.Log(usdtAddress, usdcAddress, price)
}
