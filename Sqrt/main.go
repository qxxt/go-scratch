package main

import (
	"fmt"
)

func main() {
	testData := []int{2, 4, 7, 8, 64, 984230921}
	for _, v := range testData {
		fmt.Println("Target:", v)
		fmt.Println("Square Root:", mySqrt(v))
		fmt.Println("Square Root Long:", mySqrtFloat(float64(v)))
	}
}

func mySqrtFloat(x float64) float64 {
	res := x / 2
	for z := 0.0; z != res; {
		z = res
		res = (res + x/res) / 2
	}
	return res
}

func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
