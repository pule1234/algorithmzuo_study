package A_start

import (
	"container/heap"
	"math"
)

// A星算法
// grid[i][j] == 0 代表障碍
// grid[i][j] == 1 代表道路
// 只能走上、下、左、右，不包括斜线方向
// 返回从(startX, startY)到(targetX, targetY)的最短距离
// 和Dijkstra算法的区别 ，，   在将点上堆的时候   加入的距离不是源点到该点的距离，而是源点到该点的距离加上该点到终点的距离之和
var move = []int{-1, 0, 1, 0, -1}

func minDistance(grid [][]int, startX, startY, endX, endY int) int {
	if grid[startX][startY] == 0 || grid[endX][endY] == 0 {
		return -1
	}

	n := len(grid)
	m := len(grid[0])
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[startX][startY] = 1
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	h := &hp{{startX, startY, 1 + f1(startX, startY, endX, endY)}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(pair)
		x := cur.x
		y := cur.y
		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		if x == endX && y == endY {
			return distance[x][y]
		}
		for i := 0; i < 4; i++ {
			nx := x + move[i]
			ny := y + move[i+1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] && distance[x][y]+1 < distance[nx][ny] {
				distance[nx][ny] = 1 + distance[x][y]
				heap.Push(h, pair{nx, ny, distance[x][y] + 1 + f1(nx, ny, endX, endY)})
			}
		}
	}
	return -1
}

type pair struct {
	x, y, d int
}

type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }

// 曼哈顿距离
func f1(x, y, endx, endy int) int {
	return int(math.Abs(float64(endx-x)) + math.Abs(float64(endy-y)))
}

// 对角线距离
func f2(x, y, endx, endy int) int {
	return max(int(math.Abs(float64(endx-x))), int(math.Abs(float64(endy-y))))
}

// 欧式距离   勾股定理
func f3(x, y, endx, endy int) float64 {
	return math.Sqrt(math.Pow(float64(endx-x), 2) + math.Pow(float64(endy-y), 2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
