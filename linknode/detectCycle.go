package linknode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
