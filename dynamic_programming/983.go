package dynamic_programming

import (
	"math"
)

var durations = []int{1, 7, 30}

func mincostTickets(days []int, costs []int) int {

	return f1(days, costs, 0)
}

func f1(days []int, costs []int, i int) int {
	if i == len(days) {
		return 0
	}

	ans := math.MaxInt32

	for k, j := 0, i; k < 3; k++ {
		for j < len(days) && days[i]+durations[k] > days[j] {
			j++
		}
		ans = min(ans, costs[k]+f1(days, costs, j))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 记忆化搜索
// 暴力尝试改记忆化搜索
// 从顶到底的动态规划
func mincostTickets2(days []int, costs []int) int {
	dp := make([]int, len(days))
	for i := 0; i < len(days); i++ {
		dp[i] = math.MaxInt32
	}
	return f2(days, costs, dp, 0)
}

func f2(days, costs, dp []int, i int) int {
	if i == len(days) {
		return 0
	}

	if dp[i] != math.MaxInt32 {
		return dp[i]
	}

	ans := math.MaxInt32
	for k, j := 0, i; k < 3; k++ {
		for j < len(days) && days[i]+durations[k] > days[j] {
			j++
		}
		ans = min(ans, costs[k]+f2(days, costs, dp, j))
	}

	dp[i] = ans
	return ans
}

// 严格位置依赖的动态规划
// 从底到顶的动态规划
const MAXN = 366

var dp [MAXN]int

func mincostTickets3(days []int, costs []int) int {
	n := len(days)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		for k, j := 0, i; k < 3; k++ {
			for j < len(days) && days[i]+durations[k] > days[j] {
				j++
			}
			dp[i] = min(dp[i], costs[k]+dp[j])
		}
	}
	return dp[0]
}
