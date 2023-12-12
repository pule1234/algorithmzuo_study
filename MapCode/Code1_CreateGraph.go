package main

import (
	"fmt"
)

// 点的最大数量
const MAXN = 11

// 边的最大数量
// 只有链式前向星方式建图需要这个数量
// 注意如果无向图的最大数量是m条边，数量要准备m*2
// 因为一条无向边要加两条有向边
const MAXM = 21

// 邻接矩阵方式建图
var graph1 [MAXN][MAXM]int

// 邻接表的方式建图
var graph2 [][][]int

// 链式前向星方式建图 、
// 下标为点编号，职位头部边号
var head [MAXN]int

// 下标为边的编号，值为下一条边的编号
var next [MAXM]int

// 下标为边的编号，值为去往的点
var to [MAXM]int

// 如果边有权重， 那么需要这个数组
var weight [MAXM]int

var cnt int

func build(n int) {
	//邻接矩阵清空
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			graph1[i][j] = 0
		}
	}

	//邻接表清空和准备
	graph2 = make([][][]int, n+1)
	//下标需要支持1~n, 所以加入n+1个列表，0下标准备但是不用
	// 链式前向星清空
	cnt = 1
	for i := 1; i <= n; i++ {
		head[i] = 0
	}
}

// 链式前向星加边   addEdge(edge[0], edge[1], edge[2])
func addEdge(u, v, w int) {
	// u -> v , 边权重是w
	next[cnt] = head[u]
	to[cnt] = v
	weight[cnt] = w
	head[u] = cnt
	cnt++
}

// 三种方式建立有向图带权图
func directGraph(edges [][]int) {
	//邻接矩阵建图
	for _, edge := range edges {
		graph1[edge[0]][edge[1]] = edge[2]
	}

	for _, edge := range edges {
		graph2[edge[0]] = append(graph2[edge[0]], []int{edge[1], edge[2]})
	}

	//链式前向星建图
	for _, edge := range edges {
		addEdge(edge[0], edge[1], edge[2])
	}
}

// 三种方式建立无向图带权图
func undirectGraph(edges [][]int) {
	//邻接矩阵建图
	for _, edge := range edges {
		graph1[edge[0]][edge[1]] = edge[2]
		graph1[edge[1]][edge[0]] = edge[2]
	}

	//邻接表建图
	for _, edge := range edges {
		graph2[edge[0]] = append(graph2[edge[0]], []int{edge[1], edge[2]})
		graph2[edge[1]] = append(graph2[edge[1]], []int{edge[0], edge[2]})
	}

	//链式前向星
	for _, edge := range edges {
		addEdge(edge[0], edge[1], edge[2])
		addEdge(edge[1], edge[0], edge[2])
	}
}

func traversal(n int) {
	fmt.Println("邻接矩阵遍历:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(graph1[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("邻接表遍历:")
	for i := 1; i <= n; i++ {
		fmt.Println(i, "(邻居、边权) : ")
		for _, edge := range graph2[i] {
			fmt.Print("(", edge[0], ",", edge[1], ") ")
		}
		fmt.Println()
	}

	fmt.Println("链式前向星：")
	for i := 0; i <= n; i++ {
		fmt.Println(i, "(邻居、边权) : ")
		// 注意这个for循环，链式前向星的方式遍历
		for ei := head[i]; ei > 0; ei = next[ei] {
			fmt.Print("(", to[ei], ",", weight[ei], ")")
		}
		fmt.Println()
	}
}

func main() {
	n1 := 4
	edges1 := [][]int{{1, 3, 6}, {4, 3, 4}, {2, 4, 2}, {1, 2, 7}, {2, 3, 5}, {3, 1, 1}}
	build(n1)
	directGraph(edges1)
	traversal(n1)

	fmt.Println("++++++++++++++++++++++++++++++++++++++")

	n2 := 5
	edges2 := [][]int{{3, 5, 4}, {4, 1, 1}, {3, 4, 2}, {5, 2, 4}, {2, 3, 7}, {1, 5, 5}, {4, 2, 6}}
	build(n2)
	undirectGraph(edges2)
	traversal(n2)
}
