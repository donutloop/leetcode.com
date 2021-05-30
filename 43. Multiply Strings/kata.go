package kata

import "math/big"

func multiply(num1 string, num2 string) string {
	a, _ := big.NewInt(0).SetString(num1, 10)
	b, _ := big.NewInt(0).SetString(num2, 10)
	y := big.NewInt(0).Mul(a, b)
	return y.String()
}