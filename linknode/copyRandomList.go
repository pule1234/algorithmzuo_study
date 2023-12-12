package linknode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	node := head
	// 将该链表上的拷贝节点放在该节点的后面
	for node != nil {
		copynode := &Node{Val: node.Val}
		copynode.Next = node.Next
		node.Next = copynode
		node = copynode.Next
	}

	node = head
	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		} else {
			node.Next.Random = nil
		}
		node = node.Next.Next
	}

	node = head
	newhead := head.Next
	copynode := newhead

	for node != nil {
		if node != head {
			copynode.Next = node.Next
			copynode = copynode.Next
		}
		node.Next = node.Next.Next
		node = node.Next
	}

	return newhead
}
