package linknode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	count := 0

	for a != nil {
		a = a.Next
		count++
	}

	for b != nil {
		b = b.Next
		count--
	}

	if a != b {
		return nil
	}

	if count >= 0 {
		a = headA
		b = headB
	} else {
		a = headB
		b = headA
	}

	for count > 0 {
		count--
		a = a.Next
	}

	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}
