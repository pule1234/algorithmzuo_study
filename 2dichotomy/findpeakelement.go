package _dichotomy

// 找峰值，在数组中找到数组的峰值， n-1< n  , n < n+1, 返回数组中的任意一个即可
func findpeakelement(nums []int) int {
	//首先判断index =0 的位置是不是峰值
	if nums[0] < nums[1] {
		return 0
	}
	// 判断index = len（nums） -1 的位置是不是峰值
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}

	// 若最右侧和最左侧均不存在峰值点，那么点一定在中间
	n := len(nums)
	left, right, ans, m := 1, n-2, -1, 0
	for left <= right {
		m = left + (right-left)/2
		if nums[m-1] > nums[m] { // 则说明在m的左侧一定存在峰值点
			right = m - 1
		} else if nums[m] < nums[m+1] { // 则说明m
			left = m + 1
		} else {
			ans = m
			return ans
		}
	}
	return -1
}
