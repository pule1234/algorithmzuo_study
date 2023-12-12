package Dijkstra

const (
	MAXN = 31
	MAXM = 31
	MAXK = 6
)

var move = []int{-1, 0, 1, 0, -1}
var grid [MAXN][]byte

var visited [MAXN][MAXM][1 << MAXK]bool

// 0 : 行
// 1 : 列
// 2 : 收集钥匙的状态
var queue [MAXN * MAXM * (1 << MAXK)][3]int

var l, r, n, m, key int

func build(g []string) {
	l, r, key = 0, 0, 0
	n = len(g)
	m = len(g[0])
	for i := 0; i < n; i++ {
		grid[i] = []byte(g[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 为起点
			if grid[i][j] == '@' {
				queue[r][0] = i
				queue[r][1] = j
				queue[r][2] = 0
				r++
			}
			// 为钥匙
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				key |= (1 << (grid[i][j] - 'a'))
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for s := 0; s <= key; s++ {
				visited[i][j][s] = false
			}
		}
	}
}

func shortestPathAllKeys(grid []string) int {
	build(grid)
	level := 1
	for l < r {
		for k, size := 0, r-l; k < size; k++ {
			x := queue[l][0]
			y := queue[l][1]
			s := queue[l][2]
			l++
			for i := 0; i < 4; i++ {
				nx := x + move[i]
				ny := y + move[i+1]
				ns := s

				if nx < 0 || nx == n || ny < 0 || ny == m || grid[nx][ny] == '#' {
					//越界或者是墙
					continue
				}
				if grid[nx][ny] >= 'A' && grid[nx][ny] <= 'F' && ((ns & (1 << (grid[nx][ny] - 'A'))) == 0) {
					//遇到了锁但是没有对应的钥匙
					continue
				}
				if grid[nx][ny] >= 'a' && grid[nx][ny] <= 'f' {
					ns |= (1 << (grid[nx][ny] - 'a'))
				}

				if ns == key {
					// 找到了所有的钥匙
					return level
				}

				if !visited[nx][ny][ns] {
					visited[nx][ny][ns] = true
					queue[r][0] = nx
					queue[r][1] = ny
					queue[r][2] = ns
					r++
				}
			}
		}
		level++
	}
	return -1
}
