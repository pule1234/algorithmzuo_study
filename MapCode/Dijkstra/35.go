package Dijkstra

import (
	"container/heap"
	"math"
)

func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	graph := make([][][]int, n)

	for _, path := range paths {
		graph[path[0]] = append(graph[path[0]], []int{path[1], path[2]})
		graph[path[1]] = append(graph[path[1]], []int{path[0], path[2]})
	}

	// （点，到达这个点的电量）
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, cnt+1)
		for j := 0; j < cnt+1; j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	distance[start][0] = 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, cnt+1)
	}
	// 0 : 当前点
	// 1 : 来到当前点的电量
	// 2 : 花费时间
	h := &hp{{start, 0, 0}}
	for h.Len() > 0 {
		record := heap.Pop(h).(pair)
		cur := record.x
		power := record.y
		cost := record.z
		if visited[cur][power] {
			continue
		}

		if cur == end {
			// 到达了终点
			return cost
		}

		visited[cur][power] = true
		if power < cnt { // 当前所剩的电量小于总电量
			//充一格电
			if !visited[cur][power+1] && cost+charge[cur] < distance[cur][power+1] {
				distance[cur][power+1] = cost + charge[cur]
				heap.Push(h, pair{cur, power + 1, cost + charge[cur]})
			}
		}
		for _, edge := range graph[cur] {
			// 不充电区别的城市
			nextCity := edge[0]
			restPower := power - edge[1]
			nextCost := cost + edge[1]
			if restPower >= 0 && !visited[nextCity][restPower] && nextCost < distance[nextCity][restPower] {
				distance[nextCity][restPower] = nextCost
				heap.Push(h, pair{nextCity, restPower, nextCost})
			}
		}
	}
	return -1
}

type pair struct {
	x, y, z int
}

type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].z < h[j].z }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
