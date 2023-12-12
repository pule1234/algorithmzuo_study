package longgestt_increasing_Subsequences

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		} else {
			return envelopes[i][1] > envelopes[j][1]
		}
	})
	n := len(envelopes)
	length := 0
	ends := make([]int, n)

	for i := 0; i < n; i++ {
		find := bs1(ends, length, envelopes[i][1])
		if find == -1 {
			ends[length] = envelopes[i][1]
			length++
		} else {
			ends[find] = envelopes[i][1]
		}
	}
	return length
}

func bs1(ends []int, len, num int) int {
	l, r, ans := 0, len-1, -1

	for l <= r {
		m := (l + r) / 2
		if ends[m] >= num {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return ans
}
