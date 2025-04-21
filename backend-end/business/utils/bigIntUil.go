package utils

import (
	"fmt"
	"math"
	"math/big"
)

func BigIntToUint64(b *big.Int) (uint64, error) {
	// 检查是否为负数
	if b.Sign() < 0 {
		return 0, fmt.Errorf("cannot convert negative big.Int to uint64: %s", b.String())
	}

	// 检查是否超出 uint64 范围
	maxUint64 := new(big.Int).SetUint64(math.MaxUint64)
	if b.Cmp(maxUint64) > 0 {
		return 0, fmt.Errorf("big.Int value %s exceeds uint64 max", b.String())
	}

	return b.Uint64(), nil
}
