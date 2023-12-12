package _linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleNode struct {
	Val  int
	Next *DoubleNode
	last *DoubleNode
}
