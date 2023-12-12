package stack

var mod = 1000000007
var maxn = 30001

// 找到左边比他小的最近的值，找到右边离他最近比他小的值，，  中间这个区域的所有包含他的子集的最小值就是他自己
func sumSubarrayMins(arr []int) int {
	stack := make([]int, maxn)
	r := 0
	ans := 0

	for i := 0; i < len(arr); i++ {
		for r > 0 && arr[stack[r-1]] >= arr[i] {
			r--
			cur := stack[r]
			left := 0
			if r == 0 {
				left = -1
			} else {
				left = stack[r-1]
			}
			ans = (ans + (cur-left)*(i-cur)*arr[cur]) % mod
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
		ans = (ans + (cur-left)*(len(arr)-cur)*arr[cur]) % mod
	}

	return ans
}
