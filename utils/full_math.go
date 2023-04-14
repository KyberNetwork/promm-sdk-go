package utils

import (
	"math/big"

	"github.com/KyberNetwork/promm-sdk-go/constants"
)

func MulDivRoundingUp(a, b, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	result := new(big.Int).Div(product, denominator)
	if new(big.Int).Rem(product, denominator).Cmp(constants.Zero) != 0 {
		result.Add(result, constants.One)
	}
	return result
}

func MulDivRoundingDown(a, b, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	return product.Quo(product, denominator)
}

func MulDiv(a, b, denominator *big.Int) *big.Int {
	product := new(big.Int).Mul(a, b)
	return product.Div(product, denominator)
}

func GetSmallerRootOfQuadEqn(a, b, c *big.Int) *big.Int {
	// smallerRoot = (b - sqrt(b * b - a * c)) / a;
	tmp1 := new(big.Int).Mul(b, b)
	tmp2 := new(big.Int).Mul(a, c)
	tmp3 := tmp1.Sqrt(tmp1.Sub(tmp1, tmp2))
	tmp3.Sub(b, tmp3)
	return tmp3.Div(tmp3, a)
}
