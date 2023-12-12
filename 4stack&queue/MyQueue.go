package _stack_queue

type MyQueue struct {
	instack  []int
	outstack []int
}

func Constructor() MyQueue {
	return MyQueue{
		instack:  make([]int, 0),
		outstack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack)
}

func (q *MyQueue) in2out() {
	for len(q.instack) > 0 {
		q.outstack = append(q.outstack, q.instack[len(q.instack)-1])
		q.instack = q.instack[:len(q.instack)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outstack) == 0 {
		this.in2out()
	}
	x := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.outstack) == 0 {
		this.in2out()
	}
	return this.outstack[len(this.outstack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.instack) == 0 && len(this.outstack) == 0
}
