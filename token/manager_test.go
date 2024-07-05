package token

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	UniAddress = "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"
)

func TestGetToken(t *testing.T) {
	rpcApi := "https://ethereum.blockpi.network/v1/rpc/public"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	c, err := ethclient.DialContext(ctx, rpcApi)
	if err != nil {
		log.Fatal(err)
	}

	tdb := NewTokenDB(c)

	token, err := tdb.GetToken(ctx, UniAddress)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", token)

	assert.Equal(t, "UNI", token.Symbol)
}
