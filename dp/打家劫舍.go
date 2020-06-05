package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 4, 1, 2}))
}

// 状态转移方程:
// dp[i] = MAX(dp[i-2] + A, dp[i-1])

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	l := make([]int, len(nums))
	l[0] = nums[0]
	l[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		l[i] = max(l[i-2]+nums[i], l[i-1])
	}

	return l[len(nums)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
