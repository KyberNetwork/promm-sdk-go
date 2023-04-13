package utils

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"

	"github.com/KyberNetwork/promm-sdk-go/constants"
)

var ErrInvalidInput = errors.New("invalid input")
var powers = []int64{128, 64, 32, 16, 8, 4, 2, 1}
var powerBigInts = make([]*big.Int, len(powers))

func init() {
	for i := range powers {
		powerBigInts[i] = big.NewInt(powers[i])
	}
}

func MostSignificantBit(x *big.Int) (int64, error) {
	if x.Cmp(constants.Zero) <= 0 {
		return 0, ErrInvalidInput
	}
	if x.Cmp(entities.MaxUint256) > 0 {
		return 0, ErrInvalidInput
	}
	var msb int64
	for i, power := range powers {
		min := new(big.Int).Exp(constants.Two, powerBigInts[i], nil)
		if x.Cmp(min) >= 0 {
			x = new(big.Int).Rsh(x, uint(power))
			msb += power
		}
	}
	return msb, nil
}
