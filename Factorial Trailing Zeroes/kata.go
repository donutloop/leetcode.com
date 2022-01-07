package Factorial_Trailing_Zeroes

import "math/big"

func trailingZeroes(n int) int {
	s := factorial(big.NewInt(int64(n)))
	var c int
	b := s.String()
	for i := len(b) - 1; i > 0; i-- {
		if b[i] == '0' {
			c++
		} else {
			break
		}
	}
	return c
}

func init() {
	factorialResults = make(map[string]*big.Int)
}

var factorialResults map[string]*big.Int

func factorial(n *big.Int) *big.Int {

	if _, ok := factorialResults[n.String()]; ok {
		return factorialResults[n.String()]
	}

	if n.Cmp(big.NewInt(0)) > 0 {
		y := big.NewInt(0).Mul(n, factorial(big.NewInt(0).Sub(n, big.NewInt(1))))
		factorialResults[n.String()] = y
		return y
	}
	return big.NewInt(1)
}
