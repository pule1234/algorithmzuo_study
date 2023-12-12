package Sliding_window

// Q W E R
// 0 1 2 3
// "W Q Q R R E"
// 1 0 0 3 3 2
// cnts[1] = 1;
// cnts[0] = 2;
// cnts[2] = 1;
// cnts[3] = 2;
func balancedString(s string) int {
	n := len(s)
	arr := make([]int, n)
	cnts := make([]int, 4)

	for i := 0; i < n; i++ {
		c := s[i]
		if c == 'W' {
			arr[i] = 1
		} else if c == 'E' {
			arr[i] = 2
		} else if c == 'R' {
			arr[i] = 3
		}
		cnts[arr[i]]++ // 代表每一个字符对应的数量
	}
	//每个字符应该出现的次数
	require := n / 4
	//最少应该修改的字符数量
	ans := n

	for l, r := 0, 0; l < n; l++ {
		for !ok(cnts, r-l, require) && r < n {
			cnts[arr[r]]--
			r++
		}
		if ok(cnts, r-l, require) {
			ans = min(ans, r-l)
		}
		cnts[arr[l]]++
	}

	return ans

}

func ok(cnts []int, len int, require int) bool {
	for i := 0; i < 4; i++ {
		if cnts[i] > require { // 窗口以外的词频有一种大于require
			return false
		}
		len -= require - cnts[i]
	}

	return len == 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
