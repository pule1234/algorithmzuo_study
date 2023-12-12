package queue

import "math"

const maxn = 100001

var deque [maxn][2]int
var h, t int

func findMaxValueOfEquation(points [][]int, k int) int {
	h = 0
	t = 0
	n := len(points)
	ans := math.MinInt32
	for i := 0; i < n; i++ {
		x := points[i][0]
		y := points[i][1]
		for h < t && deque[h][0]+k < x {
			h++
		}

		if h < t {
			ans = max(ans, x+y+deque[h][1]-deque[h][0])
		}

		for h < t && deque[t-1][1]-deque[t-1][0] <= y-x {
			t--
		}
		deque[t][0] = x
		deque[t][1] = y
		t++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
