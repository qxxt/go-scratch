package main

import (
	"fmt"
	"math"
)

func main() {
	testData := []struct {
		base float64
		pow  int
	}{
		{1, 2147483647},
		{-1, -2147483647},
		{1, 213},
		{0, 0},
		{2, 20},
	}
	for _, v := range testData {
		fmt.Println("base", v.base, "pow", v.pow)
		fmt.Println("mypow  ", myPow(v.base, v.pow))
		fmt.Println("mathPow", math.Pow(v.base, float64(v.pow)))
		fmt.Println("===")
	}
}

// Binary exponentiation as explained from:
// https://leetcode.com/problems/powx-n/editorial/
// Implemented in GO
// O(Log N)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n *= -1
		x = 1 / x
	}
	res := x

	for n > 1 {
		if res == 0 {
			return res
		}

		res *= x
		if n%2 == 1 {
			n--
		} else {
			x *= x
			n /= 2
		}
	}
	return res
}
