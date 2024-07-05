package utils

import "github.com/ethereum/go-ethereum/common"

func SortTokens(tokenA, tokenB string) (token0, token1 string) {
	token0 = tokenA
	token1 = tokenB
	if common.HexToAddress(tokenB).String() < common.HexToAddress(tokenA).String() {
		token0 = tokenB
		token1 = tokenA
	}
	return
}
