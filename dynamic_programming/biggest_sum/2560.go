package biggest_sum

func minCapability(nums []int, k int) int {

	n, l, r := len(nums), nums[0], nums[0]

	for i := 1; i < n; i++ {
		l = min(l, nums[i])
		r = max(r, nums[r])
	}

	m, ans := 0, 0
	for l <= r {
		m = (l + r) / 2
		if mostRob1(nums, n, m) >= k {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}

func mostRob1(nums []int, n, ability int) int {
	if n == 1 {
		if nums[0] <= ability {
			return 1
		}
		return 0
	}

	if n == 2 {
		if nums[0] <= ability || nums[1] <= ability {
			return 1
		}

		return 0
	}

	dp := make([]int, n)
	if nums[0] <= ability {
		dp[0] = 1
	}

	if nums[0] <= ability || nums[1] <= ability {
		dp[1] = 1
	}

	for i := 2; i < n; i++ {
		var cur = dp[i-2]
		if nums[i] <= ability {
			cur = 1 + dp[i-2]
		}

		dp[i] = max(dp[i-1], cur)
	}
	return dp[n-1]
}
