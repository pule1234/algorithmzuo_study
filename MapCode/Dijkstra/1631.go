package Dijkstra

import (
	"container/heap"
	"math"
)

var move = []int{-1, 0, 1, 0, -1}

func minimumEffortPath(heights [][]int) int {

	// （0，0）为原点 ， （x ,y ）为要去的点
	n := len(heights)
	m := len(heights[0])
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[0][0] = 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	h := &hp{{0, 0, 0}}

	for h.Len() > 0 {
		record := heap.Pop(h).(pair)
		x := record.x
		y := record.y
		d := record.d
		if visited[x][y] {
			continue
		}
		if x == n-1 && y == m-1 {
			return d
		}

		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + move[i]
			ny := y + move[i+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] {
				nd := max(d, int(math.Abs(float64(heights[x][y]-heights[nx][ny]))))
				if nd < distance[nx][ny] {
					distance[nx][ny] = nd
					heap.Push(h, pair{nx, ny, nd})
				}
			}
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type pair struct {
	x, y, d int
}

type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
