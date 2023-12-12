package tree_dp

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	max       int
	min       int
	sum       int
	isBst     bool
	maxBstSum int
}

func maxSumBST(root *TreeNode) int {
	return f(root).maxBstSum
}

func f(x *TreeNode) Info {
	if x == nil {
		return Info{max: math.MinInt32, min: math.MaxInt32, sum: 0, isBst: true, maxBstSum: 0}
	}

	infol := f(x.Left)
	infor := f(x.Right)

	// 整合左右子树的信息
	max := Max(x.Val, Max(infol.max, infor.max))
	min := Min(x.Val, Min(infol.min, infor.min))
	sum := infol.sum + infor.sum + x.Val

	isBst := infol.isBst && infor.isBst && infol.max < x.Val && x.Val < infor.min
	maxbstsum := Max(infol.maxBstSum, infor.maxBstSum)
	if isBst {
		maxbstsum = Max(maxbstsum, sum)
	}

	return Info{max, min, sum, isBst, maxbstsum}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
