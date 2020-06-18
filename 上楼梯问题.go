package main

import "fmt"

func main() {
	fmt.Println(waysToStep(6))
}
// 有 一步 两步 三步 步幅
// 找到状态转移方程:
// dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
// for 循环就完了

func waysToStep(n int) int {

	dp := make([]int, n+1)
	dp[1] = 1

	if n >= 2 {  // 这里要 大于等于
		dp[2] = 2
	}

	if n >= 3 {
		dp[3] = 4
	}


	for i:=4; i< n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		dp[i] = dp[i] % 1000000007
	}

	return dp[n]
}