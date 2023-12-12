package Nested

// 含有嵌套的表达式求值 计算器
func calculate(str string) int {
	where = 0
	bys := []byte(str)
	return f(bys, 0)
}

var where int

// s[i....]开始计算，遇到字符串终止 或者 遇到)停止
// 返回 : 自己负责的这一段，计算的结果
// 返回之间，更新全局变量where，为了上游函数知道从哪继续！
func f(s []byte, i int) int {
	cur := 0
	numbers := make([]int, 0)
	ops := make([]byte, 0)
	for i < len(s) && s[i] != ')' { //在范围内，并且没有遇到右括号
		if s[i] >= '0' && s[i] <= '9' { // 遇到的是数字
			cur = cur*10 + int(s[i]-'0')
			i++
		} else if s[i] != '(' {
			//遇到了运算符
			push(numbers, ops, cur, s[i])
			i++
			cur = 0
		} else {
			// i (.....)
			// 遇到了左括号！
			cur = f(s, i+1)
			i = where + 1
		}
	}
	push(numbers, ops, cur, '+')
	where = i
	return computer(numbers, ops)
}

func push(numbers []int, ops []byte, cur int, op byte) {
	n := len(numbers)
	if n == 0 || ops[n-1] == '+' || ops[n-1] == '-' {
		numbers = append(numbers, cur)
		ops = append(ops, op)
	} else { // 遇到的是乘和除
		topnumber := numbers[n-1]
		topop := ops[n-1]
		if topop == '*' {
			numbers = numbers[:n-1]
			numbers = append(numbers, topnumber*cur)
		} else {
			numbers = numbers[:n-1]
			numbers = append(numbers, topnumber/cur)
		}
		ops = ops[:n-1]
		ops = append(ops, op)
	}
}

func computer(numbers []int, ops []byte) int {
	n := len(numbers)
	ans := numbers[0]
	for i := 1; i < n; i++ {
		if ops[i-1] == '+' {
			ans += numbers[i]
		} else {
			ans += -numbers[i]
		}
	}
	return ans
}
