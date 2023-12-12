package linknode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 目前中点就是slow
	pre := slow
	cur := pre.Next
	var next *ListNode
	pre.Next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 上面的过程已经把链表调整成从左右两侧往中间指
	ans := true
	left := head
	right := pre
	for left != nil && right != nil {
		if left.Val != right.Val {
			ans = false
			return ans
		}
		left = left.Next
		right = right.Next
	}

	//，把链表调整回原来的样子再返回判断结果
	cur = pre.Next
	pre.Next = nil
	next = nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return ans
}
