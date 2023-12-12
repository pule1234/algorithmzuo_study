package Sliding_window

import "math"

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum-nums[l] >= target {
			sum -= nums[l]
			ans--
		}

		if sum >= target {
			ans = min(ans, r-l+1)
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
