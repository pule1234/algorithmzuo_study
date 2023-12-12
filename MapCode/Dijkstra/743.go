package Dijkstra

import (
	"container/heap"
	"math"
)

//// 动态建图和普通堆的实现
//func networkDelayTime(times [][]int, n int, k int) (ans int) {
//	// 建图
//	type edge struct {
//		to, time int
//	}
//	g := make([][]edge, n)
//	for _, t := range times {
//		x, y := t[0]-1, t[1]-1
//		g[x] = append(g[x], edge{y, t[2]})
//	}
//
//	const inf int = math.MaxInt64 / 2
//	dist := make([]int, n)
//	// 对每一个位置设置最大值
//	for i := range dist {
//		dist[i] = inf
//	}
//
//	dist[k-1] = 0
//	h := &hp{{0, k - 1}}
//	for h.Len() > 0 {
//		p := heap.Pop(h).(pair)
//		x := p.x
//		if dist[x] < p.d {
//			continue
//		}
//		for _, e := range g[x] { // x 为当前的点
//			y := e.to // 要到的点
//			if d := dist[x] + e.time; d < dist[y] {
//				dist[y] = d
//				heap.Push(h, pair{d, y})
//			}
//		}
//	}
//
//	for _, d := range dist {
//		if d == inf {
//			return -1
//		}
//		ans = max(ans, d)
//	}
//	return
//}
//
//// ，0  源点到当前点的距离  1 ：当前节点
//type pair struct {
//	d, x int
//}
//
//type hp []pair
//
//func (h hp) Len() int              { return len(h) }
//func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
//func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
//func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//
//	return b
//}

type pair struct {
	x, d int
}

type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }

func networkDelayTime(times [][]int, n int, k int) (ans int) {

	graph := make([][][]int, n)
	for _, time := range times {
		graph[time[0]-1] = append(graph[time[0]-1], []int{time[1] - 1, time[2]})
	}

	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt32
	}

	distance[k-1] = 0

	h := &hp{{k - 1, 0}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		cur := p.x
		cost := p.d
		if distance[cur] < cost {
			continue
		}

		for _, edge := range graph[cur] {
			to := edge[0] // 要到的点
			d := distance[cur] + edge[1]
			if d < distance[to] {
				distance[to] = d
				heap.Push(h, pair{to, d})
			}
		}
	}

	for _, dis := range distance {
		if dis == math.MaxInt32 {
			return -1
		}
		ans = max(ans, dis)
	}

	return ans
}
