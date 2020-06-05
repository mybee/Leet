package main

func main() {

}

func minPathSum(grid [][]int) int {
	// 状态转移方程:
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]

	dp := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		p := make([]int, len(grid[i]))
		dp[i] = p
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
			} else if i == 0 {
				dp[0][j] = dp[0][j-1] + grid[0][j]
			} else if j == 0 {
				dp[i][0] = dp[i-1][0] + grid[i][0]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}

		}
	}

	return dp[len(grid)-1][len(dp[len(grid)-1])-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
