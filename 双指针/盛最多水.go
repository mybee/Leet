package main

import "fmt"

func main() {

	fmt.Println(maxArea([]int{
		1, 3, 5, 6, 2,
	}))

}

func maxArea(height []int) int {
	l, r := 0, len(height) -1
	maxC := 0
	for l < r {
		m := min(height[l], height[r]) * (r-l)
		maxC = max(maxC, m)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxC
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
