package uniswap_v2

import (
	"context"
	"testing"
	"time"

	"github.com/viaweb3/onchain-quoter/token"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethUsdtPair = common.HexToAddress("0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852")
	token0      = token.NewToken("WETH", wethAddress, 18)
	token1      = token.NewToken("USDT", usdtAddress, 6)
)

var p, _ = newTestPair()
var client *ethclient.Client

func newTestPair() (*Pair, error) {
	var err error

	i := PairOpts{
		Token0: token0,
		Token1: token1,
	}

	rpcApi := "https://ethereum.blockpi.network/v1/rpc/public"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err = ethclient.DialContext(ctx, rpcApi)
	if err != nil {
		return nil, err
	}

	p, err := NewPair(client, "WETHUSDT", ethUsdtPair, i)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func TestNewPair(t *testing.T) {
	p, err := newTestPair()
	if err != nil {
		t.Fatalf("error creating pair: %s", err)
	}

	t.Log("new pair name:", p.Name)
}

func TestUpdateState(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := p.UpdateState(ctx)
	if err != nil {
		t.Fatalf("error updating state: %s", err)
	}

	emptyState := PairState{}

	if p.State == emptyState {
		t.Fatal("empty state")
	}

	t.Log("pool state:", p.State)
}

func TestPriceOf(t *testing.T) {
	price, err := p.PriceOf(token0.Address)
	if err != nil {
		t.Fatalf("error getting price: %s", err)
	}

	t.Log("price of", token0.Symbol, price)
}

func TestEncodeDecodePool(t *testing.T) {
	poolBytes, err := p.Encode()
	if err != nil {
		t.Fatal(err)
	}

	newPool, err := Decode(poolBytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", newPool)
}
