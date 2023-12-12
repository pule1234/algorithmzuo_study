package Bellman_Ford

import "math"

// Bellman-Ford算法应用（不是模版）
// 测试链接 : https://leetcode.cn/problems/cheapest-flights-within-k-stops/
// Bellman-Ford算法
// 针对此题改写了松弛逻辑，课上讲了细节
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		cur[i] = math.MaxInt32
	}
	cur[src] = 0
	// 中转的次数
	for i := 0; i <= k; i++ {
		next := make([]int, n)
		copy(next, cur)
		for _, edge := range flights {
			if cur[edge[0]] != math.MaxInt32 {
				// 更新边
				next[edge[1]] = min(next[edge[1]], cur[edge[0]]+edge[2])
			}
		}
		cur = next
	}

	if cur[dst] == math.MaxInt32 {
		return -1
	} else {
		return cur[dst]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
