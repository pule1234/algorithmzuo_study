package stack

var height [201]int
var stack [201]int
var r int

// 每次遍历一层，表示，将当前这一层作为柱状图的底，并使用heights数组记录
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	ans := 0
	for i := 0; i < m; i++ {
		height[i] = 0
	}
	for i := 0; i < n; i++ {
		//来到第i行，长方形一定要用i行做底
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		ans = max(ans, largestRectangleArea(m))
	}

	return ans
}

func largestRectangleArea(m int) int {
	r = 0
	ans := 0

	for i := 0; i < m; i++ {
		for r > 0 && height[stack[r-1]] >= height[i] {
			r--
			cur := stack[r]
			var left int
			if r == 0 {
				left = -1
			} else {
				left = stack[i-1]
			}
			ans = max(ans, height[cur]*(i-left-1))
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
		ans = max(ans, height[cur]*(m-left-1))
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
