package whirlpools

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type WhirlpoolRewardInfo struct {
	Mint                  solana.PublicKey
	Vault                 solana.PublicKey
	Authority             solana.PublicKey
	EmissionsPerSecondX64 bin.Uint128
	GrowthGlobalX64       bin.Uint128
}

type PoolState struct {
	Blob                       [8]uint8
	WhirlpoolsConfig           solana.PublicKey
	WhirlpoolBump              [1]uint8
	TickSpacing                uint16
	TickSpacingSeed            [2]uint8
	FeeRate                    uint16
	ProtocolFeeRate            uint16
	Liquidity                  bin.Uint128
	SqrtPrice                  bin.Uint128
	TickCurrentIndex           int32
	ProtocolFeeOwedA           uint64
	ProtocolFeeOwedB           uint64
	TokenMintA                 solana.PublicKey
	TokenVaultA                solana.PublicKey
	FeeGrowthGlobalA           bin.Uint128
	TokenMintB                 solana.PublicKey
	TokenVaultB                solana.PublicKey
	FeeGrowthGlobalB           bin.Uint128
	RewardLastUpdatedTimestamp uint64
	RewardInfos                [3]WhirlpoolRewardInfo
}
