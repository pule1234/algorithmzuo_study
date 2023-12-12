package Sliding_window

func longestSubstring(s string, k int) int {
	str := []byte(s)
	n := len(str)
	cnts := make([]int, 256)
	ans := 0

	for require := 1; require <= 26; require++ {

		for i := 0; i < 256; i++ {
			cnts[i] = 0
		}
		for l, r, collect, satisfy := 0, 0, 0, 0; r < n; r++ {
			cnts[s[r]]++
			if cnts[int(str[r])] == 1 { // 证明这个字符此时是新来的
				collect++
			}

			if cnts[int(str[r])] == k {
				satisfy++
			}

			for collect > require {
				if cnts[int(str[l])] == 1 {
					collect--
				}
				if cnts[int(str[l])] == k {
					satisfy--
				}
				cnts[int(str[l])]--
				l++
			}
			if satisfy == require {
				ans = max(ans, r-l+1)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
