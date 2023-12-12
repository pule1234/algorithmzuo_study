package biggest_sum

import "math"

// 删掉1个数字后长度为k的子数组最大累加和
// 给定一个数组nums，求必须删除一个数字后的新数组中
// 长度为k的子数组最大累加和，删除哪个数字随意
// 对数器验证

func maxSum2(nums []int, k int) int {
	n := len(nums)
	if n <= k {
		return 0
	}

	// 求每一个长度为k+ 1 的子数组，然后再减去子数组中最小的值，最后求最大值
	// 单调队列   保存索引
	window := make([]int, n)
	l, r := 0, 0
	// 窗口累加和
	sum := 0
	ans := math.MinInt32

	for i := 0; i < n; i++ {
		// 单调队列 : i位置进入单调队列
		//nums[window[r-1]] 表示队列中最右侧元素对应的数组 nums 中的值。
		//nums[i] 表示当前遍历到的数组 nums 的元素值 // 维持单调递增
		for l < r && nums[window[r-1]] >= nums[i] {
			r--
		}
		window[r] = i
		r++
		sum += nums[i]
		if i >= k { // 窗口建立了
			ans = max(ans, sum-nums[window[l]])
			if window[l] == i-k {
				// 单调队列 : 如果单调队列最左侧的位置过期了，从队列中弹出
				l++
			}
			sum -= nums[i-k]
		}
	}

	return ans
}

func maxSum2(nums []int, k int) int {
	n := len(nums)
	if n <= k {
		return 0
	}

	window := make([]int, n)
	l, r := 0, 0
	ans := math.MinInt32
	sum := 0
	for i := 0; i < n; i++ {
		for l < r && nums[window[r-1]] >= nums[i] {
			r--
		}
		window[r] = i
		r++

		sum += nums[i]
		if i >= k {
			ans = max(ans, sum-nums[window[l]])
			if window[l] == i-k {
				l++
			}
			sum -= nums[window[l]]
		}
	}
	return ans
}
