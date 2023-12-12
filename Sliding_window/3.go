package Sliding_window

func lengthOfLongestSubstring(s string) int {
	str := []byte(s)
	n := len(str)

	//创建一个数组，表示每一种字符上一次出现的位置
	last := make([]int, 256)

	for i := 0; i < len(last); i++ {
		last[i] = -1 // 表示没有在上一次出现
	}

	ans := 0
	for l, r := 0, 0; r < n; r++ {
		// 看str【r】 之前有没有出现过，若出现过的话，比较大小，取最大值
		l = max(l, last[int(str[r])]+1)
		ans = max(ans, r-l+1)
		last[str[r]] = r
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
