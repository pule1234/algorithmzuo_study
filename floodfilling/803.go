package floodfilling

var n, m int
var grid [][]int

func hitBricks(g [][]int, hits [][]int) []int {
	grid = g
	n = len(g)
	m = len(g[0])
	ans := make([]int, len(hits))

	if n == 1 {
		return ans
	}

	for _, hit := range hits {
		grid[hit[0]][hit[1]]--
	}

	for i := 0; i < m; i++ {
		dfs(0, i)
	}

	for i := len(hits) - 1; i >= 0; i-- {
		row := hits[i][0]
		col := hits[i][1]
		grid[row][col]++
		if worth(row, col) {
			ans[i] = dfs(row, col) - 1
		}
	}
	return ans
}

// 集击中该坐标，应该掉落的数量
func dfs(i, j int) int {
	if i < 0 || i == n || j < 0 || j == m || grid[i][j] != 1 {
		return 0
	}

	grid[i][j] = 2
	return 1 + dfs(i+1, j) + dfs(i, j+1) + dfs(i-1, j) + dfs(i, j-1)
}

// 添加完之后的结果是否值得
func worth(i, j int) bool {
	return grid[i][j] == 1 &&
		(i == 0 ||
			(i > 0 && grid[i-1][j] == 2) ||
			(i < n-1 && grid[i+1][j] == 2) ||
			(j > 0 && grid[i][j-1] == 2) ||
			(j < m-1 && grid[i][j+1] == 2))
}
