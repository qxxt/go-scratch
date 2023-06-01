package main

import "fmt"

func main() {
	id := search([]int{5}, 5)
	if id == -1 {
		fmt.Println("not found")
	}

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	id = search(arr, arr[0]-1)
	if id == -1 {
		fmt.Println("not found")
	}

	for _, v := range arr {
		id := search(arr, v)
		if id == -1 {
			fmt.Println("not found")
			continue
		}
		fmt.Printf("idx: %v, val: %v\n", id, arr[id])
	}

	id = search(arr, arr[len(arr)-1]+1)
	if id == -1 {
		fmt.Println("not found")
	}
}

// O(log N)
func search(nums []int, target int) int {
	var (
		low  = 0
		high = len(nums) - 1
		mid  = low
	)

	if nums[mid] < target {
		low++
		mid = high
	}

	for low <= high {
		if cur := nums[mid]; cur == target {
			return mid
		} else if cur < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

		mid = (high + low) / 2
	}
	return -1
}
