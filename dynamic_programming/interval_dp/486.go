package interval_dp

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	first := f(nums, 0, n-1, dp)

	second := sum - first
	return first >= second
}

func f(nums []int, l, r int, dp [][]int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}

	ans := 0
	if l == r {
		ans = nums[l]
	} else if l+1 == r {
		ans = max(nums[l], nums[r])
	} else {
		p1 := nums[l] + min(f(nums, l+2, r, dp), f(nums, l+1, r-1, dp))
		p2 := nums[r] + min(f(nums, l, r-2, dp), f(nums, l+1, r-1, dp))
		ans = max(p1, p2)
	}

	dp[l][r] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	for i := 0; i < n-1; i++ {
		dp[i][i] = nums[i]
		dp[i][i+1] = max(nums[i], nums[i+1])
	}

	dp[n-1][n-1] = nums[n-1]
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = max(nums[i]+min(dp[i+2][j], dp[i+1][j-1]), nums[j]+min(dp[i+1][j-1], dp[i][j-2]))
		}
	}

	fisrt := dp[0][n-1]
	second := sum - fisrt
	return fisrt >= second
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
