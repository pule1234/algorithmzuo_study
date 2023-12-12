package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	indegree := make([]int, n)

	for i := 0; i < n; i++ {
		indegree[favorite[i]]++
	}

	queue := make([]int, n)
	l, r := 0, 0

	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	// 不包括i在内，i之前的最长链的长度
	deep := make([]int, n)
	for l < r {
		cur := queue[l]
		l++
		next := favorite[cur]
		deep[next] = max(deep[next], deep[cur]+1)
		if indegree[next]--; indegree[next] == 0 {
			queue[r] = next
			r++
		}
	}

	// 目前图中的点，不在环上的点，都删除了！ indegree[i] == 0
	// 可能性1 : 所有小环(中心个数 == 2)，算上中心点 + 延伸点，总个数
	sumOfSmallRings := 0
	// 可能性2 : 所有大环(中心个数 > 2)，只算中心点，最大环的中心点个数
	var bigRings = 0

	for i := 0; i < n; i++ {
		//只关心环
		if indegree[i] > 0 {
			ringSize := 1
			indegree[i] = 0

			for j := favorite[i]; j != i; j = favorite[j] {
				ringSize++
				indegree[j] = 0
			}

			if ringSize == 2 {
				sumOfSmallRings += 2 + deep[i] + deep[favorite[i]]
			} else {
				bigRings = max(bigRings, ringSize)
			}
		}
	}

	return max(sumOfSmallRings, bigRings)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}