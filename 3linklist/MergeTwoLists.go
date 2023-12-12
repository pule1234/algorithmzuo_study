package _linklist

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func MergeTwoList(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return dummy.Next
}

func MergeTwoList2(list1, list2 *ListNode) *ListNode {
	var head *ListNode

	if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		head = list1
	} else {
		head = list2
	}

	pre := head
	cur1 := head.Next
	var cur2 *ListNode
	if head == list1 {
		cur2 = list2
	} else {
		cur2 = list1
	}

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}

	if cur1 == nil {
		pre.Next = cur2
	} else {
		pre.Next = cur1
	}

	return head
}
