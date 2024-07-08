package raydium

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type RewardInfo struct {
	RewardState           uint8
	OpenTime              uint64
	EndTime               uint64
	LastUpdateTime        uint64
	EmissionsPerSecondX64 bin.Uint128
	RewardTotalEmissioned uint64
	RewardClaimed         uint64
	TokenMint             solana.PublicKey
	TokenVault            solana.PublicKey
	Authority             solana.PublicKey
	RewardGrowthGlobalX64 bin.Uint128
}

type PoolState struct {
	Blob                      [8]uint8
	Bump                      [1]uint8
	AmmConfig                 solana.PublicKey
	Owner                     solana.PublicKey
	TokenMint0                solana.PublicKey
	TokenMint1                solana.PublicKey
	TokenVault0               solana.PublicKey
	TokenVault1               solana.PublicKey
	ObservationKey            solana.PublicKey
	MintDecimals0             uint8
	MintDecimals1             uint8
	TickSpacing               uint16
	Liquidity                 bin.Uint128
	SqrtPriceX64              bin.Uint128
	TickCurrent               int32
	ObservationIndex          uint16
	ObservationUpdateDuration uint16
	FeeGrowthGlobal0X64       bin.Uint128
	FeeGrowthGlobal1X64       bin.Uint128
	ProtocolFeesToken0        uint64
	ProtocolFeesToken1        uint64
	SwapInAmountToken0        bin.Uint128
	SwapOutAmountToken1       bin.Uint128
	SwapInAmountToken1        bin.Uint128
	SwapOutAmountToken0       bin.Uint128
	Status                    uint8
	Padding                   [7]uint8
	RewardInfos               [3]RewardInfo
	TickArrayBitmap           [16]uint64
	TotalFeesToken0           uint64
	TotalFeesClaimedToken0    uint64
	TotalFeesToken1           uint64
	TotalFeesClaimedToken1    uint64
	FundFeesToken0            uint64
	FundFeesToken1            uint64
	OpenTime                  uint64
	Padding1                  [25]uint64
	Padding2                  [32]uint64
}
