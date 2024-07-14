package standard

import (
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

type CpmmPoolInfo struct {
	Blob               [8]uint8
	AmmConfig          solana.PublicKey
	PoolCreator        solana.PublicKey
	Token0Vault        solana.PublicKey
	Token1Vault        solana.PublicKey
	LpMint             solana.PublicKey
	Token0Mint         solana.PublicKey
	Token1Mint         solana.PublicKey
	Token0Program      solana.PublicKey
	Token1Program      solana.PublicKey
	ObservationKey     solana.PublicKey
	AuthBump           uint8
	Status             uint8
	LpMintDecimals     uint8
	Mint0Decimals      uint8
	Mint1Decimals      uint8
	LpSupply           uint64
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64
	FundFeesToken0     uint64
	FundFeesToken1     uint64
	OpenTime           uint64
	Padding            [32]uint64
}
