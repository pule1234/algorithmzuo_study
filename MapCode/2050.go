package main

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n+1)

	indegree := make([]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = []int{}
	}
	for _, value := range relations {
		graph[value[0]] = append(graph[value[0]], value[1])
		indegree[value[1]]++
	}

	queue := make([]int, n)
	l, r := 0, 0
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	cost := make([]int, n+1)
	ans := 0

	for l < r {
		cur := queue[l]
		l++
		cost[cur] += time[cur-1]
		ans = max(ans, cost[cur])

		for _, next := range graph[cur] {
			cost[next] = max(cost[next], cost[cur])
			if indegree[next]--; indegree[next] == 0 {
				queue[r] = next
				r++
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
