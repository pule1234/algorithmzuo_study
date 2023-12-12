package biggest_sum

import (
	"math"
)

func getMaxMatrix(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	a, b, c, d := 0, 0, 0, 0
	ans := math.MinInt32

	for up := 0; up < n; up++ {
		nums := make([]int, m)
		for down := up; down < n; down++ {
			for l, r, pre := 0, 0, math.MinInt32; r < m; r++ {
				nums[r] += matrix[down][r]
				if pre >= 0 {
					pre += nums[r]
				} else {
					pre = nums[r]
					l = r
				}

				if pre > ans {
					ans = pre
					a, b, c, d = up, l, down, r
				}
			}
		}
	}

	return []int{a, b, c, d}
}
