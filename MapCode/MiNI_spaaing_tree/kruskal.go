package MiNI_spaaing_tree

import "sort"

const MAXN = 5001
const MAXM = 200001

var father [MAXN]int

var edges [MAXM][3]int

var n, m int

func build() {
	for i := 1; i <= n; i++ {
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
	} else {
		return false
	}
}

func kruskal() bool {
	sort.Slice(edges, func(x, y int) bool {
		return edges[x][2] < edges[y][2]
	})

	ans := 0
	edgesCnt := 0
	for _, edge := range edges {
		if uion(edge[0], edge[1]) {
			edgesCnt++
			ans += edge[2]
		}
	}

	return edgesCnt == n-1
}
