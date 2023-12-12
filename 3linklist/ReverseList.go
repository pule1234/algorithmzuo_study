package _linklist

// 翻转链表  单链表
func ReverseList(node *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode

	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre
}

// 双链表 翻转指针
func ReverseDoubleList(node *DoubleNode) *DoubleNode {
	var pre *DoubleNode
	var next *DoubleNode

	for node != nil {
		next = node.Next
		node.Next = pre
		node.last = next
		pre = node
		node = next
	}

	return pre
}
