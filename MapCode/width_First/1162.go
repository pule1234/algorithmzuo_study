package width_First

const MAXN = 101
const MAXM = 101

var queue [MAXN * MAXM][2]int // 遇到的每个点进队列，从头进从尾出
var l, r int

var visited [MAXN][MAXM]bool // 记录该点是否遇到过

func maxDistance(grid [][]int) int {
	move := []int{-1, 0, 1, 0, -1} // 使用数组表示坐标的移动
	l = 0
	r = 0

	n := len(grid)
	m := len(grid[0])
	seas := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				visited[i][j] = true
				queue[r][0] = i
				queue[r][1] = j
				r++
			} else {
				visited[i][j] = false
				seas++
			}
		}
	}

	if seas == 0 || seas == n*m {
		return -1
	}

	level := 0

	for l < r {
		level++
		size := r - l
		for k := 0; k < size; k++ {
			// 从队列中取出坐标点
			x := queue[l][0]
			y := queue[l][1]
			l++
			for i := 0; i < 4; i++ {
				// 坐标点移动
				nx := x + move[i]
				ny := y + move[i+1]
				// 判断是否越界 是否出现过
				if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] {
					visited[nx][ny] = true
					queue[r][0] = nx
					queue[r][1] = ny
					r++
				}
			}
		}
	}

	return level - 1
}
