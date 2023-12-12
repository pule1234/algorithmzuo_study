package linknode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	end := teamEnd(start, k)
	if end == nil {
		return head
	}
	// 第一组很特殊因为牵扯到换头的问题
	head = end
	reverse(start, end)
	// 反转之后start 变成上一个链的尾节点
	lastTeamEnd := start
	for lastTeamEnd.Next != nil {
		start = lastTeamEnd.Next
		end = teamEnd(start, k)
		if end == nil {
			return head
		}
		reverse(start, end)
		lastTeamEnd.Next = end
		lastTeamEnd = start
	}

	return head
}

// 找到当前组的结束节点
func teamEnd(s *ListNode, k int) *ListNode {
	k--
	for k > 0 && s != nil {
		k--
		s = s.Next
	}
	return s
}

// 将当前组的链表翻转
func reverse(s, e *ListNode) {
	e = e.Next
	var pre *ListNode = nil
	cur := s
	var next *ListNode = nil
	for cur != e {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	s.Next = e
}
