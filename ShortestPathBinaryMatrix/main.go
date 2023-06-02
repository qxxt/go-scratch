package main

import (
	"fmt"
)

func main() {
	testData := [][][]int{
		{
			{0, 1},
			{1, 0},
		},
		{
			{0, 0, 0},
			{1, 1, 1},
			{1, 1, 0},
		},
		{
			{0, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		},
		{
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		},
		{
			// 0  1  2  3  4  5
			{0, 1, 1, 0, 0, 0}, // 0
			{0, 1, 0, 1, 1, 0}, // 1
			{0, 1, 1, 0, 1, 0}, // 2
			{0, 0, 0, 1, 1, 0}, // 3
			{1, 1, 1, 1, 1, 0}, // 4
			{1, 1, 1, 1, 1, 0}, // 5
		},
		{
			// 0  1  2  3  4  5
			{0, 1, 1, 0, 0, 0}, // 0
			{0, 0, 0, 0, 1, 0}, // 1
			{0, 0, 0, 0, 1, 0}, // 2
			{0, 0, 0, 1, 0, 0}, // 3
			{1, 1, 1, 1, 0, 0}, // 4
			{1, 1, 1, 1, 1, 0}, // 5
		},
	}

	for i := range testData {
		fmt.Println(shortestPathBinaryMatrix(testData[i]))
	}
}

func shortestPathBinaryMatrix(grid [][]int) int {
	var (
		gridLength = len(grid) - 1
		x, y       int
		queue      = [][2]int{{0, 0}}
	)

	if grid[0][0]+grid[gridLength][gridLength] != 0 {
		return -1
	}

	moves := [][2]int{
		{1, 1},
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}

	grid[x][y] = 1
	for x+y != gridLength*2 {
		if len(queue) == 0 {
			return -1
		}

		x, y = queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, move := range moves {
			nextX, nextY := x+move[0], y+move[1]
			if nextX < 0 || nextY < 0 || nextX > gridLength || nextY > gridLength || grid[nextX][nextY] != 0 {
				continue
			}

			grid[nextX][nextY] = grid[x][y] + 1
			queue = append(queue, [2]int{nextX, nextY})
		}
	}

	return grid[x][y]
}
