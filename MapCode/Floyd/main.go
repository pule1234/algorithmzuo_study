package Floyd

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MAXN = 101
const MAXM = 10001

var path = [MAXM]int{}
var distance = [MAXN][MAXM]int{}
var n, m, ans int

// 初始时设置两点之间的最大值为无穷大，表示任何路都不存在
func build() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			distance[i][j] = math.MaxInt32
		}
	}
}

func floyd() {
	// 枚举每个跳板
	// 注意，跳板要最先枚举！跳板要最先枚举！跳板要最先枚举！    判断以该点为跳板，到别的点的距离是否会变小
	for bridge := 0; bridge < n; bridge++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// i -> .....bridge .... -> j
				// distance[i][j]能不能缩短
				// distance[i][j] = min ( distance[i][j] , distance[i][bridge] + distance[bridge][j])
				if distance[i][bridge] != math.MaxInt32 && distance[bridge][j] != math.MaxInt32 && distance[i][j] > distance[i][bridge]+distance[bridge][j] {
					distance[i][j] = distance[i][bridge] + distance[bridge][j]
				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		_, err := fmt.Fscanf(reader, "%d", n)
		if err != nil {
			break
		}
		fmt.Fscanf(reader, "%d", &m)

		for i := 0; i < m; i++ {
			fmt.Fscanf(reader, "%d", &path[i])
			path[i]--
		}
		build()

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Fscanf(reader, "%d", &distance[i][j])
			}
		}
		floyd()

		ans = 0
		for i := 1; i < m; i++ {
			ans += distance[path[i-1]][path[i]]
		}
		fmt.Fprintln(writer, ans)
	}
}
