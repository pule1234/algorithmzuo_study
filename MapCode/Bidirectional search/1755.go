package Bidirectional_search

import (
	"math"
	"sort"
)

const MAXN = 1 << 20

var lsum [MAXN]int

var rsum [MAXN]int
// fill 记录有几种组合
var fill int

// i : 开头， E: 结尾     s : 前置的和     sum为原始的nums有几种组合
func collect(nums []int, i, e, s int, sum []int) {
	if i == e {
		sum[fill] = s
		fill++
	} else {
		// nums[i.....]这一组，相同的数字有几个
		j := i + 1
		for j < e && nums[j] == nums[i] {
			j++
		}
		// nums[ 1 1 1 1 1 2....
		//       i         j
		// j表示下一个不同数字的第一个位置
		for k := 0; k <= j-i; k++ {
			// k = 0个
			// k = 1个
			// k = 2个
			collect(nums, j, e, s+k*nums[i], sum)
		}
	}
}

func minAbsDifference(nums []int, goal int) int {
	n := len(nums)
	//记录大于0的数值的和
	var min int64

	//记录小于0 的数值的和
	var max int64

	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			max += int64(nums[i])
		} else {
			min += int64(nums[i])
		}
	}

	if max < int64(goal) {
		return int(math.Abs(float64(max - int64(goal))))
	}

	if min > int64(goal) {
		return int(math.Abs(float64(min - int64(goal))))
	}
	// 原始数组排序，为了后面递归的时候，还能剪枝
	// 常数优化
	sort.Ints(nums)
	fill = 0
	collect(nums, 0, n>>1, 0, lsum[:])
	lsize := fill
	fill = 0
	collect(nums, n>>1, n, 0, rsum[:])
	rsize := fill

	sort.Slice(lsum[:lsize], func(i, j int) bool {
		return lsum[i] < lsum[j]
	})

	sort.Slice(rsum[:rsize], func(i, j int) bool {
		return rsum[i] < rsum[j]
	})

	ans := int(math.Abs(float64(goal)))
	for i, j := 0, rsize-1; i < lsize; i++ {
		for j > 0 && math.Abs(float64(goal-lsum[i]-rsum[j-1])) <= math.Abs(float64(goal-lsum[i]-rsum[j])) {
			j-- // 表示右边的指针向左移动一格， 取一个更小的值
		}
		ans = Min(ans, int(math.Abs(float64(goal-lsum[i]-rsum[j]))))
	}

	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
