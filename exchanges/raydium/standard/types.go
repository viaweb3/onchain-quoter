package standard

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type Account struct {
	Mint                 solana.PublicKey
	Owner                solana.PublicKey
	Amount               uint64
	DelegateOption       uint32
	Delegate             solana.PublicKey
	State                uint8
	IsNativeOption       uint32
	IsNative             uint64
	DelegatedAmount      uint64
	CloseAuthorityOption uint32
	CloseAuthority       solana.PublicKey
}

type LiquidityV4State struct {
	Status                 uint64
	Nonce                  uint64
	MaxOrder               uint64
	Depth                  uint64
	BaseDecimal            uint64
	QuoteDecimal           uint64
	State                  uint64
	ResetFlag              uint64
	MinSize                uint64
	VolMaxCutRatio         uint64
	AmountWaveRatio        uint64
	BaseLotSize            uint64
	QuoteLotSize           uint64
	MinPriceMultiplier     uint64
	MaxPriceMultiplier     uint64
	SystemDecimalValue     uint64
	MinSeparateNumerator   uint64
	MinSeparateDenominator uint64
	TradeFeeNumerator      uint64
	TradeFeeDenominator    uint64
	PnlNumerator           uint64
	PnlDenominator         uint64
	SwapFeeNumerator       uint64
	SwapFeeDenominator     uint64
	BaseNeedTakePnl        uint64
	QuoteNeedTakePnl       uint64
	QuoteTotalPnl          uint64
	BaseTotalPnl           uint64
	PoolOpenTime           uint64
	PunishPcAmount         uint64
	PunishCoinAmount       uint64
	OrderbookToInitTime    uint64
	// u128('poolTotalDepositPc'),
	// u128('poolTotalDepositCoin'),
	SwapBaseInAmount   bin.Uint128
	SwapQuoteOutAmount bin.Uint128
	SwapBase2QuoteFee  uint64
	SwapQuoteInAmount  bin.Uint128
	SwapBaseOutAmount  bin.Uint128
	SwapQuote2BaseFee  uint64
	// amm vault
	BaseVault  solana.PublicKey
	QuoteVault solana.PublicKey
	// mint
	BaseMint  solana.PublicKey
	QuoteMint solana.PublicKey
	LpMint    solana.PublicKey
	// market
	OpenOrders      solana.PublicKey
	MarketId        solana.PublicKey
	MarketProgramId solana.PublicKey
	TargetOrders    solana.PublicKey
	WithdrawQueue   solana.PublicKey
	LpVault         solana.PublicKey
	Owner           solana.PublicKey
	// true circulating supply without lock up
	LpReserve uint64
	Padding   [3]uint64
}
