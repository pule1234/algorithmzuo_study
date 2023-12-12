package main

// 暴力递归
func findTargetSumWays(nums []int, target int) int {
	return f1(nums, target, 0, 0)
}

func f1(nums []int, target int, i, sum int) int {
	if i == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}

	return f1(nums, target, i+1, sum+nums[i]) + f1(nums, target, i+1, sum-nums[i])
}

// 记忆化搜索
func findTargetSumWays(nums []int, target int) int {
	dp := make(map[int]map[int]int)
	return f2(nums, target, 0, 0, dp)
}

func f2(nums []int, target int, i, j int, dp map[int]map[int]int) int {
	if i == len(nums) {
		if j == target {
			return 1
		} else {
			return 0
		}
	}

	if _, ok := dp[i]; ok {
		if _, ok1 := dp[i][j]; ok1 {
			return dp[i][j]
		}
	}

	ans := f2(nums, target, i+1, j+nums[i], dp) + f2(nums, target, i+1, j-nums[i], dp)
	if dp[i] == nil {
		dp[i] = map[int]int{}
	}

	dp[i][j] = ans

	return ans
}

// 普通尝试
// 严格位置依赖的动态规划
// 平移技巧
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	if target < -s || target > s {
		return 0
	}

	n := len(nums)
	m := 2*n + 1
	// 原本的dp[i][j]含义:
	// nums[0...i-1]范围上，已经形成的累加和是sum
	// nums[i...]范围上，每个数字可以标记+或者-
	// 最终形成累加和为target的不同表达式数目
	// 因为sum可能为负数，为了下标不出现负数，
	// "原本的dp[i][j]"由dp表中的dp[i][j + s]来表示
	// 也就是平移操作！
	// 一切"原本的dp[i][j]"一律平移到dp表中的dp[i][j + s]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	// 原本的dp[n][target] = 1，平移！
	dp[n][target+s] = 1
	for i := n - 1; i >= 0; i-- {
		for j := -s; j <= s; j++ {
			if j+nums[i]+s < m {
				dp[i][j+s] = dp[i+1][j+nums[i]+s]
			}
			if j-nums[i]+s >= 0 {
				dp[i][j+s] += dp[i+1][j-nums[i]+s]
			}
		}
	}

	return dp[0][s]
}
