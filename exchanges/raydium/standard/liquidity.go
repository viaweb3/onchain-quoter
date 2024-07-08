package standard

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/viaweb3/onchain-quoter/exchanges/raydium"
	"github.com/viaweb3/onchain-quoter/token"
	"math/big"
)

var (
	ErrWrongToken       = errors.New("raydium Standard pool: PriceOf: token not in pool")
	ErrorWrongTypeAMMV4 = errors.New("raydium Standard pool: Not Raydium AMM_V4 pool")
)

type LiquidityOpts struct {
	Token0 token.Token
	Token1 token.Token
}
type LiquidityExtras struct {
	BaseReserve  uint64
	QuoteReserve uint64
	MintAAmount  uint64
	MintBAmount  uint64
	PoolPrice    float64
}

type Liquidity struct {
	Name       string
	Address    solana.PublicKey
	client     *rpc.Client
	Immutables LiquidityOpts
	State      LiquidityV4State
	Extras     LiquidityExtras
}

// NewPool Just like Uniswap-V2
func NewPool(client *rpc.Client, name string, poolAddress string, immutables LiquidityOpts) (*Liquidity, error) {
	return &Liquidity{
		Name:       name,
		Address:    solana.MustPublicKeyFromBase58(poolAddress),
		client:     client,
		Immutables: immutables,
		State:      LiquidityV4State{},
		Extras:     LiquidityExtras{},
	}, nil
}

func (p *Liquidity) UpdateState(ctx context.Context) error {
	account, err := p.client.GetAccountInfo(ctx, p.Address)
	if err != nil {
		return err
	}

	if account.Value.Owner.String() != raydium.AMM_V4 {
		return ErrorWrongTypeAMMV4
	}

	err = bin.NewBorshDecoder(account.GetBinary()).Decode(&p.State)
	if err != nil {
		return err
	}

	return nil
}

func (p *Liquidity) UpdateVault(ctx context.Context) error {
	accounts, err := p.client.GetMultipleAccounts(ctx, p.State.BaseVault, p.State.QuoteVault)
	if err != nil {
		return err
	}

	for _, account := range accounts.Value {
		var acc Account
		if err = bin.NewBorshDecoder(account.Data.GetBinary()).Decode(&acc); err != nil {
			return err
		}
		if acc.Mint == p.State.BaseMint {
			p.Extras.BaseReserve = acc.Amount - p.State.BaseNeedTakePnl
			p.Extras.MintAAmount = acc.Amount
		}
		if acc.Mint == p.State.QuoteMint {
			p.Extras.QuoteReserve = acc.Amount - p.State.QuoteNeedTakePnl
			p.Extras.MintBAmount = acc.Amount
		}
	}

	baseDecimalFactor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.State.BaseDecimal)), nil))
	quoteDecimalFactor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.State.QuoteDecimal)), nil))

	baseReserveFloat := new(big.Float).SetInt(big.NewInt(int64(p.Extras.BaseReserve)))
	quoteReserveFloat := new(big.Float).SetInt(big.NewInt(int64(p.Extras.QuoteReserve)))

	baseAdjusted := new(big.Float).Quo(baseReserveFloat, baseDecimalFactor)
	quoteAdjusted := new(big.Float).Quo(quoteReserveFloat, quoteDecimalFactor)

	p.Extras.PoolPrice, _ = new(big.Float).Quo(quoteAdjusted, baseAdjusted).Float64()

	return nil
}

func (p *Liquidity) PriceOf(token string) (float64, error) {
	if token == p.Immutables.Token0.Address {
		return p.Extras.PoolPrice, nil
	} else if token == p.Immutables.Token1.Address {
		one := new(big.Float).SetFloat64(1)
		printF, _ := new(big.Float).Quo(one, new(big.Float).SetFloat64(p.Extras.PoolPrice)).Float64()
		return printF, nil
	}

	return 0, ErrWrongToken
}

func Decode(poolBytes []byte) (*Liquidity, error) {
	buf := bytes.NewBuffer(poolBytes)
	dec := gob.NewDecoder(buf)

	var liquidity Liquidity
	if err := dec.Decode(&liquidity); err != nil {
		return nil, err
	}

	return &liquidity, nil
}

func (p Liquidity) Encode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
