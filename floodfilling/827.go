package floodfilling

func largestIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	id := 2

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dfs(grid, n, m, i, j, id)
				id++
			}
		}
	}

	sizes := make([]int, id)
	ans := 0
	visited := make([]bool, id)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > 1 {
				sizes[grid[i][j]]++
				ans = max(ans, sizes[grid[i][j]])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				visited = make([]bool, id)
				merge := 1
				merge += check(grid, visited, sizes, i-1, j)
				merge += check(grid, visited, sizes, i+1, j)
				merge += check(grid, visited, sizes, i, j-1)
				merge += check(grid, visited, sizes, i, j+1)
				ans = max(ans, merge)
			}
		}
	}

	return ans
}

func dfs(grid [][]int, n, m, i, j, id int) {
	if i < 0 || i == n || j < 0 || j == m || grid[i][j] != 1 {
		return
	}
	grid[i][j] = id
	dfs(grid, n, m, i-1, j, id)
	dfs(grid, n, m, i+1, j, id)
	dfs(grid, n, m, i, j-1, id)
	dfs(grid, n, m, i, j+1, id)
}

func check(grid [][]int, visited []bool, sizes []int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && !visited[grid[i][j]] {
		visited[grid[i][j]] = true
		return sizes[grid[i][j]]
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
