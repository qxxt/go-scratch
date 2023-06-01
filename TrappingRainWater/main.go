package main

import "fmt"

func main() {
	height := []int{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	var (
		left, right              = 0, len(height) - 1
		res, leftWall, rightWall int
	)

	for left < right {
		if height[left] < height[right] {
			if leftWall < height[left] {
				leftWall = height[left]
			} else {
				res += leftWall - height[left]
			}
			left++
		} else {
			if rightWall < height[right] {
				rightWall = height[right]
			} else {
				res += rightWall - height[right]
			}
			right--
		}
	}

	return res
}
