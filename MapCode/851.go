package main

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(richer)
	graph := make([][]int, n)
	indegree := make([]int, n)

	for _, r := range richer {
		graph[r[0]] = append(graph[r[0]], r[1])
		indegree[r[1]]++
	}

	queue := make([]int, n)
	l, r := 0, 0

	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i
	}

	for l < r {
		cur := queue[l]
		l++
		for _, next := range graph[cur] {
			if quiet[ans[cur]] < quiet[ans[next]] {
				ans[next] = ans[cur]
			}
			if indegree[next]--; indegree[next] == 0 {
				queue[r] = next
				r++
			}
		}
	}
	return ans
}
