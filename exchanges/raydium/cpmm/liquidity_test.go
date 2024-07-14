package standard

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/viaweb3/onchain-quoter/token"
	"testing"
	"time"
)

//https://api-v3.raydium.io/docs/

var (
	SOLRAYPool = "4y81XN75NGct6iUYkBp2ixQKtXdrQxxMVgFbFF9w5n4u"
	token0     = token.NewToken("SOL", "So11111111111111111111111111111111111111112", 9)
	token1     = token.NewToken("RAY", "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R", 6)
)

var p, _ = newTestPool()
var client *rpc.Client

func newTestPool() (*Liquidity, error) {
	var err error

	i := LiquidityOpts{
		Token0: token0,
		Token1: token1,
	}

	client = rpc.New(rpc.MainNetBeta.RPC)
	if err != nil {
		return nil, err
	}

	p, err := NewPool(client, "SOLRAY", SOLRAYPool, i)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func TestNewPool(t *testing.T) {
	p, err := newTestPool()
	if err != nil {
		t.Fatalf("error creating pool: %s", err)
	}

	t.Log("new pool name:", p.Name)
}

func TestUpdateState(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := p.UpdateState(ctx)
	if err != nil {
		t.Fatalf("error updating state: %s", err)
	}

	emptyState := CpmmPoolInfo{}

	if p.State == emptyState {
		t.Fatal("empty state")
	}

	t.Log("pool state:", p.State)
}

func TestUpdateVault(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := p.UpdateVault(ctx)
	if err != nil {
		t.Fatalf("error updating state: %s", err)
	}

	emptyExtras := LiquidityExtras{}

	if p.Extras == emptyExtras {
		t.Fatal("empty extras")
	}

	t.Log("pool extras:", p.Extras)
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
