package queue

import "math"

var sum [100001]int
var deque [100001]int
var h, t int

func shortestSubarray(nums []int, k int) int {

	n := len(nums)
	for i := 0; i < n; i++ { // 记录前缀和的值
		sum[i+1] = sum[i] + nums[i]
	}
	h = 0
	t = 0
	ans := math.MaxInt32
	for i := 0; i <= n; i++ {
		// 前0个数前缀和
		// 前1个数前缀和
		// 前2个数前缀和
		// ...
		// 前n个数前缀和
		// 和deque中的第一个数相减
		for h != t && sum[i]-sum[deque[h]] >= k {
			ans = min(ans, i-deque[h])
			h++
		}

		for h != t && sum[deque[t-1]] >= sum[i] {
			t--
		}
		deque[t] = i
		t++
	}
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return b

}
