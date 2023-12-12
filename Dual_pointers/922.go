package Dual_pointers

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	for even, odd := 0, 1; even < n && odd < n; {
		if (nums[n-1] & 1) == 1 {
			//奇数
			swap(nums, odd, n-1)
			odd += 2
		} else {
			swap(nums, even, n-1)
			even += 2
		}
	}
	return nums
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
