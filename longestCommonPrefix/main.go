package main

import "fmt"

func main() {
	strs := []string{""}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) (res string) {
	var (
		letterIndex int
		letter      byte
	)

	for i := 0; i < len(strs[0]); i++ {
		letter = strs[0][i]
		for j := 1; ; j++ {
			if j == len(strs) {
				letterIndex++
				break
			}

			if len(strs[j]) <= i || strs[j][i] != letter {
				i = len(strs[0])
				break
			}
		}
	}

	return strs[0][:letterIndex]
}
