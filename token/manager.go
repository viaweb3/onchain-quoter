package token

import (
	"context"
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/common"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/viaweb3/onchain-quoter/bindings/erc20"
)

var (
	ErrNotConnected = errors.New("not connected to a chain, use Connect method")
)

type TokenManager interface {
	GetToken(ctx context.Context, address string) (Token, error)
}

type tokenManager struct {
	client *ethclient.Client
	cache  *bigcache.BigCache
}

func NewTokenDB(client *ethclient.Client) TokenManager {
	c, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		panic(err)
	}

	return &tokenManager{
		cache:  c,
		client: client,
	}
}

// Adds a token to the cache
func (tc *tokenManager) add(token Token) {
	encoded, _ := token.Encode()
	tc.cache.Set(token.Address, encoded)
}

// Gets cached token by address if it's present.
func (tc tokenManager) get(address string) (Token, bool) {
	encoded, err := tc.cache.Get(address)
	if err != nil {
		return Token{}, false
	}
	t, err := Decode(encoded)
	if err != nil {
		return Token{}, false
	}
	return t, true
}

func (tc *tokenManager) GetToken(ctx context.Context, address string) (Token, error) {
	// Check cache
	if token, ok := tc.get(address); ok {
		return token, nil
	}

	// Check if we're connected to a chain
	if tc.client == nil {
		return Token{}, ErrNotConnected
	}

	token, err := erc20.NewErc20Caller(common.HexToAddress(address), tc.client)
	if err != nil {
		return Token{}, fmt.Errorf("getting token: %w", err)
	}

	opts := &bind.CallOpts{Context: ctx}
	sym, err := token.Symbol(opts)
	if err != nil {
		return Token{}, fmt.Errorf("getting token: reading name: %w", err)
	}

	decimals, err := token.Decimals(opts)
	if err != nil {
		return Token{}, fmt.Errorf("getting token: reading decimals: %w", err)
	}

	newToken := Token{
		Address:  address,
		Symbol:   sym,
		Decimals: int64(decimals),
	}

	tc.add(newToken)

	return newToken, nil
}
