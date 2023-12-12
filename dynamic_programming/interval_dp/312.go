package interval_dp

func maxCoins(nums []int) int {
	n := len(nums)
	arr := make([]int, n+2)
	arr[0] = 1
	arr[n+1] = 1

	for i := 1; i <= len(nums); i++ {
		arr[i] = nums[i-1]
	}
	dp := make([][]int, n+2)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+2)
		for j := i; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	return f(arr, 1, n, dp)
}

func f(arr []int, l, r int, dp [][]int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}

	ans := 0
	if l == r {
		ans = arr[l-1] * arr[l] * arr[r+1]
	} else {
		ans = max(arr[l-1]*arr[l]*arr[r+1]+f(arr, l+1, r, dp), arr[l-1]*arr[r]*arr[r+1]+f(arr, l, r-1, dp))
		for k := l + 1; k < r; k++ {
			ans = max(ans, arr[l-1]*arr[k]*arr[r+1]+f(arr, l, k-1, dp)+f(arr, k+1, r, dp))
		}
	}
	dp[l][r] = ans

	return ans
}
