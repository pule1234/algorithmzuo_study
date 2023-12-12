package two

// 无法改成动态规划
func exist(board [][]byte, word string) bool {
	str := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if f(board, i, j, str, 0) {
				return true
			}
		}
	}
}

// 因为board会改其中的字符
// 用来标记哪些字符无法再用
// 带路径的递归无法改成动态规划或者说没必要
// 从(i,j)出发，来到w[k]，请问后续能不能把word走出来w[k...]
func f(board [][]byte, i, j int, str []byte, k int) bool {
	if k == len(str) { // k遍历完了str
		return true
	}

	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != str[k] {
		return false
	}

	// 不越界
	tmp := board[i][j]
	board[i][j] = 0
	ans := f(board, i-1, j, str, k+1) || f(board, i, j-1, str, k+1) || f(board, i+1, j, str, k+1) || f(board, i, j+1, str, k+1)
	board[i][j] = tmp
	return ans
}
