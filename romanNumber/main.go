package main

import "fmt"

func main() {
	fmt.Println(romanToInt("DCXXI"))
}

func romanToInt(s string) (res int) {
	bit := 0
	for i := len(s) - 1; i != -1; i-- {
		switch s[i] {
		case 'M':
			bit |= 1
			res += 1000
		case 'D':
			bit |= 1
			res += 500
		case 'C':
			if bit&1 != 0 {
				res -= 100
				bit ^= 1
			} else {
				res += 100
				bit |= 2
			}
		case 'L':
			res += 50
			bit |= 2
		case 'X':
			if bit&2 != 0 {
				res -= 10
				bit ^= 2
			} else {
				res += 10
			}
			bit |= 4
		case 'V':
			res += 5
			bit |= 4
		case 'I':
			if bit&4 != 0 {
				res -= 1
				bit ^= 4
			} else {
				res += 1
			}
		}
	}
	return res
}
