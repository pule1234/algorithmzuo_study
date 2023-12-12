package linknode

var start *ListNode
var end *ListNode

func sortList(head *ListNode) *ListNode {
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	// l1...r1 每组的左部分
	// l2...r2 每组的右部分
	// next 下一组的开头
	// lastTeamEnd 上一组的结尾
	var l1 *ListNode
	var l2 *ListNode
	var r1 *ListNode
	var r2 *ListNode
	var next *ListNode
	var lastteamend *ListNode
	for step := 1; step < n; step <<= 1 {
		// 第一组很特殊，因为要决定整个链表的头，所以单独处理
		l1 = head
		r1 = findEnd(l1, step)
		l2 = r1.Next
		r2 = findEnd(l2, step)
		next = r2.Next
		r1.Next = nil
		r2.Next = nil
		merge(l1, l2, r1, r2)
		head = start
		lastteamend = end
		for next != nil {
			l1 = next
			r1 = findEnd(l1, step)
			l2 = r1.Next
			if l2 == nil {
				lastteamend.Next = r1
				break
			}
			r2 = findEnd(l2, step)
			next = r2.Next
			r1.Next = nil
			r2.Next = nil
			merge(l1, l2, r1, r2)
			lastteamend.Next = start
			lastteamend = end
		}
	}
	return head

}

func findEnd(s *ListNode, k int) *ListNode {
	k--
	for k > 0 && s.Next != nil {
		s = s.Next
		k--
	}
	return s
}

// l1...r1 -> null : 有序的左部分
// l2...r2 -> null : 有序的右部分
// 整体merge在一起，保证有序
// 并且把全局变量start设置为整体的头，全局变量end设置为整体的尾
func merge(l1 *ListNode, l2 *ListNode, r1 *ListNode, r2 *ListNode) {
	var pre *ListNode
	if l1.Val <= l2.Val {
		start = l1
		pre = l1
		l1 = l1.Next
	} else {
		start = l2
		pre = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		pre.Next = l1
		end = r1
	} else {
		pre.Next = l2
		end = r2
	}
}
