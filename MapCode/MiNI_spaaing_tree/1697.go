package MiNI_spaaing_tree

import "sort"

const MAXN = 100001

var questions [MAXN][4]int

var father [MAXN]int

func build(n int) {
	for i := 0; i < n; i++ {
		father[i] = i
	}
}

func find(i int) int {
	if i != father[i] {
		father[i] = find(father[i])
	}
	return father[i]
}

func uion(x, y int) {
	father[find(x)] = find(y)
}

func isSameSet(x, y int) bool {
	return find(x) == find(y)
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	m := len(edgeList)
	k := len(queries)
	for i := 0; i < k; i++ {
		questions[i][0] = queries[i][0]
		questions[i][1] = queries[i][1]
		questions[i][2] = queries[i][2]
		questions[i][3] = i
	}

	sort.Slice(questions[:k], func(i, j int) bool {
		return questions[i][2] < questions[j][2]
	})

	build(n)
	ans := make([]bool, k)
	for i, j := 0, 0; i < k; i++ {
		// i : 问题编号
		// j : 边的编号
		for ; j < m && edgeList[j][2] < questions[i][2]; j++ {
			uion(edgeList[j][0], edgeList[j][1])
		}
		ans[questions[i][3]] = isSameSet(questions[i][0], questions[i][1])
	}

	return ans
}
