package stack

var height [50001]int
var stack [50001]int

func numSubmat(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])

	ans := 0
	for i := 0; i < m; i++ {
		height[i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		ans += countFromBottom(m)
	}

	return ans
}

func countFromBottom(m int) int {
	r := 0
	ans := 0

	for i := 0; i < m; i++ {
		for r > 0 && height[stack[r-1]] >= height[i] { // 出栈
			r--
			cur := stack[r]
			if height[cur] > height[i] {
				// 只有height[cur] > height[i]才结算
				// 如果是因为height[cur]==height[i]导致cur位置从栈中弹出
				// 那么不结算！等i位置弹出的时候再说！
				var left int
				if r == 0 {
					left = -1
				} else {
					left = stack[r-1]
				}
				length := i - left - 1
				var down int
				if left == -1 {
					down = 0
				} else {
					down = height[left]
				}
				bottom := max(down, height[i])
				ans += (height[cur] - bottom) * length * (length + 1) / 2
			}
		}
		stack[r] = i
		r++
	}

	for r > 0 {
		r--
		cur := stack[r]
		var left int
		if r == 0 {
			left = -1
		} else {
			left = stack[r-1]
		}

		length := m - left - 1
		var down int
		if left == -1 {
			down = 0
		} else {
			down = height[left]
		}
		ans += (height[cur] - down) * length * (length + 1) / 2
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
