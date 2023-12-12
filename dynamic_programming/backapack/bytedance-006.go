package main

import (
	"bufio"
	"fmt"
	"os"
)

// 夏季特惠
// 某公司游戏平台的夏季特惠开始了，你决定入手一些游戏
// 现在你一共有X元的预算，平台上所有的 n 个游戏均有折扣
// 标号为 i 的游戏的原价a_i元，现价只要b_i元
// 也就是说该游戏可以优惠 a_i - b_i，并且你购买该游戏能获得快乐值为w_i
// 由于优惠的存在，你可能做出一些冲动消费导致最终买游戏的总费用超过预算
// 只要满足 : 获得的总优惠金额不低于超过预算的总金额
// 那在心理上就不会觉得吃亏。
// 现在你希望在心理上不觉得吃亏的前提下，获得尽可能多的快乐值。
// 测试链接 : https://leetcode.cn/problems/tJau2o/

const MAXN = 501
const MAXX = 100001

var cost [MAXN]int
var val [MAXN]int64
var dp [MAXX]int64

var n, m, x int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		fmt.Sscanf(input, "%d %d", &n, &x)
		m = 1

		var ans int64
		var happy int64
		for i := 1; i <= n; i++ {
			scanner.Scan()
			input = scanner.Text()
			var pre, cur, well int
			fmt.Sscanf(input, "%d %d %d", &pre, &cur, &happy)
			well = pre - cur - cur

			if well >= 0 {
				x += well
				ans += happy
			} else {
				cost[m] = -well
				val[m] = happy
				m++
			}
		}
		ans += int64(computer1())
		fmt.Println(ans)
	}

}

func computer1() int {
	for i := 1; i <= n; i++ {
		for j := x; j >= cost[i]; j-- {
			dp[j] = int64(MAX(int(dp[j]), int(dp[j-cost[i]]+val[i])))
		}
	}

	return int(dp[x])
}

func MAX(a, b int) int {
	if a > b {
		return a
	}
	return b
}
