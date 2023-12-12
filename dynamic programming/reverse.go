package dynamic_programming

// 用递归函数逆序栈
func reverse(stack []int) {
	if len(stack) == 0 {
		return
	}
	num := bottomOut(stack)
	reverse(stack)
	stack = append(stack, num)
}

// 栈底元素移除掉，上面的元素盖下来
// 返回移除掉的栈底元素
func bottomOut(stack []int) int {
	ans := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if len(stack) == 0 {
		return ans
	} else {
		last := bottomOut(stack)
		stack = append(stack, ans)
		return last
	}
}
