package floodfilling

// 洪水填充的做法
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	island := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				island++
				dfs(grid, n, m, i, j)
			}
		}
	}

	return island
}

func dfs(grid [][]byte, n, m, i, j int) {
	if i < 0 || i == n || j < 0 || j == m || grid[i][j] != '1' {
		return
	}

	// grid[i]【j】 == 1
	grid[i][j] = 0

	dfs(grid, n, m, i-1, j)
	dfs(grid, n, m, i+1, j)
	dfs(grid, n, m, i, j-1)
	dfs(grid, n, m, i, j+1)
}
