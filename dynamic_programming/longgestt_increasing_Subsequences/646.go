package longgestt_increasing_Subsequences

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	n := len(pairs)
	ends := make([]int, n)

	length := 0
	for i := 0; i < n; i++ {
		find := bs(ends, length, pairs[i][0])
		if find == -1 {
			ends[length] = pairs[i][1]
			length++
		} else {
			ends[find] = min(ends[find], pairs[i][1])
		}
	}
	return length
}

func bs(ends []int, len, num int) int {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
