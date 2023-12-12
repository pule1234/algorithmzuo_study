package floodfilling

// 找到所有被x围绕的区域
func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])

	for j := 0; j < m; j++ {
		if board[0][j] == 'O' { // 第一行 {}
			dfs(board, n, m, 0, j)
		}
		if board[n-1][j] == 'O' {
			dfs(board, n, m, n-1, j)
		}
	}

	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dfs(board, n, m, i, 0)
		}
		if board[i][m-1] == 'O' {
			dfs(board, n, m, i, m-1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, n, m, i, j int) {
	if i < 0 || i == n || j < 0 || j == m || board[i][j] != 'O' {
		return
	}

	//board[i][j] = 'O'
	board[i][j] = 'F'
	dfs(board, n, m, i+1, j)
	dfs(board, n, m, i-1, j)
	dfs(board, n, m, i, j+1)
	dfs(board, n, m, i, j-1)

}
