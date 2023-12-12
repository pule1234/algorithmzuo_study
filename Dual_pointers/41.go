package Dual_pointers

// 0 ~ n 的范围  在理想情况下应该是 出现1 ~ n+1 的 数字
// 首先要设置垃圾区域， 垃圾区域为右边   最左边界为r
// 然后 设置已处理好的区域      最右边界为l
// 最后 中间为待处理的区域
func firstMissingPositive(nums []int) int {
	n := len(nums)
	l, r := 0, n // 初始状态为理想状态，即没有垃圾区

	for l < r {
		if nums[l] == l+1 {
			l++
		} else if nums[l] <= l || nums[l] > r || nums[nums[l]-1] == nums[l] { //第三个条件表示有重复值
			r--
			swap(nums, l, r)
		} else { //表示nums【l】在l - r 的范围中，直接将nums【l】 放在应该出现的位置
			swap(nums, l, nums[l]-1) // nums[l] - 1 为nums[l]应该在的下标位置
		}
	}

	return l + 1
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
