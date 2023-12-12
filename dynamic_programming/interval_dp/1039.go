package interval_dp

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return f(values, 0, n-1, dp)
}

func f(values []int, l, r int, dp [][]int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}

	ans := math.MaxInt32

	if l == r || l+1 == r {
		ans = 0
	} else {
		for m := l + 1; m < r; m++ {
			ans = min(ans, f(values, l, m, dp)+f(values, m, r, dp)+values[m]*values[l]*values[r])
		}
	}
	dp[l][r] = ans
	return ans
}

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			dp[l][r] = math.MaxInt32
			for m := l + 1; m < r; m++ {
				dp[l][r] = min(dp[l][r], dp[l][m]+dp[m][r]+values[l]*values[r]*values[m])
			}
		}
	}

	return dp[0][n-1]
}
