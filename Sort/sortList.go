package Sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2

	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	}

	if temp2 != nil {
		temp.Next = temp2
	}

	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head

	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow

	return merge(sort(slow, mid), sort(mid, fast))
}
