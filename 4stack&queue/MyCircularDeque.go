package _stack_queue

import "fmt"

type MyCircularDeque struct {
	left, right, limit, size int
	queue                    []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{0, 0, k, 0, make([]int, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.left == 0 {
		this.queue[len(this.queue)-1] = value
		this.left = len(this.queue) - 1
	} else {
		if this.IsEmpty() {
			this.queue[this.left] = value
		} else {
			this.left--
			this.queue[this.left] = value
		}
	}
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.right == len(this.queue)-1 {
		this.right = 0
		this.queue[this.right] = value
	} else {
		if this.IsEmpty() {
			this.queue[this.right] = value
		} else {
			this.right++
			this.queue[this.right] = value
		}
	}
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	if this.left == len(this.queue)-1 {
		this.left = 0
	} else {
		this.left++
	}
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.right == 0 {
		this.right = this.limit - 1
	} else {
		this.right--
	}
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	fmt.Println(this.queue)
	fmt.Println("left = ", this.left)
	return this.queue[this.left]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	fmt.Println(this.queue)
	fmt.Println("this.right = ", this.right)
	return this.queue[this.right]
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularDeque) IsFull() bool {
	if this.size == this.limit {
		return true
	}
	return false
}
