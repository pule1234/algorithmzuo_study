package interval_dp

import (
	"math"
	"sort"
)

func minCost(n int, cuts []int) int {
	m := len(cuts)
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i] < cuts[j]
	})
	arr := make([]int, m+2)
	arr[0] = 0
	arr[m+1] = n
	for i := 1; i <= m; i++ {
		arr[i] = cuts[i-1]
	}
	dp := make([][]int, m+2)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, m+2)
		for j := 1; j <= m; j++ {
			dp[i][j] = -1
		}
	}

	return f(arr, 1, m, dp)
}

func f(arr []int, l, r int, dp [][]int) int {
	if l > r {
		return 0
	}

	if l == r {
		return arr[r+1] - arr[l-1]
	}
	if dp[l][r] != -1 {
		return dp[l][r]
	}

	ans := math.MaxInt32
	for k := l; k <= r; k++ {
		ans = min(ans, f(arr, l, k-1, dp)+f(arr, k+1, r, dp))
	}

	ans += arr[r+1] - arr[l-1]
	dp[l][r] = ans
	return ans
}
