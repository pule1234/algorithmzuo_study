package MiNI_spaaing_tree

import "sort"

// 水资源分配优化
// 村里面一共有 n 栋房子。我们希望通过建造水井和铺设管道来为所有房子供水。
// 对于每个房子 i，我们有两种可选的供水方案：一种是直接在房子内建造水井
// 成本为 wells[i - 1] （注意 -1 ，因为 索引从0开始 ）
// 另一种是从另一口井铺设管道引水，数组 pipes 给出了在房子间铺设管道的成本，
// 其中每个 pipes[j] = [house1j, house2j, costj]
// 代表用管道将 house1j 和 house2j连接在一起的成本。连接是双向的。
// 请返回 为所有房子都供水的最低总成本
// 测试链接 : https://leetcode.cn/problems/optimize-water-distribution-in-a-village/
const MAXN = 10010

var edges [MAXN][3]int

var cnt int

var father [MAXN]int

func build(n int) {
	for i := 0; i <= n; i++ {
		father[i] = i
	}
}

func find(i int) int {
	if i != father[i] {
		father[i] = find(father[i])
	}

	return father[i]
}

func uion(x, y int) bool {
	fx := find(x)
	fy := find(y)

	if fx != fy {
		father[fx] = fy
		return true
	}

	return false
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	build(n)
	for i := 0; i < n; i++ {
		edges[cnt][0] = 0
		edges[cnt][1] = i + 1
		edges[cnt][2] = wells[i]
		cnt++
	}

	for i := 0; i < len(pipes); i++ {
		edges[cnt][0] = pipes[i][0]
		edges[cnt][1] = pipes[i][1]
		edges[cnt][2] = pipes[i][2]
	}

	sort.Slice(edges, func(x, y int) bool {
		return edges[x][2] < edges[y][2]
	})

	ans := 0
	for i := 0; i < cnt; i++ {
		if uion(edges[i][0], edges[i][1]) {
			ans += edges[i][2]
		}
	}

	return ans

}
