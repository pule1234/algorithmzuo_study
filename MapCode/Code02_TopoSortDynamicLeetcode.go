package main

// 建图 设置入度和出度，所有课程的入度为0即可
// /prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)

	indegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		indegree[edge[0]]++
	}

	queue := make([]int, numCourses)
	l, r := 0, 0

	//入度为0的节点，先进入队列
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue[r] = i
			r++
		}
	}

	var cnt int
	for l < r {
		cur := queue[l]
		l++
		cnt++

		for _, next := range graph[cur] {
			if indegree[next]--; indegree[next] == 0 {
				queue[r] = next
				r++
			}
		}
	}

	if cnt == numCourses {
		return queue
	} else {
		return []int{}
	}
}
