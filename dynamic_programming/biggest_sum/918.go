package biggest_sum

func maxSubarraySumCircular(nums []int) int {
	n, all, maxsum, minsum := len(nums), nums[0], nums[0], nums[0]
	var maxpre = nums[0]
	var minpre = nums[0]
	for i := 1; i < n; i++ {
		all += nums[i]
		maxpre = max(nums[i], nums[i]+maxpre)
		maxsum = max(maxsum, maxpre)

		minpre = min(nums[i], nums[i]+minpre)
		minsum = min(minsum, minpre)
	}

	if all == minsum {
		return maxsum
	}
	return max(maxsum, all-minsum)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
