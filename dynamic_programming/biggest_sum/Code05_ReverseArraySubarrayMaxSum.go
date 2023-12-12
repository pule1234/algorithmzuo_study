package biggest_sum

// 可以翻转1次的情况下子数组最大累加和
// 给定一个数组nums，
// 现在允许你随意选择数组连续一段进行翻转，也就是子数组逆序的调整
// 比如翻转[1,2,3,4,5,6]的[2~4]范围，得到的是[1,2,5,4,3,6]
// 返回必须随意翻转1次之后，子数组的最大累加和

func maxSumReverse2(nums []int) int {
	n := len(nums)

	// start[i] : 所有必须以i开头的子数组中，最大累加和是多少
	start := make([]int, n)
	start[n-1] = n - 1 // 最后只有一个数直接初始化

	for i := n - 2; i >= 0; i-- {
		start[i] = max(nums[i], nums[i]+start[i+1])
	}

	ans := nums[0]
	// end : 子数组必须以i-1结尾，其中的最大累加和
	end := nums[0]

	maxEnd := nums[0]
	// maxEnd :
	// 0~i-1范围上，
	// 子数组必须以0结尾，其中的最大累加和
	// 子数组必须以1结尾，其中的最大累加和
	// ...
	// 子数组必须以i-1结尾，其中的最大累加和
	// 所有情况中，最大的那个累加和就是maxEnd

	for i := 1; i < n; i++ {
		ans = max(ans, maxEnd+start[i])
		end = max(end, nums[i]+end)
		maxEnd = max(end, maxEnd)
	}

	ans = max(ans, maxEnd)
	return ans

}
