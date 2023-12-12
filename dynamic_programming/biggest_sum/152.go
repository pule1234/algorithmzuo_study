package biggest_sum

func maxProduct(nums []int) int {
	ans := nums[0]
	min := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		curmax := Max(nums[i], Max(min*nums[i], max*nums[i]))
		curmin := Min(nums[i], Min(min*nums[i], max*nums[i]))
		min = curmin
		max = curmax
		ans = Max(ans, max)
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
