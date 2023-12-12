package biggest_sum

import "math"

// 魔法卷轴
// 给定一个数组nums，其中可能有正、负、0
// 每个魔法卷轴可以把nums中连续的一段全变成0
// 你希望数组整体的累加和尽可能大
// 卷轴使不使用、使用多少随意，但一共只有2个魔法卷轴
// 请返回数组尽可能大的累加和

func maxSum2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 情况一
	p1 := 0
	for _, v := range nums {
		p1 += v
	}
	// prefix[i] : 0~i范围上一定要用1次卷轴的情况下，0~i范围上整体最大累加和多少
	prefix := make([]int, n)
	// 每一步的前缀和
	sum := nums[0]
	// maxPresum : 之前所有前缀和的最大值
	maxPresum := max(0, nums[0])
	for i := 1; i < n; i++ {
		prefix[i] = max(prefix[i-1]+nums[i], maxPresum)
		sum += nums[i]
		maxPresum = max(maxPresum, sum)
	}

	/// 情况二
	p2 := prefix[n-1]
	// suffix[i] : i~n-1范围上一定要用1次卷轴的情况下，i~n-1范围上整体最大累加和多少
	suffix := make([]int, n)
	sum = nums[n-1]
	maxPresum = max(0, sum)
	for i := n - 2; i >= 0; i-- {
		suffix[i] = max(nums[i]+suffix[i-1], maxPresum)
		sum += nums[i]
		maxPresum = max(maxPresum, sum)
	}

	// 情况三
	p3 := math.MinInt32
	for i := 1; i < n; i++ {
		p3 = max(p3, prefix[i-1]+suffix[i])
	}

	return max(p1, max(p2, p3))
}
