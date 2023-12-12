package main

// piles是一组一组的硬币
// m是容量，表示一定要进行m次操作
// dp[i][j] : 1~i组上，一共拿走j个硬币的情况下，获得的最大价值
// 1) 不要i组的硬币 : dp[i-1][j]
// 2) i组里尝试每一种方案
// 比如，i组里拿走前k个硬币的方案 : dp[i-1][j-k] + 从顶部开始前k个硬币的价值和
// 枚举每一个k，选出最大值
func maxValueOfCoins(piles [][]int, m int) int {
	n := len(piles)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		// i从1组开始（我们的设定），但是题目中的piles是从下标0开始的
		// 所以来到i的时候，piles.get(i-1)是当前组
		team := piles[i-1]
		t := min(len(team), m)
		preSum := make([]int, t+1)
		for j, sum := 0, 0; j < t; j++ {
			sum += team[j]
			preSum[j+1] = sum
		}

		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k <= min(t, j); k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k]+preSum[k])
			}
		}
	}

	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a

	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}