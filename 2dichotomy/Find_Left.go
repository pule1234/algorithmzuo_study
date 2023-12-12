package _dichotomy

// 查找大于或等于num的最左位置
func findleft(nums []int, num int) int {
	ans := -1
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left+(right-left)/2] < num {
			left = left + (right-left)/2 + 1
		} else {
			ans = left + (right-left)/2
			right = ans - 1
		}
	}
	return ans
}
