package main

import (
	"fmt"
	"sort"
)

func maxHeightDifference(n int, heights []int) int {
	sort.Ints(heights)

	totalDifference := 0

	for i := 0; i < n/2; i++ {
		totalDifference += abs(heights[i] - heights[n-1-i])
	}

	return totalDifference
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
	}

	fmt.Println(maxHeightDifference(n, heights))
}
