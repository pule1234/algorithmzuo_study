package dynamic_programming

import (
	"fmt"
	"strconv"
)

func hanoi(n int) {
	if n > 0 {
		f1(n, "左", "中", "右")
	}
}

func f1(i int, from, to, other string) {
	if i == 1 {
		fmt.Println("移动圆盘 1 从 " + from + " 到 " + to)
	} else {
		f1(i-1, from, other, to)
		fmt.Println("移动圆盘" + strconv.Itoa(i) + " 从 " + from + " 到 " + to)
		f1(i-1, other, to, from)
	}
}

func main() {
	hanoi(3)
}
