package tree_dp

import "math"

// 最大BST子树
// 给定一个二叉树，找到其中最大的二叉搜索树（BST）子树，并返回该子树的大小
// 其中，最大指的是子树节点数最多的
// 二叉搜索树（BST）中的所有节点都具备以下属性：
// 左子树的值小于其父（根）节点的值
// 右子树的值大于其父（根）节点的值
// 注意：子树必须包含其所有后代
// 测试链接 : https://leetcode.cn/problems/largest-bst-subtree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Info 定义用于存储节点信息的结构
type Info struct {
	Max        int64
	Min        int64
	IsBST      bool
	MaxBSTSize int
}

// LargestBSTSubtree 是计算最大二叉搜索子树的函数
func LargestBSTSubtree(root *TreeNode) int {
	return f(root).MaxBSTSize
}

func f(x *TreeNode) Info {
	if x == nil {
		return Info{Max: math.MinInt64, Min: math.MaxInt64, IsBST: true, MaxBSTSize: 0}
	}

	infol := f(x.Left)
	infor := f(x.Right)

	//整合左右子树的信息
	max := int64(math.Max(float64(x.Val), math.Max(float64(infol.Max), float64(infor.Max)))) // 整颗树上的最大值
	min := int64(math.Min(float64(x.Val), math.Min(float64(infol.Min), float64(infor.Min)))) //整棵树上的最小值

	isBST := infol.IsBST && infor.IsBST && infol.Max < int64(x.Val) && infor.Min > int64(x.Val)
	var maxBSTSize int
	if isBST {
		maxBSTSize = infol.MaxBSTSize + infor.MaxBSTSize + 1
	} else {
		maxBSTSize = int(math.Max(float64(infol.MaxBSTSize), float64(infor.MaxBSTSize)))
	}

	return Info{max, min, isBST, maxBSTSize}
}
