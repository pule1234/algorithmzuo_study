package longgestt_increasing_Subsequences

func lengthOfLIS(nums []int) int {
	n := len(nums)
	ends := make([]int, n)
	// len表示ends数组目前的有效区长度
	// ends[0...len-1]是有效区，有效区内的数字一定严格升序
	length := 0
	for i := 0; i < n; i++ {
		find := bs1(ends, length, nums[i])
		if find == -1 {
			ends[length] = nums[i]
			length++
		} else {
			ends[find] = nums[i]
		}
	}
	return length
}

// "最长递增子序列"使用如下二分搜索 :
// ends[0...len-1]是严格升序的，找到>=num的最左位置
// 如果不存在返回-1
func bs1(ends []int, len, num int) int {
	l, r, m, ans := 0, len-1, 0, -1
	for l <= r {
		m = (l + r) / 2
		if ends[m] >= num {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
