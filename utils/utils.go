package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

func SortTokens(tokenA, tokenB common.Address) (token0, token1 common.Address) {
	token0 = tokenA
	token1 = tokenB
	if tokenB.String() < tokenA.String() {
		token0 = tokenB
		token1 = tokenA
	}
	return
}
