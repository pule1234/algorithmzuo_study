package biggest_sum

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	return max(best(nums, 1, n-1), nums[0]+best(nums, 2, n-2))
}

func best(nums []int, l, r int) int {
	if l > r {
		return 0
	}
	if l == r {
		return nums[l]
	}

	if l+1 == r {
		return max(nums[l], nums[l+1])
	}

	prepre := nums[l]
	pre := max(nums[l], nums[l+1])
	for i := l + 2; i <= r; i++ {
		cur := max(pre, max(nums[i], nums[i]+prepre))
		prepre = pre
		pre = cur
	}

	return pre

}
