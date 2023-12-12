package interval_dp

import "math"

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}

	presum := make([]int, n+1)
	for i, j, sum := 0, 1, 0; i < n; i, j = i+1, j+1 {
		sum += stones[i]
		presum[j] = sum
	}
	// dp[l][r] : l...r范围上的石头，合并到不能再合并（份数是确定的），最小代价是多少
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for l := n - 2; l >= 0; l-- {
		var ans int
		for r := l + 1; r < n; r++ {
			ans = math.MaxInt32
			for m := l; m < r; m += k - 1 {
				ans = min(ans, dp[l][m]+dp[m+1][r])
			}
			if (r-l)%(k-1) == 0 {
				ans += presum[r+1] - presum[l]
			}
			dp[l][r] = ans
		}
	}
	return dp[0][n-1]
}
