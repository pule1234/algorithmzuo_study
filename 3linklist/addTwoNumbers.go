package _linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var cur *ListNode

	carry := 0

	for l1 != nil || l2 != nil {
		var val int
		if l1 != nil && l2 != nil {
			val = l1.Val + l2.Val + carry
		} else if l1 == nil && l2 != nil {
			val = l2.Val + carry
		} else {
			val = l1.Val + carry
		}
		carry = val / 10
		val = val % 10

		if ans == nil {
			ans = &ListNode{val, nil}
			cur = ans
		} else {
			cur.Next = &ListNode{val, nil}
			cur = cur.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
	}

	return ans
}
