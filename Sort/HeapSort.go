package Sort

func sort(nums []int) {
	heapSort1(nums)
}

// i 位置的数，向上调整大根堆
// arr[i] = x，x是新来的！往上看，直到不比父亲大，或者来到0位置(顶)
func heapInsert(nums []int, i int) {
	for nums[i] > nums[(i-1)/2] {
		nums[i], nums[(i-1)/2] = nums[(i-1)/2], nums[i]
		i = (i - 1) / 2
	}
}

// i位置的数，改变了，又想维持大根堆结构， 若
// 向下调整大根堆
// 当前堆的大小为size
func heapify(nums []int, i, size int) {
	l := i*2 + 1 // 左孩子节点
	for l < size {
		best := 0
		if l+1 < size && nums[l] > nums[l+1] {
			best = l
		} else {
			best = l + 1
		}

		if nums[best] < nums[i] {
			best = i
		}

		if best == i {
			return
		}
		nums[best], nums[i] = nums[i], nums[best]
		i = best
		l = i*2 + 1
	}
}

// 从顶到底建立大根堆，O(n * logn)
// 依次弹出堆内最大值并排好序，O(n * logn)
// 整体时间复杂度O(n * logn)
func heapSort1(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		heapInsert(nums, i)
	}

	size := n
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapify(nums, 0, size)
	}
}

func heapSort2(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
	size := n
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapify(nums, 0, size)
	}
}
