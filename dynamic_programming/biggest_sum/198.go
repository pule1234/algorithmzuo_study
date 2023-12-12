package biggest_sum

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[1], nums[0])
	}
	// dp[i] : nums[0...i]范围上可以随意选择数字，但是不能选相邻数，能得到的最大累加和
	dp := make([]int, n)

	dp[0] = nums[0]

	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], max(nums[i], nums[i]+dp[i-2]))
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[1], nums[0])
	}
	prepre := nums[0]
	pre := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		cur := max(pre, max(nums[i], prepre+nums[i-2]))
		prepre = pre
		pre = cur
	}

	return pre

}
