package biggest_sum

// 子序列累加和必须被7整除的最大累加和
// 给定一个非负数组nums，
// 可以任意选择数字组成子序列，但是子序列的累加和必须被7整除
// 返回最大累加和
// 对数器验证

func maxSum2(nums []int) int {
	n := len(nums)

	// dp[i][j] : nums[0...i-1]
	// nums前i个数形成的子序列一定要做到，子序列累加和%7 == j
	// 这样的子序列最大累加和是多少
	// 注意 : dp[i][j] == -1代表不存在这样的子序列

	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 7)
	}
	dp[0][0] = 0

	for j := 1; j < 7; j++ {
		dp[0][j] = -1
	}

	x, cur, need := 0, 0, 0
	for i := 1; i <= n; i++ {
		x = nums[i-1]
		cur = nums[i-1] % 7
		for j := 0; j < 7; j++ {
			dp[i][j] = dp[i-1][j]
			if cur <= j {
				need = j - cur
			} else {
				need = j - cur + 7
			}

			if dp[i-1][need] != -1 {
				dp[i][j] = max(dp[i][j], dp[i-1][need]+x)
			}
		}
	}

	return dp[n][0]
}
