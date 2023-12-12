package tree_dp

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Info struct {
	cnt  int // 节点数
	sum  int // 硬币数
	move int // 移动的次数
}

func distributeCoins(root *TreeNode) int {
	return f(root).move
}

func f(x *TreeNode) Info {
	if x == nil {
		return Info{0, 0, 0}
	}

	infol := f(x.Left)
	infor := f(x.Right)
	cnt := infol.cnt + infor.cnt + 1
	sum := infol.sum + infor.sum + x.Val
	move := infol.move + infor.move + int(math.Abs(float64(infol.cnt-infol.sum))+math.Abs(float64(infor.cnt-infor.sum)))

	return Info{cnt, sum, move}
}
