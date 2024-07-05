package uniswap_v2

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/viaweb3/onchain-quoter/token"
	"log"
	"testing"
	"time"
)

var (
	wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	usdtAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
)

func newConnectedUniV2() *UniswapV2 {
	var (
		factoryAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
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
	uni := newConnectedUniV2()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	price, err := uni.GetPrice(ctx, wethAddress, usdcAddress)
	if err != nil {
		t.Log(err)
	}

	t.Log(wethAddress, usdcAddress, price)

	price, err = uni.GetPrice(ctx, usdtAddress, usdcAddress)
	if err != nil {
		t.Log(err)
	}

	t.Log(usdtAddress, usdcAddress, price)
}
