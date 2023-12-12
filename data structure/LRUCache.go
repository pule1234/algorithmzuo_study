package data_structure

import "container/list"

type LRUCache struct {
	size int
	hash map[int]*list.Element
	list *list.List
}

type entery struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*list.Element{}, list.New()}
}

func (this *LRUCache) Get(key int) int {
	node := this.hash[key]
	if node == nil {
		return -1
	}
	this.list.MoveToFront(node)
	return node.Value.(entery).value
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.hash[key]; node != nil {
		node.Value = entery{key: key, value: value}
		this.list.MoveToFront(node)
		return
	}
	this.hash[key] = this.list.PushFront(entery{key, value})
	if len(this.hash) > this.size {
		delete(this.hash, this.list.Remove(this.list.Back()).(entery).value)
	}
}
